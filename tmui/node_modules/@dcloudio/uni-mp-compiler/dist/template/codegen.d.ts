import { type MiniProgramCompilerOptions } from '@dcloudio/uni-cli-shared';
import { type ElementNode, type RootNode, type TemplateChildNode } from '@vue/compiler-core';
import type { TemplateCodegenOptions } from '../options';
import type { TransformContext } from '../transform';
export interface TemplateCodegenContext {
    code: string;
    directive: string;
    scopeId?: string | null;
    event: MiniProgramCompilerOptions['event'];
    slot: MiniProgramCompilerOptions['slot'];
    lazyElement: MiniProgramCompilerOptions['lazyElement'];
    component: MiniProgramCompilerOptions['component'];
    isBuiltInComponent: TransformContext['isBuiltInComponent'];
    isMiniProgramComponent: TransformContext['isMiniProgramComponent'];
    push(code: string): void;
}
export declare function generate({ children }: RootNode, { slot, event, scopeId, emitFile, filename, directive, lazyElement, isBuiltInComponent, isMiniProgramComponent, component, }: TemplateCodegenOptions): void;
export declare function genNode(node: TemplateChildNode, context: TemplateCodegenContext): void;
export declare function genElementProps(node: ElementNode, virtualHost: boolean, context: TemplateCodegenContext): void;
