import { defineStore } from 'pinia'

const ACCOUNTS_KEY = 'blink.shop.accounts'
const SESSION_KEY = 'blink.shop.session'

function saved(key, fallback) {
  const raw = localStorage.getItem(key)
  return raw ? JSON.parse(raw) : fallback
}

async function digest(password) {
  const data = new TextEncoder().encode(password)
  const buf = await crypto.subtle.digest('SHA-256', data)
  return [...new Uint8Array(buf)].map((byte) => byte.toString(16).padStart(2, '0')).join('')
}

export const useShopSession = defineStore('shopSession', {
  state: () => ({
    accounts: saved(ACCOUNTS_KEY, []),
    userId: saved(SESSION_KEY, null),
  }),
  getters: {
    user(state) {
      return state.accounts.find((account) => account.id === state.userId) || null
    },
  },
  actions: {
    persist() {
      localStorage.setItem(ACCOUNTS_KEY, JSON.stringify(this.accounts))
      localStorage.setItem(SESSION_KEY, JSON.stringify(this.userId))
    },
    async register({ name, email, phone, password }) {
      const taken = this.accounts.some(
        (account) => account.email.toLowerCase() === email.toLowerCase()
      )
      if (taken) return 'emailTaken'

      this.accounts.push({
        id: `U${Date.now().toString(36).toUpperCase()}`,
        name,
        email,
        phone,
        passwordHash: await digest(password),
      })
      this.userId = this.accounts.at(-1).id
      this.persist()
      return ''
    },
    async login({ email, password }) {
      const passwordHash = await digest(password)
      const account = this.accounts.find(
        (item) =>
          item.email.toLowerCase() === email.toLowerCase() && item.passwordHash === passwordHash
      )
      if (!account) return 'badLogin'

      this.userId = account.id
      this.persist()
      return ''
    },
    loginWithProvider(provider) {
      const names = { google: 'Google', facebook: 'Facebook', apple: 'Apple' }
      let account = this.accounts.find((item) => item.provider === provider)

      if (!account) {
        account = {
          id: `U${Date.now().toString(36).toUpperCase()}`,
          name: names[provider],
          email: '',
          phone: '',
          provider,
          passwordHash: '',
        }
        this.accounts.push(account)
      }

      this.userId = account.id
      this.persist()
    },
    logout() {
      this.userId = null
      this.persist()
    },
    update({ name, phone }) {
      const account = this.accounts.find((item) => item.id === this.userId)
      account.name = name
      account.phone = phone
      this.persist()
    },
  },
})
