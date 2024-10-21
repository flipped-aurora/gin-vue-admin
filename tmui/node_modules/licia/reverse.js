exports = function(arr) {
    var len = arr.length;
    var ret = Array(len);
    len--;
    for (var i = 0; i <= len; i++) {
        ret[len - i] = arr[i];
    }
    return ret;
};

module.exports = exports;
