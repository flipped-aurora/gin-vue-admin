"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getDefaultVueLanguagePlugins = void 0;
const file_html_1 = require("./plugins/file-html");
const file_md_1 = require("./plugins/file-md");
const file_vue_1 = require("./plugins/file-vue");
const vue_sfc_customblocks_1 = require("./plugins/vue-sfc-customblocks");
const vue_sfc_scripts_1 = require("./plugins/vue-sfc-scripts");
const vue_sfc_styles_1 = require("./plugins/vue-sfc-styles");
const vue_sfc_template_1 = require("./plugins/vue-sfc-template");
const vue_template_html_1 = require("./plugins/vue-template-html");
const vue_tsx_1 = require("./plugins/vue-tsx");
const CompilerDOM = require("@vue/compiler-dom");
const CompilerVue2 = require("./utils/vue2TemplateCompiler");
function getDefaultVueLanguagePlugins(ts, compilerOptions, vueCompilerOptions, codegenStack) {
    const plugins = [
        file_md_1.default, // .md for VitePress
        file_html_1.default, // .html for PetiteVue
        file_vue_1.default, // .vue and others for Vue
        vue_template_html_1.default,
        vue_sfc_styles_1.default,
        vue_sfc_customblocks_1.default,
        vue_sfc_scripts_1.default,
        vue_sfc_template_1.default,
        vue_tsx_1.default,
        ...vueCompilerOptions.plugins,
    ];
    const pluginCtx = {
        modules: {
            '@vue/compiler-dom': vueCompilerOptions.target < 3
                ? {
                    ...CompilerDOM,
                    compile: CompilerVue2.compile,
                }
                : CompilerDOM,
            typescript: ts,
        },
        compilerOptions,
        vueCompilerOptions,
        codegenStack,
    };
    const pluginInstances = plugins
        .map(plugin => plugin(pluginCtx))
        .sort((a, b) => {
        const aOrder = a.order ?? 0;
        const bOrder = b.order ?? 0;
        return aOrder - bOrder;
    });
    return pluginInstances.filter((plugin) => {
        const valid = plugin.version >= 1 && plugin.version < 2;
        if (!valid) {
            console.warn(`Plugin ${JSON.stringify(plugin.name)} API version incompatible, expected 1.x but got ${JSON.stringify(plugin.version)}`);
        }
        return valid;
    });
}
exports.getDefaultVueLanguagePlugins = getDefaultVueLanguagePlugins;
//# sourceMappingURL=plugins.js.map