"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createOptimizeDeps = void 0;
function createOptimizeDeps(_options) {
    return {
        exclude: [
            'vue',
            'vuex',
            'vue-i18n',
            'vue-router',
            '@dcloudio/uni-app',
            '@dcloudio/uni-components',
            '@dcloudio/uni-i18n',
            '@dcloudio/uni-shared',
            '@dcloudio/uni-stacktracey',
            'pinia',
        ],
    };
}
exports.createOptimizeDeps = createOptimizeDeps;
