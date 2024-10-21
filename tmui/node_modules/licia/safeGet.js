var isUndef = require('./isUndef');
var castPath = require('./castPath');
exports = function(obj, path) {
    path = castPath(path, obj);
    var prop;
    prop = path.shift();
    while (!isUndef(prop)) {
        obj = obj[prop];
        if (obj == null) return;
        prop = path.shift();
    }
    return obj;
};

module.exports = exports;
