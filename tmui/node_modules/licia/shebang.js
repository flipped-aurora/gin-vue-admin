var trim = require('./trim');
var regShebang = /^#!(.*)/;
exports = function(str) {
    var match = str.match(regShebang);
    if (!match) return;
    return trim(match[1]);
};

module.exports = exports;
