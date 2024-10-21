"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniMovePlugin = void 0;
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const fast_glob_1 = __importDefault(require("fast-glob"));
function uniMovePlugin({ apply, enforce, cwd, pattern, dest, }) {
    return {
        name: 'uni:move',
        apply,
        enforce,
        async writeBundle() {
            await Promise.all(fast_glob_1.default
                .sync(pattern, {
                cwd,
            })
                .map((filename) => {
                return fs_extra_1.default.move(path_1.default.resolve(cwd, filename), path_1.default.resolve(dest, filename), {
                    overwrite: true,
                });
            }));
        },
    };
}
exports.uniMovePlugin = uniMovePlugin;
