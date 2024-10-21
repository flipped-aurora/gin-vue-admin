exports = {
    encode: function(bytes) {
        var ret = [];
        var len = bytes.length;
        var remain = len % 3;
        len = len - remain;
        for (var i = 0; i < len; i += 3) {
            ret.push(
                numToBase64(
                    (bytes[i] << 16) + (bytes[i + 1] << 8) + bytes[i + 2]
                )
            );
        }
        len = bytes.length;
        var tmp;
        if (remain === 1) {
            tmp = bytes[len - 1];
            ret.push(code[tmp >> 2]);
            ret.push(code[(tmp << 4) & 0x3f]);
            ret.push('==');
        } else if (remain === 2) {
            tmp = (bytes[len - 2] << 8) + bytes[len - 1];
            ret.push(code[tmp >> 10]);
            ret.push(code[(tmp >> 4) & 0x3f]);
            ret.push(code[(tmp << 2) & 0x3f]);
            ret.push('=');
        }
        return ret.join('');
    },
    decode: function(str) {
        var len = str.length,
            remain = 0;
        if (str[len - 2] === '=') remain = 2;
        else if (str[len - 1] === '=') remain = 1;
        var ret = new Array((len * 3) / 4 - remain);
        len = remain > 0 ? len - 4 : len;
        var i, j;
        for (i = 0, j = 0; i < len; i += 4) {
            var num = base64ToNum(str[i], str[i + 1], str[i + 2], str[i + 3]);
            ret[j++] = (num >> 16) & 0xff;
            ret[j++] = (num >> 8) & 0xff;
            ret[j++] = num & 0xff;
        }
        var tmp;
        if (remain === 2) {
            tmp =
                (codeMap[str.charCodeAt(i)] << 2) |
                (codeMap[str.charCodeAt(i + 1)] >> 4);
            ret[j++] = tmp & 0xff;
        } else if (remain === 1) {
            tmp =
                (codeMap[str.charCodeAt(i)] << 10) |
                (codeMap[str.charCodeAt(i + 1)] << 4) |
                (codeMap[str.charCodeAt(i + 2)] >> 2);
            ret[j++] = (tmp >> 8) & 0xff;
            ret[j++] = tmp & 0xff;
        }
        return ret;
    }
};
var codeMap = [];
var code = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/';
for (var i = 0, len = code.length; i < len; i++) {
    codeMap[code.charCodeAt(i)] = i;
}
function numToBase64(num) {
    return (
        code[(num >> 18) & 0x3f] +
        code[(num >> 12) & 0x3f] +
        code[(num >> 6) & 0x3f] +
        code[num & 0x3f]
    );
}
function base64ToNum(str1, str2, str3, str4) {
    return (
        (codeMap[str1.charCodeAt(0)] << 18) |
        (codeMap[str2.charCodeAt(0)] << 12) |
        (codeMap[str3.charCodeAt(0)] << 6) |
        codeMap[str4.charCodeAt(0)]
    );
}

module.exports = exports;
