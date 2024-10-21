import type * as ts from 'typescript/lib/tsserverlibrary';
import { VueCompilerOptions } from './types';
import * as CompilerDOM from '@vue/compiler-dom';
export declare function getDefaultVueLanguagePlugins(ts: typeof import('typescript/lib/tsserverlibrary'), compilerOptions: ts.CompilerOptions, vueCompilerOptions: VueCompilerOptions, codegenStack: boolean): {
    version: 1;
    name?: string | undefined;
    order?: number | undefined;
    requiredCompilerOptions?: string[] | undefined;
    parseSFC?(fileName: string, content: string): import("@vue/compiler-sfc").SFCParseResult | undefined;
    updateSFC?(oldResult: import("@vue/compiler-sfc").SFCParseResult, textChange: {
        start: number;
        end: number;
        newText: string;
    }): import("@vue/compiler-sfc").SFCParseResult | undefined;
    resolveTemplateCompilerOptions?(options: CompilerDOM.CompilerOptions): CompilerDOM.CompilerOptions;
    compileSFCTemplate?(lang: string, template: string, options: CompilerDOM.CompilerOptions): CompilerDOM.CodegenResult | undefined;
    updateSFCTemplate?(oldResult: CompilerDOM.CodegenResult, textChange: {
        start: number;
        end: number;
        newText: string;
    }): CompilerDOM.CodegenResult | undefined;
    getEmbeddedFileNames?(fileName: string, sfc: import("./types").Sfc): string[];
    resolveEmbeddedFile?(fileName: string, sfc: import("./types").Sfc, embeddedFile: import("./virtualFile/embeddedFile").VueEmbeddedFile): void;
}[];
//# sourceMappingURL=plugins.d.ts.map