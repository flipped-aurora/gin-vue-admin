# OS locale detector (os-locale-s)
[![Language grade: JavaScript](https://img.shields.io/lgtm/grade/javascript/g/jeffy-g/os-locale-s.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/jeffy-g/os-locale-s/context:javascript) ![Node.js CI](https://github.com/jeffy-g/os-locale-s/workflows/Node.js%20CI/badge.svg) ![GitHub](https://img.shields.io/github/license/jeffy-g/os-locale-s?style=flat) 

__Its a light weight version that minimizes the dependency module of `os-locale`__


> Get the system [locale](https://en.wikipedia.org/wiki/Locale_(computer_software))

Useful for localizing your module or app.

POSIX systems: The returned locale refers to the [`LC_MESSAGE`](http://www.gnu.org/software/libc/manual/html_node/Locale-Categories.html#Locale-Categories) category, suitable for selecting the language used in the user interface for message translation.

## Motivacion

I was creating a module package and needed to detect the **locale**.

At that time, I found a module `os-locale` that provides the corresponding function in the npm package.

I just wanted to detect locale token such as "cs", "de", "es", "fr" etc but  
`os-locale` installs as many as 27 npm packages, including indirectly dependent modules
  + e.g - `yarn add os-locale`

I just wanted to detect a simple locale token but thought this was overwork and decided to rewrite the code.

At the same time, has been migrate to `TypeScript` as well.

> As a result, the number of npm packages installed in `os-locale-s` has been reduced to 3 including `os-locale-s`


## Install

```
$ npm install os-locale-s
```

## Usage

```js
// node (commenjs)
const { osLocale } = require("os-locale-s");
(async () => {
    console.log(await osLocale());
    //=> 'en-US'
})();
```

```ts
// ECMA module
import { osLocale } from "os-locale-s";
(async () => {
    console.log(await osLocale());
    //=> 'en-US'
})();
```
## API

### osLocale(options?)

Returns a `Promise` for the locale.

### osLocale.sync(options?)

Returns the locale.

#### options

Type: `object`

##### spawn

Type: `boolean`\
Default: `true`

Set to `false` to avoid spawning subprocesses and instead only resolve the locale from environment variables. (`process.env`)

##### cache

Type: `boolean`\
Default: `true`

Once the locale is detected, its value is retained and reused at the second and subsequent detections.

If set to `false`, the last held value will be ignored and do locale detection again (and the resulting value is not preserved)
