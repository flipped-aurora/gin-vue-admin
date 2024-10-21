import type { Language } from '@volar/language-core';
import { VueFile } from './virtualFile/vueFile';
import { VueCompilerOptions } from './types';
import type * as ts from 'typescript/lib/tsserverlibrary';
export declare function createVueLanguage(ts: typeof import('typescript/lib/tsserverlibrary'), compilerOptions?: ts.CompilerOptions, _vueCompilerOptions?: Partial<VueCompilerOptions>, codegenStack?: boolean): Language<VueFile>;
/**
 * @deprecated planed to remove in 2.0, please use createVueLanguage instead of
 */
export declare function createLanguages(ts: typeof import('typescript/lib/tsserverlibrary'), compilerOptions?: ts.CompilerOptions, vueCompilerOptions?: Partial<VueCompilerOptions>, codegenStack?: boolean): Language[];
//# sourceMappingURL=languageModule.d.ts.map