var keys = require('./keys');
exports = function(obj) {
    var _keys = keys(obj);
    var len = _keys.length;
    var ret = Array(len);
    for (var i = 0; i < len; i++) {
        ret[i] = [_keys[i], obj[_keys[i]]];
    }
    return ret;
};

module.exports = exports;
