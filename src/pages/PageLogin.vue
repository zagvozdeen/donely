<template>
  <div class="flex min-h-dvh w-full items-center justify-center">
    <Form
      v-slot="$form"
      :resolver="resolver"
      @submit="onSubmit"
      :initial-values="form"
      class="bg-surface-900 border-surface my-8 flex w-full flex-col gap-4 rounded-2xl border p-4"
    >
      <h1 class="w-full text-center text-xl">Войти в аккаунт</h1>

      <FormField name="email" v-slot="$field" class="flex flex-col gap-1">
        <label for="email-field">Электронная почта</label>
        <InputText
          id="email-field"
          v-model="form.email"
          placeholder="Введите электронную почту"
          fluid
        />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
          $field.error?.message
        }}</Message>
      </FormField>

      <FormField name="password" v-slot="$field" class="flex flex-col gap-1">
        <label for="password-field">Пароль</label>
        <Password
          id="password-field"
          v-model="form.password"
          placeholder="Введите пароль"
          :feedback="false"
          toggle-mask
          fluid
        />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
          $field.error?.message
        }}</Message>
      </FormField>

      <Button type="submit" label="Войти" fluid />

      <span class="text-center"
        >Ещё нет аккаунта?
        <router-link :to="{ name: 'register' }" class="underline">
          Зарегистрироваться
        </router-link></span
      >
    </Form>
  </div>
</template>
<script lang="ts" setup>
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import Button from 'primevue/button'
import Message from 'primevue/message'
import { Form, FormField } from '@primevue/forms'
import { z } from 'zod'
import { zodResolver } from '@primevue/forms/resolvers/zod'
import { reactive } from 'vue'
import {useToast} from "primevue";

const toast = useToast()

const form = reactive({
  email: '',
  password: '',
})

const resolver = zodResolver(
  z.object({
    email: z.email('Введите электронный адрес'),
    password: z.string().min(1, 'Введите пароль'),
  }),
)

const onSubmit = ({ valid, values }) => {
  if (valid) {

    toast.add({severity: "info", summary: 'fsdfsd'})
  }
}
</script>
