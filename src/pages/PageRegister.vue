<template>
  <div class="flex min-h-dvh items-center">
    <Card class="w-full">
      <CardHeader>
        <CardTitle>Регистрация</CardTitle>
        <CardDescription>Заполните данные, чтобы создать аккаунт</CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit="onSubmit" class="space-y-4">
          <FormField v-slot="{ componentField }" name="first_name">
            <FormItem>
              <FormLabel>Введите имя</FormLabel>
              <FormControl>
                <Input type="text" placeholder="Иван" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="last_name">
            <FormItem>
              <FormLabel>Введите фамилию</FormLabel>
              <FormControl>
                <Input type="text" placeholder="Иванов" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

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

          <FormField v-slot="{ componentField }" name="password_confirmation">
            <FormItem>
              <FormLabel>Введите пароль ещё раз</FormLabel>
              <FormControl>
                <Input type="password" placeholder="Подтверждение пароля" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="consent" type="checkbox">
            <FormItem>
              <div class="flex gap-2">
                <FormControl>
                  <Checkbox v-bind="componentField" />
                </FormControl>
                <FormLabel class="text-muted-foreground block cursor-pointer text-sm font-normal">
                  Я соглашаюсь на обработку
                  <a href="#" class="text-foreground underline underline-offset-2">
                    персональных данных
                  </a>
                  и принимаю
                  <a href="#" class="text-foreground underline underline-offset-2">
                    условия использования
                  </a>
                </FormLabel>
              </div>
              <FormMessage />
            </FormItem>
          </FormField>

          <Button class="w-full" type="submit">Зарегистрироваться</Button>

          <p class="text-muted-foreground text-center text-sm">
            Уже есть аккаунт?
            <router-link to="/login" class="text-foreground underline-offset-2 hover:underline">
              Войти
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
import { Checkbox } from '@/components/ui/checkbox'

const formSchema = toTypedSchema(
  z
    .object({
      first_name: z
        .string({ message: 'Введите ваше имя' })
        .min(1, { message: 'Введите ваше имя' })
        .max(100),
      last_name: z
        .string({ message: 'Введите вашу фамилию' })
        .min(1, { message: 'Введите вашу фамилию' })
        .max(100),
      email: z
        .string({ message: 'Введите вашу почту' })
        .email({ message: 'Некорректная почта' })
        .min(1, { message: 'Введите вашу почту' })
        .max(100),
      password: z
        .string({ message: 'Введите пароль' })
        .min(8, { message: 'Минимальная длина пароля ― 8 символов' })
        .max(100),
      password_confirmation: z
        .string({ message: 'Подтвердите ваш пароль' })
        .min(8, { message: 'Минимальная длина пароля ― 8 символов' })
        .max(100),
      consent: z.literal(true, { message: 'Необходимо ваше согласие' }),
    })
    .refine((data) => data.password === data.password_confirmation, {
      message: 'Пароли не совпадают',
      path: ['password_confirmation'],
    }),
)

const form = useForm({
  validationSchema: formSchema,
})

const onSubmit = form.handleSubmit((values, ctx) => {
  console.log(values, ctx)
})
</script>
