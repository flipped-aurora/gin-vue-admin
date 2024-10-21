var $safeEls = require('./$safeEls');
exports = function(els) {
    els = $safeEls(els);
    var el = els[0];
    var clientRect = el.getBoundingClientRect();
    return {
        left: clientRect.left + window.pageXOffset,
        top: clientRect.top + window.pageYOffset,
        width: Math.round(clientRect.width),
        height: Math.round(clientRect.height)
    };
};

module.exports = exports;
