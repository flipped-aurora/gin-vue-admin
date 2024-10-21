var memoize = require('./memoize');
var isUndef = require('./isUndef');
var camelCase = require('./camelCase');
exports = memoize(
    function(name, value) {
        if (isUndef(value)) {
            name = camelCase(name);
            return !isUndef(style[name]);
        }
        style.cssText = '';
        style.cssText = name + ':' + value;
        return !!style.length;
    },
    function(name, value) {
        return name + ' ' + value;
    }
);
var style = document.createElement('p').style;

module.exports = exports;
