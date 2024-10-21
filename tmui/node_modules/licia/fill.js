var isUndef = require('./isUndef');
exports = function(arr, val, start, end) {
    var len = arr.length;
    if (!len) return [];
    if (isUndef(end)) end = len;
    if (isUndef(start)) start = 0;
    while (start < end) arr[start++] = val;
    return arr;
};

module.exports = exports;
