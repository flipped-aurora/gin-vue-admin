import type { DirectiveNode, ElementNode, NodeTransform } from '@vue/compiler-core';
export declare function createTransformEvent(options: Record<string, string | ((node: ElementNode, dir: DirectiveNode) => string)>): NodeTransform;
