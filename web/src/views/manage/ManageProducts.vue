<template>
  <div>
    <header class="page-header flex flex-wrap items-end justify-between gap-4">
      <div>
        <h1 class="page-title">{{ $t('manage.products.title') }}</h1>
        <p class="page-subtitle">{{ $t('manage.products.subtitle') }}</p>
      </div>
      <router-link :to="{ name: 'ManageProduct', params: { sku: 'new' } }" class="btn-primary bg-theme-text">
        {{ $t('manage.products.new') }}
      </router-link>
    </header>

    <div class="mb-4 flex flex-col gap-3 sm:flex-row">
      <input
        v-model="query"
        class="input-field sm:max-w-xs"
        type="search"
        :placeholder="$t('manage.products.search')"
      />
      <select v-model="categoryFilter" class="input-field sm:max-w-xs">
        <option value="">{{ $t('manage.products.all_categories') }}</option>
        <option v-for="category in leafCategories" :key="category.slug" :value="category.slug">
          {{ category.label }}
        </option>
      </select>
    </div>

    <div class="section !p-0 overflow-x-auto">
      <table class="min-w-full divide-y divide-theme-border text-sm">
        <thead class="bg-theme-hover text-left text-xs uppercase text-theme-textLight">
          <tr>
            <th class="px-4 py-3">{{ $t('manage.products.name') }}</th>
            <th class="px-4 py-3">{{ $t('manage.products.category') }}</th>
            <th class="px-4 py-3 text-right">{{ $t('manage.products.price') }}</th>
            <th class="px-4 py-3 text-right"></th>
          </tr>
        </thead>
        <tbody class="divide-y divide-theme-border">
          <tr v-for="product in filtered" :key="product.sku">
            <td class="px-4 py-3">
              <div class="flex items-center gap-3">
                <span class="text-xl">{{ product.emoji }}</span>
                <div>
                  <p class="font-medium">{{ product.name }}</p>
                  <p class="text-xs text-theme-textLight">{{ product.sku }} · {{ product.pack }}</p>
                </div>
              </div>
            </td>
            <td class="px-4 py-3">{{ categoryName(product.category) }}</td>
            <td class="px-4 py-3 text-right">
              <p class="font-semibold">{{ euros(product.cents) }}</p>
              <p v-if="product.wasCents" class="text-xs text-theme-textLight line-through">
                {{ euros(product.wasCents) }}
              </p>
            </td>
            <td class="px-4 py-3">
              <div class="flex justify-end gap-2">
                <router-link
                  :to="{ name: 'ManageProduct', params: { sku: product.sku } }"
                  class="btn-secondary"
                >
                  {{ $t('common.edit') }}
                </router-link>
                <button type="button" class="btn-danger" @click="pendingDelete = product">
                  {{ $t('common.delete') }}
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <p v-if="!filtered.length" class="py-8 text-center text-sm text-theme-textLight">
        {{ $t('manage.products.empty') }}
      </p>
    </div>

    <ConfirmModal
      :open="!!pendingDelete"
      :title="$t('manage.products.delete_title')"
      :message="$t('manage.products.confirm_delete', { name: pendingDelete?.name })"
      :confirm-label="$t('common.delete')"
      :busy-label="$t('common.delete')"
      @confirm="confirmDelete"
      @cancel="pendingDelete = null"
    />
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import ConfirmModal from '@/comp/ConfirmModal.vue'
import { categories, categoryBySlug, deleteProduct, euros, products } from '@/shop/catalog'

const query = ref('')
const categoryFilter = ref('')
const pendingDelete = ref(null)

const leafCategories = computed(() =>
  categories.flatMap((category) =>
    category.children.map((child) => ({ slug: child.slug, label: `${category.name} / ${child.name}` }))
  )
)

const filtered = computed(() => {
  const needle = query.value.trim().toLowerCase()

  return products.filter((product) => {
    const matchesCategory = !categoryFilter.value || product.category === categoryFilter.value
    const haystack = `${product.name} ${product.brand} ${product.sku}`.toLowerCase()
    return matchesCategory && haystack.includes(needle)
  })
})

function categoryName(slug) {
  return categoryBySlug(slug).name
}

function confirmDelete() {
  deleteProduct(pendingDelete.value.sku)
  pendingDelete.value = null
}
</script>
