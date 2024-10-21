import './runtimeHelpers';
import type { CodegenResult, TemplateCompilerOptions } from './options';
import { type DirectiveTransform, type NodeTransform } from './transform';
export type TransformPreset = [
    NodeTransform[],
    Record<string, DirectiveTransform>
];
export declare function getBaseTransformPreset(prefixIdentifiers?: boolean): TransformPreset;
export declare function compile(template: string, options: TemplateCompilerOptions): CodegenResult;
