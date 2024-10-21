var contain = require('./contain');
var toArr = require('./toArr');
exports = function(arr) {
    var ret = [];
    var args = toArr(arguments);
    var argsLen = args.length;
    for (var i = 0, len = arr.length; i < len; i++) {
        var item = arr[i];
        if (contain(ret, item)) continue;
        var j = 1;
        for (; j < argsLen; j++) {
            if (!contain(args[j], item)) break;
        }
        if (j === argsLen) ret.push(item);
    }
    return ret;
};

module.exports = exports;
