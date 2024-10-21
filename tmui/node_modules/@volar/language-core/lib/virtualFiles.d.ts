import { SourceMap } from '@volar/source-map';
import type * as ts from 'typescript/lib/tsserverlibrary';
import { MirrorMap } from './sourceMaps';
import type { FileRangeCapabilities, Language, VirtualFile } from './types';
export type VirtualFiles = ReturnType<typeof createVirtualFiles>;
export interface Source {
    fileName: string;
    languageId: string | undefined;
    snapshot: ts.IScriptSnapshot;
    root: VirtualFile;
    language: Language;
}
export declare function createVirtualFiles(languages: Language[]): {
    allSources(): Source[];
    updateSource(fileName: string, snapshot: ts.IScriptSnapshot, languageId: string | undefined): VirtualFile | undefined;
    deleteSource(fileName: string): void;
    getSource(fileName: string): Source | undefined;
    hasSource: (fileName: string) => boolean;
    getMirrorMap: (file: VirtualFile) => MirrorMap | undefined;
    getMaps: (virtualFile: VirtualFile) => Map<string, [ts.IScriptSnapshot, SourceMap<FileRangeCapabilities>]>;
    hasVirtualFile(fileName: string): boolean;
    getVirtualFile(fileName: string): readonly [VirtualFile, Source] | readonly [undefined, undefined];
};
export declare function updateVirtualFileMaps(virtualFile: VirtualFile, getSourceSnapshot: (source: string | undefined) => [string, ts.IScriptSnapshot] | undefined, map?: Map<string, [ts.IScriptSnapshot, SourceMap<FileRangeCapabilities>]>): Map<string, [ts.IScriptSnapshot, SourceMap<FileRangeCapabilities>]>;
export declare function forEachEmbeddedFile(file: VirtualFile, cb: (embedded: VirtualFile) => void): void;
//# sourceMappingURL=virtualFiles.d.ts.map