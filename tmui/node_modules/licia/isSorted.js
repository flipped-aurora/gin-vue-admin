exports = function(arr) {
    var cmp =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : exports.defComparator;
    for (var i = 0, len = arr.length; i < len - 1; i++) {
        if (cmp(arr[i], arr[i + 1]) > 0) return false;
    }
    return true;
};
exports.defComparator = function(a, b) {
    if (a < b) return -1;
    if (a > b) return 1;
    return 0;
};

module.exports = exports;
