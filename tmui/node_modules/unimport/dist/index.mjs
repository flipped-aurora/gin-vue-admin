export { b as builtinPresets, c as createUnimport, e as dedupeDtsExports, r as resolveBuiltinPresets, a as resolvePreset, d as scanDirExports, f as scanExports, s as scanFilesFromDir, v as version } from './shared/unimport.1c7b6182.mjs';
export { n as addImportToCode, c as dedupeImports, d as defineUnimportPreset, e as excludeRE, l as getMagicString, k as getString, i as importAsRE, m as matchRE, o as normalizeImports, r as resolveIdAbsolute, s as separatorRE, b as stringifyImports, a as stripCommentsAndStrings, f as stripFileExtension, t as toExports, p as toImports, h as toTypeDeclarationFile, g as toTypeDeclarationItems, j as toTypeReExports, v as vueTemplateAddon } from './shared/unimport.85ddadbb.mjs';
import 'mlly';
import 'node:fs';
import 'node:fs/promises';
import 'node:path';
import 'node:process';
import 'node:url';
import 'fast-glob';
import 'pathe';
import 'scule';
import 'node:os';
import 'pkg-types';
import 'local-pkg';
import 'magic-string';
import 'strip-literal';

async function installGlobalAutoImports(imports, options = {}) {
  const {
    globalObject = globalThis,
    overrides = false
  } = options;
  imports = Array.isArray(imports) ? imports : await imports.getImports();
  await Promise.all(
    imports.map(async (i) => {
      if (i.disabled || i.type)
        return;
      const as = i.as || i.name;
      if (overrides || !(as in globalObject)) {
        const module = await import(i.from);
        globalObject[as] = module[i.name];
      }
    })
  );
  return globalObject;
}

export { installGlobalAutoImports };
