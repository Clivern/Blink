<template>
  <div class="px-4 pt-5 md:px-6 md:pt-8">
    <div class="flex items-start justify-between gap-4">
      <div>
        <h1 class="text-2xl font-extrabold">{{ t('hello', { name: session.user.name }) }}</h1>
        <p class="mt-1 text-sm font-bold text-[#87867f]">
          {{
            session.user.provider
              ? t('signedInWith', { provider: providerLabel[session.user.provider] })
              : session.user.email
          }}
        </p>
        <p class="mt-1 text-sm font-bold text-[#87867f]">{{ t('profileNote') }}</p>
      </div>
      <button
        type="button"
        class="rounded-full bg-[#f5f4ed] px-4 py-2 text-sm font-extrabold"
        @click="signOut"
      >
        {{ t('signOut') }}
      </button>
    </div>

    <div class="mt-6 md:grid md:grid-cols-2 md:items-start md:gap-10">
      <form @submit.prevent="save">
        <h2 class="text-lg font-extrabold">{{ t('yourDetails') }}</h2>
        <label class="mt-4 block">
          <span class="text-sm font-extrabold">{{ t('name') }}</span>
          <input
            v-model="name"
            required
            minlength="2"
            class="mt-1.5 w-full rounded-2xl border border-[#e8e6dc] bg-[#f5f4ed] px-4 py-3 font-bold outline-none focus:border-[#141413]"
          />
        </label>
        <label class="mt-4 block">
          <span class="text-sm font-extrabold">{{ t('phone') }}</span>
          <input
            v-model="phone"
            required
            type="tel"
            minlength="6"
            class="mt-1.5 w-full rounded-2xl border border-[#e8e6dc] bg-[#f5f4ed] px-4 py-3 font-bold outline-none focus:border-[#141413]"
          />
        </label>
        <label class="mt-4 block">
          <span class="text-sm font-extrabold">{{ t('addressTitle') }}</span>
          <input
            v-model="address"
            required
            minlength="4"
            class="mt-1.5 w-full rounded-2xl border border-[#e8e6dc] bg-[#f5f4ed] px-4 py-3 font-bold outline-none focus:border-[#141413]"
          />
        </label>
        <button
          type="submit"
          class="mt-4 rounded-full bg-[#141413] px-4 py-2.5 text-sm font-extrabold text-[#faf9f5]"
        >
          {{ saved ? t('saved') : t('saveAddress') }}
        </button>
      </form>

      <section>
        <h2 class="mt-8 text-lg font-extrabold md:mt-0">{{ t('orders') }}</h2>
        <p v-if="orders.length === 0" class="mt-2 text-sm font-bold text-[#87867f]">
          {{ t('noOrders') }}
        </p>
        <ul v-else class="mt-3 space-y-3">
          <li v-for="order in orders" :key="order.id">
            <router-link
              :to="{ name: 'ShopOrder', params: { id: order.id } }"
              class="block rounded-2xl border border-[#e8e6dc] px-4 py-3"
            >
              <div class="flex items-center justify-between">
                <p class="font-extrabold">{{ order.id }}</p>
                <p class="text-sm font-extrabold">{{ euros(order.total) }}</p>
              </div>
              <p class="mt-1 flex gap-2 text-xs font-bold text-[#87867f]">
                <span>{{ t('orderItems', { n: order.items.length }) }}</span>
                <span dir="ltr">{{ order.eta }}</span>
              </p>
            </router-link>
          </li>
        </ul>
      </section>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useShopCart } from '@/shop/cart'
import { euros } from '@/shop/catalog'
import { t } from '@/shop/locale'
import { useShopSession } from '@/shop/session'

const providerLabel = { google: 'Google', facebook: 'Facebook', apple: 'Apple' }

const cart = useShopCart()
const session = useShopSession()
const router = useRouter()

const name = ref(session.user.name)
const phone = ref(session.user.phone)
const address = ref(cart.address)
const saved = ref(false)

const orders = computed(() => cart.orders.filter((order) => order.userId === session.user.id))

function save() {
  session.update({ name: name.value.trim(), phone: phone.value.trim() })
  cart.setAddress(address.value.trim())
  saved.value = true
}

function signOut() {
  session.logout()
  router.push({ name: 'ShopLogin' })
}
</script>
