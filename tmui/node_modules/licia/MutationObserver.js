var Class = require('./Class');
exports =
    window.MutationObserver ||
    window.WebKitMutationObserver ||
    window.MozMutationObserver;
if (!exports) {
    exports = Class({
        initialize: function MutationObserver() {},
        observe: function() {},
        disconnect: function() {},
        takeRecords: function() {}
    });
}

module.exports = exports;
