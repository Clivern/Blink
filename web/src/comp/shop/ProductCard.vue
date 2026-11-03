<template>
  <article class="flex flex-col">
    <router-link
      :to="{ name: 'ShopProduct', params: { sku: product.sku } }"
      class="relative flex aspect-square items-center justify-center rounded-2xl"
      :style="{ background: product.tint }"
    >
      <span
        v-if="product.wasCents"
        class="absolute top-2 start-2 rounded-full bg-[#ff2f78] px-2 py-0.5 text-[11px] font-extrabold text-white"
      >
        {{ t('deal') }}
      </span>
      <span class="text-6xl" aria-hidden="true">{{ product.emoji }}</span>
    </router-link>
    <div class="mt-2 flex items-start justify-between gap-2">
      <div>
        <p class="text-[15px] leading-none font-extrabold text-[#1c1c1c]">
          {{ euros(product.cents) }}
        </p>
        <p v-if="product.wasCents" class="mt-1 text-xs text-[#9a9a9a] line-through">
          {{ euros(product.wasCents) }}
        </p>
      </div>
      <QtyControl :sku="product.sku" />
    </div>
    <router-link
      :to="{ name: 'ShopProduct', params: { sku: product.sku } }"
      class="mt-1.5 text-sm leading-tight font-bold text-[#1c1c1c]"
    >
      {{ localizeProduct(product).name }}
    </router-link>
    <p class="mt-0.5 text-xs text-[#8a8a8a]">{{ localizeProduct(product).pack }}</p>
  </article>
</template>

<script setup>
import QtyControl from '@/comp/shop/QtyControl.vue'
import { euros } from '@/shop/catalog'
import { localizeProduct, t } from '@/shop/locale'

defineProps({
  product: { type: Object, required: true },
})
</script>
