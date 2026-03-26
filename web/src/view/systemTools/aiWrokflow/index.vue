<template>
  <div class="gva-table-box ai-workflow-page space-y-4">
    <el-card shadow="never">
      <div class="flex flex-wrap items-start justify-between gap-4">
        <div class="max-w-3xl">
          <h1 class="mt-2 text-2xl font-semibold text-slate-900">
            AI 需求分析与 Prompt 工作流
          </h1>
          <p class="mt-3 text-sm leading-6 text-slate-500">
            会话会自动保存到后端。刷新页面后可以继续查看历史需求、按节点回看结果，并支持回滚到任意
            Assistant 节点后重新续聊。
          </p>
        </div>
        <div class="flex flex-wrap">
          <el-button :icon="MagicStick" @click="fillExample"
            >填入示例</el-button
          >
        </div>
      </div>
    </el-card>

    <div class="grid grid-cols-1 gap-4 xl:grid-cols-[380px_minmax(0,1fr)]">
      <el-card shadow="never" class="self-start">
        <template #header>
          <div class="flex items-center justify-between gap-3">
            <div>
              <p class="text-base font-semibold text-slate-800">输入与历史</p>
              <p class="mt-1 text-xs text-slate-500">
                {{
                  currentSession.id
                    ? `会话 #${currentSession.id}`
                    : '未保存新会话'
                }}
              </p>
            </div>
            <el-tag
              effect="plain"
              :type="activeTab === 'analysis' ? 'primary' : 'success'"
            >
              {{ activeTab === 'analysis' ? '当前：需求分析' : '当前：工作流' }}
            </el-tag>
          </div>
        </template>

        <div class="space-y-4">
          <div
            class="space-y-3 rounded border border-slate-200 bg-white p-3 shadow-sm"
          >
            <div
              class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between"
            >
              <div>
                <p class="text-sm font-semibold text-slate-800">
                  先选择当前工作模式
                </p>
                <p class="mt-1 text-xs text-slate-500">
                  很多用户会忽略这里，现在可以直接点下面的大卡片切换。
                </p>
              </div>
              <span
                class="w-fit shrink-0 whitespace-nowrap rounded-full bg-slate-100 px-3 py-1 text-[11px] font-semibold text-slate-500"
                >可随时切换</span
              >
            </div>
            <div class="grid grid-cols-1 gap-3 md:grid-cols-2 xl:grid-cols-1">
              <button
                type="button"
                class="rounded border px-4 py-4 text-left transition"
                :class="
                  activeTab === 'analysis'
                    ? 'border-sky-400 bg-sky-50 shadow-sm shadow-sky-100'
                    : 'border-slate-200 bg-slate-50 hover:border-slate-300 hover:bg-white'
                "
                @click="switchTab('analysis')"
              >
                <div class="flex items-center justify-between gap-3">
                  <div>
                    <p
                      class="text-base font-semibold"
                      :class="
                        activeTab === 'analysis'
                          ? 'text-sky-700'
                          : 'text-slate-800'
                      "
                    >
                      需求分析
                    </p>
                    <p
                      class="mt-2 text-xs leading-5"
                      :class="
                        activeTab === 'analysis'
                          ? 'text-sky-700/80'
                          : 'text-slate-500'
                      "
                    >
                      先把你的需求拆成模块、字段、待确认问题，再决定怎么落地。
                    </p>
                  </div>
                  <span
                    class="shrink-0 whitespace-nowrap rounded-full px-3 py-1 text-[11px] font-semibold leading-none"
                    :class="
                      activeTab === 'analysis'
                        ? 'bg-sky-600 text-white'
                        : 'bg-white text-slate-500 border border-slate-200'
                    "
                  >
                    {{ activeTab === 'analysis' ? '当前模式' : '切换到这里' }}
                  </span>
                </div>
              </button>

              <button
                type="button"
                class="rounded border px-4 py-4 text-left transition"
                :class="
                  activeTab === 'workflow'
                    ? 'border-emerald-400 bg-emerald-50 shadow-sm shadow-emerald-100'
                    : 'border-slate-200 bg-slate-50 hover:border-slate-300 hover:bg-white'
                "
                @click="switchTab('workflow')"
              >
                <div class="flex items-center justify-between gap-3">
                  <div>
                    <p
                      class="text-base font-semibold"
                      :class="
                        activeTab === 'workflow'
                          ? 'text-emerald-700'
                          : 'text-slate-800'
                      "
                    >
                      工作流
                    </p>
                    <p
                      class="mt-2 text-xs leading-5"
                      :class="
                        activeTab === 'workflow'
                          ? 'text-emerald-700/80'
                          : 'text-slate-500'
                      "
                    >
                      基于需求或分析结果，直接生成分步骤 Prompt 和执行路线。
                    </p>
                  </div>
                  <span
                    class="shrink-0 whitespace-nowrap rounded-full px-3 py-1 text-[11px] font-semibold leading-none"
                    :class="
                      activeTab === 'workflow'
                        ? 'bg-emerald-600 text-white'
                        : 'bg-white text-slate-500 border border-slate-200'
                    "
                  >
                    {{ activeTab === 'workflow' ? '当前模式' : '切换到这里' }}
                  </span>
                </div>
              </button>
            </div>
          </div>

          <div
            class="space-y-3 rounded border border-slate-200 bg-slate-50 p-3"
          >
            <div class="flex items-center">
              <el-input
                v-model="historyKeywords[activeTab]"
                class="flex-1 mr-2"
                clearable
                placeholder="搜索历史需求"
                @keyup.enter="loadSessionList(activeTab)"
              />
              <el-button
                plain
                class="!px-3"
                :icon="RefreshRight"
                :loading="historyLoading"
                title="刷新列表"
                @click="loadSessionList(activeTab)"
              />
              <el-button
                plain
                class="!px-3"
                :icon="Plus"
                title="新建会话"
                @click="startNewConversation(activeTab)"
              />
            </div>
            <div
              v-if="currentSessionList.length"
              class="max-h-[280px] space-y-2 overflow-auto pr-1"
            >
              <div
                v-for="item in currentSessionList"
                :key="item.ID || item.id"
                class="cursor-pointer rounded border p-3"
                :class="
                  isSessionActive(item)
                    ? 'border-sky-300 bg-sky-50'
                    : 'border-slate-200 bg-white'
                "
                @click="openSession(item)"
              >
                <div class="flex items-start justify-between gap-3">
                  <div class="min-w-0 flex-1">
                    <div class="flex flex-wrap items-center gap-2">
                      <p class="truncate text-sm font-semibold text-slate-800">
                        {{ item.title || '未命名需求' }}
                      </p>
                      <el-tag size="small" effect="plain">{{
                        item.tab === 'analysis' ? '分析' : '工作流'
                      }}</el-tag>
                    </div>
                    <p class="mt-2 text-xs leading-5 text-slate-500">
                      {{ item.summary || '暂无摘要' }}
                    </p>
                    <p class="mt-2 text-[11px] text-slate-400">
                      更新时间：{{
                        formatTime(item.UpdatedAt || item.updatedAt)
                      }}
                    </p>
                  </div>
                  <el-button
                    link
                    type="danger"
                    :icon="Delete"
                    @click.stop="removeSession(item)"
                  />
                </div>
              </div>
            </div>
            <el-empty v-else description="当前 tab 还没有保存的会话。" />
          </div>

          <el-form label-position="top">
            <template v-if="activeTab === 'analysis'">
              <el-form-item label="原始需求">
                <el-input
                  v-model="analysisForm.requirement"
                  type="textarea"
                  :rows="10"
                  maxlength="4000"
                  show-word-limit
                  placeholder="例如：我要做一个请假管理系统，支持员工提交请假申请、负责人审批、人事统计。"
                />
              </el-form-item>
              <el-form-item label="目标形态">
                <el-radio-group v-model="analysisForm.packageType">
                  <el-radio-button label="auto">自动判断</el-radio-button>
                  <el-radio-button label="package">package</el-radio-button>
                  <el-radio-button label="plugin">plugin</el-radio-button>
                </el-radio-group>
              </el-form-item>
              <el-form-item label="业务场景">
                <el-input
                  v-model="analysisForm.businessScene"
                  placeholder="例如：OA、ERP、内部后台"
                />
              </el-form-item>
              <el-form-item label="额外约束">
                <el-input
                  v-model="analysisForm.extraConstraints"
                  type="textarea"
                  :rows="4"
                  placeholder="例如：需要附件字段、状态字典、导出统计。"
                />
              </el-form-item>
              <el-form-item label="是否存在客户端页面">
                <el-switch
                  v-model="analysisForm.hasClientPage"
                  active-text="有"
                  inactive-text="无"
                />
              </el-form-item>
              <template v-if="analysisForm.hasClientPage">
                <el-form-item label="客户端页面说明">
                  <el-input
                    v-model="analysisForm.clientPageDescription"
                    type="textarea"
                    :rows="4"
                    placeholder="例如：需要一个员工端 H5 页面，用于提交请假申请、查看审批进度、上传附件。"
                  />
                </el-form-item>
                <el-form-item label="客户端页面额外约束">
                  <el-input
                    v-model="analysisForm.clientPageConstraints"
                    type="textarea"
                    :rows="4"
                    placeholder="例如：移动端优先，字段尽量少，审批状态要醒目，附件支持图片和 PDF。"
                  />
                </el-form-item>
              </template>
              <div class="flex flex-wrap">
                <el-button
                  type="primary"
                  :loading="analysisLoading"
                  :icon="MagicStick"
                  @click="runAnalysis"
                  >开始分析</el-button
                >
                <el-button
                  :disabled="!hasAnalysisResult"
                  @click="pushAnalysisToWorkflow"
                  >带入工作流</el-button
                >
                <el-button @click="startNewConversation('analysis')"
                  >清空</el-button
                >
              </div>
            </template>

            <template v-else>
              <el-form-item label="需求或分析结果">
                <el-input
                  v-model="workflowForm.source"
                  type="textarea"
                  :rows="10"
                  maxlength="6000"
                  show-word-limit
                  placeholder="可以直接粘贴原始需求，也可以把需求分析的结构化结果带进来。"
                />
              </el-form-item>
              <el-form-item label="工作流目标">
                <el-select v-model="workflowForm.flowType" class="w-full">
                  <el-option label="GVA 代码生成" value="gva_codegen" />
                  <el-option label="GVA 功能完善" value="gva_polish" />
                  <el-option label="MCP 使用指导" value="mcp_assist" />
                </el-select>
              </el-form-item>
              <el-form-item label="额外约束">
                <el-input
                  v-model="workflowForm.extraConstraints"
                  type="textarea"
                  :rows="4"
                  placeholder="例如：每一步都给出可复制 Prompt，并说明预期输出。"
                />
              </el-form-item>
              <div class="flex flex-wrap">
                <el-button
                  type="primary"
                  :loading="workflowLoading"
                  :icon="MagicStick"
                  @click="runWorkflow"
                  >生成</el-button
                >
                <el-button
                  :disabled="!workflowResult.steps.length"
                  :icon="DocumentCopy"
                  @click="copyAllWorkflowPrompts"
                  >复制全部</el-button
                >
                <el-button @click="startNewConversation('workflow')"
                  >清空</el-button
                >
              </div>
            </template>

            <el-divider />

            <div class="space-y-3 rounded bg-slate-50 p-3">
              <div class="text-xs leading-6 text-slate-500">
                当前会话会持久化到后端。点击历史需求即可恢复内容，刷新页面后也会自动加载最近一次会话。
              </div>
              <div class="flex flex-wrap">
                <el-button plain @click="startNewConversation(activeTab)"
                  >新建会话</el-button
                >
                <el-button
                  plain
                  :disabled="!currentSession.conversationId"
                  @click="
                    copyText(
                      currentSession.conversationId,
                      '已复制 conversation_id'
                    )
                  "
                  >复制会话 ID</el-button
                >
              </div>
              <div
                class="rounded-xl border border-dashed border-slate-300 bg-white px-3 py-2 text-xs text-slate-500"
              >
                conversation_id:
                <span class="break-all font-mono text-slate-700">{{
                  currentSession.conversationId ||
                  '当前会话还没有 conversation_id'
                }}</span>
              </div>
              <div
                class="rounded-xl border border-dashed border-slate-300 bg-white px-3 py-2 text-xs text-slate-500"
              >
                当前节点:
                <span class="break-all font-mono text-slate-700">{{
                  currentSession.currentNodeId || '默认显示最新 Assistant 节点'
                }}</span>
              </div>
            </div>

            <el-divider />

            <el-collapse>
              <el-collapse-item title="额外透传参数" name="settings">
                <el-form-item label="额外透传参数 JSON">
                  <el-input
                    v-model="settings.extraPayload"
                    type="textarea"
                    :rows="4"
                    placeholder='例如：{"tenant":"default","provider":"dify"}'
                  />
                </el-form-item>
              </el-collapse-item>
            </el-collapse>
          </el-form>
        </div>
      </el-card>

      <el-card shadow="never" class="min-h-[760px]">
        <template #header>
          <div class="flex items-center justify-between gap-3">
            <div>
              <p class="text-base font-semibold text-slate-800">
                {{ activeTab === 'analysis' ? '需求分析结果' : '工作流结果' }}
              </p>
              <p class="mt-1 text-xs text-slate-500">
                {{
                  activeTab === 'analysis'
                    ? '支持从会话节点回看任意版本的分析结果。'
                    : '支持从会话节点回看任意版本的工作流结果。'
                }}
              </p>
            </div>
            <div class="flex gap-2">
              <el-button
                plain
                :loading="currentDumpLoading"
                :disabled="!canDumpCurrentMarkdown"
                @click="dumpCurrentMarkdown"
                >{{ activeTab === 'analysis' ? '落盘分析' : '落盘 Prompt' }}</el-button
              >
              <el-button
                v-if="selectedAssistantMessage"
                plain
                @click="rollbackToMessage(selectedAssistantMessage.id)"
                >回滚到当前节点</el-button
              >
              <el-button
                v-if="activeTab === 'analysis'"
                :disabled="!hasAnalysisResult"
                :icon="DocumentCopy"
                @click="copyAnalysisResult"
                >复制结构化结果</el-button
              >
              <el-button
                v-else
                :disabled="!workflowResult.rawText"
                :icon="RefreshRight"
                @click="copyWorkflowRaw"
                >复制原始结果</el-button
              >
            </div>
          </div>
        </template>

        <div class="space-y-4">
          <div
            v-if="currentLoading"
            class="rounded border border-sky-200 bg-sky-50 px-4 py-3 text-sm text-sky-700"
          >
            <div class="mb-2 flex items-center gap-2">
              <el-icon class="animate-spin"><RefreshRight /></el-icon>
              <span class="font-semibold">AI 思考中...</span>
            </div>
            <pre
              v-if="streamingPreviewText"
              ref="streamingPreviewRef"
              class="mt-2 max-h-[200px] overflow-auto whitespace-pre-wrap rounded bg-slate-950 p-3 text-xs leading-5 text-slate-100"
            >{{ streamingPreviewText }}</pre>
            <div v-else class="text-xs text-sky-500">等待模型响应...</div>
          </div>
          <div class="rounded border border-slate-200 bg-slate-50 p-4">
            <div class="mb-3 flex items-center justify-between">
              <h3 class="text-sm font-semibold text-slate-800">会话记录</h3>
              <el-tag effect="plain"
                >{{ currentSession.messages.length }} 条</el-tag
              >
            </div>
            <div
              v-if="currentSession.messages.length"
              ref="conversationListRef"
              class="max-h-[320px] space-y-3 overflow-auto pr-1"
            >
              <div
                v-for="item in currentSessionDisplayMessages"
                :key="item.id"
                class="rounded px-4 py-3"
                :class="
                  item.role === 'user'
                    ? 'bg-sky-600 text-white'
                    : item.isSelected
                    ? 'border border-sky-300 bg-sky-50 text-slate-700'
                    : 'border border-slate-200 bg-white text-slate-700'
                "
              >
                <div class="mb-2 flex items-start justify-between gap-3">
                  <div>
                    <div
                      class="text-xs font-semibold uppercase tracking-[0.2em]"
                      :class="
                        item.role === 'user' ? 'text-sky-100' : 'text-slate-400'
                      "
                    >
                      {{ item.role === 'user' ? 'User' : 'Assistant' }}
                    </div>
                    <div
                      v-if="item.createdAt"
                      class="mt-1 text-[11px]"
                      :class="
                        item.role === 'user'
                          ? 'text-sky-100/80'
                          : 'text-slate-400'
                      "
                    >
                      {{ formatTime(item.createdAt) }}
                    </div>
                  </div>
                  <div v-if="item.role === 'assistant'" class="flex flex-wrap">
                    <el-button
                      link
                      type="primary"
                      @click.stop="selectMessageNode(item.id)"
                      >查看节点</el-button
                    >
                    <el-button
                      link
                      type="danger"
                      @click.stop="rollbackToMessage(item.id)"
                      >回滚到这里</el-button
                    >
                  </div>
                </div>
                <div
                  v-if="item.role === 'user'"
                  class="whitespace-pre-wrap text-sm leading-6"
                >
                  {{ item.content }}
                </div>
                <div v-else class="space-y-3">
                  <template v-if="item.isStreaming">
                    <div
                      class="rounded-xl border border-sky-200 bg-sky-50 px-3 py-2 text-xs leading-5 text-sky-700"
                    >
                      正在流式返回，当前内容实时展开中...
                    </div>
                    <pre
                      class="overflow-auto whitespace-pre-wrap rounded bg-slate-950 p-4 text-xs leading-6 text-slate-100"
                      >{{ item.content || '...' }}</pre
                    >
                  </template>
                  <template v-else>
                  <div
                    class="rounded-xl border border-dashed border-slate-200 bg-white px-3 py-2 text-xs leading-5 text-slate-500"
                  >
                    {{
                      item.display.preview ||
                      'Assistant 返回内容已折叠，按需展开查看。'
                    }}
                  </div>
                  <div class="flex gap-2">
                    <el-tag
                      v-if="item.display.hasThink"
                      effect="plain"
                      type="warning"
                      class="cursor-pointer"
                      @click="openDrawer('Think', item.display.think)"
                    >Think</el-tag>
                    <el-tag
                      v-if="item.display.hasRawFile"
                      effect="plain"
                      class="cursor-pointer"
                      @click="openDrawer('原始文件', item.display.rawFile)"
                    >原始文件</el-tag>
                  </div>
                  </template>
                </div>
              </div>
            </div>
            <el-empty v-else description="当前会话还没有消息。" />
          </div>

          <div class="rounded border border-slate-200 bg-white p-4">
            <div class="mb-3 flex items-center justify-between">
              <h3 class="text-sm font-semibold text-slate-800">继续追问</h3>
              <span class="text-xs text-slate-400">{{
                currentSession.conversationId
                  ? '会自动续用当前 conversation_id'
                  : '若已回滚，将基于当前节点结果重新开启会话'
              }}</span>
            </div>
            <el-input
              v-model="followUpInput"
              type="textarea"
              :rows="4"
              maxlength="3000"
              show-word-limit
              :placeholder="
                activeTab === 'analysis'
                  ? '例如：补充一下，需要销假流程和抄送功能。'
                  : '例如：把步骤再细化，增加字段确认和字典设计。'
              "
            />
            <div class="mt-3 flex flex-wrap">
              <el-button
                type="primary"
                :loading="currentLoading"
                :disabled="!canSendFollowUp"
                @click="sendFollowUp"
                >继续追问</el-button
              >
              <el-button @click="followUpInput = ''">清空输入</el-button>
            </div>
          </div>

          <template v-if="activeTab === 'analysis'">
            <template v-if="hasAnalysisResult">
              <div class="rounded bg-slate-50 p-4">
                <p
                  class="text-xs font-semibold uppercase tracking-[0.2em] text-sky-600"
                >
                  Summary
                </p>
                <p
                  class="mt-2 whitespace-pre-wrap text-sm leading-6 text-slate-700"
                >
                  {{
                    analysisResult.summary || '未解析到摘要，将保留原始结果。'
                  }}
                </p>
                <div class="mt-3 flex flex-wrap gap-2">
                  <el-tag
                    v-if="analysisResult.recommendedPackageType"
                    effect="plain"
                    type="primary"
                    >推荐: {{ analysisResult.recommendedPackageType }}</el-tag
                  >
                  <el-tag effect="plain"
                    >模块 {{ analysisResult.modules.length }}</el-tag
                  >
                  <el-tag
                    v-if="analysisClientPages.length"
                    effect="plain"
                    type="success"
                    >客户端页面 {{ analysisClientPages.length }}</el-tag
                  >
                  <el-tag effect="plain"
                    >待确认 {{ analysisResult.missingInfo.length }}</el-tag
                  >
                </div>
              </div>
              <el-collapse class="mt-4">
                <el-collapse-item
                  v-if="analysisResult.missingInfo.length"
                  title="待补充信息"
                  name="analysis-missing"
                  ><div class="flex flex-wrap gap-2">
                    <el-tag
                      v-for="item in analysisResult.missingInfo"
                      :key="item"
                      effect="plain"
                      type="warning"
                      >{{ item }}</el-tag
                    >
                  </div></el-collapse-item
                >
                <el-collapse-item
                  v-if="analysisResult.modules.length"
                  title="模块清单"
                  name="analysis-modules"
                >
                  <div class="space-y-3">
                    <el-collapse>
                      <el-collapse-item
                        v-for="(module, index) in analysisResult.modules"
                        :key="`${module.name}-${index}`"
                        :name="`${module.name || index}`"
                        :title="`${index + 1}. ${module.label || module.name || `模块 ${index + 1}`}`"
                      >
                        <div class="space-y-4">
                          <p class="text-sm leading-6 text-slate-600">
                            {{ module.description || '暂无模块说明。' }}
                          </p>
                          <div
                            v-if="module.fields.length"
                            class="grid grid-cols-1 gap-3 md:grid-cols-2"
                          >
                            <div
                              v-for="field in module.fields"
                              :key="`${module.name}-${field.name}`"
                              class="rounded-2xl border border-slate-200 bg-slate-50 p-3"
                            >
                              <div
                                class="flex flex-wrap items-center justify-between gap-2"
                              >
                                <div>
                                  <p class="text-sm font-semibold text-slate-800">
                                    {{ field.label || field.name }}
                                  </p>
                                  <p class="mt-1 text-xs text-slate-500">
                                    {{ field.name }}
                                  </p>
                                </div>
                                <el-tag size="small" effect="plain">
                                  {{ field.type || 'string' }}
                                </el-tag>
                              </div>
                              <p class="mt-3 text-sm leading-6 text-slate-600">
                                {{ field.description || '暂无字段说明。' }}
                              </p>
                              <div class="mt-3 flex flex-wrap gap-2">
                                <el-tag
                                  v-if="field.required"
                                  size="small"
                                  type="danger"
                                  effect="plain"
                                >
                                  必填
                                </el-tag>
                                <el-tag
                                  v-if="field.dictionary"
                                  size="small"
                                  type="success"
                                  effect="plain"
                                >
                                  字典: {{ field.dictionary }}
                                </el-tag>
                                <el-tag
                                  v-if="field.relation"
                                  size="small"
                                  type="warning"
                                  effect="plain"
                                >
                                  关联: {{ field.relation }}
                                </el-tag>
                              </div>
                            </div>
                          </div>
                        </div>
                      </el-collapse-item>
                    </el-collapse>
                  </div>
                </el-collapse-item>
                <el-collapse-item
                  v-if="analysisClientPages.length"
                  title="客户端页面"
                  name="analysis-client-pages"
                >
                  <div class="space-y-3">
                    <el-collapse>
                      <el-collapse-item
                        v-for="(page, index) in analysisClientPages"
                        :key="`${page.name}-${index}`"
                        :name="`${page.name || index}`"
                        :title="`${index + 1}. ${page.label || page.name || `页面 ${index + 1}`}`"
                      >
                        <div class="space-y-4">
                          <div class="flex flex-wrap gap-2">
                            <el-tag v-if="page.pageType" effect="plain" type="primary">
                              {{ page.pageType }}
                            </el-tag>
                            <el-tag
                              v-for="moduleName in page.targetModules"
                              :key="`${page.name}-${moduleName}`"
                              effect="plain"
                              type="success"
                            >
                              {{ moduleName }}
                            </el-tag>
                          </div>
                          <p class="text-sm leading-6 text-slate-600">
                            {{ page.description || '暂无页面说明。' }}
                          </p>
                          <div
                            v-if="page.fields.length"
                            class="grid grid-cols-1 gap-3 md:grid-cols-2"
                          >
                            <div
                              v-for="field in page.fields"
                              :key="`${page.name}-${field.name}`"
                              class="rounded-2xl border border-slate-200 bg-slate-50 p-3"
                            >
                              <div class="flex flex-wrap items-center justify-between gap-2">
                                <div>
                                  <p class="text-sm font-semibold text-slate-800">
                                    {{ field.label || field.name }}
                                  </p>
                                  <p class="mt-1 text-xs text-slate-500">
                                    {{ field.name }}
                                  </p>
                                </div>
                                <el-tag size="small" effect="plain">
                                  {{ field.displayType || 'text' }}
                                </el-tag>
                              </div>
                              <p class="mt-3 text-sm leading-6 text-slate-600">
                                {{ field.description || '暂无页面字段说明。' }}
                              </p>
                              <div class="mt-3 flex flex-wrap gap-2">
                                <el-tag
                                  v-if="field.required"
                                  size="small"
                                  type="danger"
                                  effect="plain"
                                >
                                  必填
                                </el-tag>
                                <el-tag
                                  v-if="field.sourceModule || field.sourceField"
                                  size="small"
                                  type="info"
                                  effect="plain"
                                >
                                  {{ `${field.sourceModule || '-'}.${
                                    field.sourceField || '-'
                                  }` }}
                                </el-tag>
                              </div>
                            </div>
                          </div>
                          <div v-if="page.interactions.length" class="space-y-2">
                            <h4 class="text-sm font-semibold text-slate-800">页面交互</h4>
                            <div
                              v-for="item in page.interactions"
                              :key="`${page.name}-${item}`"
                              class="rounded border border-slate-200 bg-white px-4 py-3 text-sm leading-6 text-slate-600"
                            >
                              {{ item }}
                            </div>
                          </div>
                          <div v-if="page.relations.length" class="space-y-2">
                            <h4 class="text-sm font-semibold text-slate-800">字段映射关系</h4>
                            <div
                              v-for="item in page.relations"
                              :key="`${page.name}-${item}`"
                              class="rounded border border-slate-200 bg-white px-4 py-3 text-sm leading-6 text-slate-600"
                            >
                              {{ item }}
                            </div>
                          </div>
                        </div>
                      </el-collapse-item>
                    </el-collapse>
                  </div>
                </el-collapse-item>
                <el-collapse-item
                  v-if="analysisResult.suggestions.length"
                  title="实施建议"
                  name="analysis-suggestions"
                  ><div class="space-y-2">
                    <div
                      v-for="item in analysisResult.suggestions"
                      :key="item"
                      class="rounded border border-slate-200 bg-white px-4 py-3 text-sm leading-6 text-slate-600"
                    >
                      {{ item }}
                    </div>
                  </div></el-collapse-item
                >
                </el-collapse>
            </template>
            <el-empty
              v-else
              description="选择会话节点后，这里会展示对应版本的需求分析结果。"
            />
          </template>

          <template v-else>
            <template
              v-if="workflowResult.steps.length || workflowResult.rawText"
            >
              <div class="rounded bg-slate-50 p-4">
                <p
                  class="text-xs font-semibold uppercase tracking-[0.2em] text-sky-600"
                >
                  Workflow Summary
                </p>
                <p
                  class="mt-2 whitespace-pre-wrap text-sm leading-6 text-slate-700"
                >
                  {{ workflowResult.summary || '已生成 Prompt 工作流。' }}
                </p>
                <div class="mt-3 flex flex-wrap gap-2">
                  <el-tag effect="plain"
                    >步骤 {{ workflowResult.steps.length }}</el-tag
                  >
                </div>
              </div>
              <div v-if="workflowResult.steps.length" class="mt-4 space-y-3">
                <div
                  v-for="(step, index) in workflowResult.steps"
                  :key="`${step.title}-${index}`"
                  class="rounded border border-slate-200 bg-white p-4 shadow-sm shadow-slate-100"
                >
                  <div class="flex flex-wrap items-start justify-between gap-3">
                    <div class="flex items-start gap-3">
                      <div
                        class="flex h-9 w-9 items-center justify-center rounded-full bg-sky-600 text-sm font-semibold text-white"
                      >
                        {{ index + 1 }}
                      </div>
                      <div>
                        <h3 class="text-base font-semibold text-slate-800">
                          {{ step.title || `步骤 ${index + 1}` }}
                        </h3>
                        <p class="mt-1 text-sm leading-6 text-slate-500">
                          {{ step.goal || '建议按顺序执行这一轮 Prompt。' }}
                        </p>
                      </div>
                    </div>
                    <div class="flex flex-wrap gap-2">
                      <el-tag
                        v-if="step.suggestedTool"
                        effect="plain"
                        type="success"
                        >{{ step.suggestedTool }}</el-tag
                      >
                      <el-tag
                        :type="step.autoExecutable ? 'danger' : 'info'"
                        effect="plain"
                        >{{
                          step.autoExecutable ? '可自动执行' : '建议人工确认'
                        }}</el-tag
                      >
                    </div>
                  </div>
                  <div class="mt-4 rounded bg-slate-950 p-4">
                    <div
                      class="mb-2 text-xs font-semibold uppercase tracking-[0.18em] text-slate-400"
                    >
                      Prompt
                    </div>
                    <pre
                      class="max-h-[260px] overflow-auto whitespace-pre-wrap text-xs leading-6 text-slate-100"
                      >{{ step.prompt || '模型未返回该步骤 Prompt。' }}</pre
                    >
                  </div>
                  <div
                    v-if="step.expectedOutput"
                    class="mt-4 rounded border border-dashed border-slate-300 bg-slate-50 p-4"
                  >
                    <div
                      class="mb-2 text-xs font-semibold uppercase tracking-[0.18em] text-slate-400"
                    >
                      Expected Output
                    </div>
                    <p
                      class="whitespace-pre-wrap text-sm leading-6 text-slate-600"
                    >
                      {{ step.expectedOutput }}
                    </p>
                  </div>
                </div>
              </div>
            </template>
            <el-empty
              v-else
              description="选择会话节点后，这里会展示对应版本的工作流结果。"
            />
          </template>
        </div>
      </el-card>
    </div>
    <el-drawer
      v-model="drawerVisible"
      :title="drawerTitle"
      size="50%"
      destroy-on-close
    >
      <pre
        class="h-full overflow-auto whitespace-pre-wrap rounded bg-slate-950 p-4 text-xs leading-6 text-slate-100"
      >{{ drawerContent }}</pre>
    </el-drawer>
  </div>
