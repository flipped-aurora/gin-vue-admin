import { type CompilerError, type SourceLocation } from '@vue/compiler-core';
export declare const enum MPErrorCodes {
    X_V_ON_NO_ARGUMENT = 0,
    X_V_ON_DYNAMIC_EVENT = 1,
    X_V_BIND_NO_ARGUMENT = 2,
    X_V_BIND_DYNAMIC_ARGUMENT = 3,
    X_V_BIND_MODIFIER_PROP = 4,
    X_V_BIND_MODIFIER_ATTR = 5,
    X_V_IS_NOT_SUPPORTED = 6,
    X_NOT_SUPPORTED = 7,
    X_DYNAMIC_COMPONENT_NOT_SUPPORTED = 8
}
export interface MPCompilerError extends CompilerError {
    code: MPErrorCodes;
}
export declare function createMPCompilerError(code: MPErrorCodes, loc?: SourceLocation, additionalMessage?: string): MPCompilerError;
