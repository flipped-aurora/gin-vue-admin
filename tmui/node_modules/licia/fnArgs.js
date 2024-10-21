var startWith = require('./startWith');
var last = require('./last');
var lowerCase = require('./lowerCase');
var isObj = require('./isObj');
var type = require('./type');
exports = function(types, args) {
    var argsLen = args.length;
    var typesLen = types.length;
    var minLen = typesLen;
    var maxLen = typesLen;
    for (var i = 0; i < typesLen; i++) {
        var _type = types[i].split('|');
        if (startWith(_type[0], '?')) {
            _type[0] = _type[0].slice(1);
            if (minLen === typesLen) {
                minLen = i;
            }
        }
        if (i === typesLen - 1 && startWith(_type[0], '...')) {
            maxLen = Infinity;
            _type[0] = _type[0].slice(3);
            if (minLen === typesLen) {
                minLen = i;
            }
        }
        types[i] = _type;
    }
    if (argsLen < minLen) {
        throw Error(
            'Expected at least '
                .concat(minLen, ' args but got ')
                .concat(argsLen)
        );
    } else if (argsLen > maxLen) {
        throw Error(
            'Expected at most '.concat(maxLen, ' args but got ').concat(argsLen)
        );
    }
    for (var _i = 0; _i < argsLen; _i++) {
        var arg = args[_i];
        if (_i >= typesLen) {
            validateArg(arg, last(types), _i);
        } else {
            validateArg(arg, types[_i], _i);
        }
    }
};
function validateArg(value, types, num) {
    var isValid = false;
    for (var i = 0, len = types.length; i < len; i++) {
        var t = lowerCase(types[i]);
        if (
            t === 'any' ||
            (t === 'object' && isObj(value)) ||
            type(value) === t
        ) {
            isValid = true;
            break;
        }
    }
    if (!isValid) {
        throw TypeError(
            'Argument '.concat(num, ' should be type ').concat(types.join('|'))
        );
    }
}

module.exports = exports;
