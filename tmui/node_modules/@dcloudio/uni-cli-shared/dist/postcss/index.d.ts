import type { Plugin } from 'postcss';
import type { Options } from 'autoprefixer';
import uniPostcssScopedPlugin from './plugins/stylePluginScoped';
import uniPostcssPlugin, { type UniAppCssProcessorOptions } from './plugins/uniapp';
export { uniPostcssPlugin };
export { uniPostcssScopedPlugin };
export declare function initPostcssPlugin({ uniApp, autoprefixer, }?: {
    uniApp?: UniAppCssProcessorOptions;
    autoprefixer?: Options | false;
}): Plugin[];
