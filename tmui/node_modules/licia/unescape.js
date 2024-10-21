var escape = require('./escape');
var keys = require('./keys');
var invert = require('./invert');
exports = function(str) {
    return regTest.test(str) ? str.replace(regReplace, replaceFn) : str;
};
var map = invert(escape.map);
var regSrc = '(?:' + keys(map).join('|') + ')';
var regTest = new RegExp(regSrc);
var regReplace = new RegExp(regSrc, 'g');
function replaceFn(match) {
    return map[match];
}

module.exports = exports;
