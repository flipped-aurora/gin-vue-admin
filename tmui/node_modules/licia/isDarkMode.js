var MediaQuery = require('./MediaQuery');
var m = new MediaQuery('(prefers-color-scheme: dark)');
exports = function() {
    return m.isMatch();
};

module.exports = exports;
