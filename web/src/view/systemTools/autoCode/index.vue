<template>
  <div>
    <warning-bar
      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
      title="此功能为开发环境使用，不建议发布到生产，具体使用效果请点我观看。"
    />
    <div class="gva-search-box" v-if="!isAdd">
      <div class="text-lg mb-2 text-gray-600">
        使用AI创建<a
          class="text-blue-600 text-sm ml-4"
          href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
          target="_blank"
          >获取AiPath</a
        >
      </div>
      <div class="relative">
        <el-input
          v-model="prompt"
          type="textarea"
          :rows="5"
          :maxlength="2000"
          :placeholder="`现已完全免费\n试试复制一张图片然后按下ctrl+v或者commend+v\n试试描述你的表，让AI帮你完成。\n此功能需要到插件市场个人中心获取自己的AI-Path，把AI-Path填入config.yaml下的autocode-->ai-path，重启项目即可使用。\n按下 Ctrl+Enter 或 Cmd+Enter 直接生成`"
          resize="none"
          @focus="handleFocus"
          @blur="handleBlur"
        />

        <div class="flex absolute right-28 bottom-2">
          <el-tooltip effect="light">
            <template #content>
              <div>
                【完全免费】前往<a
                  class="text-blue-600"
                  href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
                  target="_blank"
              >插件市场个人中心</a
              >申请AIPath，填入config.yaml的ai-path属性即可使用。
              </div>
            </template>
            <el-button
                :disabled="form.onlyTemplate"
                type="primary"
                @click="eyeFunc()"
            >
              <el-icon size="18">
                <ai-gva />
              </el-icon>
              识图
            </el-button>
          </el-tooltip>
        </div>

        <div class="flex absolute right-2 bottom-2">
          <el-tooltip effect="light">
            <template #content>
              <div>
                【完全免费】前往<a
                  class="text-blue-600"
                  href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
                  target="_blank"
                  >插件市场个人中心</a
                >申请AIPath，填入config.yaml的ai-path属性即可使用。
              </div>
            </template>
            <el-button
              :disabled="form.onlyTemplate"
              type="primary"
              @click="llmAutoFunc()"
            >
              <el-icon size="18">
                <ai-gva />
              </el-icon>
              生成
            </el-button>
          </el-tooltip>
        </div>
      </div>
    </div>
    <!-- 从数据库直接获取字段 -->
    <div class="gva-search-box" v-if="!isAdd">
      <div class="text-lg mb-2 text-gray-600">从数据库创建</div>
      <el-form
        ref="getTableForm"
        :inline="true"
        :model="dbform"
        label-width="120px"
      >
        <el-row class="w-full">
          <el-col :span="6">
            <el-form-item label="业务库" prop="selectDBtype" class="w-full">
              <template #label>
                <el-tooltip
                  content="注：需要提前到db-list自行配置多数据库，如未配置需配置后重启服务方可使用。（此处可选择对应库表，可理解为从哪个库选择表）"
                  placement="bottom"
                  effect="light"
                >
                  <div>
                    业务库 <el-icon><QuestionFilled /></el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-select
                v-model="dbform.businessDB"
                clearable
                placeholder="选择业务库"
                @change="getDbFunc"
                class="w-full"
              >
                <el-option
                  v-for="item in dbList"
                  :key="item.aliasName"
                  :value="item.aliasName"
                  :label="item.aliasName"
                  :disabled="item.disable"
                >
                  <div>
                    <span>{{ item.aliasName }}</span>
                    <span
                      style="float: right; color: #8492a6; font-size: 13px"
                      >{{ item.dbName }}</span
                    >
                  </div>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="数据库名" prop="structName" class="w-full">
              <el-select
                v-model="dbform.dbName"
                clearable
                filterable
                placeholder="请选择数据库"
                class="w-full"
                @change="getTableFunc"
              >
                <el-option
                  v-for="item in dbOptions"
                  :key="item.database"
                  :label="item.database"
                  :value="item.database"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="表名" prop="structName" class="w-full">
              <el-select
                v-model="dbform.tableName"
                :disabled="!dbform.dbName"
                class="w-full"
                filterable
                placeholder="请选择表"
              >
                <el-option
                  v-for="item in tableOptions"
                  :key="item.tableName"
                  :label="item.tableName"
                  :value="item.tableName"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item class="w-full">
              <div class="flex justify-end w-full">
                <el-button type="primary" @click="getColumnFunc">
                  使用此表
                </el-button>
              </div>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
    <div class="gva-search-box">
      <!-- 初始版本自动化代码工具 -->
      <div class="text-lg mb-2 text-gray-600">自动化结构</div>
      <el-form
        :disabled="isAdd"
        ref="autoCodeForm"
        :rules="rules"
        :model="form"
        label-width="120px"
        :inline="true"
      >
        <el-row class="w-full">
          <el-col :span="6">
            <el-form-item label="结构名称" prop="structName" class="w-full">
              <div class="flex gap-2 w-full">
                <el-input
                  v-model="form.structName"
                  placeholder="首字母自动转换大写"
                />
                <el-button
                  :disabled="form.onlyTemplate"
                  type="primary"
                  @click="llmAutoFunc(true)"
                >
                  <el-icon size="18">
                    <ai-gva />
                  </el-icon>
                  生成
                </el-button>
              </div>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="abbreviation" prop="abbreviation" class="w-full">
              <template #label>
                <el-tooltip
                  content="简称会作为入参对象名和路由group"
                  placement="bottom"
                  effect="light"
                >
                  <div>
                    结构简称 <el-icon><QuestionFilled /></el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-input
                v-model="form.abbreviation"
                placeholder="请输入Struct简称"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="中文名称" prop="description" class="w-full">
              <el-input
                v-model="form.description"
                placeholder="中文描述作为自动api描述"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="表名" prop="tableName" class="w-full">
              <el-input
                v-model="form.tableName"
                placeholder="指定表名（非必填）"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="w-full">
          <el-col :span="6">
            <el-form-item prop="packageName" class="w-full">
              <template #label>
                <el-tooltip
                  content="生成文件的默认名称(建议为驼峰格式,首字母小写,如sysXxxXxxx)"
                  placement="bottom"
                  effect="light"
                >
                  <div>
                    文件名称 <el-icon><QuestionFilled /></el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-input
                v-model="form.packageName"
                placeholder="请输入文件名称"
                @blur="toLowerCaseFunc(form, 'packageName')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              label="选择模板"
              prop="package"
              class="w-full relative"
            >
              <el-select v-model="form.package" class="w-full pr-12" filterable>
                <el-option
                  v-for="item in pkgs"
                  :key="item.ID"
                  :value="item.packageName"
                  :label="item.packageName"
                />
              </el-select>
              <span class="absolute right-0">
                <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getPkgs"
                >
                  <refresh />
                </el-icon>
                <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goPkgs"
                >
                  <document-add />
                </el-icon>
              </span>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="业务库" prop="businessDB" class="w-full">
              <template #label>
                <el-tooltip
                  content="注：需要提前到db-list自行配置多数据库，此项为空则会使用gva本库创建自动化代码(global.GVA_DB),填写后则会创建指定库的代码(global.MustGetGlobalDBByDBName(dbname))"
                  placement="bottom"
                  effect="light"
                >
                  <div>
                    业务库 <el-icon><QuestionFilled /></el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-select
                v-model="form.businessDB"
                clearable
                placeholder="选择业务库"
                class="w-full"
              >
                <el-option
                  v-for="item in dbList"
                  :key="item.aliasName"
                  :value="item.aliasName"
                  :label="item.aliasName"
                  :disabled="item.disable"
                >
                  <div>
                    <span>{{ item.aliasName }}</span>
                    <span
                      style="float: right; color: #8492a6; font-size: 13px"
                      >{{ item.dbName }}</span
                    >
                  </div>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
    <div class="gva-search-box">
      <el-collapse class="no-border-collapse">
        <el-collapse-item>
          <template #title>
            <div class="text-lg text-gray-600 font-normal">
              专家模式
            </div>
          </template>
          <template #icon="{ isActive }">
          <span class="text-lg ml-auto mr-4 font-normal">
            {{ isActive ? '收起' : '展开' }}
          </span>
          </template>
          <div class="p-4">
            <!-- 基础设置组 -->
            <div class="border-b border-gray-200 last:border-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">基础设置</h3>
              <el-row :gutter="20">
                <el-col :span="3">
                  <el-tooltip
                      content="注：会自动在结构体global.Model其中包含主键和软删除相关操作配置"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item label="使用GVA结构">
                      <el-checkbox v-model="form.gvaModel" @change="useGva" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                      content="注：会自动产生页面内的按钮权限配置，若不在角色管理中进行按钮分配则按钮不可见"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item label="创建按钮权限">
                      <el-checkbox :disabled="!form.generateWeb" v-model="form.autoCreateBtnAuth" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-form-item label="生成前端">
                    <el-checkbox v-model="form.generateWeb" />
                  </el-form-item>
                </el-col>
                <el-col :span="3">
                  <el-form-item label="生成后端">
                    <el-checkbox disabled v-model="form.generateServer" />
                  </el-form-item>
                </el-col>
              </el-row>
            </div>

            <!-- 自动化设置组 -->
            <div class="border-b border-gray-200 last:border-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">自动化设置</h3>
              <el-row :gutter="20">
                <el-col :span="3">
                  <el-tooltip
                      content="注：把自动生成的API注册进数据库"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item label="自动创建API">
                      <el-checkbox  :disabled="!form.generateServer" v-model="form.autoCreateApiToSql" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                      content="注：把自动生成的菜单注册进数据库"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item label="自动创建菜单">
                      <el-checkbox :disabled="!form.generateWeb" v-model="form.autoCreateMenuToSql" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                      content="注：自动同步数据库表结构，如果不需要可以选择关闭"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item label="同步表结构">
                      <el-checkbox  :disabled="!form.generateServer" v-model="form.autoMigrate" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
              </el-row>
            </div>

            <!-- 高级设置组 -->
            <div class="border-b border-gray-200 last:border-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">高级设置</h3>
              <el-row :gutter="20">
                <el-col :span="3">
                  <el-tooltip
                      content="注：会自动在结构体添加 created_by updated_by deleted_by，方便用户进行资源权限控制"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item label="创建资源标识">
                      <el-checkbox v-model="form.autoCreateResource" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                      content="注：使用基础模板将不会生成任何结构体和CURD,仅仅配置enter等属性方便自行开发非CURD逻辑"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item label="基础模板">
                      <el-checkbox v-model="form.onlyTemplate" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
              </el-row>
            </div>

            <!-- 树形结构设置 -->
            <div class="last:pb-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">树形结构设置</h3>
              <el-row :gutter="20" align="middle">
                <el-col :span="24">
                    <el-form-item label="树型结构">
                      <div class="flex items-center gap-4">
                        <el-tooltip
                            content="注：会自动创建parentID来进行父子关系关联,仅支持主键为int类型"
                            placement="top"
                            effect="light"
                        >
                          <el-checkbox v-model="form.isTree" />
                        </el-tooltip>
                        <el-input
                            v-model="form.treeJson"
                            :disabled="!form.isTree"
                            placeholder="前端展示json属性"
                            class="flex-1"
                        />
                      </div>
                    </el-form-item>
                </el-col>
              </el-row>
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <!-- 组件列表 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          @click="editAndAddField()"
          :disabled="form.onlyTemplate"
        >
          新增字段
        </el-button>
      </div>
      <div class="draggable">
        <el-table :data="form.fields" row-key="fieldName">
          <el-table-column
            v-if="!isAdd"
            fixed="left"
            align="left"
            type="index"
            width="60"
          >
            <template #default>
              <el-icon class="cursor-grab drag-column">
                <MoreFilled />
              </el-icon>
            </template>
          </el-table-column>
          <el-table-column
            fixed="left"
            align="left"
            type="index"
            label="序列"
            width="60"
          />
          <el-table-column
            fixed="left"
            align="left"
            type="index"
            label="主键"
            width="60"
          >
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.primaryKey" />
            </template>
          </el-table-column>
          <el-table-column
            fixed="left"
            align="left"
            prop="fieldName"
            label="字段名称"
            width="160"
          >
            <template #default="{ row }">
              <el-input disabled v-model="row.fieldName" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldDesc"
            label="中文名"
            width="160"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.fieldDesc" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="defaultValue"
            label="默认值"
            width="160"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.defaultValue" />
            </template>
          </el-table-column>
          <el-table-column align="left" prop="require" label="必填">
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.require" />
            </template>
          </el-table-column>
          <el-table-column align="left" prop="sort" label="排序">
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.sort" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="form"
            width="100"
            label="新建/编辑"
          >
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.form" />
            </template>
          </el-table-column>
          <el-table-column align="left" prop="table" label="表格">
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.table" />
            </template>
          </el-table-column>
          <el-table-column align="left" prop="desc" label="详情">
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.desc" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="excel"
            width="100"
            label="导入/导出"
            v-if="!isAdd"
          >
            <template #default="{ row }">
              <el-checkbox v-model="row.excel" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldJson"
            width="160px"
            label="字段Json"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.fieldJson" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldType"
            label="字段类型"
            width="160"
          >
            <template #default="{ row }">
              <el-select
                v-model="row.fieldType"
                style="width: 100%"
                placeholder="请选择字段类型"
                :disabled="row.disabled"
                clearable
              >
                <el-option
                  v-for="item in typeOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldIndexType"
            label="索引类型"
            width="160"
          >
            <template #default="{ row }">
              <el-select
                v-model="row.fieldIndexType"
                style="width: 100%"
                placeholder="请选择字段索引类型"
                :disabled="row.disabled"
                clearable
              >
                <el-option
                  v-for="item in typeIndexOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="dataTypeLong"
            label="字段长度/枚举值"
            width="160"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.dataTypeLong" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="columnName"
            label="数据库字段"
            width="160"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.columnName" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="comment"
            label="数据库字段描述"
            width="160"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.comment" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="fieldSearchType"
            label="搜索条件"
            width="130"
          >
            <template #default="{ row }">
              <el-select
                v-model="row.fieldSearchType"
                style="width: 100%"
                placeholder="请选择字段查询条件"
                clearable
                :disabled="row.fieldType === 'json' || row.disabled"
              >
                <el-option
                  v-for="item in typeSearchOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                  :disabled="canSelect(row.fieldType,item.value)"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column align="left" label="操作" width="300" fixed="right">
            <template #default="scope">
              <el-button
                v-if="!scope.row.disabled"
                type="primary"
                link
                icon="edit"
                @click="editAndAddField(scope.row)"
              >
                高级编辑
              </el-button>
              <el-button
                v-if="!scope.row.disabled"
                type="primary"
                link
                icon="delete"
                @click="deleteField(scope.$index)"
              >
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!-- 组件列表 -->
      <div class="gva-btn-list justify-end mt-4">
        <el-button type="primary" :disabled="isAdd" @click="exportJson()">
          导出json
        </el-button>
        <el-upload
          class="flex items-center"
          :before-upload="importJson"
          :show-file-list="false"
          :headers="{'x-token': token}"
          accept=".json"
        >
          <el-button type="primary" class="mx-2" :disabled="isAdd"
            >导入json</el-button
          >
        </el-upload>
        <el-button type="primary" :disabled="isAdd" @click="clearCatch()">
          清除暂存
        </el-button>
        <el-button type="primary" :disabled="isAdd" @click="catchData()">
          暂存
        </el-button>
        <el-button type="primary" :disabled="isAdd" @click="enterForm(false)">
          生成代码
        </el-button>
        <el-button type="primary" @click="enterForm(true)">
          {{ isAdd ? '查看代码' : '预览代码' }}
        </el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-drawer v-model="dialogFlag" size="70%" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">组件内容</span>
          <div>
            <el-button @click="closeDialog"> 取 消 </el-button>
            <el-button type="primary" @click="enterDialog"> 确 定 </el-button>
          </div>
        </div>
      </template>

      <FieldDialog
        v-if="dialogFlag"
        ref="fieldDialogNode"
        :dialog-middle="dialogMiddle"
        :type-options="typeOptions"
        :type-search-options="typeSearchOptions"
        :type-index-options="typeIndexOptions"
      />
    </el-drawer>

    <el-drawer v-model="previewFlag" size="80%" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">操作栏</span>
          <div>
            <el-button type="primary" @click="selectText"> 全选 </el-button>
            <el-button type="primary" @click="copy"> 复制 </el-button>
          </div>
        </div>
      </template>
      <PreviewCodeDialog
        v-if="previewFlag"
        :is-add="isAdd"
        ref="previewNode"
        :preview-code="preViewCode"
      />
    </el-drawer>
  </div>
