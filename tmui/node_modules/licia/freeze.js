var keys = require('./keys');
exports = function(obj) {
    if (Object.freeze) return Object.freeze(obj);
    keys(obj).forEach(function(prop) {
        if (!Object.getOwnPropertyDescriptor(obj, prop).configurable) return;
        Object.defineProperty(obj, prop, {
            writable: false,
            configurable: false
        });
    });
    return obj;
};

module.exports = exports;
