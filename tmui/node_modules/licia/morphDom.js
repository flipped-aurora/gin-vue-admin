var toEl = require('./toEl');
var isStr = require('./isStr');
var h = require('./h');
var isNull = require('./isNull');
var each = require('./each');

exports = function(from, to) {
    if (isStr(to)) {
        to = toEl(to);
    }
    var morphed = from;
    var morphedType = morphed.nodeType;
    var toType = to.nodeType;
    if (morphedType === toType) {
        if (morphedType === ELEMENT_NODE) {
            if (morphed.nodeName !== to.nodeName) {
                morphed = h(to.nodeName);
                moveChildren(from, morphed);
            }
        } else if (morphedType === TEXT_NODE || morphedType === COMMENT_NODE) {
            if (morphed.nodeValue !== to.nodeValue) {
                morphed.nodeValue = to.nodeValue;
            }
            return morphed;
        }
    } else {
        morphed = to;
    }
    if (morphed !== to) {
        morphEl(morphed, to);
    }
    if (from.parentNode) {
        from.parentNode.replaceChild(morphed, from);
    }
    return morphed;
};
var ELEMENT_NODE = 1;
var TEXT_NODE = 3;
var COMMENT_NODE = 8;
function morphEl(from, to) {
    morphAttrs(from, to);
    morphChildren(from, to);
}
function morphAttrs(from, to) {
    var attrs = to.attributes;
    each(attrs, function(_ref) {
        var name = _ref.name,
            value = _ref.value;
        var fromVal = from.getAttribute(name);
        if (fromVal !== value) {
            from.setAttribute(name, value);
        }
    });
    attrs = from.attributes;
    var removedAttrNames = [];
    each(attrs, function(_ref2) {
        var name = _ref2.name;
        if (isNull(to.getAttribute(name))) {
            removedAttrNames.push(name);
        }
    });
    each(removedAttrNames, function(name) {
        return from.removeAttribute(name);
    });
}
function morphChildren(from, to) {
    var curToChild = to.firstChild;
    var curFromChild = from.firstChild;
    var toNextSibling;
    var fromNextSibling;
    outer: while (curToChild) {
        toNextSibling = curToChild.nextSibling;
        while (curFromChild) {
            fromNextSibling = curFromChild.nextSibling;
            var isCompatible = false;
            var curFromType = curFromChild.nodeType;
            var curToType = curToChild.nodeType;
            if (curFromType === curToType) {
                if (curFromType === ELEMENT_NODE) {
                    if (curFromChild.nodeName === curToChild.nodeName) {
                        isCompatible = true;
                        morphEl(curFromChild, curToChild);
                    }
                } else if (
                    curFromType === TEXT_NODE ||
                    curFromType === COMMENT_NODE
                ) {
                    isCompatible = true;
                    if (curFromChild.nodeValue !== curToChild.nodeValue) {
                        curFromChild.nodeValue = curToChild.nodeValue;
                    }
                }
            }
            if (isCompatible) {
                curToChild = toNextSibling;
                curFromChild = fromNextSibling;
                continue outer;
            }
            from.removeChild(curFromChild);
            curFromChild = fromNextSibling;
        }
        from.appendChild(curToChild);
        curFromChild = fromNextSibling;
        curToChild = toNextSibling;
    }
    if (curFromChild) {
        while (curFromChild) {
            fromNextSibling = curFromChild.nextSibling;
            from.removeChild(curFromChild);
            curFromChild = fromNextSibling;
        }
    }
}
function moveChildren(from, to) {
    var curChild = from.firstChild;
    while (curChild) {
        var nextChild = curChild.nextSibling;
        to.appendChild(curChild);
        curChild = nextChild;
    }
}

module.exports = exports;
