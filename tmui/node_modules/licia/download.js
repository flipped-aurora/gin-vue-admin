var createUrl = require('./createUrl');
exports = function(data, name, type) {
    type = type || 'text/plain';
    var el = document.createElement('a');
    el.setAttribute(
        'href',
        createUrl(data, {
            type: type
        })
    );
    el.setAttribute('download', name);
    el.addEventListener('click', function(e) {
        e.stopImmediatePropagation();
    });
    document.body.appendChild(el);
    el.click();
    document.body.removeChild(el);
};

module.exports = exports;
