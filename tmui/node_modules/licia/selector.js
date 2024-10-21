var trim = require('./trim');
var each = require('./each');
var identity = require('./identity');
var map = require('./map');

var whitespace = '[\\x20\\t\\r\\n\\f]';
var identifier = '(?:\\\\[\\da-fA-F]{1,6}'.concat(
    whitespace,
    '?|\\\\[^\\r\\n\\f]|[\\w-]|[^\0-\\x7f])+'
);
var attributes = '\\['
    .concat(whitespace, '*(')
    .concat(identifier, ')(?:')
    .concat(whitespace, '*([*^$|!~]?=)')
    .concat(
        whitespace,
        '*(?:\'((?:\\\\.|[^\\\\\'])*)\'|"((?:\\\\.|[^\\\\"])*)"|('
    )
    .concat(identifier, '))|)')
    .concat(whitespace, '*\\]');
var pseudos = '::?('
    .concat(
        identifier,
        ')(?:\\(((\'((?:\\\\.|[^\\\\\'])*)\'|"((?:\\\\.|[^\\\\"])*)")|((?:\\\\.|[^\\\\()[\\]]|'
    )
    .concat(attributes, ')*)|.*)\\)|)');
var regComma = new RegExp('^'.concat(whitespace, '*,').concat(whitespace, '*'));
var regCombinators = new RegExp(
    '^'
        .concat(whitespace, '*([>+~]|')
        .concat(whitespace, ')')
        .concat(whitespace, '*')
);
var matchExpr = {
    id: {
        reg: new RegExp('^#('.concat(identifier, ')')),
        value: function(raw) {
            return raw.slice(1);
        },
        toStr: function(value) {
            return '#'.concat(value);
        }
    },
    class: {
        reg: new RegExp('^\\.('.concat(identifier, ')')),
        value: function(raw) {
            return raw.slice(1);
        },
        toStr: function(value) {
            return '.'.concat(value);
        }
    },
    tag: {
        reg: new RegExp('^('.concat(identifier, '|[*])')),
        value: identity
    },
    attribute: {
        reg: new RegExp('^'.concat(attributes)),
        value: function(raw) {
            return raw.slice(1, raw.length - 1);
        },
        toStr: function(value) {
            return '['.concat(value, ']');
        }
    },
    pseudo: {
        reg: new RegExp('^'.concat(pseudos)),
        value: identity
    }
};
each(matchExpr, function(item) {
    if (!item.value) item.value = identity;
    if (!item.toStr) item.toStr = identity;
});
function parse(selector) {
    selector = trim(selector);
    var groups = [];
    var tokens;
    var match;
    var matched;
    while (selector) {
        if (!matched || (match = regComma.exec(selector))) {
            if (match) {
                selector = selector.slice(match[0].length);
            }
            tokens = [];
            groups.push(tokens);
        }
        matched = false;
        if ((match = regCombinators.exec(selector))) {
            matched = match.shift();
            selector = selector.slice(matched.length);
            matched = trim(matched);
            if (!matched) matched = ' ';
            tokens.push({
                value: matched,
                type: 'combinator'
            });
        }
        each(matchExpr, function(_ref, type) {
            var reg = _ref.reg,
                value = _ref.value;
            if ((match = reg.exec(selector))) {
                matched = match.shift();
                selector = selector.slice(matched.length);
                matched = trim(matched);
                tokens.push({
                    value: value(matched),
                    type: type
                });
            }
        });
        if (!matched) {
            break;
        }
    }
    return groups;
}
function stringify(groups) {
    return map(groups, function(group) {
        group = map(group, function(_ref2) {
            var type = _ref2.type,
                value = _ref2.value;
            if (type === 'combinator') {
                return value === ' ' ? value : ' '.concat(value, ' ');
            }
            return matchExpr[type].toStr(value);
        });
        return group.join('');
    }).join(', ');
}
exports = {
    parse: parse,
    stringify: stringify
};

module.exports = exports;
