---
# @config-manager:start zdp-frontmatter
inclusion: fileMatch
fileMatchPattern: "**/*.{js,jsx,ts,tsx}"
# @config-manager:end zdp-frontmatter
---

<!-- @config-manager:start zdp -->
# JavaScript/TypeScript Zero Diagnostics

All JavaScript and TypeScript source files must be free of diagnostic errors and warnings. Code is not considered complete until `getDiagnostics` reports zero issues for every file touched during a change.

## Philosophy

* **Small, surgical edits** are strongly preferred over wide-scale refactoring. Without a larger spec to guide the work, keep changes minimal and focused on the specific diagnostic being resolved.
* **Fix the problem, not the symptom.** Strongly favor resolving the root cause over suppressing the error. Suppression is a last resort reserved for cases where a fix is genuinely impossible or would introduce worse problems.
* The goal is **zero remaining diagnostic issues** when running `oxfmt --write --no-error-on-unmatched-pattern .`, `oxlint --import-plugin --jsdoc-plugin --jsx-a11y-plugin --react-perf-plugin --fix --no-error-on-unmatched-pattern .`, and `tsc --skipLibCheck --noEmit **/*.ts` from the project root.

## Default to using bun instead of node.js

* Use `bun <file>` instead of `node <file>` or `ts-node <file>`.
* Use `bun test` instead of `jest` or `vitest`.
* Use `bun build <file.html|file.ts|file.css>` instead of `webpack` or `esbuild`.
* Use `bun install` instead of `npm install` or `yarn install` or `pnpm install`.
* Use `bun run <script>` instead of `npm run <script>` or `yarn run <script>` or `pnpm run <script>`.
* Use `bunx <package> <command>` instead of `npx <package> <command>`.
* Bun automatically loads .env, so don't use dotenv.

## APIs

* `Bun.serve()` supports WebSockets, HTTPS, and routes. Don't use `express`.
* `bun:sqlite` for SQLite. Don't use `better-sqlite3`.
* `Bun.redis` for Redis. Don't use `ioredis`.
* `Bun.sql` for Postgres. Don't use `pg` or `postgres.js`.
* `WebSocket` is built-in. Don't use `ws`.
* Prefer `Bun.file` over `node:fs`'s readFile/writeFile
* Bun.$`ls` instead of execa.

## Verification workflow

1. After editing any `.js`, `.jsx`, `.ts`, or `.tsx` file, run `getDiagnostics` on that file.
2. If diagnostics are reported, fix every one before moving on.
3. After all fixes, run `oxfmt --write --no-error-on-unmatched-pattern .` in the directory where `.oxfmtrc.jsonc` exists, `oxlint --import-plugin --jsdoc-plugin --jsx-a11y-plugin --react-perf-plugin --fix --no-error-on-unmatched-pattern .` in the directory where `.oxlintrc.jsonc` exists to auto-fix what the linter can and surface anything remaining.
4. Run `tsc --skipLibCheck --noEmit **/*.ts` in the directory where `tsconfig.json` exists to confirm the build is clean.
5. Run `bun test` for the affected package to confirm no regressions.

Do not present work as finished while diagnostics remain.

## Linter output handling

* When running `oxfmt` or `oxlint`, focus on the final lines of the output beginning with `{number} issues:` to understand the scope of work.
* Save the full results to a temporary file for easy reference during the resolution process.
* None of the changes should break the code. The project must still compile after every change.

## Resolution workflow

### Oxfmt

When resolving lint issues across the project (not just a single file), follow this process:

1. `cd` into the directory where `.oxfmtrc.jsonc` exists.
2. Run `oxfmt --write --no-error-on-unmatched-pattern .` (no path filter) to produce the full report.
3. Parse the output to count how many errors each individual linter produced. Group errors by linter name (the identifier in parentheses at the end of each line).
4. Sort the linters by error count ascending — fewest errors first.
5. Resolve linters in that order, starting with the linter that has the fewest errors and working up to the linter with the most. This maximizes early progress and keeps changesets small.
6. After clearing all errors for a given linter, re-run `oxfmt --write --no-error-on-unmatched-pattern .` to confirm the count dropped and no new issues were introduced.
7. Continue until the full run reports zero issues.

