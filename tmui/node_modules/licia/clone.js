var isObj = require('./isObj');
var isArr = require('./isArr');
var extend = require('./extend');
exports = function(obj) {
    if (!isObj(obj)) return obj;
    return isArr(obj) ? obj.slice() : extend({}, obj);
};

module.exports = exports;
