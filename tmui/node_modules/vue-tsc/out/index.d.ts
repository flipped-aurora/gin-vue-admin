import type * as ts from 'typescript/lib/tsserverlibrary';
import * as vue from '@vue/language-core';
export type Hook = (program: _Program) => void;
export type _Program = ts.Program & {
    __vue: ProgramContext;
};
interface ProgramContext {
    projectVersion: number;
    options: ts.CreateProgramOptions;
    languageHost: vue.TypeScriptLanguageHost;
    vueCompilerOptions: Partial<vue.VueCompilerOptions>;
    langaugeContext: vue.LanguageContext;
    languageService: ts.LanguageService;
}
export declare function createProgram(options: ts.CreateProgramOptions): _Program;
export {};
//# sourceMappingURL=index.d.ts.map