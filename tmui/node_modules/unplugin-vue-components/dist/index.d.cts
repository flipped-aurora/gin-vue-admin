import { Options } from './types.cjs';
export { ComponentInfo, ComponentResolveResult, ComponentResolver, ComponentResolverFunction, ComponentResolverObject, ComponentsImportMap, ImportInfo, ImportInfoLegacy, Matcher, PublicPluginAPI, ResolvedOptions, SideEffectsInfo, SupportedTransformer, Transformer, TypeImport } from './types.cjs';
import * as unplugin from 'unplugin';
import '@rollup/pluginutils';
import '@antfu/utils';

declare const _default: unplugin.UnpluginInstance<Options, boolean>;

declare function pascalCase(str: string): string;
declare function camelCase(str: string): string;
declare function kebabCase(key: string): string;

export { Options, camelCase, _default as default, kebabCase, pascalCase };
