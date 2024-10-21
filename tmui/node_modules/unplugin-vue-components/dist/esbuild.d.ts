import * as esbuild from 'esbuild';
import { Options } from './types.js';
import '@rollup/pluginutils';
import 'unplugin';
import '@antfu/utils';

declare const _default: (options: Options) => esbuild.Plugin;

export { _default as default };
