import { type DirectiveTransform, type ElementNode, type ExpressionNode, type TransformContext } from '@vue/compiler-core';
export declare function defaultMatch(name: string, node: ElementNode, context: TransformContext): boolean;
interface CreateTransformOnOptions {
    match: typeof defaultMatch;
}
/**
 * 百度、快手小程序的自定义组件，不支持动态事件绑定，故转换为静态事件 + dataset
 * @param baseTransformOn
 * @returns
 */
export declare function createTransformOn(baseTransformOn: DirectiveTransform, { match }?: CreateTransformOnOptions): DirectiveTransform;
export declare function createCustomEventExpr(): import("@vue/compiler-core").SimpleExpressionNode;
export declare function addEventOpts(event: string, value: ExpressionNode, node: ElementNode, context: TransformContext): void;
export declare const ATTR_DATASET_EVENT_OPTS = "data-e-o";
export declare const STRINGIFY_JSON: unique symbol;
export {};
