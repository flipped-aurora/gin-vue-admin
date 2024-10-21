exports = function() {
    var options =
        arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};
    return new Promise(function(resolve) {
        var input = document.createElement('input');
        input.style.position = 'fixed';
        input.style.bottom = '0';
        input.style.left = '0';
        input.style.visibility = 'hidden';
        input.setAttribute('type', 'file');
        if (options.accept) {
            input.setAttribute('accept', options.accept);
        }
        if (options.multiple) {
            input.setAttribute('multiple', '');
        }
        document.body.appendChild(input);
        input.addEventListener('change', function() {
            document.body.removeChild(input);
            resolve(input.files);
        });
        input.click();
    });
};

module.exports = exports;