</template>

<script setup>
import { computed, nextTick, onMounted, reactive, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Delete,
  DocumentCopy,
  MagicStick,
  Plus,
  Promotion,
  RefreshRight
} from '@element-plus/icons-vue'
import {
  analyzeRequirementByAISSEStream,
  dumpAIWorkflowMarkdown,
  deleteAIWorkflowSession,
  generatePromptFlowByAISSEStream,
  getAIWorkflowSessionDetail,
  getAIWorkflowSessionList,
  saveAIWorkflowSession
} from '@/api/autoCode'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({ name: 'AIWorkflow' })

const router = useRouter()
const userStore = useUserStore()
const SETTINGS_KEY = 'gva_ai_workflow_settings'
const ACTIVE_SESSION_KEY = 'gva_ai_workflow_active_session_ids'
const TAB_MODE_MAP = {
  analysis: 'analysisChat',
  workflow: 'workflowPromptChat'
}
const FLOW_TYPE_LABEL_MAP = {
  gva_codegen: 'GVA 代码生成',
  gva_polish: 'GVA 功能完善',
  mcp_assist: 'MCP 使用指导'
}
const defaultSettings = { extraPayload: '' }
const newAnalysisForm = () => ({
  requirement: '',
  packageType: 'auto',
  businessScene: '',
  extraConstraints: '',
  hasClientPage: false,
  clientPageDescription: '',
  clientPageConstraints: ''
})
const newWorkflowForm = () => ({
  source: '',
  flowType: 'gva_codegen',
  extraConstraints: ''
})
const emptyAnalysis = () => ({
  summary: '',
  recommendedPackageType: '',
  missingInfo: [],
  suggestions: [],
  modules: [],
  clientPages: [],
  rawText: '',
  rawJson: ''
})
const emptyWorkflow = () => ({
  summary: '',
  steps: [],
  rawText: '',
  rawJson: ''
})
const newSession = (tab) => ({
  id: 0,
  tab,
  title: '',
  summary: '',
  conversationId: '',
  messageId: '',
  currentNodeId: '',
  formData: {},
  resultData: {},
  messages: []
})
const parseJSON = (value) => {
  try {
    return JSON.parse(value)
  } catch (error) {
    return null
  }
}
const clone = (value, fallback = {}) => {
  try {
    return JSON.parse(JSON.stringify(value ?? fallback))
  } catch (error) {
    return fallback
  }
}
const firstText = (...values) =>
  values.find((item) => typeof item === 'string' && item.trim()) || ''
