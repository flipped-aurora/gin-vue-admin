JavaScript QRCode reader for HTML5 enabled browser.
===================================================

This was started as a port of Lazarsoft’s qrcode reader.

It is a maintained fork, so feel free to open issues or PR !

[![Build Status](https://travis-ci.org/edi9999/jsqrcode.svg?branch=master&style=flat)](https://travis-ci.org/edi9999/jsqrcode)

Installation
============

```
npm install qrcode-reader
```

Usage
=====

```
var QrCode = require('qrcode-reader');
```

Create a new instance of QrCode:

```
var qr = new QrCode();
```

Set its callback to a custom function:

```
qr.callback = function(error, result) {
  if(error) {
    console.log(error)
    return;
  }
  console.log(result)
}
```

Passing image data in node
==========================

You have to use an external imageparser

1.	You can use `npm install --save jimp` which doesn't have any dependency (runs in pure node).

	```javascript
	var Jimp = require("jimp");
	var buffer = fs.readFileSync(__dirname + '/image.png');
	Jimp.read(buffer, function(err, image) {
	    if (err) {
	        console.error(err);
	        // TODO handle error
	    }
	    var qr = new QrCode();
	    qr.callback = function(err, value) {
	        if (err) {
	            console.error(err);
	            // TODO handle error
	        }
	        console.log(value.result);
	        console.log(value);
	    };
	    qr.decode(image.bitmap);
	});
	```

2.	You can use `npm install --save image-parser`, which depends on lwip or graphicsmagick

	```javascript
	var ImageParser = require("image-parser");
	var buffer = fs.readFileSync(__dirname + '/image.png');
	var img = new ImageParser(img);
	img.parse(function(err) {
	    if (err) {
	        console.error(err);
	        // TODO handle error
	    }
	    var qr = new QrCode();
	    qr.callback = function(err, value) {
	        if (err) {
	            console.error(err);
	            // TODO handle error
	            return done(err);
	        }
	        console.log(value.result);
	        console.log(value);
	    };
	    qr.decode({width: img.width(), height: img.height()}, img._imgBuffer);
	});
	```

Passing image data in the browser
=================================

Since the browser contains the Image global, we can use it to open images with URL, Data URI, ...

Decode an image by its URL or Data URI:

```

qr.decode(url or DataURL);

```

Decode an image by context.getImageData: Works with [web workers](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API/Using_web_workers).

```

var context = canvas.getContext("2d"); var data = context.getImageData(0, 0, width, height);

qr.decode(data);

```

====================

If you want, you can build the script yourself.

First clone the repository, then from the directory of this repository, do:

```

npm install

```

To run the build process and generate a JavaScript file called `dist/index.js` you can run from node:

```

npm run build

```

To run the tests:

```

npm test

```

Make it work in the browser
===========================

The generated file `dist/index.js` works in the browser.

You will have access to the global variable `QrCode` if you do the following in your HTML:

\`\`\`

<script src="dist/index.js"></script> \`\`\`

See [examples/browser-upload/index.html](examples/browser-upload/index.html) for a very basic example using a file upload.

Changelog
=========

See [`CHANGELOG.md`](CHANGELOG.md)\.
