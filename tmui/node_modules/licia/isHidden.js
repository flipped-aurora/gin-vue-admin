var root = require('./root');
var getComputedStyle = root.getComputedStyle;
var document = root.document;
exports = function(el) {
    var _ref =
            arguments.length > 1 && arguments[1] !== undefined
                ? arguments[1]
                : {},
        _ref$display = _ref.display,
        display = _ref$display === void 0 ? true : _ref$display,
        _ref$visibility = _ref.visibility,
        visibility = _ref$visibility === void 0 ? false : _ref$visibility,
        _ref$opacity = _ref.opacity,
        opacity = _ref$opacity === void 0 ? false : _ref$opacity,
        _ref$size = _ref.size,
        size = _ref$size === void 0 ? false : _ref$size,
        _ref$viewport = _ref.viewport,
        viewport = _ref$viewport === void 0 ? false : _ref$viewport,
        _ref$overflow = _ref.overflow,
        overflow = _ref$overflow === void 0 ? false : _ref$overflow;
    var computedStyle = getComputedStyle(el);
    if (display) {
        var tagName = el.tagName;
        if (
            tagName === 'BODY' ||
            tagName === 'HTML' ||
            computedStyle.position === 'fixed'
        ) {
            if (computedStyle.display === 'none') {
                return true;
            } else {
                var cur = el;
                while ((cur = cur.parentElement)) {
                    var _computedStyle = getComputedStyle(cur);
                    if (_computedStyle.display === 'none') {
                        return true;
                    }
                }
            }
        } else if (el.offsetParent === null) {
            return true;
        }
    }
    if (visibility && computedStyle.visibility === 'hidden') {
        return true;
    }
    if (opacity) {
        if (computedStyle.opacity === '0') {
            return true;
        } else {
            var _cur = el;
            while ((_cur = _cur.parentElement)) {
                var _computedStyle2 = getComputedStyle(_cur);
                if (_computedStyle2.opacity === '0') {
                    return true;
                }
            }
        }
    }
    var clientRect = el.getBoundingClientRect();
    if (size && (clientRect.width === 0 || clientRect.height === 0)) {
        return true;
    }
    if (viewport) {
        var containerRect = {
            top: 0,
            left: 0,
            right: document.documentElement.clientWidth,
            bottom: document.documentElement.clientHeight
        };
        return isOutside(clientRect, containerRect);
    }
    if (overflow) {
        var _cur2 = el;
        while ((_cur2 = _cur2.parentElement)) {
            var _computedStyle3 = getComputedStyle(_cur2);
            var _overflow = _computedStyle3.overflow;
            if (_overflow === 'scroll' || _overflow === 'hidden') {
                var curRect = _cur2.getBoundingClientRect();
                if (isOutside(clientRect, curRect)) return true;
            }
        }
    }
    return false;
};
function isOutside(clientRect, containerRect) {
    return (
        clientRect.right < containerRect.left ||
        clientRect.left > containerRect.right ||
        clientRect.bottom < containerRect.top ||
        clientRect.top > containerRect.bottom
    );
}

module.exports = exports;
