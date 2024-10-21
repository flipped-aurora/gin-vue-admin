var defaults = require('./defaults');
var isNum = require('./isNum');
var isUndef = require('./isUndef');
var decodeUriComponent = require('./decodeUriComponent');
var defOpts = {
    path: '/'
};
function setCookie(key, val, options) {
    if (!isUndef(val)) {
        options = options || {};
        options = defaults(options, defOpts);
        if (isNum(options.expires)) {
            var expires = new Date();
            expires.setMilliseconds(
                expires.getMilliseconds() + options.expires * 864e5
            );
            options.expires = expires;
        }
        val = encodeURIComponent(val);
        key = encodeURIComponent(key);
        document.cookie = [
            key,
            '=',
            val,
            options.expires && '; expires=' + options.expires.toUTCString(),
            options.path && '; path=' + options.path,
            options.domain && '; domain=' + options.domain,
            options.secure ? '; secure' : ''
        ].join('');
        return exports;
    }
    var cookies = document.cookie ? document.cookie.split('; ') : [];
    var result = key ? undefined : {};
    for (var i = 0, len = cookies.length; i < len; i++) {
        var c = cookies[i];
        var parts = c.split('=');
        var name = decodeUriComponent(parts.shift());
        c = parts.join('=');
        c = decodeUriComponent(c);
        if (key === name) {
            result = c;
            break;
        }
        if (!key) result[name] = c;
    }
    return result;
}
exports = {
    get: setCookie,
    set: setCookie,
    remove: function(key, options) {
        options = options || {};
        options.expires = -1;
        return setCookie(key, '', options);
    }
};

module.exports = exports;
