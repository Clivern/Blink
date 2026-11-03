import { createRouter, createWebHistory } from 'vue-router'
import { useShopSession } from '@/shop/session'
import { loadUserFromStorage, loadWorkspaceFromStorage } from '@/utils/storage'
import { canManageWorkspace } from '@/lib/permission'
import { loadEdition, isSaaS } from '@/lib/edition'

const routes = [
  {
    path: '/admin',
    redirect: '/admin/dashboard',
  },
  {
    path: '/admin/login',
    name: 'Login',
    component: () => import('@/views/admin/Login.vue'),
    meta: { requiresGuest: true, title: 'Login', description: 'Login to platform' },
  },
  {
    path: '/admin/setup',
    name: 'Setup',
    component: () => import('@/views/admin/Setup.vue'),
    meta: { requiresGuest: true, title: 'Setup', description: 'Setup platform' },
  },
  {
    path: '/admin/dashboard',
    name: 'Dashboard',
    component: () => import('@/views/admin/Dashboard.vue'),
    meta: {
      requiresAuth: true,
      requiresWorkspace: true,
      title: 'Dashboard',
      description: 'Platform dashboard',
    },
  },
  {
    path: '/admin/select-workspace',
    name: 'SwitchWorkspaces',
    component: () => import('@/views/admin/SelectWorkspace.vue'),
    meta: {
      requiresAuth: true,
      title: 'Switch Workspaces',
      description: 'Choose or create a workspace',
    },
  },
  {
    path: '/admin/workspaces',
    name: 'Workspaces',
    component: () => import('@/views/admin/Workspaces.vue'),
    meta: {
      requiresAuth: true,
      requiresWorkspace: true,
      title: 'Workspaces',
      description: 'Workspaces',
    },
  },
  {
    path: '/admin/profile',
    name: 'Profile',
    component: () => import('@/views/admin/Profile.vue'),
    meta: { requiresAuth: true, title: 'Profile', description: 'Your account profile' },
  },
  {
    path: '/admin/kv',
    name: 'KV',
    component: () => import('@/views/admin/KV.vue'),
    meta: {
      requiresAuth: true,
      requiresWorkspace: true,
      title: 'KV',
      description: 'Workspace key-value config',
    },
  },
  {
    path: '/admin/members',
    name: 'Members',
    component: () => import('@/views/admin/Members.vue'),
    meta: {
      requiresAuth: true,
      requiresWorkspace: true,
      requiresWorkspaceAdmin: true,
      title: 'Members',
      description: 'Members',
    },
  },
  {
    path: '/admin/settings',
    name: 'Settings',
    component: () => import('@/views/admin/Settings.vue'),
    meta: {
      requiresAuth: true,
      requiresWorkspace: true,
      requiresWorkspaceAdmin: true,
      title: 'Settings',
      description: 'Settings',
    },
  },
  {
    path: '/',
    component: () => import('@/views/shop/ShopLayout.vue'),
    meta: { shop: true, title: 'Shop' },
    children: [
      {
        path: '',
        name: 'ShopHome',
        component: () => import('@/views/shop/ShopHome.vue'),
        meta: { title: 'Shop' },
      },
      {
        path: 'c/:slug',
        name: 'ShopCategory',
        component: () => import('@/views/shop/ShopCategory.vue'),
        meta: { title: 'Category' },
      },
      {
        path: 'p/:sku',
        name: 'ShopProduct',
        component: () => import('@/views/shop/ShopProduct.vue'),
        meta: { title: 'Product' },
      },
      {
        path: 'search',
        name: 'ShopSearch',
        component: () => import('@/views/shop/ShopSearch.vue'),
        meta: { title: 'Search' },
      },
      {
        path: 'cart',
        name: 'ShopCart',
        component: () => import('@/views/shop/ShopCart.vue'),
        meta: { title: 'Basket' },
      },
      {
        path: 'checkout',
        name: 'ShopCheckout',
        component: () => import('@/views/shop/ShopCheckout.vue'),
        meta: { title: 'Checkout' },
        beforeEnter: (to) => {
          const session = useShopSession()
          if (!session.user) return { name: 'ShopLogin', query: { redirect: to.fullPath } }
        },
      },
      {
        path: 'order/:id',
        name: 'ShopOrder',
        component: () => import('@/views/shop/ShopOrder.vue'),
        meta: { title: 'Order' },
      },
      {
        path: 'account',
        name: 'ShopAccount',
        component: () => import('@/views/shop/ShopAccount.vue'),
        meta: { title: 'You' },
        beforeEnter: () => {
          const session = useShopSession()
          if (!session.user) return { name: 'ShopLogin' }
        },
      },
      {
        path: 'login',
        name: 'ShopLogin',
        component: () => import('@/views/shop/ShopLogin.vue'),
        meta: { title: 'Sign in' },
        beforeEnter: () => {
          const session = useShopSession()
          if (session.user) return { name: 'ShopAccount' }
        },
      },
      {
        path: 'info/:slug',
        name: 'ShopPage',
        component: () => import('@/views/shop/ShopPage.vue'),
        meta: { title: 'Blink' },
      },
    ],
  },
  {
    path: '/shop',
    redirect: '/',
  },
  {
    path: '/shop/:pathMatch(.*)*',
    redirect: (to) => {
      const rest = to.params.pathMatch
      const suffix = Array.isArray(rest) ? rest.join('/') : rest
      return suffix ? `/${suffix}` : '/'
    },
  },
  {
    path: '/admin/404',
    name: 'NotFound',
    component: () => import('@/views/admin/NotFound.vue'),
    meta: { title: 'Page not found' },
  },
  {
    path: '/admin/500',
    name: 'ServerError',
    component: () => import('@/views/admin/ServerError.vue'),
    meta: { title: 'Server error' },
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFoundCatchAll',
    component: () => import('@/views/admin/NotFound.vue'),
    meta: { title: 'Page not found' },
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, _from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    }
    if (to.hash) {
      return { el: to.hash, behavior: 'smooth' }
    }
    return { top: 0 }
  },
})

function isSafeRedirect(redirect) {
  return typeof redirect === 'string' && redirect.startsWith('/') && !redirect.startsWith('//')
}

// Navigation guard - reads auth state directly from storage
router.beforeEach(async (to) => {
  await loadEdition()

  const currentUser = loadUserFromStorage()
  const currentWorkspace = loadWorkspaceFromStorage()
  const isAuthenticated = !!currentUser

  if (to.meta.requiresSaaS && !isSaaS()) {
    return '/admin/404'
  }
  if (to.meta.requiresAuth && !isAuthenticated) {
    return { path: '/admin/login', query: { redirect: to.fullPath } }
  }
  if (to.meta.requiresGuest && isAuthenticated) {
    const redirect = to.query.redirect
    return isSafeRedirect(redirect) ? redirect : '/admin/select-workspace'
  }
  if (to.meta.requiresWorkspace && !currentWorkspace) {
    return '/admin/select-workspace'
  }
  if (to.meta.requiresWorkspaceAdmin && !canManageWorkspace(currentUser, currentWorkspace)) {
    return '/admin/404'
  }
})

export default router
