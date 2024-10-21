"use strict";Object.defineProperty(exports, "__esModule", {value: true}); function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; } function _nullishCoalesce(lhs, rhsFn) { if (lhs != null) { return lhs; } else { return rhsFn(); } } function _optionalChain(ops) { let lastAccessLHS = undefined; let value = ops[0]; let i = 1; while (i < ops.length) { const op = ops[i]; const fn = ops[i + 1]; i += 2; if ((op === 'optionalAccess' || op === 'optionalCall') && value == null) { return undefined; } if (op === 'access' || op === 'optionalAccess') { lastAccessLHS = value; value = fn(value); } else if (op === 'call' || op === 'optionalCall') { value = fn((...args) => value.call(lastAccessLHS, ...args)); lastAccessLHS = undefined; } } return value; }






var _chunk6IWOTYMJcjs = require('./chunk-6IWOTYMJ.cjs');


var _chunkOBGZSXTJcjs = require('./chunk-OBGZSXTJ.cjs');

// src/core/resolvers/antdv.ts
var matchComponents = [
  {
    pattern: /^Affix/,
    styleDir: "affix"
  },
  {
    pattern: /^Avatar/,
    styleDir: "avatar"
  },
  {
    pattern: /^AutoComplete/,
    styleDir: "auto-complete"
  },
  {
    pattern: /^Alert/,
    styleDir: "alert"
  },
  {
    pattern: /^Anchor/,
    styleDir: "anchor"
  },
  {
    pattern: /^App/,
    styleDir: "app"
  },
  {
    pattern: /^Badge/,
    styleDir: "badge"
  },
  {
    pattern: /^Breadcrumb/,
    styleDir: "breadcrumb"
  },
  {
    pattern: /^Button/,
    styleDir: "button"
  },
  {
    pattern: /^Checkbox/,
    styleDir: "checkbox"
  },
  {
    pattern: /^Calendar/,
    styleDir: "calendar"
  },
  {
    pattern: /^Card/,
    styleDir: "card"
  },
  {
    pattern: /^Carousel/,
    styleDir: "carousel"
  },
  {
    pattern: /^Collapse/,
    styleDir: "collapse"
  },
  {
    pattern: /^Comment/,
    styleDir: "comment"
  },
  {
    pattern: /^Descriptions/,
    styleDir: "descriptions"
  },
  {
    pattern: /^RangePicker|^WeekPicker|^MonthPicker/,
    styleDir: "date-picker"
  },
  {
    pattern: /^Divider/,
    styleDir: "divider"
  },
  {
    pattern: /^Drawer/,
    styleDir: "drawer"
  },
  {
    pattern: /^Dropdown/,
    styleDir: "dropdown"
  },
  {
    pattern: /^Empty/,
    styleDir: "empty"
  },
  {
    pattern: /^Flex/,
    styleDir: "flex"
  },
  {
    pattern: /^FloatButton/,
    styleDir: "float-button"
  },
  {
    pattern: /^Form/,
    styleDir: "form"
  },
  {
    pattern: /^Grid/,
    styleDir: "grid"
  },
  {
    pattern: /^InputNumber/,
    styleDir: "input-number"
  },
  {
    pattern: /^Input|^Textarea/,
    styleDir: "input"
  },
  {
    pattern: /^Statistic/,
    styleDir: "statistic"
  },
  {
    pattern: /^CheckableTag/,
    styleDir: "tag"
  },
  {
    pattern: /^TimeRangePicker/,
    styleDir: "time-picker"
  },
  {
    pattern: /^Layout/,
    styleDir: "layout"
  },
  {
    pattern: /^Menu|^SubMenu/,
    styleDir: "menu"
  },
  {
    pattern: /^Table/,
    styleDir: "table"
  },
  {
    pattern: /^TimePicker|^TimeRangePicker/,
    styleDir: "time-picker"
  },
  {
    pattern: /^Radio/,
    styleDir: "radio"
  },
  {
    pattern: /^Image/,
    styleDir: "image"
  },
  {
    pattern: /^List/,
    styleDir: "list"
  },
  {
    pattern: /^Tab/,
    styleDir: "tabs"
  },
  {
    pattern: /^Mentions/,
    styleDir: "mentions"
  },
  {
    pattern: /^Step/,
    styleDir: "steps"
  },
  {
    pattern: /^Skeleton/,
    styleDir: "skeleton"
  },
  {
    pattern: /^Select/,
    styleDir: "select"
  },
  {
    pattern: /^TreeSelect/,
    styleDir: "tree-select"
  },
  {
    pattern: /^Tree|^DirectoryTree/,
    styleDir: "tree"
  },
  {
    pattern: /^Typography/,
    styleDir: "typography"
  },
  {
    pattern: /^Timeline/,
    styleDir: "timeline"
  },
  {
    pattern: /^Upload/,
    styleDir: "upload"
  },
  {
    pattern: /^Qrcode/,
    styleDir: "qrcode"
  },
  {
    pattern: /^Space/,
    styleDir: "space"
  }
];
function getStyleDir(compName) {
  let styleDir;
  const total = matchComponents.length;
  for (let i = 0; i < total; i++) {
    const matcher = matchComponents[i];
    if (compName.match(matcher.pattern)) {
      styleDir = matcher.styleDir;
      break;
    }
  }
  if (!styleDir)
    styleDir = _chunk6IWOTYMJcjs.kebabCase.call(void 0, compName);
  return styleDir;
}
function getSideEffects(compName, options) {
  const {
    importStyle = true,
    importLess = false
  } = options;
  if (!importStyle)
    return;
  const lib = options.cjs ? "lib" : "es";
  const packageName = _optionalChain([options, 'optionalAccess', _ => _.packageName]) || "ant-design-vue";
  if (importStyle === "less" || importStyle === "css-in-js" || importLess) {
    const styleDir = getStyleDir(compName);
    return `${packageName}/${lib}/${styleDir}/style`;
  } else {
    const styleDir = getStyleDir(compName);
    return `${packageName}/${lib}/${styleDir}/style/css`;
  }
}
var primitiveNames = ["Affix", "Anchor", "AnchorLink", "AutoComplete", "AutoCompleteOptGroup", "AutoCompleteOption", "Alert", "Avatar", "AvatarGroup", "BackTop", "Badge", "BadgeRibbon", "Breadcrumb", "BreadcrumbItem", "BreadcrumbSeparator", "Button", "ButtonGroup", "Calendar", "Card", "CardGrid", "CardMeta", "Collapse", "CollapsePanel", "Carousel", "Cascader", "Checkbox", "CheckboxGroup", "Col", "Comment", "ConfigProvider", "DatePicker", "MonthPicker", "WeekPicker", "RangePicker", "QuarterPicker", "Descriptions", "DescriptionsItem", "Divider", "Dropdown", "DropdownButton", "Drawer", "Empty", "Form", "FormItem", "FormItemRest", "Grid", "Input", "InputGroup", "InputPassword", "InputSearch", "Textarea", "Image", "ImagePreviewGroup", "InputNumber", "Layout", "LayoutHeader", "LayoutSider", "LayoutFooter", "LayoutContent", "List", "ListItem", "ListItemMeta", "Menu", "MenuDivider", "MenuItem", "MenuItemGroup", "SubMenu", "Mentions", "MentionsOption", "Modal", "Statistic", "StatisticCountdown", "PageHeader", "Pagination", "Popconfirm", "Popover", "Progress", "Radio", "RadioButton", "RadioGroup", "Rate", "Result", "Row", "Select", "SelectOptGroup", "SelectOption", "Skeleton", "SkeletonButton", "SkeletonAvatar", "SkeletonInput", "SkeletonImage", "Slider", "Space", "Spin", "Steps", "Step", "Switch", "Table", "TableColumn", "TableColumnGroup", "TableSummary", "TableSummaryRow", "TableSummaryCell", "Transfer", "Tree", "TreeNode", "DirectoryTree", "TreeSelect", "TreeSelectNode", "Tabs", "TabPane", "Tag", "CheckableTag", "TimePicker", "TimeRangePicker", "Timeline", "TimelineItem", "Tooltip", "Typography", "TypographyLink", "TypographyParagraph", "TypographyText", "TypographyTitle", "Upload", "UploadDragger", "LocaleProvider", "FloatButton", "FloatButtonGroup", "Qrcode", "Watermark", "Segmented", "Tour", "SpaceCompact", "StyleProvider", "Flex", "App"];
var antdvNames;
function genAntdNames(primitiveNames2) {
  antdvNames = new Set(primitiveNames2);
}
genAntdNames(primitiveNames);
function isAntdv(compName) {
  return antdvNames.has(compName);
}
function getImportName(compName) {
  if (compName === "Qrcode")
    return "QRCode";
  else if (compName === "SpaceCompact")
    return "Compact";
  return compName;
}
function AntDesignVueResolver(options = {}) {
  const originPrefix = _nullishCoalesce(options.prefix, () => ( "A"));
  return {
    type: "component",
    resolve: (name) => {
      if (options.resolveIcons && name.match(/(Outlined|Filled|TwoTone)$/)) {
        return {
          name,
          from: "@ant-design/icons-vue"
        };
      }
      const [compName, prefix] = [name.slice(originPrefix.length), name.slice(0, originPrefix.length)];
      if (prefix === originPrefix && isAntdv(compName) && !_optionalChain([options, 'optionalAccess', _2 => _2.exclude, 'optionalAccess', _3 => _3.includes, 'call', _4 => _4(compName)])) {
        const { cjs = false, packageName = "ant-design-vue" } = options;
        const path = `${packageName}/${cjs ? "lib" : "es"}`;
        return {
          name: getImportName(compName),
          from: path,
          sideEffects: getSideEffects(compName, options)
        };
      }
    }
  };
}

