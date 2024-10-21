import { FileRangeCapabilities } from '@volar/language-core';
import { Mapping } from '@volar/source-map';
import type * as ts from 'typescript/lib/tsserverlibrary';
import { Sfc } from '../types';
export declare function computedMappings(snapshot: () => ts.IScriptSnapshot, sfc: Sfc): () => Mapping<FileRangeCapabilities>[];
//# sourceMappingURL=computedMappings.d.ts.map