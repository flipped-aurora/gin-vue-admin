import { type DirectiveNode, type ElementNode, type ExpressionNode, type ForNode, type SimpleExpressionNode } from '@vue/compiler-core';
import { type TransformContext } from '../transform';
export declare const transformFor: import("../transform").NodeTransform;
export declare function processFor(node: ElementNode, dir: DirectiveNode, context: TransformContext, processCodegen?: (forNode: ForNode) => (() => void) | undefined): (() => void) | undefined;
export interface ForParseResult {
    source: ExpressionNode;
    value: ExpressionNode | undefined;
    key: ExpressionNode | undefined;
    index: ExpressionNode | undefined;
    finalized: boolean;
}
export declare function parseForExpression(input: SimpleExpressionNode, context: TransformContext): ForParseResult | undefined;
export declare function createForLoopParams({ value, key, index }: ForParseResult, memoArgs: ExpressionNode): ExpressionNode[];