// node_modules/.pnpm/compare-versions@6.1.1/node_modules/compare-versions/lib/esm/utils.js
var semver = /^[v^~<>=]*?(\d+)(?:\.([x*]|\d+)(?:\.([x*]|\d+)(?:\.([x*]|\d+))?(?:-([\da-z\-]+(?:\.[\da-z\-]+)*))?(?:\+[\da-z\-]+(?:\.[\da-z\-]+)*)?)?)?$/i;
var validateAndParse = (version) => {
  if (typeof version !== "string") {
    throw new TypeError("Invalid argument expected string");
  }
  const match = version.match(semver);
  if (!match) {
    throw new Error(`Invalid argument not valid semver ('${version}' received)`);
  }
  match.shift();
  return match;
};
var isWildcard = (s) => s === "*" || s === "x" || s === "X";
var tryParse = (v) => {
  const n = parseInt(v, 10);
  return isNaN(n) ? v : n;
};
var forceType = (a, b) => typeof a !== typeof b ? [String(a), String(b)] : [a, b];
var compareStrings = (a, b) => {
  if (isWildcard(a) || isWildcard(b))
    return 0;
  const [ap, bp] = forceType(tryParse(a), tryParse(b));
  if (ap > bp)
    return 1;
  if (ap < bp)
    return -1;
  return 0;
};
var compareSegments = (a, b) => {
  for (let i = 0; i < Math.max(a.length, b.length); i++) {
    const r = compareStrings(a[i] || "0", b[i] || "0");
    if (r !== 0)
      return r;
  }
  return 0;
};

// node_modules/.pnpm/compare-versions@6.1.1/node_modules/compare-versions/lib/esm/compareVersions.js
var compareVersions = (v1, v2) => {
  const n1 = validateAndParse(v1);
  const n2 = validateAndParse(v2);
  const p1 = n1.pop();
  const p2 = n2.pop();
  const r = compareSegments(n1, n2);
  if (r !== 0)
    return r;
  if (p1 && p2) {
    return compareSegments(p1.split("."), p2.split("."));
  } else if (p1 || p2) {
    return p1 ? -1 : 1;
  }
  return 0;
};

// node_modules/.pnpm/compare-versions@6.1.1/node_modules/compare-versions/lib/esm/compare.js
var compare = (v1, v2, operator) => {
  assertValidOperator(operator);
  const res = compareVersions(v1, v2);
  return operatorResMap[operator].includes(res);
};
var operatorResMap = {
  ">": [1],
  ">=": [0, 1],
  "=": [0],
  "<=": [-1, 0],
  "<": [-1],
  "!=": [-1, 1]
};
var allowedOperators = Object.keys(operatorResMap);
var assertValidOperator = (op) => {
  if (typeof op !== "string") {
    throw new TypeError(`Invalid operator type, expected string but got ${typeof op}`);
  }
  if (allowedOperators.indexOf(op) === -1) {
    throw new Error(`Invalid operator, expected one of ${allowedOperators.join("|")}`);
  }
};

