"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.formatInstallHBuilderXPluginTips = exports.moduleAliasFormatter = exports.installHBuilderXPlugin = exports.initModuleAlias = void 0;
const path_1 = __importDefault(require("path"));
const module_alias_1 = __importDefault(require("module-alias"));
const env_1 = require("./env");
const hbxPlugins = {
    typescript: 'compile-typescript/node_modules/typescript',
    less: 'compile-less/node_modules/less',
    sass: 'compile-dart-sass/node_modules/sass',
    stylus: 'compile-stylus/node_modules/stylus',
    pug: 'compile-pug-cli/node_modules/pug',
};
function initModuleAlias() {
    const compilerSfcPath = path_1.default.resolve(__dirname, '../../lib/@vue/compiler-sfc');
    const serverRendererPath = require.resolve('@vue/server-renderer');
    module_alias_1.default.addAliases({
        '@vue/shared': require.resolve('@vue/shared'),
        '@vue/shared/dist/shared.esm-bundler.js': require.resolve('@vue/shared/dist/shared.esm-bundler.js'),
        '@vue/compiler-dom': require.resolve('@vue/compiler-dom'),
        '@vue/compiler-sfc': compilerSfcPath,
        '@vue/server-renderer': serverRendererPath,
        'vue/compiler-sfc': compilerSfcPath,
        'vue/server-renderer': serverRendererPath,
    });
    if (process.env.VITEST) {
        module_alias_1.default.addAliases({
            vue: '@dcloudio/uni-h5-vue',
            'vue/package.json': '@dcloudio/uni-h5-vue/package.json',
        });
    }
    if ((0, env_1.isInHBuilderX)()) {
        // 又是为了复用 HBuilderX 的插件逻辑，硬编码映射
        Object.keys(hbxPlugins).forEach((lang) => {
            const realPath = path_1.default.resolve(process.env.UNI_HBUILDERX_PLUGINS, hbxPlugins[lang]);
            module_alias_1.default.addAlias(lang, 
            // @ts-expect-error
            () => {
                try {
                    require.resolve(realPath);
                }
                catch (e) {
                    const msg = exports.moduleAliasFormatter.format(`Preprocessor dependency "${lang}" not found. Did you install it?`);
                    console.error(msg);
                    process.exit(0);
                }
                return realPath;
            });
        });
        // web 平台用了 vite 内置 css 插件，该插件会加载预编译器如scss、less等，需要转向到 HBuilderX 的对应编译器插件
        if (process.env.UNI_PLATFORM === 'h5' ||
            process.env.UNI_PLATFORM === 'web') {
            // https://github.com/vitejs/vite/blob/main/packages/vite/src/node/packages.ts#L92
            // 拦截预编译器
            const join = path_1.default.join;
            path_1.default.join = function (...paths) {
                if (paths.length === 4) {
                    // path.join(basedir, 'node_modules', pkgName, 'package.json')
                    // const basedir = paths[0]
                    const nodeModules = paths[1]; // = node_modules
                    const pkgName = paths[2];
                    const packageJson = paths[3]; // = package.json
                    if (nodeModules === 'node_modules' &&
                        packageJson === 'package.json' &&
                        hbxPlugins[pkgName]) {
                        return path_1.default.resolve(process.env.UNI_HBUILDERX_PLUGINS, hbxPlugins[pkgName], packageJson);
                    }
                }
                return join(...paths);
            };
            // https://github.com/vitejs/vite/blob/892916d040a035edde1add93c192e0b0c5c9dd86/packages/vite/src/node/plugins/css.ts#L1481
            // const oldSync = resovle.sync
            // resovle.sync = (id: string, opts?: SyncOpts) => {
            //   if ((hbxPlugins as any)[id]) {
            //     return path.resolve(
            //       process.env.UNI_HBUILDERX_PLUGINS,
            //       hbxPlugins[id as keyof typeof hbxPlugins]
            //     )
            //   }
            //   return oldSync(id, opts)
            // }
        }
    }
}
exports.initModuleAlias = initModuleAlias;
function supportAutoInstallPlugin() {
    return !!process.env.HX_Version;
}
function installHBuilderXPlugin(plugin) {
    if (!supportAutoInstallPlugin()) {
        return;
    }
    return console.error(`%HXRunUniAPPPluginName%${plugin}%HXRunUniAPPPluginName%`);
}
exports.installHBuilderXPlugin = installHBuilderXPlugin;
const installPreprocessorTips = {};
exports.moduleAliasFormatter = {
    test(msg) {
        return msg.includes('Preprocessor dependency');
    },
    format(msg) {
        let lang = '';
        let preprocessor = '';
        if (msg.includes(`"pug"`)) {
            lang = 'pug';
            preprocessor = 'compile-pug-cli';
        }
        else if (msg.includes(`"sass"`)) {
            lang = 'sass';
            preprocessor = 'compile-dart-sass';
        }
        else if (msg.includes(`"less"`)) {
            lang = 'less';
            preprocessor = 'compile-less';
        }
        else if (msg.includes('"stylus"')) {
            lang = 'stylus';
            preprocessor = 'compile-stylus';
        }
        else if (msg.includes('"typescript"')) {
            lang = 'typescript';
            preprocessor = 'compile-typescript';
        }
        if (lang) {
            // 仅提醒一次
            if (installPreprocessorTips[lang]) {
                return '';
            }
            installPreprocessorTips[lang] = true;
            installHBuilderXPlugin(preprocessor);
            return formatInstallHBuilderXPluginTips(lang, preprocessor);
        }
        return msg;
    },
};
function formatInstallHBuilderXPluginTips(lang, preprocessor) {
    return `预编译器错误：代码使用了${lang}语言，但未安装相应的编译器插件，${supportAutoInstallPlugin() ? '正在从' : '请前往'}插件市场安装该插件:
https://ext.dcloud.net.cn/plugin?name=${preprocessor}`;
}
exports.formatInstallHBuilderXPluginTips = formatInstallHBuilderXPluginTips;
