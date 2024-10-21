var Class = require('./Class');
var swap = require('./swap');
var isSorted = require('./isSorted');
exports = Class({
    initialize: function Heap() {
        var cmp =
            arguments.length > 0 && arguments[0] !== undefined
                ? arguments[0]
                : isSorted.defComparator;
        this._cmp = cmp;
        this.clear();
    },
    clear: function() {
        this._data = [];
        this.size = 0;
    },
    add: function(item) {
        this._data.push(item);
        this.size++;
        this._heapifyUp(this.size - 1);
        return this.size;
    },
    poll: function() {
        var data = this._data;
        if (this.size > 0) {
            var item = data[0];
            data[0] = data[this.size - 1];
            this.size--;
            this._heapifyDown(0);
            return item;
        }
    },
    peek: function() {
        if (this.size > 0) {
            return this._data[0];
        }
    },
    _heapifyUp: function(idx) {
        var data = this._data;
        var parent = parentIdx(idx);
        while (idx > 0 && this._cmp(data[parent], data[idx]) > 0) {
            swap(data, parent, idx);
            idx = parent;
            parent = parentIdx(idx);
        }
    },
    _heapifyDown: function(idx) {
        var size = this.size;
        var cmp = this._cmp;
        var data = this._data;
        while (leftChildIdx(idx) < size) {
            var smallerIdx = leftChildIdx(idx);
            var rightChild = rightChildIdx(idx);
            if (
                rightChild < size &&
                cmp(data[rightChildIdx], data[smallerIdx]) < 0
            ) {
                smallerIdx = rightChild;
            }
            if (cmp(data[idx], data[smallerIdx]) < 0) {
                break;
            } else {
                swap(data, idx, smallerIdx);
            }
            idx = smallerIdx;
        }
    }
});
function parentIdx(idx) {
    return Math.floor((idx - 1) / 2);
}
function leftChildIdx(idx) {
    return 2 * idx + 1;
}
function rightChildIdx(idx) {
    return 2 * idx + 2;
}

module.exports = exports;
