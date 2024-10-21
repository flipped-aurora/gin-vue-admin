import type { Node, Statement, TSCallSignatureDeclaration, TSFunctionType, TSMethodSignature, TSModuleDeclaration, TSPropertySignature } from '@babel/types';
import { type ScriptCompileContext } from './context';
import type { ImportBinding } from '../compileScript';
/**
 * TypeResolveContext is compatible with ScriptCompileContext
 * but also allows a simpler version of it with minimal required properties
 * when resolveType needs to be used in a non-SFC context, e.g. in a babel
 * plugin. The simplest context can be just:
 * ```ts
 * const ctx: SimpleTypeResolveContext = {
 *   filename: '...',
 *   source: '...',
 *   options: {},
 *   error() {},
 *   ast: []
 * }
 * ```
 */
export type SimpleTypeResolveContext = Pick<ScriptCompileContext, 'source' | 'filename' | 'error' | 'options'> & Partial<Pick<ScriptCompileContext, 'scope' | 'globalScopes' | 'deps' | 'fs'>> & {
    ast: Statement[];
};
export type TypeResolveContext = ScriptCompileContext | SimpleTypeResolveContext;
type Import = Pick<ImportBinding, 'source' | 'imported'>;
interface WithScope {
    _ownerScope: TypeScope;
}
type ScopeTypeNode = Node & WithScope & {
    _ns?: TSModuleDeclaration & WithScope;
};
export declare class TypeScope {
    filename: string;
    source: string;
    offset: number;
    imports: Record<string, Import>;
    types: Record<string, ScopeTypeNode>;
    declares: Record<string, ScopeTypeNode>;
    constructor(filename: string, source: string, offset?: number, imports?: Record<string, Import>, types?: Record<string, ScopeTypeNode>, declares?: Record<string, ScopeTypeNode>);
    resolvedImportSources: Record<string, string>;
    exportedTypes: Record<string, ScopeTypeNode>;
    exportedDeclares: Record<string, ScopeTypeNode>;
}
export interface MaybeWithScope {
    _ownerScope?: TypeScope;
}
interface ResolvedElements {
    props: Record<string, (TSPropertySignature | TSMethodSignature) & {
        _ownerScope: TypeScope;
    }>;
    calls?: (TSCallSignatureDeclaration | TSFunctionType)[];
}
/**
 * Resolve arbitrary type node to a list of type elements that can be then
 * mapped to runtime props or emits.
 */
export declare function resolveTypeElements(ctx: TypeResolveContext, node: Node & MaybeWithScope & {
    _resolvedElements?: ResolvedElements;
}, scope?: TypeScope, typeParameters?: Record<string, Node>): ResolvedElements;
/**
 * @private
 */
export declare function invalidateTypeCache(filename: string): void;
export declare function fileToScope(ctx: TypeResolveContext, filename: string, asGlobal?: boolean): TypeScope;
export declare function recordImports(body: Statement[]): Record<string, Import>;
export declare function inferRuntimeType(ctx: TypeResolveContext, node: Node & MaybeWithScope, from?: 'defineProps' | 'defineModel', scope?: TypeScope): string[];
export declare function resolveUnionType(ctx: TypeResolveContext, node: Node & MaybeWithScope & {
    _resolvedElements?: ResolvedElements;
}, scope?: TypeScope): Node[];
export {};
