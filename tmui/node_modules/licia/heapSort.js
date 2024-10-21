var Heap = require('./Heap');
var isSorted = require('./isSorted');
exports = function(arr) {
    var cmp =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : isSorted.defComparator;
    var heap = new Heap(cmp);
    var len = arr.length;
    for (var i = 0; i < len; i++) {
        heap.add(arr[i]);
    }
    for (var _i = 0; _i < len; _i++) {
        arr[_i] = heap.poll();
    }
    return arr;
};

module.exports = exports;
