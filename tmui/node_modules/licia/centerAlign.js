var longest = require('./longest');
var isArr = require('./isArr');
var isUndef = require('./isUndef');
var map = require('./map');
var lpad = require('./lpad');
exports = function(str, width) {
    var ret = str;
    if (!isArr(ret)) {
        ret = ret.split(regLineBreak);
    }
    if (isUndef(width)) width = longest(str);
    ret = map(ret, function(str) {
        var len = str.length;
        return lpad(str, floor((width - len) / 2) + len);
    });
    return ret.join('\n');
};
var regLineBreak = /\n/g;
var floor = Math.floor;

module.exports = exports;
