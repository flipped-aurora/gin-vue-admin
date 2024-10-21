'use strict';

const addons = require('./unimport.8977d7da.cjs');
const mlly = require('mlly');
const node_fs = require('node:fs');
const promises = require('node:fs/promises');
const path = require('node:path');
const process = require('node:process');
const node_url = require('node:url');
const fg = require('fast-glob');
const pathe = require('pathe');
const scule = require('scule');
const os = require('node:os');
const pkgTypes = require('pkg-types');
const localPkg = require('local-pkg');

function _interopDefaultCompat (e) { return e && typeof e === 'object' && 'default' in e ? e.default : e; }

const path__default = /*#__PURE__*/_interopDefaultCompat(path);
const process__default = /*#__PURE__*/_interopDefaultCompat(process);
const fg__default = /*#__PURE__*/_interopDefaultCompat(fg);
const os__default = /*#__PURE__*/_interopDefaultCompat(os);

const version = "3.13.1";

function configureAddons(opts) {
  const addons$1 = [];
  if (Array.isArray(opts.addons)) {
    addons$1.push(...opts.addons);
  } else {
    const addonsMap = /* @__PURE__ */ new Map();
    if (opts.addons?.addons?.length) {
      let i = 0;
      for (const addon of opts.addons.addons) {
        addonsMap.set(addon.name || `external:custom-${i++}`, addon);
      }
    }
    if (opts.addons?.vueTemplate) {
      if (!addonsMap.has(addons.VUE_TEMPLATE_NAME)) {
        addonsMap.set(addons.VUE_TEMPLATE_NAME, addons.vueTemplateAddon());
      }
    }
    if (opts.addons?.vueDirectives) {
      if (!addonsMap.has(addons.VUE_DIRECTIVES_NAME)) {
        addonsMap.set(addons.VUE_DIRECTIVES_NAME, addons.vueDirectivesAddon(
          typeof opts.addons.vueDirectives === "object" ? opts.addons.vueDirectives : void 0
        ));
      }
    }
    addons$1.push(...addonsMap.values());
  }
  return addons$1;
}

async function detectImportsRegex(code, ctx, options) {
  const s = addons.getMagicString(code);
  const original = s.original;
  const strippedCode = addons.stripCommentsAndStrings(
    original,
    // Do not strip comments if they are virtual import names
    options?.transformVirtualImports !== false && ctx.options.virtualImports?.length ? {
      filter: (i) => !ctx.options.virtualImports.includes(i),
      fillChar: "-"
    } : void 0
  );
  const syntax = mlly.detectSyntax(strippedCode);
  const isCJSContext = syntax.hasCJS && !syntax.hasESM;
  let matchedImports = [];
  const occurrenceMap = /* @__PURE__ */ new Map();
  const map = await ctx.getImportMap();
  if (options?.autoImport !== false) {
    Array.from(strippedCode.matchAll(addons.matchRE)).forEach((i) => {
      if (i[1] === ".")
        return null;
      const end = strippedCode[i.index + i[0].length];
      const before = strippedCode[i.index - 1];
      if (end === ":" && !["?", "case"].includes(i[1].trim()) && before !== ":")
        return null;
      const name = i[2];
      const occurrence = i.index + i[1].length;
      if (occurrenceMap.get(name) || Number.POSITIVE_INFINITY > occurrence)
        occurrenceMap.set(name, occurrence);
    });
    for (const regex of addons.excludeRE) {
      for (const match of strippedCode.matchAll(regex)) {
        const segments = [...match[1]?.split(addons.separatorRE) || [], ...match[2]?.split(addons.separatorRE) || []];
        for (const segment of segments) {
          const identifier = segment.replace(addons.importAsRE, "").trim();
          occurrenceMap.delete(identifier);
        }
      }
    }
    const identifiers = new Set(occurrenceMap.keys());
    matchedImports = Array.from(identifiers).map((name) => {
      const item = map.get(name);
      if (item && !item.disabled)
        return item;
      occurrenceMap.delete(name);
      return null;
    }).filter(Boolean);
    for (const addon of ctx.addons)
      matchedImports = await addon.matchImports?.call(ctx, identifiers, matchedImports) || matchedImports;
  }
  if (options?.transformVirtualImports !== false && ctx.options.virtualImports?.length) {
    const virtualImports = parseVirtualImportsRegex(strippedCode, map, ctx.options.virtualImports);
    virtualImports.ranges.forEach(([start, end]) => {
      s.remove(start, end);
    });
    matchedImports.push(...virtualImports.imports);
  }
  const firstOccurrence = Math.min(...Array.from(occurrenceMap.entries()).map((i) => i[1]));
  return {
    s,
    strippedCode,
    isCJSContext,
    matchedImports,
    firstOccurrence
  };
}
function parseVirtualImportsRegex(strippedCode, importMap, virtualImports) {
  const imports = [];
  const ranges = [];
  if (virtualImports?.length) {
    mlly.findStaticImports(strippedCode).filter((i) => virtualImports.includes(i.specifier)).map((i) => mlly.parseStaticImport(i)).forEach((i) => {
      ranges.push([i.start, i.end]);
      Object.entries(i.namedImports || {}).forEach(([name, as]) => {
        const original = importMap.get(name);
        if (!original)
          throw new Error(`[unimport] failed to find "${name}" imported from "${i.specifier}"`);
        imports.push({
          from: original.from,
          name: original.name,
          as
        });
      });
    });
  }
  return {
    imports,
    ranges
  };
}

