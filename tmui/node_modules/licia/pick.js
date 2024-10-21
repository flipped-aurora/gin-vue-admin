var isStr = require('./isStr');
var isArr = require('./isArr');
var contain = require('./contain');
var each = require('./each');
exports = function(obj, filter, omit) {
    if (isStr(filter)) filter = [filter];
    if (isArr(filter)) {
        var keys = filter;
        filter = function(val, key) {
            return contain(keys, key);
        };
    }
    var ret = {};
    var iteratee = function(val, key) {
        if (filter(val, key)) ret[key] = val;
    };
    if (omit) {
        iteratee = function(val, key) {
            if (!filter(val, key)) ret[key] = val;
        };
    }
    each(obj, iteratee);
    return ret;
};

module.exports = exports;
