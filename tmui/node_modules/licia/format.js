var restArgs = require('./restArgs');
var toInt = require('./toInt');
var toNum = require('./toNum');
var toStr = require('./toStr');
exports = restArgs(function(str, values) {
    var ret = '';
    for (var i = 0, len = str.length; i < len; i++) {
        var c = str[i];
        if (c !== '%' || values.length === 0) {
            ret += c;
            continue;
        }
        i++;
        var val = values.shift();
        switch (str[i]) {
            case 'i':
            case 'd':
                ret += toInt(val);
                break;
            case 'f':
                ret += toNum(val);
                break;
            case 's':
                ret += toStr(val);
                break;
            case 'o':
                ret += tryStringify(val);
                break;
            default:
                i--;
                values.unshift(val);
                ret += c;
        }
    }
    return ret;
});
function tryStringify(obj) {
    try {
        return JSON.stringify(obj);
    } catch (err) {
        return '[Error Stringify]';
    }
}

module.exports = exports;
