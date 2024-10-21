var upperCase = require('./upperCase');
var invert = require('./invert');
exports = {
    encode: function(txt) {
        var len = txt.length;
        var ret = Array(len);
        for (var i = 0; i < len; i++) {
            var c = upperCase(txt[i]);
            ret[i] = map[c] || '?';
        }
        return ret.join(' ');
    },
    decode: function(morse) {
        var ret = morse.split(' ');
        for (var i = 0, len = ret.length; i < len; i++) {
            ret[i] = decodeMap[ret[i]] || ' ';
        }
        return ret.join('');
    }
};

var map = {
    A: '.-',
    B: '-...',
    C: '-.-.',
    D: '-..',
    E: '.',
    F: '..-.',
    G: '--.',
    H: '....',
    I: '..',
    J: '.---',
    K: '-.-',
    L: '.-..',
    M: '--',
    N: '-.',
    O: '---',
    P: '.--.',
    Q: '--.-',
    R: '.-.',
    S: '...',
    T: '-',
    U: '..-',
    V: '...-',
    W: '.--',
    X: '-..-',
    Y: '-.--',
    Z: '--..',
    Á: '.--.-',
    Ä: '.-.-',
    É: '..-..',
    Ñ: '--.--',
    Ö: '---.',
    Ü: '..--',
    '1': '.----',
    '2': '..---',
    '3': '...--',
    '4': '....-',
    '5': '.....',
    '6': '-....',
    '7': '--...',
    '8': '---..',
    '9': '----.',
    '0': '-----',
    ',': '--..--',
    '.': '.-.-.-',
    '?': '..--..',
    ';': '-.-.-',
    ':': '---...',
    '/': '-..-.',
    '-': '-....-',
    "'": '.----.',
    '()': '-.--.-',
    _: '..--.-',
    '@': '.--.-.',
    ' ': '.......'
};
var decodeMap = invert(map);

module.exports = exports;
