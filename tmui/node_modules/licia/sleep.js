exports = function(timeout) {
    return new Promise(function(resolve) {
        return setTimeout(resolve, timeout);
    });
};

module.exports = exports;
