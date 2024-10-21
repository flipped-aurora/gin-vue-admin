var regSpace = /^\s+/;
exports = function(str, chars) {
    if (chars == null) {
        if (str.trimLeft) {
            return str.trimLeft();
        }
        return str.replace(regSpace, '');
    }
    var start = 0;
    var len = str.length;
    var charLen = chars.length;
    var found = true;
    var i;
    var c;
    while (found && start < len) {
        found = false;
        i = -1;
        c = str.charAt(start);
        while (++i < charLen) {
            if (c === chars[i]) {
                found = true;
                start++;
                break;
            }
        }
    }
    return start >= len ? '' : str.substr(start, len);
};

module.exports = exports;
