import type { NodeTransform, TransformContext } from '../transform';
import { type ExpressionNode, type SimpleExpressionNode } from '@vue/compiler-core';
export declare const transformExpression: NodeTransform;
export declare function processExpression(node: SimpleExpressionNode, context: TransformContext, asParams?: boolean, asRawStatements?: boolean, localVars?: Record<string, number>): ExpressionNode;
export declare function stringifyExpression(exp: ExpressionNode | string | (ExpressionNode | string)[]): string;