async function detectImports(code, ctx, options) {
  if (options?.parser === "acorn")
    return import('../chunks/detect-acorn.cjs').then((r) => r.detectImportsAcorn(code, ctx, options));
  return detectImportsRegex(code, ctx, options);
}

async function scanFilesFromDir(dir, options) {
  const dirs = (Array.isArray(dir) ? dir : [dir]).map((d) => pathe.normalize(d));
  const fileFilter = options?.fileFilter || (() => true);
  const filePatterns = options?.filePatterns || ["*.{ts,js,mjs,cjs,mts,cts}"];
  const result = await Promise.all(
    // Do multiple glob searches to persist the order of input dirs
    dirs.map(
      async (i) => await fg__default(
        [i, ...filePatterns.map((p) => pathe.join(i, p))],
        {
          absolute: true,
          cwd: options?.cwd || process__default.cwd(),
          onlyFiles: true,
          followSymbolicLinks: true
        }
      ).then(
        (r) => r.map((f) => pathe.normalize(f)).sort()
      )
    )
  );
  return Array.from(new Set(result.flat())).filter(fileFilter);
}
async function scanDirExports(dir, options) {
  const files = await scanFilesFromDir(dir, options);
  const includeTypes = options?.types ?? true;
  const fileExports = await Promise.all(files.map((i) => scanExports(i, includeTypes)));
  const exports = fileExports.flat();
  const deduped = dedupeDtsExports(exports);
  return deduped;
}
const FileExtensionLookup = [
  ".mts",
  ".cts",
  ".ts",
  ".mjs",
  ".cjs",
  ".js"
];
function dedupeDtsExports(exports) {
  return exports.filter((i) => {
    if (!i.type)
      return true;
    if (i.declarationType === "enum")
      return true;
    return !exports.find((e) => e.as === i.as && e.name === i.name && !e.type);
  });
}
async function scanExports(filepath, includeTypes, seen = /* @__PURE__ */ new Set()) {
  if (seen.has(filepath)) {
    console.warn(`[unimport] "${filepath}" is already scanned, skipping`);
    return [];
  }
  seen.add(filepath);
  const imports = [];
  const code = await promises.readFile(filepath, "utf-8");
  const exports = mlly.findExports(code);
  const defaultExport = exports.find((i) => i.type === "default");
  if (defaultExport) {
    let name = pathe.parse(filepath).name;
    if (name === "index")
      name = pathe.parse(filepath.split("/").slice(0, -1).join("/")).name;
    const as = /[-_.]/.test(name) ? scule.camelCase(name) : name;
    imports.push({ name: "default", as, from: filepath });
  }
  async function toImport(exports2, additional) {
    for (const exp of exports2) {
      if (exp.type === "named") {
        for (const name of exp.names)
          imports.push({ name, as: name, from: filepath, ...additional });
      } else if (exp.type === "declaration") {
        if (exp.name) {
          imports.push({ name: exp.name, as: exp.name, from: filepath, ...additional });
          if (exp.declarationType === "enum")
            imports.push({ name: exp.name, as: exp.name, from: filepath, type: true, declarationType: exp.declarationType, ...additional });
        }
      } else if (exp.type === "star" && exp.specifier) {
        if (exp.name) {
          imports.push({ name: exp.name, as: exp.name, from: filepath, ...additional });
        } else {
          const subfile = exp.specifier;
          let subfilepath = pathe.resolve(pathe.dirname(filepath), subfile);
          let subfilepathResolved = false;
          for (const ext of FileExtensionLookup) {
            if (node_fs.existsSync(`${subfilepath}${ext}`)) {
              subfilepath = `${subfilepath}${ext}`;
              break;
            } else if (node_fs.existsSync(`${subfilepath}/index${ext}`)) {
              subfilepath = `${subfilepath}/index${ext}`;
              break;
            }
          }
          if (node_fs.existsSync(subfilepath)) {
            subfilepathResolved = true;
          } else {
            try {
              subfilepath = await mlly.resolve(exp.specifier);
              subfilepath = node_url.fileURLToPath(subfilepath).replaceAll(path__default.sep, "/");
              if (node_fs.existsSync(subfilepath)) {
                subfilepathResolved = true;
              }
            } catch {
            }
          }
          if (!subfilepathResolved) {
            console.warn(`[unimport] failed to resolve "${subfilepath}", skip scanning`);
            continue;
          }
          const nested = await scanExports(subfilepath, includeTypes, seen);
          imports.push(...additional ? nested.map((i) => ({ ...i, ...additional })) : nested);
        }
      }
    }
  }
  const isDts = filepath.match(/\.d\.[mc]?ts$/);
  if (isDts) {
    if (includeTypes) {
      await toImport(exports, { type: true });
      await toImport(mlly.findTypeExports(code), { type: true });
    }
  } else {
    await toImport(exports);
    if (includeTypes)
      await toImport(mlly.findTypeExports(code), { type: true });
  }
  return imports;
}

