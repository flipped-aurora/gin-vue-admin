import type { LogErrorOptions, LogOptions } from 'vite';
export interface Formatter<T extends LogOptions = LogOptions> {
    test: (msg: string, options?: T) => boolean;
    format: (msg: string, options?: T) => string;
}
export declare function formatErrMsg(msg: string, options?: LogErrorOptions): string;
export declare const removeNVueInfoFormatter: Formatter;
export declare function formatInfoMsg(msg: string, options?: LogOptions & {
    nvue?: boolean;
}): string;
export declare function formatWarnMsg(msg: string, options?: LogOptions): string;
