import { Plugin } from 'vite';
import { Options, PublicPluginAPI } from './types.js';
import '@rollup/pluginutils';
import 'unplugin';
import '@antfu/utils';

declare const _default: (options?: Options | undefined) => Plugin & {
    api: PublicPluginAPI;
};

export { _default as default };