const CACHE_PATH = /* @__PURE__ */ pathe.join(os__default.tmpdir(), "unimport");
let CACHE_WRITEABLE;
async function resolvePackagePreset(preset) {
  const scanned = await extractExports(preset.package, preset.url, preset.cache);
  const filtered = scanned.filter((name) => {
    for (const item of preset.ignore || []) {
      if (typeof item === "string" && item === name)
        return false;
      if (item instanceof RegExp && item.test(name))
        return false;
      if (typeof item === "function" && item(name) === false)
        return false;
    }
    return true;
  });
  return filtered.map((name) => ({
    from: preset.package,
    name
  }));
}
async function extractExports(name, url, cache = true) {
  const packageJsonPath = await pkgTypes.resolvePackageJSON(name, { url });
  const packageJson = await pkgTypes.readPackageJSON(packageJsonPath);
  const version = packageJson.version;
  const cachePath = pathe.join(CACHE_PATH, `${name}@${version}`, "exports.json");
  if (cache && CACHE_WRITEABLE === void 0) {
    try {
      CACHE_WRITEABLE = isWritable(CACHE_PATH);
    } catch {
      CACHE_WRITEABLE = false;
    }
  }
  const useCache = cache && version && CACHE_WRITEABLE;
  if (useCache && node_fs.existsSync(cachePath))
    return JSON.parse(await node_fs.promises.readFile(cachePath, "utf-8"));
  const scanned = await mlly.resolveModuleExportNames(name, { url });
  if (useCache) {
    await node_fs.promises.mkdir(pathe.dirname(cachePath), { recursive: true });
    await node_fs.promises.writeFile(cachePath, JSON.stringify(scanned), "utf-8");
  }
  return scanned;
}
function isWritable(filename) {
  try {
    node_fs.accessSync(filename, node_fs.constants.W_OK);
    return true;
  } catch {
    return false;
  }
}

const dateFns = addons.defineUnimportPreset({
  from: "date-fns",
  imports: [
    "add",
    "addBusinessDays",
    "addDays",
    "addHours",
    "addISOWeekYears",
    "addMilliseconds",
    "addMinutes",
    "addMonths",
    "addQuarters",
    "addSeconds",
    "addWeeks",
    "addYears",
    "areIntervalsOverlapping",
    "clamp",
    "closestIndexTo",
    "closestTo",
    "compareAsc",
    "compareDesc",
    "constants",
    "daysToWeeks",
    "differenceInBusinessDays",
    "differenceInCalendarDays",
    "differenceInCalendarISOWeekYears",
    "differenceInCalendarISOWeeks",
    "differenceInCalendarMonths",
    "differenceInCalendarQuarters",
    "differenceInCalendarWeeks",
    "differenceInCalendarYears",
    "differenceInDays",
    "differenceInHours",
    "differenceInISOWeekYears",
    "differenceInMilliseconds",
    "differenceInMinutes",
    "differenceInMonths",
    "differenceInQuarters",
    "differenceInSeconds",
    "differenceInWeeks",
    "differenceInYears",
    "eachDayOfInterval",
    "eachHourOfInterval",
    "eachMinuteOfInterval",
    "eachMonthOfInterval",
    "eachQuarterOfInterval",
    "eachWeekOfInterval",
    "eachWeekendOfInterval",
    "eachWeekendOfMonth",
    "eachWeekendOfYear",
    "eachYearOfInterval",
    "endOfDay",
    "endOfDecade",
    "endOfHour",
    "endOfISOWeek",
    "endOfISOWeekYear",
    "endOfMinute",
    "endOfMonth",
    "endOfQuarter",
    "endOfSecond",
    "endOfToday",
    "endOfTomorrow",
    "endOfWeek",
    "endOfYear",
    "endOfYesterday",
    "format",
    "formatDistance",
    "formatDistanceStrict",
    "formatDistanceToNow",
    "formatDistanceToNowStrict",
    "formatDuration",
    "formatISO",
    "formatISO9075",
    "formatISODuration",
    "formatRFC3339",
    "formatRFC7231",
    "formatRelative",
    "fromUnixTime",
    "getDate",
    "getDay",
    "getDayOfYear",
    "getDaysInMonth",
    "getDaysInYear",
    "getDecade",
    "getDefaultOptions",
    "getHours",
    "getISODay",
    "getISOWeek",
    "getISOWeekYear",
    "getISOWeeksInYear",
    "getMilliseconds",
    "getMinutes",
    "getMonth",
    "getOverlappingDaysInIntervals",
    "getQuarter",
    "getSeconds",
    "getTime",
    "getUnixTime",
    "getWeek",
    "getWeekOfMonth",
    "getWeekYear",
    "getWeeksInMonth",
    "getYear",
    "hoursToMilliseconds",
    "hoursToMinutes",
    "hoursToSeconds",
    "intervalToDuration",
    "intlFormat",
    "intlFormatDistance",
    "isAfter",
    "isBefore",
    "isDate",
    "isEqual",
    "isExists",
    "isFirstDayOfMonth",
    "isFriday",
    "isFuture",
    "isLastDayOfMonth",
    "isLeapYear",
    "isMatch",
    "isMonday",
    "isPast",
    "isSameDay",
    "isSameHour",
    "isSameISOWeek",
    "isSameISOWeekYear",
    "isSameMinute",
    "isSameMonth",
    "isSameQuarter",
    "isSameSecond",
    "isSameWeek",
    "isSameYear",
    "isSaturday",
    "isSunday",
    "isThisHour",
    "isThisISOWeek",
    "isThisMinute",
    "isThisMonth",
    "isThisQuarter",
    "isThisSecond",
    "isThisWeek",
    "isThisYear",
    "isThursday",
    "isToday",
    "isTomorrow",
    "isTuesday",
    "isValid",
    "isWednesday",
    "isWeekend",
    "isWithinInterval",
    "isYesterday",
    "lastDayOfDecade",
    "lastDayOfISOWeek",
    "lastDayOfISOWeekYear",
    "lastDayOfMonth",
    "lastDayOfQuarter",
    "lastDayOfWeek",
    "lastDayOfYear",
    "lightFormat",
    "locale",
    "max",
    "milliseconds",
    "millisecondsToHours",
    "millisecondsToMinutes",
    "millisecondsToSeconds",
    "min",
    "minutesToHours",
    "minutesToMilliseconds",
    "minutesToSeconds",
    "monthsToQuarters",
    "monthsToYears",
    "nextDay",
    "nextFriday",
    "nextMonday",
    "nextSaturday",
    "nextSunday",
    "nextThursday",
    "nextTuesday",
    "nextWednesday",
    "parse",
    "parseISO",
    "parseJSON",
    "previousDay",
    "previousFriday",
    "previousMonday",
    "previousSaturday",
    "previousSunday",
    "previousThursday",
    "previousTuesday",
    "previousWednesday",
    "quartersToMonths",
    "quartersToYears",
    "roundToNearestMinutes",
    "secondsToHours",
    "secondsToMilliseconds",
    "secondsToMinutes",
    "set",
    "setDate",
    "setDay",
    "setDayOfYear",
    "setDefaultOptions",
    "setHours",
    "setISODay",
    "setISOWeek",
    "setISOWeekYear",
    "setMilliseconds",
    "setMinutes",
    "setMonth",
    "setQuarter",
    "setSeconds",
    "setWeek",
    "setWeekYear",
    "setYear",
    "startOfDay",
    "startOfDecade",
    "startOfHour",
    "startOfISOWeek",
    "startOfISOWeekYear",
    "startOfMinute",
    "startOfMonth",
    "startOfQuarter",
    "startOfSecond",
    "startOfToday",
    "startOfTomorrow",
    "startOfWeek",
    "startOfWeekYear",
    "startOfYear",
    "startOfYesterday",
    "sub",
    "subBusinessDays",
    "subDays",
    "subHours",
    "subISOWeekYears",
    "subMilliseconds",
    "subMinutes",
    "subMonths",
    "subQuarters",
    "subSeconds",
    "subWeeks",
    "subYears",
    "toDate",
    "weeksToDays",
    "yearsToMonths",
    "yearsToQuarters"
  ]
});

