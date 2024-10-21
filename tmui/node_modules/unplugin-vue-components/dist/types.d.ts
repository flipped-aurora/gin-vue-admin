import { FilterPattern } from '@rollup/pluginutils';
import { TransformResult } from 'unplugin';
import { Awaitable } from '@antfu/utils';

interface ImportInfoLegacy {
    /**
     * @deprecated renamed to `as`
     */
    name?: string;
    /**
     * @deprecated renamed to `name`
     */
    importName?: string;
    /**
     * @deprecated renamed to `from`
     */
    path: string;
    sideEffects?: SideEffectsInfo;
}
interface ImportInfo {
    as?: string;
    name?: string;
    from: string;
}
type SideEffectsInfo = (ImportInfo | string)[] | ImportInfo | string | undefined;
interface ComponentInfo extends ImportInfo {
    sideEffects?: SideEffectsInfo;
}
type ComponentResolveResult = Awaitable<string | ComponentInfo | null | undefined | void>;
type ComponentResolverFunction = (name: string) => ComponentResolveResult;
interface ComponentResolverObject {
    type: 'component' | 'directive';
    resolve: ComponentResolverFunction;
}
type ComponentResolver = ComponentResolverFunction | ComponentResolverObject;
type Matcher = (id: string) => boolean | null | undefined;
type Transformer = (code: string, id: string, path: string, query: Record<string, string>) => Awaitable<TransformResult | null>;
type SupportedTransformer = 'vue3' | 'vue2';
interface PublicPluginAPI {
    /**
     * Resolves a component using the configured resolvers.
     */
    findComponent: (name: string, filename?: string) => Promise<ComponentInfo | undefined>;
    /**
     * Obtain an import statement for a resolved component.
     */
    stringifyImport: (info: ComponentInfo) => string;
}
/**
 * Plugin options.
 */
interface Options {
    /**
     * RegExp or glob to match files to be transformed
     */
    include?: FilterPattern;
    /**
     * RegExp or glob to match files to NOT be transformed
     */
    exclude?: FilterPattern;
    /**
     * RegExp or string to match component names that will NOT be imported
     */
    excludeNames?: FilterPattern;
    /**
     * Relative paths to the directory to search for components.
     * @default 'src/components'
     */
    dirs?: string | string[];
    /**
     * Valid file extensions for components.
     * @default ['vue']
     */
    extensions?: string | string[];
    /**
     * Glob patterns to match file names to be detected as components.
     *
     * When specified, the `dirs`, `extensions`, and `directoryAsNamespace` options will be ignored.
     */
    globs?: string | string[];
    /**
     * Search for subdirectories
     * @default true
     */
    deep?: boolean;
    /**
     * Allow subdirectories as namespace prefix for components
     * @default false
     */
    directoryAsNamespace?: boolean;
    /**
     * Collapse same prefixes (camel-sensitive) of folders and components
     * to prevent duplication inside namespaced component name.
     *
     * Works when `directoryAsNamespace: true`
     * @default false
     */
    collapseSamePrefixes?: boolean;
    /**
     * Subdirectory paths for ignoring namespace prefixes
     *
     * Works when `directoryAsNamespace: true`
     * @default "[]"
     */
    globalNamespaces?: string[];
    /**
     * Pass a custom function to resolve the component importing path from the component name.
     *
     * The component names are always in PascalCase
     */
    resolvers?: (ComponentResolver | ComponentResolver[])[];
    /**
     * Apply custom transform over the path for importing
     */
    importPathTransform?: (path: string) => string | undefined;
    /**
     * Transformer to apply
     *
     * @default 'vue3'
     */
    transformer?: SupportedTransformer;
    /**
     * Generate TypeScript declaration for global components
     *
     * Accept boolean or a path related to project root
     *
     * @see https://github.com/vuejs/core/pull/3399
     * @see https://github.com/johnsoncodehk/volar#using
     * @default true
     */
    dts?: boolean | string;
    /**
     * Do not emit warning on component overriding
     *
     * @default false
     */
    allowOverrides?: boolean;
    /**
     * auto import for directives.
     *
     * default: `true` for Vue 3, `false` for Vue 2
     *
     * Babel is needed to do the transformation for Vue 2, it's disabled by default for performance concerns.
     * To install Babel, run: `npm install -D @babel/parser`
     * @default undefined
     */
    directives?: boolean;
    /**
     * Only provide types of components in library (registered globally)
     */
    types?: TypeImport[];
    /**
     * Vue version of project. It will detect automatically if not specified.
     */
    version?: 2 | 2.7 | 3;
}
type ResolvedOptions = Omit<Required<Options>, 'resolvers' | 'extensions' | 'dirs' | 'globalComponentsDeclaration'> & {
    resolvers: ComponentResolverObject[];
    extensions: string[];
    dirs: string[];
    resolvedDirs: string[];
    globs: string[];
    dts: string | false;
    root: string;
};
type ComponentsImportMap = Record<string, string[] | undefined>;
interface TypeImport {
    from: string;
    names: string[];
}

export type { ComponentInfo, ComponentResolveResult, ComponentResolver, ComponentResolverFunction, ComponentResolverObject, ComponentsImportMap, ImportInfo, ImportInfoLegacy, Matcher, Options, PublicPluginAPI, ResolvedOptions, SideEffectsInfo, SupportedTransformer, Transformer, TypeImport };
