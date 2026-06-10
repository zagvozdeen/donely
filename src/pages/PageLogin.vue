<template>
  <div class="flex min-h-dvh w-full items-center justify-center">
    <Form
        v-slot="$form"
      @submit.prevent="onSubmit"
      class="bg-surface-900 border-surface my-8 w-full rounded-2xl border p-4"
    >
      <h1 class="mb-4 w-full text-center text-xl">Войти в аккаунт</h1>
      <!--      <n-form-item label="Почта" path="email">-->
      <div class="flex flex-col gap-2">
        <FormField name="email-field" v-slot="$field" class="flex flex-col gap-1">
          <label for="email-field">Электронная почта</label>
          <InputText
            id="email-field"
            v-model="form.email"
            placeholder="Введите электронную почту"
            fluid
          />
        </FormField>

        <!--      <FloatLabel variant="on">-->
        <!--        <InputText id="email-field" v-model="form.email" type="email" fluid />-->
        <!--        <label for="email-field">Почта</label>-->
        <!--      </FloatLabel>-->
        <!--      </n-form-item>-->

        <!--      <n-form-item label="Пароль" path="password">-->
        <FormField name="password-field" v-slot="$field" class="flex flex-col gap-1">
          <label for="password-field">Пароль</label>
          <Password
            id="password-field"
            v-model="form.password"
            :feedback="false"
            placeholder="Введите пароль"
            toggle-mask
            fluid
          />
        </FormField>

        <!--      <FloatLabel variant="on">-->
        <!--        <Password id="password-field" v-model="form.password" :feedback="false" toggle-mask fluid />-->
        <!--        <label for="password-field">Пароль</label>-->
        <!--      </FloatLabel>-->
        <!--      </n-form-item>-->

        <Button type="submit" label="Войти" fluid />

        <!--      <n-form-item :show-feedback="false" :show-label="false">-->
        <!--        <n-button attr-type="submit" type="primary" class="flex-1"> Войти </n-button>-->
        <!--      </n-form-item>-->

        <span class="text-center"
          >Ещё нет аккаунта?
          <router-link :to="{ name: 'register' }" class="underline">
            Зарегистрироваться
          </router-link></span
        >
      </div>
    </Form>
  </div>
</template>
<script lang="ts" setup>
import InputText from 'primevue/inputtext'
import FloatLabel from 'primevue/floatlabel'
import Password from 'primevue/password'
import Button from 'primevue/button'
import { Form, FormField } from '@primevue/forms'

import { reactive } from 'vue'

const form = reactive({
  email: '',
  password: '',
})

// import { NForm, NFormItem, NInput, NButton, FormInst, FormRules, useMessage } from 'naive-ui'
// import { reactive, ref } from 'vue'
// import { useGetMe, useLogin } from '@/composables/useAuthStore.ts'
// import { useRouter } from 'vue-router'
// import { PageExpose } from '@/types.ts'
// import { Submitter } from '@/composables/useForm.ts'
//
// defineExpose<PageExpose>({
//   title: 'Вход',
// })

// const router = useRouter()
// const message = useMessage()
// const submitter = new Submitter()

// const formRef = ref<FormInst>()
// const formValue = reactive({
//   email: null as string | null,
//   password: null as string | null,
// })
// const formRules: FormRules = {
//   email: {
//     required: true,
//     type: 'email',
//     message: 'Введите почту',
//   },
//   password: {
//     required: true,
//     message: 'Введите пароль',
//   },
// }

// const clearForm = () => {
// formValue.email = null
// formValue.password = null
// }

const onSubmit = () => {
  // submitter.submit(formRef.value, async () => {
  //   await useLogin(formValue)
  //   await useGetMe()
  //
  //   clearForm()
  //
  //   message.success('Вы успешно вошли в систему')
  //
  //   await router.push({ name: 'main' })
  // })
}
</script>
