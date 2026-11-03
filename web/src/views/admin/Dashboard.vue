<template>
  <div class="min-h-screen bg-theme-bg">
    <AppNav />

    <main class="w-full px-4 sm:px-6 lg:px-8 py-8">
      <header class="mb-8">
        <h1 class="text-2xl sm:text-3xl font-semibold text-theme-text tracking-tight">
          {{ $t('dashboard.title') }}
        </h1>
        <p class="mt-1 text-sm text-theme-textLight">
          {{ $t('dashboard.subtitle') }}
        </p>
        <p v-if="user?.email" class="mt-2 text-sm text-theme-text">
          {{ $t('dashboard.welcome_back') }}
          <span class="font-medium">{{ user.name || user.email }}</span>
        </p>
      </header>

      <section class="mb-8 grid grid-cols-1 gap-6 lg:grid-cols-2">
        <div class="section overflow-hidden !p-0">
          <div class="border-b border-theme-border px-6 py-4">
            <div class="flex flex-wrap items-center justify-between gap-2">
              <div>
                <h2 class="text-lg font-semibold text-theme-text">
                  {{ $t('dashboard.attention_title') }}
                </h2>
                <p class="mt-0.5 text-sm text-theme-textLight">
                  {{ $t('dashboard.attention_subtitle') }}
                </p>
              </div>
            </div>
          </div>

          <ul class="divide-y divide-theme-border max-h-[28rem] overflow-y-auto">
            <li v-for="item in attentionItems" :key="item.id" class="flex gap-4 px-6 py-4">
              <span
                class="mt-1.5 h-2 w-2 shrink-0 rounded-full"
                :class="attentionDotClass(item.level)"
                aria-hidden="true"
              />
              <div class="min-w-0 flex-1">
                <p class="text-sm font-medium text-theme-text">{{ item.title }}</p>
                <p class="mt-0.5 text-sm text-theme-textLight leading-relaxed">
                  {{ item.description }}
                </p>
                <router-link
                  v-if="item.to"
                  :to="item.to"
                  class="mt-2 inline-flex text-sm font-medium text-primary-800 hover:underline"
                >
                  {{ item.action }}
                </router-link>
              </div>
            </li>
          </ul>
        </div>

        <div class="section overflow-hidden !p-0">
          <div class="border-b border-theme-border px-6 py-4">
            <h2 class="text-lg font-semibold text-theme-text">
              {{ $t('dashboard.whats_new_title') }}
            </h2>
            <p class="mt-0.5 text-sm text-theme-textLight">
              {{ $t('dashboard.whats_new_subtitle') }}
            </p>
          </div>
          <div class="px-6 py-5">
            <p class="text-xs font-medium uppercase tracking-wide text-theme-textLight">
              {{ $t('dashboard.whats_new_tip_label') }}
            </p>
            <p class="mt-2 text-base font-medium text-theme-text">
              {{ $t('dashboard.whats_new_tip_title') }}
            </p>
            <p class="mt-1 text-sm leading-relaxed text-theme-textLight">
              {{ $t('dashboard.whats_new_tip_body') }}
            </p>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import AppNav from '@/comp/AppNav.vue'
import { kv_api, workspace_access_key_api, workspace_invite_api } from '@/api'
import { user } from '@/lib/auth'
import { useWorkspaceContext } from '@/lib/permission'

const { t } = useI18n()

const { currentWorkspace, canManage } = useWorkspaceContext()

const accessKeysTotal = ref(null)
const pendingInvites = ref(null)
const kvTotal = ref(null)

async function loadSetupSignals() {
  if (canManage.value) {
    try {
      const res = await workspace_access_key_api.list(currentWorkspace.id, {
        limit: 1,
        offset: 0,
      })
      accessKeysTotal.value = res.data._meta?.total ?? (res.data.keys || []).length
    } catch (_) {
      // The key check is skipped when the keys cannot be read.
    }

    try {
      const res = await workspace_invite_api.list(currentWorkspace.id, {
        limit: 100,
        offset: 0,
      })
      pendingInvites.value = (res.data.invites || []).filter(
        (invite) => invite.status === 'pending'
      ).length
    } catch (_) {
      // The invite check is skipped when the invites cannot be read.
    }
  }

  try {
    const res = await kv_api.list(currentWorkspace.id, { limit: 1, offset: 0 })
    kvTotal.value = res.data._meta?.total ?? (res.data.items || []).length
  } catch (_) {
    // The knowledge check is skipped when the keys cannot be read.
  }

}

onMounted(() => {
  loadSetupSignals()
})

const attentionReady = computed(
  () =>
    kvTotal.value !== null &&
    (!canManage.value || (accessKeysTotal.value !== null && pendingInvites.value !== null))
)

const attentionItems = computed(() => {
  const items = []

  if (canManage.value) {
    if (accessKeysTotal.value === 0) {
      items.push({
        id: 'key',
        level: 'info',
        title: t('dashboard.attention_no_key_title'),
        description: t('dashboard.attention_no_key_desc'),
        action: t('dashboard.attention_no_key_action'),
        to: { name: 'Settings' },
      })
    }

    if (pendingInvites.value > 0) {
      items.push({
        id: 'invites',
        level: 'info',
        title: t('dashboard.attention_invites_title'),
        description: t('dashboard.attention_invites_desc', { count: pendingInvites.value }),
        action: t('dashboard.attention_invites_action'),
        to: { name: 'Members' },
      })
    }
  }

  if (kvTotal.value === 0) {
    items.push({
      id: 'kb',
      level: 'info',
      title: t('dashboard.attention_empty_kb_title'),
      description: t('dashboard.attention_empty_kb_desc'),
      action: t('dashboard.attention_empty_kb_action'),
      to: { name: 'KV' },
    })
  }

  if (!items.length && attentionReady.value) {
    items.push({
      id: 'clear',
      level: 'info',
      title: t('dashboard.attention_all_clear_title'),
      description: t('dashboard.attention_all_clear_desc'),
    })
  }

  return items
})

function attentionDotClass(level) {
  return (
    {
      error: 'bg-red-500',
      warn: 'bg-amber-500',
      info: 'bg-primary-500',
    }[level] || 'bg-theme-textLight'
  )
}
</script>
