var isPlainObj = require('./isPlainObj');
var each = require('./each');
var cloneDeep = require('./cloneDeep');
exports = function(obj) {
    var i = 0;
    var ret = obj;
    var len = arguments.length;
    while (++i < len) {
        obj = arguments[i];
        if (isPlainObj(ret) && isPlainObj(obj)) {
            each(obj, function(val, key) {
                if (
                    key === '__proto__' ||
                    key === 'constructor' ||
                    key === 'prototype'
                ) {
                    return;
                }
                ret[key] = exports(ret[key], obj[key]);
            });
        } else {
            ret = cloneDeep(obj);
        }
    }
    return ret;
};

module.exports = exports;
