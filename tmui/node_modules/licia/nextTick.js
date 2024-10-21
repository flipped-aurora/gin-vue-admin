if (typeof process === 'object' && process.nextTick && !false) {
    exports = process.nextTick;
} else if (typeof setImmediate === 'function') {
    exports = function(cb) {
        setImmediate(ensureCallable(cb));
    };
} else {
    exports = function(cb) {
        setTimeout(ensureCallable(cb), 0);
    };
}
function ensureCallable(fn) {
    if (typeof fn !== 'function')
        throw new TypeError(fn + ' is not a function');
    return fn;
}

module.exports = exports;
