<template>
  <div>
    <h1>This is main page</h1>

    <div>{{ user }}</div>

    <router-link to="/logout">Выйти</router-link>
  </div>
</template>

<script lang="ts" setup>
import { useAuthStore } from '@/stores/auth.store.ts'

const { user } = useAuthStore()
</script>
