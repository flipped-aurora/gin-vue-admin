"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _defineProperty2 = _interopRequireDefault(require("@babel/runtime/helpers/defineProperty"));

var _bmpJs = _interopRequireDefault(require("bmp-js"));

var _utils = require("@jimp/utils");

var MIME_TYPE = 'image/bmp';
var MIME_TYPE_SECOND = 'image/x-ms-bmp';

function toAGBR(image) {
  return (0, _utils.scan)(image, 0, 0, image.bitmap.width, image.bitmap.height, function (x, y, index) {
    var red = this.bitmap.data[index + 0];
    var green = this.bitmap.data[index + 1];
    var blue = this.bitmap.data[index + 2];
    var alpha = this.bitmap.data[index + 3];
    this.bitmap.data[index + 0] = alpha;
    this.bitmap.data[index + 1] = blue;
    this.bitmap.data[index + 2] = green;
    this.bitmap.data[index + 3] = red;
  }).bitmap;
}

function fromAGBR(bitmap) {
  return (0, _utils.scan)({
    bitmap: bitmap
  }, 0, 0, bitmap.width, bitmap.height, function (x, y, index) {
    var alpha = this.bitmap.data[index + 0];
    var blue = this.bitmap.data[index + 1];
    var green = this.bitmap.data[index + 2];
    var red = this.bitmap.data[index + 3];
    this.bitmap.data[index + 0] = red;
    this.bitmap.data[index + 1] = green;
    this.bitmap.data[index + 2] = blue;
    this.bitmap.data[index + 3] = bitmap.is_with_alpha ? alpha : 0xff;
  }).bitmap;
}

var decode = function decode(data) {
  return fromAGBR(_bmpJs["default"].decode(data));
};

var encode = function encode(image) {
  return _bmpJs["default"].encode(toAGBR(image)).data;
};

var _default = function _default() {
  var _decoders, _encoders;

  return {
    mime: (0, _defineProperty2["default"])({}, MIME_TYPE, ['bmp']),
    constants: {
      MIME_BMP: MIME_TYPE,
      MIME_X_MS_BMP: MIME_TYPE_SECOND
    },
    decoders: (_decoders = {}, (0, _defineProperty2["default"])(_decoders, MIME_TYPE, decode), (0, _defineProperty2["default"])(_decoders, MIME_TYPE_SECOND, decode), _decoders),
    encoders: (_encoders = {}, (0, _defineProperty2["default"])(_encoders, MIME_TYPE, encode), (0, _defineProperty2["default"])(_encoders, MIME_TYPE_SECOND, encode), _encoders)
  };
};

exports["default"] = _default;
module.exports = exports.default;
//# sourceMappingURL=index.js.map