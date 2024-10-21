"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _utils = require("@jimp/utils");

/**
 * Creates a circle out of an image.
 * @param {object} options (optional) r: radius of effect
 * @param {function(Error, Jimp)} cb (optional) a callback for when complete
 * @returns {Jimp} this for chaining of methods
 */
var _default = function _default() {
  return {
    fisheye: function fisheye() {
      var _this = this;

      var options = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {
        r: 2.5
      };
      var cb = arguments.length > 1 ? arguments[1] : undefined;

      if (typeof options === 'function') {
        cb = options;
        options = {
          r: 2.5
        };
      }

      var source = this.cloneQuiet();
      var _source$bitmap = source.bitmap,
          width = _source$bitmap.width,
          height = _source$bitmap.height;
      source.scanQuiet(0, 0, width, height, function (x, y) {
        var hx = x / width;
        var hy = y / height;
        var r = Math.sqrt(Math.pow(hx - 0.5, 2) + Math.pow(hy - 0.5, 2));
        var rn = 2 * Math.pow(r, options.r);
        var cosA = (hx - 0.5) / r;
        var sinA = (hy - 0.5) / r;
        var newX = Math.round((rn * cosA + 0.5) * width);
        var newY = Math.round((rn * sinA + 0.5) * height);
        var color = source.getPixelColor(newX, newY);

        _this.setPixelColor(color, x, y);
      });
      /* Set center pixel color, otherwise it will be transparent */

      this.setPixelColor(source.getPixelColor(width / 2, height / 2), width / 2, height / 2);

      if ((0, _utils.isNodePattern)(cb)) {
        cb.call(this, null, this);
      }

      return this;
    }
  };
};

exports["default"] = _default;
module.exports = exports.default;
//# sourceMappingURL=index.js.map