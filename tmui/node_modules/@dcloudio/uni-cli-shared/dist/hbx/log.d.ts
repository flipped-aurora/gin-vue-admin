import type { LogErrorOptions } from 'vite';
import type { Formatter } from '../logs/format';
export declare function formatAtFilename(filename: string, line?: number, column?: number): string;
export declare const h5ServeFormatter: Formatter;
export declare const removeInfoFormatter: Formatter;
export declare const removeWarnFormatter: Formatter;
export declare const errorFormatter: Formatter<LogErrorOptions>;
