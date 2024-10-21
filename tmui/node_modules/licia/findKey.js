var safeCb = require('./safeCb');
var keys = require('./keys');
exports = function(obj, predicate, ctx) {
    predicate = safeCb(predicate, ctx);
    var _keys = keys(obj);
    var key;
    for (var i = 0, len = _keys.length; i < len; i++) {
        key = _keys[i];
        if (predicate(obj[key], key, obj)) return key;
    }
};

module.exports = exports;
