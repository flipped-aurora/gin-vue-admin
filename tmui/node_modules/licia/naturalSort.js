var startWith = require('./startWith');
var root = require('./root');
var toStr = require('./toStr');
exports = function(arr) {
    return arr.sort(naturalOrderComparator);
};

function naturalOrderComparator(a, b) {
    a = toStr(a);
    b = toStr(b);
    if (startWith(a, '_') && !startWith(b, '_')) {
        return 1;
    }
    if (startWith(b, '_') && !startWith(a, '_')) {
        return -1;
    }
    var chunk = /^\d+|^\D+/;
    var chunka, chunkb, anum, bnum;

    while (true) {
        if (a) {
            if (!b) {
                return 1;
            }
        } else {
            if (b) {
                return -1;
            }
            return 0;
        }
        chunka = a.match(chunk)[0];
        chunkb = b.match(chunk)[0];
        anum = !root.isNaN(chunka);
        bnum = !root.isNaN(chunkb);
        if (anum && !bnum) {
            return -1;
        }
        if (bnum && !anum) {
            return 1;
        }
        if (anum && bnum) {
            var diff = chunka - chunkb;
            if (diff) {
                return diff;
            }
            if (chunka.length !== chunkb.length) {
                if (!+chunka && !+chunkb) {
                    return chunka.length - chunkb.length;
                }
                return chunkb.length - chunka.length;
            }
        } else if (chunka !== chunkb) {
            return chunka < chunkb ? -1 : 1;
        }
        a = a.substring(chunka.length);
        b = b.substring(chunkb.length);
    }
}

module.exports = exports;
