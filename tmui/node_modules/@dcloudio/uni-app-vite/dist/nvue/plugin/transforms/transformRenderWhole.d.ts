import { type NodeTransform, type RootNode } from '@vue/compiler-core';
export declare const transformRenderWhole: NodeTransform;
/**
 * 仅当根节点只有一个，标签在白名单，且开发者未主动配置的情况下，才补充
 * @param node
 */
export declare function addRenderWhole(node: RootNode): void;
