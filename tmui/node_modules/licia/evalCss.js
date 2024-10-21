exports = function(css) {
    var style = document.createElement('style');
    style.textContent = css;
    style.type = 'text/css';
    document.head.appendChild(style);
    return style;
};

module.exports = exports;