const firstParagraph = (text) =>
  String(text || '')
    .split(/\n{2,}/)
    .map((item) => item.trim())
    .find(Boolean) || ''
const formatPayload = (payload) => {
  try {
    return JSON.stringify(payload, null, 2)
  } catch (error) {
    return String(payload || '')
  }
}
const normalizeText = (value) =>
  String(value || '')
    .replace(/\r\n/g, '\n')
    .trim()
const toArray = (value) => (Array.isArray(value) ? value : [])
const loadStorage = (key, fallback) => {
  try {
    return { ...fallback, ...JSON.parse(localStorage.getItem(key) || '{}') }
  } catch (error) {
    return { ...fallback }
  }
}
const activeTab = ref('analysis')
const followUpInput = ref('')
const historyLoading = ref(false)
const analysisLoading = ref(false)
const workflowLoading = ref(false)
const dumpLoading = reactive({ analysis: false, workflow: false })
const conversationListRef = ref(null)
const settings = reactive(loadStorage(SETTINGS_KEY, defaultSettings))
const activeSessionIds = reactive(
  loadStorage(ACTIVE_SESSION_KEY, { analysis: 0, workflow: 0 })
)
const historyKeywords = reactive({ analysis: '', workflow: '' })
const sessionLists = reactive({ analysis: [], workflow: [] })
const sessions = reactive({
  analysis: newSession('analysis'),
  workflow: newSession('workflow')
})
const analysisForm = reactive(newAnalysisForm())
const workflowForm = reactive(newWorkflowForm())
const analysisResult = ref(emptyAnalysis())
const workflowResult = ref(emptyWorkflow())
const currentSession = computed(() => sessions[activeTab.value])
const currentSessionList = computed(() => sessionLists[activeTab.value])
const currentLoading = computed(() =>
  activeTab.value === 'analysis' ? analysisLoading.value : workflowLoading.value
)
const streamingPreviewText = ref('')
const streamingPreviewRef = ref(null)
const drawerVisible = ref(false)
const drawerTitle = ref('')
const drawerContent = ref('')
const openDrawer = (title, content) => {
  drawerTitle.value = title
  drawerContent.value = content
  drawerVisible.value = true
}
const analysisClientPages = computed(() =>
  Array.isArray(analysisResult.value?.clientPages)
    ? analysisResult.value.clientPages
    : []
)
const hasAnalysisResult = computed(() =>
  Boolean(
    analysisResult.value.summary ||
      analysisResult.value.modules.length ||
      analysisClientPages.value.length ||
      analysisResult.value.missingInfo.length ||
      analysisResult.value.rawText
  )
)
const canSendFollowUp = computed(() =>
  Boolean(currentSession.value.messages.length && followUpInput.value.trim())
)
const selectedAssistantMessage = computed(
  () =>
    currentSession.value.messages.find(
      (item) =>
        item.role === 'assistant' &&
        item.id === currentSession.value.currentNodeId
    ) || null
)
const currentUser = computed(() =>
  String(
    userStore.userInfo.ID ||
      userStore.userInfo.uuid ||
      userStore.userInfo.id ||
      'gva-ai-workflow'
  )
)
const hasWorkflowResult = computed(() =>
  Boolean(
    workflowResult.value.summary ||
      workflowResult.value.steps.length ||
      workflowResult.value.rawText
  )
)
const canDumpCurrentMarkdown = computed(() =>
  activeTab.value === 'analysis' ? hasAnalysisResult.value : hasWorkflowResult.value
)
const currentDumpLoading = computed(() => dumpLoading[activeTab.value])
const isSessionActive = (item) =>
  Number(item.ID || item.id || 0) ===
  Number(activeSessionIds[activeTab.value] || 0)
