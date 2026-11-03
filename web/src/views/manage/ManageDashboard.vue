<template>
  <div>
    <header class="mb-8">
      <h1 class="text-2xl sm:text-3xl font-semibold text-theme-text tracking-tight">
        {{ $t('manage.dashboard.title') }}
      </h1>
      <p class="mt-1 text-sm text-theme-textLight">{{ $t('manage.dashboard.subtitle') }}</p>
    </header>

    <section class="mb-8 grid grid-cols-2 gap-4 lg:grid-cols-3">
      <div v-for="card in cards" :key="card.label" class="stat-card">
        <p class="stat-label">{{ card.label }}</p>
        <p class="stat-value">{{ card.value }}</p>
      </div>
    </section>

    <section class="mb-8 grid grid-cols-1 gap-6 lg:grid-cols-2">
      <div class="section overflow-hidden !p-0">
        <div class="border-b border-theme-border px-6 py-4">
          <h2 class="text-lg font-semibold text-theme-text">{{ $t('manage.dashboard.attention_title') }}</h2>
          <p class="mt-0.5 text-sm text-theme-textLight">{{ $t('manage.dashboard.attention_subtitle') }}</p>
        </div>

        <ul class="divide-y divide-theme-border max-h-[28rem] overflow-y-auto">
          <li v-for="item in attentionItems" :key="item.id" class="flex gap-4 px-6 py-4">
            <span
              class="mt-1.5 h-2 w-2 shrink-0 rounded-full"
              :class="item.level === 'warn' ? 'bg-amber-500' : 'bg-primary-500'"
              aria-hidden="true"
            />
            <div class="min-w-0 flex-1">
              <p class="text-sm font-medium text-theme-text">{{ item.title }}</p>
              <p class="mt-0.5 text-sm text-theme-textLight leading-relaxed">{{ item.description }}</p>
              <router-link
                v-if="item.to"
                :to="item.to"
                class="mt-2 inline-flex text-sm font-medium text-primary-800 hover:underline"
              >
                {{ item.action }}
              </router-link>
            </div>
          </li>
        </ul>
      </div>

      <div class="section overflow-hidden !p-0">
        <div class="flex items-center justify-between border-b border-theme-border px-6 py-4">
          <div>
            <h2 class="text-lg font-semibold text-theme-text">{{ $t('manage.dashboard.recent_orders') }}</h2>
            <p class="mt-0.5 text-sm text-theme-textLight">{{ $t('manage.dashboard.recent_subtitle') }}</p>
          </div>
          <router-link :to="{ name: 'ManageOrders' }" class="text-sm font-medium text-primary-800 hover:underline">
            {{ $t('manage.dashboard.view_all') }}
          </router-link>
        </div>

        <ul class="divide-y divide-theme-border">
          <li v-for="order in recentOrders" :key="order.id" class="flex items-center gap-4 px-6 py-3 text-sm">
            <div class="min-w-0 flex-1">
              <router-link
                :to="{ name: 'ManageOrder', params: { id: order.id } }"
                class="font-medium text-theme-text hover:underline"
              >
                {{ order.id }}
              </router-link>
              <p class="text-xs text-theme-textLight truncate">{{ order.name }}</p>
            </div>
            <span class="text-theme-textLight">{{ $t(`manage.status.${orderStatus(order)}`) }}</span>
            <span class="font-semibold">{{ euros(order.total) }}</span>
          </li>
        </ul>
        <p v-if="!cart.orders.length" class="px-6 py-8 text-center text-sm text-theme-textLight">
          {{ $t('manage.dashboard.no_orders') }}
        </p>
      </div>

      <div class="section overflow-hidden !p-0">
        <div class="border-b border-theme-border px-6 py-4">
          <h2 class="text-lg font-semibold text-theme-text">{{ $t('manage.dashboard.orders_by_status') }}</h2>
        </div>
        <ul class="divide-y divide-theme-border">
          <li v-for="item in statusCounts" :key="item.status" class="flex justify-between px-6 py-3 text-sm">
            <span>{{ $t(`manage.status.${item.status}`) }}</span>
            <span class="font-semibold">{{ item.count }}</span>
          </li>
        </ul>
      </div>

      <div class="section overflow-hidden !p-0">
        <div class="border-b border-theme-border px-6 py-4">
          <h2 class="text-lg font-semibold text-theme-text">{{ $t('manage.dashboard.top_products') }}</h2>
        </div>
        <ul class="divide-y divide-theme-border">
          <li v-for="item in topProducts" :key="item.sku" class="flex items-center gap-3 px-6 py-3 text-sm">
            <span class="text-xl">{{ item.emoji }}</span>
            <span class="flex-1">{{ item.name }}</span>
            <span class="text-theme-textLight">{{ $t('manage.dashboard.units_sold', { count: item.qty }) }}</span>
          </li>
        </ul>
      </div>
    </section>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { categories, euros, products } from '@/shop/catalog'
