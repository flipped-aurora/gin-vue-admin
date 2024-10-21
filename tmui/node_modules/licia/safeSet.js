var castPath = require('./castPath');
var isUndef = require('./isUndef');
var toStr = require('./toStr');
var isSymbol = require('./isSymbol');
var isStr = require('./isStr');
exports = function(obj, path, val) {
    path = castPath(path, obj);
    var lastProp = path.pop();
    var prop;
    prop = path.shift();
    while (!isUndef(prop)) {
        if (!isStr(prop) && !isSymbol(prop)) {
            prop = toStr(prop);
        }
        if (
            prop === '__proto__' ||
            prop === 'constructor' ||
            prop === 'prototype'
        ) {
            return;
        }
        if (!obj[prop]) obj[prop] = {};
        obj = obj[prop];
        prop = path.shift();
    }
    obj[lastProp] = val;
};

module.exports = exports;
