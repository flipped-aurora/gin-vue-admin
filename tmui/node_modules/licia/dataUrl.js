var isDataUrl = require('./isDataUrl');
var trim = require('./trim');
var endWith = require('./endWith');
var startWith = require('./startWith');
var contain = require('./contain');
var decodeUriComponent = require('./decodeUriComponent');
var defaults = require('./defaults');
var isStr = require('./isStr');
var convertBin = require('./convertBin');
exports = {
    parse: function(dataUrl) {
        if (!isDataUrl(dataUrl)) {
            return null;
        }
        dataUrl = dataUrl.slice('data:'.length);
        var commaIdx = dataUrl.indexOf(',');
        var mime = trim(dataUrl.slice(0, commaIdx));
        var data = trim(dataUrl.slice(commaIdx + 1));
        var base64 = false;
        if (endWith(mime, ';base64')) {
            base64 = true;
            mime = mime.slice(0, -';base64'.length);
        }
        var charset = '';
        if (contain(mime, 'charset=')) {
            charset = mime.split('charset=')[1];
            mime = mime.split(';')[0];
        }
        if (!mime) {
            mime = 'text/plain';
        }
        if (!base64 && startWith(mime, 'text/') && contain(data, '%')) {
            data = decodeUriComponent(data);
        }
        return {
            data: data,
            mime: mime,
            charset: charset,
            base64: base64
        };
    },
    stringify: function(data, mime) {
        var options =
            arguments.length > 2 && arguments[2] !== undefined
                ? arguments[2]
                : {};
        defaults(options, {
            base64: true,
            charset: ''
        });
        var result = 'data:' + mime;
        if (options.charset && startWith(mime, 'text/')) {
            result += ';charset=' + options.charset;
        }
        if (!isStr(data)) {
            data = convertBin(data, 'base64');
            options.base64 = true;
        }
        if (options.base64) {
            result += ';base64';
        } else if (startWith(mime, 'text/') || !mime) {
            data = encodeURIComponent(data);
        }
        return result + ',' + data;
    }
};

module.exports = exports;