### Oxlint

When resolving lint issues across the project (not just a single file), follow this process:

1. `cd` into the directory where `.oxlintrc.jsonc` exists.
2. Run `oxlint --import-plugin --jsdoc-plugin --jsx-a11y-plugin --react-perf-plugin --fix --no-error-on-unmatched-pattern .` (no path filter) to produce the full report.
3. Parse the output to count how many errors each individual linter produced. Group errors by linter name (the identifier in parentheses at the end of each line).
4. Sort the linters by error count ascending — fewest errors first.
5. Resolve linters in that order, starting with the linter that has the fewest errors and working up to the linter with the most. This maximizes early progress and keeps changesets small.
6. After clearing all errors for a given linter, re-run `oxlint --import-plugin --jsdoc-plugin --jsx-a11y-plugin --react-perf-plugin --fix --no-error-on-unmatched-pattern .` to confirm the count dropped and no new issues were introduced.
7. Continue until the full run reports zero issues.

### tsc

When resolving lint issues across the project (not just a single file), follow this process:

1. `cd` into the directory where `tsconfig.json` exists.
2. Run `tsc --skipLibCheck --noEmit **/*.ts` (no path filter) to produce the full report.
3. Parse the output to count how many errors each individual linter produced. Group errors by linter name (the identifier in parentheses at the end of each line).
4. Sort the linters by error count ascending — fewest errors first.
5. Resolve linters in that order, starting with the linter that has the fewest errors and working up to the linter with the most. This maximizes early progress and keeps changesets small.
6. After clearing all errors for a given linter, re-run `tsc --skipLibCheck --noEmit **/*.ts` to confirm the count dropped and no new issues were introduced.
7. Continue until the full run reports zero issues.

### Comment line length

Comments, including any whitespace (where tabs count as 2 spaces), must not have individual lines longer than 80 characters. Wrap to the next line instead of continuing on the same line.

Wrong:

```javascript
// validateManagedMarkers checks all profiles at once, used by the validate
// command to give a comprehensive report rather than stopping at the first error.
```

Right:

```javascript
// validateManagedMarkers checks all profiles at once, used by the validate
// command to give a comprehensive report rather than stopping at the first
// error.
```

### Cognitive complexity (gocognit)

Functions with complexity > 20 trigger a warning. Reduce complexity by extracting helper functions for distinct logical branches (e.g., "handle existing destination", "initialize from stub"). Each helper should have a single responsibility and a clear doc comment.

Look for opportunities to extract sections of code into separate functions in order to reduce the overall complexity of the solution.

### Magic numbers

Convert raw numbers into appropriately-named constants, then use those constants in place of the raw numbers. If this cannot be done cleanly, report the issue in the final summary instead of suppressing it.

### Commented-out code

Remove commented-out code blocks. If the code is needed for reference, move it behind a build tag or delete it and rely on version control.

### Testing

Use `bun test` to run tests.

```ts
// index.test.ts
import { test, expect } from "bun:test";

test("hello world", () => {
  expect(1).toBe(1);
});
```

## Code conventions

These rules apply to all JavaScript and TypeScript code in the project, independent of linter enforcement.

### Strict equality

Strongly prefer `===` and `!==` over `==` and `!=`. Loose equality introduces subtle type coercion bugs. Only use loose equality when you have an explicit, documented reason (e.g., intentional `null`/`undefined` coalescing with `== null`).

### Unicode-aware regular expressions

When writing a regular expression, enable Unicode-aware behavior and stricter pattern parsing. Use the `v` flag (ES2024 unicodeSets) when you need set operations or properties of strings; otherwise use the `u` flag. This catches common mistakes like unescaped special characters and ensures correct handling of astral code points.

