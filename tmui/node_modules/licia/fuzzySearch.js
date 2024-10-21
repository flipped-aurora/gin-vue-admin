var filter = require('./filter');
var map = require('./map');
var isStr = require('./isStr');
var safeGet = require('./safeGet');
var levenshtein = require('./levenshtein');
var pluck = require('./pluck');
exports = function(needle, haystacks) {
    var options =
        arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : {};
    if (!options.caseSensitive) {
        needle = needle.toLowerCase();
    }
    haystacks = map(haystacks, function(haystack) {
        var string = toStr(haystack, options);
        if (!options.caseSensitive) {
            string = string.toLowerCase();
        }
        return {
            value: haystack,
            levenshtein: levenshtein(needle, string),
            string: string
        };
    });
    haystacks = filter(haystacks, function(haystack) {
        return hasAllLetters(needle, haystack.string, options);
    });
    haystacks.sort(function(a, b) {
        return a.levenshtein - b.levenshtein;
    });
    return pluck(haystacks, 'value');
};
function toStr(haystack, options) {
    if (isStr(haystack)) return haystack;
    return safeGet(haystack, options.key) || '';
}
function hasAllLetters(needle, haystack) {
    var hLen = haystack.length;
    var nLen = needle.length;
    if (nLen > hLen) return false;
    if (nLen === hLen) return needle === haystack;
    for (var i = 0, j = 0; i < nLen; i++) {
        var c = needle.charCodeAt(i);
        var has = false;
        while (j < hLen) {
            if (haystack.charCodeAt(j++) === c) {
                has = true;
                break;
            }
        }
        if (!has) return false;
    }
    return true;
}

module.exports = exports;
