import { createRouter, createWebHistory } from 'vue-router'
import PageMain from '@/pages/PageMain.vue'
import PageLogin from '@/pages/PageLogin.vue'
import PageRegister from '@/pages/PageRegister.vue'

export default createRouter({
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
  ],
})