const formatTime = (value) => {
  const date = new Date(value)
  return Number.isNaN(date.getTime())
    ? String(value || '-')
    : date.toLocaleString()
}

watch(
  () => ({ ...settings }),
  (value) => localStorage.setItem(SETTINGS_KEY, JSON.stringify(value)),
  { deep: true }
)
watch(
  () => ({ ...activeSessionIds }),
  (value) => localStorage.setItem(ACTIVE_SESSION_KEY, JSON.stringify(value)),
  { deep: true }
)
watch(
  () => ({ ...analysisForm }),
  (value) => {
    sessions.analysis.formData = clone(value)
  },
  { deep: true }
)
watch(
  () => ({ ...workflowForm }),
  (value) => {
    sessions.workflow.formData = clone(value)
  },
  { deep: true }
)
watch(
  analysisResult,
  (value) => {
    sessions.analysis.resultData = clone(value)
    sessions.analysis.summary = firstText(value.summary)
  },
  { deep: true }
)
watch(
  workflowResult,
  (value) => {
    sessions.workflow.resultData = clone(value)
    sessions.workflow.summary = firstText(value.summary)
  },
  { deep: true }
)
watch(
  () => [activeTab.value, currentSession.value.messages.length],
  async () => {
    await nextTick()
    if (conversationListRef.value)
      conversationListRef.value.scrollTop =
        conversationListRef.value.scrollHeight
  },
  { flush: 'post' }
)