// src/core/resolvers/element-plus.ts
function getSideEffectsLegacy(partialName, options) {
  const { importStyle } = options;
  if (!importStyle)
    return;
  if (importStyle === "sass") {
    return [
      "element-plus/packages/theme-chalk/src/base.scss",
      `element-plus/packages/theme-chalk/src/${partialName}.scss`
    ];
  } else if (importStyle === true || importStyle === "css") {
    return [
      "element-plus/lib/theme-chalk/base.css",
      `element-plus/lib/theme-chalk/el-${partialName}.css`
    ];
  }
}
function getSideEffects2(dirName, options) {
  const { importStyle, ssr, nightly } = options;
  const themeFolder = nightly ? "@element-plus/nightly/theme-chalk" : "element-plus/theme-chalk";
  const esComponentsFolder = nightly ? "@element-plus/nightly/es/components" : "element-plus/es/components";
  if (importStyle === "sass") {
    return ssr ? [`${themeFolder}/src/base.scss`, `${themeFolder}/src/${dirName}.scss`] : [`${esComponentsFolder}/base/style/index`, `${esComponentsFolder}/${dirName}/style/index`];
  } else if (importStyle === true || importStyle === "css") {
    return ssr ? [`${themeFolder}/base.css`, `${themeFolder}/el-${dirName}.css`] : [`${esComponentsFolder}/base/style/css`, `${esComponentsFolder}/${dirName}/style/css`];
  }
}
function resolveComponent(name, options) {
  if (options.exclude && name.match(options.exclude))
    return;
  if (!name.match(/^El[A-Z]/))
    return;
  if (name.match(/^ElIcon.+/)) {
    return {
      name: name.replace(/^ElIcon/, ""),
      from: "@element-plus/icons-vue"
    };
  }
  const partialName = _chunk6IWOTYMJcjs.kebabCase.call(void 0, name.slice(2));
  const { version, ssr, nightly } = options;
  if (compare(version, "1.1.0-beta.1", ">=") || nightly) {
    return {
      name,
      from: `${nightly ? "@element-plus/nightly" : "element-plus"}/${ssr ? "lib" : "es"}`,
      sideEffects: getSideEffects2(partialName, options)
    };
  } else if (compare(version, "1.0.2-beta.28", ">=")) {
    return {
      from: `element-plus/es/el-${partialName}`,
      sideEffects: getSideEffectsLegacy(partialName, options)
    };
  } else {
    return {
      from: `element-plus/lib/el-${partialName}`,
      sideEffects: getSideEffectsLegacy(partialName, options)
    };
  }
}
function resolveDirective(name, options) {
  if (!options.directives)
    return;
  const directives2 = {
    Loading: { importName: "ElLoadingDirective", styleName: "loading" },
    Popover: { importName: "ElPopoverDirective", styleName: "popover" },
    InfiniteScroll: { importName: "ElInfiniteScroll", styleName: "infinite-scroll" }
  };
  const directive = directives2[name];
  if (!directive)
    return;
  const { version, ssr, nightly } = options;
  if (compare(version, "1.1.0-beta.1", ">=") || nightly) {
    return {
      name: directive.importName,
      from: `${nightly ? "@element-plus/nightly" : "element-plus"}/${ssr ? "lib" : "es"}`,
      sideEffects: getSideEffects2(directive.styleName, options)
    };
  }
}
var noStylesComponents = ["ElAutoResizer"];
function ElementPlusResolver(options = {}) {
  let optionsResolved;
  async function resolveOptions() {
    if (optionsResolved)
      return optionsResolved;
    optionsResolved = {
      ssr: false,
      version: await _chunk6IWOTYMJcjs.getPkgVersion.call(void 0, "element-plus", "2.2.2"),
      importStyle: "css",
      directives: true,
      exclude: void 0,
      noStylesComponents: options.noStylesComponents || [],
      nightly: false,
      ...options
    };
    return optionsResolved;
  }
  return [
    {
      type: "component",
      resolve: async (name) => {
        const options2 = await resolveOptions();
        if ([...options2.noStylesComponents, ...noStylesComponents].includes(name))
          return resolveComponent(name, { ...options2, importStyle: false });
        else return resolveComponent(name, options2);
      }
    },
    {
      type: "directive",
      resolve: async (name) => {
        return resolveDirective(name, await resolveOptions());
      }
    }
  ];
}

// src/core/resolvers/element-ui.ts
function getSideEffects3(partialName, options) {
  const { importStyle = "css" } = options;
  if (!importStyle)
    return;
  if (importStyle === "sass") {
    return [
      "element-ui/packages/theme-chalk/src/base.scss",
      `element-ui/packages/theme-chalk/src/${partialName}.scss`
    ];
  } else {
    return [
      "element-ui/lib/theme-chalk/base.css",
      `element-ui/lib/theme-chalk/${partialName}.css`
    ];
  }
}
function ElementUiResolver(options = {}) {
  return {
    type: "component",
    resolve: (name) => {
      if (options.exclude && name.match(options.exclude))
        return;
      if (/^El[A-Z]/.test(name)) {
        const compName = name.slice(2);
        const partialName = _chunk6IWOTYMJcjs.kebabCase.call(void 0, compName);
        if (partialName === "collapse-transition") {
          return {
            from: `element-ui/lib/transitions/${partialName}`
          };
        }
        return {
          from: `element-ui/lib/${partialName}`,
          sideEffects: getSideEffects3(partialName, options)
        };
      }
    }
  };
}

// src/core/resolvers/headless-ui.ts
var components = [
  "Combobox",
  "ComboboxButton",
  "ComboboxInput",
  "ComboboxLabel",
  "ComboboxOption",
  "ComboboxOptions",
  "Dialog",
  "DialogDescription",
  "DialogOverlay",
  "DialogPanel",
  "DialogTitle",
  "Disclosure",
  "DisclosureButton",
  "DisclosurePanel",
  "FocusTrap",
  "Listbox",
  "ListboxButton",
  "ListboxLabel",
  "ListboxOption",
  "ListboxOptions",
  "Menu",
  "MenuButton",
  "MenuItem",
  "MenuItems",
  "Popover",
  "PopoverButton",
  "PopoverGroup",
  "PopoverOverlay",
  "PopoverPanel",
  "Portal",
  "PortalGroup",
  "RadioGroup",
  "RadioGroupDescription",
  "RadioGroupLabel",
  "RadioGroupOption",
  "Switch",
  "SwitchDescription",
  "SwitchGroup",
  "SwitchLabel",
  "Tab",
  "TabGroup",
  "TabList",
  "TabPanel",
  "TabPanels",
  "TransitionChild",
  "TransitionRoot"
];
function HeadlessUiResolver(options = {}) {
  const { prefix = "" } = options;
  return {
    type: "component",
    resolve: (name) => {
      if (name.startsWith(prefix)) {
        const componentName = name.substring(prefix.length);
        if (components.includes(componentName)) {
          return {
            name: componentName,
            from: "@headlessui/vue"
          };
        }
      }
    }
  };
}

