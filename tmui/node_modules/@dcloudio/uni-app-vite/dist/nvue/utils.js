"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.esbuildGlobals = exports.globals = exports.external = void 0;
function external(appService) {
    return appService ? ['vue'] : ['vue', 'vuex', 'pinia'];
}
exports.external = external;
function globals(appService) {
    return appService
        ? { vue: 'Vue' }
        : {
            vue: 'Vue',
            vuex: 'Vuex',
            pinia: 'Pinia',
        };
}
exports.globals = globals;
function esbuildGlobals(appService) {
    return appService
        ? { vue: 'Vue' }
        : {
            vue: 'Vue',
            vuex: 'uni.Vuex',
            pinia: 'uni.Pinia',
        };
}
exports.esbuildGlobals = esbuildGlobals;
