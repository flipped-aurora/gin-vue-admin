var root = require('./root');
var touchEvents = {
    down: 'touchstart',
    move: 'touchmove',
    up: 'touchend'
};
var mouseEvents = {
    down: 'mousedown',
    move: 'mousemove',
    up: 'mouseup'
};
var pointerEvents = {
    down: 'pointerdown',
    move: 'pointermove',
    up: 'pointerup'
};
var hasPointerSupport = 'PointerEvent' in root;
var hasTouchSupport = 'ontouchstart' in root;
exports = function(type) {
    if (hasPointerSupport) {
        return pointerEvents[type];
    }
    return hasTouchSupport ? touchEvents[type] : mouseEvents[type];
};

module.exports = exports;
