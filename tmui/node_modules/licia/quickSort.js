var swap = require('./swap');
var isSorted = require('./isSorted');
exports = function(arr) {
    var cmp =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : isSorted.defComparator;
    return quickSort(arr, 0, arr.length - 1, cmp);
};
function quickSort(arr, left, right, cmp) {
    if (arr.length <= 1) return arr;
    var idx = partition(arr, left, right, cmp);
    if (left < idx - 1) quickSort(arr, left, idx - 1, cmp);
    if (idx < right) quickSort(arr, idx, right, cmp);
    return arr;
}
function partition(arr, left, right, cmp) {
    var pivot = arr[floor((right + left) / 2)];
    while (left <= right) {
        while (cmp(arr[left], pivot) < 0) left++;
        while (cmp(arr[right], pivot) > 0) right--;
        if (left <= right) {
            swap(arr, left, right);
            left++;
            right--;
        }
    }
    return left;
}
var floor = Math.floor;

module.exports = exports;
