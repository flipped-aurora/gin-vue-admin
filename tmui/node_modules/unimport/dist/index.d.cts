export { v as vueTemplateAddon } from './shared/unimport.a3a6779e.cjs';
import { U as UnimportOptions, a as Unimport, I as Import, b as InstallGlobalOptions, S as ScanDirExportsOptions, P as Preset, B as BuiltinPresetName, c as InlinePreset, T as TypeDeclarationOptions, M as MagicStringResult } from './shared/unimport.c45f29b8.cjs';
export { q as Addon, m as AddonVueDirectivesOptions, A as AddonsOptions, D as DetectImportResult, g as ImportCommon, r as ImportInjectionResult, f as ImportName, o as InjectImportsOptions, k as InjectionUsageRecord, e as ModuleId, i as PackagePreset, n as PathFromResolver, h as PresetImport, p as Thenable, j as UnimportContext, l as UnimportMeta, d as builtinPresets } from './shared/unimport.c45f29b8.cjs';
import { StripLiteralOptions } from 'strip-literal';
import MagicString from 'magic-string';
import 'mlly';

const version = "3.13.1";

declare function createUnimport(opts: Partial<UnimportOptions>): Unimport;

declare function installGlobalAutoImports(imports: Import[] | Unimport, options?: InstallGlobalOptions): Promise<any>;

declare function scanFilesFromDir(dir: string | string[], options?: ScanDirExportsOptions): Promise<string[]>;
declare function scanDirExports(dir: string | string[], options?: ScanDirExportsOptions): Promise<Import[]>;
declare function dedupeDtsExports(exports: Import[]): Import[];
declare function scanExports(filepath: string, includeTypes: boolean, seen?: Set<string>): Promise<Import[]>;

declare function resolvePreset(preset: Preset): Promise<Import[]>;
declare function resolveBuiltinPresets(presets: (BuiltinPresetName | Preset)[]): Promise<Import[]>;

declare const excludeRE: RegExp[];
declare const importAsRE: RegExp;
declare const separatorRE: RegExp;
/**
 *                                                                             |       |
 *                    destructing   case&ternary    non-call     inheritance   |  id   |
 *                         ↓             ↓             ↓             ↓         |       |
 */
declare const matchRE: RegExp;
declare function stripCommentsAndStrings(code: string, options?: StripLiteralOptions): string;

declare function defineUnimportPreset(preset: InlinePreset): InlinePreset;
declare function stringifyImports(imports: Import[], isCJS?: boolean): string;
declare function dedupeImports(imports: Import[], warn: (msg: string) => void): Import[];
declare function toExports(imports: Import[], fileDir?: string, includeType?: boolean): string;
declare function stripFileExtension(path: string): string;
declare function toTypeDeclarationItems(imports: Import[], options?: TypeDeclarationOptions): string[];
declare function toTypeDeclarationFile(imports: Import[], options?: TypeDeclarationOptions): string;
declare function toTypeReExports(imports: Import[], options?: TypeDeclarationOptions): string;
declare function getString(code: string | MagicString): string;
declare function getMagicString(code: string | MagicString): MagicString;
declare function addImportToCode(code: string | MagicString, imports: Import[], isCJS?: boolean, mergeExisting?: boolean, injectAtLast?: boolean, firstOccurrence?: number, onResolved?: (imports: Import[]) => void | Import[], onStringified?: (str: string, imports: Import[]) => void | string): MagicStringResult;
declare function normalizeImports(imports: Import[]): Import[];
declare function resolveIdAbsolute(id: string, parentId?: string): Promise<string>;
/**
 * @deprecated renamed to `stringifyImports`
 */
declare const toImports: typeof stringifyImports;

export { BuiltinPresetName, Import, InlinePreset, InstallGlobalOptions, MagicStringResult, Preset, ScanDirExportsOptions, TypeDeclarationOptions, Unimport, UnimportOptions, addImportToCode, createUnimport, dedupeDtsExports, dedupeImports, defineUnimportPreset, excludeRE, getMagicString, getString, importAsRE, installGlobalAutoImports, matchRE, normalizeImports, resolveBuiltinPresets, resolveIdAbsolute, resolvePreset, scanDirExports, scanExports, scanFilesFromDir, separatorRE, stringifyImports, stripCommentsAndStrings, stripFileExtension, toExports, toImports, toTypeDeclarationFile, toTypeDeclarationItems, toTypeReExports, version };
