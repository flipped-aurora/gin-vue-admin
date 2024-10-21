var objToStr = require('./objToStr');
var isNaN = require('./isNaN');
var lowerCase = require('./lowerCase');
var isBuffer = require('./isBuffer');
exports = function(val) {
    var lower =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : true;
    var ret;
    if (val === null) ret = 'Null';
    if (val === undefined) ret = 'Undefined';
    if (isNaN(val)) ret = 'NaN';
    if (isBuffer(val)) ret = 'Buffer';
    if (!ret) {
        ret = objToStr(val).match(regObj);
        if (ret) ret = ret[1];
    }
    if (!ret) return '';
    return lower ? lowerCase(ret) : ret;
};
var regObj = /^\[object\s+(.*?)]$/;

module.exports = exports;
