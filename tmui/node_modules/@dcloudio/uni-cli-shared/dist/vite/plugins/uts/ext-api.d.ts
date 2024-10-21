import type { Plugin } from 'vite';
import { type Injects, parseUniExtApis } from '../../../uni_modules';
export declare const parseUniExtApisOnce: typeof parseUniExtApis;
export declare function uniUTSExtApiReplace(): Plugin;
/**
 * { 'uni.getBatteryInfo': ['@/uni_modules/uni-getbatteryinfo/utssdk/web/index.uts','getBatteryInfo'] }
 * { '@/uni_modules/uni-getbatteryinfo/utssdk/web/index.ts': [['getBatteryInfo', 'uni_getBatteryInfo']] }
 * @param injects
 */
export declare function injectsToAutoImports(injects: Injects): {
    from: string;
    imports: [string, string][];
}[];
