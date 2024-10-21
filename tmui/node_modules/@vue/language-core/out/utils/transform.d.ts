import type * as ts from 'typescript/lib/tsserverlibrary';
import { VueCompilerOptions } from '../types';
export declare function walkInterpolationFragment(ts: typeof import('typescript/lib/tsserverlibrary'), code: string, ast: ts.SourceFile, cb: (fragment: string, offset: number | undefined, isJustForErrorMapping?: boolean) => void, localVars: Map<string, number>, identifiers: Set<string>, vueOptions: VueCompilerOptions): {
    text: string;
    isShorthand: boolean;
    offset: number;
}[];
export declare function collectVars(ts: typeof import('typescript/lib/tsserverlibrary'), node: ts.Node, result: string[]): void;
//# sourceMappingURL=transform.d.ts.map