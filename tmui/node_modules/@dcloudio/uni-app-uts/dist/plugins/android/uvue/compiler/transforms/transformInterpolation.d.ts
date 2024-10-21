import type { NodeTransform } from '../transform';
import { type InterpolationNode, type TemplateChildNode, type TextNode } from '@vue/compiler-core';
export declare function isText(node: TemplateChildNode): node is TextNode | InterpolationNode;
export declare const transformInterpolation: NodeTransform;
