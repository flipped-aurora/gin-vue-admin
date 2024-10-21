import { Arrayable, Awaitable } from '@antfu/utils';
import { FilterPattern } from '@rollup/pluginutils';
import { Import, InlinePreset } from 'unimport';
import * as unimport_dist_shared_unimport_b55a67ec from 'unimport/dist/shared/unimport.b55a67ec';

declare const presets: {
    ahooks: () => ImportsMap;
    '@vueuse/core': () => ImportsMap;
    '@vueuse/math': () => ImportsMap;
    '@vueuse/head': ImportsMap;
    mobx: ImportsMap;
    'mobx-react-lite': ImportsMap;
    preact: ImportsMap;
    quasar: ImportsMap;
    react: ImportsMap;
    'react-router': ImportsMap;
    'react-router-dom': ImportsMap;
    'react-i18next': ImportsMap;
    svelte: ImportsMap;
    'svelte/animate': ImportsMap;
    'svelte/easing': ImportsMap;
    'svelte/motion': ImportsMap;
    'svelte/store': ImportsMap;
    'svelte/transition': ImportsMap;
    'vee-validate': ImportsMap;
    vitepress: ImportsMap;
    'vue-router': ImportsMap;
    'vue-router/composables': ImportsMap;
    vuex: ImportsMap;
    'uni-app': ImportsMap;
    'solid-js': ImportsMap;
    '@solidjs/router': ImportsMap;
    'solid-app-router': ImportsMap;
    jotai: ImportsMap;
    'jotai/utils': ImportsMap;
    recoil: ImportsMap;
    '@vue/composition-api': unimport_dist_shared_unimport_b55a67ec.a;
    pinia: unimport_dist_shared_unimport_b55a67ec.a;
    'vue-demi': unimport_dist_shared_unimport_b55a67ec.a;
    'vue-i18n': unimport_dist_shared_unimport_b55a67ec.a;
    'vue-router-composables': unimport_dist_shared_unimport_b55a67ec.a;
    vue: unimport_dist_shared_unimport_b55a67ec.a;
    'vue/macros': unimport_dist_shared_unimport_b55a67ec.a;
    vitest: unimport_dist_shared_unimport_b55a67ec.a;
    rxjs: unimport_dist_shared_unimport_b55a67ec.a;
};
type PresetName = keyof typeof presets;

interface ImportLegacy {
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
interface ImportExtended extends Import {
    sideEffects?: SideEffectsInfo;
    __source?: 'dir' | 'resolver';
}
type ImportNameAlias = [string, string];
type SideEffectsInfo = Arrayable<ResolverResult | string> | undefined;
interface ResolverResult {
    as?: string;
    name?: string;
    from: string;
}
type ResolverFunction = (name: string) => Awaitable<string | ResolverResult | ImportExtended | null | undefined | void>;
interface ResolverResultObject {
    type: 'component' | 'directive';
    resolve: ResolverFunction;
}
/**
 * Given a identifier name, returns the import path or an import object
 */
type Resolver = ResolverFunction | ResolverResultObject;
/**
 * module, name, alias
 */
type ImportsMap = Record<string, (string | ImportNameAlias)[]>;
type ESLintGlobalsPropValue = boolean | 'readonly' | 'readable' | 'writable' | 'writeable';
interface ESLintrc {
    /**
     * @default false
     */
    enabled?: boolean;
    /**
     * Filepath to save the generated eslint config
     *
     * @default './.eslintrc-auto-import.json'
     */
    filepath?: string;
    /**
     * @default true
     */
    globalsPropValue?: ESLintGlobalsPropValue;
}
interface Options {
    /**
     * Preset names or custom imports map
     *
     * @default []
     */
    imports?: Arrayable<ImportsMap | PresetName | InlinePreset>;
    /**
     * Identifiers to be ignored
     */
    ignore?: (string | RegExp)[];
    /**
     * Inject the imports at the end of other imports
     *
     * @default true
     */
    injectAtEnd?: boolean;
    /**
     * Path for directories to be auto imported
     */
    dirs?: string[];
    /**
     * Pass a custom function to resolve the component importing path from the component name.
     *
     * The component names are always in PascalCase
     */
    resolvers?: Arrayable<Arrayable<Resolver>>;
    /**
     * Filepath to generate corresponding .d.ts file.
     * Default enabled when `typescript` is installed locally.
     * Set `false` to disable.
     *
     * @default './auto-imports.d.ts'
     */
    dts?: string | boolean;
    /**
     * Auto import inside Vue templates
     *
     * @see https://github.com/unjs/unimport/pull/15
     * @see https://github.com/unjs/unimport/pull/72
     * @default false
     */
    vueTemplate?: boolean;
    /**
     * Set default export alias by file name
     *
     * @default false
     */
    defaultExportByFilename?: boolean;
    /**
     * Rules to include transforming target.
     *
     * @default [/\.[jt]sx?$/, /\.vue\??/]
     */
    include?: FilterPattern;
    /**
     * Rules to exclude transforming target.
     *
     * @default [/node_modules/, /\.git/]
     */
    exclude?: FilterPattern;
    /**
     * Generate corresponding .eslintrc-auto-import.json file.
     */
    eslintrc?: ESLintrc;
}

export { ESLintGlobalsPropValue, ESLintrc, ImportExtended, ImportLegacy, ImportNameAlias, ImportsMap, Options, PresetName, Resolver, ResolverFunction, ResolverResult, ResolverResultObject, SideEffectsInfo };
