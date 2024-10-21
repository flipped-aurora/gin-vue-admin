var Emitter = require('./Emitter');
var isStr = require('./isStr');
var isObj = require('./isObj');
var each = require('./each');
var toArr = require('./toArr');
exports = Emitter.extend({
    initialize: function Store(data) {
        this.callSuper(Emitter, 'initialize', arguments);
        this._data = data || {};
        this.save(this._data);
    },
    set: function(key, val) {
        var data;
        if (isStr(key)) {
            data = {};
            data[key] = val;
        } else if (isObj(key)) {
            data = key;
        }
        var self = this;
        each(data, function(val, key) {
            var oldVal = self._data[key];
            self._data[key] = val;
            self.emit('change', key, val, oldVal);
        });
        this.save(this._data);
    },
    get: function(key) {
        var data = this._data;
        if (isStr(key)) return data[key];
        var ret = {};
        each(key, function(val) {
            ret[val] = data[val];
        });
        return ret;
    },
    remove: function(key) {
        key = toArr(key);
        var data = this._data;
        each(key, function(val) {
            delete data[val];
        });
        this.save(data);
    },
    clear: function() {
        this._data = {};
        this.save(this._data);
    },
    each: function(fn) {
        each(this._data, fn);
    },

    save: function(data) {
        this._data = data;
    }
});

module.exports = exports;
