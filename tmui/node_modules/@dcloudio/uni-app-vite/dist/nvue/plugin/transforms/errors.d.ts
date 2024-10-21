import { type CompilerError, type SourceLocation } from '@vue/compiler-core';
export declare const enum NVueErrorCodes {
    X_V_SHOW = 0,
    X_V_MODEL_DYNAMIC_TYPE = 1,
    X_V_MODEL_AND_V_BIND = 2
}
interface NVueCompilerError extends CompilerError {
    code: NVueErrorCodes;
}
export declare function createNVueCompilerError(code: NVueErrorCodes, loc?: SourceLocation, additionalMessage?: string): NVueCompilerError;
export {};
