"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.normalizeAppConfigService = exports.normalizeAppNVuePagesJson = exports.normalizeAppPagesJson = void 0;
const code_1 = require("./code");
const definePage_1 = require("./definePage");
const uniConfig_1 = require("./uniConfig");
const uniRoutes_1 = require("./uniRoutes");
function normalizeAppPagesJson(pagesJson, platform = 'app') {
    return (0, definePage_1.definePageCode)(pagesJson, platform);
}
exports.normalizeAppPagesJson = normalizeAppPagesJson;
function normalizeAppNVuePagesJson(pagesJson) {
    return (0, definePage_1.defineNVuePageCode)(pagesJson);
}
exports.normalizeAppNVuePagesJson = normalizeAppNVuePagesJson;
function normalizeAppConfigService(pagesJson, manifestJson) {
    return `
  ;(function(){
  let u=void 0,isReady=false,onReadyCallbacks=[],isServiceReady=false,onServiceReadyCallbacks=[];
  const __uniConfig = ${(0, uniConfig_1.normalizeAppUniConfig)(pagesJson, manifestJson)};
  const __uniRoutes = ${(0, uniRoutes_1.normalizeAppUniRoutes)(pagesJson)}.map(uniRoute=>(uniRoute.meta.route=uniRoute.path,__uniConfig.pages.push(uniRoute.path),uniRoute.path='/'+uniRoute.path,uniRoute));
  __uniConfig.styles=${process.env.UNI_NVUE_APP_STYLES || '[]'};//styles
  __uniConfig.onReady=function(callback){if(__uniConfig.ready){callback()}else{onReadyCallbacks.push(callback)}};Object.defineProperty(__uniConfig,"ready",{get:function(){return isReady},set:function(val){isReady=val;if(!isReady){return}const callbacks=onReadyCallbacks.slice(0);onReadyCallbacks.length=0;callbacks.forEach(function(callback){callback()})}});
  __uniConfig.onServiceReady=function(callback){if(__uniConfig.serviceReady){callback()}else{onServiceReadyCallbacks.push(callback)}};Object.defineProperty(__uniConfig,"serviceReady",{get:function(){return isServiceReady},set:function(val){isServiceReady=val;if(!isServiceReady){return}const callbacks=onServiceReadyCallbacks.slice(0);onServiceReadyCallbacks.length=0;callbacks.forEach(function(callback){callback()})}});
  service.register("uni-app-config",{create(a,b,c){if(!__uniConfig.viewport){var d=b.weex.config.env.scale,e=b.weex.config.env.deviceWidth,f=Math.ceil(e/d);Object.assign(__uniConfig,{viewport:f,defaultFontSize:16})}return{instance:{__uniConfig:__uniConfig,__uniRoutes:__uniRoutes,${code_1.globalCode}}}}}); 
  })();
  `;
}
exports.normalizeAppConfigService = normalizeAppConfigService;
