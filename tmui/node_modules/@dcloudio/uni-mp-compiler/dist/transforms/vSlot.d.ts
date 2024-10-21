import { type BinaryExpression } from '@babel/types';
import { type ComponentNode, type CompoundExpressionNode, type DirectiveNode } from '@vue/compiler-core';
import type { CodegenScope, CodegenVForScope } from '../options';
import { type NodeTransform, type TransformContext } from '../transform';
export declare const transformSlot: NodeTransform;
export declare function rewriteVSlot(dir: DirectiveNode, context: TransformContext): void;
export declare function findSlotName(slotDir: DirectiveNode): string | CompoundExpressionNode;
export declare function findCurrentPath(id: string, scope: CodegenScope): BinaryExpression | import("@babel/types").StringLiteral;
/**
 * 目前无用
 * @param vForScope
 * @param parentScope
 * @param context
 */
export declare function rewriteScopedSlotVForScope(vForScope: CodegenVForScope, parentScope: CodegenScope, context: TransformContext): void;
export declare function createVSlotCallExpression(slotComponent: ComponentNode, vForScope: CodegenVForScope, context: TransformContext): import("@babel/types").CallExpression;
