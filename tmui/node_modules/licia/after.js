exports = function(n, fn) {
    return function() {
        if (--n < 1) return fn.apply(this, arguments);
    };
};

module.exports = exports;
