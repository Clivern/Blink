<template>
  <div>
    <header class="page-header">
      <h1 class="page-title">{{ $t('manage.customers.title') }}</h1>
      <p class="page-subtitle">{{ $t('manage.customers.subtitle') }}</p>
    </header>

    <div class="section !p-0 overflow-x-auto">
      <table class="min-w-full divide-y divide-theme-border text-sm">
        <thead class="bg-theme-hover text-left text-xs uppercase text-theme-textLight">
          <tr>
            <th class="px-4 py-3">{{ $t('manage.customers.name') }}</th>
            <th class="px-4 py-3">{{ $t('manage.customers.email') }}</th>
            <th class="px-4 py-3">{{ $t('manage.customers.phone') }}</th>
            <th class="px-4 py-3 text-right">{{ $t('manage.customers.orders') }}</th>
            <th class="px-4 py-3 text-right">{{ $t('manage.customers.spent') }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-theme-border">
          <tr v-for="customer in customers" :key="customer.id">
            <td class="px-4 py-3 font-medium">{{ customer.name }}</td>
            <td class="px-4 py-3">{{ customer.email || '—' }}</td>
            <td class="px-4 py-3">{{ customer.phone || '—' }}</td>
            <td class="px-4 py-3 text-right">{{ customer.orderCount }}</td>
            <td class="px-4 py-3 text-right font-semibold">{{ euros(customer.spent) }}</td>
          </tr>
        </tbody>
      </table>
      <p v-if="!customers.length" class="py-8 text-center text-sm text-theme-textLight">
        {{ $t('manage.customers.empty') }}
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { euros } from '@/shop/catalog'
import { useShopCart } from '@/shop/cart'
import { useShopSession } from '@/shop/session'

const cart = useShopCart()
const session = useShopSession()

const customers = computed(() =>
  session.accounts.map((account) => {
    const orders = cart.orders.filter((order) => order.userId === account.id)

    return {
      ...account,
      orderCount: orders.length,
      spent: orders.reduce((sum, order) => sum + order.total, 0),
    }
  })
)
</script>
