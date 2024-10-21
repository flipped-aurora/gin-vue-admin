var unique = require('./unique');
var trim = require('./trim');
var map = require('./map');
var toArr = require('./toArr');
exports = function(str) {
    var urlList = toArr(str.match(regUrl));
    return unique(
        map(urlList, function(url) {
            return trim(url);
        })
    );
};
var regUrl = /((https?)|(ftp)):\/\/[\w.]+[^ \f\n\r\t\v"\\<>[\]\u2100-\uFFFF(),]*/gi;

module.exports = exports;
