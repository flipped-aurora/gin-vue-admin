var trim = require('./trim');
var each = require('./each');
var isUndef = require('./isUndef');
var isArr = require('./isArr');
var map = require('./map');
var isEmpty = require('./isEmpty');
var filter = require('./filter');
var isObj = require('./isObj');
exports = {
    parse: function(str) {
        var ret = {};
        str = trim(str).replace(regIllegalChars, '');
        each(str.split('&'), function(param) {
            var parts = param.split('=');
            var key = parts.shift(),
                val = parts.length > 0 ? parts.join('=') : null;
            key = decodeURIComponent(key);
            val = decodeURIComponent(val);
            if (isUndef(ret[key])) {
                ret[key] = val;
            } else if (isArr(ret[key])) {
                ret[key].push(val);
            } else {
                ret[key] = [ret[key], val];
            }
        });
        return ret;
    },
    stringify: function(obj, arrKey) {
        return filter(
            map(obj, function(val, key) {
                if (isObj(val) && isEmpty(val)) return '';
                if (isArr(val)) return exports.stringify(val, key);
                return (
                    (arrKey
                        ? encodeURIComponent(arrKey)
                        : encodeURIComponent(key)) +
                    '=' +
                    encodeURIComponent(val)
                );
            }),
            function(str) {
                return str.length > 0;
            }
        ).join('&');
    }
};
var regIllegalChars = /^(\?|#|&)/g;

module.exports = exports;
