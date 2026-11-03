<template>
  <div class="px-4 pt-5 md:px-6 md:pt-8">
    <h1 class="text-2xl font-extrabold">{{ t('basket') }}</h1>

    <div v-if="cart.detailed.length === 0" class="px-6 pt-16 text-center">
      <p class="text-5xl">🧺</p>
      <p class="mt-4 text-lg font-extrabold">{{ t('emptyTitle') }}</p>
      <p class="mt-1 text-sm font-bold text-[#8a8a8a]">
        {{ t('emptyBody') }}
      </p>
      <router-link
        :to="{ name: 'ShopHome' }"
        class="mt-6 inline-flex rounded-full bg-[#ff2f78] px-5 py-3 text-sm font-extrabold text-white"
      >
        {{ t('startShopping') }}
      </router-link>
    </div>

    <div v-else class="lg:grid lg:grid-cols-[minmax(0,1fr)_20rem] lg:items-start lg:gap-10">
      <div>
        <div
          class="mt-4 rounded-2xl px-4 py-3 text-sm font-extrabold"
          :class="remaining === 0 ? 'bg-[#e8f8ee] text-[#157a45]' : 'bg-[#fff4e8] text-[#9a4d00]'"
        >
          <template v-if="remaining === 0">{{ t('freeUnlocked') }}</template>
          <template v-else>{{ t('addForFree', { amount: euros(remaining) }) }}</template>
          <div class="mt-2 h-1.5 overflow-hidden rounded-full bg-white/70">
            <div class="h-full rounded-full bg-current" :style="{ width: `${progress}%` }" />
          </div>
        </div>

        <ul class="mt-4 divide-y divide-[#f2f2f2]">
          <li v-for="line in cart.detailed" :key="line.sku" class="flex items-center gap-3 py-3">
            <router-link
              :to="{ name: 'ShopProduct', params: { sku: line.sku } }"
              class="flex h-16 w-16 shrink-0 items-center justify-center rounded-2xl text-3xl"
              :style="{ background: line.product.tint }"
            >
              {{ line.product.emoji }}
            </router-link>
            <div class="min-w-0 flex-1">
              <p class="truncate font-extrabold">{{ localizeProduct(line.product).name }}</p>
              <p class="text-xs font-bold text-[#8a8a8a]">
                {{ localizeProduct(line.product).pack }}
              </p>
              <p class="mt-1 text-sm font-extrabold">{{ euros(line.product.cents * line.qty) }}</p>
            </div>
            <QtyControl :sku="line.sku" />
          </li>
        </ul>
      </div>

      <aside
        class="mt-4 lg:sticky lg:top-24 lg:mt-0 lg:rounded-3xl lg:border lg:border-[#f0f0f0] lg:p-5"
      >
        <dl class="space-y-2 text-sm font-bold lg:mt-0">
          <div class="flex justify-between">
            <dt class="text-[#6d6d6d]">{{ t('subtotal') }}</dt>
            <dd>{{ euros(cart.subtotal) }}</dd>
          </div>
          <div class="flex justify-between">
            <dt class="text-[#6d6d6d]">{{ t('delivery') }}</dt>
            <dd>{{ cart.delivery === 0 ? t('free') : euros(cart.delivery) }}</dd>
          </div>
          <div class="flex justify-between pt-2 text-base font-extrabold">
            <dt>{{ t('total') }}</dt>
            <dd>{{ euros(cart.total) }}</dd>
          </div>
        </dl>

        <router-link
          :to="{ name: 'ShopCheckout' }"
          class="mt-5 flex w-full items-center justify-center rounded-full bg-[#ff2f78] py-3.5 text-base font-extrabold text-white"
        >
          {{ t('checkoutCta', { amount: euros(cart.total) }) }}
        </router-link>
      </aside>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import QtyControl from '@/comp/shop/QtyControl.vue'
import { useShopCart } from '@/shop/cart'
import { euros, FREE_DELIVERY_OVER } from '@/shop/catalog'
import { localizeProduct, t } from '@/shop/locale'

const cart = useShopCart()
const remaining = computed(() => Math.max(0, FREE_DELIVERY_OVER - cart.subtotal))
const progress = computed(() => Math.min(100, (cart.subtotal / FREE_DELIVERY_OVER) * 100))
</script>
