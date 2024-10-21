var isStr = require('./isStr');
var isNaN = require('./isNaN');
var isFinite = require('./isFinite');
var isArr = require('./isArr');
exports = function(val) {
    if (isStr(val)) val = val.replace(regComma, '');
    return !isNaN(parseFloat(val)) && isFinite(val) && !isArr(val);
};
var regComma = /,/g;

module.exports = exports;
