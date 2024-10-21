/* global HTMLImageElement */
/* global HTMLCanvasElement */
/* global SVGElement */

import getOptionsFromElement from "./getOptionsFromElement.js";
import renderers from "../renderers";

import { InvalidElementException } from "../exceptions/exceptions.js";

// Takes an element and returns an object with information about how
// it should be rendered
// This could also return an array with these objects
// {
//   element: The element that the renderer should draw on
//   renderer: The name of the renderer
//   afterRender (optional): If something has to done after the renderer
//     completed, calls afterRender (function)
//   options (optional): Options that can be defined in the element
// }

function getRenderProperties(element) {
	// If the element is a string, query select call again
	if (typeof element === "string") {
		return querySelectedRenderProperties(element);
	}
	// If element is array. Recursivly call with every object in the array
	else if (Array.isArray(element)) {
		var returnArray = [];
		for (let i = 0; i < element.length; i++) {
			returnArray.push(getRenderProperties(element[i]));
		}
		return returnArray;
	}
	// If element, render on canvas and set the uri as src
	else if (typeof HTMLCanvasElement !== 'undefined' && element instanceof HTMLImageElement) {
		return newCanvasRenderProperties(element);
	}
	// If SVG
	else if (
		(element && element.nodeName && element.nodeName.toLowerCase() === 'svg') ||
		(typeof SVGElement !== 'undefined' && element instanceof SVGElement)
	) {
		return {
			element: element,
			options: getOptionsFromElement(element),
			renderer: renderers.SVGRenderer
		};
	}
	// If canvas (in browser)
	else if (typeof HTMLCanvasElement !== 'undefined' && element instanceof HTMLCanvasElement) {
		return {
			element: element,
			options: getOptionsFromElement(element),
			renderer: renderers.CanvasRenderer
		};
	}
	// If canvas (in node)
	else if (element && element.getContext) {
		return {
			element: element,
			renderer: renderers.CanvasRenderer
		};
	}
	else if (element && typeof element === 'object' && !element.nodeName) {
		return {
			element: element,
			renderer: renderers.ObjectRenderer
		};
	}
	else {
		throw new InvalidElementException();
	}
}

function querySelectedRenderProperties(string) {
	var selector = document.querySelectorAll(string);
	if (selector.length === 0) {
		return undefined;
	}
	else {
		let returnArray = [];
		for (let i = 0; i < selector.length; i++) {
			returnArray.push(getRenderProperties(selector[i]));
		}
		return returnArray;
	}
}


function newCanvasRenderProperties(imgElement) {
	var canvas = document.createElement('canvas');
	return {
		element: canvas,
		options: getOptionsFromElement(imgElement),
		renderer: renderers.CanvasRenderer,
		afterRender: function () {
			imgElement.setAttribute("src", canvas.toDataURL());
		}
	};
}

export default getRenderProperties;
