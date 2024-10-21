var restArgs = require('./restArgs');
exports = restArgs(function(fn, wait, args) {
    return setTimeout(function() {
        return fn.apply(null, args);
    }, wait);
});

module.exports = exports;
