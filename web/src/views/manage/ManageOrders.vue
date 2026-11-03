<template>
  <div>
    <header class="page-header flex flex-wrap items-end justify-between gap-4">
      <div>
        <h1 class="page-title">{{ $t('manage.orders.title') }}</h1>
        <p class="page-subtitle">{{ $t('manage.orders.subtitle') }}</p>
      </div>
      <select v-model="statusFilter" class="input-field sm:max-w-xs">
        <option value="">{{ $t('manage.orders.all_statuses') }}</option>
        <option v-for="status in ORDER_STATUSES" :key="status" :value="status">
          {{ $t(`manage.status.${status}`) }}
        </option>
      </select>
    </header>

    <div class="section !p-0 overflow-x-auto">
      <table class="min-w-full divide-y divide-theme-border text-sm">
        <thead class="bg-theme-hover text-left text-xs uppercase text-theme-textLight">
          <tr>
            <th class="px-4 py-3">{{ $t('manage.orders.order') }}</th>
            <th class="px-4 py-3">{{ $t('manage.orders.customer') }}</th>
            <th class="px-4 py-3">{{ $t('manage.orders.status') }}</th>
            <th class="px-4 py-3 text-right">{{ $t('manage.orders.items') }}</th>
            <th class="px-4 py-3 text-right">{{ $t('manage.orders.total') }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-theme-border">
          <tr v-for="order in filtered" :key="order.id">
            <td class="px-4 py-3">
              <router-link
                :to="{ name: 'ManageOrder', params: { id: order.id } }"
                class="font-medium hover:underline"
              >
                {{ order.id }}
              </router-link>
              <p class="text-xs text-theme-textLight">{{ formatDate(order.placedAt) }}</p>
            </td>
            <td class="px-4 py-3">{{ order.name }}</td>
            <td class="px-4 py-3">{{ $t(`manage.status.${orderStatus(order)}`) }}</td>
            <td class="px-4 py-3 text-right">{{ itemCount(order) }}</td>
            <td class="px-4 py-3 text-right font-semibold">{{ euros(order.total) }}</td>
          </tr>
        </tbody>
      </table>
      <p v-if="!filtered.length" class="py-8 text-center text-sm text-theme-textLight">
        {{ $t('manage.orders.empty') }}
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { euros } from '@/shop/catalog'
import { ORDER_STATUSES, orderStatus, useShopCart } from '@/shop/cart'

const cart = useShopCart()
const statusFilter = ref('')

const filtered = computed(() =>
  cart.orders.filter((order) => !statusFilter.value || orderStatus(order) === statusFilter.value)
)

function itemCount(order) {
  return order.items.reduce((sum, item) => sum + item.qty, 0)
}

function formatDate(value) {
  return new Date(value).toLocaleString()
}
</script>
