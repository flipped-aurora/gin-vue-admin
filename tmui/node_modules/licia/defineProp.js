var castPath = require('./castPath');
var isStr = require('./isStr');
var isObj = require('./isObj');
var each = require('./each');
exports = function(obj, prop, descriptor) {
    if (isStr(prop)) {
        defineProp(obj, prop, descriptor);
    } else if (isObj(prop)) {
        each(prop, function(descriptor, prop) {
            defineProp(obj, prop, descriptor);
        });
    }
    return obj;
};
function defineProp(obj, prop, descriptor) {
    var path = castPath(prop, obj);
    var lastProp = path.pop();

    while ((prop = path.shift())) {
        if (!obj[prop]) obj[prop] = {};
        obj = obj[prop];
    }
    Object.defineProperty(obj, lastProp, descriptor);
}

module.exports = exports;
