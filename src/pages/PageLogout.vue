<template>
  <div class="flex min-h-dvh items-center justify-center gap-2 select-none">
    <Spinner />
    <span class="text-muted-foreground text-sm">Выход из аккаунта...</span>
  </div>
</template>

<script lang="ts" setup>
import { Spinner } from '@/components/ui/spinner'
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { toast } from 'vue-sonner'

const router = useRouter()

onMounted(() => {
  localStorage.removeItem('token')
  router.push({ name: 'login' })
  toast.info('Вы вышли из аккаунта')
})
</script>
