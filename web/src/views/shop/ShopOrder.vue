<template>
  <div class="mx-auto max-w-lg px-4 pt-8 text-center md:pt-16">
    <p class="text-5xl">🛵</p>
    <p class="mt-4 text-xs font-extrabold tracking-wide text-[#141413] uppercase">
      {{ t('orderLabel', { id: order.id }) }}
    </p>
    <h1 class="mt-1 text-2xl font-extrabold">{{ t('onTheWay') }}</h1>
    <p class="mt-2 text-sm font-bold text-[#5e5d59]">
      {{ t('orderLine', { window: order.window, address: order.address }) }}
    </p>

    <ul class="mt-6 space-y-3 text-start">
      <li v-for="item in order.items" :key="item.sku" class="flex items-center gap-3">
        <span class="text-2xl">{{ item.emoji }}</span>
        <span class="flex-1 font-bold">{{ item.qty }} × {{ itemName(item) }}</span>
        <span class="font-extrabold">{{ euros(item.cents * item.qty) }}</span>
      </li>
    </ul>

    <p class="mt-6 text-start text-lg font-extrabold">
      {{ t('totalAmount', { amount: euros(order.total) }) }}
    </p>
    <p class="mt-1 text-start text-xs font-bold text-[#87867f]">
      {{ t('savedDevice') }}
    </p>

    <router-link
      :to="{ name: 'ShopHome' }"
      class="mt-6 flex w-full items-center justify-center rounded-full bg-[#141413] py-3.5 text-base font-extrabold text-[#faf9f5]"
    >
      {{ t('backToStore') }}
    </router-link>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useShopCart } from '@/shop/cart'
import { euros, productBySku } from '@/shop/catalog'
import { localizeProduct, t } from '@/shop/locale'

const route = useRoute()
const cart = useShopCart()
const order = computed(() => cart.orders.find((item) => item.id === route.params.id))

function itemName(item) {
  const product = productBySku(item.sku)
  return product ? localizeProduct(product).name : item.name
}
</script>
