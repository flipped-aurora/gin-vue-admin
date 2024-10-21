var isUndef = require('./isUndef');
var castPath = require('./castPath');
exports = function(obj, path) {
    path = castPath(path, obj);
    var prop, ret;

    while ((prop = path.shift())) {
        ret = obj[prop];
        if (path.length === 0) delete obj[prop];
        obj = ret;
        if (isUndef(obj)) return;
    }
    return ret;
};

module.exports = exports;
