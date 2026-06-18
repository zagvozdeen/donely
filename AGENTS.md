# Repository Guidelines

## Project Structure & Module Organization

This is a full-stack Vue 3, Vite, TypeScript frontend with a Go backend.
Frontend application code lives in `src/`. Entry points are `src/main.ts` and
`src/App.vue`. Route definitions are in `src/router/index.ts`, page-level views
are in `src/pages/`, Pinia stores are in `src/stores/`, shared helpers are in
`src/lib/`, and global styling is in `src/assets/styles.css`. shadcn-vue UI
components are generated under `src/components/ui/`; keep reusable UI primitives
there and page-specific layout in `src/pages/`. Static public assets belong in
`public/`.

The Go module is `github.com/zagvozdeen/donely` and targets Go 1.26. Backend
entrypoint code lives in `cmd/main.go`. Internal application packages live under
`internal/`: HTTP routes and auth in `internal/api/`, configuration loading in
`internal/config/`, PostgreSQL setup and embedded Goose migrations in
`internal/db/`, logging in `internal/logger/`, and persistence code in
`internal/store/`. Shared backend primitives that can be imported outside
`internal/` live in `pkg/`. Database migrations belong in
`internal/db/migrations/`.

## Build, Test, and Development Commands

- `npm run dev`: start the Vite development server.
- `npm run build`: run `vue-tsc -b` type checking, then create a production Vite
  build.
- `npm run lint`: format with `oxfmt` and apply lint fixes with `oxlint --fix`.
- `docker compose up -d`: start the local PostgreSQL service used by the Go
  API. Values are read from `.env`.
- `go run ./cmd`: start the Go API server on port `8000`; it loads `.env`,
  connects to PostgreSQL, and runs embedded migrations on startup.
- `go test ./...`: run backend package tests and compile checks.
- `go fmt ./...`: format Go files. Run this before submitting Go changes.
- `go mod tidy`: update module metadata after adding or removing Go
  dependencies.

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

Use idiomatic Go formatting from `gofmt`/`go fmt`. Keep package names lowercase
and short. Keep command wiring in `cmd/`, application-only packages in
`internal/`, and reusable public backend packages in `pkg/`. Pass
`context.Context` explicitly to database and request-scoped backend work, and
wrap returned errors with `%w` when adding context.

## Testing Guidelines

No frontend test suite is present yet. When adding frontend tests, add a package
script and document it here. Co-locate component tests near the component or use
a dedicated `src/__tests__/` directory. Name tests after the behavior being
verified, not the implementation detail.

Go test files are not present yet, but `go test ./...` should be used for
backend compile checks and any future backend tests. Add Go tests next to the
package being tested, using `_test.go` files and behavior-focused test names.

## Commit & Pull Request Guidelines

Recent commits use Conventional Commit-style prefixes, especially `feat:`. Keep
commit messages short and imperative, for example `feat: add register form
validation` or `fix: handle checkbox consent state`.

Pull requests should include a concise description, the commands run for
verification, and screenshots or short recordings for UI changes. Link related
issues when applicable and call out any new environment or configuration needs.

## Agent-Specific Instructions

Before editing, check for existing project patterns, generated shadcn-vue
components, and existing Go package boundaries. Keep changes focused, avoid
unrelated refactors, and do not overwrite local user changes.
