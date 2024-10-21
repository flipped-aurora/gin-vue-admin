import { type ImportDeclaration, type Program } from '@babel/types';
import type { PluginContext } from 'rollup';
type BindingComponents = Record<string, {
    tag: string;
    type: 'unknown' | 'setup' | 'self';
}>;
interface MainDescriptor {
    imports: ImportDeclaration[];
    script: string;
    template: string;
}
export declare function parseMainDescriptor(filename: string, ast: Program, resolve: ParseDescriptor['resolve']): Promise<MainDescriptor>;
export declare function updateMiniProgramComponentsByScriptFilename(scriptFilename: string, inputDir: string, normalizeComponentName: (name: string) => string): void;
export declare function updateMiniProgramComponentsByTemplateFilename(templateFilename: string, inputDir: string, normalizeComponentName: (name: string) => string): void;
export declare function updateMiniProgramGlobalComponents(filename: string, ast: Program, { inputDir, resolve, normalizeComponentName, }: {
    inputDir: string;
    resolve: ParseDescriptor['resolve'];
    normalizeComponentName: (name: string) => string;
}): Promise<{
    imports: ImportDeclaration[];
}>;
export declare function updateMiniProgramComponentsByMainFilename(mainFilename: string, inputDir: string, normalizeComponentName: (name: string) => string): void;
export interface TemplateDescriptor {
    bindingComponents: BindingComponents;
    imports: ImportDeclaration[];
}
/**
 * 解析 template
 * @param filename
 * @param code
 * @param ast
 * @param options
 * @returns
 */
export declare function parseTemplateDescriptor(filename: string, ast: Program, options: ParseDescriptor): Promise<TemplateDescriptor>;
interface ParseDescriptor {
    resolve: PluginContext['resolve'];
    isExternal: boolean;
}
export interface ScriptDescriptor extends TemplateDescriptor {
    setupBindingComponents: BindingComponents;
}
/**
 * 解析 script
 * @param filename
 * @param code
 * @param ast
 * @param options
 * @returns
 */
export declare function parseScriptDescriptor(filename: string, ast: Program, options: ParseDescriptor): Promise<ScriptDescriptor>;
/**
 * static import => dynamic import
 * @param code
 * @param imports
 * @param dynamicImport
 * @returns
 */
export declare function transformDynamicImports(code: string, imports: ImportDeclaration[], { id, sourceMap, dynamicImport, }: {
    id?: string;
    sourceMap?: boolean;
    dynamicImport: (name: string, source: string) => string;
}): Promise<{
    code: string;
    map: null;
}>;
export {};
