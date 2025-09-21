<template>
  <div>
    <warning-bar
      :title="t('view.systemTools.autoCode.autoCodeNote')"
      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
    />
    <div class="gva-search-box" v-if="!isAdd">
      <div class="text-lg mb-2 text-gray-600">
        {{ t('view.systemTools.autoCode.createdByAI')
        }}<a
          class="text-blue-600 text-sm ml-4"
          href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
          target="_blank"
          >{{ t('view.systemTools.autoCode.getAiPath') }}</a
        >
      </div>
      <div class="relative">
        <el-input
          v-model="prompt"
          type="textarea"
          :rows="5"
          :maxlength="2000"
          :placeholder="t('view.systemTools.autoCode.aiCodeNote')"
          resize="none"
          @focus="handleFocus"
          @blur="handleBlur"
        />

        <div class="flex absolute bottom-2" style="right:9rem !important;">
          <el-tooltip effect="light">
            <template #content>
              <div>
                {{ t('view.systemTools.autoCode.aiNote1') }}<a
                  class="text-blue-600"
                  href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
                  target="_blank"
              >{{ t('view.systemTools.autoCode.aiNote2') }}</a
              >{{ t('view.systemTools.autoCode.aiNote3') }}
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
              {{ t('view.systemTools.autoCode.imageRecognition') }}
            </el-button>
          </el-tooltip>
        </div>

        <div class="flex absolute right-2 bottom-2">
          <el-tooltip effect="light">
            <template #content>
              <div>
                {{ t('view.systemTools.autoCode.aiNote1')
                }}<a
                  class="text-blue-600"
                  href="https://plugin.gin-vue-admin.com/#/layout/userInfo/center"
                  target="_blank"
                  >{{ t('view.systemTools.autoCode.aiNote2') }}</a
                >{{ t('view.systemTools.autoCode.aiNote3') }}
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
              {{ t('view.systemTools.autoCode.generate') }}
            </el-button>
          </el-tooltip>
        </div>
      </div>
    </div>   
    <!-- 从数据库直接获取字段 -->
    <div class="gva-search-box" v-if="!isAdd">
      <div class="text-lg mb-2 text-gray-600">
        {{ t('view.systemTools.autoCode.createdFromDB') }}
      </div>
      <el-form
        ref="getTableForm"
        :inline="true"
        :model="dbform"
        label-width="140px"
      >
        <el-row class="w-full">
          <el-col :span="6">
            <el-form-item
              :label="t('view.systemTools.autoCode.businessLibrary')"
              class="w-full"
              prop="selectDBtype"
            >
              <template #label>
                <el-tooltip
                  :content="
                    t('view.systemTools.autoCode.businessLibraryNotice')
                  "
                  effect="light"
                  placement="bottom"
                >
                  <div>
                    {{ t('view.systemTools.autoCode.businessLibrary') }}
                    <el-icon>
                      <QuestionFilled />
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-select
                v-model="dbform.businessDB"
                :placeholder="
                  t('view.systemTools.autoCode.selectBusinessLibrary')
                "
                class="w-full"
                clearable
                @change="getDbFunc"
              >
                <el-option
                  v-for="item in dbList"
                  :key="item.aliasName"
                  :disabled="item.disable"
                  :label="item.aliasName"
                  :value="item.aliasName"
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
            <el-form-item
              :label="t('view.systemTools.autoCode.dbName')"
              class="w-full"
              prop="structName"
            >
              <el-select
                v-model="dbform.dbName"
                :placeholder="t('view.systemTools.autoCode.selectDB')"
                class="w-full"
                clearable
                filterable
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
            <el-form-item
              :label="t('view.systemTools.autoCode.tableName')"
              class="w-full"
              prop="structName"
            >
              <el-select
                v-model="dbform.tableName"
                :disabled="!dbform.dbName"
                :placeholder="t('view.systemTools.autoCode.selectTable')"
                class="w-full"
                filterable
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
                  {{ t('view.systemTools.autoCode.selectTableBtn') }}
                </el-button>
              </div>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
    <div class="gva-search-box">
      <!-- 初始版本自动化代码工具 -->
      <div class="text-lg mb-2 text-gray-600">
        {{ t('view.systemTools.autoCode.automationStructure') }}
      </div>
      <el-form
        ref="autoCodeForm"
        :disabled="isAdd"
        :inline="true"
        :model="form"
        :rules="rules"
        label-width="170px"
      >
        <el-row class="w-full">
          <el-col :span="7">
            <el-form-item
              :label="t('view.systemTools.autoCode.structureName')"
              class="w-full"
              prop="structName"
            >
              <div class="flex gap-2 w-full">
                <el-input
                  v-model="form.structName"
                  :placeholder="t('view.systemTools.autoCode.capitalizeFirstLetterAutomatically')"
                />
                <el-button
                  :disabled="form.onlyTemplate"
                  type="primary"
                  @click="llmAutoFunc(true)"
                >
                  <el-icon size="18">
                    <ai-gva />
                  </el-icon>
                  {{ t('view.systemTools.autoCode.generate') }}
                </el-button>
              </div>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="abbreviation" prop="abbreviation" class="w-full">
              <template #label>
                <el-tooltip
                  :content="
                    t('view.systemTools.autoCode.objectNameAndRouteGroup')
                  "
                  effect="light"
                  placement="bottom"
                >
                  <div>
                    {{ t('view.systemTools.autoCode.structureSimpleName') }}
                    <el-icon>
                      <QuestionFilled />
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-input
                v-model="form.abbreviation"
                :placeholder="t('view.systemTools.autoCode.structNameInput')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              :label="t('view.systemTools.autoCode.StructureOverview')"
              class="w-full"
              prop="description"
            >
              <el-input
                v-model="form.description"
                :placeholder="
                  t('view.systemTools.autoCode.structChineseNameNote')
                "
              />
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item
              :label="t('view.systemTools.autoCode.tableName')"
              class="w-full"
              prop="tableName"
            >
              <el-input
                v-model="form.tableName"
                :placeholder="t('view.systemTools.autoCode.tableNameNote')"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="w-full">
          <el-col :span="7">
            <el-form-item class="w-full" prop="packageName">
              <template #label>
                <el-tooltip
                  :content="t('view.systemTools.autoCode.fileNameNote')"
                  effect="light"
                  placement="bottom"
                >
                  <div>
                    {{ t('view.systemTools.autoCode.fileName') }}
                    <el-icon>
                      <QuestionFilled />
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-input
                v-model="form.packageName"
                :placeholder="t('view.systemTools.autoCode.fineNameInput')"
                @blur="toLowerCaseFunc(form, 'packageName')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
              :label="t('view.systemTools.autoCode.templateChoose')"
              class="w-full relative"
              prop="package"
            >
              <el-select v-model="form.package" class="w-full pr-12" filterable>
                <el-option
                  v-for="item in pkgs"
                  :key="item.ID"
                  :label="item.packageName"
                  :value="item.packageName"
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
            <el-form-item
              :label="t('view.systemTools.autoCode.businessLibrary')"
              class="w-full"
              prop="businessDB"
            >
              <template #label>
                <el-tooltip
                  :content="t('view.systemTools.autoCode.libraryNote')"
                  effect="light"
                  placement="bottom"
                >
                  <div>
                    {{ t('view.systemTools.autoCode.businessLibrary') }}
                    <el-icon>
                      <QuestionFilled />
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-select
                v-model="form.businessDB"
                clearable
                :placeholder="t('view.systemTools.autoCode.selectBusinessLibrary')"
                class="w-full"
              >
                <el-option
                  v-for="item in dbList"
                  :key="item.aliasName"
                  :disabled="item.disable"
                  :label="item.aliasName"
                  :value="item.aliasName"
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
              {{ t('view.systemTools.autoCode.expertMode') }}
            </div>
          </template>
          <template #icon="{ isActive }">
          <span class="text-lg ml-auto mr-4 font-normal">
            {{ isActive ? t('general.collapse') : t('general.expand') }}
          </span>
          </template>
          <div class="p-4">
            <!-- 基础设置组 -->
            <div class="border-b border-gray-200 last:border-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">{{ t('view.systemTools.autoCode.basicSettings') }}</h3>
              <el-row :gutter="20">
                <el-col :span="3">
                  <el-tooltip
                      :content="t('view.systemTools.autoCode.useGvaNote')"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item :label="t('view.systemTools.autoCode.groupInfos.useGvaStructure')">
                      <el-checkbox v-model="form.gvaModel" @change="useGva" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                      :content="t('view.systemTools.autoCode.groupInfos.note4')"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item :label="t('view.systemTools.autoCode.groupInfos.createButtonPermissions')">
                      <el-checkbox :disabled="!form.generateWeb" v-model="form.autoCreateBtnAuth" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-form-item :label="t('view.systemTools.autoCode.generateFrontend')">
                    <el-checkbox v-model="form.generateWeb" />
                  </el-form-item>
                </el-col>
                <el-col :span="3">
                  <el-form-item :label="t('view.systemTools.autoCode.generateBackend')">
                    <el-checkbox disabled v-model="form.generateServer" />
                  </el-form-item>
                </el-col>
              </el-row>
            </div>

            <!-- 自动化设置组 -->
            <div class="border-b border-gray-200 last:border-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">{{ t('view.systemTools.autoCode.automationSettings') }}</h3>
              <el-row :gutter="20">
                <el-col :span="3">
                  <el-tooltip
                      :content="t('view.systemTools.autoCode.groupInfos.note1')"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item :label="t('view.systemTools.autoCode.groupInfos.autoCreateApi')">
                      <el-checkbox  :disabled="!form.generateServer" v-model="form.autoCreateApiToSql" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                      :content="t('view.systemTools.autoCode.groupInfos.note2')"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item :label="t('view.systemTools.autoCode.groupInfos.autoCreateMenu')">
                      <el-checkbox :disabled="!form.generateWeb" v-model="form.autoCreateMenuToSql" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                      :content="t('view.systemTools.autoCode.groupInfos.note3')"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item :label="t('view.systemTools.autoCode.groupInfos.syncTableStructure')">
                      <el-checkbox  :disabled="!form.generateServer" v-model="form.autoMigrate" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
              </el-row>
            </div>

            <!-- 高级设置组 -->
            <div class="border-b border-gray-200 last:border-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">{{ t('view.systemTools.autoCode.advancedSettings') }}</h3>
              <el-row :gutter="20">
                <el-col :span="3">
                  <el-tooltip
                      :content="t('view.systemTools.autoCode.groupInfos.note5')"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item :label="t('view.systemTools.autoCode.groupInfos.createResourceIdentifier')">
                      <el-checkbox v-model="form.autoCreateResource" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
                <el-col :span="3">
                  <el-tooltip
                      :content="t('view.systemTools.autoCode.groupInfos.note6')"
                      placement="top"
                      effect="light"
                  >
                    <el-form-item :label="t('view.systemTools.autoCode.groupInfos.basicTemplate')">
                      <el-checkbox v-model="form.onlyTemplate" />
                    </el-form-item>
                  </el-tooltip>
                </el-col>
              </el-row>
            </div>

            <!-- 树形结构设置 -->
            <div class="last:pb-0">
              <h3 class="text-lg font-medium mb-4 text-gray-700">{{ t('view.systemTools.autoCode.treeStructureSettings') }}</h3>
              <el-row :gutter="20" align="middle">
                <el-col :span="24">
                    <el-form-item :label="t('view.systemTools.autoCode.groupInfos.treeStructure')">
                      <div class="flex items-center gap-4">
                        <el-tooltip
                            :content="t('view.systemTools.autoCode.groupInfos.parentIdNote')"
                            placement="top"
                            effect="light"
                        >
                          <el-checkbox v-model="form.isTree" />
                        </el-tooltip>
                        <el-input
                            v-model="form.treeJson"
                            :disabled="!form.isTree"
                            :placeholder="t('view.systemTools.autoCode.groupInfos.frontendJsonAttr')"
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
          :disabled="form.onlyTemplate"
          type="primary"
          @click="editAndAddField()"
        >
          {{ t('view.systemTools.autoCode.addField') }}
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
            :label="t('view.systemTools.autoCode.fieldIndex')"
            width="80"
          />
          <el-table-column
            fixed="left"
            align="left"
            type="index"
            :label="t('view.systemTools.autoCode.primaryKey')"
            width="120"
          >
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.primaryKey" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.systemTools.autoCode.fieldName')"
            prop="fieldName"
            width="160"
          >
            <template #default="{ row }">
              <el-input disabled v-model="row.fieldName" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.systemTools.autoCode.chineseName')"
            prop="fieldDesc"
            width="160"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.fieldDesc" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.systemTools.autoCode.defaultValue')"
            prop="defaultValue"
            width="160"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.defaultValue" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.systemTools.autoCode.required')"
            prop="require"
            width="100"
          >
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.require" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.superAdmin.menu.sort')"
            prop="sort"
          >
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.sort" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.systemTools.autoCode.createEdit')"
            prop="form"
            width="100"
          >
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.form" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.systemTools.autoCode.table')"
            prop="table"
          >
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.table" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.dictionary.sysDictionary.details')"
            prop="desc"
          >
            <template #default="{ row }">
              <el-checkbox :disabled="row.disabled" v-model="row.desc" />
            </template>
          </el-table-column>
          <el-table-column
            v-if="!isAdd"
            align="left"
            prop="excel"
            width="120"
            :label="t('view.systemTools.autoCode.importExport')"
          >
            <template #default="{ row }">
              <el-checkbox v-model="row.excel" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.systemTools.autoCode.fieldJson')"
            prop="fieldJson"
            width="160px"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.fieldJson" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.systemTools.autoCode.fieldType')"
            prop="fieldType"
            width="160"
          >
            <template #default="{ row }">
              <el-select
                v-model="row.fieldType"
                clearable
                :placeholder="t('view.systemTools.autoCode.selectFieldType')"
                :disabled="row.disabled"
                style="width: 100%"
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
            :label="t('view.systemTools.autoCode.indexType')"
            prop="fieldIndexType"
            width="160"
          >
            <template #default="{ row }">
              <el-select
                v-model="row.fieldIndexType"
                clearable
                :placeholder="t('view.systemTools.autoCode.selectIndexType')"
                style="width: 100%"
                :disabled="row.disabled"
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
            :label="t('view.systemTools.autoCode.columnName')"
            prop="columnName"
            width="160"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.columnName" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('view.systemTools.autoCode.comment')"
            prop="comment"
            width="160"
          >
            <template #default="{ row }">
              <el-input :disabled="row.disabled" v-model="row.comment" />
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            :label="t('general.searchCriteria')"
            prop="fieldSearchType"
            width="130"
          >
            <template #default="{ row }">
              <el-select
                v-model="row.fieldSearchType"
                style="width: 100%"
                :placeholder="
                  t('view.systemTools.autoCode.selectSearchCondition')
                "
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
          <el-table-column
            align="left"
            fixed="right"
            :label="t('components.commandMenu.operate')"
            width="300"
          >
            <template #default="scope">
              <el-button
                v-if="!scope.row.disabled"
                icon="edit"
                link
                type="primary"
                @click="editAndAddField(scope.row)"
              >
                {{ t('view.systemTools.autoCode.advancedEdit') }}
              </el-button>
              <el-button
                v-if="!scope.row.disabled"
                icon="delete"
                link
                type="primary"
                @click="deleteField(scope.$index)"
              >
                {{ t('general.delete') }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!-- 组件列表 -->
      <div class="gva-btn-list justify-end mt-4">
        <el-button type="primary" :disabled="isAdd" @click="exportJson()">
          {{ t('view.systemTools.autoCode.exportJson') }}
        </el-button>
        <el-upload
          :before-upload="importJson"
          :show-file-list="false"
          :headers="{'x-token': token}"
          accept=".json"
          class="flex items-center"
        >
          <el-button class="mx-2" type="primary" :disabled="isAdd">{{
            t('view.systemTools.autoCode.importJson')
          }}</el-button>
        </el-upload>
        <el-button type="primary" :disabled="isAdd" @click="clearCatch()">
          {{ t('view.systemTools.autoCode.clearTemp') }}
        </el-button>
        <el-button type="primary" :disabled="isAdd" @click="catchData()">
          {{ t('view.systemTools.autoCode.temporary') }}
        </el-button>
        <el-button type="primary" :disabled="isAdd" @click="enterForm(false)">
          {{ t('view.systemTools.autoCode.generateCode') }}
        </el-button>
        <el-button type="primary" @click="enterForm(true)">
          {{ t('view.systemTools.autoCode.codePreview') }}
        </el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-drawer v-model="dialogFlag" :show-close="false" size="70%">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{
            t('view.systemTools.autoCode.componentContent')
          }}</span>
          <div>
            <el-button @click="closeDialog">
              {{ t('general.close') }}
            </el-button>
            <el-button type="primary" @click="enterDialog">
              {{ t('general.confirm') }}
            </el-button>
          </div>
        </div>
      </template>

      <FieldDialog
        v-if="dialogFlag"
        ref="fieldDialogNode"
        :dialog-middle="dialogMiddle"
        :type-index-options="typeIndexOptions"
        :type-options="typeOptions"
        :type-search-options="typeSearchOptions"
      />
    </el-drawer>

    <el-drawer v-model="previewFlag" :show-close="false" size="80%">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{
            t('view.systemTools.autoCode.actionBar')
          }}</span>
          <div>
            <el-button type="primary" @click="selectText">
              {{ t('general.selectAll') }}
            </el-button>
            <el-button type="primary" @click="copy">
              {{ t('view.systemTools.autoCode.copy') }}
            </el-button>
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

