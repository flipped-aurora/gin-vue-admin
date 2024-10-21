var ucs2 = require('./ucs2');

exports = {
    encode: function(str) {
        var codePoints = ucs2.decode(str);
        var byteArr = '';
        for (var i = 0, len = codePoints.length; i < len; i++) {
            byteArr += encodeCodePoint(codePoints[i]);
        }
        return byteArr;
    },
    decode: function(str, safe) {
        byteArr = ucs2.decode(str);
        byteIdx = 0;
        byteCount = byteArr.length;
        codePoint = 0;
        bytesSeen = 0;
        bytesNeeded = 0;
        lowerBoundary = 0x80;
        upperBoundary = 0xbf;
        var codePoints = [];
        var tmp;
        while ((tmp = decodeCodePoint(safe)) !== false) {
            codePoints.push(tmp);
        }
        return ucs2.encode(codePoints);
    }
};
var fromCharCode = String.fromCharCode;
function encodeCodePoint(codePoint) {
    if ((codePoint & 0xffffff80) === 0) {
        return fromCharCode(codePoint);
    }
    var ret = '',
        count,
        offset;

    if ((codePoint & 0xfffff800) === 0) {
        count = 1;
        offset = 0xc0;
    } else if ((codePoint & 0xffff0000) === 0) {
        count = 2;
        offset = 0xe0;
    } else if ((codePoint & 0xffe00000) == 0) {
        count = 3;
        offset = 0xf0;
    }
    ret += fromCharCode((codePoint >> (6 * count)) + offset);
    while (count > 0) {
        var tmp = codePoint >> (6 * (count - 1));
        ret += fromCharCode(0x80 | (tmp & 0x3f));
        count--;
    }
    return ret;
}
var byteArr,
    byteIdx,
    byteCount,
    codePoint,
    bytesSeen,
    bytesNeeded,
    lowerBoundary,
    upperBoundary;
function decodeCodePoint(safe) {
    while (true) {
        if (byteIdx >= byteCount && bytesNeeded) {
            if (safe) return goBack();
            throw new Error('Invalid byte index');
        }
        if (byteIdx === byteCount) return false;
        var byte = byteArr[byteIdx];
        byteIdx++;
        if (!bytesNeeded) {
            if ((byte & 0x80) === 0) {
                return byte;
            }

            if ((byte & 0xe0) === 0xc0) {
                bytesNeeded = 1;
                codePoint = byte & 0x1f;
            } else if ((byte & 0xf0) === 0xe0) {
                if (byte === 0xe0) lowerBoundary = 0xa0;
                if (byte === 0xed) upperBoundary = 0x9f;
                bytesNeeded = 2;
                codePoint = byte & 0xf;
            } else if ((byte & 0xf8) === 0xf0) {
                if (byte === 0xf0) lowerBoundary = 0x90;
                if (byte === 0xf4) upperBoundary = 0x8f;
                bytesNeeded = 3;
                codePoint = byte & 0x7;
            } else {
                if (safe) return goBack();
                throw new Error('Invalid UTF-8 detected');
            }
            continue;
        }
        if (byte < lowerBoundary || byte > upperBoundary) {
            if (safe) {
                byteIdx--;
                return goBack();
            }
            throw new Error('Invalid continuation byte');
        }
        lowerBoundary = 0x80;
        upperBoundary = 0xbf;
        codePoint = (codePoint << 6) | (byte & 0x3f);
        bytesSeen++;
        if (bytesSeen !== bytesNeeded) continue;
        var tmp = codePoint;
        codePoint = 0;
        bytesNeeded = 0;
        bytesSeen = 0;
        return tmp;
    }
}
function goBack() {
    var start = byteIdx - bytesSeen - 1;
    byteIdx = start + 1;
    codePoint = 0;
    bytesNeeded = 0;
    bytesSeen = 0;
    lowerBoundary = 0x80;
    upperBoundary = 0xbf;
    return byteArr[start];
}

module.exports = exports;
