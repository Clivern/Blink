<template>
  <div>
    <header class="sticky top-0 z-20 bg-[#faf9f5]/95 px-3 py-3 backdrop-blur md:static md:px-6 md:pt-6">
      <button
        type="button"
        class="flex h-10 w-10 items-center justify-center rounded-full bg-[#f5f4ed]"
        aria-label="Back"
        @click="router.back()"
      >
        <svg
          width="18"
          height="18"
          viewBox="0 0 24 24"
          fill="none"
          class="shop-back"
          stroke="currentColor"
          stroke-width="2.4"
        >
          <path d="M15 5 8 12l7 7" />
        </svg>
      </button>
    </header>

    <div class="lg:grid lg:grid-cols-2 lg:items-start lg:gap-10 lg:px-6">
      <div
        class="mx-4 flex aspect-[5/4] items-center justify-center rounded-3xl lg:mx-0 lg:aspect-square"
        :style="{ background: product.tint }"
      >
        <span class="text-8xl lg:text-9xl" aria-hidden="true">{{ product.emoji }}</span>
      </div>

      <div>
        <div class="px-4 pt-5 pb-4 lg:px-0 lg:pt-2">
          <p class="text-sm font-extrabold text-[#141413]">{{ shown.brand }}</p>
          <h1 class="mt-1 text-2xl font-extrabold">{{ shown.name }}</h1>
          <p class="mt-1 text-sm font-bold text-[#87867f]">
            {{ shown.pack }} · {{ shown.unitPrice }}
          </p>
          <div class="mt-3 flex items-baseline gap-2">
            <p class="text-2xl font-extrabold">{{ euros(product.cents) }}</p>
            <p v-if="product.wasCents" class="text-sm font-bold text-[#87867f] line-through">
              {{ euros(product.wasCents) }}
            </p>
          </div>
          <p class="mt-4 text-[15px] leading-relaxed font-semibold text-[#30302e]">
            {{ shown.blurb }}
          </p>
          <div class="mt-4 flex flex-wrap gap-2">
            <router-link
              v-if="parent"
              :to="{ name: 'ShopCategory', params: { slug: parent.slug } }"
              class="inline-flex rounded-full bg-[#f5f4ed] px-3 py-1.5 text-sm font-extrabold"
            >
              {{ categoryLabel(parent) }}
            </router-link>
            <router-link
              :to="{ name: 'ShopCategory', params: { slug: product.category } }"
              class="inline-flex rounded-full bg-[#f5f4ed] px-3 py-1.5 text-sm font-extrabold"
            >
              {{ categoryLabel(category) }}
            </router-link>
          </div>
        </div>

        <div class="px-4 lg:px-0">
          <QtyControl :sku="product.sku" large />
        </div>
      </div>
    </div>

    <section v-if="related.length" class="mt-6">
      <h2 class="px-4 text-lg font-extrabold md:px-6">{{ t('goesWith') }}</h2>
      <div
        class="mt-3 grid grid-cols-2 gap-x-3 gap-y-5 px-4 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 md:px-6"
      >
        <ProductCard v-for="item in related" :key="item.sku" :product="item" />
      </div>
    </section>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import ProductCard from '@/comp/shop/ProductCard.vue'
import QtyControl from '@/comp/shop/QtyControl.vue'
import { categoryBySlug, euros, productBySku, productsIn } from '@/shop/catalog'
import { categoryLabel, localizeProduct, t } from '@/shop/locale'

const route = useRoute()
const router = useRouter()

const product = computed(() => productBySku(route.params.sku))
const shown = computed(() => localizeProduct(product.value))
const category = computed(() => categoryBySlug(product.value.category))
const parent = computed(() =>
  category.value.parent ? categoryBySlug(category.value.parent) : null
)
const related = computed(() =>
  productsIn(product.value.category)
    .filter((item) => item.sku !== product.value.sku)
    .slice(0, 5)
)
</script>
