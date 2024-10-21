var isNil = require('./isNil');
exports = function(fn) {
    if (isNil(fn)) return '';
    try {
        return fnToStr.call(fn);
    } catch (e) {}
    try {
        return fn + '';
    } catch (e) {}
    return '';
};
var fnToStr = Function.prototype.toString;

module.exports = exports;
