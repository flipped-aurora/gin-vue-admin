var Class = require('./Class');
var keys = require('./keys');
var isObj = require('./isObj');
exports = function(val, parents) {
    if (!isObj(val)) {
        return false;
    }
    if (parents && parents.contains(val)) {
        return true;
    }
    parents = new Node(val, parents);
    var _keys = keys(val);
    for (var i = 0, len = _keys.length; i < len; i++) {
        if (exports(val[_keys[i]], parents)) {
            return true;
        }
    }
    return false;
};

var Node = Class({
    initialize: function Node(val, next) {
        this.val = val;
        this.next = next;
    },
    contains: function(val) {
        var cursor = this;
        while (cursor) {
            if (cursor.val === val) {
                return true;
            }
            cursor = cursor.next;
        }
        return false;
    }
});

module.exports = exports;
