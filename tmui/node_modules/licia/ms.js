var toNum = require('./toNum');
var isStr = require('./isStr');
exports = function(str) {
    if (isStr(str)) {
        var match = str.match(regStrTime);
        if (!match) return 0;
        return toNum(match[1]) * factor[match[2] || 'ms'];
    } else {
        var num = str;
        var suffix = 'ms';
        for (var i = 0, len = suffixList.length; i < len; i++) {
            if (num >= factor[suffixList[i]]) {
                suffix = suffixList[i];
                break;
            }
        }
        return +(num / factor[suffix]).toFixed(2) + suffix;
    }
};
var factor = {
    ms: 1,
    s: 1000
};
factor.m = factor.s * 60;
factor.h = factor.m * 60;
factor.d = factor.h * 24;
factor.y = factor.d * 365.25;
var suffixList = ['y', 'd', 'h', 'm', 's'];
var regStrTime = /^((?:\d+)?\.?\d+) *(s|m|h|d|y)?$/;

module.exports = exports;
