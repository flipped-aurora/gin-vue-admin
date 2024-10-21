"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.isNodePattern = isNodePattern;
exports.throwError = throwError;
exports.scan = scan;
exports.scanIterator = scanIterator;

var _regenerator = _interopRequireDefault(require("@babel/runtime/regenerator"));

require("regenerator-runtime/runtime");

var _marked =
/*#__PURE__*/
_regenerator["default"].mark(scanIterator);

function isNodePattern(cb) {
  if (typeof cb === 'undefined') {
    return false;
  }

  if (typeof cb !== 'function') {
    throw new TypeError('Callback must be a function');
  }

  return true;
}

function throwError(error, cb) {
  if (typeof error === 'string') {
    error = new Error(error);
  }

  if (typeof cb === 'function') {
    return cb.call(this, error);
  }

  throw error;
}

function scan(image, x, y, w, h, f) {
  // round input
  x = Math.round(x);
  y = Math.round(y);
  w = Math.round(w);
  h = Math.round(h);

  for (var _y = y; _y < y + h; _y++) {
    for (var _x = x; _x < x + w; _x++) {
      var idx = image.bitmap.width * _y + _x << 2;
      f.call(image, _x, _y, idx);
    }
  }

  return image;
}

function scanIterator(image, x, y, w, h) {
  var _y, _x, idx;

  return _regenerator["default"].wrap(function scanIterator$(_context) {
    while (1) {
      switch (_context.prev = _context.next) {
        case 0:
          // round input
          x = Math.round(x);
          y = Math.round(y);
          w = Math.round(w);
          h = Math.round(h);
          _y = y;

        case 5:
          if (!(_y < y + h)) {
            _context.next = 17;
            break;
          }

          _x = x;

        case 7:
          if (!(_x < x + w)) {
            _context.next = 14;
            break;
          }

          idx = image.bitmap.width * _y + _x << 2;
          _context.next = 11;
          return {
            x: _x,
            y: _y,
            idx: idx,
            image: image
          };

        case 11:
          _x++;
          _context.next = 7;
          break;

        case 14:
          _y++;
          _context.next = 5;
          break;

        case 17:
        case "end":
          return _context.stop();
      }
    }
  }, _marked);
}
//# sourceMappingURL=index.js.map