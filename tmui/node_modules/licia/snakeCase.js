var splitCase = require('./splitCase');
exports = function(str) {
    return splitCase(str).join('_');
};

module.exports = exports;
