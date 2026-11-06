<template>
  <button
    v-if="qty === 0 && large"
    type="button"
    class="w-full rounded-full bg-[#141413] py-3.5 text-base font-extrabold text-[#faf9f5] active:scale-[0.99]"
    @click.stop="add"
  >
    {{ t('addToBasket') }}
  </button>
  <div
    v-else-if="large"
    class="flex items-center justify-between rounded-full bg-[#141413] px-2 text-[#faf9f5]"
  >
    <button
      type="button"
      class="h-12 w-12 text-2xl"
      aria-label="Decrease quantity"
      @click.stop="dec"
    >
      −
    </button>
    <span class="text-base font-extrabold">{{ t('inBasket', { n: qty }) }}</span>
    <button
      type="button"
      class="h-12 w-12 text-2xl"
      aria-label="Increase quantity"
      @click.stop="add"
    >
      +
    </button>
  </div>
  <button
    v-else-if="qty === 0"
    type="button"
    class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-[#141413] text-lg leading-none font-bold text-[#faf9f5]"
    :aria-label="`Add ${sku}`"
    @click.stop="add"
  >
    +
  </button>
  <div
    v-else
    class="flex h-8 shrink-0 items-center rounded-full bg-[#141413] text-[#faf9f5]"
    @click.stop
  >
    <button
      type="button"
      class="h-8 w-7 text-lg leading-none"
      aria-label="Decrease quantity"
      @click="dec"
    >
      −
    </button>
    <span class="min-w-4 text-center text-sm font-extrabold">{{ qty }}</span>
    <button
      type="button"
      class="h-8 w-7 text-lg leading-none"
      aria-label="Increase quantity"
      @click="add"
    >
      +
    </button>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useShopCart } from '@/shop/cart'
import { t } from '@/shop/locale'

const props = defineProps({
  sku: { type: String, required: true },
  large: { type: Boolean, default: false },
})

const cart = useShopCart()
const qty = computed(() => cart.lines.find((line) => line.sku === props.sku)?.qty || 0)

function add() {
  cart.add(props.sku)
}

function dec() {
  cart.setQty(props.sku, qty.value - 1)
}
</script>
