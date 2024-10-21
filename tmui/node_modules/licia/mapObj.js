var safeCb = require('./safeCb');
var keys = require('./keys');
exports = function(obj, iterator, ctx) {
    iterator = safeCb(iterator, ctx);
    var _keys = keys(obj);
    var len = _keys.length;
    var ret = {};
    for (var i = 0; i < len; i++) {
        var curKey = _keys[i];
        ret[curKey] = iterator(obj[curKey], curKey, obj);
    }
    return ret;
};

module.exports = exports;
