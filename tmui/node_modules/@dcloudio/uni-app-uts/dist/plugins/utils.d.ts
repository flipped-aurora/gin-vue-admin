import { type UniVitePlugin } from '@dcloudio/uni-cli-shared';
import type { ResolvedConfig } from 'vite';
export declare function createUniOptions(platform: 'android' | 'ios'): UniVitePlugin['uni'];
export declare function isManifest(id: string): boolean;
export declare function isPages(id: string): boolean;
export declare function configResolved(config: ResolvedConfig, isAndroidX?: boolean): void;
export declare function relativeInputDir(filename: string): string;
export declare function normalizeManifestJson(userManifestJson: Record<string, any>): {
    id: any;
    name: any;
    description: any;
    version: {
        name: any;
        code: any;
    };
    'uni-app-x': any;
    app: any;
};
export declare function updateManifestModules(manifest: Record<string, any>, modules: string[]): Record<string, any>;
export declare function setGlobalPageOrientation(value: string): void;
export declare function getGlobalPageOrientation(): string;
export declare function addExtApiComponents(components: string[]): void;
export declare function getExtApiComponents(): Set<string>;
