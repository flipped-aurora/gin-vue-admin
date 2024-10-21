var safeCb = require('./safeCb');
exports = function(arr, predicate, ctx, dir) {
    dir = dir || 1;
    predicate = safeCb(predicate, ctx);
    var len = arr.length;
    var i = dir > 0 ? 0 : len - 1;
    while (i >= 0 && i < len) {
        if (predicate(arr[i], i, arr)) return i;
        i += dir;
    }
    return -1;
};

module.exports = exports;
