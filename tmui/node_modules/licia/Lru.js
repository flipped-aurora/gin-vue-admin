var LinkedList = require('./LinkedList');
var PseudoMap = require('./PseudoMap');
var Class = require('./Class');
exports = Class({
    initialize: function Lru(max) {
        this._max = max;
        this._list = new LinkedList();
        this._map = new PseudoMap();
    },
    has: function(key) {
        return this._map.has(key);
    },
    remove: function(key) {
        var map = this._map;
        if (this.has(key)) {
            var node = map.get(key);
            this._list.rmNode(node);
            map.delete(key);
        }
    },
    get: function(key) {
        var list = this._list;
        var map = this._map;
        var ret;
        if (this.has(key)) {
            var node = map.get(key);
            ret = node.value.val;
            list.rmNode(node);
            list.unshift(node.value);
            map.set(key, list.head);
        }
        return ret;
    },
    set: function(key, val) {
        var list = this._list;
        var map = this._map;
        if (this.has(key)) {
            var node = map.get(key);
            list.rmNode(node);
            list.unshift({
                key: key,
                val: val
            });
            map.set(key, list.head);
        } else {
            list.unshift({
                key: key,
                val: val
            });
            map.set(key, list.head);
            if (list.size > this._max) {
                var item = list.pop();
                map.delete(item.key);
            }
        }
    },
    clear: function() {
        this._map = new PseudoMap();
        this._list = new LinkedList();
    }
});

module.exports = exports;
