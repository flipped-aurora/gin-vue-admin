exports = function(fn, wait, immediate) {
    var timeout;
    return function() {
        var ctx = this;
        var args = arguments;
        var throttler = function() {
            timeout = null;
            fn.apply(ctx, args);
        };
        if (!immediate) clearTimeout(timeout);
        if (!immediate || !timeout) timeout = setTimeout(throttler, wait);
    };
};

module.exports = exports;
