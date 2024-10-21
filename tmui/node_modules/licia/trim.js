var ltrim = require('./ltrim');
var rtrim = require('./rtrim');
exports = function(str, chars) {
    if (chars == null && str.trim) {
        return str.trim();
    }
    return ltrim(rtrim(str, chars), chars);
};

module.exports = exports;
