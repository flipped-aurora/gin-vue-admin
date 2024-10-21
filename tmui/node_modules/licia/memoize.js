var has = require('./has');
exports = function(fn, hashFn) {
    var memoize = function(key) {
        var cache = memoize.cache;
        var address = '' + (hashFn ? hashFn.apply(this, arguments) : key);
        if (!has(cache, address)) cache[address] = fn.apply(this, arguments);
        return cache[address];
    };
    memoize.cache = {};
    return memoize;
};

module.exports = exports;
