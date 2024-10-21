var idxOf = require('./idxOf');
var isStr = require('./isStr');
var isArrLike = require('./isArrLike');
var values = require('./values');
exports = function(arr, val) {
    if (isStr(arr)) return arr.indexOf(val) > -1;
    if (!isArrLike(arr)) arr = values(arr);
    return idxOf(arr, val) >= 0;
};

module.exports = exports;
