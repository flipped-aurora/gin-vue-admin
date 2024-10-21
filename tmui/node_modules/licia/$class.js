var toArr = require('./toArr');
var some = require('./some');
var $safeEls = require('./$safeEls');
var isStr = require('./isStr');
var each = require('./each');
exports = {
    add: function(els, name) {
        els = $safeEls(els);
        var names = safeName(name);
        each(els, function(el) {
            var classList = [];
            each(names, function(name) {
                if (!exports.has(el, name)) classList.push(name);
            });
            if (classList.length !== 0) {
                el.className += (el.className ? ' ' : '') + classList.join(' ');
            }
        });
    },
    has: function(els, name) {
        els = $safeEls(els);
        var regName = new RegExp('(^|\\s)' + name + '(\\s|$)');
        return some(els, function(el) {
            return regName.test(el.className);
        });
    },
    toggle: function(els, name) {
        els = $safeEls(els);
        each(els, function(el) {
            if (!exports.has(el, name)) return exports.add(el, name);
            exports.remove(el, name);
        });
    },
    remove: function(els, name) {
        els = $safeEls(els);
        var names = safeName(name);
        each(els, function(el) {
            each(names, function(name) {
                el.classList.remove(name);
            });
        });
    }
};
function safeName(name) {
    return isStr(name) ? name.split(/\s+/) : toArr(name);
}

module.exports = exports;
