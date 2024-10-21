import type { Container } from 'postcss';
import type { Document as Document_2 } from 'postcss';
import type { Plugin as Plugin_2 } from 'postcss';
import postcss from 'postcss';
import type { Root } from 'postcss';

export declare function expand(options: NormalizeOptions): Plugin_2;

export declare function normalize(opts?: NormalizeOptions): Plugin_2;

declare interface NormalizeOptions {
    logLevel?: 'NOTE' | 'WARNING' | 'ERROR';
    type?: 'nvue' | 'uvue';
    platform?: typeof process.env.UNI_UTS_PLATFORM;
}

export declare function objectifier(node: Root | Document_2 | Container | null, { trim }?: {
    trim: boolean;
}): Record<string, unknown>;

export declare function parse(input: string, options?: ParseOptions): Promise<{
    code: string;
    messages: postcss.Message[];
}>;

declare interface ParseOptions extends NormalizeOptions {
    filename?: string;
    map?: boolean;
    mapOf?: string;
    ts?: boolean;
    chunk?: number;
    noCode?: boolean;
    trim?: boolean;
}

export { }
