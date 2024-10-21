# `@vue/tsconfig`

TSConfigs for Vue projects to extend.

Requires TypeScript >= 4.5.

Install:

```sh
npm add -D @vue/tsconfig
```

Add one of the available configurations to your `tsconfig.json`:

The base configuration (runtime-agnostic):

```json
"extends": "@vue/tsconfig/tsconfig.json"
```

Configuration for Browser environment:

```json
"extends": "@vue/tsconfig/tsconfig.web.json"
```

Configuration for Node environment:

```json
"extends": "@vue/tsconfig/tsconfig.node.json"
```
