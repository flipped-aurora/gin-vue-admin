exports = function(predicate) {
    return function() {
        return !predicate.apply(this, arguments);
    };
};

module.exports = exports;
