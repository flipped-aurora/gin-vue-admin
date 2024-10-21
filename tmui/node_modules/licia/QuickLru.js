var isUndef = require('./isUndef');
var Class = require('./Class');
exports = Class({
    initialize: function QuickLru(max) {
        this._max = max;
        this._cache = {};
        this._oldCache = {};
        this._size = 0;
    },
    has: function(key) {
        return !isUndef(this._cache[key]) || !isUndef(this._oldCache[key]);
    },
    remove: function(key) {
        if (!isUndef(this._cache[key])) this._cache[key] = undefined;
        if (!isUndef(this._oldCache[key])) this._oldCache[key] = undefined;
    },
    get: function(key) {
        if (!isUndef(this._cache[key])) {
            return this._cache[key];
        }
        var val = this._oldCache[key];
        if (!isUndef(val)) {
            this._update(key, val);
            return val;
        }
    },
    set: function(key, val) {
        if (!isUndef(this._cache[key])) {
            this._cache[key] = val;
        } else {
            this._update(key, val);
        }
    },
    clear: function() {
        this._cache = {};
        this._oldCache = {};
    },
    _update: function(key, val) {
        this._cache[key] = val;
        this._size++;
        if (this._size > this._max) {
            this._size = 0;
            this._oldCache = this._cache;
            this._cache = {};
        }
    }
});

module.exports = exports;
