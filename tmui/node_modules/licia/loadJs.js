exports = function(src, cb) {
    var script = document.createElement('script');
    script.src = src;
    script.onload = function() {
        var isNotLoaded =
            script.readyState &&
            script.readyState != 'complete' &&
            script.readyState != 'loaded';
        cb && cb(!isNotLoaded);
    };
    script.onerror = function() {
        cb(false);
    };
    document.body.appendChild(script);
};

module.exports = exports;
