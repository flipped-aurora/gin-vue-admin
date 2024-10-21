import type { Plugin, ResolvedConfig } from 'vite';
import { type MiniProgramFilterOptions } from '@dcloudio/uni-cli-shared';
export declare function getFiltersCache(resolvedConfig: ResolvedConfig): MiniProgramFilterOptions[];
export declare function uniRenderjsPlugin({ lang }: {
    lang?: string;
}): Plugin;
