import type * as ts from 'typescript/lib/tsserverlibrary';
import type { VueCompilerOptions } from '../types';
export type ParsedCommandLine = ts.ParsedCommandLine & {
    vueOptions: Partial<VueCompilerOptions>;
};
export declare function createParsedCommandLineByJson(ts: typeof import('typescript/lib/tsserverlibrary'), parseConfigHost: ts.ParseConfigHost, rootDir: string, json: any, configFileName?: string): ParsedCommandLine;
export declare function createParsedCommandLine(ts: typeof import('typescript/lib/tsserverlibrary'), parseConfigHost: ts.ParseConfigHost, tsConfigPath: string): ParsedCommandLine;
export declare function resolveVueCompilerOptions(vueOptions: Partial<VueCompilerOptions>): VueCompilerOptions;
//# sourceMappingURL=ts.d.ts.map