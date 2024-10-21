var memoize = require('./memoize');
exports = memoize(function(n) {
    return n < 2 ? n : exports(n - 1) + exports(n - 2);
});

module.exports = exports;
