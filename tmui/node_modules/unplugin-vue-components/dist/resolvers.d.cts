import { ComponentResolver, ComponentResolveResult } from './types.cjs';
import { FilterPattern } from '@rollup/pluginutils';
import 'unplugin';
import '@antfu/utils';

interface AntDesignVueResolverOptions {
    /**
     * exclude components that do not require automatic import
     *
     * @default []
     */
    exclude?: string[];
    /**
     * import style along with components
     *
     * @default 'css'
     */
    importStyle?: boolean | 'css' | 'less' | 'css-in-js';
    /**
     * resolve `ant-design-vue' icons
     *
     * requires package `@ant-design/icons-vue`
     *
     * @default false
     */
    resolveIcons?: boolean;
    /**
     * @deprecated use `importStyle: 'css'` instead
     */
    importCss?: boolean;
    /**
     * @deprecated use `importStyle: 'less'` instead
     */
    importLess?: boolean;
    /**
     * use commonjs build default false
     */
    cjs?: boolean;
    /**
     * rename package
     *
     * @default 'ant-design-vue'
     */
    packageName?: string;
    /**
     * customize prefix of component
     * @default 'A'
     */
    prefix?: string;
}
/**
 * Resolver for Ant Design Vue
 *
 * Requires ant-design-vue@v2.2.0-beta.6 or later
 *
 * See https://github.com/antfu/unplugin-vue-components/issues/26#issuecomment-789767941 for more details
 *
 * @author @yangss3
 * @link https://antdv.com/
 */
declare function AntDesignVueResolver(options?: AntDesignVueResolverOptions): ComponentResolver;

interface ElementPlusResolverOptions {
    /**
     * import style css or sass with components
     *
     * @default 'css'
     */
    importStyle?: boolean | 'css' | 'sass';
    /**
     * use commonjs lib & source css or scss for ssr
     */
    ssr?: boolean;
    /**
     * specify element-plus version to load style
     *
     * @default installed version
     */
    version?: string;
    /**
     * auto import for directives
     *
     * @default true
     */
    directives?: boolean;
    /**
     * exclude component name, if match do not resolve the name
     */
    exclude?: RegExp;
    /**
     * a list of component names that have no styles, so resolving their styles file should be prevented
     */
    noStylesComponents?: string[];
    /**
     * nightly version
     */
    nightly?: boolean;
}
/**
 * Resolver for Element Plus
 *
 * See https://github.com/antfu/vite-plugin-components/pull/28 for more details
 * See https://github.com/antfu/vite-plugin-components/issues/117 for more details
 *
 * @author @develar @nabaonan @sxzz
 * @link https://element-plus.org/ for element-plus
 *
 */
declare function ElementPlusResolver(options?: ElementPlusResolverOptions): ComponentResolver[];

interface ElementUiResolverOptions {
    /**
     * import style css or sass with components
     *
     * @default 'css'
     */
    importStyle?: boolean | 'css' | 'sass';
    /**
     * exclude component name, if match do not resolve the name
     */
    exclude?: RegExp;
}
/**
 * Resolver for Element-UI
 * @link https://element.eleme.cn/#/zh-CN
 * @version @element-ui ^2.15.3, @vue ^2.6.14
 * @author @nabaonan
 */
declare function ElementUiResolver(options?: ElementUiResolverOptions): ComponentResolver;

interface HeadlessUiResolverOptions {
    /**
     * prefix for headless ui components used in templates
     *
     * @default ""
     */
    prefix?: string;
}
/**
 * Resolver for headlessui
 *
 * @link https://github.com/tailwindlabs/headlessui
 */
declare function HeadlessUiResolver(options?: HeadlessUiResolverOptions): ComponentResolver;

interface IduxResolverOptions {
    /**
     * exclude components that do not require automatic import
     *
     * @default []
     */
    exclude?: string[];
    /**
     * import style along with components
     */
    importStyle?: 'css' | 'less';
    /**
     * theme for import style
     *
     * @default 'default' for 1.x version
     */
    importStyleTheme?: string;
    /**
     * The scope of the packages.
     *
     * @default '@idux'
     */
    scope?: string;
    /**
     * specify idux version to load style
     *
     * @default installed version
     */
    version?: string;
}
/**
 * Resolver for `@idux/cdk`, `@idux/components` and ``@idux/pro``
 *
 * @link https://idux.site
 */
declare function IduxResolver(options?: IduxResolverOptions): ComponentResolver;

/**
 * Resolver for Inkline
 *
 * @author @alexgrozav
 * @link https://github.com/inkline/inkline
 */
declare function InklineResolver(): ComponentResolver;

/**
 * Resolver for Naive UI
 *
 * @author @antfu
 * @link https://www.naiveui.com/
 */
declare function NaiveUiResolver(): ComponentResolver;

