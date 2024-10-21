var isArr = require('./isArr');
exports = function(arr) {
    return flat(arr, []);
};
function flat(arr, res) {
    var len = arr.length,
        i = -1,
        cur;
    while (len--) {
        cur = arr[++i];
        isArr(cur) ? flat(cur, res) : res.push(cur);
    }
    return res;
}

module.exports = exports;
