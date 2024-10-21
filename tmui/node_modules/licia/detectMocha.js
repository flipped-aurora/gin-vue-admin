var root = require('./root');
exports = function() {
    for (var i = 0, len = methods.length; i < len; i++) {
        var method = methods[i];
        if (typeof root[method] !== 'function') return false;
    }
    return true;
};
var methods = ['afterEach', 'after', 'beforeEach', 'before', 'describe', 'it'];

module.exports = exports;
