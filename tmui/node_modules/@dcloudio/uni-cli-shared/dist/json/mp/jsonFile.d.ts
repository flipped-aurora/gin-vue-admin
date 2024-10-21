import type { ComponentJson, MiniProgramComponentsType, PageWindowOptions, UsingComponents } from './types';
export declare function isMiniProgramPageFile(file: string, inputDir?: string): boolean;
export declare function isMiniProgramPageSfcFile(file: string, inputDir?: string): boolean;
export declare function hasJsonFile(filename: string): boolean;
export declare function getComponentJsonFilenames(): string[];
export declare function findJsonFile(filename: string): Record<string, any> | ComponentJson | PageWindowOptions | undefined;
export declare function findUsingComponents(filename: string): UsingComponents | undefined;
export declare function normalizeJsonFilename(filename: string): string;
export declare function findChangedJsonFiles(supportGlobalUsingComponents?: boolean): Map<string, string>;
export declare function addMiniProgramAppJson(appJson: Record<string, any>): void;
export declare function addMiniProgramPageJson(filename: string, json: PageWindowOptions): void;
export declare function addMiniProgramComponentJson(filename: string, json: ComponentJson): void;
export declare function addMiniProgramUsingComponents(filename: string, json: UsingComponents): void;
export declare function isMiniProgramUsingComponent(name: string, options: {
    filename: string;
    inputDir: string;
    componentsDir?: string;
}): boolean;
interface MiniProgramComponents {
    [name: string]: MiniProgramComponentsType;
}
export declare function findMiniProgramUsingComponents({ filename, inputDir, componentsDir, }: {
    filename: string;
    inputDir: string;
    componentsDir?: string;
}): MiniProgramComponents;
export declare function findUsingComponentsJson(pathInpages: string, componentsDir: string): Record<any, any>;
export {};
