var each = require('./each');
var $safeEls = require('./$safeEls');
exports = function(els) {
    els = $safeEls(els);
    each(els, function(el) {
        if (isHidden(el)) {
            el.style.display = getDefDisplay(el.nodeName);
        }
    });
};
function isHidden(el) {
    return getComputedStyle(el, '').getPropertyValue('display') == 'none';
}
var elDisplay = {};
function getDefDisplay(elName) {
    var el, display;
    if (!elDisplay[elName]) {
        el = document.createElement(elName);
        document.documentElement.appendChild(el);
        display = getComputedStyle(el, '').getPropertyValue('display');
        el.parentNode.removeChild(el);
        display == 'none' && (display = 'block');
        elDisplay[elName] = display;
    }
    return elDisplay[elName];
}

module.exports = exports;
