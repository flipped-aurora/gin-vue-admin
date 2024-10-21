import * as webpack from 'webpack';
import { Options } from './types.cjs';
import '@rollup/pluginutils';
import 'unplugin';
import '@antfu/utils';

declare const _default: (options: Options) => webpack.WebpackPluginInstance;

export { _default as default };
