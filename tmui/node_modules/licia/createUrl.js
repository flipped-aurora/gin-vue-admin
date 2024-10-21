var defaults = require('./defaults');
var isBlob = require('./isBlob');
var isFile = require('./isFile');
var Blob = require('./Blob');
var toArr = require('./toArr');
exports = function(data, options) {
    options = options || {};
    defaults(options, defOpts);
    if (!isBlob(data) && !isFile(data)) {
        data = new Blob(toArr(data), options);
    }
    return URL.createObjectURL(data);
};
var defOpts = {
    type: 'text/plain'
};

module.exports = exports;
