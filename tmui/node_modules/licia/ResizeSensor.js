var SingleEmitter = require('./SingleEmitter');
var h = require('./h');
var $event = require('./$event');
var $css = require('./$css');
var contain = require('./contain');
var extend = require('./extend');
var root = require('./root');
if (root.ResizeObserver && !false) {
    exports = SingleEmitter.extend({
        initialize: function ResizeSensor(el) {
            var _this = this;
            if (el._resizeSensor) {
                return el._resizeSensor;
            }
            this.callSuper(SingleEmitter, 'initialize');
            var resizeObserver = new root.ResizeObserver(function() {
                return _this.emit();
            });
            resizeObserver.observe(el);
            el._resizeSensor = this;
            this._resizeObserver = resizeObserver;
            this._el = el;
        },
        destroy: function() {
            var el = this._el;
            if (!el._resizeSensor) {
                return;
            }
            this.rmAllListeners();
            delete el._resizeSensor;
            this._resizeObserver.unobserve(el);
        }
    });
} else {
    exports = SingleEmitter.extend({
        initialize: function ResizeSensor(el) {
            if (el._resizeSensor) {
                return el._resizeSensor;
            }
            this.callSuper(SingleEmitter, 'initialize');
            this._el = el;
            el._resizeSensor = this;
            if (
                !contain(
                    ['absolute', 'relative', 'fixed', 'sticky'],
                    $css(el, 'position')
                )
            ) {
                $css(el, 'position', 'relative');
            }
            this._appendResizeSensor();
            this._bindEvent();
        },
        destroy: function() {
            var el = this._el;
            if (!el._resizeSensor) {
                return;
            }
            this.rmAllListeners();
            delete el._resizeSensor;
            el.removeChild(this._resizeSensorEl);
        },
        _appendResizeSensor: function() {
            var el = this._el;
            var style = {
                pointerEvents: 'none',
                position: 'absolute',
                left: '0px',
                top: '0px',
                right: '0px',
                bottom: '0px',
                overflow: 'hidden',
                zIndex: '-1',
                visibility: 'hidden',
                maxWidth: '100%'
            };
            var styleChild = {
                position: 'absolute',
                left: '0px',
                top: '0px',
                transition: '0s'
            };
            var expandChildEl = h('div', {
                style: styleChild
            });
            var expandEl = h(
                'div.resize-sensor-expand',
                {
                    style: style
                },
                expandChildEl
            );
            var shrinkEl = h(
                'div.resize-sensor-shrink',
                {
                    style: style
                },
                h('div', {
                    style: extend(
                        {
                            width: '200%',
                            height: '200%'
                        },
                        styleChild
                    )
                })
            );
            var resizeSensorEl = h(
                'div.resize-sensor',
                {
                    dir: 'ltr',
                    style: style
                },
                expandEl,
                shrinkEl
            );
            this._expandEl = expandEl;
            this._expandChildEl = expandChildEl;
            this._shrinkEl = shrinkEl;
            this._resizeSensorEl = resizeSensorEl;
            el.appendChild(resizeSensorEl);
            this._resetExpandShrink();
        },
        _bindEvent: function() {
            var _this2 = this;
            $event.on(this._expandEl, 'scroll', function() {
                return _this2._onScroll();
            });
            $event.on(this._shrinkEl, 'scroll', function() {
                return _this2._onScroll();
            });
        },
        _onScroll: function() {
            this.emit();
            this._resetExpandShrink();
        },
        _resetExpandShrink: function() {
            var el = this._el;
            var width = el.offsetWidth;
            var height = el.offsetHeight;
            $css(this._expandChildEl, {
                width: width + 10,
                height: height + 10
            });
            extend(this._expandEl, {
                scrollLeft: width + 10,
                scrollTop: height + 10
            });
            extend(this._shrinkEl, {
                scrollLeft: width + 10,
                scrollTop: height + 10
            });
        }
    });
}

module.exports = exports;
