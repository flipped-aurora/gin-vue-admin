var swap = require('./swap');
var isSorted = require('./isSorted');
exports = function(arr) {
    var cmp =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : isSorted.defComparator;
    var len = arr.length;
    var gap = Math.floor(len / 2);
    while (gap > 0) {
        for (var i = gap; i <= len - gap; i++) {
            for (var j = i; j > 0; j -= gap) {
                if (cmp(arr[j], arr[j - gap]) < 0) {
                    swap(arr, j, j - gap);
                } else {
                    break;
                }
            }
        }
        gap = Math.floor(gap / 2);
    }
    return arr;
};

module.exports = exports;