const extractJson = (text) => {
  if (!text || typeof text !== 'string') return null
  const direct = parseJSON(text.trim())
  if (direct) return direct
  const code = text.match(/```(?:json)?\s*([\s\S]*?)```/i)?.[1]
  if (code) {
    const parsed = parseJSON(code.trim())
    if (parsed) return parsed
  }
  const starts = [text.indexOf('{'), text.indexOf('[')].filter((i) => i >= 0)
  const ends = [text.lastIndexOf('}'), text.lastIndexOf(']')].filter(
    (i) => i >= 0
  )
  return starts.length && ends.length
    ? parseJSON(text.slice(Math.min(...starts), Math.max(...ends) + 1))
    : null
}

const normalizeStrings = (value) => {
  if (!value) return []
  if (Array.isArray(value))
    return value
      .map((item) =>
        typeof item === 'string'
          ? item.trim()
          : firstText(
              item?.label,
              item?.name,
              item?.title,
              JSON.stringify(item)
            )
      )
      .filter(Boolean)
  if (typeof value === 'string')
    return value
      .split(/\n|,|；|;/)
      .map((item) => item.trim())
      .filter(Boolean)
  return []
}

const normalizeFields = (value) =>
  toArray(value).map((field, index) =>
    typeof field === 'string'
      ? {
          name: field,
          label: field,
          type: 'string',
          required: false,
          description: ''
        }
      : {
          name: firstText(
            field.name,
            field.fieldName,
            field.key,
            `field_${index + 1}`
          ),
          label: firstText(
            field.label,
            field.fieldLabel,
            field.title,
            field.name,
            `字段 ${index + 1}`
          ),
          type: firstText(field.type, field.fieldType, 'string'),
          required: Boolean(field.required),
          description: firstText(field.description, field.comment, field.desc),
          dictionary: firstText(
            field.dictionary,
            field.dictionaryName,
            field.dict
          ),
          relation: firstText(
            field.relation,
            field.relationship,
            field.association
          )
        }
  )

const normalizeModules = (value) =>
  toArray(value).map((module, index) =>
    typeof module === 'string'
      ? {
          name: `module_${index + 1}`,
          label: module,
          description: '',
          fields: []
        }
      : {
          name: firstText(
            module.name,
            module.moduleName,
            `module_${index + 1}`
          ),
          label: firstText(
            module.label,
            module.title,
            module.cnName,
            module.name,
            `模块 ${index + 1}`
          ),
          description: firstText(
            module.description,
            module.goal,
            module.summary
          ),
          fields: normalizeFields(
            module.fields || module.fieldList || module.columns || []
          )
        }
  )

const normalizeClientPages = (value) =>
  toArray(value).map((page, index) => ({
    name: firstText(page.name, page.pageName, `page_${index + 1}`),
    label: firstText(page.label, page.title, page.name, `页面 ${index + 1}`),
    description: firstText(page.description, page.summary, page.goal),
    pageType: firstText(page.pageType, page.type, page.viewType),
    targetModules: normalizeStrings(
      page.targetModules || page.modules || page.relatedModules
    ),
    fields: toArray(page.fields).map((field, fieldIndex) => ({
      name: firstText(field.name, field.fieldName, `page_field_${fieldIndex + 1}`),
      label: firstText(
        field.label,
        field.title,
        field.name,
        `页面字段 ${fieldIndex + 1}`
      ),
      sourceModule: firstText(field.sourceModule, field.module, field.moduleName),
      sourceField: firstText(field.sourceField, field.field, field.backendField),
      displayType: firstText(field.displayType, field.component, field.type, 'text'),
      required: Boolean(field.required),
      description: firstText(field.description, field.comment, field.desc)
    })),
    interactions: normalizeStrings(
      page.interactions || page.actions || page.behaviors
    ),
    relations: normalizeStrings(
      page.relations || page.mappings || page.fieldRelations
    )
  }))

const normalizeSteps = (value, fallback) => {
  const list = toArray(value)
  if (!list.length) {
    return String(fallback || '')
      .split(/\n(?=(?:步骤|Step)\s*\d+|(?:\d+)\.\s)/i)
      .map((item) => item.trim())
      .filter(Boolean)
      .map((item, index) => ({
        title: `步骤 ${index + 1}`,
        goal: '',
        prompt: item,
        expectedOutput: '',
        suggestedTool: '',
        autoExecutable: false
      }))
  }
  return list.map((step, index) =>
    typeof step === 'string'
      ? {
          title: `步骤 ${index + 1}`,
          goal: '',
          prompt: step,
          expectedOutput: '',
          suggestedTool: '',
          autoExecutable: false
        }
      : {
          title: firstText(
            step.title,
            step.name,
            step.stepName,
            `步骤 ${index + 1}`
          ),
          goal: firstText(step.goal, step.description, step.objective),
          prompt: firstText(
            step.prompt,
            step.content,
            step.instruction,
            step.text
          ),
          expectedOutput: firstText(
            step.expectedOutput,
            step.expected,
            step.output
          ),
          suggestedTool: firstText(step.suggestedTool, step.tool, step.action),
          autoExecutable: Boolean(step.autoExecutable)
        }
  )
}

const normalizeAnalysis = (payload, rawText = '') => ({
  summary: firstText(
    payload?.summary,
    payload?.overview,
    payload?.analysisSummary,
    firstParagraph(rawText)
  ),
  recommendedPackageType: firstText(
    payload?.recommendedPackageType,
    payload?.packageType,
    payload?.targetType
  ),
  missingInfo: normalizeStrings(
    payload?.missingInfo ||
      payload?.missing_info ||
      payload?.questions ||
      payload?.clarifyQuestions
  ),
  suggestions: normalizeStrings(
    payload?.suggestions ||
      payload?.recommendations ||
      payload?.advice ||
      payload?.tips
  ),
  clientPages: normalizeClientPages(
    payload?.clientPages || payload?.pages || payload?.client_pages
  ),
  modules: normalizeModules(
    payload?.modules ||
      payload?.moduleList ||
      payload?.entities ||
      payload?.items
  ),
  rawText,
  rawJson: formatPayload(payload || rawText)
})

const normalizeWorkflow = (payload, rawText = '') => ({
  summary: firstText(
    payload?.summary,
    payload?.overview,
    payload?.workflowSummary,
    firstParagraph(rawText)
  ),
  steps: normalizeSteps(
    payload?.steps ||
      payload?.workflow ||
      payload?.prompts ||
      payload?.promptFlow ||
      payload?.promptList,
    rawText
  ),
  rawText,
  rawJson: formatPayload(payload || rawText)
})

const parseDisplay = (content) => {
  const text = normalizeText(content)
  const think = normalizeText(
    text.match(/<think(?:ing)?>\s*([\s\S]*?)<\/think(?:ing)?>/i)?.[1]
  )
  const rest = normalizeText(
    text.replace(/<think(?:ing)?>\s*[\s\S]*?<\/think(?:ing)?>/gi, '')
  )
  const raw = extractJson(rest || text)
    ? formatPayload(extractJson(rest || text))
    : rest || text
  return {
    think,
    rawFile: raw,
    preview: `${(raw || think).slice(0, 140)}${
      (raw || think).length > 140 ? '...' : ''
    }`,
    hasThink: Boolean(think),
    hasRawFile: Boolean(raw)
  }
}

const currentSessionDisplayMessages = computed(() =>
  currentSession.value.messages.map((item) => ({
    ...item,
    display: item.role === 'assistant' ? parseDisplay(item.content) : null,
    isStreaming:
      currentLoading.value &&
      item.role === 'assistant' &&
      item.id === currentSession.value.currentNodeId,
    isSelected:
      item.role === 'assistant' &&
      item.id === currentSession.value.currentNodeId
  }))
)

const STRUCTURED_RESULT_KEYS = [
  'summary',
  'overview',
  'analysisSummary',
  'workflowSummary',
  'recommendedPackageType',
  'packageType',
  'targetType',
  'modules',
  'moduleList',
  'entities',
  'items',
  'clientPages',
  'client_pages',
  'pages',
  'missingInfo',
  'missing_info',
  'questions',
  'clarifyQuestions',
  'suggestions',
  'recommendations',
  'advice',
  'tips',
  'steps',
  'workflow',
  'prompts',
  'promptFlow',
  'promptList'
]

const hasStructuredResultShape = (value) =>
  Boolean(
    value &&
      typeof value === 'object' &&
      STRUCTURED_RESULT_KEYS.some((key) =>
        Object.prototype.hasOwnProperty.call(value, key)
      )
  )

