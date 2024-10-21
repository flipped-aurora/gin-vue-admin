'use strict';

const path = require('node:path');
const process = require('node:process');
const pathe = require('pathe');
const scule = require('scule');
const MagicString = require('magic-string');
const mlly = require('mlly');
const stripLiteral = require('strip-literal');

function _interopDefaultCompat (e) { return e && typeof e === 'object' && 'default' in e ? e.default : e; }

const process__default = /*#__PURE__*/_interopDefaultCompat(process);
const MagicString__default = /*#__PURE__*/_interopDefaultCompat(MagicString);

const excludeRE = [
  // imported/exported from other module
  /\b(import|export)\b([\w$*{},\s]+?)\bfrom\s*["']/g,
  // defined as function
  /\bfunction\s*([\w$]+)\s*\(/g,
  // defined as class
  /\bclass\s*([\w$]+)\s*\{/g,
  // defined as local variable
  // eslint-disable-next-line regexp/no-super-linear-backtracking
  /\b(?:const|let|var)\s+?(\[.*?\]|\{.*?\}|.+?)\s*?[=;\n]/gs
];
const importAsRE = /^.*\sas\s+/;
const separatorRE = /[,[\]{}\n]|\b(?:import|export)\b/g;
const matchRE = /(^|\.\.\.|(?:\bcase|\?)\s+|[^\w$/)]|\bextends\s+)([\w$]+)\s*(?=[.()[\]}:;?+\-*&|`<>,\n]|\b(?:instanceof|in)\b|$|(?<=extends\s+\w+)\s+\{)/g;
const regexRE = /\/\S*?(?<!\\)(?<!\[[^\]]*)\/[gimsuy]*/g;
function stripCommentsAndStrings(code, options) {
  return stripLiteral.stripLiteral(code, options).replace(regexRE, 'new RegExp("")');
}

function defineUnimportPreset(preset) {
  return preset;
}
const safePropertyName = /^[a-z$_][\w$]*$/i;
function stringifyWith(withValues) {
  let withDefs = "";
  for (let entries = Object.entries(withValues), l = entries.length, i = 0; i < l; i++) {
    const [prop, value] = entries[i];
    withDefs += safePropertyName.test(prop) ? prop : JSON.stringify(prop);
    withDefs += `: ${JSON.stringify(String(value))}`;
    if (i + 1 !== l)
      withDefs += ", ";
  }
  return `{ ${withDefs} }`;
}
function stringifyImports(imports, isCJS = false) {
  const map = toImportModuleMap(imports);
  return Object.entries(map).flatMap(([name, importSet]) => {
    const entries = [];
    const imports2 = Array.from(importSet).filter((i) => {
      if (!i.name || i.as === "") {
        let importStr;
        if (isCJS) {
          importStr = `require('${name}');`;
        } else {
          importStr = `import '${name}'`;
          if (i.with)
            importStr += ` with ${stringifyWith(i.with)}`;
          importStr += ";";
        }
        entries.push(importStr);
        return false;
      } else if (i.name === "default" || i.name === "=") {
        let importStr;
        if (isCJS) {
          importStr = i.name === "=" ? `const ${i.as} = require('${name}');` : `const { default: ${i.as} } = require('${name}');`;
        } else {
          importStr = `import ${i.as} from '${name}'`;
          if (i.with)
            importStr += ` with ${stringifyWith(i.with)}`;
          importStr += ";";
        }
        entries.push(importStr);
        return false;
      } else if (i.name === "*") {
        let importStr;
        if (isCJS) {
          importStr = `const ${i.as} = require('${name}');`;
        } else {
          importStr = `import * as ${i.as} from '${name}'`;
          if (i.with)
            importStr += ` with ${stringifyWith(i.with)}`;
          importStr += ";";
        }
        entries.push(importStr);
        return false;
      } else if (!isCJS && i.with) {
        entries.push(`import { ${stringifyImportAlias(i)} } from '${name}' with ${stringifyWith(i.with)};`);
        return false;
      }
      return true;
    });
    if (imports2.length) {
      const importsAs = imports2.map((i) => stringifyImportAlias(i, isCJS));
      entries.push(
        isCJS ? `const { ${importsAs.join(", ")} } = require('${name}');` : `import { ${importsAs.join(", ")} } from '${name}';`
      );
    }
    return entries;
  }).join("\n");
}
function dedupeImports(imports, warn) {
  const map = /* @__PURE__ */ new Map();
  const indexToRemove = /* @__PURE__ */ new Set();
  imports.filter((i) => !i.disabled).forEach((i, idx) => {
    if (i.declarationType === "enum")
      return;
    const name = i.as ?? i.name;
    if (!map.has(name)) {
      map.set(name, idx);
      return;
    }
    const other = imports[map.get(name)];
    if (other.from === i.from) {
      indexToRemove.add(idx);
      return;
    }
    const diff = (other.priority || 1) - (i.priority || 1);
    if (diff === 0)
      warn(`Duplicated imports "${name}", the one from "${other.from}" has been ignored and "${i.from}" is used`);
    if (diff <= 0) {
      indexToRemove.add(map.get(name));
      map.set(name, idx);
    } else {
      indexToRemove.add(idx);
    }
  });
  return imports.filter((_, idx) => !indexToRemove.has(idx));
}
function toExports(imports, fileDir, includeType = false) {
  const map = toImportModuleMap(imports, includeType);
  return Object.entries(map).flatMap(([name, imports2]) => {
    if (isFilePath(name))
      name = name.replace(/\.[a-z]+$/i, "");
    if (fileDir && pathe.isAbsolute(name)) {
      name = pathe.relative(fileDir, name);
      if (!name.match(/^[./]/))
        name = `./${name}`;
    }
    const entries = [];
    const filtered = Array.from(imports2).filter((i) => {
      if (i.name === "*") {
        entries.push(`export * as ${i.as} from '${name}';`);
        return false;
      }
      return true;
    });
    if (filtered.length)
      entries.push(`export { ${filtered.map((i) => stringifyImportAlias(i, false)).join(", ")} } from '${name}';`);
    return entries;
  }).join("\n");
}
function stripFileExtension(path) {
  return path.replace(/\.[a-z]+$/i, "");
}
function toTypeDeclarationItems(imports, options) {
  return imports.map((i) => {
    const from = options?.resolvePath?.(i) || stripFileExtension(i.typeFrom || i.from);
    let typeDef = "";
    if (i.with)
      typeDef += `import('${from}', { with: ${stringifyWith(i.with)} })`;
    else
      typeDef += `import('${from}')`;
    if (i.name !== "*" && i.name !== "=")
      typeDef += `['${i.name}']`;
    return `const ${i.as}: typeof ${typeDef}`;
  }).sort();
}
function toTypeDeclarationFile(imports, options) {
  const items = toTypeDeclarationItems(imports, options);
  const {
    exportHelper = true
  } = options || {};
  let declaration = "";
  if (exportHelper)
    declaration += "export {}\n";
  declaration += `declare global {
${items.map((i) => `  ${i}`).join("\n")}
}`;
  return declaration;
}
function makeTypeModulesMap(imports, resolvePath2) {
  const modulesMap = /* @__PURE__ */ new Map();
  const resolveImportFrom = typeof resolvePath2 === "function" ? (i) => {
    return resolvePath2(i) || stripFileExtension(i.typeFrom || i.from);
  } : (i) => stripFileExtension(i.typeFrom || i.from);
  for (const import_ of imports) {
    const from = resolveImportFrom(import_);
    let module = modulesMap.get(from);
    if (!module) {
      module = { typeImports: /* @__PURE__ */ new Set(), starTypeImport: void 0 };
      modulesMap.set(from, module);
    }
    if (import_.name === "*") {
      if (import_.as)
        module.starTypeImport = import_;
    } else {
      module.typeImports.add(import_);
    }
  }
  return modulesMap;
}
function toTypeReExports(imports, options) {
  const importsMap = makeTypeModulesMap(imports, options?.resolvePath);
  const code = Array.from(importsMap).flatMap(([from, module]) => {
    const { starTypeImport, typeImports } = module;
    const strings = [];
    if (typeImports.size) {
      const typeImportNames = Array.from(typeImports).map(({ name, as }) => {
        if (as && as !== name)
          return `${name} as ${as}`;
        return name;
      });
      strings.push(
        "// @ts-ignore",
        `export type { ${typeImportNames.join(", ")} } from '${from}'`
      );
    }
    if (starTypeImport) {
      strings.push(
        "// @ts-ignore",
        `export type * as ${starTypeImport.as} from '${from}'`
      );
    }
    if (strings.length) {
      strings.push(
        // This is a workaround for a TypeScript issue where type-only re-exports are not properly initialized.
        `import('${from}')`
      );
    }
    return strings;
  });
  return `// for type re-export
declare global {
${code.map((i) => `  ${i}`).join("\n")}
}`;
}
function stringifyImportAlias(item, isCJS = false) {
  return item.as === void 0 || item.name === item.as ? item.name : isCJS ? `${item.name}: ${item.as}` : `${item.name} as ${item.as}`;
}
function toImportModuleMap(imports, includeType = false) {
  const map = {};
  for (const _import of imports) {
    if (_import.type && !includeType)
      continue;
    if (!map[_import.from])
      map[_import.from] = /* @__PURE__ */ new Set();
    map[_import.from].add(_import);
  }
  return map;
}
function getString(code) {
  if (typeof code === "string")
    return code;
  return code.toString();
}
function getMagicString(code) {
  if (typeof code === "string")
    return new MagicString__default(code);
  return code;
}
function addImportToCode(code, imports, isCJS = false, mergeExisting = false, injectAtLast = false, firstOccurrence = Number.POSITIVE_INFINITY, onResolved, onStringified) {
  let newImports = [];
  const s = getMagicString(code);
  let _staticImports;
  const strippedCode = stripCommentsAndStrings(s.original);
  function findStaticImportsLazy() {
    if (!_staticImports) {
      _staticImports = mlly.findStaticImports(s.original).filter((i) => Boolean(strippedCode.slice(i.start, i.end).trim())).map((i) => mlly.parseStaticImport(i));
    }
    return _staticImports;
  }
  function hasShebang() {
    const shebangRegex = /^#!.+/;
    return shebangRegex.test(s.original);
  }
  if (mergeExisting && !isCJS) {
    const existingImports = findStaticImportsLazy();
    const map = /* @__PURE__ */ new Map();
    imports.forEach((i) => {
      const target = existingImports.find((e) => e.specifier === i.from && e.imports.startsWith("{"));
      if (!target)
        return newImports.push(i);
      if (!map.has(target))
        map.set(target, []);
      map.get(target).push(i);
    });
    for (const [target, items] of map.entries()) {
      const strings = items.map((i) => `${stringifyImportAlias(i)}, `);
      const importLength = target.code.match(/^\s*import\s*\{/)?.[0]?.length;
      if (importLength)
        s.appendLeft(target.start + importLength, ` ${strings.join("").trim()}`);
    }
  } else {
    newImports = imports;
  }
  newImports = onResolved?.(newImports) ?? newImports;
  let newEntries = stringifyImports(newImports, isCJS);
  newEntries = onStringified?.(newEntries, newImports) ?? newEntries;
  if (newEntries) {
    const insertionIndex = injectAtLast ? findStaticImportsLazy().reverse().find((i) => i.end <= firstOccurrence)?.end ?? 0 : 0;
    if (insertionIndex > 0)
      s.appendRight(insertionIndex, `
${newEntries}
`);
    else if (hasShebang())
      s.appendLeft(s.original.indexOf("\n") + 1, `
${newEntries}
`);
    else
      s.prepend(`${newEntries}
`);
  }
  return {
    s,
    get code() {
      return s.toString();
    }
  };
}
function normalizeImports(imports) {
  for (const _import of imports)
    _import.as = _import.as ?? _import.name;
  return imports;
}
function resolveIdAbsolute(id, parentId) {
  return mlly.resolvePath(id, {
    url: parentId
  });
}
function isFilePath(path) {
  return path.startsWith(".") || pathe.isAbsolute(path) || path.includes("://");
}
const toImports = stringifyImports;

const contextRE$1 = /\b_ctx\.([$\w]+)\b/g;
const UNREF_KEY = "__unimport_unref_";
const VUE_TEMPLATE_NAME = "unimport:vue-template";
function vueTemplateAddon() {
  const self = {
    name: VUE_TEMPLATE_NAME,
    async transform(s, id) {
      if (!s.original.includes("_ctx.") || s.original.includes(UNREF_KEY))
        return s;
      const matches = Array.from(s.original.matchAll(contextRE$1));
      const imports = await this.getImports();
      let targets = [];
      for (const match of matches) {
        const name = match[1];
        const item = imports.find((i) => i.as === name);
        if (!item)
          continue;
        const start = match.index;
        const end = start + match[0].length;
        const tempName = `__unimport_${name}`;
        s.overwrite(start, end, `(${JSON.stringify(name)} in _ctx ? _ctx.${name} : ${UNREF_KEY}(${tempName}))`);
        if (!targets.find((i) => i.as === tempName)) {
          targets.push({
            ...item,
            as: tempName
          });
        }
      }
      if (targets.length) {
        targets.push({
          name: "unref",
          from: "vue",
          as: UNREF_KEY
        });
        for (const addon of this.addons) {
          if (addon === self)
            continue;
          targets = await addon.injectImportsResolved?.call(this, targets, s, id) ?? targets;
        }
        let injection = stringifyImports(targets);
        for (const addon of this.addons) {
          if (addon === self)
            continue;
          injection = await addon.injectImportsStringified?.call(this, injection, targets, s, id) ?? injection;
        }
        s.prepend(injection);
      }
      return s;
    },
    async declaration(dts, options) {
      const imports = await this.getImports();
      const items = imports.map((i) => {
        if (i.type || i.dtsDisabled)
          return "";
        const from = options?.resolvePath?.(i) || i.from;
        return `readonly ${i.as}: UnwrapRef<typeof import('${from}')${i.name !== "*" ? `['${i.name}']` : ""}>`;
      }).filter(Boolean).sort();
      const extendItems = items.map((i) => `    ${i}`).join("\n");
      return `${dts}
// for vue template auto import
import { UnwrapRef } from 'vue'
declare module 'vue' {
  interface ComponentCustomProperties {
${extendItems}
  }
}`;
    }
  };
  return self;
}

const contextRE = /resolveDirective as _resolveDirective/;
const contextText = `${contextRE.source}, `;
const directiveRE = /(?:var|const) (\w+) = _resolveDirective\("([\w.-]+)"\);?\s*/g;
const VUE_DIRECTIVES_NAME = "unimport:vue-directives";
function vueDirectivesAddon(options = {}) {
  function isDirective(importEntry) {
    let isDirective2 = importEntry.meta?.vueDirective === true;
    if (isDirective2) {
      return true;
    }
    isDirective2 = options.isDirective?.(normalizePath(process__default.cwd(), importEntry.from), importEntry) ?? false;
    if (isDirective2) {
      importEntry.meta ?? (importEntry.meta = {});
      importEntry.meta.vueDirective = true;
    }
    return isDirective2;
  }
  const self = {
    name: VUE_DIRECTIVES_NAME,
    async transform(s, id) {
      if (!s.original.match(contextRE))
        return s;
      const matches = Array.from(s.original.matchAll(directiveRE)).sort((a, b) => b.index - a.index);
      if (!matches.length)
        return s;
      let targets = [];
      for await (const [
        begin,
        end,
        importEntry
      ] of findDirectives(
        isDirective,
        matches,
        this.getImports()
      )) {
        s.overwrite(begin, end, "");
        targets.push(importEntry);
      }
      if (!targets.length)
        return s;
      s.replace(contextText, "");
      for (const addon of this.addons) {
        if (addon === self)
          continue;
        targets = await addon.injectImportsResolved?.call(this, targets, s, id) ?? targets;
      }
      let injection = stringifyImports(targets);
      for (const addon of this.addons) {
        if (addon === self)
          continue;
        injection = await addon.injectImportsStringified?.call(this, injection, targets, s, id) ?? injection;
      }
      s.prepend(injection);
      return s;
    },
    async declaration(dts, options2) {
      const directivesMap = await this.getImports().then((imports) => {
        return imports.filter(isDirective).reduce((acc, i) => {
          if (i.type || i.dtsDisabled)
            return acc;
          let name;
          if (i.name === "default" && (i.as === "default" || !i.as)) {
            const file = path.basename(i.from);
            const idx = file.indexOf(".");
            name = idx > -1 ? file.slice(0, idx) : file;
          } else {
            name = i.as ?? i.name;
          }
          name = name[0] === "v" ? scule.camelCase(name) : scule.camelCase(`v-${name}`);
          if (!acc.has(name)) {
            acc.set(name, i);
          }
          return acc;
        }, /* @__PURE__ */ new Map());
      });
      if (!directivesMap.size)
        return dts;
      const directives = Array.from(directivesMap.entries()).map(([name, i]) => `    ${name}: typeof import('${options2?.resolvePath?.(i) || i.from}')['${i.name}']`).sort().join("\n");
      return `${dts}
// for vue directives auto import
declare module 'vue' {
  interface ComponentCustomProperties {
${directives}
  }
  interface GlobalDirectives {
${directives}
  }
}`;
    }
  };
  return self;
}
function resolvePath(cwd, path) {
  return path[0] === "." ? pathe.resolve(cwd, path) : path;
}
function normalizePath(cwd, path) {
  return resolvePath(cwd, path).replace(/\\/g, "/");
}
async function* findDirectives(isDirective, regexArray, importsPromise) {
  const imports = (await importsPromise).filter(isDirective);
  if (!imports.length)
    return;
  const symbols = regexArray.reduce((acc, regex) => {
    const [all, symbol, resolveDirectiveName] = regex;
    if (acc.has(symbol))
      return acc;
    acc.set(symbol, [
      regex.index,
      regex.index + all.length,
      scule.kebabCase(resolveDirectiveName)
    ]);
    return acc;
  }, /* @__PURE__ */ new Map());
  for (const [symbol, data] of symbols.entries()) {
    yield* findDirective(imports, symbol, data);
  }
}
function* findDirective(imports, symbol, [begin, end, importName]) {
  let resolvedName;
  for (const i of imports) {
    if (i.name === "default" && (i.as === "default" || !i.as)) {
      const file = path.basename(i.from);
      const idx = file.indexOf(".");
      resolvedName = scule.kebabCase(idx > -1 ? file.slice(0, idx) : file);
    } else {
      resolvedName = scule.kebabCase(i.as ?? i.name);
    }
    if (resolvedName[0] === "v") {
      resolvedName = resolvedName.slice(resolvedName[1] === "-" ? 2 : 1);
    }
    if (resolvedName === importName) {
      yield [
        begin,
        end,
        { ...i, name: i.name, as: symbol }
      ];
      return;
    }
  }
}

exports.VUE_DIRECTIVES_NAME = VUE_DIRECTIVES_NAME;
exports.VUE_TEMPLATE_NAME = VUE_TEMPLATE_NAME;
exports.addImportToCode = addImportToCode;
exports.dedupeImports = dedupeImports;
exports.defineUnimportPreset = defineUnimportPreset;
exports.excludeRE = excludeRE;
exports.getMagicString = getMagicString;
exports.getString = getString;
exports.importAsRE = importAsRE;
exports.matchRE = matchRE;
exports.normalizeImports = normalizeImports;
exports.resolveIdAbsolute = resolveIdAbsolute;
exports.separatorRE = separatorRE;
exports.stringifyImports = stringifyImports;
exports.stripCommentsAndStrings = stripCommentsAndStrings;
exports.stripFileExtension = stripFileExtension;
exports.toExports = toExports;
exports.toImports = toImports;
exports.toTypeDeclarationFile = toTypeDeclarationFile;
exports.toTypeDeclarationItems = toTypeDeclarationItems;
exports.toTypeReExports = toTypeReExports;
exports.vueDirectivesAddon = vueDirectivesAddon;
exports.vueTemplateAddon = vueTemplateAddon;
