import { createVirtualFiles } from './virtualFiles';
import { Language, TypeScriptLanguageHost } from './types';
export interface LanguageContext {
    rawHost: TypeScriptLanguageHost;
    host: TypeScriptLanguageHost;
    virtualFiles: ReturnType<typeof createVirtualFiles>;
}
export declare function createLanguageContext(rawHost: TypeScriptLanguageHost, languages: Language<any>[]): LanguageContext;
//# sourceMappingURL=languageContext.d.ts.map