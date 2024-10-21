var each = require('./each');
exports = function(obj) {
    var ret = [];
    each(obj, function(val) {
        ret.push(val);
    });
    return ret;
};

module.exports = exports;
