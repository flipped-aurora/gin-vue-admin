var isStr = require('./isStr');
var defaults = require('./defaults');
var extend = require('./extend');

exports = function(el, type, opts) {
    if (isStr(el)) {
        opts = type;
        type = el;
        el = document;
    }
    opts = opts || {};
    defaults(opts, defOpts);
    var event = document.createEvent('Event');
    event.initEvent(type, opts.bubbles, opts.cancelable);
    delete opts.bubbles;
    delete opts.cancelable;
    extend(event, opts);
    el.dispatchEvent(event);
};
var defOpts = {
    bubbles: true,
    cancelable: true
};

module.exports = exports;
