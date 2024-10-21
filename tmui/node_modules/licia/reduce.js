var optimizeCb = require('./optimizeCb');
var isArrLike = require('./isArrLike');
var isUndef = require('./isUndef');
var keys = require('./keys');
exports = createReduce(1);
exports.create = createReduce;
function createReduce(dir) {
    return function(obj, iterator, initial, ctx) {
        iterator = optimizeCb(iterator, ctx);
        var i, len, key;
        if (isArrLike(obj)) {
            len = obj.length;
            i = dir > 0 ? 0 : len - 1;
            if (isUndef(initial)) {
                initial = obj[i];
                i += dir;
            }
            for (; i < len && i >= 0; i += dir) {
                initial = iterator(initial, obj[i], i, obj);
            }
        } else {
            var _keys = keys(obj);
            len = _keys.length;
            i = dir > 0 ? 0 : len - 1;
            if (isUndef(initial)) {
                initial = obj[_keys[i]];
                i += dir;
            }
            for (; i < len && i >= 0; i += dir) {
                key = _keys[i];
                initial = iterator(initial, obj[key], key, obj);
            }
        }
        return initial;
    };
}

module.exports = exports;
