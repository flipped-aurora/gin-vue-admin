<template>
  <div :data-clazz="model.clazz">
    <div class="panelTitle">{{ i18n['userTask'] }}</div>
    <div class="panelBody">
      <DefaultDetail :model="model" :onChange="onChange" :readOnly="readOnly" />
      <div class="panelRow">
        <div>选择角色：</div>
        <el-select
          style="width: 90%; font-size: 12px"
          placeholder="请选择角色（与用户互斥）"
          :value="model.assignType"
          :disabled="readOnly"
          @change="
            (e) => {
              onChange('assignValue', '')
              onChange('assignType', e)
            }
          "
        >
          <el-option key="user" value="user" :label="'用户'" />
          <el-option key="authority" value="authority" :label="'角色'" />
        </el-select>
      </div>
      <div v-if="model.assignType == 'user'" class="panelRow">
        <div>选择用户：</div>
        <el-select
          style="width: 90%; font-size: 12px"
          :placeholder="'请选择用户'"
          :disabled="readOnly"
          :value="model.assignValue"
          clearable
          multiple
          :filterable="true"
          :filter-method="filterUsers"
          @change="(e) => onChange('assignValue', e)"
        >
          <el-option
            v-for="user in usersCopy"
            :key="user.id"
            :label="user.name"
            :value="user.id"
          />
        </el-select>
      </div>
      <div v-if="model.assignType == 'authority'" class="panelRow">
        <div>选择角色：</div>
        <el-select
          style="width: 90%; font-size: 12px"
          :placeholder="'请选择角色'"
          :disabled="readOnly"
          multiple
          :value="model.assignValue"
          clearable
          :filterable="true"
          :filter-method="filterAuthorities"
          @change="(e) => onChange('assignValue', e)"
        >
          <el-option
            v-for="authority in authoritiesCopy"
            :key="authority.id"
            :label="authority.name"
            :value="authority.id"
          />
        </el-select>
      </div>
      <div class="panelRow">
        <div style="display: inline">{{ i18n['userTask.dueDate'] }}：</div>
        <el-date-picker
          type="datetime"
          style="width: 90%; min-width: null"
          :placeholder="i18n['userTask.dueDate.placeholder']"
          :disabled="readOnly"
          :value="model.dueDate"
          @input="(value) => onChange('dueDate', value)"
        />
      </div>
      <!-- <div class="panelRow">
        <el-checkbox
          @change="(value) => onChange('isSequential', value)"
          :disabled="readOnly"
          :value="!!model.isSequential"
          >{{ i18n['userTask.counterSign'] }}</el-checkbox
        >
      </div> -->
      
      <div class="panelRow">
        <div>视图文件路径（以view开头）</div>
        <el-input
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          placeholder="请输入视图文件路径（以view开头）"
          type="view"
          :value="model.view"
          @input="
            (value) => {
              onChange('view', value)
            }
          "
        />
      </div>
      <div class="panelRow">
        <div>步骤：</div>
        <el-input
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          :value="model.step"
          placeholder="请输入步骤"
          @input="
            (value) => {
              onChange('step', value)
            }
          "
        />
      </div>
      <div class="panelRow">
        <div>详情说明：</div>
        <el-input
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          placeholder="请输入详情内容"
          type="textarea"
          :value="model.description"
          @input="
            (value) => {
              onChange('description', value)
            }
          "
        />
      </div>
    </div>
  </div>
</template>
<script>
  import DefaultDetail from "./DefaultDetail";
  export default {
    inject: ['i18n'],
    components: {
      DefaultDetail
    },
    props: {
      model: {
        type:Object,
        default: ()=>({}),
      },
      users: {
        type: Array,
        default: ()=>([]),
      },
      authority: {
        type: Array,
        default: ()=>([]),
      },
      authorities: {
        type: Array,
        default: ()=>([]),
      },
      groups: {
        type: Array,
        default: ()=>([]),
      },
      onChange: {
        type: Function,
        default: ()=>{}
      },
      readOnly:{
        type: Boolean,
        default: false,
      }
    },
    data() {
      return {
        usersCopy: this.users,
        groupsCopy: this.groups,
        authoritiesCopy:this.authorities
      }
    },
    methods: {
      filterUsers(input) {
        if (input) {
          this.usersCopy = this.users.filter((item) => {
            if (item.name.indexOf(input) >-1 ) {
              return true
            }
          })
        } else {
          this.usersCopy = this.users;
        }
      },
      filterAuthorities(input) {
        if (input) {
          this.authoritiesCopy = this.authorities.filter((item) => {
            console.log(item)
            if (item.name.indexOf(input) >-1 ) {
              return true
            }
          })
        } else {
          this.authoritiesCopy = this.authorities;
        }
      },
      filterGroups(input) {
        if (input) {
          this.groupsCopy = this.groups.filter((item) => {
            if (!!~item.name.indexOf(input) || !!~item.name.toLowerCase().indexOf(input.toLowerCase())) {
              return true
            }
          })
        } else {
          this.groupsCopy = this.groups;
        }
      }
    }
  }
</script>
