<div align="center">
  <img width="200" height="200"
    src="https://s3.amazonaws.com/pix.iemoji.com/images/emoji/apple/ios-11/256/crayon.png">
  <h1>@jimp/plugin-threshold</h1>
  <p>Lightens an image.</p>
</div>

This is useful as a simplified method for processing scanned drawings, signatures, etc

## Usage

- @param {number} options object
  - max: A number auto limited between 0 - 255
  - replace: (optional) A number auto limited between 0 - 255 (default 255)
  - autoGreyscale: (optional) A boolean whether to apply greyscale beforehand (default true)
- @param {function(Error, Jimp)} cb (optional) a callback for when complete

```js
import jimp from 'jimp';

async function main() {
  const image = await jimp.read('test/image.png');

  image.threshold({ max: 150 });
  // or
  image.threshold({ max: 200, replace: 200, autoGreyscale: false });
}

main();
```
