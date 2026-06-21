<template>
  <div>
    <header class="flex items-center justify-between py-5">
      <h1 class="text-lg font-bold select-none">Главная</h1>
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <button>
            <Avatar>
              <AvatarFallback>{{ avatarFallback }}</AvatarFallback>
            </Avatar>
          </button>
        </DropdownMenuTrigger>
        <DropdownMenuContent class="w-50" align="end">
<!--          <DropdownMenuGroup>-->
<!--            <DropdownMenuItem> Profile </DropdownMenuItem>-->
<!--          </DropdownMenuGroup>-->
<!--          <DropdownMenuSeparator />-->
          <router-link to="/logout">
            <DropdownMenuItem class="cursor-pointer"> Выйти </DropdownMenuItem>
          </router-link>
        </DropdownMenuContent>
      </DropdownMenu>
    </header>

    <div class="flex flex-col gap-4">
      <Card>
        <CardHeader>
          <CardTitle> Ваша активность </CardTitle>
        </CardHeader>
        <CardContent> test </CardContent>
      </Card>

      <Card v-if="user">
        <CardHeader>
          <CardTitle>Ваш профиль</CardTitle>
          <CardDescription>Карточка с вашими данными</CardDescription>
        </CardHeader>
        <CardContent>
          <ul>
            <li class="flex justify-between">
              <span class="font-medium">ID:</span>
              <span>{{ user.id }}</span>
            </li>
            <li class="flex justify-between">
              <span class="font-medium">UUID:</span>
              <span>{{ user.uuid }}</span>
            </li>
            <li class="flex justify-between">
              <span class="font-medium">Электронная почта:</span>
              <span>{{ user.email }}</span>
            </li>
            <li class="flex justify-between">
              <span class="font-medium">Имя:</span>
              <span>{{ user.first_name }}</span>
            </li>
            <li class="flex justify-between">
              <span class="font-medium">Фамилия:</span>
              <span>{{ user.last_name }}</span>
            </li>
            <li class="flex justify-between">
              <span class="font-medium">Изменен:</span>
              <span>{{ user.updated_at }}</span>
            </li>
            <li class="flex justify-between">
              <span class="font-medium">Создан:</span>
              <span>{{ user.created_at }}</span>
            </li>
          </ul>
        </CardContent>
      </Card>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useAuthStore } from '@/stores/auth.store.ts'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Avatar, AvatarFallback } from '@/components/ui/avatar'
import { computed } from 'vue'
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from '@/components/ui/card'

const { user } = useAuthStore()

const avatarFallback = computed(() => {
  if (user) {
    const firstNameLetter = user.first_name[0]
    const firstSurnameLetter = user.last_name[0]
    if (firstNameLetter && firstSurnameLetter) {
      return firstNameLetter.toUpperCase() + firstSurnameLetter.toUpperCase()
    }
  }
  return '...'
})
</script>
