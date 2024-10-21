var isObj = require('./isObj');
var isArr = require('./isArr');
var isFn = require('./isFn');
var has = require('./has');
exports = function(val) {
    if (!isObj(val)) return false;
    var ctor = val.constructor;
    if (!isFn(ctor)) return false;
    if (!has(ctor.prototype, 'isPrototypeOf')) return false;
    return !isArr(val) && !isFn(val);
};

module.exports = exports;
