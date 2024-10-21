var size = require('./size');
exports = function(arr) {
    if (arr.length < 1) return;
    var ret = arr[0],
        retSize = size(arr[0]);
    for (var i = 1, len = arr.length; i < len; i++) {
        var elSize = size(arr[i]);
        if (elSize > retSize) {
            ret = arr[i];
            retSize = elSize;
        }
    }
    return ret;
};

module.exports = exports;
