var toInt = require('./toInt');
var max = require('./max');
exports = function(v1, v2) {
    v1 = v1.split('.');
    v2 = v2.split('.');
    var len = max(v1.length, v2.length);
    for (var i = 0; i < len; i++) {
        var num1 = toInt(v1[i]);
        var num2 = toInt(v2[i]);
        if (num1 > num2) return 1;
        if (num1 < num2) return -1;
    }
    return 0;
};

module.exports = exports;
