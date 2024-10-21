exports = function() {
    var orig = Error.prepareStackTrace;
    Error.prepareStackTrace = function(_, stack) {
        return stack;
    };
    var stack = new Error().stack.slice(1);
    Error.prepareStackTrace = orig;
    return stack;
};

module.exports = exports;
