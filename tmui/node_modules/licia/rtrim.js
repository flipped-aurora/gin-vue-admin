exports = function(str, chars) {
    if (chars == null) {
        if (str.trimRight) {
            return str.trimRight();
        }
        chars = ' \r\n\t\f\v';
    }
    var end = str.length - 1;
    var charLen = chars.length;
    var found = true;
    var i;
    var c;
    while (found && end >= 0) {
        found = false;
        i = -1;
        c = str.charAt(end);
        while (++i < charLen) {
            if (c === chars[i]) {
                found = true;
                end--;
                break;
            }
        }
    }
    return end >= 0 ? str.substring(0, end + 1) : '';
};

module.exports = exports;
