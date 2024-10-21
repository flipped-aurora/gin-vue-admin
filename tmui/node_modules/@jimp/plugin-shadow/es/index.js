"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _utils = require("@jimp/utils");

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
var _default = function _default() {
  return {
    shadow: function shadow() {
      var _this = this;

      var options = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};
      var cb = arguments.length > 1 ? arguments[1] : undefined;

      if (typeof options === 'function') {
        cb = options;
        options = {};
      }

      var _options = options,
          _options$opacity = _options.opacity,
          opacity = _options$opacity === void 0 ? 0.7 : _options$opacity,
          _options$size = _options.size,
          size = _options$size === void 0 ? 1.1 : _options$size,
          _options$x = _options.x,
          x = _options$x === void 0 ? -25 : _options$x,
          _options$y = _options.y,
          y = _options$y === void 0 ? 25 : _options$y,
          _options$blur = _options.blur,
          blur = _options$blur === void 0 ? 5 : _options$blur; // clone the image

      var orig = this.clone();
      var shadow = this.clone(); // turn all it's pixels black

      shadow.scan(0, 0, shadow.bitmap.width, shadow.bitmap.height, function (x, y, idx) {
        shadow.bitmap.data[idx] = 0x00;
        shadow.bitmap.data[idx + 1] = 0x00;
        shadow.bitmap.data[idx + 2] = 0x00; // up the opacity a little,

        shadow.bitmap.data[idx + 3] = shadow.constructor.limit255(shadow.bitmap.data[idx + 3] * opacity);
        _this.bitmap.data[idx] = 0x00;
        _this.bitmap.data[idx + 1] = 0x00;
        _this.bitmap.data[idx + 2] = 0x00;
        _this.bitmap.data[idx + 3] = 0x00;
      }); // enlarge it. This creates a "shadow".

      shadow.resize(shadow.bitmap.width * size, shadow.bitmap.height * size).blur(blur); // Then blit the "shadow" onto the background and the image on top of that.

      this.composite(shadow, x, y);
      this.composite(orig, 0, 0);

      if ((0, _utils.isNodePattern)(cb)) {
        cb.call(this, null, this);
      }

      return this;
    }
  };
};

exports["default"] = _default;
//# sourceMappingURL=index.js.map