<script>
  export default {
    name: 'AutoCode'
  }
</script>

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
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

  const { t } = useI18n() // added by mohamed hassan to support multilingual
  
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
      ElMessage.error(t('view.systemTools.autoPkg.entStructName'))
      return
    }
    if (!flag && !prompt.value) {
      ElMessage.error(t('general.enterDescription'))
      return
    }

    if (form.value.fields.length > 0) {
      const res = await ElMessageBox.confirm(
        t('view.systemTools.autoCode.aiClearDataNote'),
        t('general.hint'),
        {
          confirmButtonText: t('general.confirm'),
          cancelButtonText: t('general.cancel'),
          type: 'warning'
        }
      )
      if (res !== 'confirm') {
        return
      }
    }

    const res = await llmAuto({
      prompt: flag
        ? t('view.systemTools.autoPkg.structNameIs') + form.value.structName
        : prompt.value
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
      label: t('view.systemTools.autoCode.fieldDialog.string'),
      value: 'string'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.richText'),
      value: 'richtext'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.integer'),
      value: 'int'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.boolean'),
      value: 'bool'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.float'),
      value: 'float64'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.time'),
      value: 'time.Time'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.enum'),
      value: 'enum'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.singleImage'),
      value: 'picture'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.multipleImages'),
      value: 'pictures'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.video'),
      value: 'video'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.file'),
      value: 'file'
    },
    {
      label: 'JSON',
      value: 'json'
    },
    {
      label: t('view.systemTools.autoCode.fieldDialog.array'),
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
      {
        required: true,
        message: t('view.systemTools.autoCode.entStructName'),
        trigger: 'blur'
      }
    ],
    abbreviation: [
      {
        required: true,
        message: t('view.systemTools.autoCode.entStructAbbreviation'),
        trigger: 'blur'
      }
    ],
    description: [
      {
        required: true,
        message: t('view.systemTools.autoCode.entStructDesc'),
        trigger: 'blur'
      }
    ],
    packageName: [
      {
        required: true,
        message: t('view.systemTools.autoCode.entFileName'),
        trigger: 'blur'
      }
    ],
    package: [
      {
        required: true,
        message: t('view.systemTools.autoCode.selectPackage'),
        trigger: 'blur'
      }
    ]
  })

  const dialogMiddle = ref({})
  const bk = ref({})
  const dialogFlag = ref(false)
  const previewFlag = ref(false)

  const useGva = (e) => {
    if (e && form.value.fields.length) {
      ElMessageBox.confirm(
        t('view.systemTools.autoCode.gvaStructureNote'),
        t('view.systemTools.autoCode.note'),
        {
          confirmButtonText: t('general.continue'),
          cancelButtonText: t('general.cancel'),
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
    ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
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
        message: t('view.systemTools.autoCode.fillJsonDataNote')
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
          message: t('view.systemTools.autoCode.errNoFields')
        })
        return false
      }

      if (
        !form.value.gvaModel &&
        form.value.fields.every((item) => !item.primaryKey)
      ) {
        ElMessage({
          type: 'error',
          message: t('view.systemTools.autoCode.primaryKeyRequirement')
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
          message: t('view.systemTools.autoCode.errSameFiledName')
        })
        return false
      }

      if (
        form.value.fields.some((item) => item.fieldJson === form.value.package)
      ) {
        ElMessage({
          type: 'error',
          message: t('view.systemTools.autoCode.errJsonFieldNameAsTemplate')
        })
        return false
      }

      if (form.value.fields.some((item) => !item.fieldType)) {
        ElMessage({
          type: 'error',
          message: t('view.systemTools.autoCode.fillFieldTypes')
        })
        return false
      }

      if (form.value.package === form.value.abbreviation) {
        ElMessage({
          type: 'error',
          message: t('view.systemTools.autoCode.packageNameConflict')
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
            message: t('view.systemTools.autoCode.errSameStructDescAbbr')
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
            message: t('view.systemTools.autoCode.codeGenMoveSuccess')
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
      form.value.description = tbHump + ' ' + t('view.systemTools.autoCode.table')
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
              fieldDesc:
                item.columnComment ||
                fbHump + t('view.systemTools.autoCode.field'),
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
      console.log(add)
      isAdd.value = add
      form.value = JSON.parse(res.data.meta)
      if (isAdd.value) {
        form.value.fields.forEach((item) => {
          item.disabled = true
        })
      }

      console.log(form.value.fields)
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
        ElMessage.success(t('view.systemTools.autoCode.jsonImportSuccess'))
      } catch (_) {
        ElMessage.error(t('view.systemTools.autoCode.invalidJsonFile'))
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
          t('view.systemTools.autoCode.basicTemplateNote'),
          t('view.systemTools.autoCode.note'),
          {
            confirmButtonText: t('general.continue'),
            cancelButtonText: t('general.cancel'),
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
