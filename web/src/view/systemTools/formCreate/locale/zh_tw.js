export default {
    name: "zh-tw",
    form: {
        field: "字段 ID",
        title: "字段名稱",
        info: "提示信息",
        control: "聯動數據",
        labelPosition: "標籤的位置",
        labelStyle: "標籤的樣式",
        labelSuffix: "標籤的後綴",
        size: "表單的尺寸",
        event: "表單事件",
        labelWidth: "標籤的寬度",
        hideRequiredAsterisk: "隱藏必填字段的標籤旁邊的紅色星號",
        showMessage: "顯示校驗錯誤信息",
        inlineMessage: "以行內形式展示校驗信息",
        submitBtn: "是否顯示表單提交按鈕",
        resetBtn: "是否顯示表單重置按鈕",
        appendChild: "添加子級",
        formMode: "表單模式",
        formName: "表單名稱",
        componentMode: "生成組件",
        document: "幫助文檔"
    },
    computed: {
        fieldUsed: "【{label}】在計算公式中被使用，請先修改對應公式",
        fieldExist: "【{label}】字段已存在",
        fieldEmpty: "字段名稱不能為空",
        fieldChar: "字段名稱必須以字母開頭"
    },
    validate: {
        type: "字段類型",
        typePlaceholder: "請選擇",
        trigger: "觸發方式",
        mode: "驗證方式",
        modes: {
            required: "必填",
            pattern: "正則表達式",
            validator: "自定義驗證",
            min: "最小值",
            max: "最大值",
            len: "長度"
        },
        types: {
            string: "字符串",
            array: "多選",
            number: "數字",
            integer: "整數",
            float: "小數",
            object: "合集",
            date: "日期",
            url: "URL鏈接",
            email: "郵箱地址"
        },
        message: "錯誤信息",
        auto: "自動獲取",
        autoRequired: "請輸入{title}",
        autoMode: "請輸入正確的{title}",
        requiredPlaceholder: "請輸入提示語",
        required: "是否必填",
        rule: "驗證規則"
    },
    tableOptions: {
        handle: "操作",
        add: "添加",
        empty1: "點擊右下角",
        empty2: "按鈕添加一列",
        rmCol: "刪除當前列",
        rmRow: "刪除當前行",
        splitRow: "拆分成行",
        splitCol: "拆分成列",
        mergeBottom: "向下合併",
        mergeRight: "向右合併",
        addTop: "添加上列",
        addBottom: "添加下列",
        addLeft: "添加左列",
        addRight: "添加右列",
        keyValue: "鍵值對"
    },
    struct: {
        title: "編輯數據",
        only: "【{label}】只允許添加一個",
        errorMsg: "輸入的內容語法錯誤",
        configured: "已配置"
    },
    event: {
        title: "設置事件",
        create: "創建事件",
        list: "事件列表",
        placeholder: "請輸入事件的名稱",
        saveMsg: "請先保存當前正在編輯的事件",
        type: "類型",
        info: "說明",
        label: "字段",
        inject: {
            api: "當前表單的api",
            rule: "當前表單的生成規則",
            self: "組件的生成規則",
            option: "表單的配置",
            args: "事件的原始參數"
        }
    },
    fetch: {
        title: "設置數據源",
        create: "創建數據源",
        config: "請求配置",
        action: "請求鏈接",
        actionRequired: "請輸入正確的鏈接",
        placeholder: "請輸入數據源的名稱",
        method: "請求方式",
        data: "附帶數據",
        headers: "請求頭部",
        parse: "數據處理",
        response: "接口返回的數據",
        onError: "錯誤處理",
        remote: "遠程數據",
        static: "靜態數據",
        optionsType: {
            fetch: "遠程數據",
            struct: "靜態數據"
        }
    },
    style: {
        width: "寬度",
        height: "高度",
        color: "顏色",
        backgroundColor: "背景色",
        margin: "外邊距",
        padding: "內邊距",
        borderRadius: "圓角",
        border: "邊框",
        solid: "實線",
        dashed: "虛線",
        dotted: "點狀虛線",
        double: "雙實線",
        opacity: "透明度",
        scale: "縮放",
        shadow: {
            name: "陰影",
            x: "x軸偏移量",
            y: "y軸偏移量",
            vague: "模糊半徑",
            extend: "擴散半徑",
            inset: "向內",
            external: "向外",
            mode: "模式",
            classic: "經典",
            flat: "扁平",
            solid: "立體"
        },
        font: {
            name: "字體",
            size: "大小",
            align: "對齊方式",
            height: "行高",
            spacing: "字間距",
            preview: "樣式預覽"
        },
        decoration: {
            name: "修飾",
            underline: "下劃線",
            "line-through": "刪除線",
            overline: "上劃線"
        },
        weight: {
            name: "粗細",
            300: "細體",
            400: "常規體",
            500: "���黑體",
            700: "中粗體"
        }
    },
    designer: {
        component: "組件配置",
        id: "唯一值",
        name: "編號",
        type: "組件類型",
        form: "表單配置",
        style: "組件樣式配置",
        rule: "基礎配置",
        advanced: "高級配置",
        props: "屬性配置",
        validate: "驗證配置",
        event: "事件配置",
        clearWarn: "清空後將不能恢復，確定要清空嗎？",
        childEmpty: "點擊右下角 \\e789 按鈕添加一列",
        dragEmpty: "拖拽左側列表中的組件到此處",
        unload: "確定離開當前頁面嗎?"
    },
    menu: {
        main: "基礎組件",
        aide: "輔助組件",
        layout: "佈局組件",
        component: "組件",
        subform: "子表單組件",
        tree: "大綱"
    },
    props: {
        disabled: "禁用",
        time: "時間",
        email: "郵箱",
        number: "數字",
        globalData: "全局數據",
        mobile: "移動端",
        reactive: "響應式",
        pc: "電腦端",
        title: "標題",
        content: "內容",
        collection: "合集",
        group: "分組",
        custom: "自定義",
        change: "改變",
        blur: "失去焦點",
        preview: "預覽",
        clear: "清空",
        cancel: "取消",
        close: "關閉",
        ok: "確定",
        save: "保存",
        refresh: "刷新",
        submit: "提交",
        reset: "重置",
        copy: "複製",
        delete: "刪除",
        hide: "隱藏",
        show: "顯示",
        position: "位置",
        render: "渲染",
        large: "大",
        default: "默認",
        small: "小",
        always: "常顯",
        never: "不顯示",
        hover: "懸浮",
        click: "點擊",
        button: "按鈕",
        year: "年份",
        month: "月份",
        date: "日期",
        dates: "日期多選",
        week: "一周",
        datetime: "日期時間",
        "datetime-local": "日期時間",
        datetimerange: "日期時間區間",
        daterange: "日期區間",
        monthrange: "月份區間",
        left: "左對齊",
        right: "右對齊",
        top: "頂部",
        text: "文字",
        picture: "圖片",
        "picture-card": "卡片",
        center: "居中",
        vertical: "豎向",
        horizontal: "橫向",
        manage: "管理",
        key: "鍵名",
        name: "名稱",
        value: "值",
        inputData: "默認值",
        append: "插入",
        options: "選項數據",
        option: "選項",
        callback: "回調",
        _self: "當前窗口",
        _blank: "新的窗口",
        _parent: "父級窗口",
        _top: "頂級窗口"
    },
    com: {
        cascader: {
            name: "級聯選擇器",
            event: {
                expandChange: "當展開節點發生變化時觸發",
                removeTag: "在多選模式下，移除Tag時觸發"
            },
            props: {
                props: "配置選項",
                placeholder: "輸入框佔位文本",
                disabled: "是否禁用",
                clearable: "是否支持清空選項",
                showAllLevels: "輸入框中是否顯示選中值的完整路徑",
                collapseTags: "多選模式下是否折疊Tag",
                collapseTagsTooltip: "當鼠標懸停於折疊標籤的文本時，是否顯示所有選中的標籤",
                separator: "選項分隔符",
                filterable: "該選項是否可以被搜索",
                tagType: "標籤類型"
            },
            propsOpt: {
                multiple: "是否多選",
                expandTrigger: "次級菜單的展開方式",
                checkStrictly: "是否嚴格的遵守父子節點不互相關聯",
                emitPath: "在選中節點改變時，是否返回由該節點所在的各級菜單的值所組成的數組",
                value: "指定選項的值為選項對象的某個屬性值",
                label: "指定選項標籤為選項對象的某個屬性值",
                children: "指定選項的子選項為選項對象的某個屬性值",
                disabled: "指定選項的禁用為選項對象的某個屬性值",
                leaf: "指定選項的葉子節點的標誌位為選項對象的某個屬性值"
            }
        },
        checkbox: {
            name: "多選框",
            props: {
                input: "是否可以填寫",
                type: "按鈕類型",
                disabled: "是否禁用",
                min: "可被選中的最小數量",
                max: "可被選中的最大數量",
                textColor: "當按鈕為活躍狀態時的字體顏色",
                fill: "當按鈕為活躍狀態時的邊框和背景顏色"
            }
        },
        col: {
            name: "佈局格子",
            props: {
                span: "柵格佔據的列數",
                offset: "柵格左側的間隔格數",
                push: "柵格向右移動格數",
                pull: "柵格向左移動格數"
            }
        },
        colorPicker: {
            name: "顏色選擇器",
            event: {
                activeChange: "面板中當前顯示的顏色發生變化時觸發"
            },
            props: {
                disabled: "是否禁用",
                showAlpha: "是否支持透明度選擇",
                colorFormat: "顏色的格式",
                predefine: "預定義顏色"
            }
        },
        datePicker: {
            name: "日期",
            props: {
                pickerOptions: "當前時間日期選擇器特有的選項",
                readonly: "完全只讀",
                disabled: "禁用",
                type: "顯示類型",
                editable: "文本框可輸入",
                clearable: "是否顯示清除按鈕",
                placeholder: "非範圍選擇時的佔位內容",
                startPlaceholder: "範圍選擇時開始日期的佔位內容",
                endPlaceholder: "範圍選擇時結束日期的佔位內容",
                format: "顯示在輸入框中的格式",
                align: "對齊方式",
                rangeSeparator: "選擇範圍時的分隔符",
                unlinkPanels: "在範圍選擇器裡取消兩個日期面板之間的聯動"
            }
        },
        dateRange: {
            name: "日期區間"
        },
        timeRange: {
            name: "時間區間"
        },
        elAlert: {
            name: "提示",
            description: "說明文字",
            props: {
                title: "標題",
                type: "主題",
                description: "輔助性文字",
                closable: "是否可關閉",
                center: "文字是否居中",
                closeText: "關閉按鈕自定義文本",
                showIcon: "是否顯示圖標",
                effect: "選擇提供的主題"
            }
        },
        elButton: {
            name: "按鈕",
            props: {
                formCreateChild: "內容",
                size: "尺寸",
                type: "類型",
                plain: "是否樸素按鈕",
                round: "是否圓角按鈕",
                circle: "是否圓形按鈕",
                loading: "是否加載中狀態",
                disabled: "是否禁用狀態"
            }
        },
        elCard: {
            name: "卡片",
            props: {
                header: "標題",
                shadow: "陰影顯示時機"
            }
        },
        elCollapse: {
            name: "摺疊面板",
            event: {
                change: "切換當前活躍面板，在手風琴模式下其類型是string，在其他模式下是array"
            },
            props: {
                accordion: "是否手風琴模式"
            }
        },
        elCollapseItem: {
            name: "面板",
            props: {
                title: "面板標題",
                name: "唯一標誌符",
                disabled: "是否禁用"
            }
        },
        elDivider: {
            name: "��割線",
            props: {
                formCreateChild: "設置分割線文案",
                contentPosition: "設置分割線文案的位置"
            }
        },
        elTabPane: {
            name: "選項卡",
            props: {
                label: "選項卡標題",
                disabled: "是否禁用",
                name: "選項卡的標識符",
                lazy: "標籤是否延遲渲染"
            }
        },
        elTabs: {
            name: "標籤頁",
            event: {
                tabClick: "tab 被選中時觸發",
                tabChange: "activeName 改變時觸發",
                tabRemove: "點擊 tab 移除按鈕時觸發",
                tabAdd: "點擊 tab ��增按鈕時觸發",
                edit: "點擊 tab 的新增或移除按鈕後觸發"
            },
            props: {
                type: "風格類型",
                closable: "標籤是否可關閉",
                tabPosition: "選項卡所在位置",
                stretch: "標籤的寬度是否自撐開"
            }
        },
        elTag: {
            name: "標籤",
            props: {
                formCreateNative: "是否顯示標題",
                formCreateTitle: "標題",
                formCreateChild: "標籤內容",
                type: "標籤的類型",
                size: "標籤的尺寸",
                effect: "標籤的主題",
                closable: "是否可關閉",
                disableTransitions: "是否禁用漸變動畫",
                hit: "是否有邊框描邊",
                round: "是否為圓形",
                color: "背景色"
            }
        },
        elTransfer: {
            name: "穿梭框",
            event: {
                leftCheckChange: "左側列表元素被用戶選中 / 取消選中時觸發",
                rightCheckChange: "右側列表元素被用戶選中 / 取消選中時觸發"
            },
            props: {
                filterable: "是否可搜索",
                filterPlaceholder: "搜索框佔位符",
                targetOrder: "右側列表元素的排序策略",
                targetOrderInfo: "若為 original，則保持與數據相同的順序；若為 push，則新加入的元素排在最後；若為 unshift，則新加入的元素排在最前",
                titles: "自定義列表標題",
                buttonTexts: "自定義按鈕文案",
                props: "數據源的字段別名"
            }
        },
        elTreeSelect: {
            name: "樹形選擇",
            event: {
                removeTag: "多選模式下移除tag時觸發"
            },
            props: {
                multiple: "是否多選",
                disabled: "是否禁用",
                clearable: "是否可以清空選項",
                collapseTags: "多選時是否將選中值按文字的形式展示",
                multipleLimit: "多選時用戶最多可以選擇的項目數，為 0 則不限制",
                placeholder: "佔位符",
                props: "配置選項",
                renderAfterExpand: "是否在第一次展開某個樹節點後才渲染其子節點",
                defaultExpandAll: "是否默認展開所有節點",
                expandOnClickNode: "是否在點擊節點的時候展開或者收縮節點",
                checkOnClickNode: "是否在點擊節點的時候選中節點",
                nodeKey: "每個樹節點用來作為唯一標識的屬性，整棵樹應該是唯一的"
            }
        },
        fcEditor: {
            name: "富文本框",
            props: {
                disabled: "是否禁用"
            }
        },
        fcRow: {
            name: "柵格佈局",
            props: {
                gutter: "柵格間隔",
                type: "flex佈局模式",
                justify: "flex佈局下的水平排列方式",
                align: "flex佈局下的垂���排列方式"
            }
        },
        fcTable: {
            name: "表格佈局",
            props: {
                border: "是否顯示邊框",
                borderColor: "邊框顏色",
                borderWidth: "邊框寬度"
            }
        },
        fcTableGrid: {
            name: "格子"
        },
        group: {
            name: "子表單",
            props: {
                disabled: "是否禁用",
                syncDisabled: "是否與子表單強制同步禁用狀態",
                expand: "設置默認展開幾項",
                button: "是否顯示操作按鈕",
                sortBtn: "是否顯示排序按鈕",
                min: "設置最小添加幾項",
                max: "設置最多添加幾項"
            }
        },
        html: {
            name: "HTML",
            props: {
                formCreateNative: "是否顯示標題",
                formCreateTitle: "標題",
                formCreateChild: "內容"
            }
        },
        input: {
            name: "輸入框",
            event: {
                change: "當值改變時，當組件失去焦點或用戶按Enter時觸發"
            },
            props: {
                type: "類型",
                maxlength: "最大輸入長度",
                minlength: "最小輸入長度",
                placeholder: "輸入框佔位文本",
                clearable: "是否顯示清除按鈕",
                disabled: "是否禁用",
                readonly: "是否只讀"
            }
        },
        inputNumber: {
            name: "計數器",
            props: {
                precision: "數值精度",
                min: "設置計數器允許的最小值",
                max: "設置計數器允許的最大值",
                step: "計數器步長",
                stepStrictly: "是否只能輸入 step 的倍數",
                disabled: "是否禁用計數���",
                controls: "是否使用控制按鈕",
                controlsPosition: "控制按鈕位置",
                placeholder: "輸入框佔位文本"
            }
        },
        password: {
            name: "密碼輸入框",
            event: {
                change: "當值改變時，當組件失去焦點或用戶按Enter時觸發"
            },
            props: {
                disabled: "是否禁用",
                readonly: "是否只讀",
                maxlength: "最大輸入長度",
                minlength: "最小輸入長度",
                placeholder: "輸入框佔位文本",
                clearable: "是否顯示清除按鈕"
            }
        },
        radio: {
            name: "單選框",
            props: {
                input: "是否可以填寫",
                disabled: "是否禁用",
                type: "按鈕形式",
                textColor: "按鈕形式激活時的文字顏色",
                fill: "按鈕形式激活時的填充色和邊框色"
            }
        },
        rate: {
            name: "評分",
            props: {
                max: "最大分值",
                disabled: "是否禁用",
                allowHalf: "是否允許半選",
                voidColor: "未選中時圖標的顏色",
                disabledVoidColor: "只讀時未選中時圖標的顏色",
                voidIconClass: "未選中時圖標的類名",
                disabledVoidIconClass: "只讀時未選中時圖標的類名",
                showScore: "是否顯示當前分數",
                textColor: "輔助文字的顏色",
                scoreTemplate: "分數顯示模板"
            }
        },
        select: {
            name: "選擇器",
            event: {
                removeTag: "多選模式下移除tag時觸發"
            },
            props: {
                multiple: "是否多選",
                disabled: "是否禁用",
                clearable: "是否可以清空選項",
                collapseTags: "多選時是否將選中值按文字的形式展示",
                multipleLimit: "多選時用戶最多可以選擇的項目數，為 0 則不限制",
                placeholder: "佔位符",
                filterable: "是否可搜索",
                allowCreate: "是否允許用戶創建新條目",
                noMatchText: "搜索條件無匹配時顯示的文字",
                noDataText: "選項為空時顯示的文字",
                reserveKeyword: "多選且可搜索時，是否在選中一個選項後保留當前的搜索關鍵詞",
                defaultFirstOption: "在輸入框按下回車，選擇第一個匹配項",
                remote: "其中的選項是否從服務器遠程加載",
                remoteMethod: "自定義遠程搜索方法"
            }
        },
        slider: {
            name: "滑塊",
            props: {
                min: "最小值",
                max: "最大值",
                disabled: "是否禁用",
                step: "步長",
                showInput: "是否顯示輸��框，僅在非範圍選擇時有效",
                showInputControls: "在顯示輸入框的情況下，是否顯示輸入框的控制按鈕",
                showStops: "是否顯示間斷點",
                range: "是否為範圍選擇",
                vertical: "是否豎向模式",
                height: "Slider 高度，豎向模式時必填"
            }
        },
        space: {
            name: "間距",
            props: {
                height: "高度"
            }
        },
        subForm: {
            name: "分組",
            props: {
                disabled: "是否禁用",
                syncDisabled: "是否與子表單強制同步禁用狀態"
            }
        },
        switch: {
            name: "開關",
            props: {
                disabled: "是否禁用",
                width: "寬度（px）",
                activeText: "打開時的文字描述",
                inactiveText: "關閉時的文字描述",
                activeValue: "打開時的值",
                inactiveValue: "關閉時的值",
                activeColor: "打開時的背景色",
                inactiveColor: "關閉時的背景色"
            }
        },
        tableForm: {
            name: "表格表單",
            props: {
                disabled: "是否禁用",
                max: "最多添加幾行，為 0 則不限制"
            }
        },
        tableFormColumn: {
            name: "表格格子",
            label: "自定義名稱",
            props: {
                label: "標題",
                width: "寬度",
                color: "顏色"
            }
        },
        text: {
            name: "文字",
            props: {
                formCreateNative: "是否顯示標題",
                formCreateTitle: "標題",
                formCreateChild: "內容"
            }
        },
        textarea: {
            name: "多行輸入框",
            event: {
                change: "當值改變時，當組件失去焦點或用戶按Enter時觸發"
            },
            props: {
                disabled: "是否禁用",
                readonly: "是否只讀",
                maxlength: "最大輸入長度",
                minlength: "最小輸入長度",
                showWordLimit: "是否顯示統計字數",
                placeholder: "輸入框佔位文本",
                rows: "輸入框行數",
                autosize: "高度是否自適應"
            }
        },
        timePicker: {
            name: "時間",
            props: {
                pickerOptions: "當前時間日期選擇器特有的選項",
                readonly: "完全只讀",
                disabled: "禁用",
                editable: "文本框可輸入",
                clearable: "是否顯示清除按鈕",
                placeholder: "非範圍選擇時的佔位內容",
                startPlaceholder: "範圍選擇時開始日期的佔位內容",
                endPlaceholder: "範圍選擇時結束日期的佔位內容",
                isRange: "是否為時間範圍選擇",
                arrowControl: "是否使用箭頭進行時間選擇",
                align: "對齊方式"
            }
        },
        tree: {
            name: "樹形控件",
            event: {
                nodeClick: "當節點被點擊的時候觸發",
                nodeContextmenu: "當某一節點被鼠標右鍵點擊時會觸發該事件",
                checkChange: "當複選框被點擊的時候觸發",
                check: "點擊節點複選框之後觸發",
                currentChange: "當前選中節點變化時觸發的事件",
                nodeExpand: "節點被展開時觸發的事件",
                nodeCollapse: "節點被關閉時觸發的事件",
                nodeDragStart: "節點開始拖拽時觸發的事件",
                nodeDragEnter: "拖拽進入其他節點時觸發的事件",
                nodeDragLeave: "拖拽離開某個節點時觸發的事件",
                nodeDragOver: "在拖拽節點時觸發的事件（類似瀏覽器的 mouseover 事件）",
                nodeDragEnd: "拖拽結束時（可能未成功）觸發的事件",
                nodeDrop: "拖拽成功完成時觸發的事件"
            },
            props: {
                emptyText: "內容為空的時候展示的文本",
                props: "配置選項",
                renderAfterExpand: "是否在第一次展開某個樹節點後才渲染其子節點",
                defaultExpandAll: "是否默認展開所有節點",
                expandOnClickNode: "是否在點擊節點的時候展開或者收縮節點，如果為 false，則只有點箭頭圖標的時候才會展開或者收縮節點。",
                checkOnClickNode: "是否在點擊節點的時候選中節點",
                autoExpandParent: "展開子節點的時候是否自動展開父節點",
                checkStrictly: "在顯示複選框的情況下，是否嚴格的遵循父子不互相關聯的做法",
                accordion: "是否每次只打開一個同級樹節點展開",
                indent: "相鄰級節點間的水平縮進(px)",
                nodeKey: "每個樹節點用來作為唯一標識的屬性，整棵樹應該是唯一的"
            }
        },
        upload: {
            name: "上傳",
            info: "在onSuccess方法中將接口返回的url賦值給file.url",
            event: {
                remove: "文件列表移除文件時觸發"
            },
            props: {
                listType: "上傳類型",
                multiple: "是否支持多選文件",
                action: "上傳的地址(必填)",
                beforeUpload: "上傳文件之前的回調",
                onSuccess: "上傳成功回調",
                headers: "設置上傳的請求頭部",
                data: "上傳時附帶的額外參數",
                name: "上傳的文件字段名",
                withCredentials: "支持發送 cookie 憑證信息",
                accept: "接受上傳的文件類型",
                autoUpload: "是否在選取文件後立即進行上傳",
                disabled: "是否禁用",
                limit: "最大允許上傳個數"
            }
        }
    }
}