const pinia = addons.defineUnimportPreset({
  from: "pinia",
  imports: [
    // https://pinia.esm.dev/api/modules/pinia.html#functions
    "acceptHMRUpdate",
    "createPinia",
    "defineStore",
    "getActivePinia",
    "mapActions",
    "mapGetters",
    "mapState",
    "mapStores",
    "mapWritableState",
    "setActivePinia",
    "setMapStoreSuffix",
    "storeToRefs"
  ]
});

const preact = addons.defineUnimportPreset({
  from: "preact",
  imports: [
    "useState",
    "useCallback",
    "useMemo",
    "useEffect",
    "useRef",
    "useContext",
    "useReducer"
  ]
});

const quasar = addons.defineUnimportPreset({
  from: "quasar",
  imports: [
    // https://quasar.dev/vue-composables
    "useQuasar",
    "useDialogPluginComponent",
    "useFormChild",
    "useMeta"
  ]
});

const react = addons.defineUnimportPreset({
  from: "react",
  imports: [
    "useState",
    "useCallback",
    "useMemo",
    "useEffect",
    "useRef",
    "useContext",
    "useReducer"
  ]
});

const ReactRouterHooks = [
  "useOutletContext",
  "useHref",
  "useInRouterContext",
  "useLocation",
  "useNavigationType",
  "useNavigate",
  "useOutlet",
  "useParams",
  "useResolvedPath",
  "useRoutes"
];
const reactRouter = addons.defineUnimportPreset({
  from: "react-router",
  imports: [
    ...ReactRouterHooks
  ]
});

const reactRouterDom = addons.defineUnimportPreset({
  from: "react-router-dom",
  imports: [
    ...ReactRouterHooks,
    // react-router-dom only hooks
    "useLinkClickHandler",
    "useSearchParams",
    // react-router-dom Component
    // call once in general
    // 'BrowserRouter',
    // 'HashRouter',
    // 'MemoryRouter',
    "Link",
    "NavLink",
    "Navigate",
    "Outlet",
    "Route",
    "Routes"
  ]
});

