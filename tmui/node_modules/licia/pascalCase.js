var camelCase = require('./camelCase');
var upperFirst = require('./upperFirst');
exports = function(str) {
    return upperFirst(camelCase(str));
};

module.exports = exports;
