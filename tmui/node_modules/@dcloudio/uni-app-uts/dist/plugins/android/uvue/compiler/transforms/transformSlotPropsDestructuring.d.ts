import { type CompoundExpressionNode } from '@vue/compiler-core';
import type { CodegenContext } from '../codegen';
export declare const SLOT_PROPS_NAME = "slotProps";
export declare function isDestructuringSlotProps(isSlot: boolean, params: CompoundExpressionNode): boolean;
export declare function createDestructuringSlotProps(params: CompoundExpressionNode, context: CodegenContext): void;
