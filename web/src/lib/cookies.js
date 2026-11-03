import { ref } from 'vue'

const CONSENT_STORAGE_KEY = 'blink_cookie_consent'
const CONSENT_VERSION = 1

export const showCookieBanner = ref(false)
export const showCookiePreferences = ref(false)

function readConsent() {
  const parsed = JSON.parse(localStorage.getItem(CONSENT_STORAGE_KEY))
  if (!parsed || parsed.version !== CONSENT_VERSION) return null
  return parsed
}

export function hasConsent() {
  return readConsent() !== null
}

export function saveConsent(categories) {
  const payload = {
    version: CONSENT_VERSION,
    updatedAt: new Date().toISOString(),
    categories: {
      essential: true,
      preferences: !!categories.preferences,
      analytics: !!categories.analytics,
    },
  }
  localStorage.setItem(CONSENT_STORAGE_KEY, JSON.stringify(payload))
  showCookieBanner.value = false
  showCookiePreferences.value = false
  return payload
}

export function acceptAllCookies() {
  return saveConsent({ preferences: true, analytics: true })
}

export function rejectOptionalCookies() {
  return saveConsent({ preferences: false, analytics: false })
}

export function initCookieConsent() {
  if (!readConsent()) {
    showCookieBanner.value = true
  }
}

export function consentCategoriesForEdit() {
  const consent = readConsent()
  return {
    essential: true,
    preferences: !!consent?.categories?.preferences,
    analytics: !!consent?.categories?.analytics,
  }
}