// src/core/resolvers/idux.ts
var _localpkg = require('local-pkg');
var specialComponents = {
  CdkClickOutside: "click-outside",
  CdkDraggable: "drag-drop",
  CdkResizable: "resize",
  CdkResizableHandle: "resize",
  CdkResizeObserver: "resize",
  CdkVirtualScroll: "scroll",
  IxAutoComplete: "auto-complete",
  IxBackTop: "back-top",
  IxCol: "grid",
  IxDatePicker: "date-picker",
  IxDateRangePicker: "date-picker",
  IxInputNumber: "input-number",
  IxLoadingBar: "loading-bar",
  IxLoadingBarProvider: "loading-bar",
  IxRow: "grid",
  IxTab: "tabs",
  IxTimePicker: "time-picker",
  IxTimeRangePicker: "time-picker",
  IxTreeSelect: "tree-select"
};
function IduxResolver(options = {}) {
  return {
    type: "component",
    resolve: async (name) => {
      const { importStyle, importStyleTheme, exclude = [], scope = "@idux" } = options;
      if (exclude.includes(name))
        return;
      const packageName = getPackageName(name);
      if (!packageName)
        return;
      const resolvedVersion = await _chunk6IWOTYMJcjs.getPkgVersion.call(void 0, `${scope}/${packageName}`, "2.0.0");
      let dirname = specialComponents[name];
      if (!dirname) {
        const nameIndex = packageName === "pro" ? 2 : 1;
        dirname = _chunk6IWOTYMJcjs.kebabCase.call(void 0, name).split("-")[nameIndex];
      }
      const path = `${scope}/${packageName}/${dirname}`;
      const sideEffects = packageName === "cdk" ? void 0 : getSideEffects4(resolvedVersion, path, importStyle, importStyleTheme);
      return { name, from: path, sideEffects };
    }
  };
}
function getPackageName(name) {
  let packageName;
  if (name.match(/^Cdk[A-Z]/))
    packageName = "cdk";
  else if (name.match(/^IxPro[A-Z]/))
    packageName = "pro";
  else if (name.match(/^Ix[A-Z]/))
    packageName = "components";
  return packageName;
}
function getSideEffects4(version, path, importStyle, importStyleTheme) {
  if (!importStyle)
    return;
  if (compare(version, "2.0.0-beta.0", "<"))
    return getLegacySideEffects(path, importStyle, importStyleTheme);
  const styleRoot = `${path}/style`;
  const themeRoot = `${path}/theme`;
  const styleImport = `${styleRoot}/${importStyle === "css" ? "index_css" : "index"}`;
  if (!_localpkg.resolveModule.call(void 0, styleImport))
    return;
  const themeImport = `${themeRoot}/${importStyleTheme}.css`;
  if (!importStyleTheme || !_localpkg.resolveModule.call(void 0, themeImport))
    return styleImport;
  return [styleImport, `${themeRoot}/${importStyleTheme}`];
}
function getLegacySideEffects(path, importStyle, importStyleTheme = "default") {
  const styleImport = `${path}/style/themes/${importStyle === "css" ? `${importStyleTheme}_css` : importStyleTheme}`;
  if (!_localpkg.resolveModule.call(void 0, styleImport))
    return;
  return styleImport;
}

// src/core/resolvers/inkline.ts
function InklineResolver() {
  return {
    type: "component",
    resolve: (name) => {
      if (name.match(/^I[A-Z]/)) {
        return {
          name,
          from: "@inkline/inkline"
        };
      }
    }
  };
}

// src/core/resolvers/naive-ui.ts
function NaiveUiResolver() {
  return {
    type: "component",
    resolve: (name) => {
      if (name.match(/^(N[A-Z]|n-[a-z])/))
        return { name, from: "naive-ui" };
    }
  };
}

// src/core/resolvers/prime-vue.ts
var components2 = [
  "Accordion",
  "AccordionTab",
  "AutoComplete",
  "Avatar",
  "AvatarGroup",
  "Badge",
  "BlockUI",
  "Breadcrumb",
  "Button",
  "Calendar",
  "Card",
  "Carousel",
  "CascadeSelect",
  "Chart",
  "Checkbox",
  "Chip",
  "Chips",
  "ColorPicker",
  "Column",
  "ColumnGroup",
  // 'ConfirmDialog',
  // 'ConfirmPopup',
  // These must be registered globally in order for the confirm service to work properly
  "ContextMenu",
  "DataTable",
  "DataView",
  "DataViewLayoutOptions",
  "DeferredContent",
  "Dialog",
  "Divider",
  "Dock",
  "Dropdown",
  "Editor",
  "Fieldset",
  "FileUpload",
  "FloatLabel",
  "FullCalendar",
  "Galleria",
  "IconField",
  "IconField",
  "Image",
  "InlineMessage",
  "Inplace",
  "InputGroup",
  "InputGroupAddon",
  "InputIcon",
  "InputMask",
  "InputNumber",
  "InputOtp",
  "InputSwitch",
  "InputText",
  "Knob",
  "Listbox",
  "MegaMenu",
  "Menu",
  "Menubar",
  "Message",
  "MeterGroup",
  "MultiSelect",
  "OrderList",
  "OrganizationChart",
  "OverlayPanel",
  "Paginator",
  "Panel",
  "PanelMenu",
  "Password",
  "PickList",
  "ProgressBar",
  "ProgressSpinner",
  "RadioButton",
  "Rating",
  "Row",
  "ScrollPanel",
  "ScrollTop",
  "SelectButton",
  "Sidebar",
  "Skeleton",
  "Slider",
  "SpeedDial",
  "SplitButton",
  "Splitter",
  "SplitterPanel",
  "Stepper",
  "StepperPanel",
  "Steps",
  "TabMenu",
  "TabPanel",
  "TabView",
  "Tag",
  "Terminal",
  "TerminalService",
  "Textarea",
  "TieredMenu",
  "Timeline",
  "Timelist",
  // 'Toast',
  // Toast must be registered globally in order for the toast service to work properly
  "ToggleButton",
  "Toolbar",
  // 'Tooltip',
  // Tooltip must be registered globally in order for the tooltip service to work properly
  "Tree",
  "TreeSelect",
  "TreeTable",
  "TriStateCheckbox",
  "VirtualScroller"
];
function PrimeVueResolver(options = {}) {
  return {
    type: "component",
    resolve: (name) => {
      const sideEffects = [];
      if (options.importStyle)
        sideEffects.push("primevue/resources/primevue.min.css");
      if (options.importIcons)
        sideEffects.push("primeicons/primeicons.css");
      if (options.importTheme) {
        sideEffects.push(
          `primevue/resources/themes/${options.importTheme}/theme.css`
        );
      }
      if (options.prefix) {
        if (!name.startsWith(options.prefix))
          return;
        name = name.substring(options.prefix.length);
      }
      if (components2.includes(name)) {
        return {
          from: `primevue/${name.toLowerCase()}`,
          sideEffects
        };
      }
    }
  };
}

// src/core/resolvers/vant.ts
var moduleType = _chunk6IWOTYMJcjs.isSSR ? "lib" : "es";
function getSideEffects5(dirName, options) {
  const { importStyle = true } = options;
  if (!importStyle || _chunk6IWOTYMJcjs.isSSR)
    return;
  if (importStyle === "less")
    return `vant/${moduleType}/${dirName}/style/less`;
  if (importStyle === "css")
    return `vant/${moduleType}/${dirName}/style/index`;
  return `vant/${moduleType}/${dirName}/style/index`;
}
function VantResolver(options = {}) {
  return {
    type: "component",
    resolve: (name) => {
      if (name.startsWith("Van")) {
        const partialName = name.slice(3);
        return {
          name: partialName,
          from: `vant/${moduleType}`,
          sideEffects: getSideEffects5(_chunk6IWOTYMJcjs.kebabCase.call(void 0, partialName), options)
        };
      }
    }
  };
}

