# unplugin-vue-components

[![NPM version](https://img.shields.io/npm/v/unplugin-vue-components?color=a1b858&label=)](https://www.npmjs.com/package/unplugin-vue-components)

On-demand components auto importing for Vue.

###### Features

- 💚 Supports both Vue 2 and Vue 3 out-of-the-box.
- ✨ Supports both components and directives.
- ⚡️ Supports Vite, Webpack, Rspack, Vue CLI, Rollup, esbuild and more, powered by <a href="https://github.com/unjs/unplugin">unplugin</a>.
- 🏝 Tree-shakable, only registers the components you use.
- 🪐 Folder names as namespaces.
- 🦾 Full TypeScript support.
- 🌈 [Built-in resolvers](#importing-from-ui-libraries) for popular UI libraries.
- 😃 Works perfectly with [unplugin-icons](https://github.com/antfu/unplugin-icons).

<br>

<p align="center">
  <a href="https://cdn.jsdelivr.net/gh/antfu/static/sponsors.svg">
    <img src='https://cdn.jsdelivr.net/gh/antfu/static/sponsors.svg'/>
  </a>
</p>

<br>

## Installation

```bash
npm i unplugin-vue-components -D
```

> **`vite-plugin-components` has been renamed to `unplugin-vue-components`**, see the [migration guide](#migrate-from-vite-plugin-components).

<details>
<summary>Vite</summary><br>

```ts
// vite.config.ts
import Components from 'unplugin-vue-components/vite'

export default defineConfig({
  plugins: [
    Components({ /* options */ }),
  ],
})
```

<br></details>

<details>
<summary>Rollup</summary><br>

```ts
// rollup.config.js
import Components from 'unplugin-vue-components/rollup'

export default {
  plugins: [
    Components({ /* options */ }),
  ],
}
```

<br></details>

<details>
<summary>Webpack</summary><br>

```ts
// webpack.config.js
module.exports = {
  /* ... */
  plugins: [
    require('unplugin-vue-components/webpack').default({ /* options */ }),
  ],
}
```

<br></details>

<details>
<summary>Rspack</summary><br>

```ts
// rspack.config.js
module.exports = {
  /* ... */
  plugins: [
    require('unplugin-vue-components/rspack').default({ /* options */ }),
  ],
}
```

<br></details>

<details>
<summary>Nuxt</summary><br>

You might not need this plugin for Nuxt. Use [`@nuxt/components`](https://github.com/nuxt/components) instead.

<br></details>

<details>
<summary>Vue CLI</summary><br>

```ts
// vue.config.js
module.exports = {
  /* ... */
  plugins: [
    require('unplugin-vue-components/webpack').default({ /* options */ }),
  ],
}
```

You can also rename the Vue configuration file to `vue.config.mjs` and use static import syntax (you should use latest `@vue/cli-service ^5.0.8`):

```ts
// vue.config.mjs
import Components from 'unplugin-vue-components/webpack'

export default {
  configureWebpack: {
    plugins: [
      Components({ /* options */ }),
    ],
  },
}
```

<br></details>

<details>
<summary>esbuild</summary><br>

```ts
// esbuild.config.js
import { build } from 'esbuild'
import Components from 'unplugin-vue-components/esbuild'

build({
  /* ... */
  plugins: [
    Components({
      /* options */
    }),
  ],
})
```

<br></details>

## Usage

Use components in templates as you would usually do, it will import components on demand, and there is no `import` and `component registration` required anymore! If you register the parent component asynchronously (or lazy route), the auto-imported components will be code-split along with their parent.

It will automatically turn this

```html
<template>
  <div>
    <HelloWorld msg="Hello Vue 3.0 + Vite" />
  </div>
</template>

<script>
  export default {
    name: 'App',
  }
</script>
```

into this

```html
<template>
  <div>
    <HelloWorld msg="Hello Vue 3.0 + Vite" />
  </div>
</template>

<script>
  import HelloWorld from './src/components/HelloWorld.vue'

  export default {
    name: 'App',
    components: {
      HelloWorld,
    },
  }
</script>
```

> **Note**
> By default this plugin will import components in the `src/components` path. You can customize it using the `dirs` option.

## TypeScript

To get TypeScript support for auto-imported components, there is [a PR](https://github.com/vuejs/core/pull/3399) to Vue 3 extending the interface of global components. Currently, [Volar](https://github.com/johnsoncodehk/volar) has supported this usage already. If you are using Volar, you can change the config as following to get the support.

```ts
Components({
  dts: true, // enabled by default if `typescript` is installed
})
```

Once the setup is done, a `components.d.ts` will be generated and updates automatically with the type definitions. Feel free to commit it into git or not as you want.

> **Make sure you also add `components.d.ts` to your `tsconfig.json` under `include`.**

## Importing from UI Libraries

We have several built-in resolvers for popular UI libraries like **Vuetify**, **Ant Design Vue**, and **Element Plus**, where you can enable them by:

Supported Resolvers:

- [Ant Design Vue](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/antdv.ts)
- [Arco Design Vue](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/arco.ts)
- [BootstrapVue](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/bootstrap-vue.ts)
- [Element Plus](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/element-plus.ts)
- [Element UI](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/element-ui.ts)
- [Headless UI](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/headless-ui.ts)
- [IDux](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/idux.ts)
- [Inkline](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/inkline.ts)
- [Ionic](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/ionic.ts)
- [Naive UI](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/naive-ui.ts)
- [Prime Vue](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/prime-vue.ts)
- [Quasar](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/quasar.ts)
- [TDesign](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/tdesign.ts)
- [Vant](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/vant.ts)
- [Varlet UI](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/varlet-ui.ts)
- [VEUI](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/veui.ts)
- [View UI](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/view-ui.ts)
- [Vuetify](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/vuetify.ts) &mdash; Prefer first-party plugins when possible: [v3 + vite](https://www.npmjs.com/package/vite-plugin-vuetify), [v3 + webpack](https://www.npmjs.com/package/webpack-plugin-vuetify), [v2 + webpack](https://npmjs.com/package/vuetify-loader)
- [VueUse Components](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/vueuse.ts)
- [VueUse Directives](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/vueuse-directive.ts)
- [Dev UI](https://github.com/antfu/unplugin-vue-components/blob/main/src/core/resolvers/devui.ts)

```ts
// vite.config.js
import Components from 'unplugin-vue-components/vite'
import {
  AntDesignVueResolver,
  ElementPlusResolver,
  VantResolver,
} from 'unplugin-vue-components/resolvers'

// your plugin installation
Components({
  resolvers: [
    AntDesignVueResolver(),
    ElementPlusResolver(),
    VantResolver(),
  ],
})
```

You can also write your own resolver quickly:

```ts
Components({
  resolvers: [
    // example of importing Vant
    (componentName) => {
      // where `componentName` is always CapitalCase
      if (componentName.startsWith('Van'))
        return { name: componentName.slice(3), from: 'vant' }
    },
  ],
})
```

> [We no longer accept new resolvers](./src/core/resolvers/_READ_BEFORE_CONTRIBUTE.md).

## Types for global registered components

Some libraries might register some global components for you to use anywhere (e.g. Vue Router provides `<RouterLink>` and `<RouterView>`). Since they are global available, there is no need for this plugin to import them. However, those are commonly not TypeScript friendly, and you might need to register their types manually.

Thus `unplugin-vue-components` provided a way to only register types for global components.

```ts
Components({
  dts: true,
  types: [{
    from: 'vue-router',
    names: ['RouterLink', 'RouterView'],
  }],
})
```

So the `RouterLink` and `RouterView` will be presented in `components.d.ts`.

By default, `unplugin-vue-components` detects supported libraries automatically (e.g. `vue-router`) when they are installed in the workspace. If you want to disable it completely, you can pass an empty array to it:

```ts
Components({
  // Disable type only registration
  types: [],
})
```

## Migrate from `vite-plugin-components`

`package.json`

```diff
{
  "devDependencies": {
-   "vite-plugin-components": "*",
+   "unplugin-vue-components": "^0.14.0",
  }
}
```

`vite.config.js`

```diff
- import Components, { ElementPlusResolver } from 'vite-plugin-components'
+ import Components from 'unplugin-vue-components/vite'
+ import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

export default {
  plugins: [
    /* ... */
    Components({
      /* ... */

      // `customComponentsResolvers` has renamed to `resolver`
-     customComponentsResolvers: [
+     resolvers: [
        ElementPlusResolver(),
      ],

      // `globalComponentsDeclaration` has renamed to `dts`
-     globalComponentsDeclaration: true,
+     dts: true,

      // `customLoaderMatcher` is depreacted, use `include` instead
-     customLoaderMatcher: id => id.endsWith('.md'),
+     include: [/\.vue$/, /\.vue\?vue/, /\.md$/],
    }),
  ],
}
```

## Configuration

The following show the default values of the configuration

```ts
Components({
  // relative paths to the directory to search for components.
  dirs: ['src/components'],

  // valid file extensions for components.
  extensions: ['vue'],

  // Glob patterns to match file names to be detected as components.
  // When specified, the `dirs`, `extensions`, and `directoryAsNamespace` options will be ignored.
  // If you want to exclude components being registered, use negative globs with leading `!`.
  globs: ['src/components/*.{vue}'],

  // search for subdirectories
  deep: true,

  // resolvers for custom components
  resolvers: [],

  // generate `components.d.ts` global declarations,
  // also accepts a path for custom filename
  // default: `true` if package typescript is installed
  dts: false,

  // Allow subdirectories as namespace prefix for components.
  directoryAsNamespace: false,

  // Collapse same prefixes (camel-sensitive) of folders and components
  // to prevent duplication inside namespaced component name.
  // works when `directoryAsNamespace: true`
  collapseSamePrefixes: false,

  // Subdirectory paths for ignoring namespace prefixes.
  // works when `directoryAsNamespace: true`
  globalNamespaces: [],

  // auto import for directives
  // default: `true` for Vue 3, `false` for Vue 2
  // Babel is needed to do the transformation for Vue 2, it's disabled by default for performance concerns.
  // To install Babel, run: `npm install -D @babel/parser`
  directives: true,

  // Transform path before resolving
  importPathTransform: v => v,

  // Allow for components to override other components with the same name
  allowOverrides: false,

  // Filters for transforming targets (components to insert the auto import)
  // Note these are NOT about including/excluding components registered - use `globs` or `excludeNames` for that
  include: [/\.vue$/, /\.vue\?vue/],
  exclude: [/[\\/]node_modules[\\/]/, /[\\/]\.git[\\/]/, /[\\/]\.nuxt[\\/]/],

  // Filters for component names that will not be imported
  // Use for globally imported async components or other conflicts that the plugin cannot detect
  excludeNames: [/^Async.+/],

  // Vue version of project. It will detect automatically if not specified.
  // Acceptable value: 2 | 2.7 | 3
  version: 2.7,

  // Only provide types of components in library (registered globally)
  types: []
})
```

## Example

[Vitesse](https://github.com/antfu/vitesse) starter template.

## Thanks

Thanks to [@brattonross](https://github.com/brattonross), this project is heavily inspired by [vite-plugin-voie](https://github.com/vamplate/vite-plugin-voie).

## License

[MIT](./LICENSE) License © 2020-PRESENT [Anthony Fu](https://github.com/antfu)
