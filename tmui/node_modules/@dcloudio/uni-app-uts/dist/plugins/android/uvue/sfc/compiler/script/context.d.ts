import type { CallExpression, Node, ObjectPattern, Program, TSInterfaceDeclaration, TSTypeLiteral } from '@babel/types';
import type { SFCDescriptor } from '@vue/compiler-sfc';
import { type ParserPlugin } from '@babel/parser';
import type { ImportBinding, SFCScriptCompileOptions } from '../compileScript';
import type { PropsDestructureBindings } from './defineProps';
import type { ModelDecl } from './defineModel';
import type { BindingMetadata } from '@vue/compiler-core';
import MagicString from 'magic-string';
import type { TypeScope } from './resolveType';
export declare class ScriptCompileContext {
    descriptor: SFCDescriptor;
    options: Partial<SFCScriptCompileOptions>;
    scriptAst: Program | null;
    scriptSetupAst: Program | null;
    source: string;
    filename: string;
    s: MagicString;
    startOffset: number | undefined;
    endOffset: number | undefined;
    scope?: TypeScope;
    globalScopes?: TypeScope[];
    userImports: Record<string, ImportBinding>;
    hasDefinePropsCall: boolean;
    hasDefineEmitCall: boolean;
    hasDefineExposeCall: boolean;
    hasDefaultExportName: boolean;
    hasDefaultExportRender: boolean;
    hasDefineOptionsCall: boolean;
    hasDefineSlotsCall: boolean;
    hasDefineModelCall: boolean;
    propsCall: CallExpression | undefined;
    propsDecl: Node | undefined;
    propsRuntimeDecl: Node | undefined;
    propsTypeDecl: Node | undefined;
    propsDestructureDecl: ObjectPattern | undefined;
    propsDestructuredBindings: PropsDestructureBindings;
    propsDestructureRestId: string | undefined;
    propsRuntimeDefaults: Node | undefined;
    propsInterfaceDecl: TSInterfaceDeclaration | undefined;
    emitsRuntimeDecl: Node | undefined;
    emitsTypeDecl: Node | undefined;
    emitDecl: Node | undefined;
    modelDecls: Record<string, ModelDecl>;
    optionsRuntimeDecl: Node | undefined;
    slotsRuntimeDecl: TSTypeLiteral | undefined;
    bindingMetadata: BindingMetadata;
    helperImports: Set<string>;
    helper(key: string): string;
    /**
     * to be exposed on compiled script block for HMR cache busting
     */
    deps?: Set<string>;
    /**
     * cache for resolved fs
     */
    fs?: NonNullable<SFCScriptCompileOptions['fs']>;
    constructor(descriptor: SFCDescriptor, options: Partial<SFCScriptCompileOptions>);
    getString(node: Node, scriptSetup?: boolean): string;
    error(msg: string, node: Node, scope?: TypeScope): never;
}
export declare function resolveParserPlugins(lang: string, userPlugins?: ParserPlugin[], dts?: boolean): (import("@babel/parser").ParserPluginWithOptions | ("asyncDoExpressions" | "asyncGenerators" | "bigInt" | "classPrivateMethods" | "classPrivateProperties" | "classProperties" | "classStaticBlock" | "decimal" | "decorators-legacy" | "deferredImportEvaluation" | "decoratorAutoAccessors" | "destructuringPrivate" | "doExpressions" | "dynamicImport" | "explicitResourceManagement" | "exportDefaultFrom" | "exportNamespaceFrom" | "flow" | "flowComments" | "functionBind" | "functionSent" | "importMeta" | "jsx" | "logicalAssignment" | "importAssertions" | "importAttributes" | "importReflection" | "moduleBlocks" | "moduleStringNames" | "nullishCoalescingOperator" | "numericSeparator" | "objectRestSpread" | "optionalCatchBinding" | "optionalChaining" | "partialApplication" | "placeholders" | "privateIn" | "regexpUnicodeSets" | "sourcePhaseImports" | "throwExpressions" | "topLevelAwait" | "v8intrinsic" | "decorators" | "estree" | "moduleAttributes" | "optionalChainingAssign" | "pipelineOperator" | "recordAndTuple" | "typescript"))[];
