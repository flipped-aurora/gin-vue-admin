import { type RootNode, type TemplateChildNode } from '@vue/compiler-core';
/**
 * 将scroll-view、list-view内的<view slot="refresher">转为vue支持的用法，此transform需要再较早时机执行，暂时放在transformTag前。此时node.tag还没有加上v-uni-前缀
 * @param node
 * @param context
 * @returns
 */
export declare function transformRefresherSlot(node: RootNode | TemplateChildNode): void;
