<template>
  <div>
    <header class="px-4 pt-5 md:px-6 md:pt-8">
      <div class="flex items-center justify-between md:hidden">
        <p class="text-[28px] leading-none font-extrabold tracking-tight text-[#141413]">blink</p>
        <span class="flex items-center gap-2">
          <span class="rounded-full bg-[#f5f4ed] px-3 py-1 text-sm font-extrabold text-[#141413]">
            {{ t('eta') }}
          </span>
          <LocaleSwitch />
        </span>
      </div>
      <button type="button" class="mt-4 block w-full text-start" @click="openAddress">
        <p class="text-xs font-bold text-[#87867f]">{{ t('deliveringTo') }}</p>
        <p class="mt-0.5 flex items-center gap-1 text-base font-extrabold">
          {{ cart.address }}
          <svg
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2.4"
          >
            <path d="m6 9 6 6 6-6" />
          </svg>
        </p>
      </button>
      <router-link
        :to="{ name: 'ShopSearch' }"
        class="mt-4 flex items-center gap-2 rounded-2xl bg-[#f5f4ed] px-4 py-3 text-sm font-bold text-[#87867f] md:hidden"
      >
        <svg
          width="18"
          height="18"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2.2"
        >
          <circle cx="11" cy="11" r="6.5" />
          <path d="m16 16 4 4" />
        </svg>
        {{ t('searchPlaceholder') }}
      </router-link>
    </header>

    <div
      class="mt-5 flex gap-3 overflow-x-auto px-4 pb-1 [-ms-overflow-style:none] [scrollbar-width:none] md:flex-wrap md:gap-4 md:overflow-visible md:px-6 [&::-webkit-scrollbar]:hidden"
    >
      <router-link
        v-for="category in categories"
        :key="category.slug"
        :to="{ name: 'ShopCategory', params: { slug: category.slug } }"
        class="flex w-[76px] shrink-0 flex-col items-center gap-1.5"
      >
        <span
          class="flex h-[68px] w-[68px] items-center justify-center rounded-2xl text-3xl"
          :style="{ background: category.tint }"
        >
          {{ category.emoji }}
        </span>
        <span class="text-center text-[11px] leading-tight font-bold">{{
          categoryLabel(category)
        }}</span>
      </router-link>
    </div>

    <section
      class="mx-4 mt-5 overflow-hidden rounded-3xl bg-[#141413] px-5 py-5 text-[#faf9f5] md:mx-6 md:px-8 md:py-8"
    >
      <p class="text-xs font-extrabold tracking-wide uppercase opacity-80">{{ t('fromNearby') }}</p>
      <h1 class="mt-1 text-2xl leading-tight font-extrabold md:max-w-xl md:text-4xl">
        {{ t('hero') }}
      </h1>
      <p class="mt-2 text-sm font-bold opacity-90">
        {{ t('freeFrom', { amount: euros(FREE_DELIVERY_OVER) }) }}
      </p>
    </section>

    <section v-for="section in sections" :key="section.slug" class="mt-7">
      <div class="mb-3 flex items-end justify-between px-4 md:px-6">
        <h2 class="text-lg font-extrabold">{{ section.title }}</h2>
        <router-link
          :to="{ name: 'ShopCategory', params: { slug: section.slug } }"
          class="text-sm font-extrabold text-[#141413]"
        >
          {{ t('seeAll') }}
        </router-link>
      </div>
      <div
        class="grid grid-cols-2 gap-x-3 gap-y-5 px-4 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 md:px-6"
      >
        <ProductCard v-for="product in section.products" :key="product.sku" :product="product" />
      </div>
    </section>

    <div
      v-if="editing"
      class="fixed inset-0 z-40 flex items-end justify-center bg-black/40 md:items-center"
      @click.self="editing = false"
    >
      <form
        class="w-full max-w-md rounded-t-3xl bg-[#faf9f5] px-4 pt-5 pb-8 md:rounded-3xl md:p-6"
        @submit.prevent="saveAddress"
      >
        <p class="text-lg font-extrabold">{{ t('addressTitle') }}</p>
        <p class="mt-1 text-sm text-[#5e5d59]">
          {{ t('addressHelp') }}
        </p>
        <input
          v-model="draft"
          required
          minlength="4"
          class="mt-4 w-full rounded-2xl border border-[#e8e6dc] bg-[#f5f4ed] px-4 py-3 text-base font-bold outline-none focus:border-[#141413]"
          :placeholder="t('addressPlaceholder')"
        />
        <button
          type="submit"
          class="mt-4 w-full rounded-full bg-[#141413] py-3.5 text-base font-extrabold text-[#faf9f5]"
        >
          {{ t('saveAddress') }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import LocaleSwitch from '@/comp/shop/LocaleSwitch.vue'
import ProductCard from '@/comp/shop/ProductCard.vue'
import { useShopCart } from '@/shop/cart'
import { categories, categoryBySlug, euros, FREE_DELIVERY_OVER, productsIn } from '@/shop/catalog'
import { categoryLabel, t } from '@/shop/locale'

const cart = useShopCart()
const editing = ref(false)
const draft = ref('')

const sections = computed(() => [
  {
    slug: 'deals',
    title: categoryLabel(categoryBySlug('deals')),
    products: productsIn('deals').slice(0, 10),
  },
  {
    slug: 'fruit',
    title: categoryLabel(categoryBySlug('fruit')),
    products: productsIn('fruit').slice(0, 10),
  },
  {
    slug: 'drinks',
    title: t('somethingToDrink'),
    products: productsIn('drinks').slice(0, 10),
  },
])

function openAddress() {
  draft.value = cart.address
  editing.value = true
}

function saveAddress() {
  cart.setAddress(draft.value.trim())
  editing.value = false
}
</script>
