var partial = require('./partial');
exports = function(fn, wrapper) {
    return partial(wrapper, fn);
};

module.exports = exports;
