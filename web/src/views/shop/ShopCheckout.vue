<template>
  <div class="px-4 pt-5 pb-6 md:mx-auto md:max-w-xl md:px-6 md:pt-8">
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
    <h1 class="mt-3 text-2xl font-extrabold">{{ t('checkout') }}</h1>
    <p class="mt-1 text-sm font-bold text-[#8a8a8a]">
      {{ cart.count === 1 ? t('oneItem') : t('nItems', { n: cart.count }) }} · {{ t('arriving') }}
    </p>

    <form class="mt-5 space-y-4" @submit.prevent="submit">
      <p v-if="error" class="rounded-2xl bg-red-50 px-4 py-3 text-sm font-bold text-red-700">
        {{ error }}
      </p>

      <label class="block">
        <span class="text-sm font-extrabold">{{ t('name') }}</span>
        <input
          v-model="name"
          required
          minlength="2"
          autocomplete="name"
          class="mt-1.5 w-full rounded-2xl border border-[#ececec] bg-[#fafafa] px-4 py-3 font-bold outline-none focus:border-[#1c1c1c]"
        />
      </label>

      <label class="block">
        <span class="text-sm font-extrabold">{{ t('phone') }}</span>
        <input
          v-model="phone"
          required
          type="tel"
          minlength="6"
          autocomplete="tel"
          class="mt-1.5 w-full rounded-2xl border border-[#ececec] bg-[#fafafa] px-4 py-3 font-bold outline-none focus:border-[#1c1c1c]"
        />
      </label>

      <label class="block">
        <span class="text-sm font-extrabold">{{ t('address') }}</span>
        <input
          v-model="address"
          required
          minlength="4"
          autocomplete="street-address"
          class="mt-1.5 w-full rounded-2xl border border-[#ececec] bg-[#fafafa] px-4 py-3 font-bold outline-none focus:border-[#1c1c1c]"
        />
      </label>

      <fieldset>
        <legend class="text-sm font-extrabold">{{ t('window') }}</legend>
        <label
          v-for="option in windows"
          :key="option.id"
          class="mt-2 flex items-center gap-3 rounded-2xl border border-[#ececec] px-4 py-3 font-bold"
        >
          <input v-model="windowChoice" type="radio" name="window" required :value="option.id" />
          {{ option.label }}
        </label>
      </fieldset>

      <label class="block">
        <span class="text-sm font-extrabold">{{ t('note') }}</span>
        <textarea
          v-model="note"
          rows="2"
          class="mt-1.5 w-full rounded-2xl border border-[#ececec] bg-[#fafafa] px-4 py-3 font-bold outline-none focus:border-[#1c1c1c]"
          :placeholder="t('notePlaceholder')"
        />
      </label>

      <div
        class="fixed inset-x-0 bottom-[4.15rem] z-20 bg-white px-4 pt-3 pb-2 md:static md:mt-2 md:px-0"
      >
        <button
          type="submit"
          class="w-full rounded-full bg-[#ff2f78] py-3.5 text-base font-extrabold text-white disabled:opacity-60"
          :disabled="saving"
        >
          {{ saving ? t('placing') : t('placeOrder', { amount: euros(cart.total) }) }}
        </button>
        <p class="mt-2 text-center text-[11px] font-bold text-[#9a9a9a]">
          {{ t('preview') }}
        </p>
      </div>
    </form>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useShopCart } from '@/shop/cart'
import { euros } from '@/shop/catalog'
import { t } from '@/shop/locale'
import { useShopSession } from '@/shop/session'

const cart = useShopCart()
const session = useShopSession()
const router = useRouter()

const name = ref(session.user.name)
const phone = ref(session.user.phone)
const address = ref(cart.address)
const note = ref('')
const windowChoice = ref('asap')
const saving = ref(false)
const error = ref('')

const windows = computed(() => [
  { id: 'asap', label: t('windowAsap') },
  { id: '30', label: t('window30') },
  { id: '60', label: t('window60') },
])

function submit() {
  if (cart.count === 0) {
    error.value = t('emptyError')
    return
  }

  error.value = ''
  saving.value = true
  cart.setAddress(address.value.trim())

  const order = cart.place({
    name: name.value.trim(),
    phone: phone.value.trim(),
    userId: session.user.id,
    note: note.value.trim(),
    window: windows.value.find((option) => option.id === windowChoice.value).label,
  })

  router.push({ name: 'ShopOrder', params: { id: order.id } })
}
</script>
