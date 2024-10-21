var isStr = require('./isStr');
var toArr = require('./toArr');
var min = require('./min');
var map = require('./map');
var trim = require('./trim');
exports = function(literals) {
    if (isStr(literals)) literals = toArr(literals);
    var str = '';
    for (
        var _len = arguments.length,
            placeholders = new Array(_len > 1 ? _len - 1 : 0),
            _key = 1;
        _key < _len;
        _key++
    ) {
        placeholders[_key - 1] = arguments[_key];
    }
    for (var i = 0, len = literals.length; i < len; i++) {
        str += literals[i];
        if (placeholders[i]) str += placeholders[i];
    }
    var lines = str.split('\n');
    var indentLens = [];
    for (var _i = 0, _len2 = lines.length; _i < _len2; _i++) {
        var line = lines[_i];
        var _indent = line.match(regStartSpaces);
        if (_indent) {
            indentLens.push(_indent[1].length);
        }
    }
    var indent = indentLens.length > 0 ? min.apply(null, indentLens) : 0;
    return trim(
        map(lines, function(line) {
            return line[0] === ' ' ? line.slice(indent) : line;
        }).join('\n')
    );
};
var regStartSpaces = /^(\s+)\S+/;

module.exports = exports;
