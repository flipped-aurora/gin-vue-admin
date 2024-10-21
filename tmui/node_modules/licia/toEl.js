var doc = document;
exports = function(str) {
    var fragment = doc.createElement('body');
    fragment.innerHTML = str;
    return fragment.childNodes[0];
};
if (doc.createRange && doc.body) {
    var range = doc.createRange();
    range.selectNode(doc.body);
    if (range.createContextualFragment) {
        exports = function(str) {
            return range.createContextualFragment(str).childNodes[0];
        };
    }
}

module.exports = exports;
