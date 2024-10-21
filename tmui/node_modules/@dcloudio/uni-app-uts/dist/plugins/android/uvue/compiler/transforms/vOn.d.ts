import type { DirectiveTransform } from '../transform';
import { type DirectiveNode, type ExpressionNode, type SimpleExpressionNode } from '@vue/compiler-core';
export interface VOnDirectiveNode extends DirectiveNode {
    arg: ExpressionNode;
    exp: SimpleExpressionNode | undefined;
}
export declare const transformOn: DirectiveTransform;
