"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getAppStyleIsolation = exports.getAppCodeSpliting = exports.getAppRenderer = void 0;
function getAppRenderer(manifestJson) {
    const platformOptions = manifestJson['app-plus'];
    if (platformOptions && platformOptions.renderer === 'native') {
        return 'native';
    }
    return '';
}
exports.getAppRenderer = getAppRenderer;
function getAppCodeSpliting(manifestJson) {
    if (manifestJson['app-plus']?.optimization?.codeSpliting === true) {
        return true;
    }
    return false;
}
exports.getAppCodeSpliting = getAppCodeSpliting;
function getAppStyleIsolation(manifestJson) {
    return (manifestJson['app-plus']?.optimization?.styleIsolation ?? 'apply-shared');
}
exports.getAppStyleIsolation = getAppStyleIsolation;
