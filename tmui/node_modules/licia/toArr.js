var isArrLike = require('./isArrLike');
var map = require('./map');
var isArr = require('./isArr');
var isStr = require('./isStr');
exports = function(val) {
    if (!val) return [];
    if (isArr(val)) return val;
    if (isArrLike(val) && !isStr(val)) return map(val);
    return [val];
};

module.exports = exports;
