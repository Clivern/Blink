<template>
  <div class="shop-root min-h-screen bg-white" :dir="shopLocale === 'ar' ? 'rtl' : 'ltr'">
    <header
      class="sticky top-0 z-30 hidden border-b border-[#f0f0f0] bg-white/95 backdrop-blur md:block"
    >
      <div class="mx-auto flex h-16 max-w-6xl items-center gap-6 px-6">
        <router-link
          :to="{ name: 'ShopHome' }"
          class="text-2xl leading-none font-extrabold tracking-tight text-[#ff2f78]"
        >
          blink
        </router-link>
        <router-link
          :to="{ name: 'ShopSearch' }"
          class="flex max-w-md flex-1 items-center gap-2 rounded-2xl bg-[#f4f4f4] px-4 py-2.5 text-sm font-bold text-[#8d8d8d]"
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
        <span class="rounded-full bg-[#ffe4ef] px-3 py-1 text-sm font-extrabold text-[#ff2f78]">
          {{ t('eta') }}
        </span>
        <LocaleSwitch />
        <nav class="flex items-center gap-1">
          <router-link
            v-for="tab in tabs"
            :key="tab.name"
            :to="{ name: tab.name }"
            class="relative flex items-center gap-1.5 rounded-full px-3 py-2 text-sm font-extrabold"
            :class="isActive(tab) ? 'bg-[#ffe4ef] text-[#ff2f78]' : 'text-[#5c5c5c]'"
          >
            {{ tab.label }}
            <span
              v-if="tab.name === 'ShopCart' && cart.count"
              class="flex h-5 min-w-5 items-center justify-center rounded-full bg-[#ff2f78] px-1 text-[11px] font-extrabold text-white"
            >
              {{ cart.count }}
            </span>
          </router-link>
        </nav>
      </div>
    </header>

    <div class="mx-auto w-full max-w-6xl">
      <div v-if="route.name !== 'ShopHome'" class="flex justify-end px-4 pt-3 md:hidden">
        <LocaleSwitch />
      </div>
      <main>
        <router-view />
      </main>

      <footer class="mt-16 border-t border-[#f0f0f0] px-4 pt-10 pb-40 md:px-6 md:pb-12">
        <div class="grid gap-8 sm:grid-cols-3">
          <div>
            <p class="text-2xl font-extrabold tracking-tight text-[#ff2f78]">blink</p>
            <p class="mt-2 max-w-xs text-sm font-bold text-[#6d6d6d]">
              {{ t('tagline') }}
            </p>
          </div>
          <div>
            <p class="text-sm font-extrabold">Blink</p>
            <ul class="mt-3 space-y-2">
              <li v-for="page in companyPages" :key="page.slug">
                <router-link
                  :to="{ name: 'ShopPage', params: { slug: page.slug } }"
                  class="text-sm font-bold text-[#5c5c5c]"
                >
                  {{ localizePage(page).title }}
                </router-link>
              </li>
            </ul>
          </div>
          <div>
            <p class="text-sm font-extrabold">{{ t('legal') }}</p>
            <ul class="mt-3 space-y-2">
              <li v-for="page in legalPages" :key="page.slug">
                <router-link
                  :to="{ name: 'ShopPage', params: { slug: page.slug } }"
                  class="text-sm font-bold text-[#5c5c5c]"
                >
                  {{ localizePage(page).title }}
                </router-link>
              </li>
            </ul>
          </div>
        </div>
        <p class="mt-10 text-xs font-bold text-[#9a9a9a]">{{ t('rights') }}</p>
      </footer>
    </div>

    <router-link
      v-if="cart.count && showBasket"
      :to="{ name: 'ShopCart' }"
      class="fixed bottom-[4.5rem] left-4 right-4 z-30 flex items-center gap-3 rounded-full bg-[#1c1c1c] px-4 py-3 text-white shadow-lg md:hidden"
    >
      <span
        class="flex h-7 min-w-7 items-center justify-center rounded-full bg-[#ff2f78] px-1.5 text-sm font-extrabold"
      >
        {{ cart.count }}
      </span>
      <span class="flex-1 text-sm font-extrabold">{{ t('viewBasket') }}</span>
      <span class="text-sm font-extrabold">{{ euros(cart.total) }}</span>
    </router-link>

    <nav
      class="fixed inset-x-0 bottom-0 z-30 flex justify-around border-t border-[#f0f0f0] bg-white px-2 pt-2 pb-[max(0.6rem,env(safe-area-inset-bottom))] md:hidden"
    >
      <router-link
        v-for="tab in tabs"
        :key="tab.name"
        :to="{ name: tab.name }"
        class="relative flex w-16 flex-col items-center gap-0.5 text-[11px] font-bold"
        :class="isActive(tab) ? 'text-[#ff2f78]' : 'text-[#8a8a8a]'"
      >
        <svg
          v-if="tab.name === 'ShopHome'"
          width="22"
          height="22"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <path d="M4 10.5 12 4l8 6.5V20a1 1 0 0 1-1 1h-5v-6H10v6H5a1 1 0 0 1-1-1v-9.5Z" />
        </svg>
        <svg
          v-else-if="tab.name === 'ShopSearch'"
          width="22"
          height="22"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <circle cx="11" cy="11" r="6.5" />
          <path d="m16 16 4 4" />
        </svg>
        <svg
          v-else-if="tab.name === 'ShopCart'"
          width="22"
          height="22"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <path d="M6 7h15l-1.5 9h-12L5 4H2" />
          <circle cx="9" cy="20" r="1.2" fill="currentColor" />
          <circle cx="18" cy="20" r="1.2" fill="currentColor" />
        </svg>
        <svg
          v-else
          width="22"
          height="22"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <circle cx="12" cy="8" r="3.2" />
          <path d="M5 19.5c1.4-3 3.8-4.5 7-4.5s5.6 1.5 7 4.5" />
        </svg>
        <span>{{ tab.label }}</span>
        <span
          v-if="tab.name === 'ShopCart' && cart.count"
          class="absolute top-0 end-2 flex h-4 min-w-4 items-center justify-center rounded-full bg-[#ff2f78] px-1 text-[10px] font-extrabold text-white"
        >
          {{ cart.count }}
        </span>
      </router-link>
    </nav>
  </div>