const rxjs = addons.defineUnimportPreset({
  from: "rxjs",
  imports: [
    "of",
    "from",
    "map",
    "tap",
    "filter",
    "forkJoin",
    "throwError",
    "catchError",
    "Observable",
    "mergeMap",
    "switchMap",
    "merge",
    "zip",
    "take",
    "takeUntil",
    "first",
    "lastValueFrom",
    "skip",
    "skipUntil",
    "distinct",
    "distinctUntilChanged",
    "throttle",
    "throttleTime",
    "retry",
    "retryWhen",
    "timeout",
    "delay",
    "debounce",
    "debounceTime",
    "find",
    "every"
  ]
});

const solidCore = addons.defineUnimportPreset({
  from: "solid-js",
  imports: [
    "createSignal",
    "createEffect",
    "createMemo",
    "createResource",
    "onMount",
    "onCleanup",
    "onError",
    "untrack",
    "batch",
    "on",
    "createRoot",
    "mergeProps",
    "splitProps",
    "useTransition",
    "observable",
    "mapArray",
    "indexArray",
    "createContext",
    "useContext",
    "children",
    "lazy",
    "createDeferred",
    "createRenderEffect",
    "createSelector",
    "For",
    "Show",
    "Switch",
    "Match",
    "Index",
    "ErrorBoundary",
    "Suspense",
    "SuspenseList"
  ]
});
const solidStore = addons.defineUnimportPreset({
  from: "solid-js/store",
  imports: [
    "createStore",
    "produce",
    "reconcile",
    "createMutable"
  ]
});
const solidWeb = addons.defineUnimportPreset({
  from: "solid-js/web",
  imports: [
    "Dynamic",
    "hydrate",
    "render",
    "renderToString",
    "renderToStringAsync",
    "renderToStream",
    "isServer",
    "Portal"
  ]
});
const solid = addons.defineUnimportPreset({
  from: "solid-js",
  imports: [
    solidCore,
    solidStore,
    solidWeb
  ]
});

const solidAppRouter = addons.defineUnimportPreset({
  from: "solid-app-router",
  imports: [
    "Link",
    "NavLink",
    "Navigate",
    "Outlet",
    "Route",
    "Router",
    "Routes",
    "_mergeSearchString",
    "createIntegration",
    "hashIntegration",
    "normalizeIntegration",
    "pathIntegration",
    "staticIntegration",
    "useHref",
    "useIsRouting",
    "useLocation",
    "useMatch",
    "useNavigate",
    "useParams",
    "useResolvedPath",
    "useRouteData",
    "useRoutes",
    "useSearchParams"
  ]
});

const svelteAnimate = addons.defineUnimportPreset({
  from: "svelte/animate",
  imports: [
    "flip"
  ]
});
const svelteEasing = addons.defineUnimportPreset({
  from: "svelte/easing",
  imports: [
    "back",
    "bounce",
    "circ",
    "cubic",
    "elastic",
    "expo",
    "quad",
    "quart",
    "quint",
    "sine"
  ].reduce((acc, e) => {
    acc.push(`${e}In`, `${e}Out`, `${e}InOut`);
    return acc;
  }, ["linear"])
});
const svelteStore = addons.defineUnimportPreset({
  from: "svelte/store",
  imports: [
    "writable",
    "readable",
    "derived",
    "get"
  ]
});
const svelteMotion = addons.defineUnimportPreset({
  from: "svelte/motion",
  imports: [
    "tweened",
    "spring"
  ]
});
const svelteTransition = addons.defineUnimportPreset({
  from: "svelte/transition",
  imports: [
    "fade",
    "blur",
    "fly",
    "slide",
    "scale",
    "draw",
    "crossfade"
  ]
});
const svelte = addons.defineUnimportPreset({
  from: "svelte",
  imports: [
    // lifecycle
    "onMount",
    "beforeUpdate",
    "afterUpdate",
    "onDestroy",
    // tick
    "tick",
    // context
    "setContext",
    "getContext",
    "hasContext",
    "getAllContexts",
    // event dispatcher
    "createEventDispatcher"
  ]
});

const uniApp = addons.defineUnimportPreset({
  from: "@dcloudio/uni-app",
  imports: [
    "onAddToFavorites",
    "onBackPress",
    "onError",
    "onHide",
    "onLaunch",
    "onLoad",
    "onNavigationBarButtonTap",
    "onNavigationBarSearchInputChanged",
    "onNavigationBarSearchInputClicked",
    "onNavigationBarSearchInputConfirmed",
    "onNavigationBarSearchInputFocusChanged",
    "onPageNotFound",
    "onPageScroll",
    "onPullDownRefresh",
    "onReachBottom",
    "onReady",
    "onResize",
    "onShareAppMessage",
    "onShareTimeline",
    "onShow",
    "onTabItemTap",
    "onThemeChange",
    "onUnhandledRejection",
    "onUnload"
  ]
});

const veeValidate = addons.defineUnimportPreset({
  from: "vee-validate",
  imports: [
    // https://vee-validate.logaretm.com/v4/guide/composition-api/api-review
    // https://github.com/logaretm/vee-validate/blob/main/packages/vee-validate/src/index.ts
    "validate",
    "defineRule",
    "configure",
    "useField",
    "useForm",
    "useFieldArray",
    "useResetForm",
    "useIsFieldDirty",
    "useIsFieldTouched",
    "useIsFieldValid",
    "useIsSubmitting",
    "useValidateField",
    "useIsFormDirty",
    "useIsFormTouched",
    "useIsFormValid",
    "useValidateForm",
    "useSubmitCount",
    "useFieldValue",
    "useFormValues",
    "useFormErrors",
    "useFieldError",
    "useSubmitForm",
    "FormContextKey",
    "FieldContextKey"
  ]
});

