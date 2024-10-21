import type { SFCParseResult } from '@vue/compiler-sfc';
import type * as ts from 'typescript/lib/tsserverlibrary';
import { VueLanguagePlugin } from '../types';
export declare function computedVueSfc(plugins: ReturnType<VueLanguagePlugin>[], fileName: string, snapshot: () => ts.IScriptSnapshot): () => SFCParseResult | undefined;
//# sourceMappingURL=computedVueSfc.d.ts.map