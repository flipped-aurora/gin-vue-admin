import type { Program } from '@babel/types';
import type MagicString from 'magic-string';
export declare function hasConsole(content: string): boolean;
export declare function rewriteConsole(ast: Program, s: MagicString, { fileName, startLine, startOffset, }: {
    fileName: string;
    startLine: number;
    startOffset: number;
}): void;
