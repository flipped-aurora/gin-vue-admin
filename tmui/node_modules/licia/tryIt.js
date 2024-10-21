var noop = require('./noop');
exports = function(fn, cb) {
    cb = cb || noop;
    try {
        cb(null, fn());
    } catch (e) {
        cb(e);
        return;
    }
};

module.exports = exports;
