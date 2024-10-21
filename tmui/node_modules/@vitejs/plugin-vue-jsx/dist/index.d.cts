import { FilterPattern, Plugin } from 'vite';
import { VueJSXPluginOptions } from '@vue/babel-plugin-jsx';

interface FilterOptions {
    include?: FilterPattern;
    exclude?: FilterPattern;
}
type Options = VueJSXPluginOptions & FilterOptions & {
    babelPlugins?: any[];
};

declare function vueJsxPlugin(options?: Options): Plugin;

export { type FilterOptions, type Options, vueJsxPlugin as default };
