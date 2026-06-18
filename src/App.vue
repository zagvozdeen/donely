<template>
  <main class="mx-auto min-h-dvh max-w-md px-4">
    <router-view />
    <Toaster position="top-center" theme="system" richColors />
  </main>
</template>

<script setup lang="ts">
import 'vue-sonner/style.css'
import { Toaster } from '@/components/ui/sonner'
</script>
