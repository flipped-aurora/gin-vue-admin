var isSorted = require('./isSorted');
exports = function(arr) {
    var cmp =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : isSorted.defComparator;
    if (arr.length <= 1) return arr;
    var middle = floor(arr.length / 2);
    var left = arr.slice(0, middle);
    var right = arr.slice(middle);
    return merge(exports(left, cmp), exports(right, cmp), cmp);
};
function merge(left, right, cmp) {
    var ret = [];
    var i = 0;
    var j = 0;
    while (i < left.length && j < right.length) {
        cmp(left[i], right[j]) < 0 ? ret.push(left[i++]) : ret.push(right[j++]);
    }
    while (i < left.length) ret.push(left[i++]);
    while (j < right.length) ret.push(right[j++]);
    return ret;
}
var floor = Math.floor;

module.exports = exports;
