var isNum = require('./isNum');
var isObj = require('./isObj');
var isFn = require('./isFn');
var isStr = require('./isStr');
exports = function(val) {
    if (isNum(val)) return val;
    if (isObj(val)) {
        var temp = isFn(val.valueOf) ? val.valueOf() : val;
        val = isObj(temp) ? temp + '' : temp;
    }
    if (!isStr(val)) return val === 0 ? val : +val;
    return +val;
};

module.exports = exports;
