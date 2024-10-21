# safeAreaInsets

使用javascript获取安全区信息.

## 先决条件

* WKWebview
* viewport-fit=cover

## 安装

### NPM

```shell
npm install safe-area-insets --save
```

### UMD

```http
https://unpkg.com/safe-area-insets/out/umd/index.min.js
```

## 使用

### 引用

```js
var safeAreaInsets = require('safe-area-insets')

```

### 安全区信息

```js
console.log('safeAreaInsets.support', safeAreaInsets.support)
console.log('safe-area-inset-top', safeAreaInsets.top)
console.log('safe-area-inset-left', safeAreaInsets.left)
console.log('safe-area-inset-right', safeAreaInsets.right)
console.log('safe-area-inset-bottom', safeAreaInsets.bottom)
```

### 监听安全区信息变化事件

```js
function callback(style){
    console.log(style)
}
// add
safeAreaInsets.onChange(callback)
// remove
safeAreaInsets.offChange(callback)
```

## 相关阅读

>[为iPhoneX设计网站](https://webkit.org/blog/7929/designing-websites-for-iphone-x/)
