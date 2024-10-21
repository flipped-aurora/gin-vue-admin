var each = require('./each');
var $safeEls = require('./$safeEls');
var isStr = require('./isStr');
exports = {
    before: insertFactory('beforebegin'),
    after: insertFactory('afterend'),
    append: insertFactory('beforeend'),
    prepend: insertFactory('afterbegin')
};
function insertFactory(type) {
    return function(nodes, val) {
        nodes = $safeEls(nodes);
        each(nodes, function(node) {
            if (isStr(val)) {
                node.insertAdjacentHTML(type, val);
            } else {
                var parentNode = node.parentNode;
                switch (type) {
                    case 'beforebegin':
                        if (parentNode) {
                            parentNode.insertBefore(val, node);
                        }
                        break;
                    case 'afterend':
                        if (parentNode) {
                            parentNode.insertBefore(val, node.nextSibling);
                        }
                        break;
                    case 'beforeend':
                        node.appendChild(val);
                        break;
                    case 'afterbegin':
                        node.prepend(val);
                        break;
                }
            }
        });
    };
}

module.exports = exports;
