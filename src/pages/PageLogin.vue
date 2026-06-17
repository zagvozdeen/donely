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

          <Button class="w-full" type="submit">Войти</Button>

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
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'

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

const form = useForm({
  validationSchema: formSchema,
})

const onSubmit = form.handleSubmit((values, ctx) => {
  console.log(values, ctx)
})
</script>
