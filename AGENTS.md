# Repository Guidelines

## Project Structure & Module Organization

This is a Vue 3, Vite, TypeScript application. Application code lives in `src/`.
Entry points are `src/main.ts` and `src/App.vue`. Route definitions are in
`src/router/index.ts`, page-level views are in `src/pages/`, Pinia stores are in
`src/stores/`, shared helpers are in `src/lib/`, and global styling is in
`src/assets/styles.css`. shadcn-vue UI components are generated under
`src/components/ui/`; keep reusable UI primitives there and page-specific layout in
`src/pages/`. Static public assets belong in `public/`.

## Build, Test, and Development Commands

- `npm run dev`: start the Vite development server.
- `npm run build`: run `vue-tsc -b` type checking, then create a production Vite
  build.
- `npm run lint`: format with `oxfmt` and apply lint fixes with `oxlint --fix`.

There is currently no `npm test` script or test framework configured. Do not
refer to tests as passing unless a test command has been added and run.

## Coding Style & Naming Conventions

Use TypeScript and Vue single-file components. Follow the formatter settings in
`oxfmt.config.ts`: 2-space indentation, single quotes, no semicolons, LF line
endings, trailing commas, and a 100-character print width. Run `npm run lint`
before submitting changes.

Use PascalCase for Vue components and page files, for example
`PageRegister.vue`. Use camelCase for variables, functions, and composables.
Prefer the configured `@/` aliases from `components.json`, such as
`@/components/ui/button` and `@/lib/utils`.

## Testing Guidelines

No test suite is present yet. When adding tests, add a package script and document
it here. Co-locate component tests near the component or use a dedicated
`src/__tests__/` directory. Name tests after the behavior being verified, not the
implementation detail.

## Commit & Pull Request Guidelines

Recent commits use Conventional Commit-style prefixes, especially `feat:`. Keep
commit messages short and imperative, for example `feat: add register form
validation` or `fix: handle checkbox consent state`.

Pull requests should include a concise description, the commands run for
verification, and screenshots or short recordings for UI changes. Link related
issues when applicable and call out any new environment or configuration needs.

## Agent-Specific Instructions

Before editing, check for existing project patterns and generated shadcn-vue
components. Keep changes focused, avoid unrelated refactors, and do not overwrite
local user changes.
