import { type DirectiveNode, type DirectiveTransform, type ElementNode, type ExpressionNode, type SimpleExpressionNode } from '@vue/compiler-core';
import type { TransformContext } from '../transform';
export interface VOnDirectiveNode extends DirectiveNode {
    arg: ExpressionNode;
    exp: SimpleExpressionNode | undefined;
}
export declare const transformOn: DirectiveTransform;
export declare function wrapperVOn(value: ExpressionNode, node: ElementNode, context: TransformContext): ExpressionNode;
