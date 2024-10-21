var safeCb = require('./safeCb');
exports = function(arr, iterator, ctx) {
    var ret = [];
    iterator = safeCb(iterator, ctx);
    var i = -1;
    var len = arr.length;
    while (++i < len) {
        var realIdx = i - ret.length;
        var val = arr[realIdx];
        if (iterator(val, i, arr)) {
            ret.push(val);
            arr.splice(realIdx, 1);
        }
    }
    return ret;
};

module.exports = exports;
