import { type Program } from '@babel/types';
export declare function hasExternalClasses(code: string): boolean;
export declare function findMiniProgramComponentExternalClasses(filename: string): string[] | undefined;
export declare function updateMiniProgramComponentExternalClasses(filename: string, classes: string[]): void;
export declare function parseExternalClasses(ast: Program): string[];
