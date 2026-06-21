<template>
  <div class="flex min-h-dvh items-center">
    <Card class="w-full">
      <CardHeader>
        <CardTitle>Вход в аккаунт</CardTitle>
        <CardDescription>Заполните почту и пароль, чтобы войти в аккаунт</CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit="onSubmit" class="space-y-4">
          <FormField v-slot="{ componentField }" name="email">
            <FormItem>
              <FormLabel>Введите электронную почту</FormLabel>
              <FormControl>
                <Input type="email" placeholder="example@example.com" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="password">
            <FormItem>
              <FormLabel>Введите пароль</FormLabel>
              <FormControl>
                <Input type="password" placeholder="Пароль" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <Button class="w-full" type="submit" :disabled="form.isSubmitting.value">
            <Spinner v-show="form.isSubmitting.value" />
            Войти
          </Button>

          <p class="text-muted-foreground text-center text-sm">
            Ещё нет аккаунта?
            <router-link to="/register" class="text-foreground underline-offset-2 hover:underline">
              Зарегистрироваться
            </router-link>
          </p>
        </form>
      </CardContent>
    </Card>
  </div>
</template>

<script lang="ts" setup>
import { Button } from '@/components/ui/button'
import { FormField, FormControl, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Spinner } from '@/components/ui/spinner'
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { useRouter } from 'vue-router'
import { toast } from 'vue-sonner'
import type { ErrorResponse } from '@/types.ts'

const formSchema = toTypedSchema(
  z.object({
    email: z
      .string({ message: 'Введите вашу почту' })
      .email({ message: 'Некорректная почта' })
      .min(1, { message: 'Введите вашу почту' })
      .max(100),
    password: z
      .string({ message: 'Введите пароль' })
      .min(8, { message: 'Минимальная длина пароля ― 8 символов' })
      .max(100),
  }),
)

const router = useRouter()
const form = useForm({
  validationSchema: formSchema,
})

const onSubmit = form.handleSubmit(async (values) => {
  try {
    const response = await fetch('/api/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(values),
    })

    if (!response.ok) {
      const error = (await response.json()) as ErrorResponse
      toast.error('При выполнении запроса произошла ошибка')
      return
    }

    const data = (await response.json()) as { token: string }
    localStorage.setItem('token', data.token)
    router.push({ name: 'main' })
    toast.success('Вы успешно вошли в аккаунт')
  } catch {
    toast.error('При выполнении произошла неизвестная ошибка')
  }
})
</script>
