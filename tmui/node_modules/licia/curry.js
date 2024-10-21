var toArr = require('./toArr');
exports = function(fn) {
    var len = fn.length;
    return function curriedFn() {
        var args = toArr(arguments);
        if (args.length < len) {
            return function() {
                return curriedFn.apply(null, args.concat(toArr(arguments)));
            };
        }
        return fn.apply(null, args);
    };
};

module.exports = exports;
