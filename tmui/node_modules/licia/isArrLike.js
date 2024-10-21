var isNum = require('./isNum');
var isFn = require('./isFn');
var MAX_ARR_IDX = Math.pow(2, 53) - 1;
exports = function(val) {
    if (!val) return false;
    var len = val.length;
    return isNum(len) && len >= 0 && len <= MAX_ARR_IDX && !isFn(val);
};

module.exports = exports;
