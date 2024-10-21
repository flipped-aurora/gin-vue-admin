var filter = require('./filter');
exports = function(arr) {
    return filter(arr, function(val) {
        return !!val;
    });
};

module.exports = exports;
