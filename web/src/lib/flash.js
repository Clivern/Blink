import { ref } from 'vue'

const message = ref(null)
const duration = ref(2500)
let timer = null

export const flashMessage = message
export const flashDuration = duration

export function showFlash(msg, ms = 2500) {
  clearTimeout(timer)
  message.value = msg
  duration.value = ms
  timer = setTimeout(() => {
    message.value = null
  }, ms)
}

export function clearFlash() {
  clearTimeout(timer)
  message.value = null
}
