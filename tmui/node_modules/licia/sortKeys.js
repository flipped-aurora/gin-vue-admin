var isSorted = require('./isSorted');
var defaults = require('./defaults');
var keys = require('./keys');
var isArr = require('./isArr');
var isObj = require('./isObj');
exports = function(obj) {
    var options =
        arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
    defaults(options, defOpts);
    var deep = options.deep,
        comparator = options.comparator;
    var visited = [];
    var visitedResult = [];
    function sort(obj) {
        var idx = visited.indexOf(obj);
        if (idx > -1) {
            return visitedResult[idx];
        }
        var result;
        if (isArr(obj)) {
            result = [];
            visited.push(obj);
            visitedResult.push(result);
            for (var i = 0, len = obj.length; i < len; i++) {
                var value = obj[i];
                if (deep && isObj(value)) {
                    result[i] = sort(value);
                } else {
                    result[i] = value;
                }
            }
        } else {
            result = {};
            visited.push(obj);
            visitedResult.push(result);
            var _keys = keys(obj).sort(comparator);
            for (var _i = 0, _len = _keys.length; _i < _len; _i++) {
                var key = _keys[_i];
                var _value = obj[key];
                if (deep && isObj(_value)) {
                    result[key] = sort(_value);
                } else {
                    result[key] = _value;
                }
            }
        }
        return result;
    }
    return sort(obj);
};
var defOpts = {
    deep: false,
    comparator: isSorted.defComparator
};

module.exports = exports;
