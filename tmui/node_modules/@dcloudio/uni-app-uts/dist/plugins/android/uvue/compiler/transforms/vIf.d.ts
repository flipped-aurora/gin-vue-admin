import { type DirectiveNode, type ElementNode, type IfBranchNode, type IfNode } from '@vue/compiler-core';
import { type TransformContext } from '../transform';
export declare const transformIf: import("../transform").NodeTransform;
export declare function processIf(node: ElementNode, dir: DirectiveNode, context: TransformContext, processCodegen?: (node: IfNode, branch: IfBranchNode, isRoot: boolean) => (() => void) | undefined): (() => void) | undefined;
