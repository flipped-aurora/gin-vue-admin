var toStr = require('./toStr');
exports = function(str) {
    return toStr(str).replace(regEscapeChars, function(char) {
        switch (char) {
            case '"':
            case "'":
            case '\\':
                return '\\' + char;
            case '\n':
                return '\\n';
            case '\r':
                return '\\r';
            // Line separator
            case '\u2028':
                return '\\u2028';
            // Paragraph separator
            case '\u2029':
                return '\\u2029';
        }
    });
};
var regEscapeChars = /["'\\\n\r\u2028\u2029]/g;

module.exports = exports;
