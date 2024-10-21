import type { DecodedSourceMap, RawSourceMap } from '@ampproject/remapping';
import type { Position, SourceLocation } from '@vue/compiler-core';
import { LinesAndColumns } from 'lines-and-columns';
export declare function slash(p: string): string;
export declare const bareImportRE: RegExp;
export declare const deepImportRE: RegExp;
export declare const isWindows: boolean;
export declare function normalizePath(id: string): string;
export declare const queryRE: RegExp;
export declare const hashRE: RegExp;
export declare const cleanUrl: (url: string) => string;
export declare const externalRE: RegExp;
export declare const isExternalUrl: (url: string) => boolean;
export declare const dataUrlRE: RegExp;
export declare const isDataUrl: (url: string) => boolean;
export declare const multilineCommentsRE: RegExp;
export declare function asyncReplace(input: string, re: RegExp, replacer: (match: RegExpExecArray) => string | Promise<string>): Promise<string>;
export declare function isObject(value: unknown): value is Record<string, any>;
export declare function pad(source: string, n?: number): string;
export declare function offsetToStartAndEnd(source: string, startOffset: number, endOffset: number): SourceLocation;
export declare function offsetToLineColumn(source: string, offset: number): {
    line: number;
    column: number;
};
export declare function offsetToLineColumnByLines(lines: LinesAndColumns, offset: number): Position;
export declare function posToNumber(source: string, pos: number | {
    line: number;
    column: number;
}): number;
export declare function locToStartAndEnd(source: string, loc: {
    start: {
        line: number;
        column: number;
    };
    end: {
        line: number;
        column: number;
    };
}): {
    start: number;
    end: number | undefined;
};
export declare function generateCodeFrame(source: string, start?: number | {
    line: number;
    column: number;
}, end?: number): string;
interface ImageCandidate {
    url: string;
    descriptor: string;
}
export declare function processSrcSet(srcs: string, replacer: (arg: ImageCandidate) => Promise<string>): Promise<string>;
export declare function combineSourcemaps(filename: string, sourcemapList: Array<DecodedSourceMap | RawSourceMap>, excludeContent?: boolean): RawSourceMap;
export {};
