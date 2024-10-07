<template>
  <div>
    <warning-bar
        :title="t('view.systemTools.autoCode.autoCodeNote')"
        href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
    />
    <div class="gva-search-box">
      <div class="text-lg mb-2 text-gray-600">{{ t('view.systemTools.autoCode.createdByAI') }}</div>
      <div class="relative">
        <el-input v-model="prompt" :maxlength="100" :placeholder="t('view.systemTools.autoCode.aiCodeNote')" :rows="5"
                  resize="none" type="textarea"
        />
        <div class="flex absolute right-2 bottom-2">
          <el-tooltip
              :content="t('view.systemTools.autoCode.aiContent')"
          >
            <el-button type="primary" @click="llmAutoFunc('xiaoqi')">
              <el-icon size="18">
                <ai-gva/>
              </el-icon>
              <span>{{ t('view.systemTools.autoCode.XiaoQi') }}</span>
            </el-button>
          </el-tooltip>
          <el-tooltip
              :content="t('view.systemTools.autoCode.XiaoMiaoDesc')"
          >
            <el-button type="primary" @click="llmAutoFunc('xiaomiao')">
              <el-icon size="18">
                <ai-gva/>
              </el-icon>
              <span>{{ t('view.systemTools.autoCode.XiaoMiao') }}</span>
            </el-button>
          </el-tooltip>
        </div>
      </div>
    </div>
    <!-- 从数据库直接获取字段 -->
    <div class="gva-search-box">
      <div class="text-lg mb-2 text-gray-600">{{ t('view.systemTools.autoCode.createdFromDB') }}</div>
      <el-form
          ref="getTableForm"
          :inline="true"
          :model="dbform"
          label-width="120px"
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
                    :content="t('view.systemTools.autoCode.businessLibraryNotice')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{ t('view.systemTools.autoCode.businessLibrary') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-select
                  v-model="dbform.businessDB"
                  :placeholder="t('view.systemTools.autoCode.selectBusinessLibrary')"
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
                    <span style="float:right;color:#8492a6;font-size:13px">{{ item.dbName }}</span>
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
            <el-form-item
                class="w-full"
            >
              <div class="flex justify-end w-full">
                <el-button
                    type="primary"
                    @click="getColumnFunc"
                >
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
      <div class="text-lg mb-2 text-gray-600">{{ t('view.systemTools.autoCode.automationStructure') }}</div>
      <el-form
          ref="autoCodeForm"
          :inline="true"
          :model="form"
          :rules="rules"
          label-width="170px"
      >
        <el-row class="w-full">
          <el-col :span="6">
            <el-form-item
                :label="t('view.systemTools.autoCode.structureName')"
                class="w-full"
                prop="structName"
            >
              <el-input
                  v-model="form.structName"
                  :placeholder="t('view.systemTools.autoCode.capitalizeFirstLetterAutomatically')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
                class="w-full"
                label="TableName"
            >
              <template #label>
                <el-tooltip
                    :content="t('view.systemTools.autoCode.objectNameAndRouteGroup')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{ t('view.systemTools.autoCode.structureSimpleName') }}
                    <el-icon>
                      <QuestionFilled/>
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
                  :placeholder="t('view.systemTools.autoCode.structChineseNameNote')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
                :label=" t('view.systemTools.autoCode.tableName')"
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
          <el-col :span="6">
            <el-form-item
                class="w-full"
                prop="packageName"
            >
              <template #label>
                <el-tooltip
                    :content="t('view.systemTools.autoCode.fileNameNote')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{ t('view.systemTools.autoCode.fileName') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-input
                  v-model="form.packageName"
                  :placeholder="t('view.systemTools.autoCode.fineNameInput')"
                  @blur="toLowerCaseFunc(form,'packageName')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item
                :label="t('view.systemTools.autoCode.templateChoose')"
                class="w-full relative"
                prop="package"
            >
              <el-select
                  v-model="form.package"
                  class="w-full pr-12"
              >
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
                  <refresh/>
                </el-icon>
                <el-icon
                    class="cursor-pointer ml-2 text-gray-600"
                    @click="goPkgs"
                >
                  <document-add/>
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
                  <div> {{ t('view.systemTools.autoCode.businessLibrary') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-select
                  v-model="form.businessDB"
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
                    <span style="float:right;color:#8492a6;font-size:13px">{{ item.dbName }}</span>
                  </div>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="3">
            <el-form-item>
              <template #label>
                <el-tooltip
                    :ontent="t('view.systemTools.autoCode.useGvaNote')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{ t('view.systemTools.autoCode.groupInfos.useGvaStructure') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-checkbox
                  v-model="form.gvaModel"
                  @change="useGva"
              />
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item>
              <template #label>
                <el-tooltip
                    :content="t('view.systemTools.autoCode.groupInfos.note1')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{ t('view.systemTools.autoCode.groupInfos.autoCreateApi') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-checkbox v-model="form.autoCreateApiToSql"/>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item>
              <template #label>
                <el-tooltip
                    :content="t('view.systemTools.autoCode.groupInfos.note2')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{ t('view.systemTools.autoCode.groupInfos.autoCreateMenu') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-checkbox v-model="form.autoCreateMenuToSql"/>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item>
              <template #label>
                <el-tooltip
                    :content="t('view.systemTools.autoCode.groupInfos.note3')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{ t('view.systemTools.autoCode.groupInfos.syncTableStructure') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-checkbox v-model="form.autoMigrate"/>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item>
              <template #label>
                <el-tooltip
                    :content="t('view.systemTools.autoCode.groupInfos.note4')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{ t('view.systemTools.autoCode.groupInfos.createButtonPermissions') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-checkbox v-model="form.autoCreateBtnAuth"/>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item>
              <template #label>
                <el-tooltip
                    :content="t('view.systemTools.autoCode.groupInfos.note5')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{ t('view.systemTools.autoCode.groupInfos.createResourceIdentifier') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-checkbox v-model="form.autoCreateResource"/>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item>
              <template #label>
                <el-tooltip
                    :content="t('view.systemTools.autoCode.groupInfos.note6')"
                    effect="light"
                    placement="bottom"
                >
                  <div> {{t('view.systemTools.autoCode.groupInfos.basicTemplate') }}
                    <el-icon>
                      <QuestionFilled/>
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-checkbox v-model="form.onlyTemplate"/>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
    <!-- 组件列表 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
            :disabled="form.onlyTemplate"
            type="primary"
            @click="editAndAddField()"
        >
          {{t('view.systemTools.autoCode.addField')}}
        </el-button>
      </div>
      <div class="draggable">
        <el-table
            :data="form.fields"
            row-key="fieldName"
        >
          <el-table-column
              align="left"
              type="index"
              width="60"
          >
            <template #default>
              <el-icon class="cursor-grab drag-column">
                <MoreFilled/>
              </el-icon>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.fieldIndex')"
              type="index"
              width="60"
          />
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.primaryKey')"
              type="index"
              width="60"
          >
            <template #default="{row}">
              <el-checkbox v-model="row.primaryKey"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.fieldName')"
              prop="fieldName"
              width="160"
          >
            <template #default="{row}">
              <el-input v-model="row.fieldName"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.chineseName')"
              prop="fieldDesc"
              width="160"
          >
            <template #default="{row}">
              <el-input v-model="row.fieldDesc"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.defaultValue')"
              prop="defaultValue"
              width="160"
          >
            <template #default="{row}">
              <el-input v-model="row.defaultValue"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.required')"
              prop="require"
          >
            <template #default="{row}">
              <el-checkbox v-model="row.require"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('menu.sort')"
              prop="sort"
          >
            <template #default="{row}">
              <el-checkbox v-model="row.sort"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.createEdit')"
              prop="form"
              width="100"
          >
            <template #default="{row}">
              <el-checkbox v-model="row.form"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.table')"
              prop="table"
          >
            <template #default="{row}">
              <el-checkbox v-model="row.table"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.dictionary.sysDictionary.details')"
              prop="desc"
          >
            <template #default="{row}">
              <el-checkbox v-model="row.desc"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.importExport')"
              prop="excel"
          >
            <template #default="{row}">
              <el-checkbox v-model="row.excel"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.fieldJson')"
              prop="fieldJson"
              width="160px"
          >
            <template #default="{row}">
              <el-input v-model="row.fieldJson"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.fieldType')"
              prop="fieldType"
              width="160"
          >
            <template #default="{row}">
              <el-select
                  v-model="row.fieldType"
                  clearable
                  :placeholder="t('view.systemTools.autoCode.selectFieldType')"
                  style="width:100%"
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
            <template #default="{row}">
              <el-select
                  v-model="row.fieldIndexType"
                  clearable
                  :placeholder="t('view.systemTools.autoCode.selectIndexType')"
                  style="width:100%"
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
              :label="t('view.systemTools.autoCode.fieldLen')"
              prop="dataTypeLong"
              width="160"
          >
            <template #default="{row}">
              <el-input v-model="row.dataTypeLong"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.columnName')"
              prop="columnName"
              width="160"
          >
            <template #default="{row}">
              <el-input v-model="row.columnName"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('view.systemTools.autoCode.comment')"
              prop="comment"
              width="160"
          >
            <template #default="{row}">
              <el-input v-model="row.comment"/>
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              :label="t('general.searchCriteria')"
              prop="fieldSearchType"
              width="130"
          >
            <template #default="{row}">
              <el-select
                  v-model="row.fieldSearchType"
                  :disabled="row.fieldType!=='json'"
                  clearable
                  :placeholder="t('view.systemTools.autoCode.selectSearchCondition')"
                  style="width:100%"
              >
                <el-option
                    v-for="item in typeSearchOptions"
                    :key="item.value"
                    :disabled="
                    (row.fieldType!=='string'&&item.value==='LIKE')||
                      ((row.fieldType!=='int'&&row.fieldType!=='time.Time'&&row.fieldType!=='float64')&&(item.value==='BETWEEN' || item.value==='NOT BETWEEN'))
                  "
                    :label="item.label"
                    :value="item.value"
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

                  icon="edit"
                  link
                  type="primary"
                  @click="editAndAddField(scope.row)"
              >
                {{ t('view.systemTools.autoCode.advancedEdit') }}
              </el-button>
              <el-button
                  icon="delete"
                  link
                  type="primary"
                  @click="deleteField(scope.$index)"
              >
                {{t('general.delete')}}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!-- 组件列表 -->
      <div class="gva-btn-list justify-end mt-4">
        <el-button
            type="primary"
            @click="exportJson()"
        >
          {{ t('view.systemTools.autoCode.exportJson') }}
        </el-button>
        <el-upload
            :before-upload="importJson"
            accept=".json"
            class="flex items-center"
            show-file-list="false"
        >
          <el-button class="mx-2" type="primary">{{ t('view.systemTools.autoCode.importJson') }}</el-button>
        </el-upload>
        <el-button
            type="primary"
            @click="clearCatch()"
        >
          {{ t('view.systemTools.autoCode.clearTemp') }}
        </el-button>
        <el-button
            type="primary"
            @click="catchData()"
        >
          {{ t('view.systemTools.autoCode.temporary') }}
        </el-button>
        <el-button
            type="primary"
            @click="enterForm(true)"
        >
          {{ t('view.systemTools.autoCode.codePreview') }}
        </el-button>
        <el-button
            type="primary"
            @click="enterForm(false)"
        >
          {{ t('view.systemTools.autoCode.generateCode') }}
        </el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-drawer
        v-model="dialogFlag"
        :show-close="false"
        size="70%"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ t('view.systemTools.autoCode.componentContent') }}</span>
          <div>
            <el-button @click="closeDialog">
              {{ t('general.close') }}
            </el-button>
            <el-button
                type="primary"
                @click="enterDialog"
            >
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

    <el-drawer
        v-model="previewFlag"
        :show-close="false"
        size="80%"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ t('view.systemTools.autoCode.actionBar') }}</span>
          <div>
            <el-button
                type="primary"
                @click="selectText"
            >
              {{ t('general.selectAll') }}
            </el-button>
            <el-button
                type="primary"
                @click="copy"
            >
              {{ t('view.systemTools.autoCode.copy') }}
            </el-button>
          </div>
        </div>
      </template>
      <PreviewCodeDialog
          v-if="previewFlag"
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
import PreviewCodeDialog from '@/view/systemTools/autoCode/component/previewCodeDialg.vue'
import {toUpperCase, toHump, toSQLLine, toLowerCase} from '@/utils/stringFun'
import {createTemp, getDB, getTable, getColumn, preview, getMeta, getPackageApi, llmAuto} from '@/api/autoCode'
import {getDict} from '@/utils/dictionary'
import {ref, watch, toRaw, onMounted, nextTick} from 'vue'
import {useRoute, useRouter} from 'vue-router'
import {ElMessage, ElMessageBox} from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'
import Sortable from 'sortablejs'
import {useI18n} from 'vue-i18n' // added by mohamed hassan to support multilanguage

