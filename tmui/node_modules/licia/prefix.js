var memoize = require('./memoize');
var camelCase = require('./camelCase');
var upperFirst = require('./upperFirst');
var has = require('./has');
var kebabCase = require('./kebabCase');
exports = memoize(function(name) {
    name = name.replace(regPrefixes, '');
    name = camelCase(name);
    if (has(style, name)) return name;
    var i = prefixes.length;
    while (i--) {
        var prefixName = prefixes[i] + upperFirst(name);
        if (has(style, prefixName)) return prefixName;
    }
    return name;
});
exports.dash = memoize(function(name) {
    var camelCaseResult = exports(name);
    return (
        (regPrefixes.test(camelCaseResult) ? '-' : '') +
        kebabCase(camelCaseResult)
    );
});
var prefixes = ['O', 'ms', 'Moz', 'Webkit'];
var regPrefixes = /^(O)|(ms)|(Moz)|(Webkit)|(-o-)|(-ms-)|(-moz-)|(-webkit-)/g;
var style = document.createElement('p').style;

module.exports = exports;
