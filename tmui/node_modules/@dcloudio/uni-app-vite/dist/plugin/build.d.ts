import type { ConfigEnv, UserConfig } from 'vite';
export declare function buildOptions({ appService, renderer, }: {
    renderer: 'native' | undefined;
    appService: boolean;
}, userConfig: UserConfig, _: ConfigEnv): UserConfig['build'];