const unwrapStructuredResult = (value, depth = 0) => {
  if (depth > 4 || value == null) return null
  if (typeof value === 'string') {
    const parsed = extractJson(value)
    return parsed?.payload && typeof parsed.payload === 'object'
      ? parsed.payload
      : parsed
  }
  if (Array.isArray(value)) return value
  if (typeof value !== 'object') return null
  if (hasStructuredResultShape(value)) return value

  const candidates = [
    value?.structured,
    value?.payload,
    value?.outputs,
    value?.output,
    value?.result,
    value?.results,
    value?.data?.outputs,
    value?.data?.output,
    value?.data?.result,
    value?.data?.payload,
    value?.data,
    value?.answer,
    value?.text,
    value?.content
  ]

  for (const candidate of candidates) {
    const resolved = unwrapStructuredResult(candidate, depth + 1)
    if (resolved) return resolved
  }

  return null
}

const fallbackAnswerText = (value) => {
  if (typeof value === 'string') return value
  if (value && typeof value === 'object') return formatPayload(value)
  return ''
}

const unwrap = (response) => {
  if (!response) return {}
  if (typeof response.code !== 'undefined') return response.data || {}
  if (
    typeof response.status !== 'undefined' &&
    typeof response.data !== 'undefined'
  )
    return response.data
  return response
}

const resolveChat = (response) => {
  const raw = unwrap(response)
  const structuredCandidate = unwrapStructuredResult(raw)
  const answerText =
    firstText(raw.answer, raw.text, raw.output, raw.content) ||
    fallbackAnswerText(structuredCandidate)
  const parsed = extractJson(answerText)
  const structuredFromText =
    parsed?.payload && typeof parsed.payload === 'object'
      ? parsed.payload
      : parsed
  return {
    answerText,
    structured: structuredFromText || structuredCandidate,
    conversationId: firstText(raw.conversation_id, raw.conversationId),
    messageId: firstText(raw.message_id, raw.messageId)
  }
}

const messageSnapshot = (tab, message) => {
  const payload = Object.keys(message.snapshot || {}).length
    ? message.snapshot
    : extractJson(message.content)
  return tab === 'analysis'
    ? normalizeAnalysis(payload, firstText(payload?.rawText, message.content))
    : normalizeWorkflow(payload, firstText(payload?.rawText, message.content))
}

const applyMessage = (tab, message) => {
  if (!message) {
    if (tab === 'analysis') analysisResult.value = emptyAnalysis()
    else workflowResult.value = emptyWorkflow()
    return
  }
  const snapshot = messageSnapshot(tab, message)
  if (tab === 'analysis') analysisResult.value = snapshot
  else workflowResult.value = snapshot
  sessions[tab].resultData = clone(snapshot)
  sessions[tab].summary = firstText(snapshot.summary)
}

const applySession = (tab) => {
  const session = sessions[tab]
  if (tab === 'analysis')
    Object.assign(analysisForm, {
      ...newAnalysisForm(),
      ...clone(session.formData)
    })
  else
    Object.assign(workflowForm, {
      ...newWorkflowForm(),
      ...clone(session.formData)
    })
  const selected =
    session.messages.find(
      (item) => item.role === 'assistant' && item.id === session.currentNodeId
    ) ||
    [...session.messages].reverse().find((item) => item.role === 'assistant')
  if (selected) {
    session.currentNodeId = selected.id
    applyMessage(tab, selected)
  } else if (tab === 'analysis') {
    analysisResult.value = Object.keys(session.resultData || {}).length
      ? normalizeAnalysis(
          session.resultData,
          firstText(session.resultData?.rawText)
        )
      : emptyAnalysis()
  } else {
    workflowResult.value = Object.keys(session.resultData || {}).length
      ? normalizeWorkflow(
          session.resultData,
          firstText(session.resultData?.rawText)
        )
      : emptyWorkflow()
  }
}

const sessionTitle = (tab) => {
  const firstUser = sessions[tab].messages.find(
    (item) => item.role === 'user'
  )?.content
  return firstText(
    firstUser,
    tab === 'analysis' ? analysisForm.requirement : workflowForm.source,
    tab === 'analysis'
      ? analysisResult.value.summary
      : workflowResult.value.summary
  ).slice(0, 120)
}

const normalizeInlineText = (value) =>
  normalizeText(value).replace(/\s+/g, ' ').trim()

const formatModuleFieldForTransfer = (field) => {
  const meta = [field.type || 'string']
  if (field.required) meta.push('必填')
  if (field.dictionary) meta.push(`字典:${field.dictionary}`)
  if (field.relation) meta.push(`关联:${field.relation}`)

  const lines = [
    `- ${field.label || field.name}（${field.name}）｜${meta.join('｜')}`
  ]

  if (field.description) {
    lines.push(`  说明：${normalizeInlineText(field.description)}`)
  }

  return lines.join('\n')
}

const buildAnalysisTransferText = () => {
  const blocks = [
    '# 需求分析结果',
    '请直接基于以下需求分析结果生成工作流，不要再输出“未明确的需求”、“待确认项”或类似表述。'
  ]

  if (analysisForm.requirement.trim()) {
    blocks.push(`原始需求：${normalizeInlineText(analysisForm.requirement)}`)
  }
  if (analysisForm.businessScene.trim()) {
    blocks.push(`业务场景：${normalizeInlineText(analysisForm.businessScene)}`)
  }
  if (analysisForm.packageType) {
    blocks.push(`目标形态：${analysisForm.packageType}`)
  }
  if (analysisForm.extraConstraints.trim()) {
    blocks.push(`额外约束：${normalizeInlineText(analysisForm.extraConstraints)}`)
  }
  if (analysisForm.hasClientPage) {
    blocks.push('存在客户端页面：是')
  }
  if (analysisForm.clientPageDescription.trim()) {
    blocks.push(
      `客户端页面说明：${normalizeInlineText(analysisForm.clientPageDescription)}`
    )
  }
  if (analysisForm.clientPageConstraints.trim()) {
    blocks.push(
      `客户端页面额外约束：${normalizeInlineText(
        analysisForm.clientPageConstraints
      )}`
    )
  }
  if (analysisResult.value.summary) {
    blocks.push(`分析摘要：${normalizeInlineText(analysisResult.value.summary)}`)
  }
  if (analysisResult.value.recommendedPackageType) {
    blocks.push(`推荐形态：${analysisResult.value.recommendedPackageType}`)
  }
  if (analysisResult.value.suggestions.length) {
    blocks.push(`实施建议：${analysisResult.value.suggestions.join('；')}`)
  }

  if (analysisResult.value.modules.length) {
    const moduleText = analysisResult.value.modules
      .map((module, index) => {
        const lines = [`${index + 1}. ${module.label || module.name || `模块 ${index + 1}`}`]

        if (module.description) {
          lines.push(`   模块说明：${normalizeInlineText(module.description)}`)
        }

        if (module.fields.length) {
          lines.push('   字段：')
          lines.push(
            ...module.fields.map((field) => `   ${formatModuleFieldForTransfer(field)}`)
          )
        }

        return lines.join('\n')
      })
      .join('\n\n')

    blocks.push(`模块清单：\n${moduleText}`)
  }

  if (analysisClientPages.value.length) {
    const pageText = analysisClientPages.value
      .map((page, index) => {
        const lines = [
          `${index + 1}. ${page.label || page.name || `页面 ${index + 1}`}`
        ]

        if (page.pageType) {
          lines.push(`   页面类型：${page.pageType}`)
        }
        if (page.description) {
          lines.push(`   页面说明：${normalizeInlineText(page.description)}`)
        }
        if (page.targetModules.length) {
          lines.push(`   关联模块：${page.targetModules.join('、')}`)
        }
        if (page.fields.length) {
          lines.push('   页面字段：')
          lines.push(
            ...page.fields.map((field) => {
              const source = [field.sourceModule, field.sourceField]
                .filter(Boolean)
                .join('.')
              const meta = [field.displayType || 'text']
              if (field.required) meta.push('必填')
              if (source) meta.push(`映射:${source}`)
              const text = `- ${field.label || field.name}（${field.name}）｜${meta.join('｜')}`
              return `   ${text}${field.description ? `\n     说明：${normalizeInlineText(field.description)}` : ''}`
            })
          )
        }
        if (page.interactions.length) {
          lines.push(`   页面交互：${page.interactions.join('；')}`)
        }
        if (page.relations.length) {
          lines.push(`   字段映射关系：${page.relations.join('；')}`)
        }

        return lines.join('\n')
      })
      .join('\n\n')

    blocks.push(`客户端页面：\n${pageText}`)
  }

  return blocks.join('\n\n')
}

const sessionPayload = (tab) => ({
  id: sessions[tab].id,
  tab,
  title: sessions[tab].title || sessionTitle(tab),
  summary:
    sessions[tab].summary || firstText(sessions[tab].resultData?.summary),
  conversationId: sessions[tab].conversationId,
  messageId: sessions[tab].messageId,
  currentNodeId: sessions[tab].currentNodeId,
  settings: { extraPayload: settings.extraPayload },
  formData: clone(sessions[tab].formData),
  resultData: clone(sessions[tab].resultData),
  messages: sessions[tab].messages.map((item) => ({
    id: item.id,
    role: item.role,
    content: item.content,
    snapshot: clone(item.snapshot),
    conversationId: item.conversationId,
    messageId: item.messageId,
    createdAt: item.createdAt
  }))
})

