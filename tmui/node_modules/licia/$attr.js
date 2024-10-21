var toArr = require('./toArr');
var isObj = require('./isObj');
var isStr = require('./isStr');
var each = require('./each');
var isUndef = require('./isUndef');
var $safeEls = require('./$safeEls');
exports = function(els, name, val) {
    els = $safeEls(els);
    var isGetter = isUndef(val) && isStr(name);
    if (isGetter) return getAttr(els[0], name);
    var attrs = name;
    if (!isObj(attrs)) {
        attrs = {};
        attrs[name] = val;
    }
    setAttr(els, attrs);
};
exports.remove = function(els, names) {
    els = $safeEls(els);
    names = toArr(names);
    each(els, function(node) {
        each(names, function(name) {
            node.removeAttribute(name);
        });
    });
};
function getAttr(el, name) {
    return el.getAttribute(name);
}
function setAttr(els, attrs) {
    each(els, function(el) {
        each(attrs, function(val, name) {
            el.setAttribute(name, val);
        });
    });
}

module.exports = exports;
