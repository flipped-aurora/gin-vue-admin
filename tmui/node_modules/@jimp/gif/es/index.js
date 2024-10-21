"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _defineProperty2 = _interopRequireDefault(require("@babel/runtime/helpers/defineProperty"));

var _omggif = _interopRequireDefault(require("omggif"));

var MIME_TYPE = 'image/gif';

var _default = function _default() {
  return {
    mime: (0, _defineProperty2["default"])({}, MIME_TYPE, ['gif']),
    constants: {
      MIME_GIF: MIME_TYPE
    },
    decoders: (0, _defineProperty2["default"])({}, MIME_TYPE, function (data) {
      var gifObj = new _omggif["default"].GifReader(data);
      var gifData = Buffer.alloc(gifObj.width * gifObj.height * 4);
      gifObj.decodeAndBlitFrameRGBA(0, gifData);
      return {
        data: gifData,
        width: gifObj.width,
        height: gifObj.height
      };
    })
  };
};

exports["default"] = _default;
//# sourceMappingURL=index.js.map