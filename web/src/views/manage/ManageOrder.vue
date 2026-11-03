<template>
  <div>
    <header class="page-header">
      <router-link :to="{ name: 'ManageOrders' }" class="text-sm text-theme-textLight hover:text-theme-text">
        ← {{ $t('manage.orders.back') }}
      </router-link>
      <h1 class="page-title mt-2">{{ $t('manage.orders.detail_title', { id: order.id }) }}</h1>
      <p class="page-subtitle">{{ formatDate(order.placedAt) }}</p>
    </header>

    <div class="grid gap-6 lg:grid-cols-[1fr_20rem]">
      <div class="section !p-0 overflow-hidden">
        <ul class="divide-y divide-theme-border">
          <li v-for="item in order.items" :key="item.sku" class="flex items-center gap-3 px-6 py-4">
            <span class="text-2xl">{{ item.emoji }}</span>
            <div class="flex-1">
              <p class="font-medium">{{ item.name }}</p>
              <p class="text-xs text-theme-textLight">{{ item.pack }} · {{ item.qty }} × {{ euros(item.cents) }}</p>
            </div>
            <span class="font-semibold">{{ euros(item.cents * item.qty) }}</span>
          </li>
        </ul>
        <dl class="space-y-2 border-t border-theme-border px-6 py-4 text-sm">
          <div class="flex justify-between">
            <dt class="text-theme-textLight">{{ $t('manage.orders.subtotal') }}</dt>
            <dd>{{ euros(order.subtotal) }}</dd>
          </div>
          <div class="flex justify-between">
            <dt class="text-theme-textLight">{{ $t('manage.orders.delivery') }}</dt>
            <dd>{{ order.delivery ? euros(order.delivery) : $t('manage.orders.free') }}</dd>
          </div>
          <div class="flex justify-between text-base font-semibold">
            <dt>{{ $t('manage.orders.total') }}</dt>
            <dd>{{ euros(order.total) }}</dd>
          </div>
        </dl>
      </div>

      <div class="space-y-6">
        <div class="section">
          <label class="form-label" for="order-status">{{ $t('manage.orders.status') }}</label>
          <select
            id="order-status"
            class="input-field"
            :value="orderStatus(order)"
            @change="cart.setStatus(order.id, $event.target.value)"
          >
            <option v-for="status in ORDER_STATUSES" :key="status" :value="status">
              {{ $t(`manage.status.${status}`) }}
            </option>
          </select>
        </div>

        <div class="section space-y-4 text-sm">
          <div>
            <p class="text-xs uppercase text-theme-textLight">{{ $t('manage.orders.customer') }}</p>
            <p class="font-medium">{{ order.name }}</p>
            <p>{{ order.phone }}</p>
          </div>
          <div>
            <p class="text-xs uppercase text-theme-textLight">{{ $t('manage.orders.address') }}</p>
            <p>{{ order.address }}</p>
          </div>
          <div>
            <p class="text-xs uppercase text-theme-textLight">{{ $t('manage.orders.window') }}</p>
            <p>{{ order.window }}</p>
          </div>
          <div>
            <p class="text-xs uppercase text-theme-textLight">{{ $t('manage.orders.note') }}</p>
            <p>{{ order.note || $t('manage.orders.no_note') }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { euros } from '@/shop/catalog'
import { ORDER_STATUSES, orderStatus, useShopCart } from '@/shop/cart'

const route = useRoute()
const cart = useShopCart()
const order = computed(() => cart.orders.find((item) => item.id === route.params.id))

function formatDate(value) {
  return new Date(value).toLocaleString()
}
</script>
