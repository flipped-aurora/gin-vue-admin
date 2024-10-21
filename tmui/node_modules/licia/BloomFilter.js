var Class = require('./Class');
var fill = require('./fill');
var fnv1a = require('./fnv1a');
var strHash = require('./strHash');
var each = require('./each');
var some = require('./some');
exports = Class({
    initialize: function() {
        var size =
            arguments.length > 0 && arguments[0] !== undefined
                ? arguments[0]
                : 1024;
        var k =
            arguments.length > 1 && arguments[1] !== undefined
                ? arguments[1]
                : 3;
        this._buckets = fill(new Array(size), 0);
        this._k = k;
        this._size = size;
    },
    add: function(val) {
        var _this = this;
        each(this._locations(val), function(location) {
            return (_this._buckets[location] = 1);
        });
    },
    test: function(val) {
        var _this2 = this;
        return !some(this._locations(val), function(location) {
            return _this2._buckets[location] === 0;
        });
    },
    _locations: function(val) {
        var ret = [];
        var size = this._size;
        var a = fnv1a(val);
        var b = strHash(val);

        for (var i = 0; i < this._k; i++) {
            ret[i] = (a + b * i) % size;
        }
        return ret;
    }
});

module.exports = exports;
