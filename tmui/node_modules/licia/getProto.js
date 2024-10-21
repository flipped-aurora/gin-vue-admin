var isObj = require('./isObj');
var isFn = require('./isFn');
var getPrototypeOf = Object.getPrototypeOf;
var ObjectCtr = {}.constructor;
exports = function(obj) {
    if (!isObj(obj)) return;
    if (getPrototypeOf && !false) return getPrototypeOf(obj);
    var proto = obj.__proto__;
    if (proto || proto === null) return proto;
    if (isFn(obj.constructor)) return obj.constructor.prototype;
    if (obj instanceof ObjectCtr) return ObjectCtr.prototype;
};

module.exports = exports;
