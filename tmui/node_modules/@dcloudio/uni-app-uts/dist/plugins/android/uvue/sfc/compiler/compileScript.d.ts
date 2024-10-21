import type { SFCDescriptor, SFCScriptBlock } from '@vue/compiler-sfc';
import type { ParserPlugin } from '@babel/parser';
import type { SFCTemplateCompileOptions } from '@vue/compiler-sfc';
export declare const normalScriptDefaultVar = "__default__";
export declare const DEFAULT_FILENAME = "anonymous.vue";
export interface SFCScriptCompileOptions {
    root: string;
    /**
     * 类型
     */
    componentType: 'app' | 'page' | 'component';
    /**
     * 是否同时支持使用 <script> 和 <script setup>
     */
    scriptAndScriptSetup?: boolean;
    /**
     * Class name
     */
    className: string;
    /**
     * Scope ID for prefixing injected CSS variables.
     * This must be consistent with the `id` passed to `compileStyle`.
     */
    id: string;
    /**
     * Production mode. Used to determine whether to generate hashed CSS variables
     */
    isProd?: boolean;
    /**
     * Enable/disable source map. Defaults to true.
     */
    sourceMap?: boolean;
    /**
     * https://babeljs.io/docs/en/babel-parser#plugins
     */
    babelParserPlugins?: ParserPlugin[];
    /**
     * A list of files to parse for global types to be made available for type
     * resolving in SFC macros. The list must be fully resolved file system paths.
     */
    globalTypeFiles?: string[];
    /**
     * Compile the template and inline the resulting render function
     * directly inside setup().
     * - Only affects `<script setup>`
     * - This should only be used in production because it prevents the template
     * from being hot-reloaded separately from component state.
     */
    inlineTemplate?: boolean;
    /**
     * Generate the final component as a variable instead of default export.
     * This is useful in e.g. @vitejs/plugin-vue where the script needs to be
     * placed inside the main module.
     */
    genDefaultAs?: string;
    /**
     * Options for template compilation when inlining. Note these are options that
     * would normally be passed to `compiler-sfc`'s own `compileTemplate()`, not
     * options passed to `compiler-dom`.
     */
    templateOptions?: Partial<SFCTemplateCompileOptions>;
    /**
     * Hoist <script setup> static constants.
     * - Only enables when one `<script setup>` exists.
     * @default true
     */
    hoistStatic?: boolean;
    /**
     * (**Experimental**) Enable macro `defineModel`
     * @default false
     */
    defineModel?: boolean;
    /**
     * (**Experimental**) Enable reactive destructure for `defineProps`
     * @default false
     */
    propsDestructure?: boolean;
    /**
     * File system access methods to be used when resolving types
     * imported in SFC macros. Defaults to ts.sys in Node.js, can be overwritten
     * to use a virtual file system for use in browsers (e.g. in REPLs)
     */
    fs?: {
        fileExists(file: string): boolean;
        readFile(file: string): string | undefined;
    };
    /**
     * (Experimental) Enable syntax transform for using refs without `.value` and
     * using destructured props with reactivity
     * @deprecated the Reactivity Transform proposal has been dropped. This
     * feature will be removed from Vue core in 3.4. If you intend to continue
     * using it, disable this and switch to the [Vue Macros implementation](https://vue-macros.sxzz.moe/features/reactivity-transform.html).
     */
    reactivityTransform?: boolean;
    /**
     * Transform Vue SFCs into custom elements.
     */
    customElement?: boolean | ((filename: string) => boolean);
}
export interface ImportBinding {
    isType: boolean;
    imported: string;
    local: string;
    source: string;
    isFromSetup: boolean;
    isUsedInTemplate: boolean;
}
/**
 * Compile `<script setup>`
 * It requires the whole SFC descriptor because we need to handle and merge
 * normal `<script>` + `<script setup>` if both are present.
 */
export declare function compileScript(sfc: SFCDescriptor, options: SFCScriptCompileOptions): SFCScriptBlock;
