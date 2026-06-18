import '@/assets/styles.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from '@/App.vue'
import { createRoutes } from '@/router'

createApp(App).use(createPinia()).use(createRoutes()).mount('#app')
