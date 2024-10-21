# safeAreaInsets

Use javascript to get the safe area insets.

Chinese documentation : [中文文档](https://github.com/zhetengbiji/safeAreaInsets/blob/master/README-CH.md)

## Precondition

* WKWebview
* viewport-fit=cover

## Install

### NPM

```
npm install safe-area-insets --save
```

### UMD

```http
https://unpkg.com/safe-area-insets/out/umd/index.min.js
```

## Use

### Require

```js
const safeAreaInsets = require('safe-area-insets')
```

### Safe Area Info

```js
console.log('safeAreaInsets.support', safeAreaInsets.support)
console.log('safe-area-inset-top', safeAreaInsets.top)
console.log('safe-area-inset-left', safeAreaInsets.left)
console.log('safe-area-inset-right', safeAreaInsets.right)
console.log('safe-area-inset-bottom', safeAreaInsets.bottom)
```

### Listening Change Event

```js
function callback(style){
    console.log(style)
}
// add
safeAreaInsets.onChange(callback)
// remove
safeAreaInsets.offChange(callback)
```

## Related Reading

>[Designing Websites for iPhone X](https://webkit.org/blog/7929/designing-websites-for-iphone-x/)