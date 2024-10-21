var Class = require('./Class');
var safeSet = require('./safeSet');
var safeGet = require('./safeGet');
var map = require('./map');
var filter = require('./filter');
var isFn = require('./isFn');
var safeDel = require('./safeDel');
var toArr = require('./toArr');
var each = require('./each');
exports = Class({
    className: 'JsonTransformer',
    initialize: function(data) {
        this._data = data || {};
    },
    set: function(key, val) {
        if (arguments.length === 1) {
            this._data = key;
            return this;
        }
        safeSet(this._data, key, val);
        return this;
    },
    get: function(key) {
        if (key == null) return this._data;
        return safeGet(this._data, key);
    },
    map: function(from, to, fn) {
        if (isFn(from)) return this.set(map(this._data, from, this));
        if (isFn(to)) {
            fn = to;
            to = from;
        }
        return this.set(to, map(this.get(from), fn, this));
    },
    filter: function(from, to, fn) {
        if (isFn(from)) return this.set(filter(this._data, from, this));
        if (isFn(to)) {
            fn = to;
            to = from;
        }
        return this.set(to, filter(this.get(from), fn, this));
    },
    remove: function(keys) {
        keys = toArr(keys);
        var data = this._data;
        each(keys, function(key) {
            safeDel(data, key);
        });
        return this;
    },
    compute: function(from, to, fn) {
        if (isFn(from)) return this.set(from.call(this, this._data));
        if (isFn(to)) return this.set(from, to.call(this, this.get(from)));
        from = map(
            toArr(from),
            function(key) {
                return safeGet(this._data, key);
            },
            this
        );
        return this.set(to, fn.apply(this, from));
    },
    toString: function() {
        return JSON.stringify(this._data);
    }
});

module.exports = exports;
