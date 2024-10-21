<div align="center">
  <img width="200" height="200"
    src="https://s3.amazonaws.com/pix.iemoji.com/images/emoji/apple/ios-11/256/crayon.png">
  <h1>@jimp/plugin-circle</h1>
  <p>Creates a circle out of an image.</p>
</div>

## Usage

- @param {function(Error, Jimp)} options (optional) radius, x, y
- @param {function(Error, Jimp)} cb (optional) a callback for when complete

```js
import jimp from 'jimp';

async function main() {
  const image = await jimp.read('test/image.png');

  image.circle();
  // or
  image.circle({ radius: 50, x: 25, y: 25 });
}

main();
```
