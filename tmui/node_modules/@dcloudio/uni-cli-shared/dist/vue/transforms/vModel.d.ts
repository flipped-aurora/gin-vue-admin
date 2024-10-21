import type { DirectiveTransform, ElementNode, TransformContext } from '@vue/compiler-core';
export declare function defaultMatch(node: ElementNode, context: TransformContext): boolean;
interface CreateTransformModelOptions {
    match: typeof defaultMatch;
}
/**
 * 百度、快手小程序的自定义组件，不支持动态事件绑定，故 v-model 也需要调整
 * @param baseTransformModel
 * @returns
 */
export declare function createTransformModel(baseTransformModel: DirectiveTransform, { match }?: CreateTransformModelOptions): DirectiveTransform;
export {};
