var type = require('./type');
var upperFirst = require('./upperFirst');
var toStr = require('./toStr');
var isUndef = require('./isUndef');
var isFn = require('./isFn');
var isRegExp = require('./isRegExp');
exports = function(obj, spaces) {
    return JSON.stringify(obj, serializer(), spaces);
};
function serializer() {
    var stack = [];
    var keys = [];
    return function(key, val) {
        if (stack.length > 0) {
            var pos = stack.indexOf(this);
            if (pos > -1) {
                stack.splice(pos + 1);
                keys.splice(pos, Infinity, key);
            } else {
                stack.push(this);
                keys.push(key);
            }
            var valPos = stack.indexOf(val);
            if (valPos > -1) {
                if (stack[0] === val) {
                    val = '[Circular ~]';
                } else {
                    val =
                        '[Circular ~.' + keys.slice(0, valPos).join('.') + ']';
                }
            }
        } else {
            stack.push(val);
        }
        if (isRegExp(val) || isFn(val)) {
            val = '[' + upperFirst(type(val)) + ' ' + toStr(val) + ']';
        } else if (isUndef(val)) {
            val = null;
        }
        return val;
    };
}

module.exports = exports;
