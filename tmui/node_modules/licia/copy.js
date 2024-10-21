var extend = require('./extend');
var noop = require('./noop');
exports = function(text, cb) {
    cb = cb || noop;
    var el = document.createElement('textarea');
    var body = document.body;
    extend(el.style, {
        fontSize: '12pt',
        border: '0',
        padding: '0',
        margin: '0',
        position: 'absolute',
        left: '-9999px'
    });
    el.value = text;
    body.appendChild(el);

    el.setAttribute('readonly', '');
    el.select();
    el.setSelectionRange(0, text.length);
    try {
        document.execCommand('copy');
        cb();
    } catch (e) {
        cb(e);
    } finally {
        body.removeChild(el);
    }
};

module.exports = exports;
