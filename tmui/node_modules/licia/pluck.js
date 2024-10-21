var map = require('./map');
var property = require('./property');
exports = function(obj, key) {
    return map(obj, property(key));
};

module.exports = exports;
