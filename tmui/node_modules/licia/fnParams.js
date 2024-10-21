var toSrc = require('./toSrc');
var stripCmt = require('./stripCmt');
var startWith = require('./startWith');
var isStr = require('./isStr');
exports = function(fn) {
    var fnStr = stripCmt(isStr(fn) ? fn : toSrc(fn));
    var open;
    var close;
    if (
        !startWith(fnStr, 'async') &&
        !startWith(fnStr, 'function') &&
        !startWith(fnStr, '(')
    ) {
        open = 0;
        close = fnStr.indexOf('=>');
    } else {
        open = fnStr.indexOf('(') + 1;
        close = fnStr.indexOf(')');
    }
    var ret = fnStr.slice(open, close);
    ret = ret.match(regArgNames);
    return ret === null ? [] : ret;
};
var regArgNames = /[^\s,]+/g;

module.exports = exports;
