exports = function() {
    var arr = arguments;
    var ret = 0;
    for (var i = 0, len = arr.length; i < len; i++) ret += arr[i];
    return ret;
};

module.exports = exports;
