import type { Plugin } from 'vite';
export declare function uniEsbuildPlugin({ appService, }: {
    renderer?: 'native';
    appService: boolean;
}): Plugin;
export declare function wrapperNVueAppStyles(code: string): string;
