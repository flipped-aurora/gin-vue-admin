import type { NodeTransform, TransformContext } from '../transform';
import { type ExpressionNode, type SlotOutletNode } from '@vue/compiler-core';
import type { PropsExpression } from '@vue/compiler-core';
export declare const transformSlotOutlet: NodeTransform;
interface SlotOutletProcessResult {
    slotName: string | ExpressionNode;
    slotProps: PropsExpression | undefined;
}
export declare function processSlotOutlet(node: SlotOutletNode, context: TransformContext): SlotOutletProcessResult;
export {};
