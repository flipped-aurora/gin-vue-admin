var delegate = require('./delegate');
var isUndef = require('./isUndef');
var $safeEls = require('./$safeEls');
var each = require('./each');
exports = {
    on: eventFactory('add'),
    off: eventFactory('remove')
};
function eventFactory(type) {
    return function(nodes, event, selector, handler) {
        nodes = $safeEls(nodes);
        if (isUndef(handler)) {
            handler = selector;
            selector = undefined;
        }
        each(nodes, function(node) {
            delegate[type](node, event, selector, handler);
        });
    };
}

module.exports = exports;
