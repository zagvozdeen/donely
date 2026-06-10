import { defineConfig } from 'oxlint'

export default defineConfig({
  plugins: ['vue', 'typescript'],
  env: {
    browser: true,
    node: true,
  },
})
