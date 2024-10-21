var splitCase = require('./splitCase');
exports = function(str) {
    return splitCase(str).join('-');
};

module.exports = exports;