const dumpCurrentMarkdown = async () => {
  const tab = activeTab.value
  const hasResult =
    tab === 'analysis' ? hasAnalysisResult.value : hasWorkflowResult.value

  if (!hasResult) {
    ElMessage.warning(
      tab === 'analysis' ? '当前还没有可落盘的分析结果' : '当前还没有可落盘的 Prompt 工作流'
    )
    return
  }

  try {
    await persistSession(tab, false)
  } catch (error) {
    // 落盘优先使用当前内存态数据，保存失败不阻断后续操作
  }

  dumpLoading[tab] = true
  try {
    const data = unwrap(await dumpAIWorkflowMarkdown(sessionPayload(tab)))
    const result = data.result || data
    const path = firstText(result.filePath, result.relativePath)
    await ElMessageBox.alert(
      path || 'Markdown 已落盘完成',
      tab === 'analysis' ? '分析结果已落盘' : 'Prompt 工作流已落盘',
      {
        confirmButtonText: '知道了'
      }
    )
  } finally {
    dumpLoading[tab] = false
  }
}

const startNewConversation = (tab = activeTab.value) => {
  sessions[tab] = newSession(tab)
  activeSessionIds[tab] = 0
  if (tab === 'analysis') {
    Object.assign(analysisForm, newAnalysisForm())
    analysisResult.value = emptyAnalysis()
  } else {
    Object.assign(workflowForm, newWorkflowForm())
    workflowResult.value = emptyWorkflow()
  }
  if (activeTab.value === tab) followUpInput.value = ''
  ElMessage.success(
    tab === 'analysis' ? '已新建需求分析会话' : '已新建工作流会话'
  )
}

const hydrateSession = (tab, raw) => {
  sessions[tab] = {
    ...newSession(tab),
    id: Number(raw?.ID || raw?.id || 0),
    tab: firstText(raw?.tab) || tab,
    title: firstText(raw?.title),
    summary: firstText(raw?.summary),
    conversationId: firstText(raw?.conversationId, raw?.conversation_id),
    messageId: firstText(raw?.messageId, raw?.message_id),
    currentNodeId: firstText(raw?.currentNodeId),
    formData: clone(raw?.formData),
    resultData: clone(raw?.resultData),
    messages: toArray(raw?.messages).map((item) => ({
      id:
        firstText(item?.id) ||
        `msg-${Date.now()}-${Math.random().toString(36).slice(2, 8)}`,
      role: firstText(item?.role) || 'assistant',
      content: firstText(item?.content),
      snapshot: clone(item?.snapshot),
      conversationId: firstText(item?.conversationId, item?.conversation_id),
      messageId: firstText(item?.messageId, item?.message_id),
      createdAt: firstText(item?.createdAt)
    }))
  }
  activeSessionIds[tab] = sessions[tab].id
  if (activeTab.value === tab) applySession(tab)
}

const loadSessionList = async (tab = activeTab.value) => {
  historyLoading.value = true
  try {
    const data = unwrap(
      await getAIWorkflowSessionList({
        page: 1,
        pageSize: 50,
        tab,
        keyword: historyKeywords[tab]
      })
    )
    sessionLists[tab] = toArray(data.list)
  } finally {
    historyLoading.value = false
  }
}

const openSession = async (item) => {
  const data = unwrap(
    await getAIWorkflowSessionDetail({ id: Number(item.ID || item.id || 0) })
  )
  if (data.session) hydrateSession(activeTab.value, data.session)
  followUpInput.value = ''
}

const persistSession = async (tab, refreshList = true) => {
  const hasResult =
    tab === 'analysis' ? hasAnalysisResult.value : hasWorkflowResult.value
  if (!sessions[tab].messages.length && !hasResult) return null
  const data = unwrap(await saveAIWorkflowSession(sessionPayload(tab)))
  if (data.session) hydrateSession(tab, data.session)
  if (refreshList) await loadSessionList(tab)
  return data.session || null
}

const removeSession = async (item) => {
  await ElMessageBox.confirm(
    '删除后将无法恢复这条历史需求，是否继续？',
    '删除会话',
    {
      type: 'warning',
      confirmButtonText: '删除',
      cancelButtonText: '取消'
    }
  )
  await deleteAIWorkflowSession({ id: Number(item.ID || item.id || 0) })
  if (isSessionActive(item)) startNewConversation(activeTab.value)
  await loadSessionList(activeTab.value)
  ElMessage.success('已删除会话')
}

const selectMessageNode = async (messageId) => {
  const message = currentSession.value.messages.find(
    (item) => item.role === 'assistant' && item.id === messageId
  )
  if (!message) return
  currentSession.value.currentNodeId = messageId
  applyMessage(activeTab.value, message)
  if (currentSession.value.id) await persistSession(activeTab.value, false)
}

const rollbackToMessage = async (messageId) => {
  const idx = currentSession.value.messages.findIndex(
    (item) => item.role === 'assistant' && item.id === messageId
  )
  if (idx < 0) return
  await ElMessageBox.confirm(
    '回滚后会删除该节点之后的所有消息，并从当前节点重新开启后续追问，是否继续？',
    '回滚会话',
    {
      type: 'warning',
      confirmButtonText: '回滚',
      cancelButtonText: '取消'
    }
  )
  const message = currentSession.value.messages[idx]
  currentSession.value.messages = currentSession.value.messages.slice(
    0,
    idx + 1
  )
  currentSession.value.currentNodeId = messageId
  currentSession.value.conversationId = ''
  currentSession.value.messageId = ''
  applyMessage(activeTab.value, message)
  await persistSession(activeTab.value)
  ElMessage.success('已回滚到所选节点，后续追问会从这里重新开始')
}

const parseExtraPayload = () => {
  if (!settings.extraPayload.trim()) return {}
  try {
    const parsed = JSON.parse(settings.extraPayload)
    return parsed && typeof parsed === 'object' ? parsed : {}
  } catch (error) {
    ElMessage.error('额外透传参数不是有效 JSON')
    throw error
  }
}

const sendChat = async ({ tab, query, inputs, onProgress }) => {
  let extra = {}
  try {
    extra = parseExtraPayload()
  } catch (error) {
    return null
  }
  const safeExtra = { ...extra }
  delete safeExtra.mode
  delete safeExtra.query
  delete safeExtra.inputs
  delete safeExtra.user
  delete safeExtra.response_mode
  delete safeExtra.conversation_id
  const requestData = {
    ...safeExtra,
    mode: TAB_MODE_MAP[tab],
    query:
      query ||
      (tab === 'analysis'
        ? analysisForm.requirement
        : `请基于当前输入生成${
            FLOW_TYPE_LABEL_MAP[workflowForm.flowType] || 'Prompt 工作流'
          }`),
    inputs,
    user: currentUser.value,
    response_mode: 'streaming',
    scene: 'gva_ai_workflow'
  }
  if (!String(requestData.query || '').trim()) {
    ElMessage.error('当前请求缺少 query，已阻止发送')
    return null
  }
  if (sessions[tab].conversationId)
    requestData.conversation_id = sessions[tab].conversationId
  try {
    return resolveChat(
      await (tab === 'analysis'
        ? analyzeRequirementByAISSEStream(requestData, {
            onMessage: onProgress
          })
        : generatePromptFlowByAISSEStream(requestData, {
            onMessage: onProgress
          }))
    )
  } catch (error) {
    ElMessage.error(error?.message || 'AI 请求失败')
    return null
  }
}

const addMessage = (tab, role, content, extras = {}) => {
  const message = {
    id: `${role}-${Date.now()}-${Math.random().toString(36).slice(2, 8)}`,
    role,
    content,
    snapshot: clone(extras.snapshot),
    conversationId: firstText(extras.conversationId),
    messageId: firstText(extras.messageId),
    createdAt: new Date().toISOString()
  }
  sessions[tab].messages.push(message)
  return message
}

const removeMessages = (tab, messageIds = []) => {
  const idSet = new Set(messageIds.filter(Boolean))
  sessions[tab].messages = sessions[tab].messages.filter(
    (item) => !idSet.has(item.id)
  )
}

const scrollConversationToBottom = async () => {
  await nextTick()
  if (conversationListRef.value) {
    conversationListRef.value.scrollTop = conversationListRef.value.scrollHeight
  }
}

const updateStreamMessage = (tab, message, payload = {}) => {
  if (!message) return
  if (payload.text) {
    message.content = payload.text
    streamingPreviewText.value = payload.text
    nextTick(() => {
      if (streamingPreviewRef.value) {
        streamingPreviewRef.value.scrollTop = streamingPreviewRef.value.scrollHeight
      }
    })
  }
  if (payload.conversationId) {
    message.conversationId = payload.conversationId
    sessions[tab].conversationId = payload.conversationId
  }
  if (payload.messageId) {
    message.messageId = payload.messageId
    sessions[tab].messageId = payload.messageId
  }
  scrollConversationToBottom()
}

