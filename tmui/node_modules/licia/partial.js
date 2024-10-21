var restArgs = require('./restArgs');
var toArr = require('./toArr');
exports = restArgs(function(fn, partials) {
    return function() {
        var args = [];
        args = args.concat(partials);
        args = args.concat(toArr(arguments));
        return fn.apply(this, args);
    };
});

module.exports = exports;
