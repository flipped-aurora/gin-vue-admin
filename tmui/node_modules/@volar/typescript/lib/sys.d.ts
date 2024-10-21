import type { ServiceEnvironment, Disposable } from '@volar/language-service';
import type * as ts from 'typescript/lib/tsserverlibrary';
export declare function createSys(ts: typeof import('typescript/lib/tsserverlibrary'), env: ServiceEnvironment): ts.System & {
    version: number;
    sync(): Promise<number>;
} & Disposable;
//# sourceMappingURL=sys.d.ts.map