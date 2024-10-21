"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.formatMiniProgramEvent = void 0;
const uni_shared_1 = require("@dcloudio/uni-shared");
function formatMiniProgramEvent(eventName, { isCatch, isCapture, isComponent, }) {
    if (isComponent) {
        // 自定义组件的自定义事件需要格式化，因为 triggerEvent 时也会格式化
        eventName = (0, uni_shared_1.customizeEvent)(eventName);
    }
    if (!isComponent && eventName === 'click') {
        eventName = 'tap';
    }
    let eventType = 'bind';
    if (isCatch) {
        eventType = 'catch';
    }
    if (isCapture) {
        return `capture-${eventType}:${eventName}`;
    }
    // bind:foo-bar
    return eventType + (isSimpleExpr(eventName) ? '' : ':') + eventName;
}
exports.formatMiniProgramEvent = formatMiniProgramEvent;
function isSimpleExpr(name) {
    if (name.startsWith('_')) {
        return false;
    }
    if (name.indexOf('-') > -1) {
        return false;
    }
    if (name.indexOf(':') > -1) {
        return false;
    }
    return true;
}
