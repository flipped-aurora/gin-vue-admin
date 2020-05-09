<template>
    <div>
        <el-form
                :inline="true"
                :model="form"
                :rules="rules"
                label-width="85px"
                ref="menuForm"
                label-position="top"
        >
            <el-form-item prop="icon" style="width:100%">
                <!--<el-input autocomplete="off" v-model="form.meta.icon">
                  <template slot="prepend">el-icon-</template>
                </el-input>-->
                <!--<el-select v-model="value" filterable placeholder="请选择">
                    <el-option
                            v-for="item in options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                    </el-option>
                </el-select>-->
                <el-select v-model="value" placeholder="请选择">
                    <el-option
                            v-for="item in options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                        <span class="icon" :class="item.label"></span>
                        <span style="float: left">{{ item.label }}</span>

                    </el-option>
                </el-select>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        name: "icon",
        data(){
          return {
              input:'',
              options: [{
                  value: 'el-icon-platform-eleme',
                  label: 'el-icon-platform-eleme'
              }, {
                  value: 'el-icon-eleme',
                  label: 'el-icon-eleme'
              }, {
                  value: 'el-icon-delete-solid',
                  label: 'el-icon-delete-solid'
              }, {
                  value: 'el-icon-delete',
                  label: 'el-icon-delete'
              }, {
                  value: 'el-icon-s-tools',
                  label: 'el-icon-s-tools'
              }],
              value: ''
          }
        },
    }
</script>

<style scoped>
    .icon{
        float: left;
        color: rgb(132, 146, 166);
        font-size: 13px;
        line-height: 34px;
        margin-right: 10px;
    }
</style>
