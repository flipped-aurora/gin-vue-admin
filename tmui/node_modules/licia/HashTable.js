var Class = require('./Class');
var LinkedList = require('./LinkedList');
var map = require('./map');
var strHash = require('./strHash');
var has = require('./has');
exports = Class({
    initialize: function HashTable() {
        var size =
            arguments.length > 0 && arguments[0] !== undefined
                ? arguments[0]
                : 32;
        this._buckets = map(Array(size), function() {
            return new LinkedList();
        });
        this._keys = {};
    },
    set: function(key, val) {
        var keyHash = this._hash(key);
        this._keys[key] = keyHash;
        var linkedList = this._buckets[keyHash];
        var node = linkedList.find(function(val) {
            return val.key === key;
        });
        if (!node) {
            linkedList.push({
                key: key,
                value: val
            });
        } else {
            node.value.value = val;
        }
    },
    get: function(key) {
        var linkedList = this._buckets[this._hash(key)];
        var node = linkedList.find(function(val) {
            return val.key === key;
        });
        if (node) {
            return node.value.value;
        }
    },
    has: function(key) {
        return has(this._keys, key);
    },
    delete: function(key) {
        var keyHash = this._hash(key);
        delete this._keys[key];
        var linkedList = this._buckets[keyHash];
        var node = linkedList.find(function(val) {
            return val.key === key;
        });
        if (node) {
            linkedList.rmNode(node);
        }
    },
    _hash: function(key) {
        return strHash(key) % this._buckets.length;
    }
});

module.exports = exports;
