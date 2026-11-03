<template>
  <div>
    <header class="sticky top-0 z-20 bg-white px-4 pt-4 pb-3 md:static md:px-6 md:pt-8">
      <form @submit.prevent>
        <label class="flex items-center gap-2 rounded-2xl bg-[#f4f4f4] px-4 py-3">
          <svg
            width="18"
            height="18"
            viewBox="0 0 24 24"
            fill="none"
            stroke="#8d8d8d"
            stroke-width="2.2"
          >
            <circle cx="11" cy="11" r="6.5" />
            <path d="m16 16 4 4" />
          </svg>
          <input
            v-model="query"
            type="search"
            autofocus
            :placeholder="t('searchStore')"
            class="w-full bg-transparent text-base font-bold outline-none placeholder:text-[#a3a3a3]"
          />
        </label>
      </form>
    </header>

    <p v-if="!query.trim()" class="px-4 pt-2 text-sm font-bold text-[#8a8a8a] md:px-6">
      {{ t('searchHint') }}
    </p>
    <p
      v-else-if="results.length === 0"
      class="px-4 pt-6 text-center text-sm font-bold text-[#8a8a8a] md:px-6"
    >
      {{ t('noMatch', { query: query.trim() }) }}
    </p>
    <div
      v-else
      class="grid grid-cols-2 gap-x-3 gap-y-5 px-4 pt-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 md:px-6"
    >
      <ProductCard v-for="product in results" :key="product.sku" :product="product" />
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import ProductCard from '@/comp/shop/ProductCard.vue'
import { products, searchProducts } from '@/shop/catalog'
import { localizeProduct, shopLocale, t } from '@/shop/locale'

const query = ref('')
const results = computed(() => {
  const needle = query.value.trim().toLowerCase()
  if (!needle) return []
  if (shopLocale.value !== 'ar') return searchProducts(query.value)

  return products.filter((product) => {
    const view = localizeProduct(product)
    const haystack = `${product.name} ${view.name} ${view.pack} ${product.brand}`.toLowerCase()
    return haystack.includes(needle)
  })
})
</script>
