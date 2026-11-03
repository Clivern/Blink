<template>
  <div>
    <header class="page-header">
      <h1 class="page-title">{{ $t('manage.categories.title') }}</h1>
      <p class="page-subtitle">{{ $t('manage.categories.subtitle') }}</p>
    </header>

    <div class="grid gap-6 lg:grid-cols-[1fr_22rem]">
      <div class="section !p-0 overflow-hidden">
        <ul class="divide-y divide-theme-border">
          <li v-for="category in categories" :key="category.slug">
            <div class="flex items-center gap-3 px-6 py-4">
              <span class="text-2xl">{{ category.emoji }}</span>
              <div class="flex-1">
                <p class="font-medium">{{ category.name }}</p>
                <p class="text-xs text-theme-textLight">
                  {{ $t('manage.categories.subcategories', { count: category.children.length }) }}
                  · {{ $t('manage.categories.products_count', { count: productCount(category.slug) }) }}
                </p>
              </div>
              <button type="button" class="btn-secondary" @click="startEdit(category.slug)">
                {{ $t('common.edit') }}
              </button>
              <button
                type="button"
                class="btn-danger"
                :disabled="!canDelete(category.slug)"
                @click="askDelete(category.slug)"
              >
                {{ $t('common.delete') }}
              </button>
            </div>

            <ul class="divide-y divide-theme-border bg-theme-hover">
              <li
                v-for="child in category.children"
                :key="child.slug"
                class="flex items-center gap-3 px-6 py-3 pl-14"
              >
                <span class="text-xl">{{ child.emoji }}</span>
                <div class="flex-1">
                  <p class="text-sm font-medium">{{ child.name }}</p>
                  <p class="text-xs text-theme-textLight">
                    {{ $t('manage.categories.products_count', { count: productCount(child.slug) }) }}
                  </p>
                </div>
                <button type="button" class="btn-secondary" @click="startEdit(child.slug)">
                  {{ $t('common.edit') }}
                </button>
                <button
                  type="button"
                  class="btn-danger"
                  :disabled="!canDelete(child.slug)"
                  @click="askDelete(child.slug)"
                >
                  {{ $t('common.delete') }}
                </button>
              </li>
            </ul>
          </li>
        </ul>
      </div>

      <form class="section space-y-4 self-start" @submit.prevent="submit">
        <h2 class="section-title !mb-0">
          {{ editingSlug ? $t('manage.categories.edit_title') : $t('manage.categories.new_title') }}
        </h2>

        <div v-if="error" class="alert-error">
          <p class="alert-text-error">{{ error }}</p>
        </div>

        <div>
          <label class="form-label" for="category-name">{{ $t('manage.categories.name') }}</label>
          <input
            id="category-name"
            v-model.trim="form.name"
            class="input-field"
            type="text"
            maxlength="40"
            required
          />
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="form-label" for="category-emoji">{{ $t('manage.categories.emoji') }}</label>
            <input
              id="category-emoji"
              v-model.trim="form.emoji"
              class="input-field"
              type="text"
              maxlength="4"
              required
            />
          </div>
          <div>
            <label class="form-label" for="category-tint">{{ $t('manage.categories.tint') }}</label>
            <input id="category-tint" v-model="form.tint" class="input-field h-10" type="color" required />
          </div>
        </div>

        <div v-if="!editingSlug">
          <label class="form-label" for="category-parent">{{ $t('manage.categories.parent') }}</label>
          <select id="category-parent" v-model="form.parent" class="input-field">
            <option value="">{{ $t('manage.categories.top_level') }}</option>
            <option v-for="category in categories" :key="category.slug" :value="category.slug">
              {{ category.name }}
            </option>
          </select>
        </div>

        <div class="flex gap-3">
          <button type="submit" class="btn-primary bg-theme-text">
            {{ editingSlug ? $t('common.save') : $t('common.create') }}
          </button>
          <button v-if="editingSlug" type="button" class="btn-secondary" @click="reset">
            {{ $t('common.cancel') }}
          </button>
        </div>
      </form>
    </div>

    <ConfirmModal
      :open="!!pendingDelete"
      :title="$t('manage.categories.delete_title')"
      :message="$t('manage.categories.confirm_delete', { name: pendingDelete?.name })"
      :confirm-label="$t('common.delete')"
      :busy-label="$t('common.delete')"
      @confirm="confirmDelete"
      @cancel="pendingDelete = null"
    />
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ConfirmModal from '@/comp/ConfirmModal.vue'
import {
  addCategory,
  categories,
  categoryUsage,
  deleteCategory,
  findCategory,
  isCategorySlugTaken,
  products,
  slugify,
  updateCategory,
} from '@/shop/catalog'

const BUILT_IN = ['deals']

const { t } = useI18n()
const editingSlug = ref(null)
const pendingDelete = ref(null)
const error = ref('')
const form = reactive({ name: '', emoji: '', tint: '#f5f5f5', parent: '' })

function productCount(slug) {
  return products.filter((product) => product.category === slug).length
}

function canDelete(slug) {
  return !BUILT_IN.includes(slug) && categoryUsage(slug) === 0
}

function startEdit(slug) {
  const category = findCategory(slug)

  editingSlug.value = slug
  error.value = ''
  Object.assign(form, { name: category.name, emoji: category.emoji, tint: category.tint, parent: '' })
}

function reset() {
  editingSlug.value = null
  error.value = ''
  Object.assign(form, { name: '', emoji: '', tint: '#f5f5f5', parent: '' })
}

function submit() {
  if (editingSlug.value) {
    updateCategory(editingSlug.value, { name: form.name, emoji: form.emoji, tint: form.tint })
    reset()
    return
  }

  if (isCategorySlugTaken(slugify(form.name))) {
    error.value = t('manage.categories.duplicate')
    return
  }

  addCategory({ name: form.name, emoji: form.emoji, tint: form.tint, parent: form.parent || null })
  reset()
}

function askDelete(slug) {
  pendingDelete.value = findCategory(slug)
}

function confirmDelete() {
  deleteCategory(pendingDelete.value.slug)
  pendingDelete.value = null
}
</script>
