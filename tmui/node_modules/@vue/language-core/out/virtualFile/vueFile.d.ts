import { FileCapabilities, FileKind, VirtualFile } from '@volar/language-core';
import { Stack } from '@volar/source-map';
import type * as ts from 'typescript/lib/tsserverlibrary';
import { VueCompilerOptions, VueLanguagePlugin } from '../types';
import { Signal } from 'computeds';
export declare class VueFile implements VirtualFile {
    fileName: string;
    initSnapshot: ts.IScriptSnapshot;
    vueCompilerOptions: VueCompilerOptions;
    plugins: ReturnType<VueLanguagePlugin>[];
    ts: typeof import('typescript/lib/tsserverlibrary');
    codegenStack: boolean;
    _snapshot: Signal<ts.IScriptSnapshot>;
    getVueSfc: () => import("@vue/compiler-sfc").SFCParseResult | undefined;
    sfc: import("../types").Sfc;
    getMappings: () => import("@volar/source-map").Mapping<import("@volar/language-core").FileRangeCapabilities>[];
    getEmbeddedFiles: () => VirtualFile[];
    capabilities: FileCapabilities;
    kind: FileKind;
    codegenStacks: Stack[];
    get embeddedFiles(): VirtualFile[];
    get mainScriptName(): string;
    get snapshot(): ts.IScriptSnapshot;
    get mappings(): import("@volar/source-map").Mapping<import("@volar/language-core").FileRangeCapabilities>[];
    constructor(fileName: string, initSnapshot: ts.IScriptSnapshot, vueCompilerOptions: VueCompilerOptions, plugins: ReturnType<VueLanguagePlugin>[], ts: typeof import('typescript/lib/tsserverlibrary'), codegenStack: boolean);
    update(newSnapshot: ts.IScriptSnapshot): void;
}
//# sourceMappingURL=vueFile.d.ts.map