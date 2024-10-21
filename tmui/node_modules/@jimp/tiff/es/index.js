"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _defineProperty2 = _interopRequireDefault(require("@babel/runtime/helpers/defineProperty"));

var _utif = _interopRequireDefault(require("utif"));

var MIME_TYPE = 'image/tiff';

var _default = function _default() {
  return {
    mime: (0, _defineProperty2["default"])({}, MIME_TYPE, ['tiff', 'tif']),
    constants: {
      MIME_TIFF: MIME_TYPE
    },
    decoders: (0, _defineProperty2["default"])({}, MIME_TYPE, function (data) {
      var ifds = _utif["default"].decode(data);

      var page = ifds[0];

      _utif["default"].decodeImages(data, ifds);

      var rgba = _utif["default"].toRGBA8(page);

      return {
        data: Buffer.from(rgba),
        width: page.t256[0],
        height: page.t257[0]
      };
    }),
    encoders: (0, _defineProperty2["default"])({}, MIME_TYPE, function (image) {
      var tiff = _utif["default"].encodeImage(image.bitmap.data, image.bitmap.width, image.bitmap.height);

      return Buffer.from(tiff);
    })
  };
};

exports["default"] = _default;
//# sourceMappingURL=index.js.map