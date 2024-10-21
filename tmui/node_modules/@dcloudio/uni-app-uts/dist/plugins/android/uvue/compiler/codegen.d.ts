import { SourceMapGenerator } from 'source-map-js';
import { type JSChildNode, type RootNode, type SSRCodegenNode, type TemplateChildNode } from '@vue/compiler-core';
import { type ParserPlugin } from '@babel/parser';
import type { CodegenOptions, CodegenResult } from './options';
type CodegenNode = TemplateChildNode | JSChildNode | SSRCodegenNode;
export interface CodegenContext extends Required<Omit<CodegenOptions, 'sourceMapGeneratedLine' | 'className' | 'originalLineOffset' | 'generatedLineOffset' | 'inMap'>> {
    source: string;
    code: string;
    easyComponentAutoImports: Record<string, [string, string]>;
    line: number;
    column: number;
    offset: number;
    indentLevel: number;
    map?: SourceMapGenerator;
    expressionPlugins: ParserPlugin[];
    helper(key: symbol): string;
    push(code: string, node?: CodegenNode): void;
    indent(): void;
    deindent(withoutNewLine?: boolean): void;
    newline(): void;
}
export declare function generate(ast: RootNode, options?: CodegenOptions & {
    genDefaultAs?: string;
}): CodegenResult;
export {};
