var isFn = require('./isFn');
var isObj = require('./isObj');
var isArr = require('./isArr');
var optimizeCb = require('./optimizeCb');
var matcher = require('./matcher');
var identity = require('./identity');
var property = require('./property');
exports = function(val, ctx, argCount) {
    if (val == null) return identity;
    if (isFn(val)) return optimizeCb(val, ctx, argCount);
    if (isObj(val) && !isArr(val)) return matcher(val);
    return property(val);
};

module.exports = exports;
