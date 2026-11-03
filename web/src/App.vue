<template>
  <div id="app" class="min-h-screen">
    <FlashNotice :message="flashMessage" :duration="flashDuration" @dismiss="clearFlash" />
    <CookieConsent v-if="!isShop" />
    <router-view />
    <AssistantPanel v-if="route.meta.requiresWorkspace" />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import AssistantPanel from '@/comp/AssistantPanel.vue'
import CookieConsent from '@/comp/CookieConsent.vue'
import FlashNotice from '@/comp/FlashNotice.vue'
import { flashMessage, flashDuration, clearFlash } from '@/lib/flash'

const route = useRoute()
const isShop = computed(() => route.matched.some((record) => record.meta.shop))
</script>

<style scoped>
/* Additional app-level styles can go here */
</style>
