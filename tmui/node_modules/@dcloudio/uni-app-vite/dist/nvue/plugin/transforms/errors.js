"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createNVueCompilerError = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const NVueErrorMessages = {
    [0 /* NVueErrorCodes.X_V_SHOW */]: 'nvue: v-show is not supported',
    [1 /* NVueErrorCodes.X_V_MODEL_DYNAMIC_TYPE */]: 'nvue: v-model with :type="" is not supported',
    [2 /* NVueErrorCodes.X_V_MODEL_AND_V_BIND */]: 'nvue: v-model with v-bind is not supported',
};
function createNVueCompilerError(code, loc, additionalMessage) {
    return (0, compiler_core_1.createCompilerError)(code, loc, NVueErrorMessages, additionalMessage);
}
exports.createNVueCompilerError = createNVueCompilerError;
