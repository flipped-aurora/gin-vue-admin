var isArrLike = require('./isArrLike');
var isArr = require('./isArr');
var isStr = require('./isStr');
var isArgs = require('./isArgs');
var keys = require('./keys');
exports = function(val) {
    if (val == null) return true;
    if (isArrLike(val) && (isArr(val) || isStr(val) || isArgs(val))) {
        return val.length === 0;
    }
    return keys(val).length === 0;
};

module.exports = exports;
