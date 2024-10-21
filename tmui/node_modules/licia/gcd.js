exports = function(a, b) {
    if (b === 0) return a;
    return exports(b, a % b);
};

module.exports = exports;
