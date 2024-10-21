exports = function(requireFn) {
    const cache = {};

    return function(name) {
        return function() {
            return cache[name] || (cache[name] = requireFn(name));
        };
    };
};

module.exports = exports;
