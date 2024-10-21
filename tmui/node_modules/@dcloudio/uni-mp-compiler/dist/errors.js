"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createMPCompilerError = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const MPErrorMessages = {
    [0 /* MPErrorCodes.X_V_ON_NO_ARGUMENT */]: 'v-on="" is not supported',
    [1 /* MPErrorCodes.X_V_ON_DYNAMIC_EVENT */]: 'v-on:[event]="" is not supported.',
    [2 /* MPErrorCodes.X_V_BIND_NO_ARGUMENT */]: 'v-bind="" is not supported.',
    [3 /* MPErrorCodes.X_V_BIND_DYNAMIC_ARGUMENT */]: 'v-bind:[name]="" is not supported.',
    [4 /* MPErrorCodes.X_V_BIND_MODIFIER_PROP */]: 'v-bind .prop is not supported',
    [5 /* MPErrorCodes.X_V_BIND_MODIFIER_ATTR */]: 'v-bind .attr is not supported',
    [8 /* MPErrorCodes.X_DYNAMIC_COMPONENT_NOT_SUPPORTED */]: '<component is=""/> is not supported',
    [7 /* MPErrorCodes.X_NOT_SUPPORTED */]: 'not supported: ',
    [6 /* MPErrorCodes.X_V_IS_NOT_SUPPORTED */]: 'v-is not supported',
};
function createMPCompilerError(code, loc, additionalMessage) {
    return (0, compiler_core_1.createCompilerError)(code, loc, MPErrorMessages, additionalMessage);
}
exports.createMPCompilerError = createMPCompilerError;
