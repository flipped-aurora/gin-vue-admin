import { Mapping, Stack } from '@volar/source-map';
import type * as ts from 'typescript/lib/tsserverlibrary';
export interface FileCapabilities {
    diagnostic?: boolean;
    foldingRange?: boolean;
    documentFormatting?: boolean;
    documentSymbol?: boolean;
    codeAction?: boolean;
    inlayHint?: boolean;
}
export interface FileRangeCapabilities {
    hover?: boolean;
    references?: boolean;
    definition?: boolean;
    rename?: boolean | {
        normalize?(newName: string): string;
        apply?(newName: string): string;
    };
    completion?: boolean | {
        additional?: boolean;
        autoImportOnly?: boolean;
    };
    diagnostic?: boolean | {
        shouldReport(): boolean;
    };
    semanticTokens?: boolean;
    referencesCodeLens?: boolean;
    displayWithLink?: boolean;
}
export interface MirrorBehaviorCapabilities {
    references?: boolean;
    definition?: boolean;
    rename?: boolean;
}
export declare namespace FileCapabilities {
    const full: FileCapabilities;
}
export declare namespace FileRangeCapabilities {
    const full: FileRangeCapabilities;
}
export declare namespace MirrorBehaviorCapabilities {
    const full: MirrorBehaviorCapabilities;
}
export declare enum FileKind {
    TextFile = 0,
    TypeScriptHostFile = 1
}
export interface VirtualFile {
    fileName: string;
    snapshot: ts.IScriptSnapshot;
    kind: FileKind;
    capabilities: FileCapabilities;
    mappings: Mapping<FileRangeCapabilities>[];
    codegenStacks: Stack[];
    mirrorBehaviorMappings?: Mapping<[MirrorBehaviorCapabilities, MirrorBehaviorCapabilities]>[];
    embeddedFiles: VirtualFile[];
}
export interface Language<T extends VirtualFile = VirtualFile> {
    resolveHost?(host: TypeScriptLanguageHost): TypeScriptLanguageHost;
    createVirtualFile(fileName: string, snapshot: ts.IScriptSnapshot, languageId: string | undefined): T | undefined;
    updateVirtualFile(virtualFile: T, snapshot: ts.IScriptSnapshot): void;
    deleteVirtualFile?(virtualFile: T): void;
}
interface LanguageHost {
    workspacePath: string;
    rootPath: string;
    getProjectVersion(): string;
    getScriptFileNames(): string[];
    getScriptSnapshot(fileName: string): ts.IScriptSnapshot | undefined;
    getLanguageId?(fileName: string): string | undefined;
}
export interface TypeScriptLanguageHost extends LanguageHost, Pick<ts.LanguageServiceHost, 'getCancellationToken' | 'getLocalizedDiagnosticMessages' | 'getCompilationSettings' | 'getProjectReferences'> {
    resolveModuleName?(path: string, impliedNodeFormat?: ts.ResolutionMode): string;
}
export {};
//# sourceMappingURL=types.d.ts.map