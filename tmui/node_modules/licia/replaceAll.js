var escapeRegExp = require('./escapeRegExp');
exports = function(str, substr, newSubstr) {
    return str.replace(new RegExp(escapeRegExp(substr), 'g'), newSubstr);
};

module.exports = exports;
