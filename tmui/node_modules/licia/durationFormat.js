var toInt = require('./toInt');
var lpad = require('./lpad');
var toStr = require('./toStr');
var floor = Math.floor;
exports = function(duration) {
    var mask =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : 'hh:mm:ss';
    duration = toInt(duration);
    var d = floor(duration / 86400000);
    var h = floor(duration / 3600000) % 24;
    var m = floor(duration / 60000) % 60;
    var s = floor(duration / 1000) % 60;
    var l = floor(duration) % 1000;
    var flags = {
        d: d,
        h: h,
        hh: padZero(h),
        m: m,
        mm: padZero(m),
        s: s,
        ss: padZero(s),
        l: l,
        ll: padZero(l, 3)
    };
    return mask.replace(regToken, function(match) {
        if (match in flags) return flags[match];
        return match.slice(1, match.length - 1);
    });
};
var padZero = function(str) {
    var len =
        arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 2;
    return lpad(toStr(str), len, '0');
};
var regToken = /d{1,2}|h{1,2}|m{1,2}|s{1,2}|l{1,2}|"[^"]*"|'[^']*'/g;

module.exports = exports;
