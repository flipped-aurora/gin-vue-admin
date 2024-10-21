var keys = require('./keys');
exports = function(str) {
    return regTest.test(str) ? str.replace(regReplace, replaceFn) : str;
};
var map = (exports.map = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#x27;',
    '`': '&#x60;'
});
var regSrc = '(?:' + keys(map).join('|') + ')';
var regTest = new RegExp(regSrc);
var regReplace = new RegExp(regSrc, 'g');
var replaceFn = function(match) {
    return map[match];
};

module.exports = exports;
