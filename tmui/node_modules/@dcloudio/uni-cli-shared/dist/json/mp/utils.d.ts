import type { AppWindowOptions, PageWindowOptions, TabBar } from './types';
export declare function parseWindowOptions(style: UniApp.PagesJsonPageStyle, platform: UniApp.PLATFORM, windowOptionsMap?: Record<string, string>): PageWindowOptions | AppWindowOptions;
export declare function parseTabBar(tabBar: UniApp.TabBarOptions, platform: UniApp.PLATFORM, tabBarOptionsMap?: Record<string, string>, tabBarItemOptionsMap?: Record<string, string>): TabBar;
