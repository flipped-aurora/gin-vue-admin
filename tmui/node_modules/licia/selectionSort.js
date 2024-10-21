var swap = require('./swap');
var isSorted = require('./isSorted');
exports = function(arr) {
    var cmp =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : isSorted.defComparator;
    var min;
    for (var i = 0, len = arr.length; i < len; i++) {
        min = i;
        for (var j = i + 1; j < len; j++) {
            if (cmp(arr[j], arr[min]) < 0) {
                min = j;
            }
        }
        if (i != min) {
            swap(arr, i, min);
        }
    }
    return arr;
};

module.exports = exports;
