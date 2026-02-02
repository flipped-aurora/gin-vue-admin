<template>
  <div class="h-full">
    <warning-bar
        href="https://plugin.gin-vue-admin.com/license"
        title="此功能仅在开发阶段使用，用户构建本项目内的skills技能库。"
    />
    <el-row :gutter="12" class="h-full">
      <el-col :xs="24" :sm="8" :md="6" :lg="5" class="flex flex-col gap-4 h-full">
        <el-card shadow="never" class="!border-none shrink-0">
          <div class="font-bold mb-2">AI 工具</div>
          <div class="flex flex-wrap gap-2">
            <div
              v-for="tool in tools"
              :key="tool.key"
              class="px-3 py-1.5 rounded-md text-sm cursor-pointer transition-all border select-none"
              :class="activeTool === tool.key
                ? 'bg-[var(--el-color-primary)] text-white border-[var(--el-color-primary)] shadow-sm'
                : 'bg-white hover:bg-gray-50 text-gray-700 border-gray-200 dark:bg-gray-800 dark:text-gray-300 dark:border-gray-700 dark:hover:bg-gray-700'"
              @click="handleToolSelect(tool.key)"
            >
              {{ tool.label }}
            </div>
          </div>
        </el-card>

        <el-card shadow="never" class="!border-none shrink-0">
          <div class="flex justify-between items-center mb-2">
            <span class="font-bold">全局约束</span>
            <el-button type="primary" link icon="Edit" @click="openGlobalConstraint">编辑</el-button>
          </div>
          <div class="text-xs text-gray-500">路径: {{ globalConstraintPath }}</div>
        </el-card>

        <el-card shadow="never" class="!border-none flex-1 mt-2 flex flex-col min-h-0">
          <div class="flex justify-between items-center mb-2">
            <span class="font-bold">Skills</span>
            <el-button type="primary" link icon="Plus" @click="openCreateDialog">新增</el-button>
          </div>
          <el-input
            v-model="skillFilter"
            size="small"
            clearable
            placeholder="搜索技能"
            class="mb-2"
            prefix-icon="Search"
          />
          <el-scrollbar class="h-[calc(100vh-380px)]">
            <el-menu :default-active="activeSkill" class="!border-none" @select="handleSkillSelect">
              <el-menu-item
                v-for="skill in filteredSkills"
                :key="skill"
                :index="skill"
                class="!h-10 !leading-10 !my-1 !mx-1 !rounded-[4px]"
              >
                <el-icon><Document /></el-icon>
                <span class="truncate" :title="skill">{{ skill }}</span>
              </el-menu-item>
            </el-menu>
          </el-scrollbar>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="16" :md="18" :lg="19" class="h-full">
        <el-card shadow="never" class="!border-none h-full flex flex-col">
          <template v-if="!activeSkill">
            <div class="h-full flex items-center justify-center">
              <el-empty description="请选择或新建一个技能" />
            </div>
          </template>
          <template v-else>
            <div class="flex justify-between items-center mb-4 pb-4 border-b border-gray-100 dark:border-gray-800">
              <div class="text-lg font-bold flex items-center gap-2">
                <span>{{ activeSkill }}</span>
                <el-tag size="small" type="info">Skill</el-tag>
              </div>
              <el-button type="primary" icon="Check" @click="saveCurrentSkill">保存配置</el-button>
            </div>

            <el-tabs v-model="activeTab" class="h-full">
              <el-tab-pane label="技能配置" name="config">
                <div
                  class="mt-4 mb-4 rounded-md border border-gray-100 dark:border-gray-800 bg-gray-50 dark:bg-gray-900/30 p-3 text-xs text-gray-600 dark:text-gray-300"
                >
                  <div class="font-medium text-gray-700 dark:text-gray-200 mb-2">填写指引</div>
                  <ul class="list-disc pl-4 space-y-1">
                    <li>Name 用小写+连字符（kebab-case），作为技能目录名与 /slash 命令。</li>
                    <li>Description 写清触发条件与关键词，这是自动匹配的核心。</li>
                    <li>正文建议按 Instructions / Examples / Guidelines 组织，写清输入、输出与约束。</li>
                    <li>需要模板/规范/脚本时，在正文中引用 templates/...、references/...、scripts/...</li>
                  </ul>
                </div>
                <el-form :model="form" label-width="160px">
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Name
                        <el-tooltip content="唯一标识，建议小写+连字符（kebab-case）" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <el-input v-model="form.name" placeholder="例如: code-comment-expert" />
                    <div class="text-xs text-gray-400 mt-1">建议 2-4 个单词，避免空格与中文。</div>
                  </el-form-item>
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Description
                        <el-tooltip content="写清使用时机/触发条件与关键字，这是最重要的一行" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <el-input
                      v-model="form.description"
                      placeholder="例如: 为代码添加双语注释，适合代码审查/重构/可读性改进"
                    />
                    <div class="text-xs text-gray-400 mt-1">建议包含：任务类型、触发场景、关键词。</div>
                  </el-form-item>
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Allowed Tools
                        <el-tooltip content="限制可用工具范围，例如: Bash(gh *), Read, Write" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <el-input v-model="form.allowedTools" placeholder="可选，例如: Bash(gh *), Read, Write" />
                    <div class="text-xs text-gray-400 mt-1">可选字段，留空后保存会移除</div>
                  </el-form-item>
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Context
                        <el-tooltip content="fork 表示独立上下文，适合复杂任务" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <el-input v-model="form.context" placeholder="可选，例如: fork" />
                    <div class="text-xs text-gray-400 mt-1">可选字段，留空后保存会移除</div>
                  </el-form-item>
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Agent
                        <el-tooltip content="context=fork 时可指定子代理，例如: Explore" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <el-input v-model="form.agent" placeholder="可选，例如: Explore / Build" />
                    <div class="text-xs text-gray-400 mt-1">可选字段，留空后保存会移除</div>
                  </el-form-item>
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Markdown 内容
                        <el-tooltip content="正文建议精简，复杂细节可拆到 templates/references/scripts" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <div class="mb-2 flex flex-wrap gap-2">
                      <el-button
                        v-for="block in quickBlocks"
                        :key="block.label"
                        size="small"
                        @click="appendMarkdown(block.content)"
                      >
                        {{ block.label }}
                      </el-button>
                      <el-button size="small" @click="insertFullTemplate">插入完整模板</el-button>
                    </div>
                    <el-input
                      v-model="form.markdown"
                      type="textarea"
                      :rows="20"
                      :placeholder="markdownPlaceholder"
                    />
                    <div class="text-xs text-gray-400 mt-1">
                      建议精简正文，把细节放到 templates/references/scripts，并在正文中通过相对路径引用。
                    </div>
                  </el-form-item>
                </el-form>
              </el-tab-pane>

              <el-tab-pane label="脚本" name="scripts" class="mt-4">
                <div class="flex justify-between items-center mb-4">
                  <div class="text-sm text-gray-500 bg-gray-50 dark:bg-gray-800 px-3 py-1 rounded">路径: scripts/</div>
                  <el-button type="primary" icon="Plus" size="small" @click="openScriptDialog">创建脚本</el-button>
                </div>
                <div class="text-xs text-gray-500 mb-3">
                  适合放可执行逻辑或校验流程，在正文中引用 <span class="font-mono">scripts/文件名</span> 使用（运行需启用 code execution）。
                </div>
                <el-table :data="scriptRows" style="width: 100%">
                  <el-table-column prop="name" label="文件名">
                    <template #default="scope">
                      <div class="flex items-center gap-2">
                        <el-icon><Document /></el-icon>
                        <span>{{ scope.row.name }}</span>
                      </div>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="180">
                    <template #default="scope">
                      <el-button type="primary" link icon="Edit" @click="openScriptEditor(scope.row.name)">编辑</el-button>
                      <el-button type="primary" link @click="insertFileSnippet('script', scope.row.name)">调用</el-button>
                    </template>
                  </el-table-column>
                </el-table>
                <el-empty v-if="scriptRows.length === 0" description="暂无脚本" />
              </el-tab-pane>

              <el-tab-pane label="资源" name="resources">
                <div class="flex justify-between items-center mb-4 mt-4">
                  <div class="text-sm text-gray-500 bg-gray-50 dark:bg-gray-800 px-3 py-1 rounded">路径: resources/</div>
                  <el-button type="primary" icon="Plus" size="small" @click="openResourceDialog">创建资源</el-button>
                </div>
                <div class="text-xs text-gray-500 mb-3">
                  适合补充背景资料或术语表，在正文中引用 <span class="font-mono">resources/文件名</span> 提示模型按需查阅。
                </div>
                <el-table :data="resourceRows" style="width: 100%">
                  <el-table-column prop="name" label="文件名">
                    <template #default="scope">
                      <div class="flex items-center gap-2">
                        <el-icon><Document /></el-icon>
                        <span>{{ scope.row.name }}</span>
                      </div>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="180">
                    <template #default="scope">
                      <el-button type="primary" link icon="Edit" @click="openResourceEditor(scope.row.name)">编辑</el-button>
                      <el-button type="primary" link @click="insertFileSnippet('resource', scope.row.name)">引用</el-button>
                    </template>
                  </el-table-column>
                </el-table>
                <el-empty v-if="resourceRows.length === 0" description="暂无资源" />
              </el-tab-pane>

              <el-tab-pane label="References" name="references">
                <div class="flex justify-between items-center mb-4 mt-4">
                  <div class="text-sm text-gray-500 bg-gray-50 dark:bg-gray-800 px-3 py-1 rounded">路径: references/</div>
                  <el-button type="primary" icon="Plus" size="small" @click="openReferenceDialog">创建参考</el-button>
                </div>
                <div class="text-xs text-gray-500 mb-3">
                  适合放规范、规则或权威资料，在正文中引用 <span class="font-mono">references/文件名</span> 指定遵循来源。
                </div>
                <el-table :data="referenceRows" style="width: 100%">
                  <el-table-column prop="name" label="文件名">
                    <template #default="scope">
                      <div class="flex items-center gap-2">
                        <el-icon><Document /></el-icon>
                        <span>{{ scope.row.name }}</span>
                      </div>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="180">
                    <template #default="scope">
                      <el-button type="primary" link icon="Edit" @click="openReferenceEditor(scope.row.name)">编辑</el-button>
                      <el-button type="primary" link @click="insertFileSnippet('reference', scope.row.name)">引用</el-button>
                    </template>
                  </el-table-column>
                </el-table>
                <el-empty v-if="referenceRows.length === 0" description="暂无参考" />
              </el-tab-pane>

              <el-tab-pane label="Templates" name="templates">
                <div class="flex justify-between items-center mb-4 mt-4">
                  <div class="text-sm text-gray-500 bg-gray-50 dark:bg-gray-800 px-3 py-1 rounded">路径: templates/</div>
                  <el-button type="primary" icon="Plus" size="small" @click="openTemplateDialog">创建模板</el-button>
                </div>
                <div class="text-xs text-gray-500 mb-3">
                  适合放输出结构或代码骨架，在正文中引用 <span class="font-mono">templates/文件名</span> 作为格式约束。
                </div>
                <el-table :data="templateRows" style="width: 100%">
                  <el-table-column prop="name" label="文件名">
                    <template #default="scope">
                      <div class="flex items-center gap-2">
                        <el-icon><Document /></el-icon>
                        <span>{{ scope.row.name }}</span>
                      </div>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="180">
                    <template #default="scope">
                      <el-button type="primary" link icon="Edit" @click="openTemplateEditor(scope.row.name)">编辑</el-button>
                      <el-button type="primary" link @click="insertFileSnippet('template', scope.row.name)">引用</el-button>
                    </template>
                  </el-table-column>
                </el-table>
                <el-empty v-if="templateRows.length === 0" description="暂无模板" />
              </el-tab-pane>
            </el-tabs>
          </template>
        </el-card>
      </el-col>
    </el-row>

    <el-dialog v-model="createDialogVisible" title="新增 Skill" width="420px">
      <el-form :model="newSkill" label-width="100px">
        <el-form-item label="Skill 名称">
          <el-input v-model="newSkill.name" placeholder="例如: code-comment-expert" />
          <div class="text-xs text-gray-400 mt-1">仅小写字母/数字/连字符（kebab-case）。</div>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="newSkill.description" placeholder="例如: 代码注释与可读性优化，适合审查/重构" />
          <div class="text-xs text-gray-400 mt-1">写清触发条件与关键词，越具体越好。</div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createSkill">创建</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="scriptDialogVisible" title="创建脚本" width="420px">
      <el-form :model="newScript" label-width="100px">
        <el-form-item label="脚本类型">
          <el-select v-model="newScript.type" placeholder="选择类型">
            <el-option label="Python (.py)" value="py" />
            <el-option label="JavaScript (.js)" value="js" />
            <el-option label="Shell (.sh)" value="sh" />
          </el-select>
        </el-form-item>
        <el-form-item label="文件名">
          <el-input v-model="newScript.name" placeholder="例如: lint" />
          <div class="text-xs text-gray-400 mt-1">无需扩展名，会按类型自动补全。</div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="scriptDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createScript">创建</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="resourceDialogVisible" title="创建资源" width="420px">
      <el-form :model="newResource" label-width="100px">
        <el-form-item label="文件名">
          <el-input v-model="newResource.name" placeholder="例如: glossary" />
          <div class="text-xs text-gray-400 mt-1">自动补全 .md。</div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="resourceDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createResource">创建</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="referenceDialogVisible" title="创建参考" width="420px">
      <el-form :model="newReference" label-width="100px">
        <el-form-item label="文件名">
          <el-input v-model="newReference.name" placeholder="例如: style-guide" />
          <div class="text-xs text-gray-400 mt-1">自动补全 .md。</div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="referenceDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createReference">创建</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="templateDialogVisible" title="创建模板" width="420px">
      <el-form :model="newTemplate" label-width="100px">
        <el-form-item label="文件名">
          <el-input v-model="newTemplate.name" placeholder="例如: output-structure" />
          <div class="text-xs text-gray-400 mt-1">自动补全 .md。</div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="templateDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createTemplate">创建</el-button>
      </template>
    </el-dialog>

    <el-drawer v-model="editorVisible" size="70%" destroy-on-close :with-header="false">
      <div class="h-full flex flex-col p-4">
        <div class="flex justify-between items-center mb-4">
          <div class="text-lg font-bold flex items-center gap-2">
            <el-icon><Edit /></el-icon>
            {{ editorTitle }}
          </div>
          <div class="flex gap-2">
            <el-button @click="editorVisible = false">取消</el-button>
            <el-button type="primary" icon="Check" @click="saveEditor">保存内容</el-button>
          </div>
        </div>
        <div class="flex-1 overflow-hidden border border-gray-200 dark:border-gray-700 rounded-md shadow-inner">
          <v-ace-editor
            v-model:value="editorContent"
            :lang="editorLang"
            theme="github_dark"
            class="w-full h-full"
            :options="{ showPrintMargin: false, fontSize: 14 }"
          />
        </div>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
  import { computed, onMounted, reactive, ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { QuestionFilled, Document, Plus, Search, Check, Edit } from '@element-plus/icons-vue'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import {
    getSkillTools,
    getSkillList,
    getSkillDetail,
    saveSkill,
    createSkillScript,
    getSkillScript,
    saveSkillScript,
    createSkillResource,
    getSkillResource,
    saveSkillResource,
    createSkillReference,
    getSkillReference,
    saveSkillReference,
    createSkillTemplate,
    getSkillTemplate,
    saveSkillTemplate,
    getGlobalConstraint,
    saveGlobalConstraint
  } from '@/api/skills'
  import { VAceEditor } from 'vue3-ace-editor'
  import 'ace-builds/src-noconflict/mode-javascript'
  import 'ace-builds/src-noconflict/mode-python'
  import 'ace-builds/src-noconflict/mode-sh'
  import 'ace-builds/src-noconflict/mode-markdown'
  import 'ace-builds/src-noconflict/theme-github_dark'

  defineOptions({
    name: 'Skills'
  })

  const tools = ref([
    { key: 'copilot', label: 'Copilot' },
    { key: 'claude', label: 'Claude' },
    { key: 'cursor', label: 'Cursor' },
    { key: 'trae', label: 'Trae' },
    { key: 'codex', label: 'Codex' }
  ])
  const activeTool = ref('claude')
  const skills = ref([])
  const activeSkill = ref('')
  const skillFilter = ref('')
  const activeTab = ref('config')
  const globalConstraintExists = ref(false)

  const toolDirMap = {
    copilot: '.aone_copilot',
    claude: '.claude',
    cursor: '.cursor',
    trae: '.trae',
    codex: '.codex'
  }

  const globalConstraintPath = computed(() => {
    if (!activeTool.value) return 'skills/README.md'
    const toolDir = toolDirMap[activeTool.value] || `.${activeTool.value}`
    return `${toolDir}/skills/README.md`
  })

  const form = reactive({
    name: '',
    description: '',
    allowedTools: '',
    context: '',
    agent: '',
    markdown: ''
  })

  const markdownPlaceholder =
    '建议结构：# Skill Title -> ## Instructions -> ## Examples -> ## Guidelines。\n' +
    '在正文中引用 templates/...、references/...、scripts/... 可按需加载。\n\n' +
    '示例：\n# Code Comment Expert\n## Instructions\n- 说明目标、输入、输出与步骤。\n\n## Examples\n- 输入: ...\n- 输出: ...\n\n## Guidelines\n- 约束、格式与质量标准。\n'

  const quickBlocks = [
    { label: '标题', content: '\n# Skill Title\n' },
    { label: '指令', content: '\n## Instructions\n- 描述该技能要做什么、如何做。\n' },
    { label: '示例', content: '\n## Examples\n- 输入: ...\n- 输出: ...\n' },
    { label: '指南', content: '\n## Guidelines\n- 约束、质量标准与注意事项。\n' },
    { label: '输出格式', content: '\n## Output Format\n1. ...\n2. ...\n' },
    { label: '引用模板', content: '\n需要结构时参考 templates/your-template.md。\n' },
    { label: '引用参考', content: '\n如需规范/术语，参考 references/your-reference.md。\n' },
    { label: '调用脚本', content: '\n如需自动化，执行 scripts/your-script.py "{输入}"。\n' }
  ]

  const scripts = ref([])
  const resources = ref([])
  const references = ref([])
  const templates = ref([])

  const scriptRows = computed(() => skillsFilesToRows(scripts.value))
  const resourceRows = computed(() => skillsFilesToRows(resources.value))
  const referenceRows = computed(() => skillsFilesToRows(references.value))
  const templateRows = computed(() => skillsFilesToRows(templates.value))

  const createDialogVisible = ref(false)
  const scriptDialogVisible = ref(false)
  const resourceDialogVisible = ref(false)
  const referenceDialogVisible = ref(false)
  const templateDialogVisible = ref(false)

  const newSkill = reactive({
    name: '',
    description: ''
  })

  const newScript = reactive({
    name: '',
    type: 'py'
  })

  const newResource = reactive({
    name: ''
  })

  const newReference = reactive({
    name: ''
  })

  const newTemplate = reactive({
    name: ''
  })

  const editorVisible = ref(false)
  const editorContent = ref('')
  const editorFileName = ref('')
  const editorType = ref('script')
  const editorLang = ref('text')

  const editorTitle = computed(() => {
    if (!editorFileName.value) {
      return editorType.value === 'constraint' ? '全局约束' : '文件编辑'
    }
    if (editorType.value === 'script') return `脚本：${editorFileName.value}`
    if (editorType.value === 'resource') return `资源：${editorFileName.value}`
    if (editorType.value === 'reference') return `参考：${editorFileName.value}`
    if (editorType.value === 'template') return `模板：${editorFileName.value}`
    if (editorType.value === 'constraint') return `全局约束：${editorFileName.value}`
    return `文件编辑：${editorFileName.value}`
  })

  const filteredSkills = computed(() => {
    if (!skillFilter.value) return skills.value
    return skills.value.filter((item) => item.toLowerCase().includes(skillFilter.value.toLowerCase()))
  })

  onMounted(async () => {
    await loadTools()
    await loadSkills()
  })

  async function loadTools() {
    try {
      const res = await getSkillTools()
      if (res.code === 0 && res.data?.tools?.length) {
        tools.value = res.data.tools
        if (!tools.value.find((item) => item.key === activeTool.value)) {
          activeTool.value = tools.value[0]?.key || 'claude'
        }
      }
    } catch (e) {
      ElMessage.warning('获取工具列表失败，使用默认列表')
    }
  }

  async function loadSkills() {
    if (!activeTool.value) return
    try {
      const res = await getSkillList({ tool: activeTool.value })
      if (res.code === 0) {
        skills.value = res.data?.skills || []
      }
    } catch (e) {
      ElMessage.error('获取技能列表失败')
    }
  }

  async function loadSkillDetail(skillName) {
    if (!activeTool.value || !skillName) return
    try {
      const res = await getSkillDetail({ tool: activeTool.value, skill: skillName })
      if (res.code === 0) {
        const detail = res.data?.detail
        activeSkill.value = detail?.skill || skillName
        form.name = detail?.meta?.name || skillName
        form.description = detail?.meta?.description || ''
        form.allowedTools = detail?.meta?.allowedTools || ''
        form.context = detail?.meta?.context || ''
        form.agent = detail?.meta?.agent || ''
        form.markdown = detail?.markdown || ''
        scripts.value = detail?.scripts || []
        resources.value = detail?.resources || []
        references.value = detail?.references || []
        templates.value = detail?.templates || []
      }
    } catch (e) {
      ElMessage.error('获取技能详情失败')
    }
  }

  async function openGlobalConstraint() {
    if (!activeTool.value) {
      ElMessage.warning('请先选择工具')
      return
    }
    try {
      const res = await getGlobalConstraint({ tool: activeTool.value })
      if (res.code === 0) {
        globalConstraintExists.value = !!res.data?.exists
        if (!globalConstraintExists.value) {
          ElMessage.info('未检测到 README.md，保存后将创建该文件')
        }
        openEditor('constraint', 'README.md', res.data?.content || '')
      }
    } catch (e) {
      ElMessage.error('读取全局约束失败')
    }
  }

  function resetDetail() {
    activeSkill.value = ''
    form.name = ''
    form.description = ''
    form.allowedTools = ''
    form.context = ''
    form.agent = ''
    form.markdown = ''
    scripts.value = []
    resources.value = []
    references.value = []
    templates.value = []
    activeTab.value = 'config'
  }

  function handleToolSelect(key) {
    activeTool.value = key
    resetDetail()
    globalConstraintExists.value = false
    loadSkills()
  }

  function handleSkillSelect(skillName) {
    loadSkillDetail(skillName)
  }

  function openCreateDialog() {
    newSkill.name = ''
    newSkill.description = ''
    createDialogVisible.value = true
  }

  async function createSkill() {
    if (!newSkill.name.trim()) {
      ElMessage.warning('请输入 Skill 名称')
      return
    }
    const payload = {
      tool: activeTool.value,
      skill: newSkill.name.trim(),
      meta: {
        name: newSkill.name.trim(),
        description: newSkill.description.trim() || '请补充技能描述',
        allowedTools: 'Bash(gh *)',
        context: 'fork',
        agent: 'Explore'
      },
      markdown: defaultSkillTemplate()
    }
    try {
      const res = await saveSkill(payload)
      if (res.code === 0) {
        ElMessage.success('创建成功')
        createDialogVisible.value = false
        await loadSkills()
        await loadSkillDetail(payload.skill)
      }
    } catch (e) {
      ElMessage.error('创建失败')
    }
  }

  async function saveCurrentSkill() {
    if (!activeSkill.value) return
    if (!form.name.trim()) {
      ElMessage.warning('Name 不能为空')
      return
    }
    const payload = {
      tool: activeTool.value,
      skill: activeSkill.value,
      meta: {
        name: form.name.trim(),
        description: form.description.trim(),
        allowedTools: form.allowedTools.trim(),
        context: form.context.trim(),
        agent: form.agent.trim()
      },
      markdown: form.markdown
    }

    let syncTools = []
    try {
      await ElMessageBox.confirm('是否同步到其他 AI 客户端工具？', '同步提示', {
        confirmButtonText: '同步',
        cancelButtonText: '仅当前',
        type: 'warning'
      })
      syncTools = tools.value
        .map((item) => item.key)
        .filter((key) => key && key !== activeTool.value)
    } catch (e) {
      syncTools = []
    }

    if (syncTools.length) {
      payload.syncTools = syncTools
    }

    try {
      const res = await saveSkill(payload)
      if (res.code === 0) {
        ElMessage.success('保存成功')
      }
    } catch (e) {
      ElMessage.error('保存失败')
    }
  }

  function appendMarkdown(content) {
    form.markdown = `${form.markdown || ''}${content}`
  }

  function insertFileSnippet(kind, fileName) {
    if (!fileName) return
    let snippet = ''
    switch (kind) {
      case 'script':
        snippet = `如需自动化处理，可执行 scripts/${fileName} "{输入}"。`
        break
      case 'resource':
        snippet = `背景资料见 resources/${fileName}。`
        break
      case 'reference':
        snippet = `请遵循 references/${fileName} 的规范。`
        break
      case 'template':
        snippet = `输出结构参考 templates/${fileName}。`
        break
      default:
        snippet = ''
    }
    if (!snippet) return
    appendMarkdown(`\n${snippet}\n`)
    ElMessage.success('已插入到 SKILL.md')
    activeTab.value = 'config'
  }

  function insertFullTemplate() {
    if (!form.markdown.trim()) {
      form.markdown = defaultSkillTemplate()
      return
    }
    form.markdown = `${form.markdown}\n${defaultSkillTemplate()}`
  }

  function openScriptDialog() {
    if (!activeSkill.value) {
      ElMessage.warning('请先选择技能')
      return
    }
    newScript.name = ''
    newScript.type = 'py'
    scriptDialogVisible.value = true
  }

  async function createScript() {
    if (!newScript.name.trim()) {
      ElMessage.warning('请输入脚本文件名')
      return
    }
    try {
      const res = await createSkillScript({
        tool: activeTool.value,
        skill: activeSkill.value,
        fileName: newScript.name.trim(),
        scriptType: newScript.type
      })
      if (res.code === 0) {
        scriptDialogVisible.value = false
        await loadSkillDetail(activeSkill.value)
        openEditor('script', res.data.fileName, res.data.content)
      }
    } catch (e) {
      ElMessage.error('创建脚本失败')
    }
  }

  async function openScriptEditor(fileName) {
    if (!fileName) return
    try {
      const res = await getSkillScript({
        tool: activeTool.value,
        skill: activeSkill.value,
        fileName
      })
      if (res.code === 0) {
        openEditor('script', fileName, res.data.content)
      }
    } catch (e) {
      ElMessage.error('读取脚本失败')
    }
  }

  function openResourceDialog() {
    if (!activeSkill.value) {
      ElMessage.warning('请先选择技能')
      return
    }
    newResource.name = ''
    resourceDialogVisible.value = true
  }

  async function createResource() {
    if (!newResource.name.trim()) {
      ElMessage.warning('请输入资源文件名')
      return
    }
    try {
      const res = await createSkillResource({
        tool: activeTool.value,
        skill: activeSkill.value,
        fileName: newResource.name.trim()
      })
      if (res.code === 0) {
        resourceDialogVisible.value = false
        await loadSkillDetail(activeSkill.value)
        openEditor('resource', res.data.fileName, res.data.content)
      }
    } catch (e) {
      ElMessage.error('创建资源失败')
    }
  }

  async function openResourceEditor(fileName) {
    if (!fileName) return
    try {
      const res = await getSkillResource({
        tool: activeTool.value,
        skill: activeSkill.value,
        fileName
      })
      if (res.code === 0) {
        openEditor('resource', fileName, res.data.content)
      }
    } catch (e) {
      ElMessage.error('读取资源失败')
    }
  }

  function openReferenceDialog() {
    if (!activeSkill.value) {
      ElMessage.warning('请先选择技能')
      return
    }
    newReference.name = ''
    referenceDialogVisible.value = true
  }

  async function createReference() {
    if (!newReference.name.trim()) {
      ElMessage.warning('请输入参考文件名')
      return
    }
    try {
      const res = await createSkillReference({
        tool: activeTool.value,
        skill: activeSkill.value,
        fileName: newReference.name.trim()
      })
      if (res.code === 0) {
        referenceDialogVisible.value = false
        await loadSkillDetail(activeSkill.value)
        openEditor('reference', res.data.fileName, res.data.content)
      }
    } catch (e) {
      ElMessage.error('创建参考失败')
    }
  }

  async function openReferenceEditor(fileName) {
    if (!fileName) return
    try {
      const res = await getSkillReference({
        tool: activeTool.value,
        skill: activeSkill.value,
        fileName
      })
      if (res.code === 0) {
        openEditor('reference', fileName, res.data.content)
      }
    } catch (e) {
      ElMessage.error('读取参考失败')
    }
  }

  function openTemplateDialog() {
    if (!activeSkill.value) {
      ElMessage.warning('请先选择技能')
      return
    }
    newTemplate.name = ''
    templateDialogVisible.value = true
  }

  async function createTemplate() {
    if (!newTemplate.name.trim()) {
      ElMessage.warning('请输入模板文件名')
      return
    }
    try {
      const res = await createSkillTemplate({
        tool: activeTool.value,
        skill: activeSkill.value,
        fileName: newTemplate.name.trim()
      })
      if (res.code === 0) {
        templateDialogVisible.value = false
        await loadSkillDetail(activeSkill.value)
        openEditor('template', res.data.fileName, res.data.content)
      }
    } catch (e) {
      ElMessage.error('创建模板失败')
    }
  }

  async function openTemplateEditor(fileName) {
    if (!fileName) return
    try {
      const res = await getSkillTemplate({
        tool: activeTool.value,
        skill: activeSkill.value,
        fileName
      })
      if (res.code === 0) {
        openEditor('template', fileName, res.data.content)
      }
    } catch (e) {
      ElMessage.error('读取模板失败')
    }
  }

  function openEditor(type, fileName, content) {
    editorType.value = type
    editorFileName.value = fileName
    editorContent.value = content || ''
    editorLang.value = detectLang(fileName)
    editorVisible.value = true
  }

  async function saveEditor() {
    if (!editorFileName.value) return
    try {
      if (editorType.value === 'script') {
        const res = await saveSkillScript({
          tool: activeTool.value,
          skill: activeSkill.value,
          fileName: editorFileName.value,
          content: editorContent.value
        })
        if (res.code === 0) {
          ElMessage.success('保存成功')
        }
      } else if (editorType.value === 'resource') {
        const res = await saveSkillResource({
          tool: activeTool.value,
          skill: activeSkill.value,
          fileName: editorFileName.value,
          content: editorContent.value
        })
        if (res.code === 0) {
          ElMessage.success('保存成功')
        }
      } else if (editorType.value === 'reference') {
        const res = await saveSkillReference({
          tool: activeTool.value,
          skill: activeSkill.value,
          fileName: editorFileName.value,
          content: editorContent.value
        })
        if (res.code === 0) {
          ElMessage.success('保存成功')
        }
      } else if (editorType.value === 'template') {
        const res = await saveSkillTemplate({
          tool: activeTool.value,
          skill: activeSkill.value,
          fileName: editorFileName.value,
          content: editorContent.value
        })
        if (res.code === 0) {
          ElMessage.success('保存成功')
        }
      } else if (editorType.value === 'constraint') {
        let syncTools = []
        if (tools.value.length > 1) {
          try {
            await ElMessageBox.confirm('是否同步到其他 AI 客户端工具？', '同步提示', {
              confirmButtonText: '同步',
              cancelButtonText: '仅当前',
              type: 'warning'
            })
            syncTools = tools.value
              .map((item) => item.key)
              .filter((key) => key && key !== activeTool.value)
          } catch (e) {
            syncTools = []
          }
        }

        const res = await saveGlobalConstraint({
          tool: activeTool.value,
          content: editorContent.value,
          syncTools
        })
        if (res.code !== 0) {
          ElMessage.error('保存失败')
          return
        }
        globalConstraintExists.value = true
        ElMessage.success(syncTools.length ? '保存并同步成功' : '保存成功')
      }
    } catch (e) {
      ElMessage.error('保存失败')
    }
  }

  function detectLang(fileName) {
    if (!fileName) return 'text'
    const lower = fileName.toLowerCase()
    if (lower.endsWith('.py')) return 'python'
    if (lower.endsWith('.js')) return 'javascript'
    if (lower.endsWith('.sh')) return 'sh'
    if (lower.endsWith('.md')) return 'markdown'
    return 'text'
  }

  function defaultSkillTemplate() {
    return (
      '# Skill Title\n' +
      '## Instructions\n- 说明目标、输入、输出与步骤。\n\n' +
      '## Examples\n- 输入: ...\n- 输出: ...\n\n' +
      '## Guidelines\n- 约束、格式与质量标准。\n\n' +
      '## Output Format\n1. ...\n2. ...\n'
    )
  }

  function skillsFilesToRows(list) {
    return (list || []).map((name) => ({ name }))
  }
</script>
