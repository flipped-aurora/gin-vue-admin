var isNum = require('./isNum');
var isUndef = require('./isUndef');
var repeat = require('./repeat');
var regLineBegin = /^(?!\s*$)/gm;
exports = function(str, char, len) {
    if (isNum(char)) {
        len = char;
        char = ' ';
    }
    if (isUndef(len)) len = 4;
    if (isUndef(char)) char = ' ';
    char = repeat(char, len);
    return str.replace(regLineBegin, char);
};

module.exports = exports;
