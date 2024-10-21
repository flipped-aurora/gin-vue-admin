exports = function(str) {
    return exports.v4(str) || exports.v6(str);
};

var v4 =
    '(?:25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]\\d|\\d)(?:\\.(?:25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]\\d|\\d)){3}';
var regV4 = new RegExp('^'.concat(v4, '$'));
var v6seg = '[a-fA-F\\d]{1,4}';
var v6 = [
    '(',
    '(?:'.concat(v6seg, ':){7}(?:').concat(v6seg, '|:)|'),
    '(?:'
        .concat(v6seg, ':){6}(?:')
        .concat(v4, '|:')
        .concat(v6seg, '|:)|'),
    '(?:'
        .concat(v6seg, ':){5}(?::')
        .concat(v4, '|(:')
        .concat(v6seg, '){1,2}|:)|'),
    '(?:'
        .concat(v6seg, ':){4}(?:(:')
        .concat(v6seg, '){0,1}:')
        .concat(v4, '|(:')
        .concat(v6seg, '){1,3}|:)|'),
    '(?:'
        .concat(v6seg, ':){3}(?:(:')
        .concat(v6seg, '){0,2}:')
        .concat(v4, '|(:')
        .concat(v6seg, '){1,4}|:)|'),
    '(?:'
        .concat(v6seg, ':){2}(?:(:')
        .concat(v6seg, '){0,3}:')
        .concat(v4, '|(:')
        .concat(v6seg, '){1,5}|:)|'),
    '(?:'
        .concat(v6seg, ':){1}(?:(:')
        .concat(v6seg, '){0,4}:')
        .concat(v4, '|(:')
        .concat(v6seg, '){1,6}|:)|'),
    '(?::((?::'
        .concat(v6seg, '){0,5}:')
        .concat(v4, '|(?::')
        .concat(v6seg, '){1,7}|:))'),
    ')(%[0-9a-zA-Z]{1,})?'
].join('');
var regV6 = new RegExp('^'.concat(v6, '$'));
exports.v4 = function(str) {
    return regV4.test(str);
};
exports.v6 = function(str) {
    return regV6.test(str);
};

module.exports = exports;
