import type { LanguageContext } from '@volar/language-core';
import type * as ts from 'typescript/lib/tsserverlibrary';
export declare function createLanguageServiceHost(ctx: LanguageContext, ts: typeof import('typescript/lib/tsserverlibrary'), sys: ts.System & {
    version?: number;
}): ts.LanguageServiceHost;
//# sourceMappingURL=languageServiceHost.d.ts.map