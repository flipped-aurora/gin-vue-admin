var safeGet = require('./safeGet');
var toStr = require('./toStr');
var regSep = /{{(.*?)}}/g;
exports = function(str, data) {
    return str.replace(regSep, function(match, key) {
        return toStr(safeGet(data, key));
    });
};

module.exports = exports;
