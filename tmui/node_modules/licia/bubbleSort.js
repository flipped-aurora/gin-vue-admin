var swap = require('./swap');
var isSorted = require('./isSorted');
exports = function(arr) {
    var cmp =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : isSorted.defComparator;
    for (var i = 0, len = arr.length; i < len; i++) {
        for (var j = i; j > 0; j--) {
            if (cmp(arr[j], arr[j - 1]) < 0) {
                swap(arr, j, j - 1);
            }
        }
    }
    return arr;
};

module.exports = exports;
