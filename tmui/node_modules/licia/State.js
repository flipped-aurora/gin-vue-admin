var Emitter = require('./Emitter');
var each = require('./each');
var some = require('./some');
var toArr = require('./toArr');
exports = Emitter.extend({
    className: 'State',
    initialize: function(initial, events) {
        this.callSuper(Emitter, 'initialize');
        this.current = initial;
        var self = this;
        each(events, function(event, key) {
            self[key] = buildEvent(key, event);
        });
    },
    is: function(state) {
        return this.current === state;
    }
});
function buildEvent(name, event) {
    var from = toArr(event.from);
    var to = event.to;
    return function() {
        var args = toArr(arguments);
        args.unshift(name);
        var hasEvent = some(
            from,
            function(val) {
                return this.current === val;
            },
            this
        );
        if (hasEvent) {
            this.current = to;
            this.emit.apply(this, args);
        } else {
            this.emit(
                'error',
                new Error(this.current + ' => ' + to + ' error'),
                name
            );
        }
    };
}

module.exports = exports;
