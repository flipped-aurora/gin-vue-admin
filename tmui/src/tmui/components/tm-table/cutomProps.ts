import { PropType } from "vue"
import type { headresItem, cellItem, dataTypeArray } from "./interface";
import type { tableDataType } from "./newInterface";
export const cutomProps = {
    /**是否显示表格头 */
    showHeader: {
        type: Boolean,
        default: true,
    },
    /**表格数据 */
    tableData: {
        type: Object as PropType<tableDataType>,
        default: (): tableDataType => {
            return {
                data: [],
                fields: {
                    columns: []
                },
                header: []
            }
        },
    },
    /**宽度 */
    width: {
        type: Number,
        default: 750,
    },
    /**表格单元格的宽度  */
    cellWidth: {
        type: Number,
        default: 160,
    },
    /**如果提供了高度，将产生上下滑动的表格。 */
    height: {
        type: Number,
        default: 0,
    },
    /**单元格的高度。 */
    cellHeight: {
        type: Number,
        default: 72,
    },
    /**头部的高度。 */
    headerHeight: {
        type: Number,
        default: 100,
    },
    /**是否显示格格下划线 */
    showBottomBorder: {
        type: Boolean,
        default: true,
    },
    /**是否显示固定列 */
    showFixed: {
        type: Boolean,
        default: false,
    },
    /**开户间隔条纹 */
    stripe: {
        type: Boolean,
        default: false,
    }
}