var isBrowser = require('./isBrowser');
var isNode = require('./isNode');
exports = function(ua) {
    if (!ua && isBrowser) {
        ua = navigator.userAgent;
    }
    function detect(keyword) {
        return ua.indexOf(keyword) > -1;
    }
    if (ua) {
        ua = ua.toLowerCase();
        if (detect('windows phone')) return 'windows phone';
        if (detect('win')) return 'windows';
        if (detect('android')) return 'android';
        if (detect('ipad') || detect('iphone') || detect('ipod')) return 'ios';
        if (detect('mac')) return 'os x';
        if (detect('linux')) return 'linux';
    } else if (isNode) {
        var _process = process,
            platform = _process.platform,
            env = _process.env;
        if (
            platform === 'win32' ||
            env.OSTYPE === 'cygwin' ||
            env.OSTYPE === 'msys'
        ) {
            return 'windows';
        }
        if (platform === 'darwin') {
            return 'os x';
        }
        if (platform === 'linux') {
            return 'linux';
        }
    }
    return 'unknown';
};

module.exports = exports;
