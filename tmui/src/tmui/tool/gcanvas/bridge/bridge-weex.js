const isWeex = typeof WXEnvironment !== 'undefined';
const isWeexIOS = isWeex && /ios/i.test(WXEnvironment.platform);
const isWeexAndroid = isWeex && !isWeexIOS;

import GLmethod from '../context-webgl/GLmethod';

const GCanvasModule =
    (typeof weex !== 'undefined' && weex.requireModule) ? (weex.requireModule('gcanvas')) :
        (typeof __weex_require__ !== 'undefined') ? (__weex_require__('@weex-module/gcanvas')) : {};

let isDebugging = false;

let isComboDisabled = false;

const logCommand = (function () {
    const methodQuery = [];
    Object.keys(GLmethod).forEach(key => {
        methodQuery[GLmethod[key]] = key;
    })
    const queryMethod = (id) => {
        return methodQuery[parseInt(id)] || 'NotFoundMethod';
    }
    const logCommand = (id, cmds) => {
        const mId = cmds.split(',')[0];
        const mName = queryMethod(mId);
        console.log(`=== callNative - componentId:${id}; method: ${mName}; cmds: ${cmds}`);
    }
    return logCommand;
})();

function joinArray(arr, sep) {
    let res = '';
    for (let i = 0; i < arr.length; i++) {
        if (i !== 0) {
            res += sep;
        }
        res += arr[i];
    }
    return res;
}

const commandsCache = {}

const GBridge = {

    callEnable: (ref, configArray) => {

        commandsCache[ref] = [];

        return GCanvasModule.enable({
            componentId: ref,
            config: configArray
        });
    },

    callEnableDebug: () => {
        isDebugging = true;
    },

    callEnableDisableCombo: () => {
        isComboDisabled = true;
    },

    callSetContextType: function (componentId, context_type) {
        GCanvasModule.setContextType(context_type, componentId);
    },

    callReset: function(id){
        GCanvasModule.resetComponent && canvasModule.resetComponent(componentId);
    },

    render: isWeexIOS ? function (componentId) {
        return GCanvasModule.extendCallNative({
            contextId: componentId,
            type: 0x60000001
        });
    } : function (componentId) {
        return callGCanvasLinkNative(componentId, 0x60000001, 'render');
    },

    render2d: isWeexIOS ? function (componentId, commands, callback) {

        if (isDebugging) {
            console.log('>>> >>> render2d ===');
            console.log('>>> commands: ' + commands);
        }
		
        GCanvasModule.render([commands, callback?true:false], componentId, callback);

    } : function (componentId, commands,callback) {

        if (isDebugging) {
            console.log('>>> >>> render2d ===');
            console.log('>>> commands: ' + commands);
        }

        callGCanvasLinkNative(componentId, 0x20000001, commands);
		if(callback){
		callback();
		}
    },

    callExtendCallNative: isWeexIOS ? function (componentId, cmdArgs) {

        throw 'should not be here anymore ' + cmdArgs;

    } : function (componentId, cmdArgs) {

        throw 'should not be here anymore ' + cmdArgs;

    },


    flushNative: isWeexIOS ? function (componentId) {

        const cmdArgs = joinArray(commandsCache[componentId], ';');
        commandsCache[componentId] = [];

        if (isDebugging) {
            console.log('>>> >>> flush native ===');
            console.log('>>> commands: ' + cmdArgs);
        }

        const result = GCanvasModule.extendCallNative({
            "contextId": componentId,
            "type": 0x60000000,
            "args": cmdArgs
        });

        const res = result && result.result;

        if (isDebugging) {
            console.log('>>> result: ' + res);
        }

        return res;

    } : function (componentId) {

        const cmdArgs = joinArray(commandsCache[componentId], ';');
        commandsCache[componentId] = [];

        if (isDebugging) {
            console.log('>>> >>> flush native ===');
            console.log('>>> commands: ' + cmdArgs);
        }

        const result = callGCanvasLinkNative(componentId, 0x60000000, cmdArgs);

        if (isDebugging) {
            console.log('>>> result: ' + result);
        }

        return result;
    },

    callNative: function (componentId, cmdArgs, cache) {

        if (isDebugging) {
            logCommand(componentId, cmdArgs);
        }

        commandsCache[componentId].push(cmdArgs);

        if (!cache || isComboDisabled) {
            return GBridge.flushNative(componentId);
        } else {
            return undefined;
        }
    },

    texImage2D(componentId, ...args) {
        if (isWeexIOS) {
            if (args.length === 6) {
                const [target, level, internalformat, format, type, image] = args;
                GBridge.callNative(
                    componentId,
                    GLmethod.texImage2D + ',' + 6 + ',' + target + ',' + level + ',' + internalformat + ',' + format + ',' + type + ',' + image.src
                )
            } else if (args.length === 9) {
                const [target, level, internalformat, width, height, border, format, type, image] = args;
                GBridge.callNative(
                    componentId,
                    GLmethod.texImage2D + ',' + 9 + ',' + target + ',' + level + ',' + internalformat + ',' + width + ',' + height + ',' + border + ',' +
                    + format + ',' + type + ',' + (image ? image.src : 0)
                )
            }
        } else if (isWeexAndroid) {
            if (args.length === 6) {
                const [target, level, internalformat, format, type, image] = args;
                GCanvasModule.texImage2D(componentId, target, level, internalformat, format, type, image.src);
            } else if (args.length === 9) {
                const [target, level, internalformat, width, height, border, format, type, image] = args;
                GCanvasModule.texImage2D(componentId, target, level, internalformat, width, height, border, format, type, (image ? image.src : 0));
            }
        }
    },

    texSubImage2D(componentId, target, level, xoffset, yoffset, format, type, image) {
        if (isWeexIOS) {
            if (arguments.length === 8) {
                GBridge.callNative(
                    componentId,
                    GLmethod.texSubImage2D + ',' + 6 + ',' + target + ',' + level + ',' + xoffset + ',' + yoffset, + ',' + format + ',' + type + ',' + image.src
                )
            }
        } else if (isWeexAndroid) {
            GCanvasModule.texSubImage2D(componentId, target, level, xoffset, yoffset, format, type, image.src);
        }
    },

    bindImageTexture(componentId, src, imageId) {
        GCanvasModule.bindImageTexture([src, imageId], componentId);
    },

    perloadImage([url, id], callback) {
        GCanvasModule.preLoadImage([url, id], function (image) {
            image.url = url;
            image.id = id;
            callback(image);
        });
    },
	
	measureText(text, fontStyle, componentId) {
	    return GCanvasModule.measureText([text, fontStyle], componentId);
	},
	
	getImageData (componentId, x, y, w, h, callback) {
		GCanvasModule.getImageData([x, y,w,h],componentId,callback);
	},
	
	putImageData (componentId, data, x, y, w, h, callback) {
		GCanvasModule.putImageData([x, y,w,h,data],componentId,callback);
	},
	
	toTempFilePath(componentId, x, y, width, height, destWidth, destHeight, fileType, quality, callback){ 
		GCanvasModule.toTempFilePath([x, y, width,height, destWidth, destHeight, fileType, quality], componentId, callback);
	}
}

export default GBridge;