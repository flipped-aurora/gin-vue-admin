'use strict';

const context = require('./shared/unimport.3c66320c.cjs');
const addons = require('./shared/unimport.8977d7da.cjs');
require('mlly');
require('node:fs');
require('node:fs/promises');
require('node:path');
require('node:process');
require('node:url');
require('fast-glob');
require('pathe');
require('scule');
require('node:os');
require('pkg-types');
require('local-pkg');
require('magic-string');
require('strip-literal');

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

exports.builtinPresets = context.builtinPresets;
exports.createUnimport = context.createUnimport;
exports.dedupeDtsExports = context.dedupeDtsExports;
exports.resolveBuiltinPresets = context.resolveBuiltinPresets;
exports.resolvePreset = context.resolvePreset;
exports.scanDirExports = context.scanDirExports;
exports.scanExports = context.scanExports;
exports.scanFilesFromDir = context.scanFilesFromDir;
exports.version = context.version;
exports.addImportToCode = addons.addImportToCode;
exports.dedupeImports = addons.dedupeImports;
exports.defineUnimportPreset = addons.defineUnimportPreset;
exports.excludeRE = addons.excludeRE;
exports.getMagicString = addons.getMagicString;
exports.getString = addons.getString;
exports.importAsRE = addons.importAsRE;
exports.matchRE = addons.matchRE;
exports.normalizeImports = addons.normalizeImports;
exports.resolveIdAbsolute = addons.resolveIdAbsolute;
exports.separatorRE = addons.separatorRE;
exports.stringifyImports = addons.stringifyImports;
exports.stripCommentsAndStrings = addons.stripCommentsAndStrings;
exports.stripFileExtension = addons.stripFileExtension;
exports.toExports = addons.toExports;
exports.toImports = addons.toImports;
exports.toTypeDeclarationFile = addons.toTypeDeclarationFile;
exports.toTypeDeclarationItems = addons.toTypeDeclarationItems;
exports.toTypeReExports = addons.toTypeReExports;
exports.vueTemplateAddon = addons.vueTemplateAddon;
exports.installGlobalAutoImports = installGlobalAutoImports;
