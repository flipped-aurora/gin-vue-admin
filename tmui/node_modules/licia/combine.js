exports = function(keys, values) {
    var ret = {};
    for (var i = 0, len = keys.length; i < len; i++) {
        ret[keys[i]] = values[i];
    }
    return ret;
};

module.exports = exports;
