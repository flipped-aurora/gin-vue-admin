var each = require('./each');
var isUndef = require('./isUndef');
var isFn = require('./isFn');
exports = function(arr, val) {
    if (isUndef(val)) val = true;
    var _isFn = isFn(val);
    var ret = {};
    each(arr, function(key) {
        ret[key] = _isFn ? val(key) : val;
    });
    return ret;
};

module.exports = exports;
