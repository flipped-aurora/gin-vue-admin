import { type ExpressionNode } from '@vue/compiler-core';
import type { TemplateCompilerOptions } from './options';
import type { TransformContext } from './transform';
import type { CompilerError } from './errors';
export declare const __DEV__ = true;
export declare const __BROWSER__ = false;
export declare const __COMPAT__ = false;
export declare function isCompatEnabled(...args: any[]): boolean;
export declare function genRenderFunctionDecl({ className, genDefaultAs, inline, }: TemplateCompilerOptions & {
    genDefaultAs?: string;
}): string;
export declare function rewriteObjectExpression(exp: ExpressionNode, context: TransformContext): import("@vue/compiler-core").SimpleExpressionNode | undefined;
export declare function onCompilerError(error: CompilerError): void;
export declare function parseSource(fileName: string, rootDir: string): string;
export declare function addEasyComponentAutoImports(easyComponentAutoImports: Record<string, [string, string]>, rootDir: string, tagName: string, fileName: string): void;
