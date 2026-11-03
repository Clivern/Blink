<template>
  <div
    class="min-h-screen bg-theme-bg text-theme-text"
    :class="{ 'theme-emerald': $route.name !== 'ManageDashboard' }"
  >
    <nav class="bg-white border-b border-theme-border">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-14">
          <div class="flex items-center">
            <router-link :to="{ name: 'ManageDashboard' }" class="shrink-0 flex items-center text-theme-text">
              <AppLogo :alt="$t('common.logo_alt')" class="h-8 w-auto" />
            </router-link>
            <div class="hidden md:ml-8 md:flex md:space-x-1">
              <router-link
                v-for="item in items"
                :key="item.name"
                :to="{ name: item.name }"
                class="nav-link"
                active-class="nav-link-active"
              >
                {{ $t(`manage.nav.${item.key}`) }}
              </router-link>
            </div>
          </div>

          <div class="flex items-center gap-4">
            <p class="hidden text-sm font-medium text-theme-text sm:block">{{ $t('manage.title') }}</p>
            <router-link :to="{ name: 'ShopHome' }" class="nav-link">
              {{ $t('manage.view_store') }}
            </router-link>
          </div>
        </div>

        <nav class="flex gap-1 overflow-x-auto pb-2 md:hidden">
          <router-link
            v-for="item in items"
            :key="item.name"
            :to="{ name: item.name }"
            class="nav-link whitespace-nowrap"
            active-class="nav-link-active"
          >
            {{ $t(`manage.nav.${item.key}`) }}
          </router-link>
        </nav>
      </div>
    </nav>

    <main class="w-full px-4 sm:px-6 lg:px-8 py-8">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import AppLogo from '@/comp/AppLogo.vue'

const items = [
  { key: 'dashboard', name: 'ManageDashboard' },
  { key: 'categories', name: 'ManageCategories' },
  { key: 'products', name: 'ManageProducts' },
  { key: 'customers', name: 'ManageCustomers' },
  { key: 'orders', name: 'ManageOrders' },
]
</script>
