var isStr = require('./isStr');
var invert = require('./invert');
exports = function(val) {
    if (isStr(val)) return codeMap[val];
    return nameMap[val];
};
var codeMap = {
    backspace: 8,
    tab: 9,
    enter: 13,
    shift: 16,
    ctrl: 17,
    alt: 18,
    'pause/break': 19,
    'caps lock': 20,
    esc: 27,
    space: 32,
    'page up': 33,
    'page down': 34,
    end: 35,
    home: 36,
    left: 37,
    up: 38,
    right: 39,
    down: 40,
    insert: 45,
    delete: 46,
    windows: 91,
    'right windows': 92,
    'windows menu': 93,
    'numpad *': 106,
    'numpad +': 107,
    'numpad -': 109,
    'numpad .': 110,
    'numpad /': 111,
    'num lock': 144,
    'scroll lock': 145,
    ';': 186,
    '=': 187,
    ',': 188,
    '-': 189,
    '.': 190,
    '/': 191,
    '`': 192,
    '[': 219,
    '\\': 220,
    ']': 221,
    "'": 222
};

// Lower case chars
for (var i = 97; i < 123; i++) codeMap[String.fromCharCode(i)] = i - 32;
// Numbers
for (var _i = 48; _i < 58; _i++) codeMap[_i - 48] = _i;
// Function keys
for (var _i2 = 1; _i2 < 13; _i2++) codeMap['f' + _i2] = _i2 + 111;
// Numpad keys
for (var _i3 = 0; _i3 < 10; _i3++) codeMap['numpad ' + _i3] = _i3 + 96;
var nameMap = invert(codeMap);

module.exports = exports;