interface PrimeVueResolverOptions {
    /**
     * import style along with components
     *
     * @default true
     */
    importStyle?: boolean;
    /**
     * import `primeicons' icons
     *
     * requires package `primeicons`
     *
     * @default true
     */
    importIcons?: boolean;
    /**
     * imports a free theme - set theme name here (e.g. saga-blue)
     *
     * @default ''
     */
    importTheme?: string;
    /**
     * prefix for components (e.g. 'P' to resolve Menu from PMenu)
     *
     * @default ''
     */
    prefix?: string;
}
/**
 * Resolver for PrimeVue - If you're using a component with the same tag as an native HTML element (e.g. button) the component must be in uppercase
 *
 * @link https://github.com/primefaces/primevue
 */
declare function PrimeVueResolver(options?: PrimeVueResolverOptions): ComponentResolver;

interface VantResolverOptions {
    /**
     * import style css or less along with components
     *
     * @default true
     */
    importStyle?: boolean | 'css' | 'less';
}
/**
 * Resolver for Vant
 *
 * @link https://github.com/youzan/vant
 */
declare function VantResolver(options?: VantResolverOptions): ComponentResolver;

interface VarletUIResolverOptions {
    /**
     * support vue version
     * vue3 use @varlet/ui, vue2 use @varlet-vue2/ui
     *
     * @default 'vue3'
     */
    version?: 'vue3' | 'vue2';
    /**
     * import style along with components
     *
     * @default 'css'
     */
    importStyle?: boolean | 'css' | 'less';
    /**
     * style entry file extname
     *
     * @default '.mjs'
     */
    styleExtname?: string;
    /**
     * auto import for directives
     *
     * @default true
     */
    directives?: boolean;
    /**
     * compatible with unplugin-auto-import
     *
     * @default false
     */
    autoImport?: boolean;
    /**
     * @deprecated use `importStyle: 'css'` instead
     */
    importCss?: boolean;
    /**
     * @deprecated use `importStyle: 'less'` instead
     */
    importLess?: boolean;
}
declare function getResolved(name: string, options: VarletUIResolverOptions): ComponentResolveResult;
/**
 * Resolver for VarletUI
 *
 * @link https://github.com/varletjs/varlet
 * @link https://github.com/varletjs/varlet-vue2
 */
declare function VarletUIResolver(options?: VarletUIResolverOptions): ComponentResolver[];

interface VeuiPeerConfig {
    /**
     * The package name of the peer module.
     */
    package: string;
    /**
     * The directory path of the peer module.
     * @default 'components'
     */
    path?: string;
    /**
     * The file name template for the peer module.
     * @default '{module}.css'
     */
    fileName?: `${string}{module}${string}`;
    /**
     * The text transform to be applied to the '{module}' part of the file name.
     * @default 'kebab-case'
     */
    transform?: 'kebab-case' | 'camelCase' | 'PascalCase' | false;
}
type SupportedLocale = 'en-US' | 'zh-Hans';
interface VeuiResolverOptions {
    /**
     * The alias of 'veui` package.
     * @default 'veui'
     */
    alias?: string;
    /**
     * Peer modules to be injected.
     */
    modules?: VeuiPeerConfig[];
    /**
     * Locale modules to be injected.
     * @default 'zh-Hans'
     */
    locale?: SupportedLocale | SupportedLocale[] | false;
    /**
     * Global modules to be injected to all components.
     * @default []
     */
    global?: string[];
}
/**
 * Resolver for VEUI
 *
 * @link https://github.com/ecomfe/veui
 */
declare function VeuiResolver(options?: VeuiResolverOptions): ComponentResolver;

/**
 * Resolver for View UI
 * @requires @originjs/vite-plugin-commonjs
 * @author @nabaonan
 * @link https://www.iviewui.com/
 * @description has known problems list below
 * - select component render error PR: https://github.com/view-design/ViewUI/pull/944,  choose can't display value,because click option trigger twice,at second time,select value turn into undefined.
 * - scroll component has a template syntax called lang='html',it is require html-loader,but vite plugin not support yet,remove it can run. relate pr: https://github.com/view-design/ViewUI/pull/985
 */
declare function ViewUiResolver(): ComponentResolver;

/**
 * Resolver for Vuetify
 *
 * @link https://github.com/vuetifyjs/vuetify
 */
declare function VuetifyResolver(): ComponentResolver;
/**
 * Resolver for Vuetify 3 Beta
 *
 * @link https://github.com/vuetifyjs/vuetify
 */
declare function Vuetify3Resolver(): ComponentResolver;

/**
 * Resolver for VueUse
 *
 * @link https://github.com/vueuse/vueuse
 */
declare function VueUseComponentsResolver(): ComponentResolver;

/**
 * Resolver for VueUse
 *
 * @link https://github.com/vueuse/vueuse
 */
declare function VueUseDirectiveResolver(): ComponentResolver;

/**
 * Resolver for Quasar
 *
 * @link https://github.com/quasarframework/quasar
 */
declare function QuasarResolver(): ComponentResolver;

interface DevResolverOptions {
    /**
     * bring in components and styles
     *
     * @default true
     */
    importStyle?: boolean;
    /**
     * auto import for directives
     *
     * @default true
     */
    directives?: boolean;
    /**
     * use umd lib file
     */
    ssr?: boolean;
}
declare function DevUiResolver(options?: DevResolverOptions): ComponentResolver[];

