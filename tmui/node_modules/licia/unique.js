var filter = require('./filter');
exports = function(arr, cmp) {
    cmp = cmp || isEqual;
    return filter(arr, function(item, idx, arr) {
        var len = arr.length;
        while (++idx < len) {
            if (cmp(item, arr[idx])) return false;
        }
        return true;
    });
};
function isEqual(a, b) {
    return a === b;
}

module.exports = exports;
