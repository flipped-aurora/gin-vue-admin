var isStr = require('./isStr');
var isObj = require('./isObj');
var kebabCase = require('./kebabCase');
var isUndef = require('./isUndef');
var contain = require('./contain');
var isNum = require('./isNum');
var $safeEls = require('./$safeEls');
var prefix = require('./prefix');
var each = require('./each');
exports = function(nodes, name, val) {
    nodes = $safeEls(nodes);
    var isGetter = isUndef(val) && isStr(name);
    if (isGetter) return getCss(nodes[0], name);
    var css = name;
    if (!isObj(css)) {
        css = {};
        css[name] = val;
    }
    setCss(nodes, css);
};
function getCss(node, name) {
    return (
        node.style[prefix(name)] ||
        getComputedStyle(node, '').getPropertyValue(name)
    );
}
function setCss(nodes, css) {
    each(nodes, function(node) {
        var cssText = ';';
        each(css, function(val, key) {
            key = prefix.dash(key);
            cssText += key + ':' + addPx(key, val) + ';';
        });
        node.style.cssText += cssText;
    });
}
var cssNumProps = [
    'column-count',
    'columns',
    'font-weight',
    'line-weight',
    'opacity',
    'z-index',
    'zoom'
];
function addPx(key, val) {
    var needPx = isNum(val) && !contain(cssNumProps, kebabCase(key));
    return needPx ? val + 'px' : val;
}

module.exports = exports;
