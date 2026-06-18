import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { User } from '@/types.ts'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const isUserLoading = ref(false)

  const getMe = () => {
    if (user.value) return

    const token = localStorage.getItem('token')
    if (!token) return

    isUserLoading.value = true

    return fetch('/api/me', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => res.json() as Promise<User>)
      .then((data) => {
        user.value = data
      })
      .catch(() => {
        user.value = null
        localStorage.removeItem('token')
      })
      .finally(() => {
        isUserLoading.value = false
      })
  }

  return {
    user,
    isUserLoading,
    getMe,
  }
})
