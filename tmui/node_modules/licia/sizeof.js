var isArr = require('./isArr');
var keys = require('./keys');
var isBuffer = require('./isBuffer');
var isNull = require('./isNull');

var strSize = 2;
var boolSize = 4;
var numSize = 8;
exports = function(obj) {
    return sizeof(obj, {
        values: []
    });
};
function sizeof(obj, _ref) {
    var values = _ref.values;
    var t = typeof obj;
    if (t === 'string') return obj.length * strSize;
    if (t === 'number') return numSize;
    if (t === 'boolean') return boolSize;
    var size = 0;
    if (t === 'object' && !isNull(obj)) {
        if (values.indexOf(obj) > -1) {
            return 0;
        }
        values.push(obj);
        if (isArr(obj)) {
            for (var i = 0, len = obj.length; i < len; i++) {
                size += sizeof(obj[i], {
                    values: values
                });
            }
        } else {
            var _keys = keys(obj);
            for (var _i = 0, _len = _keys.length; _i < _len; _i++) {
                var key = _keys[_i];
                size += key.length * strSize;
                size += sizeof(obj[key], {
                    values: values
                });
            }
        }
    }
    if (isBuffer(obj)) return obj.length;
    return size;
}

module.exports = exports;
