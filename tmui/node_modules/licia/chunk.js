exports = function(arr, size) {
    var ret = [];
    size = size || 1;
    for (var i = 0, len = Math.ceil(arr.length / size); i < len; i++) {
        var start = i * size;
        var end = start + size;
        ret.push(arr.slice(start, end));
    }
    return ret;
};

module.exports = exports;
