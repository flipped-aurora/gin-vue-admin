exports = function(arr) {
    var len = arr ? arr.length : 0;
    if (len) return arr[len - 1];
};

module.exports = exports;
