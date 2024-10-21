import { type ElementNode, type Property, type TemplateLiteral } from '@vue/compiler-core';
import { type NodeTransform, type TransformContext } from '../transform';
export interface DirectiveTransformResult {
    props: Property[];
    needRuntime?: boolean | symbol;
    ssrTagParts?: TemplateLiteral['elements'];
}
export declare const transformElement: NodeTransform;
export declare function processProps(node: ElementNode, context: TransformContext, props?: ElementNode['props']): void;
