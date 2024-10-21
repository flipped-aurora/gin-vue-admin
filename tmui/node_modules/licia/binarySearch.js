var isSorted = require('./isSorted');
exports = function(arr, val) {
    var cmp =
        arguments.length > 2 && arguments[2] !== undefined
            ? arguments[2]
            : isSorted.defComparator;
    var startIdx = 0;
    var endIdx = arr.length - 1;
    while (startIdx <= endIdx) {
        var middleIdx = startIdx + Math.floor((endIdx - startIdx) / 2);
        var middleVal = arr[middleIdx];
        if (cmp(middleVal, val) === 0) {
            return middleIdx;
        }
        if (cmp(middleVal, val) < 0) {
            startIdx = middleIdx + 1;
        } else {
            endIdx = middleIdx - 1;
        }
    }
    return -1;
};

module.exports = exports;
