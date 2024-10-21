var Class = require('./Class');
exports = Class({
    initialize: function LinkedList() {
        this.tail = null;
        this.head = null;
        this.size = 0;
    },
    push: function(val) {
        var node = new Node(val, this.tail, null, this);
        this.tail = node;
        this.head = this.head || node;
        this.size++;
        return this.size;
    },
    pop: function() {
        if (!this.tail) return;
        var node = this.tail;
        this.tail = node.prev;
        if (this.tail) {
            this.tail.next = null;
        } else {
            this.head = null;
        }
        this.size--;
        return node.value;
    },
    unshift: function(val) {
        var node = new Node(val, null, this.head, this);
        this.head = node;
        this.tail = this.tail || node;
        this.size++;
        return this.size;
    },
    shift: function() {
        if (!this.head) return;
        var node = this.head;
        this.head = node.next;
        if (this.head) {
            this.head.prev = null;
        } else {
            this.tail = null;
        }
        this.size--;
        return node.value;
    },
    rmNode: function(node) {
        if (node.list !== this) {
            throw Error('Node does not belong to this list');
        }
        var next = node.next,
            prev = node.prev;
        if (next) {
            next.prev = prev;
        }
        if (prev) {
            prev.next = next;
        }
        if (node === this.head) {
            this.head = next;
        }
        if (node === this.tail) {
            this.tail = prev;
        }
        node.list = null;
        node.prev = null;
        node.next = null;
        this.size--;
    },
    find: function(fn) {
        for (var i = 0, current = this.head; current !== null; i++) {
            if (fn(current.value)) {
                return current;
            }
            current = current.next;
        }
    },
    forEach: function(iterator, ctx) {
        ctx = arguments.length > 1 ? ctx : this;
        for (var i = 0, current = this.head; current !== null; i++) {
            iterator.call(ctx, current.value, i, this);
            current = current.next;
        }
    },
    toArr: function() {
        var arr = new Array(this.size);
        for (var i = 0, current = this.head; current !== null; i++) {
            arr[i] = current.value;
            current = current.next;
        }
        return arr;
    }
});
var Node = (exports.Node = Class({
    initialize: function Node(val, prev, next, list) {
        this.value = val;
        this.list = list;
        if (prev) {
            prev.next = this;
            this.prev = prev;
        } else {
            this.prev = null;
        }
        if (next) {
            next.prev = this;
            this.next = next;
        } else {
            this.next = null;
        }
    }
}));

module.exports = exports;
