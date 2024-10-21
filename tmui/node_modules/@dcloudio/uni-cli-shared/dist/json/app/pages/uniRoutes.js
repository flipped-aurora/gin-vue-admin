"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.normalizeAppUniRoutes = void 0;
const pages_1 = require("../../pages");
function normalizeAppUniRoutes(pagesJson) {
    return JSON.stringify((0, pages_1.normalizePagesRoute)(pagesJson));
}
exports.normalizeAppUniRoutes = normalizeAppUniRoutes;
