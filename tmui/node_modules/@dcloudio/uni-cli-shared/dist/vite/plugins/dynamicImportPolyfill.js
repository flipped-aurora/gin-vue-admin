"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.dynamicImportPolyfill = void 0;
function dynamicImportPolyfill(promise = false) {
    return {
        name: 'dynamic-import-polyfill',
        renderDynamicImport() {
            return {
                left: promise ? 'Promise.resolve(' : '(',
                right: ')',
            };
        },
    };
}
exports.dynamicImportPolyfill = dynamicImportPolyfill;