const runAnalysis = async () => {
  if (!analysisForm.requirement.trim()) return ElMessage.warning('请先输入需求')
  analysisLoading.value = true
  streamingPreviewText.value = ''
  const streamUserMessage = addMessage(
    'analysis',
    'user',
    analysisForm.requirement
  )
  const streamAssistant = addMessage('analysis', 'assistant', '')
  sessions.analysis.currentNodeId = streamAssistant.id
  scrollConversationToBottom()
  try {
    const result = await sendChat({
      tab: 'analysis',
      query: analysisForm.requirement,
      inputs: {
        requirement: analysisForm.requirement,
        packageType: analysisForm.packageType,
        businessScene: analysisForm.businessScene,
        extraConstraints: analysisForm.extraConstraints,
        hasClientPage: analysisForm.hasClientPage,
        clientPageDescription: analysisForm.clientPageDescription,
        clientPageConstraints: analysisForm.clientPageConstraints
      },
      onProgress: (payload) =>{
        updateStreamMessage('analysis', streamAssistant, payload)
      }
    })
    if (!result) {
      if (!streamAssistant.content) {
        removeMessages('analysis', [streamUserMessage.id, streamAssistant.id])
      }
      return
    }
    const snapshot = normalizeAnalysis(result.structured, result.answerText)
    sessions.analysis.conversationId =
      result.conversationId || sessions.analysis.conversationId
    sessions.analysis.messageId =
      result.messageId || sessions.analysis.messageId
    streamAssistant.content =
      result.answerText || streamAssistant.content || '模型未返回文本内容'
    streamAssistant.snapshot = snapshot
    streamAssistant.conversationId =
      result.conversationId || streamAssistant.conversationId
    streamAssistant.messageId = result.messageId || streamAssistant.messageId
    sessions.analysis.currentNodeId = streamAssistant.id
    sessions.analysis.title = sessionTitle('analysis')
    analysisResult.value = snapshot
    await persistSession('analysis')
    ElMessage.success('需求分析完成')
  } finally {
    analysisLoading.value = false
    streamingPreviewText.value = ''
  }
}

const pushAnalysisToWorkflow = async () => {
  if (!hasAnalysisResult.value) return ElMessage.warning('还没有可用的分析结果')
  await switchTab('workflow')
  workflowForm.source = buildAnalysisTransferText()
  ElMessage.success('已带入 Prompt 工作流')
}

const runWorkflow = async () => {
  if (!workflowForm.source.trim())
    return ElMessage.warning('请先输入需求或分析结果')
  workflowLoading.value = true
  streamingPreviewText.value = ''
  const streamUserMessage = addMessage('workflow', 'user', '')
  const streamAssistant = addMessage('workflow', 'assistant', '')
  sessions.workflow.currentNodeId = streamAssistant.id
  scrollConversationToBottom()
  try {
    const query = `请基于当前输入生成${
      FLOW_TYPE_LABEL_MAP[workflowForm.flowType] || 'Prompt 工作流'
    }`
    streamUserMessage.content = query
    const result = await sendChat({
      tab: 'workflow',
      query,
      inputs: {
        source: workflowForm.source,
        flowType: workflowForm.flowType,
        extraConstraints: workflowForm.extraConstraints
      },
      onProgress: (payload) =>{
        updateStreamMessage('workflow', streamAssistant, payload)
      }
    })
    if (!result) {
      if (!streamAssistant.content) {
        removeMessages('workflow', [streamUserMessage.id, streamAssistant.id])
      }
      return
    }
    const snapshot = normalizeWorkflow(result.structured, result.answerText)
    sessions.workflow.conversationId =
      result.conversationId || sessions.workflow.conversationId
    sessions.workflow.messageId =
      result.messageId || sessions.workflow.messageId
    streamAssistant.content =
      result.answerText || streamAssistant.content || '模型未返回文本内容'
    streamAssistant.snapshot = snapshot
    streamAssistant.conversationId =
      result.conversationId || streamAssistant.conversationId
    streamAssistant.messageId = result.messageId || streamAssistant.messageId
    sessions.workflow.currentNodeId = streamAssistant.id
    sessions.workflow.title = sessionTitle('workflow')
    workflowResult.value = snapshot
    await persistSession('workflow')
    ElMessage.success('Prompt 工作流已生成')
  } finally {
    workflowLoading.value = false
    streamingPreviewText.value = ''
  }
}

const sendFollowUp = async () => {
  const query = followUpInput.value.trim()
  if (!query) return ElMessage.warning('请先输入追问内容')
  if (!currentSession.value.messages.length)
    return ElMessage.warning('当前还没有会话内容，请先发起一次分析或生成')
  const tab = activeTab.value
  if (tab === 'analysis') analysisLoading.value = true
  else workflowLoading.value = true
  streamingPreviewText.value = ''
  const streamUserMessage = addMessage(tab, 'user', query)
  const streamAssistant = addMessage(tab, 'assistant', '')
  sessions[tab].currentNodeId = streamAssistant.id
  scrollConversationToBottom()
  try {
    const result = await sendChat({
      tab,
      query,
      inputs:
        tab === 'analysis'
          ? {
              requirement: analysisForm.requirement,
              packageType: analysisForm.packageType,
              businessScene: analysisForm.businessScene,
              extraConstraints: analysisForm.extraConstraints,
              hasClientPage: analysisForm.hasClientPage,
              clientPageDescription: analysisForm.clientPageDescription,
              clientPageConstraints: analysisForm.clientPageConstraints
            }
          : {
              source: workflowForm.source,
              flowType: workflowForm.flowType,
              extraConstraints: workflowForm.extraConstraints
            },
      onProgress: (payload) => updateStreamMessage(tab, streamAssistant, payload)
    })
    if (!result) {
      if (!streamAssistant.content) {
        removeMessages(tab, [streamUserMessage.id, streamAssistant.id])
      }
      return
    }
    const snapshot =
      tab === 'analysis'
        ? normalizeAnalysis(result.structured, result.answerText)
        : normalizeWorkflow(result.structured, result.answerText)
    sessions[tab].conversationId =
      result.conversationId || sessions[tab].conversationId
    sessions[tab].messageId = result.messageId || sessions[tab].messageId
    streamAssistant.content =
      result.answerText || streamAssistant.content || '模型未返回文本内容'
    streamAssistant.snapshot = snapshot
    streamAssistant.conversationId =
      result.conversationId || streamAssistant.conversationId
    streamAssistant.messageId = result.messageId || streamAssistant.messageId
    sessions[tab].currentNodeId = streamAssistant.id
    sessions[tab].title = sessionTitle(tab)
    if (tab === 'analysis') analysisResult.value = snapshot
    else workflowResult.value = snapshot
    followUpInput.value = ''
    await persistSession(tab)
    ElMessage.success('追问已发送')
  } finally {
    if (tab === 'analysis') analysisLoading.value = false
    else workflowLoading.value = false
    streamingPreviewText.value = ''
  }
}

const copyText = async (text, successMessage = '复制成功') => {
  if (!text) return ElMessage.warning('没有可复制的内容')
  try {
    await navigator.clipboard.writeText(text)
  } catch (error) {
    const textarea = document.createElement('textarea')
    textarea.value = text
    textarea.style.position = 'fixed'
    textarea.style.left = '-9999px'
    document.body.appendChild(textarea)
    textarea.focus()
    textarea.select()
    document.execCommand('copy')
    document.body.removeChild(textarea)
  }
  ElMessage.success(successMessage)
}

const copyAnalysisResult = async () =>
  copyText(
    JSON.stringify(clone(analysisResult.value), null, 2),
    '已复制结构化分析结果'
  )
const copyAllWorkflowPrompts = async () =>
  copyText(
    workflowResult.value.steps
      .map(
        (step, index) =>
          `# ${index + 1}. ${step.title || `步骤 ${index + 1}`}\n${
            step.prompt || ''
          }`
      )
      .join('\n\n'),
    '已复制全部 Prompt'
  )
const copyWorkflowRaw = async () =>
  copyText(
    workflowResult.value.rawText || workflowResult.value.rawJson,
    '已复制原始结果'
  )
const switchTab = async (tab) => {
  activeTab.value = tab
  followUpInput.value = ''
  applySession(tab)
  if (!sessionLists[tab].length) await loadSessionList(tab)
}
const fillExample = () => {
  switchTab('analysis')
  analysisForm.requirement =
    '我要做一个请假管理系统，支持员工提交请假申请、部门负责人审批、人事复核，按状态筛选并导出统计，还需要请假类型和审批状态字典。'
  analysisForm.packageType = 'auto'
  analysisForm.businessScene = 'OA 内部后台'
  analysisForm.extraConstraints =
    '需要附件字段；后续可能补充销假功能；优先生成适合 gin-vue-admin 自动代码的结构化方案。'
  analysisForm.hasClientPage = true
  analysisForm.clientPageDescription =
    '需要一个员工端 H5 页面，用于提交请假申请、查看审批进度、上传附件和发起销假。'
  analysisForm.clientPageConstraints =
    '移动端优先，表单字段尽量精简，审批状态要明显，附件支持图片和 PDF。'
  workflowForm.source = ''
  workflowForm.flowType = 'gva_codegen'
  workflowForm.extraConstraints = '每一步都给出可复制 Prompt，并说明预期输出。'
}
const goAutoCode = () => router.push({ name: 'autoCode' })

onMounted(async () => {
  await Promise.all([loadSessionList('analysis'), loadSessionList('workflow')])
  for (const tab of ['analysis', 'workflow']) {
    const preferred = Number(activeSessionIds[tab] || 0)
    const target = preferred ? { ID: preferred } : sessionLists[tab][0]
    if (target?.ID || target?.id) {
      const data = unwrap(
        await getAIWorkflowSessionDetail({ id: Number(target.ID || target.id) })
      )
      if (data.session) hydrateSession(tab, data.session)
    }
  }
  applySession(activeTab.value)
})
</script>

<style scoped>
.ai-workflow-page :deep(.el-collapse-item__header) {
  padding-left: 6px;
}
</style>
