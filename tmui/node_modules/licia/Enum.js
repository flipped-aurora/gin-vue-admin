var Class = require('./Class');
var freeze = require('./freeze');
var isArr = require('./isArr');
var each = require('./each');
var keys = require('./keys');
exports = Class({
    initialize: function Enum(map) {
        if (isArr(map)) {
            this.size = map.length;
            each(
                map,
                function(member, val) {
                    this[member] = val;
                },
                this
            );
        } else {
            this.size = keys(map).length;
            each(
                map,
                function(val, member) {
                    this[member] = val;
                },
                this
            );
        }
        freeze(this);
    }
});

module.exports = exports;