// src/core/resolvers/varlet-ui.ts
var varFunctions = ["ImagePreview", "Snackbar", "Picker", "ActionSheet", "Dialog", "Locale", "StyleProvider", "LoadingBar"];
var varDirectives = ["Ripple", "Lazy", "Hover"];
function getResolved(name, options) {
  const {
    importStyle = "css",
    importCss = true,
    importLess,
    styleExtname = ".mjs",
    autoImport = false,
    version = "vue3"
  } = options;
  const path = version === "vue2" ? "@varlet-vue2/ui" : "@varlet/ui";
  const sideEffects = [];
  if (importStyle || importCss) {
    if (importStyle === "less" || importLess)
      sideEffects.push(`${path}/es/${_chunk6IWOTYMJcjs.kebabCase.call(void 0, name)}/style/less`);
    else
      sideEffects.push(`${path}/es/${_chunk6IWOTYMJcjs.kebabCase.call(void 0, name)}/style/index${styleExtname}`);
  }
  return {
    from: path,
    name: autoImport ? name : `_${name}Component`,
    sideEffects
  };
}
function VarletUIResolver(options = {}) {
  return [
    {
      type: "component",
      resolve: (name) => {
        const { autoImport = false } = options;
        if (autoImport && varFunctions.includes(name))
          return getResolved(name, options);
        if (name.startsWith("Var"))
          return getResolved(name.slice(3), options);
      }
    },
    {
      type: "directive",
      resolve: (name) => {
        const { directives: directives2 = true } = options;
        if (!directives2)
          return;
        if (!varDirectives.includes(name))
          return;
        return getResolved(name, options);
      }
    }
  ];
}

// src/core/resolvers/veui.ts
var _path = require('path');
var _mlly = require('mlly');
var VEUI_PACKAGE_NAME = "veui";
var components3;
function VeuiResolver(options = {}) {
  const { alias = VEUI_PACKAGE_NAME } = options;
  if (!components3) {
    try {
      const componentsData = _chunkOBGZSXTJcjs.__require.call(void 0, `${alias}/components.json`);
      components3 = new Set(componentsData.map(({ name }) => name));
    } catch (e) {
      throw new Error("[unplugin-vue-components:veui] VEUI is not installed");
    }
  }
  return {
    type: "component",
    resolve: (name) => {
      if (name.match(/^Veui[A-Z]/)) {
        const componentName = name.slice(4);
        if (!components3.has(componentName))
          return;
        const sideEffects = getSideEffects6(componentName, options);
        return { name: componentName, from: alias, sideEffects };
      }
    }
  };
}
var formatters = {
  "kebab-case": _chunk6IWOTYMJcjs.kebabCase,
  "camelCase": _chunk6IWOTYMJcjs.camelCase,
  "PascalCase": _chunk6IWOTYMJcjs.pascalCase
};
var peerPaths = /* @__PURE__ */ new Map();
function assertPeerPath(peerPath) {
  if (!peerPaths.has(peerPath)) {
    try {
      _mlly.resolvePathSync.call(void 0, peerPath);
      peerPaths.set(peerPath, true);
    } catch (e2) {
      peerPaths.set(peerPath, false);
    }
  }
  return peerPaths.get(peerPath);
}
function getSideEffects6(name, {
  alias = VEUI_PACKAGE_NAME,
  modules = [],
  locale = "zh-Hans",
  global = []
}) {
  const localeModules = (locale ? Array.isArray(locale) ? locale : [locale] : []).map((locale2) => `${alias}/locale/${locale2}/${name}.js`);
  const peerModules = modules.map(
    ({
      package: pack,
      path = "components",
      fileName = "{module}.css",
      transform = "kebab-case"
    }) => {
      const peer = transform ? formatters[transform](name) : name;
      const file = fileName.replace(/\$?\{module\}/g, peer);
      return _path.normalize.call(void 0, _path.join.call(void 0, pack, path, file));
    }
  );
  return [...localeModules, ...global, ...peerModules].filter(assertPeerPath);
}

// src/core/resolvers/view-ui.ts
function getSideEffects7(componentName) {
  const sideEffects = [
    "view-design/dist/styles/iview.css",
    "popper.js/dist/umd/popper.js"
  ];
  if (/^Table|^Slider|^Tab/.test(componentName))
    sideEffects.push("element-resize-detector");
  if (componentName.startsWith("Date"))
    sideEffects.push("js-calendar");
  return sideEffects;
}
var matchComponents2 = [
  {
    pattern: /^List/,
    compDir: "list"
  }
];
function getCompDir(compName) {
  let compPath;
  const total = matchComponents2.length;
  for (let i = 0; i < total; i++) {
    const matcher = matchComponents2[i];
    if (compName.match(matcher.pattern)) {
      compPath = `${matcher.compDir}/${_chunk6IWOTYMJcjs.kebabCase.call(void 0, compName)}.vue`;
      break;
    }
  }
  if (!compPath)
    compPath = _chunk6IWOTYMJcjs.kebabCase.call(void 0, compName);
  return compPath;
}
function ViewUiResolver() {
  return {
    type: "component",
    resolve: (name) => {
      if (name.match(/^I[A-Z]/)) {
        const compName = name.slice(1);
        return {
          from: `view-design/src/components/${getCompDir(compName)}`,
          sideEffects: getSideEffects7(compName)
        };
      }
    }
  };
}

// src/core/resolvers/vuetify.ts
function VuetifyResolver() {
  return {
    type: "component",
    resolve: (name) => {
      if (name.match(/^V[A-Z]/))
        return { name, from: "vuetify/lib" };
    }
  };
}
function Vuetify3Resolver() {
  return {
    type: "component",
    resolve: (name) => {
      if (name.match(/^V[A-Z]/))
        return { name, from: "vuetify/components" };
    }
  };
}

// src/core/resolvers/vueuse.ts
var _fs = require('fs');
var _process = require('process'); var _process2 = _interopRequireDefault(_process);

var components4;
function VueUseComponentsResolver() {
  return {
    type: "component",
    resolve: (name) => {
      if (!components4) {
        let indexesJson;
        try {
          const corePath = _localpkg.resolveModule.call(void 0, "@vueuse/core") || _process2.default.cwd();
          const path = _localpkg.resolveModule.call(void 0, "@vueuse/core/indexes.json") || _localpkg.resolveModule.call(void 0, "@vueuse/metadata/index.json") || _localpkg.resolveModule.call(void 0, "@vueuse/metadata/index.json", { paths: [corePath] });
          indexesJson = JSON.parse(_fs.readFileSync.call(void 0, path, "utf-8"));
          components4 = indexesJson.functions.filter((i) => i.component && i.name).map(({ name: name2 }) => name2[0].toUpperCase() + name2.slice(1));
        } catch (error) {
          console.error(error);
          throw new Error("[vue-components] failed to load @vueuse/core, have you installed it?");
        }
      }
      if (components4 && components4.includes(name))
        return { name, as: name, from: "@vueuse/components" };
    }
  };
}

// src/core/resolvers/vueuse-directive.ts



