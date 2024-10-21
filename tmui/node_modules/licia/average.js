exports = function() {
    var arr = arguments;
    var sum = 0;
    var len = arr.length;
    for (var i = 0; i < len; i++) sum += arr[i];
    return sum / len;
};

module.exports = exports;
