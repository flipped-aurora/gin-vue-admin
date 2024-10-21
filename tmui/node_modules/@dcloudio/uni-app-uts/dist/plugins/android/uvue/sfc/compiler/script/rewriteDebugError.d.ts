import type { Program } from '@babel/types';
import type MagicString from 'magic-string';
export declare function hasDebugError(content: string): boolean;
interface RewriteDebugErrorOptions {
    fileName: string;
    startLine: number;
    startOffset: number;
}
export declare function rewriteDebugError(ast: Program, s: MagicString, options: RewriteDebugErrorOptions): void;
export {};
