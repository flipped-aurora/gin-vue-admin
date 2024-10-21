var isBrowser = require('./isBrowser');
var toInt = require('./toInt');
var keys = require('./keys');
exports = function(ua) {
    ua = ua || (isBrowser ? navigator.userAgent : '');
    ua = ua.toLowerCase();
    var ieVer = getVer(ua, 'msie ');
    if (ieVer)
        return {
            version: ieVer,
            name: 'ie'
        };
    if (regIe11.test(ua))
        return {
            version: 11,
            name: 'ie'
        };
    for (var i = 0, len = browsers.length; i < len; i++) {
        var name = browsers[i];
        var match = ua.match(regBrowsers[name]);
        if (match == null) continue;
        var version = toInt(match[1].split('.')[0]);
        if (name === 'opera') version = getVer(ua, 'version/') || version;
        return {
            name: name,
            version: version
        };
    }
    return {
        name: 'unknown',
        version: -1
    };
};
var regBrowsers = {
    edge: /edge\/([0-9._]+)/,
    firefox: /firefox\/([0-9.]+)(?:\s|$)/,
    opera: /opera\/([0-9.]+)(?:\s|$)/,
    android: /android\s([0-9.]+)/,
    ios: /version\/([0-9._]+).*mobile.*safari.*/,
    safari: /version\/([0-9._]+).*safari/,
    chrome: /(?!chrom.*opr)chrom(?:e|ium)\/([0-9.]+)(:?\s|$)/
};
var regIe11 = /trident\/7\./;
var browsers = keys(regBrowsers);
function getVer(ua, mark) {
    var idx = ua.indexOf(mark);
    if (idx > -1)
        return toInt(ua.substring(idx + mark.length, ua.indexOf('.', idx)));
}

module.exports = exports;
