var Class = require('./Class');
var lowerCase = require('./lowerCase');
var keys = require('./keys');
exports = Class({
    initialize: function(obj) {
        this._target = obj;
    },
    set: function(key, val) {
        var name = this.getKey(key);
        if (name) key = name;
        this._target[key] = val;
    },
    get: function(key) {
        key = this.getKey(key);
        if (key) {
            return this._target[key];
        }
    },
    getKey: function(key) {
        var name = lowerCase(key);
        var _keys = keys(this._target);
        for (var i = 0, len = _keys.length; i < len; i++) {
            var _key = _keys[i];
            if (lowerCase(_key) === name) return _key;
        }
    },
    remove: function(key) {
        delete this._target[this.getKey(key)];
    },
    has: function(key) {
        return !!this.getKey(key);
    }
});

module.exports = exports;
