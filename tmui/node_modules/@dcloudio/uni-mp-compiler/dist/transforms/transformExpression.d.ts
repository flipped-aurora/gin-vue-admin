import { type ExpressionNode, type SimpleExpressionNode } from '@vue/compiler-core';
import type { NodeTransform, TransformContext } from '../transform';
export declare const transformExpression: NodeTransform;
export declare function processExpression(node: SimpleExpressionNode, context: TransformContext, asParams?: boolean, asRawStatements?: boolean, localVars?: Record<string, number>): ExpressionNode;
export declare function isBuiltInIdentifier(id: string | ExpressionNode): boolean;