var directives;
function VueUseDirectiveResolver() {
  return {
    type: "directive",
    resolve: (name) => {
      if (!directives) {
        let indexesJson;
        try {
          const corePath = _localpkg.resolveModule.call(void 0, "@vueuse/core") || _process2.default.cwd();
          const path = _localpkg.resolveModule.call(void 0, "@vueuse/core/indexes.json") || _localpkg.resolveModule.call(void 0, "@vueuse/metadata/index.json") || _localpkg.resolveModule.call(void 0, "@vueuse/metadata/index.json", { paths: [corePath] });
          indexesJson = JSON.parse(_fs.readFileSync.call(void 0, path, "utf-8"));
          directives = indexesJson.functions.filter((i) => i.directive && i.name).map(({ name: name2 }) => name2[0].toUpperCase() + name2.slice(1)).map((name2) => name2.startsWith("Use") ? name2.slice(3) : name2);
        } catch (error) {
          console.error(error);
          throw new Error("[vue-components] failed to load @vueuse/core, have you installed it?");
        }
      }
      if (directives && directives.includes(name))
        return { name: `v${name}`, as: name, from: "@vueuse/components" };
    }
  };
}

// src/core/resolvers/quasar.ts


function QuasarResolver() {
  let components5 = [];
  return {
    type: "component",
    resolve: async (name) => {
      if (!components5.length) {
        const quasarApiListPath = _localpkg.resolveModule.call(void 0, "quasar/dist/transforms/api-list.json");
        if (quasarApiListPath)
          components5 = JSON.parse(await _fs.promises.readFile(quasarApiListPath, "utf-8"));
      }
      if (components5.includes(name))
        return { name, from: "quasar" };
    }
  };
}

// src/core/resolvers/devui.ts
var LIB_NAME = "vue-devui";
var HARMLESS = ["ripple"];
function resolveDirectory(name, filename) {
  return `${LIB_NAME}/${name}/${filename}`;
}
function getSideEffects8(name, filename) {
  if (HARMLESS.includes(name))
    return;
  if (["row", "col"].includes(name))
    return resolveDirectory("grid", filename);
  if (["aside", "content", "footer", "header", "layout"].includes(name))
    return resolveDirectory("layout", filename);
  if (["overlay", "fixed-overlay", "flexible-overlay"].includes(name))
    return resolveDirectory("overlay", filename);
  if (["panel", "panel-header", "panel-body"].includes(name))
    return resolveDirectory("panel", filename);
  if (["menu", "menu-item", "sub-menu"].includes(name))
    return resolveDirectory("menu", filename);
  if (["tabs", "tab"].includes(name))
    return resolveDirectory("tabs", filename);
  if (["form", "form-item"].includes(name))
    return resolveDirectory("form", filename);
  if (["collapse", "collapse-item"].includes(name))
    return resolveDirectory("collapse", filename);
  if (["steps", "step"].includes(name))
    return resolveDirectory("steps", filename);
  if (["radio", "radio-group", "radio-button"].includes(name))
    return resolveDirectory("radio", filename);
  if (["column"].includes(name))
    return resolveDirectory("table", filename);
  if (["timeline-item"].includes(name))
    return resolveDirectory("timeline", filename);
  if (["splitter-pane"].includes(name))
    return resolveDirectory("splitter", filename);
  return resolveDirectory(name, filename);
}
function componentsResolver(name, { ssr }) {
  if (!name.match(/^D[A-Z]/))
    return;
  const resolveId = _chunk6IWOTYMJcjs.kebabCase.call(void 0, name = name.slice(1));
  return {
    name,
    sideEffects: getSideEffects8(resolveId, "style.css"),
    from: getSideEffects8(resolveId, `index.${ssr ? "umd" : "es"}.js`)
  };
}
function directivesResolver(name, { ssr }) {
  const resolveId = _chunk6IWOTYMJcjs.kebabCase.call(void 0, name);
  return {
    name: `${name}Directive`,
    sideEffects: getSideEffects8(resolveId, "style.css"),
    from: resolveDirectory(resolveId, `index.${ssr ? "umd" : "es"}.js`)
  };
}
function DevUiResolver(options = {}) {
  const config = { directives: true, importStyle: true, ssr: false, ...options };
  const resolvers = [
    {
      type: "component",
      resolve: (name) => componentsResolver(name, config)
    }
  ];
  if (config.directives) {
    resolvers.push({
      type: "directive",
      resolve: (name) => directivesResolver(name, config)
    });
  }
  return resolvers;
}

// src/core/resolvers/arco.ts
var _debug = require('debug'); var _debug2 = _interopRequireDefault(_debug);
var debug = _debug2.default.call(void 0, "unplugin-vue-components:resolvers:arco");
var matchComponents3 = [
  {
    pattern: /^AnchorLink$/,
    componentDir: "anchor"
  },
  {
    pattern: /^AvatarGroup$/,
    componentDir: "avatar"
  },
  {
    pattern: /^BreadcrumbItem$/,
    componentDir: "breadcrumb"
  },
  {
    pattern: /^ButtonGroup$/,
    componentDir: "button"
  },
  {
    pattern: /^(CardMeta|CardGrid)$/,
    componentDir: "card"
  },
  {
    pattern: /^CarouselItem$/,
    componentDir: "carousel"
  },
  {
    pattern: /^CascaderPanel$/,
    componentDir: "cascader"
  },
  {
    pattern: /^CheckboxGroup$/,
    componentDir: "checkbox"
  },
  {
    pattern: /^CollapseItem$/,
    componentDir: "collapse"
  },
  {
    pattern: /^(WeekPicker|MonthPicker|YearPicker|QuarterPicker|RangePicker)$/,
    componentDir: "date-picker"
  },
  {
    pattern: /^DescriptionsItem$/,
    componentDir: "descriptions"
  },
  {
    pattern: /^(Doption|Dgroup|Dsubmenu|DropdownButton)$/,
    componentDir: "dropdown"
  },
  {
    pattern: /^FormItem$/,
    componentDir: "form"
  },
  {
    pattern: /^(Col|Row|GridItem)$/,
    componentDir: "grid"
  },
  {
    pattern: /^(ImagePreview|ImagePreviewGroup)$/,
    componentDir: "image"
  },
  {
    pattern: /^(InputGroup|InputSearch|InputPassword)$/,
    componentDir: "input"
  },
  {
    pattern: /^(LayoutHeader|LayoutContent|LayoutFooter|LayoutSider)$/,
    componentDir: "layout"
  },
  {
    pattern: /^(ListItem|ListItemMeta)$/,
    componentDir: "list"
  },
  {
    pattern: /^(MenuItem|MenuItemGroup|SubMenu)$/,
    componentDir: "menu"
  },
  {
    pattern: /^RadioGroup$/,
    componentDir: "radio"
  },
  {
    pattern: /^(Option|Optgroup)$/,
    componentDir: "select"
  },
  {
    pattern: /^(SkeletonLine|SkeletonShape)$/,
    componentDir: "skeleton"
  },
  {
    pattern: /^Countdown$/,
    componentDir: "statistic"
  },
  {
    pattern: /^Step$/,
    componentDir: "steps"
  },
  {
    pattern: /^(Thead|Td|Th|Tr|Tbody|TableColumn)$/,
    componentDir: "table"
  },
  {
    pattern: /^TagGroup$/,
    componentDir: "tag"
  },
  {
    pattern: /^TabPane$/,
    componentDir: "tabs"
  },
  {
    pattern: /^TimelineItem$/,
    componentDir: "timeline"
  },
  {
    pattern: /^(TypographyParagraph|TypographyTitle|TypographyText)$/,
    componentDir: "typography"
  }
];
function getComponentStyleDir(importName, importStyle) {
  if (["ConfigProvider", "Icon"].includes(importName))
    return void 0;
  let componentDir = _chunk6IWOTYMJcjs.kebabCase.call(void 0, importName);
  for (const item of matchComponents3) {
    if (item.pattern.test(importName)) {
      componentDir = item.componentDir;
      break;
    }
  }
  if (importStyle === "less")
    return `@arco-design/web-vue/es/${componentDir}/style/index.js`;
  if (importStyle === "css" || importStyle)
    return `@arco-design/web-vue/es/${componentDir}/style/css.js`;
}
function canResolveIcons(options) {
  if (options === void 0)
    return false;
  if (typeof options === "boolean")
    return options;
  else
    return options.enable;
}
function getResolveIconPrefix(options) {
  if (canResolveIcons(options)) {
    if (typeof options === "boolean" && options)
      return "";
    else if (options.enable)
      return _nullishCoalesce(options.iconPrefix, () => ( ""));
    else
      return "";
  }
  return "";
}
function ArcoResolver(options = {}) {
  return {
    type: "component",
    resolve: (name) => {
      if (canResolveIcons(options.resolveIcons)) {
        const iconPrefix = _chunk6IWOTYMJcjs.pascalCase.call(void 0, getResolveIconPrefix(options.resolveIcons));
        const newNameRegexp = new RegExp(`^${iconPrefix}Icon`);
        if (newNameRegexp.test(name)) {
          debug("found icon component name %s", name);
          const rawComponentName = name.slice(iconPrefix.length);
          debug("found icon component raw name %s", rawComponentName);
          return {
            name: rawComponentName,
            as: name,
            from: "@arco-design/web-vue/es/icon"
          };
        }
      }
      if (name.match(/^A[A-Z]/) && !_chunk6IWOTYMJcjs.isExclude.call(void 0, name, options.exclude)) {
        const importStyle = _nullishCoalesce(options.importStyle, () => ( "css"));
        const importName = name.slice(1);
        const config = {
          name: importName,
          from: "@arco-design/web-vue"
        };
        if (options.sideEffect !== false)
          config.sideEffects = getComponentStyleDir(importName, importStyle);
        return config;
      }
    }
  };
}

