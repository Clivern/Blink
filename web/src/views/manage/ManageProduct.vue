<template>
  <div>
    <header class="page-header">
      <router-link :to="{ name: 'ManageProducts' }" class="text-sm text-theme-textLight hover:text-theme-text">
        ← {{ $t('manage.products.back') }}
      </router-link>
      <h1 class="page-title mt-2">
        {{ isNew ? $t('manage.products.new_title') : $t('manage.products.edit_title') }}
      </h1>
    </header>

    <form class="section max-w-3xl space-y-6" @submit.prevent="submit">
      <div v-if="error" class="alert-error">
        <p class="alert-text-error">{{ error }}</p>
      </div>

      <div class="grid gap-4 sm:grid-cols-2">
        <div>
          <label class="form-label" for="product-name">{{ $t('manage.products.name') }}</label>
          <input id="product-name" v-model.trim="form.name" class="input-field" type="text" maxlength="60" required />
        </div>
        <div>
          <label class="form-label" for="product-sku">{{ $t('manage.products.sku') }}</label>
          <input
            id="product-sku"
            v-model.trim="form.sku"
            class="input-field"
            type="text"
            pattern="[a-z0-9-]+"
            :disabled="!isNew"
            required
          />
          <p class="mt-1 text-xs text-theme-textLight">{{ $t('manage.products.sku_hint') }}</p>
        </div>
        <div>
          <label class="form-label" for="product-brand">{{ $t('manage.products.brand') }}</label>
          <input id="product-brand" v-model.trim="form.brand" class="input-field" type="text" maxlength="40" required />
        </div>
        <div>
          <label class="form-label" for="product-pack">{{ $t('manage.products.pack') }}</label>
          <input id="product-pack" v-model.trim="form.pack" class="input-field" type="text" maxlength="30" required />
        </div>
        <div>
          <label class="form-label" for="product-price">{{ $t('manage.products.price') }}</label>
          <input
            id="product-price"
            v-model="form.price"
            class="input-field"
            type="number"
            step="0.01"
            min="0.01"
            required
          />
        </div>
        <div>
          <label class="form-label" for="product-was-price">{{ $t('manage.products.was_price') }}</label>
          <input
            id="product-was-price"
            v-model="form.wasPrice"
            class="input-field"
            type="number"
            step="0.01"
            min="0.01"
          />
        </div>
        <div>
          <label class="form-label" for="product-category">{{ $t('manage.products.category') }}</label>
          <select id="product-category" v-model="form.category" class="input-field" required>
            <option v-for="category in leafCategories" :key="category.slug" :value="category.slug">
              {{ category.label }}
            </option>
          </select>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="form-label" for="product-emoji">{{ $t('manage.products.emoji') }}</label>
            <input id="product-emoji" v-model.trim="form.emoji" class="input-field" type="text" maxlength="4" required />
          </div>
          <div>
            <label class="form-label" for="product-tint">{{ $t('manage.products.tint') }}</label>
            <input id="product-tint" v-model="form.tint" class="input-field h-10" type="color" required />
          </div>
        </div>
        <div>
          <label class="form-label" for="product-unit-price">{{ $t('manage.products.unit_price') }}</label>
          <input
            id="product-unit-price"
            v-model.trim="form.unitPrice"
            class="input-field"
            type="text"
            maxlength="30"
            required
          />
          <p class="mt-1 text-xs text-theme-textLight">{{ $t('manage.products.unit_price_hint') }}</p>
        </div>
      </div>

      <div>
        <label class="form-label" for="product-blurb">{{ $t('manage.products.blurb') }}</label>
        <textarea id="product-blurb" v-model.trim="form.blurb" class="input-field" rows="3" maxlength="200" required />
      </div>

      <label class="flex items-center gap-2 text-sm">
        <input v-model="form.popular" class="checkbox" type="checkbox" />
        {{ $t('manage.products.popular') }}
      </label>

      <div class="flex gap-3">
        <button type="submit" class="btn-primary bg-theme-text">
          {{ isNew ? $t('common.create') : $t('common.save') }}
        </button>
        <router-link :to="{ name: 'ManageProducts' }" class="btn-secondary">
          {{ $t('common.cancel') }}
        </router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { computed, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRoute, useRouter } from 'vue-router'
import {
  addProduct,
  categories,
  isSkuTaken,
  productBySku,
  updateProduct,
} from '@/shop/catalog'

const { t } = useI18n()
const route = useRoute()
const router = useRouter()

const isNew = computed(() => route.params.sku === 'new')
const error = ref('')

const leafCategories = categories.flatMap((category) =>
  category.children.map((child) => ({ slug: child.slug, label: `${category.name} / ${child.name}` }))
)

const existing = isNew.value ? null : productBySku(route.params.sku)

const form = reactive({
  name: existing?.name ?? '',
  sku: existing?.sku ?? '',
  brand: existing?.brand ?? '',
  pack: existing?.pack ?? '',
  price: existing ? (existing.cents / 100).toFixed(2) : '',
  wasPrice: existing?.wasCents ? (existing.wasCents / 100).toFixed(2) : '',
  category: existing?.category ?? leafCategories[0].slug,
  emoji: existing?.emoji ?? '',
  tint: existing?.tint ?? '#f5f5f5',
  unitPrice: existing?.unitPrice ?? '',
  blurb: existing?.blurb ?? '',
  popular: existing?.tags.includes('popular') ?? false,
})

function toCents(euros) {
  return Math.round(Number(euros) * 100)
}

function submit() {
  if (isNew.value && isSkuTaken(form.sku)) {
    error.value = t('manage.products.duplicate_sku')
    return
  }

  error.value = ''
  const data = {
    name: form.name,
    brand: form.brand,
    pack: form.pack,
    cents: toCents(form.price),
    wasCents: form.wasPrice ? toCents(form.wasPrice) : undefined,
    category: form.category,
    emoji: form.emoji,
    tint: form.tint,
    unitPrice: form.unitPrice,
    blurb: form.blurb,
    tags: form.popular ? ['popular'] : [],
  }

  if (isNew.value) {
    addProduct({ sku: form.sku, ...data })
  } else {
    updateProduct(route.params.sku, data)
  }

  router.push({ name: 'ManageProducts' })
}
</script>
