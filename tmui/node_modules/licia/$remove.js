var each = require('./each');
var $safeEls = require('./$safeEls');
exports = function(els) {
    els = $safeEls(els);
    each(els, function(el) {
        var parent = el.parentNode;
        if (parent) parent.removeChild(el);
    });
};

module.exports = exports;
