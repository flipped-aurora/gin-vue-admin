import { type BuildOptions, type ServerOptions } from 'vite';
import type { RollupWatcher } from 'rollup';
import type { CliOptions } from '.';
export declare function initUVueEnv(): void;
export declare function runUVueAndroidDev(options: CliOptions & ServerOptions): Promise<undefined>;
export declare function runUVueAndroidBuild(options: CliOptions & BuildOptions): Promise<void>;
/**
 * 目前的简易实现逻辑
 * node层：
 *  1. 监听项目，生成资源到临时目录 .uts/android
 *  2. uvue 文件，做解析，拆分生成 render.kt, css.kt, uts.uvue
 *  3. static 文件，copy 到最终目录
 *  4. uvue、vue、uts 文件发生变化，调用 uts 编译器
 * @param options
 */
export declare function buildUVue(options: CliOptions): Promise<RollupWatcher | void>;
