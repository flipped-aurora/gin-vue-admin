var isArr = require('./isArr');
var safeGet = require('./safeGet');
exports = function(path) {
    if (!isArr(path)) return shallowProperty(path);
    return function(obj) {
        return safeGet(obj, path);
    };
};
function shallowProperty(key) {
    return function(obj) {
        return obj == null ? void 0 : obj[key];
    };
}

module.exports = exports;
