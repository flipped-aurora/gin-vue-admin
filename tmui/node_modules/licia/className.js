var each = require('./each');
var isStr = require('./isStr');
var isNum = require('./isNum');
var isArr = require('./isArr');
var isObj = require('./isObj');
exports = function() {
    var ret = [];
    each(arguments, function(arg) {
        if (!arg) return;
        if (isStr(arg) || isNum(arg)) return ret.push(arg);
        if (isArr(arg)) return ret.push(exports.apply(null, arg));
        if (!isObj(arg)) return;
        each(arg, function(val, key) {
            if (val) ret.push(key);
        });
    });
    return ret.join(' ');
};

module.exports = exports;
