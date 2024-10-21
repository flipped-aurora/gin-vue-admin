var isInt = require('./isInt');
exports = function(num) {
    if (!isInt(num)) return false;
    return num % 2 === 0;
};

module.exports = exports;
