exports = function(arr, start, end) {
    var len = arr.length;
    if (start == null) {
        start = 0;
    } else if (start < 0) {
        start = Math.max(len + start, 0);
    } else {
        start = Math.min(start, len);
    }
    if (end == null) {
        end = len;
    } else if (end < 0) {
        end = Math.max(len + end, 0);
    } else {
        end = Math.min(end, len);
    }
    var ret = [];
    while (start < end) ret.push(arr[start++]);
    return ret;
};

module.exports = exports;
