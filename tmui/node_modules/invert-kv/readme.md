# invert-kv [![Build Status](https://travis-ci.com/sindresorhus/invert-kv.svg?branch=master)](https://travis-ci.com/sindresorhus/invert-kv)

> Invert the key/value of an object. Example: `{foo: 'bar'}` → `{bar: 'foo'}`

## Install

```
$ npm install invert-kv
```

## Usage

```js
const invertKeyValue = require('invert-kv');

invertKeyValue({foo: 'bar', '🦄': '🌈'});
//=> {bar: 'foo', '🌈': '🦄'}
```

---

<div align="center">
	<b>
		<a href="https://tidelift.com/subscription/pkg/npm-invert-kv?utm_source=npm-invert-kv&utm_medium=referral&utm_campaign=readme">Get professional support for this package with a Tidelift subscription</a>
	</b>
	<br>
	<sub>
		Tidelift helps make open source sustainable for maintainers while giving companies<br>assurances about security, maintenance, and licensing for their dependencies.
	</sub>
</div>
