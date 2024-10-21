var random = require('./random');
exports = function(arr) {
    return arr[random(0, arr.length - 1)];
};

module.exports = exports;
