<div align="center">
  <img width="200" height="200"
    src="https://s3.amazonaws.com/pix.iemoji.com/images/emoji/apple/ios-11/256/crayon.png">
  <h1>@jimp/plugin-shadow</h1>
  <p>Creates a shadow on an image.</p>
</div>

## Usage

- @param {function(Error, Jimp)} options (optional)
  - opacity - opacity of the shadow between 0 and 1
  - size,- of the shadow
  - blur - how blurry the shadow is
  - x - x position of shadow
  - y - y position of shadow
- @param {function(Error, Jimp)} cb (optional) a callback for when complete

```js
import jimp from 'jimp';

async function main() {
  const image = await jimp.read('test/image.png');

  image.shadow();
  // or
  image.shadow({ opacity: 0.8, size: 1.2, blur: 10, x: -75, y: -75 });
}

main();
```
