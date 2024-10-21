"use strict";

var _interopRequireWildcard = require("@babel/runtime/helpers/interopRequireWildcard");

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

require("core-js/modules/es.symbol");

require("core-js/modules/es.array.filter");

require("core-js/modules/es.array.for-each");

require("core-js/modules/es.array.is-array");

require("core-js/modules/es.object.define-properties");

require("core-js/modules/es.object.define-property");

require("core-js/modules/es.object.entries");

require("core-js/modules/es.object.get-own-property-descriptor");

require("core-js/modules/es.object.get-own-property-descriptors");

require("core-js/modules/es.object.keys");

require("core-js/modules/web.dom-collections.for-each");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = configure;

var _toConsumableArray2 = _interopRequireDefault(require("@babel/runtime/helpers/toConsumableArray"));

var _defineProperty2 = _interopRequireDefault(require("@babel/runtime/helpers/defineProperty"));

var _slicedToArray2 = _interopRequireDefault(require("@babel/runtime/helpers/slicedToArray"));

var _core = _interopRequireWildcard(require("@jimp/core"));

function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(source, true).forEach(function (key) { (0, _defineProperty2["default"])(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(source).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }

function configure(configuration) {
  var jimpInstance = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : _core["default"];
  var jimpConfig = {
    hasAlpha: {},
    encoders: {},
    decoders: {},
    "class": {},
    constants: {}
  };

  function addToConfig(newConfig) {
    Object.entries(newConfig).forEach(function (_ref) {
      var _ref2 = (0, _slicedToArray2["default"])(_ref, 2),
          key = _ref2[0],
          value = _ref2[1];

      jimpConfig[key] = _objectSpread({}, jimpConfig[key], {}, value);
    });
  }

  function addImageType(typeModule) {
    var type = typeModule();

    if (Array.isArray(type.mime)) {
      _core.addType.apply(void 0, (0, _toConsumableArray2["default"])(type.mime));
    } else {
      Object.entries(type.mime).forEach(function (mimeType) {
        return _core.addType.apply(void 0, (0, _toConsumableArray2["default"])(mimeType));
      });
    }

    delete type.mime;
    addToConfig(type);
  }

  function addPlugin(pluginModule) {
    var plugin = pluginModule(_core.jimpEvChange) || {};

    if (!plugin["class"] && !plugin.constants) {
      // Default to class function
      addToConfig({
        "class": plugin
      });
    } else {
      addToConfig(plugin);
    }
  }

  if (configuration.types) {
    configuration.types.forEach(addImageType);
    jimpInstance.decoders = _objectSpread({}, jimpInstance.decoders, {}, jimpConfig.decoders);
    jimpInstance.encoders = _objectSpread({}, jimpInstance.encoders, {}, jimpConfig.encoders);
    jimpInstance.hasAlpha = _objectSpread({}, jimpInstance.hasAlpha, {}, jimpConfig.hasAlpha);
  }

  if (configuration.plugins) {
    configuration.plugins.forEach(addPlugin);
  }

  (0, _core.addJimpMethods)(jimpConfig["class"], jimpInstance);
  (0, _core.addConstants)(jimpConfig.constants, jimpInstance);
  return _core["default"];
}

module.exports = exports.default;
//# sourceMappingURL=index.js.map