```typescript
// Wrong
const re = /[a-z]+/;

// Right — basic Unicode awareness
const re = /[a-z]+/u;

// Right — when set operations or string properties are needed
const re = /[\p{Letter}&&\p{Script=Latin}]+/v;
```

### Magic numbers

Avoid "magic numbers." Set non-zero/non-one numeric literals to a well-named constant, then use that constant in place of the raw number. Zero (`0`) and one (`1`) are generally acceptable inline when their meaning is obvious (array indices, increments, identity values).

### Unused parameters

Prefix unused parameters with an underscore (`_`). This signals intent and satisfies linter rules without removing the parameter from the signature (which may be required by an interface or callback contract).

```typescript
// Wrong
function handleEvent(event, context) {
  return context.status;
}

// Right
function handleEvent(_event, context) {
  return context.status;
}
```

### Shadowed declarations

Avoid shadowing variable, parameter, or function names from an outer scope. Shadowing makes code harder to reason about and is a common source of bugs. Choose distinct names or restructure to eliminate the shadow.

### Frontend

Use HTML imports with `Bun.serve()`. Don't use `vite`. HTML imports fully support React, CSS, Tailwind.

Server:

```ts
// index.ts
import index from "./index.html"

Bun.serve({
  routes: {
    "/": index,
    "/api/users/:id": {
      GET: (req) => {
        return new Response(JSON.stringify({ id: req.params.id }));
      },
    },
  },
  // optional websocket support
  websocket: {
    open: (ws) => {
      ws.send("Hello, world!");
    },
    message: (ws, message) => {
      ws.send(message);
    },
    close: (ws) => {
      // handle close
    }
  },
  development: {
    hmr: true,
    console: true,
  }
})
```

HTML files can import .tsx, .jsx or .js files directly and Bun's bundler will transpile & bundle automatically. `<link>` tags can point to stylesheets and Bun's CSS bundler will bundle.

```html
<!-- index.html -->
<html>
  <body>
    <h1>Hello, world!</h1>
    <script type="module" src="./frontend.tsx"></script>
  </body>
</html>
```

With the following `frontend.tsx`:

```tsx
// frontend.tsx
import React from "react";
import { createRoot } from "react-dom/client";

// import .css files directly and it works
import './index.css';

const root = createRoot(document.body);

export default function Frontend() {
  return <h1>Hello, world!</h1>;
}

root.render(<Frontend />);
```

Then, run index.ts

```sh
bun --hot ./index.ts
```

For more information, read the Bun API docs in `node_modules/bun-types/docs/**.mdx`.

### Comment blocks

When making any code change, review and update all comment blocks relevant to the changed code. Doc comments on functions, types, constants, and variables must remain accurate. Do not leave stale comments that describe old behavior.

### Function parameters

If a function requires more than 3 input parameters (excluding `this`), group them into an object and pass it as a single argument. `this` must never be included in a struct — always pass it as a direct argument.

```javascript
// Wrong — too many parameters
function sync(this, source, dest, profile, dryRun) error

// Right
let syncOptions = {
  'source':  'string',
  'dest':    'string',
  'profile': 'string',
  'dryRun':  true,
}

function sync(this, syncOptions)
```

### Function return values

If a function returns more than 3 values (excluding `error`), group them into an object. `this` must never be included in an object.

## Suppression

When a diagnostic cannot be resolved cleanly, use the project's lint suppression comments. Always include a justification. Suppression is a last resort — STRONGLY prefer fixing the root cause.

**Critical rules:**

* Do NOT fall back to suppression comments except as a last resort. The goal is to resolve the issues, not hide them. Even if there are a large number of call sites, fixing them is the goal.
* When function signatures are split across multiple lines, and there is just cause to suppress an error, the suppression comment must be on the same line that triggers the diagnostic error.
* `this` should never be passed as part of a struct. It must be passed as a direct argument.
* For anything deferred (not fixed), present and explain it to the user at the end of the job so the user can follow-up.

<!-- @config-manager:end zdp -->
