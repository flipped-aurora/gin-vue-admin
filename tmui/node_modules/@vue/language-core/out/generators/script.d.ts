import { FileRangeCapabilities, MirrorBehaviorCapabilities } from '@volar/language-core';
import * as SourceMaps from '@volar/source-map';
import type * as ts from 'typescript/lib/tsserverlibrary';
import type * as templateGen from '../generators/template';
import type { ScriptRanges } from '../parsers/scriptRanges';
import type { ScriptSetupRanges } from '../parsers/scriptSetupRanges';
import type { VueCompilerOptions } from '../types';
import { Sfc } from '../types';
export declare function generate(ts: typeof import('typescript/lib/tsserverlibrary'), fileName: string, script: Sfc['script'], scriptSetup: Sfc['scriptSetup'], styles: Sfc['styles'], // TODO: computed it
lang: string, scriptRanges: ScriptRanges | undefined, scriptSetupRanges: ScriptSetupRanges | undefined, htmlGen: ReturnType<typeof templateGen['generate']> | undefined, compilerOptions: ts.CompilerOptions, vueCompilerOptions: VueCompilerOptions, codegenStack: boolean): {
    codes: SourceMaps.Segment<FileRangeCapabilities>[];
    codeStacks: SourceMaps.StackNode[];
    mirrorBehaviorMappings: SourceMaps.Mapping<[MirrorBehaviorCapabilities, MirrorBehaviorCapabilities]>[];
};
//# sourceMappingURL=script.d.ts.map