const vitepress = addons.defineUnimportPreset({
  from: "vitepress",
  imports: [
    // helper methods
    "useData",
    "useRoute",
    "useRouter",
    "withBase"
  ]
});

const vitest = addons.defineUnimportPreset({
  from: "vitest",
  imports: [
    // suite
    "suite",
    "test",
    "describe",
    "it",
    // chai
    "chai",
    "expect",
    "assert",
    // utils
    "vitest",
    "vi",
    // hooks
    "beforeAll",
    "afterAll",
    "beforeEach",
    "afterEach"
  ]
});

const CommonCompositionAPI = [
  // lifecycle
  "onActivated",
  "onBeforeMount",
  "onBeforeUnmount",
  "onBeforeUpdate",
  "onErrorCaptured",
  "onDeactivated",
  "onMounted",
  "onServerPrefetch",
  "onUnmounted",
  "onUpdated",
  // setup helpers
  "useAttrs",
  "useSlots",
  // reactivity,
  "computed",
  "customRef",
  "isReadonly",
  "isRef",
  "isProxy",
  "isReactive",
  "markRaw",
  "reactive",
  "readonly",
  "ref",
  "shallowReactive",
  "shallowReadonly",
  "shallowRef",
  "triggerRef",
  "toRaw",
  "toRef",
  "toRefs",
  "toValue",
  "unref",
  "watch",
  "watchEffect",
  "watchPostEffect",
  "watchSyncEffect",
  // component
  "defineComponent",
  "defineAsyncComponent",
  "getCurrentInstance",
  "h",
  "inject",
  "nextTick",
  "provide",
  "useCssModule",
  "createApp",
  // effect scope
  "effectScope",
  "EffectScope",
  "getCurrentScope",
  "onScopeDispose",
  // types
  ...[
    "Component",
    "ComponentPublicInstance",
    "ComputedRef",
    "DirectiveBinding",
    "ExtractDefaultPropTypes",
    "ExtractPropTypes",
    "ExtractPublicPropTypes",
    "InjectionKey",
    "PropType",
    "Ref",
    "MaybeRef",
    "MaybeRefOrGetter",
    "VNode",
    "WritableComputedRef"
  ].map((name) => ({ name, type: true }))
];
const vue = addons.defineUnimportPreset({
  from: "vue",
  imports: [
    ...CommonCompositionAPI,
    // vue3 only
    "onRenderTracked",
    "onRenderTriggered",
    "resolveComponent",
    "useCssVars",
    // vue3.4+
    "useModel",
    // vue3.5+
    "onWatcherCleanup",
    "useId",
    "useTemplateRef"
  ]
});

const vueCompositionApi = addons.defineUnimportPreset({
  from: "@vue/composition-api",
  imports: CommonCompositionAPI
});

const vueDemi = addons.defineUnimportPreset({
  from: "vue-demi",
  imports: CommonCompositionAPI
});

const vueI18n = addons.defineUnimportPreset({
  from: "vue-i18n",
  imports: [
    "useI18n"
  ]
});

const vueMacros = addons.defineUnimportPreset({
  from: "vue/macros",
  imports: [
    // https://vuejs.org/guide/extras/reactivity-transform.html#refs-vs-reactive-variables
    "$",
    "$$",
    "$ref",
    "$shallowRef",
    "$toRef",
    "$customRef",
    "$computed"
  ]
});

const vueRouter = addons.defineUnimportPreset({
  from: "vue-router",
  imports: [
    "useRouter",
    "useRoute",
    "useLink",
    "onBeforeRouteLeave",
    "onBeforeRouteUpdate"
  ]
});

const vueRouterComposables = addons.defineUnimportPreset({
  from: "vue-router/composables",
  imports: [
    "useRouter",
    "useRoute",
    "useLink",
    "onBeforeRouteLeave",
    "onBeforeRouteUpdate"
  ]
});

let _cache;
const vueuseCore = () => {
  const excluded = ["toRefs", "utils"];
  if (!_cache) {
    try {
      const corePath = localPkg.resolveModule("@vueuse/core") || process__default.cwd();
      const path = localPkg.resolveModule("@vueuse/core/indexes.json") || localPkg.resolveModule("@vueuse/metadata/index.json") || localPkg.resolveModule("@vueuse/metadata/index.json", { paths: [corePath] });
      const indexesJson = JSON.parse(node_fs.readFileSync(path, "utf-8"));
      _cache = addons.defineUnimportPreset({
        from: "@vueuse/core",
        imports: indexesJson.functions.filter((i) => ["core", "shared"].includes(i.package)).map((i) => i.name).filter((i) => i && i.length >= 4 && !excluded.includes(i))
      });
    } catch (error) {
      console.error(error);
      throw new Error("[auto-import] failed to load @vueuse/core, have you installed it?");
    }
  }
  return _cache;
};

const vueuseHead = addons.defineUnimportPreset({
  from: "@vueuse/head",
  imports: [
    "useHead"
  ]
});

