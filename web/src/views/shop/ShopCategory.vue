<template>
  <div>
    <header
      class="sticky top-0 z-20 flex items-center gap-2 bg-white/95 px-3 py-3 backdrop-blur md:static md:px-6 md:pt-8"
    >
      <button
        type="button"
        class="flex h-10 w-10 items-center justify-center rounded-full bg-[#f4f4f4]"
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
      <div>
        <router-link
          v-if="parent"
          :to="{ name: 'ShopCategory', params: { slug: parent.slug } }"
          class="text-xs font-extrabold text-[#ff2f78]"
        >
          {{ categoryLabel(parent) }}
        </router-link>
        <h1 class="text-lg font-extrabold">{{ categoryLabel(category) }}</h1>
      </div>
    </header>

    <p class="px-4 pb-3 text-sm font-bold text-[#8a8a8a] md:px-6">
      {{ productsLabel(items.length) }}
    </p>

    <div
      v-if="groups.length"
      class="flex gap-3 overflow-x-auto px-4 pb-5 [-ms-overflow-style:none] [scrollbar-width:none] md:flex-wrap md:px-6 [&::-webkit-scrollbar]:hidden"
    >
      <router-link
        v-for="group in groups"
        :key="group.slug"
        :to="{ name: 'ShopCategory', params: { slug: group.slug } }"
        class="flex shrink-0 items-center gap-2 rounded-2xl bg-[#f7f7f7] py-2 pr-4 pl-2"
      >
        <span
          class="flex h-9 w-9 items-center justify-center rounded-xl text-lg"
          :style="{ background: group.tint }"
        >
          {{ group.emoji }}
        </span>
        <span class="text-sm font-extrabold">{{ categoryLabel(group) }}</span>
      </router-link>
    </div>

    <template v-if="groups.length">
      <section v-for="group in groups" :key="group.slug" class="mt-2">
        <div class="mb-3 flex items-end justify-between px-4 md:px-6">
          <h2 class="text-lg font-extrabold">{{ categoryLabel(group) }}</h2>
          <router-link
            :to="{ name: 'ShopCategory', params: { slug: group.slug } }"
            class="text-sm font-extrabold text-[#ff2f78]"
          >
            {{ t('seeAll') }}
          </router-link>
        </div>
        <div
          class="grid grid-cols-2 gap-x-3 gap-y-5 px-4 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 md:px-6"
        >
          <ProductCard v-for="product in group.products" :key="product.sku" :product="product" />
        </div>
      </section>
    </template>

    <div
      v-else
      class="grid grid-cols-2 gap-x-3 gap-y-5 px-4 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 md:px-6"
    >
      <ProductCard v-for="product in items" :key="product.sku" :product="product" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import ProductCard from '@/comp/shop/ProductCard.vue'
import { categoryBySlug, productsIn } from '@/shop/catalog'
import { categoryLabel, productsLabel, t } from '@/shop/locale'

const route = useRoute()
const router = useRouter()

const category = computed(() => categoryBySlug(route.params.slug))
const parent = computed(() =>
  category.value.parent ? categoryBySlug(category.value.parent) : null
)
const items = computed(() => productsIn(route.params.slug))
const groups = computed(() =>
  (category.value.children || []).map((child) => ({
    ...child,
    products: productsIn(child.slug),
  }))
)
</script>
