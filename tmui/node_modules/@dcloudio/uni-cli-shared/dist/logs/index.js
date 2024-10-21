"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.output = exports.resetOutput = exports.formatWarnMsg = exports.formatInfoMsg = exports.formatErrMsg = void 0;
var format_1 = require("./format");
Object.defineProperty(exports, "formatErrMsg", { enumerable: true, get: function () { return format_1.formatErrMsg; } });
Object.defineProperty(exports, "formatInfoMsg", { enumerable: true, get: function () { return format_1.formatInfoMsg; } });
Object.defineProperty(exports, "formatWarnMsg", { enumerable: true, get: function () { return format_1.formatWarnMsg; } });
let lastType;
let lastMsg;
function resetOutput(type) {
    if (type === lastType) {
        lastType = undefined;
        lastMsg = '';
    }
}
exports.resetOutput = resetOutput;
function output(type, msg) {
    if (type === lastType && msg === lastMsg) {
        return;
    }
    lastMsg = msg;
    lastType = type;
    const method = type === 'info' ? 'log' : type;
    console[method](msg);
}
exports.output = output;
