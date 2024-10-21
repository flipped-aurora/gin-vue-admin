var safeCb = require('./safeCb');
var keys = require('./keys');
var isArrLike = require('./isArrLike');
exports = function(obj, iterator, ctx) {
    iterator = safeCb(iterator, ctx);
    var _keys = !isArrLike(obj) && keys(obj);
    var len = (_keys || obj).length;
    var results = Array(len);
    for (var i = 0; i < len; i++) {
        var curKey = _keys ? _keys[i] : i;
        results[i] = iterator(obj[curKey], curKey, obj);
    }
    return results;
};

module.exports = exports;
