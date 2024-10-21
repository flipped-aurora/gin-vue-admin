import type { EmittedFile, GetModuleInfo } from 'rollup';
import type { ResolvedConfig } from 'vite';
import { type MiniProgramFilterOptions } from '@dcloudio/uni-cli-shared';
import type { UniMiniProgramPluginOptions } from '.';
export declare function getFilterFiles(resolvedConfig: ResolvedConfig, getModuleInfo: GetModuleInfo): Record<string, MiniProgramFilterOptions>;
export declare function getTemplateFiles(template: UniMiniProgramPluginOptions['template']): Record<string, string>;
export declare const emitFile: (emittedFile: EmittedFile) => string;