const {t} = useI18n() // added by mohamed hassan to support multilanguage

const getOnlyNumber = () => {
  let randomNumber = '';
  while (randomNumber.length < 16) {
    randomNumber += Math.random().toString(16).substring(2);
  }
  return randomNumber.substring(0, 16);
}

const prompt = ref("")

const llmAutoFunc = async (mode) => {
  const res = await llmAuto({prompt: prompt.value, mode: mode})
  if (res.code === 0) {
    form.value.fields = []
    const json = JSON.parse(res.data)
    for (let key in json) {
      if (key === "fields") {
        json[key].forEach(item => {
          if (item.primaryKey) {
            form.value.gvaModel = false
          }
          form.value.fields.push({
            onlyNumber: getOnlyNumber(),
            fieldName: toUpperCase(item.fieldName),
            fieldDesc: item.fieldDesc,
            fieldType: item.fieldType,
            dataType: "",
            fieldJson: item.fieldJson || item.columnName,
            primaryKey: item.primaryKey,
            dataTypeLong: item.dataTypeLong,
            columnName: item.columnName,
            comment: item.comment || item.fieldDesc,
            require: false,
            errorText: '',
            clearable: true,
            fieldSearchType: '',
            fieldIndexType: '',
            dictType: '',
            form: true,
            desc: true,
            table: true,
            excel: false,
            dataSource: {
              association: 1,
              table: '',
              label: '',
              value: ''
            }
          })
        })
      } else {
        if (mode === "xiaomiao") {
          form.value[key] = json[key]
        }
      }
    }
  }
}

