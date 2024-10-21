var Class = require('./Class');
var root = require('./root');
var defineProp = require('./defineProp');
var keys = require('./keys');
var each = require('./each');
var isArr = require('./isArr');
var isUndef = require('./isUndef');
if (root.Map && !false) {
    exports = root.Map;
} else {
    exports = Class({
        initialize: function PseudoMap(data) {
            this.clear();
            var self = this;
            defineProp(this, 'size', {
                get: function() {
                    return keys(self._data).length;
                },
                set: function() {},
                enumerable: true,
                configurable: true
            });
            if (data instanceof exports) {
                data.forEach(function(val, key) {
                    this.set(key, val);
                }, this);
            } else if (isArr(data)) {
                each(
                    data,
                    function(val) {
                        this.set(val[0], val[1]);
                    },
                    this
                );
            }
        },
        forEach: function(fn, ctx) {
            each(
                this._data,
                function(val, key) {
                    fn.call(this, val, key);
                },
                ctx
            );
        },
        has: function(key) {
            return !isUndef(this._data[key]);
        },
        get: function(key) {
            return this._data[key];
        },
        set: function(key, val) {
            this._data[key] = val;
        },
        delete: function(key) {
            delete this._data[key];
        },
        clear: function() {
            this._data = {};
        }
    });
}

module.exports = exports;