const vuex = addons.defineUnimportPreset({
  from: "vuex",
  imports: [
    // https://next.vuex.vuejs.org/api/#createstore
    "createStore",
    // https://github.com/vuejs/vuex/blob/4.0/types/logger.d.ts#L20
    "createLogger",
    // https://next.vuex.vuejs.org/api/#component-binding-helpers
    "mapState",
    "mapGetters",
    "mapActions",
    "mapMutations",
    "createNamespacedHelpers",
    // https://next.vuex.vuejs.org/api/#composable-functions
    "useStore"
  ]
});

const builtinPresets = {
  "@vue/composition-api": vueCompositionApi,
  "@vueuse/core": vueuseCore,
  "@vueuse/head": vueuseHead,
  "pinia": pinia,
  "preact": preact,
  "quasar": quasar,
  "react": react,
  "react-router": reactRouter,
  "react-router-dom": reactRouterDom,
  "svelte": svelte,
  "svelte/animate": svelteAnimate,
  "svelte/easing": svelteEasing,
  "svelte/motion": svelteMotion,
  "svelte/store": svelteStore,
  "svelte/transition": svelteTransition,
  "vee-validate": veeValidate,
  "vitepress": vitepress,
  "vue-demi": vueDemi,
  "vue-i18n": vueI18n,
  "vue-router": vueRouter,
  "vue-router-composables": vueRouterComposables,
  "vue": vue,
  "vue/macros": vueMacros,
  "vuex": vuex,
  "vitest": vitest,
  "uni-app": uniApp,
  "solid-js": solid,
  "solid-app-router": solidAppRouter,
  "rxjs": rxjs,
  "date-fns": dateFns
};

const commonProps = [
  "from",
  "priority",
  "disabled",
  "dtsDisabled",
  "meta",
  "type"
];
async function resolvePreset(preset) {
  const imports = [];
  if ("package" in preset)
    return await resolvePackagePreset(preset);
  const common = {};
  commonProps.forEach((i) => {
    if (i in preset) {
      common[i] = preset[i];
    }
  });
  for (const _import of preset.imports) {
    if (typeof _import === "string")
      imports.push({ ...common, name: _import, as: _import });
    else if (Array.isArray(_import))
      imports.push({ ...common, name: _import[0], as: _import[1] || _import[0], from: _import[2] || preset.from });
    else if (_import.imports)
      imports.push(...await resolvePreset(_import));
    else
      imports.push({ ...common, ..._import });
  }
  return imports;
}
async function resolveBuiltinPresets(presets) {
  const resolved = await Promise.all(presets.map(async (p) => {
    let preset = typeof p === "string" ? builtinPresets[p] : p;
    if (typeof preset === "function")
      preset = preset();
    return await resolvePreset(preset);
  }));
  return resolved.flat();
}

