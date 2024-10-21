exports = function(arr, a, b) {
    var tmp = arr[a];
    arr[a] = arr[b];
    arr[b] = tmp;
    return arr;
};

module.exports = exports;
