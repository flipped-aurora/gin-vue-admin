var isUndef = require('./isUndef');
var each = require('./each');
var $safeEls = require('./$safeEls');
exports = {
    html: propFactory('innerHTML'),
    text: propFactory('textContent'),
    val: propFactory('value')
};
function propFactory(name) {
    return function(nodes, val) {
        nodes = $safeEls(nodes);
        var node = nodes[0];
        if (isUndef(val)) {
            return node ? node[name] : '';
        }
        if (!node) return;
        each(nodes, function(node) {
            node[name] = val;
        });
    };
}

module.exports = exports;
