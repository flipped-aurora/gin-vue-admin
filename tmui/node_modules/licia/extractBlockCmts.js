var map = require('./map');
var trim = require('./trim');
var regBlockCmt = /(\/\*[\s\S]*?\*\/)/gm;
exports = function(str) {
    var ret = str.match(regBlockCmt);
    if (!ret) return [];
    ret = map(ret, function(comment) {
        return trim(
            map(comment.split('\n'), function(line) {
                return trim(line).replace(/^\/\*+|\*+\/$|^\*+/g, '');
            }).join('\n')
        );
    });
    return ret;
};

module.exports = exports;