// src/core/resolvers/tdesign.ts
function TDesignResolver(options = {}) {
  const pluginList = ["DialogPlugin", "LoadingPlugin", "MessagePlugin", "NotifyPlugin"];
  return {
    type: "component",
    resolve: (name) => {
      const { library = "vue", exclude } = options;
      const importFrom = options.esm ? "/esm" : "";
      if (_chunk6IWOTYMJcjs.isExclude.call(void 0, name, exclude))
        return;
      if (options.resolveIcons && name.match(/[a-z]Icon$/)) {
        return {
          name,
          from: `tdesign-icons-${library}${importFrom}`
        };
      }
      if (name.match(/^T[A-Z]/) || pluginList.includes(name)) {
        const importName = name.match(/^T[A-Z]/) ? name.slice(1) : name;
        return {
          name: importName,
          from: `tdesign-${library}${importFrom}`
        };
      }
    }
  };
}

// src/core/resolvers/layui-vue.ts
var matchComponents4 = [
  {
    pattern: /^LayAvatarList$/,
    styleDir: "avatar"
  },
  {
    pattern: /^LayBreadcrumbItem$/,
    styleDir: "breadcrumb"
  },
  {
    pattern: /^(LayCarouselItem)$/,
    styleDir: "carousel"
  },
  {
    pattern: /^(LayCheckboxGroup)$/,
    styleDir: "checkbox"
  },
  {
    pattern: /^LayCol$/,
    styleDir: "row"
  },
  {
    pattern: /^(LayCollapseItem)$/,
    styleDir: "collapse"
  },
  {
    pattern: /^LayConfigProvider$/,
    styleDir: void 0
  },
  {
    pattern: /^LayCountUp$/,
    styleDir: void 0
  },
  {
    pattern: /^(LayDropdownMenu|LayDropdownMenuItem|LayDropdownSubMenu)$/,
    styleDir: "dropdown"
  },
  {
    pattern: /^(LayFormItem)$/,
    styleDir: "form"
  },
  {
    pattern: /^(LayMenuItem|LaySubMenu)$/,
    styleDir: "menu"
  },
  {
    pattern: /^(LayRadioGroup)$/,
    styleDir: "radio"
  },
  {
    pattern: /^LaySelectOption$/,
    styleDir: "select"
  },
  {
    pattern: /^LaySkeletonItem$/,
    styleDir: "skeleton"
  },
  {
    pattern: /^LaySplitPanelItem$/,
    styleDir: "splitPanel"
  },
  {
    pattern: /^LayStepItem$/,
    styleDir: "step"
  },
  {
    pattern: /^(LayTabItem)$/,
    styleDir: "tab"
  },
  {
    pattern: /^LayTimelineItem$/,
    styleDir: "timeline"
  }
];
var layuiRE = /^Lay[A-Z]/;
var layerRE = /^(layer|LayLayer)$/;
var iconsRE = /^[A-Z]\w+Icon$/;
var libName = "@layui/layui-vue";
function lowerCamelCase(str) {
  return str.charAt(0).toLowerCase() + str.slice(1);
}
function getSideEffects9(importName, options) {
  const { importStyle = "css" } = options;
  if (!importStyle)
    return void 0;
  if (libName !== "@layui/layui-vue")
    return `${libName}/lib/index.css`;
  let styleDir = lowerCamelCase(importName.slice(3));
  for (const item of matchComponents4) {
    if (item.pattern.test(importName)) {
      styleDir = item.styleDir;
      break;
    }
  }
  if (importStyle === "css" || importStyle) {
    return styleDir ? [`@layui/layui-vue/es/${styleDir}/index.css`, "@layui/layui-vue/es/index/index.css"] : void 0;
  }
}
function resolveComponent2(importName, options) {
  let name;
  if (_chunk6IWOTYMJcjs.isExclude.call(void 0, importName, options.exclude))
    return void 0;
  if (options.resolveIcons && importName.match(iconsRE)) {
    name = importName;
    libName = "@layui/icons-vue";
  } else if (importName.match(layerRE)) {
    name = importName;
    libName = "@layui/layer-vue";
  } else if (importName.match(layuiRE) && !importName.match(iconsRE)) {
    name = importName;
    libName = "@layui/layui-vue";
  }
  return name ? { name, from: libName, sideEffects: getSideEffects9(name, options) } : void 0;
}
function LayuiVueResolver(options = {}) {
  return {
    type: "component",
    resolve: (name) => {
      return resolveComponent2(name, options);
    }
  };
}

