"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _custom = _interopRequireDefault(require("@jimp/custom"));

var _types = _interopRequireDefault(require("@jimp/types"));

var _plugins = _interopRequireDefault(require("@jimp/plugins"));

var _default = (0, _custom["default"])({
  types: [_types["default"]],
  plugins: [_plugins["default"]]
});

exports["default"] = _default;
//# sourceMappingURL=index.js.map