type DisallowResolveIconOption = undefined | false | {
    enable: false;
};
type AllowResolveIconOption = true | {
    enable: true;
    iconPrefix?: string;
};
type ResolveIconsOption = DisallowResolveIconOption | AllowResolveIconOption;
interface ArcoResolverOptions {
    /**
     * exclude components that do not require automatic import
     *
     * @default []
     */
    exclude?: string | RegExp | (string | RegExp)[];
    /**
     * import style css or less with components
     *
     * @default 'css'
     */
    importStyle?: boolean | 'css' | 'less';
    /**
     * resolve icons
     *
     * @default false
     */
    resolveIcons?: ResolveIconsOption;
    /**
     * Control style automatic import
     *
     * @default true
     */
    sideEffect?: boolean;
}
/**
 * Resolver for Arco Design Vue
 *
 * Requires arco-design/web-vue@2.11.0 or later
 *
 * @author @flsion
 * @link https://arco.design/ for arco-design
 *
 */
declare function ArcoResolver(options?: ArcoResolverOptions): ComponentResolver;

interface TDesignResolverOptions {
    /**
     * select the specified library
     * @default 'vue'
     */
    library?: 'vue' | 'vue-next' | 'react' | 'mobile-vue' | 'mobile-react';
    /**
     * resolve `tdesign-icons'
     * @default false
     */
    resolveIcons?: boolean;
    /**
     * whether to import ESM version
     * @default false
     */
    esm?: boolean;
    /**
     * exclude component name, if match do not resolve the name
     *
     */
    exclude?: FilterPattern;
}
declare function TDesignResolver(options?: TDesignResolverOptions): ComponentResolver;

interface LayuiVueResolverOptions {
    /**
     * import style along with components
     *
     * @default 'css'
     */
    importStyle?: boolean | 'css';
    /**
     * resolve '@layui/layui-vue' icons
     * requires package `@layui/icons-vue`
     *
     * @default false
     */
    resolveIcons?: boolean;
    /**
     * exclude components that do not require automatic import
     *
     */
    exclude?: FilterPattern;
}
/**
 * Resolver for layui-vue
 *
 * @link http://www.layui-vue.com/ for layui-vue
 *
 */
declare function LayuiVueResolver(options?: LayuiVueResolverOptions): ComponentResolver;

interface BootstrapVueResolverOptions {
    /**
     * Auto import for directives.
     *
     * @default true
     */
    directives?: boolean;
}
/**
 * Resolver for BootstrapVue
 *
 * @link https://github.com/bootstrap-vue/bootstrap-vue
 */
declare function BootstrapVueResolver(_options?: BootstrapVueResolverOptions): ComponentResolver[];
/**
 * Resolver for BootstrapVueNext
 *
 * @link https://github.com/bootstrap-vue/bootstrap-vue-next
 *
 * @deprecated use `import { BootstrapVueNextResolver } from 'bootstrap-vue-next'` instead
 */
declare function BootstrapVueNextResolver(_options?: BootstrapVueResolverOptions): Array<ComponentResolver>;
/**
 * Resolver for legacy BootstrapVue3 apps
 *
 * @deprecated use BootstrapVueNextResolver with https://github.com/bootstrap-vue/bootstrap-vue-next
 * @link https://www.npmjs.com/package/bootstrap-vue-3
 */
declare function BootstrapVue3Resolver(_options?: BootstrapVueResolverOptions): Array<ComponentResolver>;

/**
 * source: https://github.com/nuxt-modules/ionic/blob/main/src/imports.ts
 * @author @danielroe
 */
declare const IonicBuiltInComponents: string[];
/**
 * Resolver for ionic framework
 *
 * @author @mathsgod @reslear
 * @link https://ionicframework.com/
 */
declare function IonicResolver(): ComponentResolver;

export { type AllowResolveIconOption, AntDesignVueResolver, type AntDesignVueResolverOptions, ArcoResolver, type ArcoResolverOptions, BootstrapVue3Resolver, BootstrapVueNextResolver, BootstrapVueResolver, type BootstrapVueResolverOptions, type DevResolverOptions, DevUiResolver, type DisallowResolveIconOption, ElementPlusResolver, type ElementPlusResolverOptions, ElementUiResolver, type ElementUiResolverOptions, HeadlessUiResolver, type HeadlessUiResolverOptions, IduxResolver, type IduxResolverOptions, InklineResolver, IonicBuiltInComponents, IonicResolver, LayuiVueResolver, type LayuiVueResolverOptions, NaiveUiResolver, PrimeVueResolver, type PrimeVueResolverOptions, QuasarResolver, type ResolveIconsOption, TDesignResolver, type TDesignResolverOptions, VantResolver, type VantResolverOptions, VarletUIResolver, type VarletUIResolverOptions, VeuiResolver, type VeuiResolverOptions, ViewUiResolver, VueUseComponentsResolver, VueUseDirectiveResolver, Vuetify3Resolver, VuetifyResolver, getResolved };
