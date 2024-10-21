var Class = require('./Class');
var isStr = require('./isStr');
var each = require('./each');
var mergeArr = require('./mergeArr');
exports = Class({
    className: 'Select',
    initialize: function(selector) {
        this.length = 0;
        if (!selector) return this;
        if (isStr(selector)) return rootSelect.find(selector);
        if (selector.nodeType) {
            this[0] = selector;
            this.length = 1;
        }
    },
    find: function(selector) {
        var ret = new exports();
        this.each(function() {
            mergeArr(ret, this.querySelectorAll(selector));
        });
        return ret;
    },
    each: function(fn) {
        each(this, function(element, idx) {
            fn.call(element, idx, element);
        });
        return this;
    }
});
var rootSelect = new exports(document);

module.exports = exports;
