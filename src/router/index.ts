import { createRouter, createWebHistory } from 'vue-router'
import PageMain from '@/pages/PageMain.vue'
import PageLogin from '@/pages/PageLogin.vue'
import PageRegister from '@/pages/PageRegister.vue'
import PageLogout from '@/pages/PageLogout.vue'
import { useAuthStore } from '@/stores/auth.store.ts'

export const createRoutes = () => {
  const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
      {
        component: PageMain,
        name: 'main',
        path: '/',
      },
      {
        component: PageLogin,
        name: 'login',
        path: '/login',
      },
      {
        component: PageRegister,
        name: 'register',
        path: '/register',
      },
      {
        component: PageLogout,
        name: 'logout',
        path: '/logout',
      },
    ],
  })

  router.beforeEach(async (to) => {
    const authStore = useAuthStore()

    const token = localStorage.getItem('token')
    const isGuestPage = to.name === 'login' || to.name === 'register'

    if (!token && !isGuestPage) {
      return '/login'
    }

    if (token && isGuestPage) {
      return '/'
    }

    if (token && !authStore.user) {
      await authStore.getMe()
      if (!authStore.user) {
        return '/login'
      }
    }
  })

  return router
}
