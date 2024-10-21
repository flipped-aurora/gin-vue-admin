var optimizeCb = require('./optimizeCb');
exports = function(n, fn, ctx) {
    var ret = Array(Math.max(0, n));
    fn = optimizeCb(fn, ctx, 1);
    for (var i = 0; i < n; i++) {
        ret[i] = fn(i);
    }
    return ret;
};

module.exports = exports;
