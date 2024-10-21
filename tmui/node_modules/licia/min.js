exports = function() {
    var arr = arguments;
    var ret = arr[0];
    for (var i = 1, len = arr.length; i < len; i++) {
        if (arr[i] < ret) ret = arr[i];
    }
    return ret;
};

module.exports = exports;
