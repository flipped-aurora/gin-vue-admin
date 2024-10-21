var randomBytes = require('./randomBytes');
var defSymbols =
    'ModuleSymbhasOwnPr-0123456789ABCDEFGHIJKLNQRTUVWXYZ_cfgijkpqtvxz';
exports = function() {
    var size =
        arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 21;
    var symbols =
        arguments.length > 1 && arguments[1] !== undefined
            ? arguments[1]
            : defSymbols;
    var id = '';
    var len = symbols.length;
    var bytes = randomBytes(size);
    while (0 < size--) {
        id += symbols[bytes[size] % len];
    }
    return id;
};

module.exports = exports;
