var isObj = require('./isObj');
var isFn = require('./isFn');
var isArr = require('./isArr');
var mapObj = require('./mapObj');
exports = function(obj) {
    if (isArr(obj)) {
        return obj.map(function(val) {
            return exports(val);
        });
    }
    if (isObj(obj) && !isFn(obj)) {
        return mapObj(obj, function(val) {
            return exports(val);
        });
    }
    return obj;
};

module.exports = exports;
