var isStr = require('./isStr');
var toArr = require('./toArr');
var Select = require('./Select');
exports = function(val) {
    return toArr(isStr(val) ? new Select(val) : val);
};

module.exports = exports;
