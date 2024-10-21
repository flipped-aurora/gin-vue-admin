var freeze = require('./freeze');
var keys = require('./keys');
var isObj = require('./isObj');
exports = function(obj) {
    freeze(obj);
    keys(obj).forEach(function(prop) {
        var val = obj[prop];
        if (isObj(val) && !Object.isFrozen(val)) exports(val);
    });
    return obj;
};

module.exports = exports;
