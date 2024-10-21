import { NodeTypes, type RootNode, type TemplateChildNode, type TransformContext } from '@vue/compiler-core';
import type { COMPONENT_BIND_LINK, COMPONENT_ON_LINK } from '../../mp/constants';
export declare function createTransformComponentLink(name: typeof COMPONENT_BIND_LINK | typeof COMPONENT_ON_LINK, type?: NodeTypes.ATTRIBUTE | NodeTypes.DIRECTIVE): (node: RootNode | TemplateChildNode, context: TransformContext) => void;
