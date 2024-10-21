var each = require('./each');
exports = function(obj) {
    var ret = {};
    each(obj, function(val, key) {
        ret[val] = key;
    });
    return ret;
};

module.exports = exports;
