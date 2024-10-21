var parseHtml = require('./parseHtml');
var Stack = require('./Stack');
var isArr = require('./isArr');
var each = require('./each');
var isStr = require('./isStr');
var mapObj = require('./mapObj');
function parse(html) {
    var ret = [];
    var stack = new Stack();
    parseHtml(html, {
        start: function(tag, attrs) {
            attrs = mapObj(attrs, function(val) {
                return unescapeQuote(val);
            });
            stack.push({
                tag: tag,
                attrs: attrs
            });
        },
        end: function() {
            var node = stack.pop();
            if (!stack.size) {
                ret.push(node);
                return;
            }
            var lastNode = stack.peek();
            if (!isArr(lastNode.content)) {
                lastNode.content = [];
            }
            lastNode.content.push(node);
        },
        comment: function(text) {
            var comment = '<!--'.concat(text, '-->');
            var lastNode = stack.peek();
            if (!lastNode) {
                ret.push(comment);
                return;
            }
            if (!lastNode.content) lastNode.content = [];
            lastNode.content.push(comment);
        },
        text: function(text) {
            var lastNode = stack.peek();
            if (!lastNode) {
                ret.push(text);
                return;
            }
            if (!lastNode.content) lastNode.content = [];
            lastNode.content.push(text);
        }
    });
    return ret;
}
function stringify(tree) {
    var ret = '';
    if (isArr(tree)) {
        each(tree, function(node) {
            return (ret += stringify(node));
        });
    } else if (isStr(tree)) {
        ret = tree;
    } else {
        ret += '<'.concat(tree.tag);
        each(tree.attrs, function(val, key) {
            return (ret += ' '.concat(key, '="').concat(escapeQuote(val), '"'));
        });
        ret += '>';
        if (tree.content) ret += stringify(tree.content);
        ret += '</'.concat(tree.tag, '>');
    }
    return ret;
}
var unescapeQuote = function(str) {
    return str.replace(/&quot;/g, '"');
};
var escapeQuote = function(str) {
    return str.replace(/"/g, '&quot;');
};
exports = {
    parse: parse,
    stringify: stringify
};

module.exports = exports;
