export default {
    name: "en",
    form: {
        field: "Field",
        title: "Title",
        info: "Info",
        control: "Control",
        labelPosition: "Label position",
        labelStyle: "Label style",
        labelSuffix: "Label suffix",
        size: "Form size",
        event: "Form event",
        labelWidth: "Label width",
        hideRequiredAsterisk: "Hide the red asterisk next to the label of a required field",
        showMessage: "Display verification error message",
        inlineMessage: "Display validation information inline",
        submitBtn: "Whether to display the form submit button",
        resetBtn: "Whether to display the form reset button",
        appendChild: "Insert child",
        formMode: "Form mode",
        formName: "Form name",
        componentMode: "Component",
        document: "Document"
    },
    computed: {
        fieldUsed: "[{label}] Is used in the calculation formula, please modify the corresponding formula first",
        fieldExist: "[{label}] Field already exists",
        fieldEmpty: "Field is required",
        fieldChar: "Field must begin with a letter"
    },
    validate: {
        type: "Type",
        typePlaceholder: "Please select",
        trigger: "Trigger",
        mode: "Verification method",
        modes: {
            required: "required",
            pattern: "pattern",
            validator: "validator",
            min: "min",
            max: "max",
            len: "length"
        },
        types: {
            string: "String",
            array: "Multiple",
            number: "Number",
            integer: "Integer",
            float: "Float",
            object: "Collection",
            date: "Date",
            url: "Url",
            email: "Email"
        },
        message: "Error",
        auto: "Automatic",
        autoRequired: "Please enter {title}",
        autoMode: "Please enter the correct {title}",
        requiredPlaceholder: "Please enter",
        required: "Is it required",
        rule: "Validation"
    },
    tableOptions: {
        handle: "Operation",
        add: "Add",
        empty1: "Click the lower right corner",
        empty2: "Button to add a column",
        rmCol: "Delete current column",
        rmRow: "Delete current row",
        splitRow: "Split into rows",
        splitCol: "Split into columns",
        mergeBottom: "Merge downward",
        mergeRight: "Merge right",
        addTop: "Add top column",
        addBottom: "Add the following",
        addLeft: "Add left column",
        addRight: "Add right column",
        keyValue: "key-value"
    },
    struct: {
        title: "Edit",
        only: "[{label}] Only one allowed to be added",
        errorMsg: "The input content is syntactically incorrect",
        configured: "Configured"
    },
    event: {
        title: "Edit",
        create: "Create",
        list: "List",
        placeholder: "Please enter the name of the event",
        saveMsg: "Please save the event currently being edited",
        type: "Type",
        info: "Info",
        label: "Field",
        inject: {
            api: "API of current form",
            rule: "Generate rules for the current form",
            self: "Component generation rule",
            option: "Form configuration",
            args: "Original parameters of event"
        }
    },
    fetch: {
        title: "Set data",
        create: "Create data",
        config: "Request",
        action: "Action",
        actionRequired: "Please enter the correct link",
        placeholder: "Please enter the name of the data source",
        method: "Method",
        data: "Attached",
        headers: "Headers",
        parse: "Processing",
        response: "Data returned by the interface",
        onError: "onError",
        remote: "Remote",
        static: "Static",
        optionsType: {
            fetch: "Fetch",
            struct: "Static"
        }
    },
    style: {
        width: "Width",
        height: "Height",
        color: "Color",
        backgroundColor: "Background color",
        margin: "Margin",
        padding: "Padding",
        borderRadius: "Border radius",
        border: "Border",
        solid: "Solid",
        dashed: "Dashed",
        dotted: "Dotted",
        double: "Double",
        opacity: "Opacity",
        scale: "Scale",
        shadow: {
            name: "Shadow",
            x: "x-axis offset",
            y: "y-axis offset",
            vague: "blurred radius",
            extend: "diffusion radius",
            inset: "inward",
            external: "outward",
            mode: "Mode",
            classic: "Classic",
            flat: "Flat",
            solid: "Stereoscopic"
        },
        font: {
            name: "Font",
            size: "Size",
            align: "Align",
            height: "line-height",
            spacing: "letter-spacing",
            preview: "Preview"
        },
        decoration: {
            name: "Decoration",
            underline: "underline",
            "line-through": "line-through",
            overline: "overline"
        },
        weight: {
            name: "font-weight",
            300: "Fine",
            400: "Default",
            500: "Medium",
            700: "Bold"
        }
    },
    designer: {
        component: "Component",
        id: "Unique id",
        name: "Serial number",
        type: "Type",
        form: "Form",
        style: "Style",
        rule: "Basis",
        advanced: "Advanced",
        props: "Props",
        validate: "Validate",
        event: "Event",
        clearWarn: "It cannot be restored after clearing it. Are you sure you want to clear it? ",
        childEmpty: "Click the \\e789  button in the lower right corner to add a column",
        dragEmpty: "Drag the components from the list on the left here",
        unload: "Are you sure you want to leave the current page?"
    },
    menu: {
        main: "Basic",
        aide: "Auxiliary",
        layout: "Layout",
        component: "Component",
        subform: "Subform",
        tree: "Structure"
    },
    props: {
        disabled: "disabled",
        time: "time",
        email: "email",
        number: "number",
        globalData: "Global data",
        mobile: "Mobile",
        pc: "Pc",
        reactive: "Reactive",
        title: "Title",
        content: "Content",
        collection: "Collection",
        group: "Group",
        custom: "Custom",
        change: "Change",
        blur: "Blur",
        preview: "Preview",
        clear: "Clear",
        cancel: "Cancel",
        close: "Close",
        ok: "Ok",
        save: "Save",
        refresh: "Refresh",
        submit: "Submit",
        reset: "Reset",
        copy: "Copy",
        delete: "Delete",
        hide: "Hidden",
        show: "Show",
        position: "Position",
        render: "Render",
        large: "large",
        default: "default",
        small: "small",
        always: "always",
        never: "never",
        hover: "hover",
        click: "click",
        button: "button",
        year: "year",
        month: "month",
        date: "date",
        dates: "dates",
        week: "week",
        datetime: "datetime",
        "datetime-local": "datetime",
        datetimerange: "datetimerange",
        daterange: "daterange",
        monthrange: "monthrange",
        left: "left",
        right: "right",
        top: "top",
        text: "text",
        picture: "picture",
        "picture-card": "picture-card",
        center: "center",
        vertical: "vertical",
        horizontal: "horizontal",
        manage: "Manage",
        key: "key",
        name: "Name",
        value: "value",
        inputData: "Default value",
        append: "Append",
        options: "Options",
        option: "Option",
        callback: "Callback",
        _self: "Current Window",
        _blank: "New Window",
        _parent: "Parent Window",
        _top: "Top Window"
    },
    com: {
        cascader: {
            name: "Cascader",
            event: {
                expandChange: "Triggered when the expanded node changes",
                removeTag: "In multi-select mode, triggered when Tag is removed"
            },
            props: {
                props: "Options",
                placeholder: "Placeholder",
                disabled: "Disabled",
                clearable: "Whether clearing options are supported",
                showAllLevels: "Whether the full path of the selected value is displayed in the input box",
                collapseTags: "Whether to collapse Tags in multi-select mode",
                collapseTagsTooltip: "Whether to display all selected tags when the mouse hovers over the text of a collapsed tag",
                separator: "Separator",
                filterable: "Whether this option can be searched",
                tagType: "Type"
            },
            propsOpt: {
                multiple: "Whether there are multiple selections",
                expandTrigger: "How to expand the secondary menu",
                checkStrictly: "Whether it is strictly observed that parent and child nodes are not related to each other",
                emitPath: "When the selected node changes, whether to return an array consisting of the values of the menus at each level where the node is located",
                value: "The value of the specified option is an attribute value of the option object",
                label: "Specify the option label as a certain attribute value of the option object",
                children: "The child option of the specified option is a certain attribute value of the option object",
                disabled: "The specified option is disabled as a certain attribute value of the option object",
                leaf: "The flag bit of the leaf node of the specified option is an attribute value of the option object"
            }
        },
        checkbox: {
            name: "Checkbox",
            props: {
                input: "Whether to fill in",
                type: "Type",
                disabled: "Disabled",
                min: "Minimum number that can be checked",
                max: "The maximum number that can be checked",
                textColor: "Font color when the button is active",
                fill: "Border and background color when the button is active"
            }
        },
        col: {
            name: "Col",
            props: {
                span: "Number of columns occupied by grid",
                offset: "Number of spaces on the left side of the grid",
                push: "Move the grid to the right by the number of cells",
                pull: "Move the grid to the left by the number of cells"
            }
        },
        colorPicker: {
            name: "ColorPicker",
            event: {
                activeChange: "Triggered when the color currently displayed in the panel changes"
            },
            props: {
                disabled: "Disabled",
                showAlpha: "Whether transparency selection is supported",
                colorFormat: "Color format",
                predefine: "Predefined color"
            }
        },
        datePicker: {
            name: "Date",
            props: {
                pickerOptions: "Options specific to the current time and date picker",
                readonly: "Readonly",
                disabled: "Disabled",
                type: "Type",
                editable: "Text box can be input",
                clearable: "Whether to display the clear button",
                placeholder: "Placeholder content for non-range selection",
                startPlaceholder: "Placeholder content for the start date when selecting the range",
                endPlaceholder: "Placeholder content for the end date when selecting a range",
                format: "Format displayed in the input box",
                align: "Alignment",
                rangeSeparator: "Separator when selecting range",
                unlinkPanels: "Unlink the two date panels in the range selector"
            }
        },
        dateRange: {
            name: "DateRange"
        },
        timeRange: {
            name: "TimeRange"
        },
        elAlert: {
            name: "Alert",
            description: "Description",
            props: {
                title: "Title",
                type: "Type",
                description: "Supporting text",
                closable: "Whether it can be closed",
                center: "Whether the text is centered",
                closeText: "Close button custom text",
                showIcon: "Whether to display the icon",
                effect: "Select a provided theme"
            }
        },
        elButton: {
            name: "Button",
            props: {
                formCreateChild: "Content",
                size: "Size",
                type: "Type",
                plain: "Whether the button is plain",
                round: "Whether the button has rounded corners",
                circle: "Whether the button is round",
                loading: "Whether it is loading status",
                disabled: "Disabled"
            }
        },
        elCard: {
            name: "Card",
            props: {
                header: "Title",
                shadow: "Shadow display timing"
            }
        },
        elCollapse: {
            name: "Collapse",
            event: {
                change: "Switch the currently active panel, its type is string in accordion mode and array in other modes"
            },
            props: {
                accordion: "Whether it is in accordion mode"
            }
        },
        elCollapseItem: {
            name: "CollapseItem",
            props: {
                title: "Panel title",
                name: "Identifier",
                disabled: "Disabled"
            }
        },
        elDivider: {
            name: "Divider",
            props: {
                formCreateChild: "Set Content",
                contentPosition: "Set content position"
            }
        },
        elTabPane: {
            name: "TabPane",
            props: {
                label: "Title",
                disabled: "Disabled",
                name: "Identifier of the tab",
                lazy: "Whether the label is delayed in rendering"
            }
        },
        elTabs: {
            name: "Tabs",
            event: {
                tabClick: "Triggered when tab is selected",
                tabChange: "Triggered when activeName changes",
                tabRemove: "Triggered when the tab remove button is clicked",
                tabAdd: "Triggered when a new tab button is clicked",
                edit: "Triggered after clicking the add or remove button of the tab"
            },
            props: {
                type: "Type",
                closable: "Whether the label can be closed",
                tabPosition: "Tab position",
                stretch: "Whether the width of the label is self-stretching"
            }
        },
        elTag: {
            name: "Tag",
            props: {
                formCreateNative: "Whether to display title",
                formCreateTitle: "Title",
                formCreateChild: "Content",
                type: "Type",
                size: "Label size",
                effect: "Label theme",
                closable: "Whether it can be closed",
                disableTransitions: "Whether to disable gradient animation",
                hit: "Whether there is a border stroke",
                round: "Whether it is round",
                color: "Background color"
            }
        },
        elTransfer: {
            name: "Transfer",
            event: {
                leftCheckChange: "Triggered when the left list element is selected/unselected by the user",
                rightCheckChange: "Triggered when the right list element is selected/unselected by the user"
            },
            props: {
                filterable: "Is it searchable",
                filterPlaceholder: "Search box placeholder",
                targetOrder: "Sort strategy of list elements on the right",
                targetOrderInfo: "If it is original, keep the same order as the data; if it is push, the newly added elements will be ranked last; if it is unshift, the newly added elements will be ranked first",
                titles: "Title",
                buttonTexts: "Set button content",
                props: "Field alias of data source"
            }
        },
        elTreeSelect: {
            name: "TreeSelect",
            event: {
                removeTag: "Triggered when tag is removed in multi-select mode"
            },
            props: {
                multiple: "Whether there are multiple selections",
                disabled: "Disabled",
                clearable: "Whether the option can be cleared",
                collapseTags: "Whether to display the selected value as text during multi-selection",
                multipleLimit: "The maximum number of items that the user can select during multiple selection, if it is 0, there is no limit",
                placeholder: "Placeholder",
                props: "Options",
                renderAfterExpand: "Whether to render its child nodes after expanding a tree node for the first time",
                defaultExpandAll: "Whether to expand all nodes by default",
                expandOnClickNode: "Whether to expand or shrink nodes when clicking on them",
                checkOnClickNode: "Whether to select the node when clicking the node",
                nodeKey: "Each tree node is used as an attribute for unique identification, and the entire tree should be unique"
            }
        },
        fcEditor: {
            name: "Editor",
            props: {
                disabled: "Disabled"
            }
        },
        fcRow: {
            name: "Row",
            props: {
                gutter: "Grid interval",
                type: "Flex layout mode",
                justify: "Horizontal arrangement under flex layout",
                align: "Vertical arrangement under flex layout"
            }
        },
        fcTable: {
            name: "Table",
            props: {
                border: "Whether to display border",
                borderColor: "Border color",
                borderWidth: "Border width"
            }
        },
        fcTableGrid: {
            name: "Grid"
        },
        group: {
            name: "Subform",
            props: {
                disabled: "Disabled",
                syncDisabled: "Whether to force synchronization of the disabled state with the subform",
                expand: "Set the default expansion items",
                button: "Whether to display the operation button",
                sortBtn: "Whether to display the sort button",
                min: "Set the minimum number of items to add",
                max: "Set the maximum number of items to add"
            }
        },
        html: {
            name: "HTML",
            props: {
                formCreateNative: "Whether to display title",
                formCreateTitle: "Title",
                formCreateChild: "Content"
            }
        },
        input: {
            name: "Input",
            event: {
                change: "Triggered when the value changes, when the component loses focus or the user presses Enter"
            },
            props: {
                type: "Type",
                maxlength: "Maximum input length",
                minlength: "Minimum input length",
                placeholder: "Placeholder",
                clearable: "Whether to display the clear button",
                disabled: "Disabled",
                readonly: "Readonly"
            }
        },
        inputNumber: {
            name: "InputNumber",
            props: {
                precision: "Precision of input value",
                min: "Set the minimum value allowed for the counter",
                max: "Set the maximum allowed value of the counter",
                step: "Step",
                stepStrictly: "Whether only multiples of step can be entered",
                disabled: "Disabled",
                controls: "Whether to use control buttons",
                controlsPosition: "Control button position",
                placeholder: "Placeholder"
            }
        },
        password: {
            name: "Password",
            event: {
                change: "Triggered when the value changes, when the component loses focus or the user presses Enter"
            },
            props: {
                disabled: "Disabled",
                readonly: "Readonly",
                maxlength: "Maximum input length",
                minlength: "Minimum input length",
                placeholder: "Placeholder",
                clearable: "Whether to display the clear button"
            }
        },
        radio: {
            name: "Radio",
            props: {
                input: "Whether to fill in",
                disabled: "Disabled",
                type: "Type",
                textColor: "Text color when button form is activated",
                fill: "Fill color and border color when the button form is activated"
            }
        },
        rate: {
            name: "Rate",
            props: {
                max: "Maximum score",
                disabled: "Disabled",
                allowHalf: "Whether to allow half selection",
                voidColor: "Color of the icon when not selected",
                disabledVoidColor: "The color of the icon when it is not selected when read-only",
                voidIconClass: "Class name of the icon when not selected",
                disabledVoidIconClass: "The class name of the icon when it is not selected when read-only",
                showScore: "Whether to display the current score",
                textColor: "Color of auxiliary text",
                scoreTemplate: "Score display template"
            }
        },
        select: {
            name: "Select",
            event: {
                removeTag: "Triggered when tag is removed in multi-select mode"
            },
            props: {
                multiple: "Whether there are multiple selections",
                disabled: "Disabled",
                clearable: "Whether the option can be cleared",
                collapseTags: "Whether to display the selected value as text during multi-selection",
                multipleLimit: "The maximum number of items that the user can select when multiple-selecting, if it is 0, there is no limit",
                placeholder: "Placeholder",
                filterable: "Is it searchable",
                allowCreate: "Whether users are allowed to create new entries",
                noMatchText: "Text displayed when no search conditions match",
                noDataText: "Text displayed when option is empty",
                reserveKeyword: "When multiple selections are searchable, whether to retain the current search keyword after selecting an option",
                defaultFirstOption: "Press Enter in the input box and select the first matching item",
                remote: "Whether the options are loaded remotely from the server",
                remoteMethod: "Custom remote search methods"
            }
        },
        slider: {
            name: "Slider",
            props: {
                min: "Minimum value",
                max: "Maximum value",
                disabled: "Disabled",
                step: "Step",
                showInput: "Whether to display the input box, it is only valid during non-range selection",
                showInputControls: "Whether to display the control buttons of the input box when the input box is displayed",
                showStops: "Whether to display discontinuities",
                range: "Whether it is a range selection",
                vertical: "Whether portrait mode",
                height: "Slider height, required in portrait mode"
            }
        },
        space: {
            name: "Space",
            props: {
                height: "Height"
            }
        },
        subForm: {
            name: "Group",
            props: {
                disabled: "Disabled",
                syncDisabled: "Whether to force synchronization of the disabled state with the subform"
            }
        },
        switch: {
            name: "Switch",
            props: {
                disabled: "Disabled",
                width: "Width (px)",
                activeText: "Text description when opening",
                inactiveText: "Text description when closing",
                activeValue: "Value when opening",
                inactiveValue: "Value when closed",
                activeColor: "Background color when opening",
                inactiveColor: "Background color when closed"
            }
        },
        tableForm: {
            name: "TableForm",
            props: {
                disabled: "Disabled",
                max: "Maximum number of rows to add, if 0, there is no limit"
            }
        },
        tableFormColumn: {
            name: "TableFormColumn",
            label: "TableFormColumn",
            props: {
                label: "Title",
                width: "Width",
                color: "Color"
            }
        },
        text: {
            name: "Text",
            props: {
                formCreateNative: "Whether to display title",
                formCreateTitle: "Title",
                formCreateChild: "Content"
            }
        },
        textarea: {
            name: "Textarea",
            event: {
                change: "Triggered when the value changes, when the component loses focus or the user presses Enter"
            },
            props: {
                disabled: "Disabled",
                readonly: "Readonly",
                maxlength: "Maximum input length",
                minlength: "Minimum input length",
                showWordLimit: "Whether to display word count statistics",
                placeholder: "Placeholder",
                rows: "Number of input box rows",
                autosize: "Whether the height is adaptive"
            }
        },
        timePicker: {
            name: "Time",
            props: {
                pickerOptions: "Options specific to the current time and date picker",
                readonly: "Readonly",
                disabled: "Disabled",
                editable: "Text box can be input",
                clearable: "Whether to display the clear button",
                placeholder: "Placeholder content for non-range selection",
                startPlaceholder: "Placeholder content for the start date when selecting the range",
                endPlaceholder: "Placeholder content for the start date when selecting the range",
                isRange: "Whether to select a time range",
                arrowControl: "Whether to use arrows for time selection",
                align: "Align"
            }
        },
        tree: {
            name: "Tree",
            event: {
                nodeClick: "Triggered when the node is clicked",
                nodeContextmenu: "This event will be triggered when a node is right-clicked",
                checkChange: "Triggered when the check box is clicked",
                check: "Triggered after clicking the node checkbox",
                currentChange: "Event triggered when the currently selected node changes",
                nodeExpand: "Event triggered when a node is expanded",
                nodeCollapse: "Event triggered when a node is closed",
                nodeDragStart: "Event triggered when a node starts dragging",
                nodeDragEnter: "Event triggered when dragging into other nodes",
                nodeDragLeave: "Event triggered when dragging leaves a node",
                nodeDragOver: "Event triggered when dragging a node",
                nodeDragEnd: "Event triggered when drag ends",
                nodeDrop: "Event triggered when drag and drop is successfully completed"
            },
            props: {
                emptyText: "Text displayed when the content is empty",
                props: "Options",
                renderAfterExpand: "Whether to render its child nodes after expanding a tree node for the first time",
                defaultExpandAll: "Whether to expand all nodes by default",
                expandOnClickNode: "Whether to expand or contract the node when clicking the node, if it is false, the node will only be expanded or contracted when the arrow icon is clicked. ",
                checkOnClickNode: "Whether to select the node when clicking the node",
                autoExpandParent: "Whether to automatically expand the parent node when expanding the child node",
                checkStrictly: "When the check box is displayed, whether the parent and child are strictly not related to each other should be strictly followed",
                accordion: "Whether to open only one sibling tree node for expansion at a time",
                indent: "Horizontal indent (px) between adjacent level nodes",
                nodeKey: "Each tree node is used as an attribute for unique identification, and the entire tree should be unique"
            }
        },
        upload: {
            name: "Upload",
            info: "Assign the url returned by the interface to file.url in the onSuccess method",
            event: {
                remove: "Triggered when a file is removed from the file list"
            },
            props: {
                listType: "Upload type",
                multiple: "Whether multiple selection of files is supported",
                action: "Upload address (required)",
                beforeUpload: "Hook before uploading file",
                onSuccess: "Upload success callback",
                headers: "Set upload request headers",
                data: "Extra parameters attached when uploading",
                name: "Uploaded file field name",
                withCredentials: "Support sending cookie credential information",
                accept: "Accept uploaded file types",
                autoUpload: "Whether to upload the file immediately after selecting it",
                disabled: "Disabled",
                limit: "Maximum number of uploads allowed"
            }
        }
    }
}
