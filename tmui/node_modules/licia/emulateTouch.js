var trigger = require('./trigger');
var root = require('./root');
var hasTouchSupport = 'ontouchstart' in root;
exports = function() {
    var el =
        arguments.length > 0 && arguments[0] !== undefined
            ? arguments[0]
            : document;
    if (hasTouchSupport) return;
    if (el._isEmulated) return;
    el._isEmulated = true;
    el.addEventListener('mousedown', onMouse('touchstart'), true);
    el.addEventListener('mousemove', onMouse('touchmove'), true);
    el.addEventListener('mouseup', onMouse('touchend'), true);
};
function onMouse(type) {
    return function(e) {
        if (e.which !== 1) return;
        trigger(e.target, type, {
            altKey: e.altKey,
            ctrlKey: e.ctrlKey,
            metaKey: e.metaKey,
            shiftKey: e.shiftKey,
            touches: getActiveTouches(e),
            targetTouches: getActiveTouches(e),
            changedTouches: createTouchList(e)
        });
    };
}
function getActiveTouches(e) {
    if (e.type == 'mouseup') return createTouchList();
    return createTouchList(e);
}
function Touch(target, identifier, pos, deltaX, deltaY) {
    deltaX = deltaX || 0;
    deltaY = deltaY || 0;
    this.identifier = identifier;
    this.target = target;
    this.clientX = pos.clientX + deltaX;
    this.clientY = pos.clientY + deltaY;
    this.screenX = pos.screenX + deltaX;
    this.screenY = pos.screenY + deltaY;
    this.pageX = pos.pageX + deltaX;
    this.pageY = pos.pageY + deltaY;
}
function createTouchList(e) {
    var touchList = [];
    touchList.item = function(index) {
        return this[index] || null;
    };
    if (e) {
        touchList.push(new Touch(e.target, 1, e, 0, 0));
    }
    return touchList;
}

module.exports = exports;
