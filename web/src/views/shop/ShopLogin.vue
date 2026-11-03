<template>
  <div class="mx-auto max-w-md px-4 pt-8 md:pt-16">
    <h1 class="text-2xl font-extrabold">
      {{ mode === 'register' ? t('createAccount') : t('signIn') }}
    </h1>
    <p class="mt-1 text-sm font-bold text-[#8a8a8a]">{{ t('profileNote') }}</p>

    <div class="mt-5 space-y-3">
      <button
        v-for="provider in providers"
        :key="provider.id"
        type="button"
        class="flex w-full items-center justify-center gap-3 rounded-full border border-[#ececec] bg-white py-3 text-sm font-extrabold"
        @click="signInWith(provider.id)"
      >
        <svg
          v-if="provider.id === 'google'"
          width="18"
          height="18"
          viewBox="0 0 24 24"
          aria-hidden="true"
        >
          <path
            fill="#4285F4"
            d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
          />
          <path
            fill="#34A853"
            d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
          />
          <path
            fill="#FBBC05"
            d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
          />
          <path
            fill="#EA4335"
            d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
          />
        </svg>
        <svg
          v-else-if="provider.id === 'facebook'"
          width="18"
          height="18"
          viewBox="0 0 24 24"
          aria-hidden="true"
        >
          <path
            fill="#1877F2"
            d="M24 12.073C24 5.446 18.627 0 12 0S0 5.446 0 12.073C0 18.062 4.388 23.027 10.125 24v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796V24C19.612 23.027 24 18.062 24 12.073z"
          />
        </svg>
        <svg v-else width="18" height="18" viewBox="0 0 24 24" aria-hidden="true">
          <path
            fill="currentColor"
            d="M16.37 12.72c.03 2.9 2.54 3.86 2.57 3.88-.02.07-.4 1.38-1.33 2.73-.8 1.16-1.63 2.32-2.94 2.34-1.29.03-1.7-.76-3.18-.76-1.47 0-1.93.74-3.15.79-1.26.05-2.22-1.26-3.03-2.41-1.65-2.38-2.91-6.72-1.22-9.65.84-1.46 2.34-2.38 3.97-2.41 1.24-.02 2.41.83 3.18.83.76 0 2.18-1.03 3.67-.88.63.03 2.38.25 3.5 1.9-.09.06-2.09 1.22-2.04 3.64zM14.2 5.3c.66-.8 1.1-1.91.98-3.02-1.05.04-2.32.7-3.07 1.58-.67.78-1.26 2.02-1.1 3.21 1.17.09 2.36-.6 3.19-1.77z"
          />
        </svg>
        {{ t(provider.label) }}
      </button>
    </div>

    <p class="mt-5 text-center text-xs font-extrabold tracking-wide text-[#9a9a9a] uppercase">
      {{ t('or') }}
    </p>

    <form class="mt-5 space-y-4" @submit.prevent="submit">
      <p v-if="error" class="rounded-2xl bg-red-50 px-4 py-3 text-sm font-bold text-red-700">
        {{ t(error) }}
      </p>

      <label v-if="mode === 'register'" class="block">
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
        <span class="text-sm font-extrabold">{{ t('email') }}</span>
        <input
          v-model="email"
          required
          type="email"
          autocomplete="email"
          class="mt-1.5 w-full rounded-2xl border border-[#ececec] bg-[#fafafa] px-4 py-3 font-bold outline-none focus:border-[#1c1c1c]"
        />
      </label>

      <label v-if="mode === 'register'" class="block">
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
        <span class="text-sm font-extrabold">{{ t('password') }}</span>
        <input
          v-model="password"
          required
          type="password"
          minlength="6"
          autocomplete="current-password"
          class="mt-1.5 w-full rounded-2xl border border-[#ececec] bg-[#fafafa] px-4 py-3 font-bold outline-none focus:border-[#1c1c1c]"
        />
      </label>

      <button
        type="submit"
        class="w-full rounded-full bg-[#ff2f78] py-3.5 text-base font-extrabold text-white disabled:opacity-60"
        :disabled="saving"
      >
        {{ submitLabel }}
      </button>
    </form>

    <button
      type="button"
      class="mt-4 text-sm font-extrabold text-[#ff2f78]"
      @click="mode = mode === 'register' ? 'signin' : 'register'"
    >
      {{ mode === 'register' ? t('haveAccount') : t('needAccount') }}
    </button>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { t } from '@/shop/locale'
import { useShopSession } from '@/shop/session'

const route = useRoute()
const router = useRouter()
const session = useShopSession()

const providers = [
  { id: 'google', label: 'continueGoogle' },
  { id: 'facebook', label: 'continueFacebook' },
  { id: 'apple', label: 'continueApple' },
]

const mode = ref('signin')
const name = ref('')
const email = ref('')
const phone = ref('')
const password = ref('')
const saving = ref(false)
const error = ref('')

const submitLabel = computed(() => {
  if (saving.value && mode.value === 'register') return t('creating')
  if (saving.value) return t('signingIn')
  if (mode.value === 'register') return t('createAccount')
  return t('signIn')
})

function signInWith(provider) {
  session.loginWithProvider(provider)
  const redirect = route.query.redirect
  router.push(
    typeof redirect === 'string' && redirect.startsWith('/') && !redirect.startsWith('//')
      ? redirect
      : { name: 'ShopAccount' }
  )
}

async function submit() {
  error.value = ''
  saving.value = true

  const result =
    mode.value === 'register'
      ? await session.register({
          name: name.value.trim(),
          email: email.value.trim(),
          phone: phone.value.trim(),
          password: password.value,
        })
      : await session.login({ email: email.value.trim(), password: password.value })

  saving.value = false
  if (result) {
    error.value = result
    return
  }

  const redirect = route.query.redirect
  router.push(
    typeof redirect === 'string' && redirect.startsWith('/') && !redirect.startsWith('//')
      ? redirect
      : { name: 'ShopAccount' }
  )
}
</script>
