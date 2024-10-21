export interface ComponentJson {
    component: true;
    usingComponents?: UsingComponents;
    usingSwanComponents?: UsingComponents;
    styleIsolation?: 'apply-shared' | 'shared' | 'isolated';
}
interface ShareWindowOptions {
    navigationBarBackgroundColor?: string;
    navigationBarTextStyle?: 'white' | 'black';
    navigationBarTitleText?: string;
    navigationStyle?: 'default' | 'custom';
    backgroundColor?: string;
    backgroundTextStyle?: 'dark' | 'light';
    backgroundColorTop?: string;
    backgroundColorBottom?: string;
    enablePullDownRefresh?: boolean;
    onReachBottomDistance?: number;
    pageOrientation?: 'portrait' | 'landscape' | 'auto';
}
type Style = 'v2' | string;
type RestartStrategy = 'homePage' | 'homePageAndLatestPage' | string;
export interface PageWindowOptions extends ShareWindowOptions {
    component?: true;
    disableScroll?: boolean;
    usingComponents?: UsingComponents;
    usingSwanComponents?: UsingComponents;
    initialRenderingCache?: 'static' | string;
    style?: Style;
    singlePage?: SinglePage;
    restartStrategy?: RestartStrategy;
}
export interface AppWindowOptions extends ShareWindowOptions {
    visualEffectInBackground?: 'none' | 'hidden';
}
interface SubPackage {
    name?: string;
    root: string;
    pages: string[];
    independent?: boolean;
}
interface TabBarItem {
    pagePath: string;
    text: string;
    iconPath?: string;
    selectedIconPath?: string;
}
export interface TabBar {
    color: string;
    selectedColor: string;
    backgroundColor: string;
    borderStyle?: 'black' | 'white';
    list: TabBarItem[];
    position?: 'bottom' | 'top';
    custom?: boolean;
}
export interface NetworkTimeout {
    request?: number;
    requeconnectSocketst?: number;
    uploadFile?: number;
    downloadFile?: number;
}
interface Plugins {
    [name: string]: {
        version: string;
        provider: string;
    };
}
interface PreloadRule {
    [name: string]: {
        network: 'wifi' | 'all';
        packages: string[];
    };
}
export interface UsingComponents {
    [name: string]: string;
}
interface Permission {
    [name: string]: {
        desc: string;
    };
}
interface UseExtendedLib {
    kbone: boolean;
    weui: boolean;
}
interface EntranceDeclare {
    locationMessage: {
        path: string;
        query: string;
    };
}
interface SinglePage {
    navigationBarFit?: 'squeezed' | 'float';
}
export interface AppJson {
    entryPagePath?: string;
    pages: string[];
    window?: AppWindowOptions;
    tabBar?: TabBar;
    networkTimeout?: NetworkTimeout;
    debug?: boolean;
    functionalPages?: boolean;
    subPackages?: SubPackage[];
    workers?: string;
    requiredBackgroundModes?: string[];
    plugins?: Plugins;
    preloadRule?: PreloadRule;
    resizable?: boolean;
    usingComponents?: UsingComponents;
    permission?: Permission;
    sitemapLocation?: string;
    style?: Style;
    useExtendedLib?: UseExtendedLib;
    entranceDeclare?: EntranceDeclare;
    darkmode?: boolean;
    themeLocation?: string;
    lazyCodeLoading?: 'requiredComponents' | string;
    singlePage?: SinglePage;
    restartStrategy?: RestartStrategy;
    [name: string]: unknown;
}
export type MiniProgramComponentsType = 'plugin' | 'component' | 'dynamicLib' | 'ext' | 'xr-frame';
export {};
