import { setup_api } from '@/api'

const OSS = 'oss'
const SAAS = 'saas'

let edition = OSS
let loading = null

export function loadEdition() {
  loading ??= setup_api
    .checkInstalled()
    .then((res) => {
      edition = res.data.edition === SAAS ? SAAS : OSS
      return edition
    })
    .catch(() => edition)

  return loading
}

export function isSaaS() {
  return edition === SAAS
}
