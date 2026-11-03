import { defineStore } from 'pinia'
import { productBySku, DELIVERY_CENTS, FREE_DELIVERY_OVER } from '@/shop/catalog'
import { loadJSON } from '@/utils/storage'

const CART_KEY = 'blink.shop.cart'
const ORDER_KEY = 'blink.shop.orders'
const ADDRESS_KEY = 'blink.shop.address'

export const ORDER_STATUSES = ['placed', 'packing', 'delivering', 'delivered']

export function orderStatus(order) {
  return order.status ?? 'placed'
}

export const useShopCart = defineStore('shopCart', {
  state: () => ({
    lines: loadJSON(CART_KEY, []),
    orders: loadJSON(ORDER_KEY, []),
    address: loadJSON(ADDRESS_KEY, 'Oranienstraße 12, Kreuzberg'),
  }),
  getters: {
    count(state) {
      return state.lines.reduce((total, line) => total + line.qty, 0)
    },
    detailed(state) {
      return state.lines
        .map((line) => ({ ...line, product: productBySku(line.sku) }))
        .filter((line) => line.product)
    },
    subtotal() {
      return this.detailed.reduce((total, line) => total + line.qty * line.product.cents, 0)
    },
    delivery() {
      if (this.subtotal === 0 || this.subtotal >= FREE_DELIVERY_OVER) return 0
      return DELIVERY_CENTS
    },
    total() {
      return this.subtotal + this.delivery
    },
  },
  actions: {
    persist() {
      localStorage.setItem(CART_KEY, JSON.stringify(this.lines))
    },
    add(sku) {
      const line = this.lines.find((item) => item.sku === sku)
      if (line) line.qty += 1
      else this.lines.push({ sku, qty: 1 })
      this.persist()
    },
    setQty(sku, qty) {
      if (qty <= 0) {
        this.lines = this.lines.filter((item) => item.sku !== sku)
      } else {
        const line = this.lines.find((item) => item.sku === sku)
        if (line) line.qty = qty
        else this.lines.push({ sku, qty })
      }
      this.persist()
    },
    setAddress(address) {
      this.address = address
      localStorage.setItem(ADDRESS_KEY, JSON.stringify(address))
    },
    place(details) {
      const order = {
        id: `B${Date.now().toString(36).toUpperCase()}`,
        placedAt: new Date().toISOString(),
        address: this.address,
        items: this.detailed.map((line) => ({
          sku: line.sku,
          name: line.product.name,
          pack: line.product.pack,
          emoji: line.product.emoji,
          qty: line.qty,
          cents: line.product.cents,
        })),
        subtotal: this.subtotal,
        delivery: this.delivery,
        total: this.total,
        eta: '10–15 min',
        status: 'placed',
        ...details,
      }
      this.orders.unshift(order)
      localStorage.setItem(ORDER_KEY, JSON.stringify(this.orders))
      this.lines = []
      this.persist()
      return order
    },
    setStatus(id, status) {
      this.orders.find((order) => order.id === id).status = status
      localStorage.setItem(ORDER_KEY, JSON.stringify(this.orders))
    },
  },
})
