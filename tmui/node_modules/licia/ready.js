var fns = [];
var listener;
var doc = document;
var hack = doc.documentElement.doScroll;
var domContentLoaded = 'DOMContentLoaded';
var loaded = (hack ? /^loaded|^c/ : /^loaded|^i|^c/).test(doc.readyState);
if (!loaded) {
    doc.addEventListener(
        domContentLoaded,
        (listener = function() {
            doc.removeEventListener(domContentLoaded, listener);
            loaded = 1;

            while ((listener = fns.shift())) listener();
        })
    );
}
exports = function(fn) {
    loaded ? setTimeout(fn, 0) : fns.push(fn);
};

module.exports = exports;