</template>

<script setup>
import { computed, watch } from 'vue'
import { useRoute } from 'vue-router'
import LocaleSwitch from '@/comp/shop/LocaleSwitch.vue'
import { useShopCart } from '@/shop/cart'
import { categoryBySlug, euros, productBySku } from '@/shop/catalog'
import { categoryLabel, localizePage, localizeProduct, shopLocale, t } from '@/shop/locale'
import { pageBySlug, pages } from '@/shop/pages'

const route = useRoute()
const cart = useShopCart()

const showBasket = computed(() => !['ShopCart', 'ShopCheckout', 'ShopOrder'].includes(route.name))

const tabs = computed(() => [
  { name: 'ShopHome', label: t('shop') },
  { name: 'ShopSearch', label: t('search') },
  { name: 'ShopCart', label: t('basket') },
  { name: 'ShopAccount', label: t('you') },
])

const companyPages = pages.filter((page) => ['about', 'delivery', 'help'].includes(page.slug))
const legalPages = pages.filter((page) => ['terms', 'privacy'].includes(page.slug))

function isActive(tab) {
  if (tab.name === 'ShopHome') {
    return ['ShopHome', 'ShopCategory', 'ShopProduct'].includes(route.name)
  }
  if (tab.name === 'ShopAccount') {
    return ['ShopAccount', 'ShopLogin'].includes(route.name)
  }
  return route.name === tab.name
}

watch(
  () => [route.fullPath, shopLocale.value],
  () => {
    let title = route.name === 'ShopHome' ? t('shop') : route.meta.title

    if (route.name === 'ShopCategory') title = categoryLabel(categoryBySlug(route.params.slug))
    if (route.name === 'ShopProduct') title = localizeProduct(productBySku(route.params.sku)).name
    if (route.name === 'ShopPage') title = localizePage(pageBySlug(route.params.slug)).title
    if (route.name === 'ShopLogin') title = t('signIn')
    if (route.name === 'ShopAccount') title = t('you')

    const storeTitle = `blink - ${t('tagline')}`

    document.title = route.name === 'ShopHome' ? storeTitle : `${title} · Blink`
    document.querySelector('meta[name="description"]').setAttribute('content', t('tagline'))
  },
  { immediate: true }
)
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Cairo:wght@500;700;800&family=Nunito:wght@500;700;800&display=swap');

.shop-root {
  font-family: 'Nunito', 'Inter', ui-sans-serif, system-ui, sans-serif;
  color: #1c1c1c;
  background: #fff;
}

.shop-root[dir='rtl'] {
  font-family: 'Cairo', 'Nunito', ui-sans-serif, system-ui, sans-serif;
}

.shop-root[dir='rtl'] .shop-back {
  transform: scaleX(-1);
}
</style>
