var restArgs = require('./restArgs');
var flatten = require('./flatten');
var filter = require('./filter');
var contain = require('./contain');
exports = restArgs(function(arr, args) {
    args = flatten(args);
    return filter(arr, function(val) {
        return !contain(args, val);
    });
});

module.exports = exports;