// src/core/resolvers/bootstrap-vue.ts
var COMPONENT_ALIASES = {
  BBtn: "BButton",
  BBtnClose: "BButtonClose",
  BBtnGroup: "BButtonGroup",
  BBtnToolbar: "BButtonToolbar",
  BCheck: "BFormCheckbox",
  BCheckbox: "BFormCheckbox",
  BCheckboxGroup: "BFormCheckboxGroup",
  BCheckGroup: "BFormCheckboxGroup",
  BDatalist: "BFormDatalist",
  BDatepicker: "BFormDatepicker",
  BDd: "BDropdown",
  BDdDivider: "BDropdownDivider",
  BDdForm: "BDropdownForm",
  BDdGroup: "BDropdownGroup",
  BDdHeader: "BDropdownHeader",
  BDdItem: "BDropdownItem",
  BDdItemBtn: "BDropdownItemButton",
  BDdItemButton: "BDropdownItemButton",
  BDdText: "BDropdownText",
  BDropdownItemBtn: "BDropdownItemButton",
  BFile: "BFormFile",
  BInput: "BFormInput",
  BNavDd: "BNavItemDropdown",
  BNavDropdown: "BNavItemDropdown",
  BNavItemDd: "BNavItemDropdown",
  BNavToggle: "BNavbarToggle",
  BRadio: "BFormRadio",
  BRadioGroup: "BFormRadioGroup",
  BRating: "BFormRating",
  BSelect: "BFormSelect",
  BSelectOption: "BFormSelectOption",
  BSelectOptionGroup: "BFormSelectOptionGroup",
  BSpinbutton: "BFormSpinbutton",
  BTag: "BFormTag",
  BTags: "BFormTags",
  BTextarea: "BFormTextarea",
  BTimepicker: "BFormTimepicker"
};
function BootstrapVueResolver(_options = {}) {
  const options = { directives: true, ..._options };
  const resolvers = [{
    type: "component",
    resolve: (name) => {
      if (name.match(/^B[A-Z]/)) {
        return {
          name: COMPONENT_ALIASES[name] || name,
          from: "bootstrap-vue"
        };
      }
    }
  }];
  if (options.directives) {
    resolvers.push({
      type: "directive",
      resolve: (name) => {
        if (name.match(/^B[A-Z]/)) {
          return {
            name: `V${name}`,
            from: "bootstrap-vue"
          };
        }
      }
    });
  }
  return resolvers;
}
function BootstrapVueNextResolver(_options = {}) {
  const options = { directives: true, ..._options };
  const resolvers = [{
    type: "component",
    resolve: (name) => {
      if (name.match(/^B[A-Z]/))
        return { name, from: "bootstrap-vue-next" };
    }
  }];
  if (options.directives) {
    resolvers.push({
      type: "directive",
      resolve: (name) => {
        if (name.match(/^B[A-Z]/))
          return { name: `v${name}`, from: "bootstrap-vue-next" };
      }
    });
  }
  return resolvers;
}
function BootstrapVue3Resolver(_options = {}) {
  const options = { directives: true, ..._options };
  const resolvers = [{
    type: "component",
    resolve: (name) => {
      if (name.match(/^B[A-Z]/))
        return { name, from: "bootstrap-vue-3" };
    }
  }];
  if (options.directives) {
    resolvers.push({
      type: "directive",
      resolve: (name) => {
        if (name.match(/^B[A-Z]/))
          return { name: `V${name}`, from: "bootstrap-vue-3" };
      }
    });
  }
  return resolvers;
}

// src/core/resolvers/ionic.ts
var IonicBuiltInComponents = [
  "IonAccordion",
  "IonAccordionGroup",
  "IonActionSheet",
  "IonAlert",
  "IonApp",
  "IonAvatar",
  "IonBackButton",
  "IonBackdrop",
  "IonBadge",
  "IonBreadcrumb",
  "IonBreadcrumbs",
  "IonButton",
  "IonButtons",
  "IonCard",
  "IonCardContent",
  "IonCardHeader",
  "IonCardSubtitle",
  "IonCardTitle",
  "IonCheckbox",
  "IonChip",
  "IonCol",
  "IonContent",
  "IonDatetime",
  "IonDatetimeButton",
  "IonFab",
  "IonFabButton",
  "IonFabList",
  "IonFooter",
  "IonGrid",
  "IonHeader",
  "IonIcon",
  "IonImg",
  "IonInfiniteScroll",
  "IonInfiniteScrollContent",
  "IonInput",
  "IonItem",
  "IonItemDivider",
  "IonItemGroup",
  "IonItemOption",
  "IonItemOptions",
  "IonItemSliding",
  "IonLabel",
  "IonList",
  "IonListHeader",
  "IonLoading",
  "IonMenu",
  "IonMenuButton",
  "IonMenuToggle",
  "IonModal",
  "IonNav",
  "IonNavLink",
  "IonNote",
  "IonPage",
  "IonPicker",
  "IonPopover",
  "IonProgressBar",
  "IonRadio",
  "IonRadioGroup",
  "IonRange",
  "IonRefresher",
  "IonRefresherContent",
  "IonReorder",
  "IonReorderGroup",
  "IonRippleEffect",
  "IonRouterOutlet",
  "IonRow",
  "IonSearchbar",
  "IonSegment",
  "IonSegmentButton",
  "IonSelect",
  "IonSelectOption",
  "IonSkeletonText",
  "IonSpinner",
  "IonSplitPane",
  "IonTabBar",
  "IonTabButton",
  "IonTabs",
  "IonText",
  "IonTextarea",
  "IonThumbnail",
  "IonTitle",
  "IonToast",
  "IonToggle",
  "IonToolbar"
];
function IonicResolver() {
  return {
    type: "component",
    resolve: (name) => {
      if (IonicBuiltInComponents.includes(name)) {
        return {
          name,
          from: "@ionic/vue"
        };
      }
    }
  };
}




























exports.AntDesignVueResolver = AntDesignVueResolver; exports.ArcoResolver = ArcoResolver; exports.BootstrapVue3Resolver = BootstrapVue3Resolver; exports.BootstrapVueNextResolver = BootstrapVueNextResolver; exports.BootstrapVueResolver = BootstrapVueResolver; exports.DevUiResolver = DevUiResolver; exports.ElementPlusResolver = ElementPlusResolver; exports.ElementUiResolver = ElementUiResolver; exports.HeadlessUiResolver = HeadlessUiResolver; exports.IduxResolver = IduxResolver; exports.InklineResolver = InklineResolver; exports.IonicBuiltInComponents = IonicBuiltInComponents; exports.IonicResolver = IonicResolver; exports.LayuiVueResolver = LayuiVueResolver; exports.NaiveUiResolver = NaiveUiResolver; exports.PrimeVueResolver = PrimeVueResolver; exports.QuasarResolver = QuasarResolver; exports.TDesignResolver = TDesignResolver; exports.VantResolver = VantResolver; exports.VarletUIResolver = VarletUIResolver; exports.VeuiResolver = VeuiResolver; exports.ViewUiResolver = ViewUiResolver; exports.VueUseComponentsResolver = VueUseComponentsResolver; exports.VueUseDirectiveResolver = VueUseDirectiveResolver; exports.Vuetify3Resolver = Vuetify3Resolver; exports.VuetifyResolver = VuetifyResolver; exports.getResolved = getResolved;