function createUnimport(opts) {
  const ctx = createInternalContext(opts);
  async function generateTypeDeclarations(options) {
    const opts2 = {
      resolvePath: (i) => addons.stripFileExtension(i.typeFrom || i.from),
      ...options
    };
    const {
      typeReExports = true
    } = opts2;
    const imports = await ctx.getImports();
    let dts = addons.toTypeDeclarationFile(imports.filter((i) => !i.type && !i.dtsDisabled), opts2);
    const typeOnly = imports.filter((i) => i.type);
    if (typeReExports && typeOnly.length)
      dts += `
${addons.toTypeReExports(typeOnly, opts2)}`;
    for (const addon of ctx.addons)
      dts = await addon.declaration?.call(ctx, dts, opts2) ?? dts;
    return dts;
  }
  async function scanImportsFromFile(filepath, includeTypes = true) {
    const additions = await scanExports(filepath, includeTypes);
    await ctx.modifyDynamicImports((imports) => imports.filter((i) => i.from !== filepath).concat(additions));
    return additions;
  }
  async function scanImportsFromDir(dirs = ctx.options.dirs || [], options = ctx.options.dirsScanOptions) {
    const files = await scanFilesFromDir(dirs, options);
    const includeTypes = options?.types ?? true;
    const imports = (await Promise.all(files.map((dir) => scanExports(dir, includeTypes)))).flat();
    const deduped = dedupeDtsExports(imports);
    await ctx.modifyDynamicImports((imports2) => imports2.filter((i) => !files.includes(i.from)).concat(deduped));
    return imports;
  }
  async function injectImportsWithContext(code, id, options) {
    const result = await injectImports(code, id, ctx, {
      ...opts,
      ...options
    });
    const metadata = ctx.getMetadata();
    if (metadata) {
      result.imports.forEach((i) => {
        metadata.injectionUsage[i.name] = metadata.injectionUsage[i.name] || { import: i, count: 0, moduleIds: [] };
        metadata.injectionUsage[i.name].count++;
        if (id && !metadata.injectionUsage[i.name].moduleIds.includes(id))
          metadata.injectionUsage[i.name].moduleIds.push(id);
      });
    }
    return result;
  }
  async function init() {
    if (ctx.options.dirs?.length)
      await scanImportsFromDir();
  }
  return {
    version,
    init,
    clearDynamicImports: () => ctx.clearDynamicImports(),
    modifyDynamicImports: (fn) => ctx.modifyDynamicImports(fn),
    scanImportsFromDir,
    scanImportsFromFile,
    getImports: () => ctx.getImports(),
    getImportMap: () => ctx.getImportMap(),
    detectImports: (code) => detectImports(code, ctx),
    injectImports: injectImportsWithContext,
    generateTypeDeclarations: (options) => generateTypeDeclarations(options),
    getMetadata: () => ctx.getMetadata(),
    getInternalContext: () => ctx,
    // Deprecated
    toExports: async (filepath, includeTypes = false) => addons.toExports(await ctx.getImports(), filepath, includeTypes)
  };
}
function createInternalContext(opts) {
  let _combinedImports;
  const _map = /* @__PURE__ */ new Map();
  const addons$1 = configureAddons(opts);
  opts.addons = addons$1;
  opts.commentsDisable = opts.commentsDisable ?? ["@unimport-disable", "@imports-disable"];
  opts.commentsDebug = opts.commentsDebug ?? ["@unimport-debug", "@imports-debug"];
  let metadata;
  if (opts.collectMeta) {
    metadata = {
      injectionUsage: {}
    };
  }
  let resolvePromise;
  const ctx = {
    version,
    options: opts,
    addons: addons$1,
    staticImports: [...opts.imports || []].filter(Boolean),
    dynamicImports: [],
    modifyDynamicImports,
    clearDynamicImports,
    async getImports() {
      await resolvePromise;
      return updateImports();
    },
    async replaceImports(imports) {
      ctx.staticImports = [...imports || []].filter(Boolean);
      ctx.invalidate();
      await resolvePromise;
      return updateImports();
    },
    async getImportMap() {
      await ctx.getImports();
      return _map;
    },
    getMetadata() {
      return metadata;
    },
    invalidate() {
      _combinedImports = void 0;
    },
    resolveId: (id, parentId) => opts.resolveId?.(id, parentId)
  };
  resolvePromise = resolveBuiltinPresets(opts.presets || []).then((r) => {
    ctx.staticImports.unshift(...r);
    _combinedImports = void 0;
    updateImports();
  });
  function updateImports() {
    if (!_combinedImports) {
      let imports = addons.normalizeImports(addons.dedupeImports([...ctx.staticImports, ...ctx.dynamicImports], opts.warn || console.warn));
      for (const addon of ctx.addons) {
        if (addon.extendImports)
          imports = addon.extendImports.call(ctx, imports) ?? imports;
      }
      imports = imports.filter((i) => !i.disabled);
      _map.clear();
      for (const _import of imports) {
        if (!_import.type)
          _map.set(_import.as ?? _import.name, _import);
      }
      _combinedImports = imports;
    }
    return _combinedImports;
  }
  async function modifyDynamicImports(fn) {
    const result = await fn(ctx.dynamicImports);
    if (Array.isArray(result))
      ctx.dynamicImports = result;
    ctx.invalidate();
  }
  function clearDynamicImports() {
    ctx.dynamicImports.length = 0;
    ctx.invalidate();
  }
  return ctx;
}
async function injectImports(code, id, ctx, options) {
  const s = addons.getMagicString(code);
  if (ctx.options.commentsDisable?.some((c) => s.original.includes(c))) {
    return {
      s,
      get code() {
        return s.toString();
      },
      imports: []
    };
  }
  for (const addon of ctx.addons)
    await addon.transform?.call(ctx, s, id);
  const { isCJSContext, matchedImports, firstOccurrence } = await detectImports(s, ctx, options);
  const imports = await resolveImports(ctx, matchedImports, id);
  if (ctx.options.commentsDebug?.some((c) => s.original.includes(c))) {
    const log = ctx.options.debugLog || console.log;
    log(`[unimport] ${imports.length} imports detected in "${id}"${imports.length ? `: ${imports.map((i) => i.name).join(", ")}` : ""}`);
  }
  return {
    ...addons.addImportToCode(
      s,
      imports,
      isCJSContext,
      options?.mergeExisting,
      options?.injectAtEnd,
      firstOccurrence,
      (imports2) => {
        for (const addon of ctx.addons)
          imports2 = addon.injectImportsResolved?.call(ctx, imports2, s, id) ?? imports2;
        return imports2;
      },
      (str, imports2) => {
        for (const addon of ctx.addons)
          str = addon.injectImportsStringified?.call(ctx, str, imports2, s, id) ?? str;
        return str;
      }
    ),
    imports
  };
}
async function resolveImports(ctx, imports, id) {
  const resolveCache = /* @__PURE__ */ new Map();
  const _imports = await Promise.all(imports.map(async (i) => {
    if (!resolveCache.has(i.from))
      resolveCache.set(i.from, await ctx.resolveId(i.from, id) || i.from);
    const from = resolveCache.get(i.from);
    if (i.from === id || !from || from === "." || from === id)
      return;
    return {
      ...i,
      from
    };
  }));
  return _imports.filter(Boolean);
}

exports.builtinPresets = builtinPresets;
exports.createUnimport = createUnimport;
exports.dedupeDtsExports = dedupeDtsExports;
exports.resolveBuiltinPresets = resolveBuiltinPresets;
exports.resolvePreset = resolvePreset;
exports.scanDirExports = scanDirExports;
exports.scanExports = scanExports;
exports.scanFilesFromDir = scanFilesFromDir;
exports.version = version;
