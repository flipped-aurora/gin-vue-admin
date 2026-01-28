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
                <el-form :model="form" label-width="160px" class="mt-4">
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Name
                        <el-tooltip content="技能的名称，例如: pr-summary" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <el-input v-model="form.name" placeholder="例如: pr-summary" />
                  </el-form-item>
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Description
                        <el-tooltip content="技能的简要描述，例如: Summarize changes in a pull request" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <el-input
                      v-model="form.description"
                      placeholder="例如: Summarize changes in a pull request"
                    />
                  </el-form-item>
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Allowed Tools
                        <el-tooltip content="该技能允许使用的工具，例如: Bash(gh *)" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <el-input v-model="form.allowedTools" placeholder="可选，例如: Bash(gh *)" />
                    <div class="text-xs text-gray-400 mt-1">可选字段，留空后保存会移除</div>
                  </el-form-item>
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Context
                        <el-tooltip content="技能执行的上下文，例如: fork" placement="top">
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
                        <el-tooltip content="指定执行该技能的 Agent，例如: Explore" placement="top">
                          <el-icon class="ml-1 cursor-pointer"><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </div>
                    </template>
                    <el-input v-model="form.agent" placeholder="可选，例如: Explore" />
                    <div class="text-xs text-gray-400 mt-1">可选字段，留空后保存会移除</div>
                  </el-form-item>
                  <el-form-item>
                    <template #label>
                      <div class="flex items-center">
                        Markdown 内容
                        <el-tooltip content="SKILL.md 的具体内容，定义技能的详细逻辑" placement="top">
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
                    <div class="text-xs text-gray-400 mt-1">这里是 SKILL.md 的正文内容，可自由编辑。</div>
                  </el-form-item>
                </el-form>
              </el-tab-pane>

              <el-tab-pane label="脚本" name="scripts" class="mt-4">
                <div class="flex justify-between items-center mb-4">
                  <div class="text-sm text-gray-500 bg-gray-50 dark:bg-gray-800 px-3 py-1 rounded">路径: scripts/</div>
                  <el-button type="primary" icon="Plus" size="small" @click="openScriptDialog">创建脚本</el-button>
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
                  <el-table-column label="操作" width="120">
                    <template #default="scope">
                      <el-button type="primary" link icon="Edit" @click="openScriptEditor(scope.row.name)">编辑</el-button>
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
                <el-table :data="resourceRows" style="width: 100%">
                  <el-table-column prop="name" label="文件名">
                    <template #default="scope">
                      <div class="flex items-center gap-2">
                        <el-icon><Document /></el-icon>
                        <span>{{ scope.row.name }}</span>
                      </div>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="120">
                    <template #default="scope">
                      <el-button type="primary" link icon="Edit" @click="openResourceEditor(scope.row.name)">编辑</el-button>
                    </template>
                  </el-table-column>
                </el-table>
                <el-empty v-if="resourceRows.length === 0" description="暂无资源" />
              </el-tab-pane>
            </el-tabs>
          </template>
        </el-card>
      </el-col>
    </el-row>

    <el-dialog v-model="createDialogVisible" title="新增 Skill" width="420px">
      <el-form :model="newSkill" label-width="100px">
        <el-form-item label="Skill 名称">
          <el-input v-model="newSkill.name" placeholder="例如: pr-summary" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="newSkill.description" placeholder="可选" />
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
          <el-input v-model="newScript.name" placeholder="例如: run" />
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
          <el-input v-model="newResource.name" placeholder="例如: usage" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="resourceDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createResource">创建</el-button>
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
    saveSkillResource
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

  const form = reactive({
    name: '',
    description: '',
    allowedTools: '',
    context: '',
    agent: '',
    markdown: ''
  })

  const markdownPlaceholder =
    '建议包含：技能用途、输入、输出、步骤与示例。\n\n示例：\n## 技能用途\n请描述技能的目标与限制。\n\n## 输入\n- 输入1...\n\n## 输出\n- 输出1...\n'

  const quickBlocks = [
    { label: '用途', content: '\n## 技能用途\n请描述技能目标与适用场景。\n' },
    { label: '输入', content: '\n## 输入\n- 输入字段与格式说明。\n' },
    { label: '输出', content: '\n## 输出\n- 输出字段与格式说明。\n' },
    { label: '步骤', content: '\n## 关键步骤\n1. 第一步\n2. 第二步\n' },
    { label: '示例', content: '\n## 示例\n在此补充示例。\n' },
    { label: '注意事项', content: '\n## 注意事项\n- 需要注意的限制或风险。\n' }
  ]

  const scripts = ref([])
  const resources = ref([])

  const scriptRows = computed(() => skillsFilesToRows(scripts.value))
  const resourceRows = computed(() => skillsFilesToRows(resources.value))

  const createDialogVisible = ref(false)
  const scriptDialogVisible = ref(false)
  const resourceDialogVisible = ref(false)

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

  const editorVisible = ref(false)
  const editorContent = ref('')
  const editorFileName = ref('')
  const editorType = ref('script')
  const editorLang = ref('text')

  const editorTitle = computed(() => {
    if (!editorFileName.value) return '文件编辑'
    return `${editorType.value === 'script' ? '脚本' : '资源'}：${editorFileName.value}`
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
      }
    } catch (e) {
      ElMessage.error('获取技能详情失败')
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
    activeTab.value = 'config'
  }

  function handleToolSelect(key) {
    activeTool.value = key
    resetDetail()
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
      } else {
        const res = await saveSkillResource({
          tool: activeTool.value,
          skill: activeSkill.value,
          fileName: editorFileName.value,
          content: editorContent.value
        })
        if (res.code === 0) {
          ElMessage.success('保存成功')
        }
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
      '## 技能用途\n请在这里描述技能的目标、适用场景与限制条件。\n\n' +
      '## 输入\n- 请补充输入格式与示例。\n\n' +
      '## 输出\n- 请补充输出格式与示例。\n\n' +
      '## 关键步骤\n1. 第一步\n2. 第二步\n\n' +
      '## 示例\n在此补充一到两个典型示例。\n'
    )
  }

  function skillsFilesToRows(list) {
    return (list || []).map((name) => ({ name }))
  }
</script>

