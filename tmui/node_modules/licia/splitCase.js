var regUpperCase = /([A-Z])/g;
var regSeparator = /[_.\- ]+/g;
var regTrim = /(^-)|(-$)/g;
exports = function(str) {
    str = str
        .replace(regUpperCase, '-$1')
        .toLowerCase()
        .replace(regSeparator, '-')
        .replace(regTrim, '');
    return str.split('-');
};

module.exports = exports;
