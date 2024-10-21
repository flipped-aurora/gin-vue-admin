import { promises } from 'node:fs';
import { createFilter } from '@rollup/pluginutils';
import MagicString from 'magic-string';
import { createUnplugin } from 'unplugin';
import { c as createUnimport } from './shared/unimport.1c7b6182.mjs';
import './shared/unimport.85ddadbb.mjs';
import 'node:path';
import 'node:process';
import 'pathe';
import 'scule';
import 'mlly';
import 'strip-literal';
import 'node:fs/promises';
import 'node:url';
import 'fast-glob';
import 'node:os';
import 'pkg-types';
import 'local-pkg';

const defaultIncludes = [/\.[jt]sx?$/, /\.vue$/, /\.vue\?vue/, /\.svelte$/];
const defaultExcludes = [/[\\/]node_modules[\\/]/, /[\\/]\.git[\\/]/];
function toArray(x) {
  return x == null ? [] : Array.isArray(x) ? x : [x];
}
const unplugin = createUnplugin((options = {}) => {
  const ctx = createUnimport(options);
  const filter = createFilter(
    toArray(options.include || []).length ? options.include : defaultIncludes,
    options.exclude || defaultExcludes
  );
  const dts = options.dts === true ? "unimport.d.ts" : options.dts;
  const {
    autoImport = true
  } = options;
  return {
    name: "unimport",
    enforce: "post",
    transformInclude(id) {
      return filter(id);
    },
    async transform(code, id) {
      const s = new MagicString(code);
      await ctx.injectImports(s, id, {
        autoImport
      });
      if (!s.hasChanged())
        return;
      return {
        code: s.toString(),
        map: s.generateMap()
      };
    },
    async buildStart() {
      await ctx.init();
      if (dts)
        return promises.writeFile(dts, await ctx.generateTypeDeclarations(), "utf-8");
    }
  };
});

export { unplugin as default, defaultExcludes, defaultIncludes };
