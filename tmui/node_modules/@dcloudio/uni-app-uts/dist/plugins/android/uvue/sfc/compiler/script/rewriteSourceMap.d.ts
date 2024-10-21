import type { Program } from '@babel/types';
import type MagicString from 'magic-string';
export declare function rewriteSourceMap(ast: Program, s: MagicString, { fileName, startLine, startOffset, }: {
    fileName: string;
    startLine: number;
    startOffset: number;
}): void;
