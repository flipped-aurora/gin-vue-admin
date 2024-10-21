# localstorage-polyfill

in memory localStorage polyfill for node.js utilizing ES6 proxies

## Installation

```
npm i localstorage-polyfill -D
```
Saving for development, because primarily this package is meant for unit testing browser in node.js.

## Usage

```javascript
require('localstorage-polyfill')
// or
import 'localstorage-polyfill'
global.localStorage // now has your in memory localStorage
```

For API doc, refer to MDN.

There are other packages like node-localStorage, but none of them work properly when you do

```javascript
localStorage.c = 1
```

They only shim the methods of localStorage object, they don't imitate it's behaviour fully.
