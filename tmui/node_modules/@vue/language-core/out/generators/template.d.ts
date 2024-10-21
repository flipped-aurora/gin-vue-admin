import { FileRangeCapabilities } from '@volar/language-core';
import { Segment } from '@volar/source-map';
import * as CompilerDOM from '@vue/compiler-dom';
import * as muggle from 'muggle-string';
import type * as ts from 'typescript/lib/tsserverlibrary';
import { Sfc, VueCompilerOptions } from '../types';
type Code = Segment<FileRangeCapabilities>;
export declare function generate(ts: typeof import('typescript/lib/tsserverlibrary'), compilerOptions: ts.CompilerOptions, vueCompilerOptions: VueCompilerOptions, template: NonNullable<Sfc['template']>, shouldGenerateScopedClasses: boolean, stylesScopedClasses: Set<string>, hasScriptSetupSlots: boolean, slotsAssignName: string | undefined, propsAssignName: string | undefined, codegenStack: boolean): {
    codes: Code[];
    codeStacks: muggle.StackNode[];
    formatCodes: Code[];
    formatCodeStacks: muggle.StackNode[];
    cssCodes: Code[];
    cssCodeStacks: muggle.StackNode[];
    tagNames: Record<string, number[]>;
    accessedGlobalVariables: Set<string>;
    hasSlot: boolean;
};
export declare function walkElementNodes(node: CompilerDOM.RootNode | CompilerDOM.TemplateChildNode, cb: (node: CompilerDOM.ElementNode) => void): void;
export {};
//# sourceMappingURL=template.d.ts.map