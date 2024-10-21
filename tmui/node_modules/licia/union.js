var restArgs = require('./restArgs');
var unique = require('./unique');
var flatten = require('./flatten');
exports = restArgs(function(arrays) {
    return unique(flatten(arrays));
});

module.exports = exports;
