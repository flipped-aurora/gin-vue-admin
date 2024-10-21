import { isNodePattern } from '@jimp/utils';

/**
 * Creates a circle out of an image.
 * @param {function(Error, Jimp)} options (optional)
 * opacity - opacity of the shadow between 0 and 1
 * size,- of the shadow
 * blur - how blurry the shadow is
 * x- x position of shadow
 * y - y position of shadow
 * @param {function(Error, Jimp)} cb (optional) a callback for when complete
 * @returns {Jimp} this for chaining of methods
 */
export default () => ({
  shadow(options = {}, cb) {
    if (typeof options === 'function') {
      cb = options;
      options = {};
    }

    const { opacity = 0.7, size = 1.1, x = -25, y = 25, blur = 5 } = options;

    // clone the image
    const orig = this.clone();
    const shadow = this.clone();

    // turn all it's pixels black
    shadow.scan(
      0,
      0,
      shadow.bitmap.width,
      shadow.bitmap.height,
      (x, y, idx) => {
        shadow.bitmap.data[idx] = 0x00;
        shadow.bitmap.data[idx + 1] = 0x00;
        shadow.bitmap.data[idx + 2] = 0x00;
        // up the opacity a little,
        shadow.bitmap.data[idx + 3] = shadow.constructor.limit255(
          shadow.bitmap.data[idx + 3] * opacity
        );

        this.bitmap.data[idx] = 0x00;
        this.bitmap.data[idx + 1] = 0x00;
        this.bitmap.data[idx + 2] = 0x00;
        this.bitmap.data[idx + 3] = 0x00;
      }
    );

    // enlarge it. This creates a "shadow".
    shadow
      .resize(shadow.bitmap.width * size, shadow.bitmap.height * size)
      .blur(blur);

    // Then blit the "shadow" onto the background and the image on top of that.
    this.composite(shadow, x, y);
    this.composite(orig, 0, 0);

    if (isNodePattern(cb)) {
      cb.call(this, null, this);
    }

    return this;
  }
});
