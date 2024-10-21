var random = require('./random');
var isBrowser = require('./isBrowser');
var isNode = require('./isNode');
exports = function(size) {
    var ret = new Uint8Array(size);
    for (var i = 0; i < size; i++) ret[i] = random(0, 255);
    return ret;
};
var crypto;
if (isBrowser) {
    crypto = window.crypto || window.msCrypto;
    if (crypto) {
        exports = function(size) {
            var ret = new Uint8Array(size);
            crypto.getRandomValues(ret);
            return ret;
        };
    }
} else if (isNode) {
    crypto = eval('require')('crypto');
    exports = function(size) {
        return crypto.randomBytes(size);
    };
}

module.exports = exports;
