"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.stripOptions = void 0;
exports.stripOptions = {
    include: ['**/*.js', '**/*.ts', '**/*.tsx', '**/*.vue'],
    functions: [
        'onBeforeMount',
        'onMounted',
        'onBeforeUpdate',
        'onUpdated',
        'onActivated',
        'onDeactivated',
        'onBeforeActivate',
        'onBeforeDeactivate',
        'onBeforeUnmount',
        'onUnmounted',
        'onRenderTracked',
        'onRenderTriggered',
    ],
};
