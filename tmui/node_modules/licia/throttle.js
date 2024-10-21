var debounce = require('./debounce');
exports = function(fn, wait) {
    return debounce(fn, wait, true);
};

module.exports = exports;
