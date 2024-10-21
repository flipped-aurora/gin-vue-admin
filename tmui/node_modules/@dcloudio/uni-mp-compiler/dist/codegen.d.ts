import { type CodegenResult, type CompoundExpressionNode, type InterpolationNode, type SimpleExpressionNode, type TextNode } from '@vue/compiler-core';
import type { Expression } from '@babel/types';
import { type GeneratorOptions } from '@babel/generator';
import type { CodegenOptions, CodegenRootNode } from './options';
export declare function generate(ast: CodegenRootNode, options: CodegenOptions): Omit<CodegenResult, 'ast'>;
type CodegenNode = SimpleExpressionNode | CompoundExpressionNode | InterpolationNode | TextNode;
interface GenNodeContext {
    code: string;
    helper(key: symbol): string;
    push(code: string, node?: CodegenNode): void;
}
export declare function genBabelExpr(expr: Expression, opts?: GeneratorOptions): string;
export declare function genExpr(node: CodegenNode | symbol | string, context?: GenNodeContext): string;
export {};
