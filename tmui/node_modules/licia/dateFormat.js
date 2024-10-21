var isStr = require('./isStr');
var isDate = require('./isDate');
var toStr = require('./toStr');
var lpad = require('./lpad');
exports = function(date, mask, utc, gmt) {
    if (arguments.length === 1 && isStr(date) && !regNum.test(date)) {
        mask = date;
        date = undefined;
    }
    date = date || new Date();
    if (!isDate(date)) date = new Date(date);
    mask = toStr(exports.masks[mask] || mask || exports.masks['default']);
    var maskSlice = mask.slice(0, 4);
    if (maskSlice === 'UTC:' || maskSlice === 'GMT:') {
        mask = mask.slice(4);
        utc = true;
        if (maskSlice === 'GMT:') gmt = true;
    }
    var prefix = utc ? 'getUTC' : 'get';
    var d = date[prefix + 'Date']();
    var D = date[prefix + 'Day']();
    var m = date[prefix + 'Month']();
    var y = date[prefix + 'FullYear']();
    var H = date[prefix + 'Hours']();
    var M = date[prefix + 'Minutes']();
    var s = date[prefix + 'Seconds']();
    var L = date[prefix + 'Milliseconds']();
    var o = utc ? 0 : date.getTimezoneOffset();
    var flags = {
        d: d,
        dd: padZero(d),
        ddd: exports.i18n.dayNames[D],
        dddd: exports.i18n.dayNames[D + 7],
        m: m + 1,
        mm: padZero(m + 1),
        mmm: exports.i18n.monthNames[m],
        mmmm: exports.i18n.monthNames[m + 12],
        yy: toStr(y).slice(2),
        yyyy: y,
        h: H % 12 || 12,
        hh: padZero(H % 12 || 12),
        H: H,
        HH: padZero(H),
        M: M,
        MM: padZero(M),
        s: s,
        ss: padZero(s),
        l: padZero(L, 3),
        L: padZero(Math.round(L / 10)),
        t: H < 12 ? 'a' : 'p',
        tt: H < 12 ? 'am' : 'pm',
        T: H < 12 ? 'A' : 'P',
        TT: H < 12 ? 'AM' : 'PM',
        Z: gmt
            ? 'GMT'
            : utc
            ? 'UTC'
            : (toStr(date).match(regTimezone) || [''])
                  .pop()
                  .replace(regTimezoneClip, ''),
        o:
            (o > 0 ? '-' : '+') +
            padZero(Math.floor(Math.abs(o) / 60) * 100 + (Math.abs(o) % 60), 4),
        S: ['th', 'st', 'nd', 'rd'][
            d % 10 > 3 ? 0 : (((d % 100) - (d % 10) != 10) * d) % 10
        ]
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
var regToken = /d{1,4}|m{1,4}|yy(?:yy)?|([HhMsTt])\1?|[LloSZWN]|"[^"]*"|'[^']*'/g;
var regTimezone = /\b(?:[PMCEA][SDP]T|(?:Pacific|Mountain|Central|Eastern|Atlantic) (?:Standard|Daylight|Prevailing) Time|(?:GMT|UTC)(?:[-+]\d{4})?)\b/g;
var regNum = /\d/;
var regTimezoneClip = /[^-+\dA-Z]/g;
exports.masks = {
    default: 'ddd mmm dd yyyy HH:MM:ss',
    shortDate: 'm/d/yy',
    mediumDate: 'mmm d, yyyy',
    longDate: 'mmmm d, yyyy',
    fullDate: 'dddd, mmmm d, yyyy',
    shortTime: 'h:MM TT',
    mediumTime: 'h:MM:ss TT',
    longTime: 'h:MM:ss TT Z',
    isoDate: 'yyyy-mm-dd',
    isoTime: 'HH:MM:ss',
    isoDateTime: "yyyy-mm-dd'T'HH:MM:sso",
    isoUtcDateTime: "UTC:yyyy-mm-dd'T'HH:MM:ss'Z'",
    expiresHeaderFormat: 'ddd, dd mmm yyyy HH:MM:ss Z'
};
exports.i18n = {
    dayNames: [
        'Sun',
        'Mon',
        'Tue',
        'Wed',
        'Thu',
        'Fri',
        'Sat',
        'Sunday',
        'Monday',
        'Tuesday',
        'Wednesday',
        'Thursday',
        'Friday',
        'Saturday'
    ],
    monthNames: [
        'Jan',
        'Feb',
        'Mar',
        'Apr',
        'May',
        'Jun',
        'Jul',
        'Aug',
        'Sep',
        'Oct',
        'Nov',
        'Dec',
        'January',
        'February',
        'March',
        'April',
        'May',
        'June',
        'July',
        'August',
        'September',
        'October',
        'November',
        'December'
    ]
};

module.exports = exports;
