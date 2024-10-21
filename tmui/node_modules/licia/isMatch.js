var keys = require('./keys');
exports = function(obj, src) {
    var _keys = keys(src);
    var len = _keys.length;
    if (obj == null) return !len;
    obj = Object(obj);
    for (var i = 0; i < len; i++) {
        var key = _keys[i];
        if (src[key] !== obj[key] || !(key in obj)) return false;
    }
    return true;
};

module.exports = exports;
