var base64 = require('./base64');
var bytesToStr = require('./bytesToStr');
var strToBytes = require('./strToBytes');
exports = {
    encrypt: function(key, str) {
        return rc4(key, str, false);
    },
    decrypt: function(key, str) {
        return rc4(key, str, true);
    }
};
function rc4(key, str, decrypt) {
    key = strToBytes(key);
    if (!decrypt) {
        str = strToBytes(str);
    } else {
        str = base64.decode(str);
    }
    var result = [];
    var s = [];
    var j = 0;
    var i = 0;
    var x;
    for (i = 0; i < 256; i++) {
        s[i] = i;
    }
    for (i = 0; i < 256; i++) {
        j = (j + s[i] + key[i % key.length]) % 256;
        x = s[i];
        s[i] = s[j];
        s[j] = x;
    }
    i = 0;
    j = 0;
    for (var y = 0, len = str.length; y < len; y++) {
        i = (i + 1) % 256;
        j = (j + s[i]) % 256;
        x = s[i];
        s[i] = s[j];
        s[j] = x;
        result.push(str[y] ^ s[(s[i] + s[j]) % 256]);
    }
    return !decrypt ? base64.encode(result) : bytesToStr(result);
}

module.exports = exports;