import { ORDER_STATUSES, orderStatus, useShopCart } from '@/shop/cart'
import { useShopSession } from '@/shop/session'

const { t } = useI18n()
const cart = useShopCart()
const session = useShopSession()

const revenue = computed(() => cart.orders.reduce((sum, order) => sum + order.total, 0))
const averageOrder = computed(() => (cart.orders.length ? revenue.value / cart.orders.length : 0))

const cards = computed(() => [
  { label: t('manage.dashboard.revenue'), value: euros(revenue.value) },
  { label: t('manage.dashboard.orders'), value: cart.orders.length },
  { label: t('manage.dashboard.average_order'), value: euros(averageOrder.value) },
  { label: t('manage.dashboard.customers'), value: session.accounts.length },
  { label: t('manage.dashboard.products'), value: products.length },
  { label: t('manage.dashboard.categories'), value: categories.length },
])

const statusCounts = computed(() =>
  ORDER_STATUSES.map((status) => ({
    status,
    count: cart.orders.filter((order) => orderStatus(order) === status).length,
  }))
)

const recentOrders = computed(() => cart.orders.slice(0, 5))

const emptyCategories = computed(() =>
  categories
    .flatMap((category) => category.children)
    .filter((child) => !products.some((product) => product.category === child.slug))
)

const attentionItems = computed(() => {
  const items = []
  const toPack = statusCounts.value.find((item) => item.status === 'placed').count

  if (toPack > 0) {
    items.push({
      id: 'to-pack',
      level: 'warn',
      title: t('manage.dashboard.attention_to_pack_title', { count: toPack }),
      description: t('manage.dashboard.attention_to_pack_desc'),
      action: t('manage.dashboard.attention_to_pack_action'),
      to: { name: 'ManageOrders' },
    })
  }

  if (emptyCategories.value.length > 0) {
    items.push({
      id: 'empty-categories',
      level: 'info',
      title: t('manage.dashboard.attention_empty_categories_title', { count: emptyCategories.value.length }),
      description: t('manage.dashboard.attention_empty_categories_desc'),
      action: t('manage.dashboard.attention_empty_categories_action'),
      to: { name: 'ManageCategories' },
    })
  }

  if (!cart.orders.length) {
    items.push({
      id: 'no-orders',
      level: 'info',
      title: t('manage.dashboard.attention_no_orders_title'),
      description: t('manage.dashboard.attention_no_orders_desc'),
      action: t('manage.dashboard.attention_no_orders_action'),
      to: { name: 'ManageProducts' },
    })
  }

  if (!items.length) {
    items.push({
      id: 'clear',
      level: 'info',
      title: t('manage.dashboard.attention_clear_title'),
      description: t('manage.dashboard.attention_clear_desc'),
    })
  }

  return items
})

const topProducts = computed(() => {
  const sold = new Map()

  for (const order of cart.orders) {
    for (const item of order.items) {
      const line = sold.get(item.sku) || { sku: item.sku, name: item.name, emoji: item.emoji, qty: 0 }
      line.qty += item.qty
      sold.set(item.sku, line)
    }
  }

  return [...sold.values()].sort((a, b) => b.qty - a.qty).slice(0, 5)
})
</script>
