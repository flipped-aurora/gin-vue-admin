import type { SFCParseResult } from '@vue/compiler-sfc';
import type * as ts from 'typescript/lib/tsserverlibrary';
import { Sfc, VueLanguagePlugin } from '../types';
export declare function computedSfc(ts: typeof import('typescript/lib/tsserverlibrary'), plugins: ReturnType<VueLanguagePlugin>[], fileName: string, snapshot: () => ts.IScriptSnapshot, parsed: () => SFCParseResult | undefined): Sfc;
//# sourceMappingURL=computedSfc.d.ts.map