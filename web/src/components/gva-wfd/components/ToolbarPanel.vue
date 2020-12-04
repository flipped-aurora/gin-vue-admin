<template>
    <div class="toolbar">
        <el-tooltip :content="this.i18n['tooltip.undo']" placement="bottom">
            <span class="command" data-command="undo">
                <span class="iconfont icon-undo"/>
            </span>
        </el-tooltip>
        <el-tooltip :content="this.i18n['tooltip.redo']" placement="bottom">
            <span class="command" data-command="redo">
                <span class="iconfont icon-redo"/>
            </span>
        </el-tooltip>
        <span class="separator"/>
        <el-tooltip :content="this.i18n['tooltip.copy']" placement="bottom">
            <span class="command" data-command="copy">
                <span class="iconfont icon-copy-o"/>
            </span>
        </el-tooltip>
        <el-tooltip :content="this.i18n['tooltip.paste']" placement="bottom">
            <span class="command" data-command="paste">
                <span class="iconfont icon-paster-o"/>
            </span>
        </el-tooltip>
        <el-tooltip :content="this.i18n['tooltip.delete']" placement="bottom">
            <span class="command" data-command="delete">
                <span class="iconfont icon-delete-o"/>
            </span>
        </el-tooltip>
        <span class="separator"/>
        <el-tooltip :content="this.i18n['tooltip.zoomIn']" placement="bottom">
            <span class="command" data-command="zoomIn">
                <span class="iconfont icon-zoom-in-o"/>
            </span>
        </el-tooltip>
        <el-tooltip :content="this.i18n['tooltip.zoomOut']" placement="bottom">
            <span class="command" data-command="zoomOut">
                <span class="iconfont icon-zoom-out-o"/>
            </span>
        </el-tooltip>
        <el-tooltip :content="this.i18n['tooltip.zoomReset']" placement="bottom">
            <span class="command" data-command="resetZoom">
                <span class="iconfont icon-actual-size-o"/>
            </span>
        </el-tooltip>
        <el-tooltip :content="this.i18n['tooltip.autoFit']" placement="bottom">
            <span class="command" data-command="autoFit">
                <span class="iconfont icon-fit"/>
            </span>
        </el-tooltip>
        <span class="separator"/>
        <el-tooltip :content="this.i18n['tooltip.toFront']" placement="bottom">
            <span class="command" data-command="toFront">
                <span class="iconfont icon-to-front"/>
            </span>
        </el-tooltip>
        <el-tooltip :content="this.i18n['tooltip.toBack']" placement="bottom">
            <span class="command" data-command="toBack">
                <span class="iconfont icon-to-back"/>
            </span>
        </el-tooltip>
    </div>
</template>

<script>
  export default {
    inject: ['i18n']
  }
</script>

<style lang="scss" scoped>
    .toolbar {
        text-align: left;
        width: 100%;
        padding: 8px 0;
        background-color: #fff;
        border: 1px solid #E9E9E9;
        box-shadow: 0 8px 12px 0 rgba(0, 52, 107, 0.04);
        .command {
            display: inline-block;
            margin: 0 6px;
            line-height: 27px;
            border: 1px solid rgba(2, 2, 2, 0);
            border-radius: 2px;

            span {
                margin: 0 6px;
            }

            &:nth-of-type(1) {
                margin-left: 24px;
            }

            &:hover {
                border: 1px solid #E9E9E9;
                cursor: pointer;
            }
        }

        .separator {
            margin: 4px;
            border-left: 1px solid #E9E9E9;
        }

        .icon-select.disable {
            background: #EEE;
        }

        .disable {
            color: rgba(0, 0, 0, 0.25);

            &:hover {
                border: 1px solid rgba(2, 2, 2, 0);
                cursor: default;
            }
        }
    }

    @font-face {
        font-family: "iconfont";
        src: url('../assets/iconfont/iconfont.eot?t=1566199791492'); /* IE9 */
        src: url('../assets/iconfont/iconfont.eot?t=1566199791492#iefix') format('embedded-opentype'), /* IE6-IE8 */
        url('data:application/x-font-woff2;charset=utf-8;base64,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') format('woff2'),
        url('../assets/iconfont/iconfont.woff?t=1566199791492') format('woff'),
        url('../assets/iconfont/iconfont.ttf?t=1566199791492') format('truetype'), /* chrome, firefox, opera, Safari, Android, iOS 4.2+ */
        url('../assets/iconfont/iconfont.svg?t=1566199791492#iconfont') format('svg'); /* iOS 4.1- */
    }

    .iconfont {
        font-family: "iconfont" !important;
        font-size: 16px;
        font-style: normal;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
    }

    .icon-copy-o:before {
        content: "\e8a2";
    }

    .icon-undo:before {
        content: "\e8a5";
    }

    .icon-zoom-in-o:before {
        content: "\e8a6";
    }

    .icon-actual-size-o:before {
        content: "\e7cb";
    }

    .icon-redo:before {
        content: "\e7cc";
    }

    .icon-fit:before {
        content: "\e7cd";
    }

    .icon-delete-o:before {
        content: "\e7ce";
    }

    .icon-to-front:before {
        content: "\e7cf";
    }

    .icon-to-back:before {
        content: "\e7d0";
    }

    .icon-paster-o:before {
        content: "\e7d1";
    }

    .icon-zoom-out-o:before {
        content: "\e7d2";
    }
</style>
