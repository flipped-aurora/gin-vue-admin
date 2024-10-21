# parse-css-font

[![NPM version](http://img.shields.io/npm/v/parse-css-font.svg?style=flat)](https://www.npmjs.org/package/parse-css-font)
[![npm license](http://img.shields.io/npm/l/parse-css-font.svg?style=flat-square)](https://www.npmjs.org/package/parse-css-font)
[![Travis Build Status](https://img.shields.io/travis/jedmao/parse-css-font.svg)](https://travis-ci.org/jedmao/parse-css-font)
[![codecov](https://codecov.io/gh/jedmao/parse-css-font/branch/master/graph/badge.svg)](https://codecov.io/gh/jedmao/parse-css-font)
[![BundlePhobia Minified](https://badgen.net/bundlephobia/min/parse-css-font?label=min)](https://bundlephobia.com/result?p=parse-css-font)
[![BundlePhobia Minified + gzip](https://badgen.net/bundlephobia/minzip/parse-css-font?label=min%2Bgzip)](https://bundlephobia.com/result?p=parse-css-font)
[![code style: prettier](https://img.shields.io/badge/code_style-prettier-ff69b4.svg)](https://github.com/prettier/prettier)
[![Unicorn Approved](https://img.shields.io/badge/unicorn-approved-ff69b4.svg)](https://twitter.com/sindresorhus/status/457989012528316416?ref_src=twsrc%5Etfw&ref_url=https%3A%2F%2Fwww.quora.com%2FWhat-does-the-unicorn-approved-shield-mean-in-GitHub)

[![npm](https://nodei.co/npm/parse-css-font.svg?downloads=true)](https://nodei.co/npm/parse-css-font/)

Parses the CSS [font property](https://developer.mozilla.org/en-US/docs/Web/CSS/font#font-variant-css21).

## Installation

```
$ npm install parse-css-font [--save[-dev]]
```

## Usage

```js
const parseCSSFont = require('parse-css-font')
parseCSSFont('1rem "Roboto Condensed", sans-serif;')
/*
{
	size: '1rem',
	family: ['Roboto Condensed', 'sans-serif'],
	style: 'normal',
	variant: 'normal',
	weight: 'normal',
	stretch: 'normal',
	lineHeight: 'normal'
}
*/
```

See [the tests](https://github.com/jedmao/parse-css-font/blob/master/src/index.test.ts) for more scenarios.

### ES6/2015 import

```ts
import parseCSSFont from 'parse-css-font'
```

## Testing

```
$ npm test
```

This will run tests and generate a code coverage report. Anything less than 100% coverage will throw an error.