// 行拖拽
const rowDrop = () => {
  // 要拖拽元素的父容器
  const tbody = document.querySelector('.draggable .el-table__body-wrapper tbody')
  Sortable.create(tbody, {
    //  可被拖拽的子元素
    draggable: '.draggable .el-table__row',
    handle: '.drag-column',
    onEnd: async ({newIndex, oldIndex}) => {
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
    label: t('fieldDialog.string'),
    value: 'string'
  },
  {
    label: t('fieldDialog.richText'),
    value: 'richtext'
  },
  {
    label: t('fieldDialog.integer'),
    value: 'int'
  },
  {
    label: t('fieldDialog.boolean'),
    value: 'bool'
  },
  {
    label: t('fieldDialog.float'),
    value: 'float64'
  },
  {
    label: t('fieldDialog.time'),
    value: 'time.Time'
  },
  {
    label: t('fieldDialog.enum'),
    value: 'enum'
  },
  {
    label: t('fieldDialog.singleImage'),
    value: 'picture',
  },
  {
    label: t('fieldDialog.multipleImages'),
    value: 'pictures',
  },
  {
    label: t('fieldDialog.video'),
    value: 'video',
  },
  {
    label: t('fieldDialog.file'),
    value: 'file',
  },
  {
    label: 'JSON',
    value: 'json',
  },
  {
    label: t('fieldDialog.array'),
    value: 'array',
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
    association: 1,
    table: '',
    label: '',
    value: ''
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
  fields: []
})
const rules = ref({
  structName: [
    {required: true, message: t('view.systemTools.autoCode.entStructName'), trigger: 'blur'}
  ],
  abbreviation: [
    {required: true, message: t('view.systemTools.autoCode.entStructAbbreviation'), trigger: 'blur'}
  ],
  description: [
    {required: true, message: t('view.systemTools.autoCode.entStructDesc'), trigger: 'blur'}
  ],
  packageName: [
    {
      required: true,
      message: t('view.systemTools.autoCode.entFileName'),
      trigger: 'blur'
    }
  ],
  package: [
    {required: true, message: t('view.systemTools.autoCode.selectPackage'), trigger: 'blur'}
  ]
})

const dialogMiddle = ref({})
const bk = ref({})
const dialogFlag = ref(false)
const previewFlag = ref(false)

const useGva = (e) => {
  if (e && form.value.fields.length) {
    ElMessageBox.confirm(
        t('view.systemTools.gvaStructureNote'),
        t('view.systemTools.note'),
        {
          confirmButtonText: t('general.continue'),
          cancelButtonText: t('general.cancel'),
          type: 'warning',
        }
    )
        .then(() => {
          form.value.fields = form.value.fields.filter(item => !gormModelList.some(gormfd => gormfd === item.columnName))
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
        association: 1,
        table: '',
        label: '',
        value: ''
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
  fieldDialogNode.value.fieldDialogForm.validate(valid => {
    if (valid) {
      dialogMiddle.value.fieldName = toUpperCase(
          dialogMiddle.value.fieldName
      )
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

  if (!form.value.onlyTemplate) {

    if (form.value.fields.length <= 0) {
      ElMessage({
        type: 'error',
        message: t('view.systemTools.autoCode.errNoFields')
      })
      return false
    }

    if (!form.value.gvaModel && form.value.fields.every(item => !item.primaryKey)) {
      ElMessage({
        type: 'error',
        message:  t('view.systemTools.autoCode.primaryKeyRequirement')
      })
      return false
    }

    if (
        form.value.fields.some(item => item.fieldName === form.value.structName)
    ) {
      ElMessage({
        type: 'error',
        message: t('view.systemTools.autoCode.errSameFiledName')
      })
      return false
    }


    if (form.value.fields.some(item => !item.fieldType)) {
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


  autoCodeForm.value.validate(async valid => {
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
      delete form.value.primaryField
      if (isPreview) {
        const data = await preview(form.value)
        preViewCode.value = data.data.autoCode
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
  const res = await getDB({businessDB: dbform.value.businessDB})
  if (res.code === 0) {
    dbOptions.value = res.data.dbs
    dbList.value = res.data.dbList
  }
}
const getTableFunc = async () => {
  const res = await getTable({businessDB: dbform.value.businessDB, dbName: dbform.value.dbName})
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
      const dbtmp = dbList.value.find(item => item.aliasName === dbform.value.businessDB)
      const dbraw = toRaw(dbtmp)
      dbtype = dbraw.dbtype
    }
    form.value.gvaModel = false
    const tbHump = toHump(dbform.value.tableName)
    form.value.structName = toUpperCase(tbHump)
    form.value.tableName = dbform.value.tableName
    form.value.packageName = tbHump
    form.value.abbreviation = tbHump
    form.value.description = tbHump + t('view.systemTools.autoCode.table')
    form.value.autoCreateApiToSql = true
    form.value.fields = []
    res.data.columns &&
    res.data.columns.forEach(item => {
      if (needAppend(item)) {
        const fbHump = toHump(item.columnName)
        form.value.fields.push({
          onlyNumber: getOnlyNumber(),
          fieldName: toUpperCase(fbHump),
          fieldDesc: item.columnComment || fbHump + t('view.systemTools.autoCode.field'),
          fieldType: fdMap.value[item.dataType],
          dataType: item.dataType,
          fieldJson: fbHump,
          primaryKey: item.primaryKey,
          dataTypeLong: item.dataTypeLong && item.dataTypeLong.split(',')[0],
          columnName: dbtype === 'oracle' ? item.columnName.toUpperCase() : item.columnName,
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
            association: 1,
            table: '',
            label: '',
            value: ''
          }
        })
      }
    })
  }
}

const needAppend = (item) => {
  let isAppend = true
  if (form.value.gvaModel && gormModelList.some(gormfd => gormfd === item.columnName)) {
    isAppend = false
  }
  if (form.value.autoCreateResource && dataModelList.some(datafd => datafd === item.columnName)) {
    isAppend = false
  }
  return isAppend
}

const setFdMap = async () => {
  const fdTypes = ['string', 'int', 'bool', 'float64', 'time.Time']
  fdTypes.forEach(async fdtype => {
    const res = await getDict(fdtype)
    res && res.forEach(item => {
      fdMap.value[item.label] = fdtype
    })
  })
}
const getAutoCodeJson = async (id) => {
  const res = await getMeta({id: Number(id)})
  if (res.code === 0) {
    form.value = JSON.parse(res.data.meta)
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
  router.push({name: 'autoPkg'})
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

watch(() => route.params.id, () => {
  if (route.name === 'autoCodeEdit') {
    init()
  }
})


const catchData = () => {
  window.sessionStorage.setItem('autoCode', JSON.stringify(form.value))
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
    fields: []
  }
  await nextTick()
  window.sessionStorage.removeItem('autoCode')
}

getCatch()

const exportJson = () => {
  const dataStr = JSON.stringify(form.value, null, 2)
  const blob = new Blob([dataStr], {type: 'application/json'})
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
      ElMessage.success(t('view.systemTools.autoCode.jsonImportSuccess'))
    } catch (error) {
      ElMessage.error(t('view.systemTools.autoCode.invalidJsonFile'))
    }
  }
  reader.readAsText(file)
  return false
}

watch(() => form.value.onlyTemplate, (val) => {
  if (val) {
    ElMessageBox.confirm(t('view.systemTools.autoCode.basicTemplateNote'), t('view.systemTools.autoCode.note'), {
      confirmButtonText: t('general.continue'),
      cancelButtonText: t('general.cancel'),
      type: 'warning',
    })
        .then(() => {
          form.value.fields = []
        })
        .catch(() => {
          form.value.onlyTemplate = false
        })
  }
})

</script>
