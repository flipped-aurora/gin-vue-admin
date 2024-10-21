exports = function(start, end, step) {
    if (end == null) {
        end = start || 0;
        start = 0;
    }
    if (!step) step = end < start ? -1 : 1;
    var len = Math.max(Math.ceil((end - start) / step), 0);
    var ret = Array(len);
    for (var i = 0; i < len; i++, start += step) ret[i] = start;
    return ret;
};

module.exports = exports;
