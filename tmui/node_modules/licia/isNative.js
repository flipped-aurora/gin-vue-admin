var isObj = require('./isObj');
var isFn = require('./isFn');
var toSrc = require('./toSrc');
exports = function(val) {
    if (!isObj(val)) return false;
    if (isFn(val)) return regIsNative.test(toSrc(val));

    return regIsHostCtor.test(toSrc(val));
};
var hasOwnProperty = Object.prototype.hasOwnProperty;
var regIsNative = new RegExp(
    '^' +
        toSrc(hasOwnProperty)
            .replace(/[\\^$.*+?()[\]{}|]/g, '\\$&')
            .replace(
                /hasOwnProperty|(function).*?(?=\\\()| for .+?(?=\\\])/g,
                '$1.*?'
            ) +
        '$'
);
var regIsHostCtor = /^\[object .+?Constructor\]$/;

module.exports = exports;
