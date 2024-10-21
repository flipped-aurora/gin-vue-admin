import type { Plugin, ResolveFn } from 'vite';
import type { CssUrlReplacer } from './plugins/vitejs/plugins/css';
export declare function createEncryptCssUrlReplacer(resolve: ResolveFn): CssUrlReplacer;
export declare function uniEncryptUniModulesAssetsPlugin(): Plugin;
export declare function uniEncryptUniModulesPlugin(): Plugin;
export declare function addUniModulesExtApiComponents(relativeFilename: string, components: string[]): void;
