var noop = require('./noop');
var defaults = require('./defaults');
var toArr = require('./toArr');
var isArr = require('./isArr');
var isErr = require('./isErr');
exports = function(fn, options) {
    defaults(options, defOptions);
    return function() {
        var args = toArr(arguments);
        var newArgs = options.before.apply(this, args);
        if (isArr(newArgs)) {
            args = newArgs;
        }
        try {
            var result = fn.apply(this, args);
            var newResult = options.after.call(this, result);
            if (newResult) {
                result = newResult;
            }
            return result;
        } catch (e) {
            var newErr = options.error(e);
            if (newErr) {
                if (isErr(newErr)) {
                    throw newErr;
                } else {
                    return newErr;
                }
            }
            throw e;
        }
    };
};
var defOptions = {
    before: noop,
    after: noop,
    error: noop
};

module.exports = exports;
