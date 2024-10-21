"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.COMPONENT_DEPS_CSS = exports.API_DEPS_CSS = void 0;
const constants_1 = require("./constants");
const RESIZE_SENSOR_CSS = constants_1.BASE_COMPONENTS_STYLE_PATH + 'resize-sensor.css';
const REFRESHER_CSS = constants_1.BASE_COMPONENTS_STYLE_PATH + 'refresher.css';
exports.API_DEPS_CSS = {
    showModal: [`${constants_1.H5_API_STYLE_PATH}modal.css`],
    showToast: [`${constants_1.H5_API_STYLE_PATH}toast.css`],
    showActionSheet: [`${constants_1.H5_API_STYLE_PATH}action-sheet.css`],
    previewImage: [
        RESIZE_SENSOR_CSS,
        `${constants_1.BASE_COMPONENTS_STYLE_PATH}swiper.css`,
        `${constants_1.BASE_COMPONENTS_STYLE_PATH}swiper-item.css`,
        `${constants_1.BASE_COMPONENTS_STYLE_PATH}movable-area.css`,
        `${constants_1.BASE_COMPONENTS_STYLE_PATH}movable-view.css`,
    ],
    openLocation: [`${constants_1.H5_API_STYLE_PATH}location-view.css`],
    chooseLocation: [
        `${constants_1.H5_API_STYLE_PATH}/location-picker.css`,
        `${constants_1.BASE_COMPONENTS_STYLE_PATH}/input.css`,
        `${constants_1.H5_COMPONENTS_STYLE_PATH}/map.css`,
        `${constants_1.BASE_COMPONENTS_STYLE_PATH}/scroll-view.css`,
    ],
};
exports.COMPONENT_DEPS_CSS = {
    canvas: [RESIZE_SENSOR_CSS],
    image: [RESIZE_SENSOR_CSS],
    'movable-area': [RESIZE_SENSOR_CSS],
    'picker-view': [RESIZE_SENSOR_CSS],
    'picker-view-column': [RESIZE_SENSOR_CSS],
    'rich-text': [RESIZE_SENSOR_CSS],
    textarea: [RESIZE_SENSOR_CSS],
    'web-view': [RESIZE_SENSOR_CSS],
    picker: [
        RESIZE_SENSOR_CSS,
        `${constants_1.BASE_COMPONENTS_STYLE_PATH}picker-view.css`,
        `${constants_1.BASE_COMPONENTS_STYLE_PATH}picker-view-column.css`,
    ],
    'scroll-view': [REFRESHER_CSS],
    'list-view': [RESIZE_SENSOR_CSS, REFRESHER_CSS],
};
