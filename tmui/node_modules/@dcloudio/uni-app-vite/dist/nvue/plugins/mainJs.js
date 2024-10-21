"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniMainJsPlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const appCss_1 = require("./appCss");
function uniMainJsPlugin({ renderer, appService, }) {
    return (0, uni_cli_shared_1.defineUniMainJsPlugin)((opts) => {
        return {
            name: 'uni:app-nvue-main-js',
            enforce: 'pre',
            transform(code, id) {
                if (opts.filter(id)) {
                    if (renderer !== 'native') {
                        return {
                            code: `import './${uni_cli_shared_1.PAGES_JSON_JS}';import('${appCss_1.APP_CSS_JS}').then(()=>{})`,
                            map: { mappings: '' },
                        };
                    }
                    if (appService) {
                        code = code.includes('createSSRApp')
                            ? createApp(code)
                            : createLegacyApp(code);
                        return {
                            code: `import './${uni_cli_shared_1.MANIFEST_JSON_JS}';\nimport './${uni_cli_shared_1.PAGES_JSON_JS}';\n` +
                                code,
                            map: { mappings: '' },
                        };
                    }
                    return {
                        code: `import './${uni_cli_shared_1.PAGES_JSON_JS}';`,
                        map: { mappings: '' },
                    };
                }
            },
        };
    });
}
exports.uniMainJsPlugin = uniMainJsPlugin;
function createApp(code) {
    return `${code.replace('createSSRApp', 'createVueApp as createSSRApp')};const {app:__app__,Vuex:__Vuex__,Pinia:__Pinia__}=createApp();uni.Vuex=__Vuex__;uni.Pinia=__Pinia__;__app__._component.mpType='app';__app__._component.render=()=>{};__app__.mount('#app');`;
}
function createLegacyApp(code) {
    return `function createApp(rootComponent,rootProps){rootComponent.mpTye='app';rootComponent.render=()=>{};const app=createVueApp(rootComponent,rootProps);const oldMount=app.mount;app.mount=(container)=>{const appVm=oldMount.call(app,container);return appVm;};return app;};${code.replace('createApp', 'createVueApp')}`;
}
