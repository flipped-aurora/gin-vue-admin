var isArrLike = require('./isArrLike');
var keys = require('./keys');
var optimizeCb = require('./optimizeCb');
exports = function(obj, iterator, ctx) {
    iterator = optimizeCb(iterator, ctx);
    var i, len;
    if (isArrLike(obj)) {
        for (i = 0, len = obj.length; i < len; i++) iterator(obj[i], i, obj);
    } else {
        var _keys = keys(obj);
        for (i = 0, len = _keys.length; i < len; i++) {
            iterator(obj[_keys[i]], _keys[i], obj);
        }
    }
    return obj;
};

module.exports = exports;
