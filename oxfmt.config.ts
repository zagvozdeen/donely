import { defineConfig } from 'oxfmt'

export default defineConfig({
  arrowParens: 'always',
  bracketSameLine: false,
  bracketSpacing: true,
  endOfLine: 'lf',
  htmlWhitespaceSensitivity: 'css',
  ignorePatterns: ['dist/**', 'node_modules/**'],
  insertFinalNewline: true,
  jsxSingleQuote: true,
  objectWrap: 'preserve',
  printWidth: 100,
  proseWrap: 'preserve',
  quoteProps: 'as-needed',
  semi: false,
  singleQuote: true,
  sortPackageJson: {
    sortScripts: true,
  },
  sortTailwindcss: {
    attributes: [':class'],
    functions: ['clsx', 'cn', 'cva'],
    preserveDuplicates: false,
    preserveWhitespace: false,
  },
  tabWidth: 2,
  trailingComma: 'all',
  useTabs: false,
  vueIndentScriptAndStyle: false,
})
