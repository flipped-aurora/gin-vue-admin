# unimport

[![npm version][npm-version-src]][npm-version-href]
[![npm downloads][npm-downloads-src]][npm-downloads-href]
[![Codecov][codecov-src]][codecov-href]

> Unified utils for auto importing APIs in modules, used in [nuxt](https://github.com/nuxt/nuxt) and [unplugin-auto-import](https://github.com/antfu/unplugin-auto-import)

## Features

- Auto import register APIs for Vite, Webpack or esbuild powered by [unplugin](https://github.com/unjs/unplugin)
- TypeScript declaration file generation
- Auto import for custom APIs defined under specific directories
- Auto import for Vue template

## Install

```sh
# npm
npm install unimport

# yarn
yarn add unimport

# pnpm
pnpm install unimport
```

## Usage

### Plugin Usage

Powered by [unplugin](https://github.com/unjs/unplugin), `unimport` provides a plugin interface for bundlers.

#### Vite / Rollup

```ts
// vite.config.js / rollup.config.js
import Unimport from 'unimport/unplugin'

export default {
  plugins: [
    Unimport.vite({ /* plugin options */ })
  ]
}
```

#### Webpack

```ts
// webpack.config.js
import Unimport from 'unimport/unplugin'

module.exports = {
  plugins: [
    Unimport.webpack({ /* plugin options */ })
  ]
}
```

### Programmatic Usage

<!-- eslint-skip -->

```js
// ESM
import { createUnimport } from 'unimport'

// CommonJS
const { createUnimport } = require('unimport')
```

```js
const { injectImports } = createUnimport({
  imports: [{ name: 'fooBar', from: 'test-id' }]
})

// { code: "import { fooBar } from 'test-id';console.log(fooBar())" }
console.log(injectImports('console.log(fooBar())'))
```

## Configurations

### Imports Item

###### Named import

```ts
imports: [
  { name: 'ref', from: 'vue' },
  { name: 'useState', as: 'useSignal', from: 'react' },
]
```

Will be injected as:

```ts
import { useState as useSignal } from 'react'
import { ref } from 'vue'
```

###### Default import

```ts
imports: [
  { name: 'default', as: '_', from: 'lodash' }
]
```

Will be injected as:

```ts
import _ from 'lodash'
```

###### Namespace import

```ts
imports: [
  { name: '*', as: '_', from: 'lodash' }
]
```

Will be injected as:

```ts
import * as _ from 'lodash'
```

###### Export assignment import

This is a special case for libraries authored with [TypeScript's `export =` syntax](https://www.typescriptlang.org/docs/handbook/modules/reference.html#export--and-import--require). You don't need it the most of the time.

```ts
imports: [
  { name: '=', as: 'browser', from: 'webextension-polyfill' }
]
```

Will be injected as:

```ts
import browser from 'webextension-polyfill'
```

And the type declaration will be added as:

```ts
const browser: typeof import('webextension-polyfill')
```

###### Custom Presets

Presets are provided as a shorthand for declaring imports from the same package:

```ts
presets: [
  {
    from: 'vue',
    imports: [
      'ref',
      'reactive',
      // ...
    ]
  }
]
```

Will be equivalent as:

```ts
imports: [
  { name: 'ref', from: 'vue' },
  { name: 'reactive', from: 'vue' },
  // ...
]
```

###### Built-in Presets

`unimport` also provides some builtin presets for common libraries:

```ts
presets: [
  'vue',
  'pinia',
  'vue-i18n',
  // ...
]
```

You can check out [`src/presets`](./src/presets/) for all the options available or refer to the type declaration.

###### Exports Auto Scan

Since `unimport` v0.7.0, we also support auto scanning the examples from a local installed package, for example:

```ts
presets: [
  {
    package: 'h3',
    ignore: ['isStream', /^[A-Z]/, /^[a-z]*$/, r => r.length > 8]
  }
]
```

This will be expanded into:

```ts
imports: [
  {
    from: 'h3',
    name: 'appendHeader',
  },
  {
    from: 'h3',
    name: 'appendHeaders',
  },
  {
    from: 'h3',
    name: 'appendResponseHeader',
  },
  // ...
]
```

The `ignore` option is used to filter out the exports, it can be a string, regex or a function that returns a boolean.

By default, the result is strongly cached by the version of the package. You can disable this by setting `cache: false`.

### Type Declarations

```ts
Unimport.vite({
  dts: true // or a path to generated file
})
```

### Directory Auto Import

```ts
dirs: [
  './composables/*'
]
```

Named exports for modules under `./composables/*` will be registered for auto imports.

### Opt-out Auto Import

You can opt-out auto-import for specific modules by adding a comment:

```ts
// @unimport-disable
```

It can be customized by setting `commentsDisable`:

```ts
Unimport.vite({
  commentsDisable: [
    '@unimport-disable',
    '@custom-imports-disable',
  ]
})
```

### Acorn Parser

By default, `unimport` uses RegExp to detect unimport entries. In some cases, RegExp might not be able to detect all the entries (false positive & false negative).

We introduced a new AST-based parser powered by [acorn](https://github.com/acornjs/acorn), providing a more accurate result. The limitation is when using Acorn, it assumes all input code are valid and vanilla JavaScript code.

```ts
Unimport.vite({
  parser: 'acorn'
})
```

### Vue Template Auto Import

In Vue's template, the usage of API is in a different context than plain modules. Thus some custom transformations are required. To enable it, set `addons.vueTemplate` to `true`:

```ts
Unimport.vite({
  addons: {
    vueTemplate: true
  }
})
```

#### Caveats

When auto-import a ref, inline operations won't be auto-unwrapped.

```ts
export const counter = ref(0)
```

```html
<template>
  <!-- this is ok -->
  <div>{{ counter }}</div>

  <!-- counter here is a ref, this won't work, volar will throw -->
  <div>{{ counter + 1 }}</div>

  <!-- use this instead -->
  <div>{{ counter.value + 1 }}</div>
</template>
```

We recommend using [Volar](https://github.com/johnsoncodehk/volar) for type checking, which will help you to identify the misusage.

### Vue Directives Auto Import and TypeScript Declaration Generation

In Vue's template, the usage of directives is in a different context than plain modules. Thus some custom transformations are required. To enable it, set `addons.vueDirectives` to `true`:

```ts
Unimport.vite({
  addons: {
    vueDirectives: true
  }
})
```

#### Library Authors

When including directives in your presets, you should:
- provide the corresponding imports with `meta.vueDirective` set to `true`, otherwise, `unimport` will not be able to detect your directives.
- use named exports for your directives, or use default export and use `as` in the Import.
- set `dtsDisabled` to `true` if you provide a type declaration for your directives.

```ts
import type { InlinePreset } from 'unimport'
import { defineUnimportPreset } from 'unimport'

export const composables = defineUnimportPreset({
  from: 'my-unimport-library/composables',
  /* imports and other options */
})

export const directives = defineUnimportPreset({
  from: 'my-unimport-library/directives',
  // disable dts generation globally
  dtsEnabled: false,
  // you can declare the vue directive globally
  meta: {
    vueDirective: true
  },
  imports: [{
    name: 'ClickOutside',
    // disable dts generation per import
    dtsEnabled: false,
    // you can declare the vue directive per import
    meta: {
      vueDirective: true
    }
  }, {
    name: 'default',
    // you should declare `as` for default exports
    as: 'Focus'
  }]
})
```

#### Using Directory Scan and Local Directives

If you add a directory scan for your local directives in the project, you need to:
- provide `isDirective` in the `vueDirectives`: `unimport` will use it to detect them (will never be called for imports with `meta.vueDirective` set to `true`).
- use always named exports for your directives.

```ts
Unimport.vite({
  dirs: ['./directives/**'],
  addons: {
    vueDirectives: {
      isDirective: (normalizedImportFrom, _importEntry) => {
        return normalizedImportFrom.includes('/directives/')
      }
    }
  }
})
```

## 💻 Development

- Clone this repository
- Enable [Corepack](https://github.com/nodejs/corepack) using `corepack enable` (use `npm i -g corepack` for Node.js < 16.10)
- Install dependencies using `pnpm install`
- Run interactive tests using `pnpm dev`

## License

Made with 💛

Published under [MIT License](./LICENSE).

<!-- Badges -->
[npm-version-src]: https://img.shields.io/npm/v/unimport?style=flat-square
[npm-version-href]: https://npmjs.com/package/unimport

[npm-downloads-src]: https://img.shields.io/npm/dm/unimport?style=flat-square
[npm-downloads-href]: https://npmjs.com/package/unimport

[github-actions-src]: https://img.shields.io/github/workflow/status/unjs/unimport/ci/main?style=flat-square
[github-actions-href]: https://github.com/unjs/unimport/actions?query=workflow%3Aci

[codecov-src]: https://img.shields.io/codecov/c/gh/unjs/unimport/main?style=flat-square
[codecov-href]: https://codecov.io/gh/unjs/unimport