</template>

<script setup>
  import FieldDialog from '@/view/systemTools/autoCode/component/fieldDialog.vue'
  import PreviewCodeDialog from '@/view/systemTools/autoCode/component/previewCodeDialog.vue'
  import {
    toUpperCase,
    toHump,
    toSQLLine,
    toLowerCase
  } from '@/utils/stringFun'
  import {
    createTemp,
    getDB,
    getTable,
    getColumn,
    preview,
    getMeta,
    getPackageApi,
    llmAuto, butler, eye
  } from '@/api/autoCode'
  import { getDict } from '@/utils/dictionary'
  import { ref, watch, toRaw, onMounted, nextTick } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import Sortable from 'sortablejs'
  import { useUserStore } from "@/pinia";

  const userStore = useUserStore()

  const token = userStore.token

  const handleFocus = () => {
    document.addEventListener('keydown', handleKeydown);
    document.addEventListener('paste', handlePaste);
  }

  const handleBlur = () => {
    document.removeEventListener('keydown', handleKeydown);
    document.removeEventListener('paste', handlePaste);
  }

  const handleKeydown = (event) => {
    if ((event.ctrlKey || event.metaKey) && event.key === 'Enter') {
      llmAutoFunc()
    }
  }

  const handlePaste = (event) => {
    const items = event.clipboardData.items;
    for (let i = 0; i < items.length; i++) {
      if (items[i].type.indexOf('image') !== -1) {
        const file = items[i].getAsFile();
        const reader = new FileReader();
        reader.onload =async (e) => {
          const base64String = e.target.result;
          const res = await eye({ picture: base64String,command: 'eye' })
          if (res.code === 0) {
            prompt.value = res.data
            llmAutoFunc()
          }
        };
        reader.readAsDataURL(file);
      }
    }
  };

  const getOnlyNumber = () => {
    let randomNumber = ''
    while (randomNumber.length < 16) {
      randomNumber += Math.random().toString(16).substring(2)
    }
    return randomNumber.substring(0, 16)
  }

  const prompt = ref('')

  const eyeFunc = async () => {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = 'image/*';

    input.onchange = (event) => {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = async (e) => {
          const base64String = e.target.result;

          const res = await eye({ picture: base64String,command: 'eye' })
          if (res.code === 0) {
            prompt.value = res.data
            llmAutoFunc()
          }
        };
        reader.readAsDataURL(file);
      }
    };

    input.click();
  }


  const llmAutoFunc = async (flag) => {
    if (flag && !form.value.structName) {
      ElMessage.error('请输入结构体名称')
      return
    }
    if (!flag && !prompt.value) {
      ElMessage.error('请输入描述')
      return
    }

    if (form.value.fields.length > 0) {
      const res = await ElMessageBox.confirm(
        'AI生成会清空当前数据，是否继续?',
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
      if (res !== 'confirm') {
        return
      }
    }

    const res = await llmAuto({
      prompt: flag ? '结构体名称为' + form.value.structName : prompt.value
    })
    if (res.code === 0) {
      form.value.fields = []
      const json = JSON.parse(res.data)
      json.fields?.forEach((item) => {
        item.fieldName = toUpperCase(item.fieldName)
      })

      for (let key in json) {
        form.value[key] = json[key]
      }

      form.value.generateServer = true
      form.value.generateWeb = true

    }
  }

  const isAdd = ref(false)

  // 行拖拽
  const rowDrop = () => {
    // 要拖拽元素的父容器
    const tbody = document.querySelector(
      '.draggable .el-table__body-wrapper tbody'
    )
    Sortable.create(tbody, {
      //  可被拖拽的子元素
      draggable: '.draggable .el-table__row',
      handle: '.drag-column',
      onEnd: async ({ newIndex, oldIndex }) => {
        await nextTick()
        const currRow = form.value.fields.splice(oldIndex, 1)[0]
        form.value.fields.splice(newIndex, 0, currRow)
      }
    })
  }

  onMounted(() => {
    rowDrop()
  })

  defineOptions({
    name: 'AutoCode'
  })
  const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']

  const dataModelList = ['created_by', 'updated_by', 'deleted_by']

  const typeOptions = ref([
    {
      label: '字符串',
      value: 'string'
    },
    {
      label: '富文本',
      value: 'richtext'
    },
    {
      label: '整型',
      value: 'int'
    },
    {
      label: '布尔值',
      value: 'bool'
    },
    {
      label: '浮点型',
      value: 'float64'
    },
    {
      label: '时间',
      value: 'time.Time'
    },
    {
      label: '枚举',
      value: 'enum'
    },
    {
      label: '单图片（字符串）',
      value: 'picture'
    },
    {
      label: '多图片（json字符串）',
      value: 'pictures'
    },
    {
      label: '视频（字符串）',
      value: 'video'
    },
    {
      label: '文件（json字符串）',
      value: 'file'
    },
    {
      label: 'JSON',
      value: 'json'
    },
    {
      label: '数组',
      value: 'array'
    }
  ])

  const typeSearchOptions = ref([
    {
      label: '=',
      value: '='
    },
    {
      label: '<>',
      value: '<>'
    },
    {
      label: '>',
      value: '>'
    },
    {
      label: '<',
      value: '<'
    },
    {
      label: 'LIKE',
      value: 'LIKE'
    },
    {
      label: 'BETWEEN',
      value: 'BETWEEN'
    },
    {
      label: 'NOT BETWEEN',
      value: 'NOT BETWEEN'
    }
  ])

  const typeIndexOptions = ref([
    {
      label: 'index',
      value: 'index'
    },
    {
      label: 'uniqueIndex',
      value: 'uniqueIndex'
    }
  ])

  const fieldTemplate = {
    fieldName: '',
    fieldDesc: '',
    fieldType: '',
    dataType: '',
    fieldJson: '',
    columnName: '',
    dataTypeLong: '',
    comment: '',
    defaultValue: '',
    require: false,
    sort: false,
    form: true,
    desc: true,
    table: true,
    excel: false,
    errorText: '',
    primaryKey: false,
    clearable: true,
    fieldSearchType: '',
    fieldIndexType: '',
    dictType: '',
    dataSource: {
      dbName: '',
      association: 1,
      table: '',
      label: '',
      value: '',
      hasDeletedAt: false
    }
  }
  const route = useRoute()
  const router = useRouter()
  const preViewCode = ref({})
  const dbform = ref({
    businessDB: '',
    dbName: '',
    tableName: ''
  })
  const tableOptions = ref([])
  const addFlag = ref('')
  const fdMap = ref({})
  const form = ref({
    structName: '',
    tableName: '',
    packageName: '',
    package: '',
    abbreviation: '',
    description: '',
    businessDB: '',
    autoCreateApiToSql: true,
    autoCreateMenuToSql: true,
    autoCreateBtnAuth: false,
    autoMigrate: true,
    gvaModel: true,
    autoCreateResource: false,
    onlyTemplate: false,
    isTree: false,
    generateWeb:true,
    generateServer:true,
    treeJson: "",
    fields: []
  })
  const rules = ref({
    structName: [
      { required: true, message: '请输入结构体名称', trigger: 'blur' }
    ],
    abbreviation: [
      { required: true, message: '请输入结构体简称', trigger: 'blur' }
    ],
    description: [
      { required: true, message: '请输入结构体描述', trigger: 'blur' }
    ],
    packageName: [
      {
        required: true,
        message: '文件名称：sysXxxxXxxx',
        trigger: 'blur'
      }
    ],
    package: [{ required: true, message: '请选择package', trigger: 'blur' }]
  })
  const dialogMiddle = ref({})
  const bk = ref({})
  const dialogFlag = ref(false)
  const previewFlag = ref(false)

  const useGva = (e) => {
    if (e && form.value.fields.length) {
      ElMessageBox.confirm(
        '如果您开启GVA默认结构，会自动添加ID,CreatedAt,UpdatedAt,DeletedAt字段，此行为将自动清除您目前在下方创建的重名字段，是否继续？',
        '注意',
        {
          confirmButtonText: '继续',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
        .then(() => {
          form.value.fields = form.value.fields.filter(
            (item) =>
              !gormModelList.some((gormfd) => gormfd === item.columnName)
          )
        })
        .catch(() => {
          form.value.gvaModel = false
        })
    }
  }

  const toLowerCaseFunc = (form, key) => {
    form[key] = toLowerCase(form[key])
  }
  const previewNode = ref(null)
  const selectText = () => {
    previewNode.value.selectText()
  }
  const copy = () => {
    previewNode.value.copy()
  }
  const editAndAddField = (item) => {
    dialogFlag.value = true
    if (item) {
      addFlag.value = 'edit'
      if (!item.dataSource) {
        item.dataSource = {
          dbName: '',
          association: 1,
          table: '',
          label: '',
          value: '',
          hasDeletedAt: false
        }
      }
      bk.value = JSON.parse(JSON.stringify(item))
      dialogMiddle.value = item
    } else {
      addFlag.value = 'add'
      fieldTemplate.onlyNumber = getOnlyNumber()
      dialogMiddle.value = JSON.parse(JSON.stringify(fieldTemplate))
    }
  }

  const fieldDialogNode = ref(null)
  const enterDialog = () => {
    fieldDialogNode.value.fieldDialogForm.validate((valid) => {
      if (valid) {
        dialogMiddle.value.fieldName = toUpperCase(dialogMiddle.value.fieldName)
        if (addFlag.value === 'add') {
          form.value.fields.push(dialogMiddle.value)
        }
        dialogFlag.value = false
      } else {
        return false
      }
    })
  }
  const closeDialog = () => {
    if (addFlag.value === 'edit') {
      dialogMiddle.value = bk.value
    }
    dialogFlag.value = false
  }
  const deleteField = (index) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      form.value.fields.splice(index, 1)
    })
  }
  const autoCodeForm = ref(null)
  const enterForm = async (isPreview) => {
    if (form.value.isTree && !form.value.treeJson){
      ElMessage({
        type: 'error',
        message: '请填写树型结构的前端展示json属性'
      })
      return false
    }
    if(!form.value.generateWeb && !form.value.generateServer){
      ElMessage({
        type: 'error',
        message: '请至少选择一个生成项'
      })
      return false
    }
    if (!form.value.onlyTemplate) {
      if (form.value.fields.length <= 0) {
        ElMessage({
          type: 'error',
          message: '请填写至少一个field'
        })
        return false
      }

      if (
        !form.value.gvaModel &&
        form.value.fields.every((item) => !item.primaryKey)
      ) {
        ElMessage({
          type: 'error',
          message: '您至少需要创建一个主键才能保证自动化代码的可行性'
        })
        return false
      }

      if (
        form.value.fields.some(
          (item) => item.fieldName === form.value.structName
        )
      ) {
        ElMessage({
          type: 'error',
          message: '存在与结构体同名的字段'
        })
        return false
      }

      if (
        form.value.fields.some((item) => item.fieldJson === form.value.package)
      ) {
        ElMessage({
          type: 'error',
          message: '存在与模板同名的的字段JSON'
        })
        return false
      }

      if (form.value.fields.some((item) => !item.fieldType)) {
        ElMessage({
          type: 'error',
          message: '请填写所有字段类型后进行提交'
        })
        return false
      }

      if (form.value.package === form.value.abbreviation) {
        ElMessage({
          type: 'error',
          message: 'package和结构体简称不可同名'
        })
        return false
      }
    }

    autoCodeForm.value.validate(async (valid) => {
      if (valid) {
        for (const key in form.value) {
          if (typeof form.value[key] === 'string') {
            form.value[key] = form.value[key].trim()
          }
        }
        form.value.structName = toUpperCase(form.value.structName)
        form.value.tableName = form.value.tableName.replace(' ', '')
        if (!form.value.tableName) {
          form.value.tableName = toSQLLine(toLowerCase(form.value.structName))
        }
        if (form.value.structName === form.value.abbreviation) {
          ElMessage({
            type: 'error',
            message: 'structName和struct简称不能相同'
          })
          return false
        }
        form.value.humpPackageName = toSQLLine(form.value.packageName)

        form.value.fields?.forEach((item) => {
          item.fieldName = toUpperCase(item.fieldName)
          if (item.fieldType === 'enum') {
            // 判断一下 item.dataTypeLong 按照,切割后的每个元素是否都使用 '' 包裹，如果没包 则修改为包裹起来的 然后再转为字符串赋值给 item.dataTypeLong
            item.dataTypeLong = item.dataTypeLong.replace(/[\[\]{}()]/g, '')
            const arr = item.dataTypeLong.split(',')
            arr.forEach((ele, index) => {
              if (ele.indexOf("'") === -1) {
                arr[index] = `'${ele}'`
              }
            })
            item.dataTypeLong = arr.join(',')
          }
        })

        delete form.value.primaryField
        if (isPreview) {
          const res = await preview({
            ...form.value,
            isAdd: !!isAdd.value,
            fields: form.value.fields.filter((item) => !item.disabled)
          })
          if(res.code !== 0){
            return
          }
          preViewCode.value = res.data.autoCode
          previewFlag.value = true
        } else {
          const res = await createTemp(form.value)
          if (res.code !== 0) {
            return
          }
          ElMessage({
            type: 'success',
            message: '自动化代码创建成功，自动移动成功'
          })
          clearCatch()
        }
      }
    })
  }

  const dbList = ref([])
  const dbOptions = ref([])

  const getDbFunc = async () => {
    dbform.value.dbName = ''
    dbform.value.tableName = ''
    const res = await getDB({ businessDB: dbform.value.businessDB })
    if (res.code === 0) {
      dbOptions.value = res.data.dbs
      dbList.value = res.data.dbList
    }
  }
  const getTableFunc = async () => {
    const res = await getTable({
      businessDB: dbform.value.businessDB,
      dbName: dbform.value.dbName
    })
    if (res.code === 0) {
      tableOptions.value = res.data.tables
    }
    dbform.value.tableName = ''
  }

  const getColumnFunc = async () => {
    const res = await getColumn(dbform.value)
    if (res.code === 0) {
      let dbtype = ''
      if (dbform.value.businessDB !== '') {
        const dbtmp = dbList.value.find(
          (item) => item.aliasName === dbform.value.businessDB
        )
        const dbraw = toRaw(dbtmp)
        dbtype = dbraw.dbtype
      }
      form.value.gvaModel = false
      const tbHump = toHump(dbform.value.tableName)
      form.value.structName = toUpperCase(tbHump)
      form.value.tableName = dbform.value.tableName
      form.value.packageName = toLowerCase(tbHump)
      form.value.abbreviation = toLowerCase(tbHump)
      form.value.description = tbHump + '表'
      form.value.autoCreateApiToSql = true
      form.value.generateServer = true
      form.value.generateWeb = true
      form.value.fields = []
      res.data.columns &&
        res.data.columns.forEach((item) => {
          if (needAppend(item)) {
            const fbHump = toHump(item.columnName)
            form.value.fields.push({
              onlyNumber: getOnlyNumber(),
              fieldName: toUpperCase(fbHump),
              fieldDesc: item.columnComment || fbHump + '字段',
              fieldType: fdMap.value[item.dataType],
              dataType: item.dataType,
              fieldJson: fbHump,
              primaryKey: item.primaryKey,
              dataTypeLong:
                item.dataTypeLong && item.dataTypeLong.split(',')[0],
              columnName:
                dbtype === 'oracle'
                  ? item.columnName.toUpperCase()
                  : item.columnName,
              comment: item.columnComment,
              require: false,
              errorText: '',
              clearable: true,
              fieldSearchType: '',
              fieldIndexType: '',
              dictType: '',
              form: true,
              table: true,
              excel: false,
              desc: true,
              dataSource: {
                dbName: '',
                association: 1,
                table: '',
                label: '',
                value: '',
                hasDeletedAt: false
              }
            })
          }
        })
    }
  }

  const needAppend = (item) => {
    let isAppend = true
    if (
      form.value.gvaModel &&
      gormModelList.some((gormfd) => gormfd === item.columnName)
    ) {
      isAppend = false
    }
    if (
      form.value.autoCreateResource &&
      dataModelList.some((datafd) => datafd === item.columnName)
    ) {
      isAppend = false
    }
    return isAppend
  }

  const setFdMap = async () => {
    const fdTypes = ['string', 'int', 'bool', 'float64', 'time.Time']
    fdTypes.forEach(async (fdtype) => {
      const res = await getDict(fdtype)
      res &&
        res.forEach((item) => {
          fdMap.value[item.label] = fdtype
        })
    })
  }
  const getAutoCodeJson = async (id) => {
    const res = await getMeta({ id: Number(id) })
    if (res.code === 0) {
      const add = route.query.isAdd
      isAdd.value = add
      form.value = JSON.parse(res.data.meta)
      if (isAdd.value) {
        form.value.fields.forEach((item) => {
          item.disabled = true
        })
      }
    }
  }

  const pkgs = ref([])
  const getPkgs = async () => {
    const res = await getPackageApi()
    if (res.code === 0) {
      pkgs.value = res.data.pkgs
    }
  }

  const goPkgs = () => {
    router.push({ name: 'autoPkg' })
  }

  const init = () => {
    getDbFunc()
    setFdMap()
    getPkgs()
    const id = route.params.id
    if (id) {
      getAutoCodeJson(id)
    }
  }
  init()

  watch(()=>form.value.generateServer,()=>{
    if(!form.value.generateServer){
      form.value.autoCreateApiToSql = false
      form.value.autoMigrate = false
    }
  })

  watch(()=>form.value.generateWeb,()=>{
    if(!form.value.generateWeb){
      form.value.autoCreateMenuToSql = false
      form.value.autoCreateBtnAuth = false
    }
  })

  const catchData = () => {
    window.sessionStorage.setItem('autoCode', JSON.stringify(form.value))
    ElMessage.success('暂存成功')
  }

  const getCatch = () => {
    const data = window.sessionStorage.getItem('autoCode')
    if (data) {
      form.value = JSON.parse(data)
    }
  }

  const clearCatch = async () => {
    form.value = {
      structName: '',
      tableName: '',
      packageName: '',
      package: '',
      abbreviation: '',
      description: '',
      businessDB: '',
      autoCreateApiToSql: true,
      autoCreateMenuToSql: true,
      autoCreateBtnAuth: false,
      autoMigrate: true,
      gvaModel: true,
      autoCreateResource: false,
      onlyTemplate: false,
      isTree: false,
      treeJson: "",
      fields: []
    }
    await nextTick()
    window.sessionStorage.removeItem('autoCode')
  }

  getCatch()

  const exportJson = () => {
    const dataStr = JSON.stringify(form.value, null, 2)
    const blob = new Blob([dataStr], { type: 'application/json' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'form_data.json'
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    URL.revokeObjectURL(url)
  }

  const importJson = (file) => {
    const reader = new FileReader()
    reader.onload = (e) => {
      try {
        form.value = JSON.parse(e.target.result)
        form.value.generateServer = true
        form.value.generateWeb = true
        ElMessage.success('JSON 文件导入成功')
      } catch (_) {
        ElMessage.error('无效的 JSON 文件')
      }
    }
    reader.readAsText(file)
    return false
  }

  watch(
    () => form.value.onlyTemplate,
    (val) => {
      if (val) {
        ElMessageBox.confirm(
          '使用基础模板将不会生成任何结构体和CURD,仅仅配置enter等属性方便自行开发非CURD逻辑',
          '注意',
          {
            confirmButtonText: '继续',
            cancelButtonText: '取消',
            type: 'warning'
          }
        )
          .then(() => {
            form.value.fields = []
          })
          .catch(() => {
            form.value.onlyTemplate = false
          })
      }
    }
  )

  const canSelect = (fieldType,item) => {
    if (fieldType === 'richtext') {
      return item !== 'LIKE';
    }

    if (fieldType !== 'string' && item === 'LIKE') {
      return true;
    }

    const nonNumericTypes = ['int', 'time.Time', 'float64'];
    if (!nonNumericTypes.includes(fieldType) && ['BETWEEN', 'NOT BETWEEN'].includes(item)) {
      return true;
    }

    return false;
  }
</script>

<style>
.no-border-collapse{
  @apply border-none;
  .el-collapse-item__header{
    @apply border-none;
  }
  .el-collapse-item__wrap{
    @apply border-none;
  }
  .el-collapse-item__content{
    @apply pb-0;
  }
}
</style>
