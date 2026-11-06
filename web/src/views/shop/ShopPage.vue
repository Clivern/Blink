<template>
  <article class="mx-auto max-w-2xl px-4 pt-6 md:px-6 md:pt-10">
    <h1 class="text-3xl font-extrabold">{{ shown.title }}</h1>
    <p class="mt-3 text-base font-bold text-[#5e5d59]">{{ shown.lead }}</p>
    <section v-for="section in shown.sections" :key="section.heading" class="mt-8">
      <h2 class="text-lg font-extrabold">{{ section.heading }}</h2>
      <p class="mt-2 text-[15px] leading-relaxed font-semibold text-[#30302e]">
        {{ section.body }}
      </p>
    </section>
  </article>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { localizePage } from '@/shop/locale'
import { pageBySlug } from '@/shop/pages'

const route = useRoute()
const shown = computed(() => localizePage(pageBySlug(route.params.slug)))
</script>
