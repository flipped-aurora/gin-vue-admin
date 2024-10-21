var toArr = require('./toArr');
exports = function() {
    var args = toArr(arguments);
    var ret = [];
    for (var i = 0, len = args.length; i < len; i++) {
        ret = ret.concat(toArr(args[i]));
    }
    return ret;
};

module.exports = exports;
