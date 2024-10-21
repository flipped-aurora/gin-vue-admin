var Emitter = require('./Emitter');
var State = require('./State');
var easing = require('./easing');
var now = require('./now');
var each = require('./each');
var raf = require('./raf');
var isFn = require('./isFn');
exports = Emitter.extend({
    className: 'Tween',
    initialize: function(target) {
        this.callSuper(Emitter, 'initialize', arguments);
        this._target = target;
        this._dest = {};
        this._duration = 0;
        this._progress = 0;
        this._origin = {};
        this._diff = {};
        this._ease = easing['linear'];
        this._state = new State('pause', {
            play: {
                from: 'pause',
                to: 'play'
            },
            pause: {
                from: 'play',
                to: 'pause'
            }
        });
    },
    to: function(props, duration, ease) {
        var origin = {};
        var target = this._target;
        var diff = {};
        ease = ease || this._ease;
        this._dest = props;
        this._duration = duration || this._duration;
        this._ease = isFn(ease) ? ease : easing[ease];
        each(props, function(val, key) {
            origin[key] = target[key];
            diff[key] = val - origin[key];
        });
        this._origin = origin;
        this._diff = diff;
        return this;
    },
    progress: function(progress) {
        var ease = this._ease;
        var target = this._target;
        var origin = this._origin;
        var diff = this._diff;
        var dest = this._dest;
        var self = this;
        if (progress != null) {
            progress = progress < 1 ? progress : 1;
            this._progress = progress;
            each(dest, function(val, key) {
                target[key] = origin[key] + diff[key] * ease(progress);
            });
            self.emit('update', target);
            return this;
        }
        return this._progress;
    },
    play: function() {
        var state = this._state;
        if (state.is('play')) return;
        state.play();
        var startTime = now();
        var progress = this._progress;
        var duration = this._duration * (1 - progress);
        var target = this._target;
        var self = this;
        function render() {
            if (state.is('pause')) return;
            var time = now();
            self.progress(progress + (time - startTime) / duration);
            if (self._progress === 1) {
                state.pause();
                self.emit('end', target);
                return;
            }
            raf(render);
        }
        raf(render);
        return this;
    },
    pause: function() {
        var state = this._state;
        if (state.is('pause')) return;
        state.pause();
        return this;
    },
    paused: function() {
        return this._state.is('pause');
    }
});

module.exports = exports;
