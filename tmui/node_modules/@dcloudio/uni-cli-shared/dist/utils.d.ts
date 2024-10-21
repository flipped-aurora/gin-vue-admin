export { default as hash } from 'hash-sum';
import { type ElementNode, type RootNode, type TemplateChildNode } from '@vue/compiler-core';
import type { ParserPlugin } from '@babel/parser';
export { camelize, capitalize, isArray } from '@vue/shared';
export declare const isWindows: boolean;
export declare function normalizePath(id: string): string;
export declare function checkElementNodeTag(node: RootNode | TemplateChildNode | null | undefined, tag: string): node is ElementNode;
/**
 * 根据 path 返回合法 js 变量
 * @param str pages.json.page.path
 * @returns
 */
export declare function normalizeIdentifier(str: string): string;
export declare function normalizePagePath(pagePath: string, platform: UniApp.PLATFORM): string | undefined;
export declare function removeExt(str: string): string;
export declare function normalizeNodeModules(str: string): string;
export declare function normalizeMiniProgramFilename(filename: string, inputDir?: string): string;
export declare function normalizeParsePlugins(importer: string, babelParserPlugins?: ParserPlugin[]): (import("@babel/parser").ParserPluginWithOptions | ("jsx" | "typescript" | "asyncDoExpressions" | "asyncGenerators" | "bigInt" | "classPrivateMethods" | "classPrivateProperties" | "classProperties" | "classStaticBlock" | "decimal" | "decorators-legacy" | "deferredImportEvaluation" | "decoratorAutoAccessors" | "destructuringPrivate" | "doExpressions" | "dynamicImport" | "explicitResourceManagement" | "exportDefaultFrom" | "exportNamespaceFrom" | "flow" | "flowComments" | "functionBind" | "functionSent" | "importMeta" | "logicalAssignment" | "importAssertions" | "importAttributes" | "importReflection" | "moduleBlocks" | "moduleStringNames" | "nullishCoalescingOperator" | "numericSeparator" | "objectRestSpread" | "optionalCatchBinding" | "optionalChaining" | "partialApplication" | "placeholders" | "privateIn" | "regexpUnicodeSets" | "sourcePhaseImports" | "throwExpressions" | "topLevelAwait" | "v8intrinsic" | "decorators" | "estree" | "moduleAttributes" | "optionalChainingAssign" | "pipelineOperator" | "recordAndTuple"))[];
export declare function pathToGlob(pathString: string, glob: string, options?: {
    windows?: boolean;
    escape?: boolean;
}): string;
export declare function resolveSourceMapPath(outputDir?: string, platform?: UniApp.PLATFORM): string;
export declare function installDepTips(type: 'dependencies' | 'devDependencies', module: string, version?: string): string;
/**
 * 根据路径判断是否为 App.(u?)vue
 * @param {string} filename 相对、绝对路径
 * @returns
 */
export declare function isAppVue(filename: string): boolean;
export declare function resolveAppVue(inputDir: string): string;
export declare function parseImporter(importer: string): string;
export declare function createResolveErrorMsg(source: string, importer: string): string;
export declare function enableSourceMap(): boolean;
export declare function requireUniHelpers(): any;
export declare function normalizeEmitAssetFileName(fileName: string): string;
export declare function createShadowImageUrl(cdn: number, type?: string): string;
