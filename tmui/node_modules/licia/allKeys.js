var keys = require('./keys');
var getProto = require('./getProto');
var unique = require('./unique');
var getOwnPropertyNames = Object.getOwnPropertyNames;
var getOwnPropertySymbols = Object.getOwnPropertySymbols;
exports = function(obj) {
    var _ref =
            arguments.length > 1 && arguments[1] !== undefined
                ? arguments[1]
                : {},
        _ref$prototype = _ref.prototype,
        prototype = _ref$prototype === void 0 ? true : _ref$prototype,
        _ref$unenumerable = _ref.unenumerable,
        unenumerable = _ref$unenumerable === void 0 ? false : _ref$unenumerable,
        _ref$symbol = _ref.symbol,
        symbol = _ref$symbol === void 0 ? false : _ref$symbol;
    var ret = [];
    if ((unenumerable || symbol) && getOwnPropertyNames) {
        var getKeys = keys;
        if (unenumerable && getOwnPropertyNames) getKeys = getOwnPropertyNames;
        do {
            ret = ret.concat(getKeys(obj));
            if (symbol && getOwnPropertySymbols) {
                ret = ret.concat(getOwnPropertySymbols(obj));
            }
        } while (
            prototype &&
            (obj = getProto(obj)) &&
            obj !== Object.prototype
        );
        ret = unique(ret);
    } else {
        if (prototype) {
            for (var key in obj) ret.push(key);
        } else {
            ret = keys(obj);
        }
    }
    return ret;
};

module.exports = exports;
