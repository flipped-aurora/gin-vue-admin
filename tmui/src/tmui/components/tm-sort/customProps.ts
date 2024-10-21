import { PropType } from "vue"
import { itemType } from "./interface"
export default {
    width: {
        type: Number,
        default: 750
    },
    height: {
        type: Number,
        default: 174
    },
    /**项目的高度 */
    itemHeight: {
        type: Number,
        default: 88
    },
    /**一排几个项目 */
    column: {
        type: Number,
        default: 3
    },
    color: {
        type: String,
        default: 'white'
    },
    fontSize: {
        type: Number,
        default: 24
    },
    iconSize: {
        type: Number,
        default: 36
    },
    list: {
        type: Array as PropType<Array<itemType>>,
        default: () => []
    }
}