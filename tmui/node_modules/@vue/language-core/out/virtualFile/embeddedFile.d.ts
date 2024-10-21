import { FileCapabilities, FileKind, FileRangeCapabilities, MirrorBehaviorCapabilities } from '@volar/language-core';
import { Mapping, Segment, StackNode } from '@volar/source-map';
export declare class VueEmbeddedFile {
    fileName: string;
    content: Segment<FileRangeCapabilities>[];
    contentStacks: StackNode[];
    parentFileName?: string;
    kind: FileKind;
    capabilities: FileCapabilities;
    mirrorBehaviorMappings: Mapping<[MirrorBehaviorCapabilities, MirrorBehaviorCapabilities]>[];
    constructor(fileName: string, content: Segment<FileRangeCapabilities>[], contentStacks: StackNode[]);
}
//# sourceMappingURL=embeddedFile.d.ts.map