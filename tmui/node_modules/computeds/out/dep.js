"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Dep = void 0;
class Dep extends Map {
    constructor(computed) {
        super();
        this.computed = computed;
    }
}
exports.Dep = Dep;
