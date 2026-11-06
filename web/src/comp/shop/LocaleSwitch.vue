<template>
  <button
    type="button"
    class="rounded-full border border-[#e8e6dc] px-3 py-1 text-sm font-extrabold"
    @click="setShopLocale(shopLocale === 'ar' ? 'en' : 'ar')"
  >
    {{ shopLocale === 'ar' ? 'EN' : 'عربي' }}
  </button>
</template>

<script setup>
import { setShopLocale, shopLocale } from '@/shop/locale'
</script>
