var isFn = require('./isFn');
exports = function(obj) {
    var ret = [];
    for (var key in obj) {
        if (isFn(obj[key])) ret.push(key);
    }
    return ret.sort();
};

module.exports = exports;
