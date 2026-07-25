<template>
  <div class="gva-table-box config-layout">
    <aside class="config-menu">
      <div class="menu-group">
        <div class="menu-group-title">系统设置</div>
        <div
          class="menu-item"
          :class="{ on: activeKey === 'basic' }"
          @click="activeKey = 'basic'"
        >
          基础设置
        </div>
        <div
          class="menu-item"
          :class="{ on: activeKey === 'switches' }"
          @click="activeKey = 'switches'"
        >
          功能开关
        </div>
      </div>
      <div class="menu-group">
        <div class="menu-group-title">服务</div>
        <div
          class="menu-item"
          :class="{ on: activeKey === 'jwt' }"
          @click="activeKey = 'jwt'"
        >
          JWT 签名
        </div>
        <div
          class="menu-item"
          :class="{ on: activeKey === 'zap' }"
          @click="activeKey = 'zap'"
        >
          Zap 日志
        </div>
        <div
          v-if="config.system['use-redis']"
          class="menu-item"
          :class="{ on: activeKey === 'redis' }"
          @click="activeKey = 'redis'"
        >
          Redis
        </div>
        <div
          class="menu-item"
          :class="{ on: activeKey === 'email' }"
          @click="activeKey = 'email'"
        >
          邮箱配置
        </div>
      </div>
      <div class="menu-group">
        <div class="menu-group-title">存储</div>
        <div
          class="menu-item"
          :class="{ on: activeKey === 'db' }"
          @click="activeKey = 'db'"
        >
          数据库配置
        </div>
        <div
          class="menu-item"
          :class="{ on: activeKey === 'oss' }"
          @click="activeKey = 'oss'"
        >
          OSS 配置
        </div>
        <div
          v-if="config.system['use-mongo']"
          class="menu-item"
          :class="{ on: activeKey === 'mongo' }"
          @click="activeKey = 'mongo'"
        >
          Mongo 数据库
        </div>
      </div>
      <div class="menu-group">
        <div class="menu-group-title">其他</div>
        <div
          class="menu-item"
          :class="{ on: activeKey === 'autocode' }"
          @click="activeKey = 'autocode'"
        >
          自动化代码
        </div>
      </div>
    </aside>

    <div class="config-content">
      <div class="action-bar">
        <el-button type="primary" @click="update">立即更新</el-button>
        <el-button type="primary" plain @click="reload">重载服务</el-button>
      </div>
      <el-form
        ref="form"
        :model="config"
        label-position="top"
        class="config-form"
      >
        <!--  基础设置  -->
        <section v-if="activeKey === 'basic'" class="config-section">
          <div class="section-header">
            <span class="section-bar"></span>
            <div>
              <h3 class="section-title">基础设置</h3>
              <p class="section-desc">服务端口、数据库与存储类型、限流等核心参数</p>
            </div>
          </div>
          <div class="section-grid">
            <el-form-item label="端口值">
              <el-input-number
                v-model="config.system.addr"
                placeholder="请输入端口值"
              />
            </el-form-item>
            <el-form-item label="数据库类型">
              <el-select v-model="config.system['db-type']">
                <el-option value="mysql" />
                <el-option value="pgsql" />
                <el-option value="mssql" />
                <el-option value="sqlite" />
                <el-option value="oracle" />
              </el-select>
            </el-form-item>
            <el-form-item label="Oss类型">
              <el-select v-model="config.system['oss-type']">
                <el-option value="local" label="本地" />
                <el-option value="qiniu" label="七牛" />
                <el-option value="tencent-cos" label="腾讯云COS" />
                <el-option value="aliyun-oss" label="阿里云OSS" />
                <el-option value="huawei-obs" label="华为云OBS" />
                <el-option value="cloudflare-r2" label="cloudflare R2" />
                <el-option value="minio">MinIO</el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="限流次数">
              <el-input-number v-model.number="config.system['iplimit-count']" />
            </el-form-item>
            <el-form-item label="限流时间">
              <el-input-number v-model.number="config.system['iplimit-time']" />
            </el-form-item>
            <el-tooltip
              content="请修改完成后，注意一并修改前端env环境下的VITE_BASE_PATH"
              placement="top-start"
            >
              <el-form-item label="全局路由前缀">
                <el-input
                  v-model.trim="config.system['router-prefix']"
                  placeholder="请输入全局路由前缀"
                />
              </el-form-item>
            </el-tooltip>
          </div>
        </section>

        <!--  功能开关  -->
        <section v-if="activeKey === 'switches'" class="config-section">
          <div class="section-header">
            <span class="section-bar"></span>
            <div>
              <h3 class="section-title">功能开关</h3>
              <p class="section-desc">按需开启或关闭系统功能</p>
            </div>
          </div>
          <div class="section-grid">
            <el-form-item label="多点登录拦截">
              <el-switch v-model="config.system['use-multipoint']" />
            </el-form-item>
            <el-form-item label="开启redis">
              <el-switch v-model="config.system['use-redis']" />
            </el-form-item>
            <el-form-item label="开启Mongo">
              <el-switch v-model="config.system['use-mongo']" />
            </el-form-item>
            <el-form-item label="严格角色模式">
              <el-switch v-model="config.system['use-strict-auth']" />
            </el-form-item>
            <el-form-item label="禁用自动迁移数据库表结构">
              <el-switch v-model="config.system['disable-auto-migrate']" />
            </el-form-item>
          </div>
        </section>

        <!--  JWT 签名  -->
        <section v-if="activeKey === 'jwt'" class="config-section">
          <div class="section-header">
            <span class="section-bar"></span>
            <div>
              <h3 class="section-title">JWT 配置</h3>
              <p class="section-desc">Token 签名密钥与有效期设置</p>
            </div>
          </div>
          <div class="section-grid">
            <el-form-item label="jwt签名">
              <el-input
                v-model.trim="config.jwt['signing-key']"
                placeholder="请输入jwt签名"
              >
                <template #append>
                  <el-button @click="getUUID">生成</el-button>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item label="有效期">
              <el-input
                v-model.trim="config.jwt['expires-time']"
                placeholder="请输入有效期"
              />
            </el-form-item>
            <el-form-item label="缓冲期">
              <el-input
                v-model.trim="config.jwt['buffer-time']"
                placeholder="请输入缓冲期"
              />
            </el-form-item>
            <el-form-item label="签发者">
              <el-input
                v-model.trim="config.jwt.issuer"
                placeholder="请输入签发者"
              />
            </el-form-item>
          </div>
        </section>

        <!--  Zap 日志  -->
        <template v-if="activeKey === 'zap'">
          <section class="config-section">
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">日志输出</h3>
                <p class="section-desc">Zap 日志的级别、格式与输出位置</p>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="级别">
                <el-select v-model="config.zap.level">
                  <el-option value="off" label="关闭" />
                  <el-option value="fatal" label="致命" />
                  <el-option value="error" label="错误" />
                  <el-option value="warn" label="警告" />
                  <el-option value="info" label="信息" />
                  <el-option value="debug" label="调试" />
                  <el-option value="trace" label="跟踪" />
                </el-select>
              </el-form-item>
              <el-form-item label="输出">
                <el-select v-model="config.zap.format">
                  <el-option value="console" label="console" />
                  <el-option value="json" label="json" />
                </el-select>
              </el-form-item>
              <el-form-item label="编码级">
                <el-select v-model="config.zap['encode-level']">
                  <el-option
                    value="LowercaseLevelEncoder"
                    label="LowercaseLevelEncoder"
                  />
                  <el-option
                    value="LowercaseColorLevelEncoder"
                    label="LowercaseColorLevelEncoder"
                  />
                  <el-option
                    value="CapitalLevelEncoder"
                    label="CapitalLevelEncoder"
                  />
                  <el-option
                    value="CapitalColorLevelEncoder"
                    label="CapitalColorLevelEncoder"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="栈名">
                <el-input
                  v-model.trim="config.zap['stacktrace-key']"
                  placeholder="请输入栈名"
                />
              </el-form-item>
              <el-form-item label="日志前缀">
                <el-input
                  v-model.trim="config.zap.prefix"
                  placeholder="请输入日志前缀"
                />
              </el-form-item>
              <el-form-item label="日志文件夹">
                <el-input
                  v-model.trim="config.zap.director"
                  placeholder="请输入日志文件夹"
                />
              </el-form-item>
              <el-form-item label="日志留存时间(默认以天为单位)">
                <el-input-number v-model="config.zap['retention-day']" />
              </el-form-item>
            </div>
          </section>
          <section class="config-section">
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">开关</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="显示行">
                <el-switch v-model="config.zap['show-line']" />
              </el-form-item>
              <el-form-item label="输出控制台">
                <el-switch v-model="config.zap['log-in-console']" />
              </el-form-item>
            </div>
          </section>
        </template>

        <!--  Redis  -->
        <section
          v-if="activeKey === 'redis' && config.system['use-redis']"
          class="config-section"
        >
          <div class="section-header">
            <span class="section-bar"></span>
            <div>
              <h3 class="section-title">Redis 连接</h3>
            </div>
          </div>
          <div class="section-grid">
            <el-form-item label="库">
              <el-input-number v-model="config.redis.db" min="0" max="16" />
            </el-form-item>
            <el-form-item label="地址">
              <el-input
                v-model.trim="config.redis.addr"
                placeholder="请输入地址"
              />
            </el-form-item>
            <el-form-item label="密码">
              <el-input
                v-model.trim="config.redis.password"
                placeholder="请输入密码"
              />
            </el-form-item>
          </div>
        </section>

        <!--  邮箱配置  -->
        <section v-if="activeKey === 'email'" class="config-section">
          <div class="section-header">
            <span class="section-bar"></span>
            <div>
              <h3 class="section-title">SMTP 配置</h3>
              <p class="section-desc">邮件发送服务参数</p>
            </div>
          </div>
          <div class="section-grid">
            <el-form-item label="接收者邮箱">
              <el-input
                v-model="config.email.to"
                placeholder="可多个，以逗号分隔"
              />
            </el-form-item>
            <el-form-item label="端口">
              <el-input-number v-model="config.email.port" />
            </el-form-item>
            <el-form-item label="发送者邮箱">
              <el-input
                v-model.trim="config.email.from"
                placeholder="请输入发送者邮箱"
              />
            </el-form-item>
            <el-form-item label="host">
              <el-input
                v-model.trim="config.email.host"
                placeholder="请输入host"
              />
            </el-form-item>
            <el-form-item label="secret">
              <el-input
                v-model.trim="config.email.secret"
                placeholder="请输入secret"
              />
            </el-form-item>
            <el-form-item label="是否为ssl">
              <el-switch v-model="config.email['is-ssl']" />
            </el-form-item>
            <el-form-item label="是否LoginAuth认证">
              <el-switch v-model="config.email['is-loginauth']" />
            </el-form-item>
            <el-form-item label="测试邮件">
              <el-button @click="email">测试邮件</el-button>
            </el-form-item>
          </div>
        </section>

        <!--  数据库配置  -->
        <template v-if="activeKey === 'db'">
          <section
            v-if="config.system['db-type'] === 'mysql'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">MySQL</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="用户名">
                <el-input
                  v-model.trim="config.mysql.username"
                  placeholder="请输入用户名"
                />
              </el-form-item>
              <el-form-item label="密码">
                <el-input
                  v-model.trim="config.mysql.password"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item label="地址">
                <el-input
                  v-model.trim="config.mysql.path"
                  placeholder="请输入地址"
                />
              </el-form-item>
              <el-form-item label="数据库名称">
                <el-input
                  v-model.trim="config.mysql['db-name']"
                  placeholder="请输入数据库名称"
                />
              </el-form-item>
              <el-form-item label="前缀">
                <el-input
                  v-model.trim="config.mysql['prefix']"
                  placeholder="默认为空"
                />
              </el-form-item>
              <el-form-item label="复数表">
                <el-switch v-model="config.mysql['singular']" />
              </el-form-item>
              <el-form-item label="引擎">
                <el-input
                  v-model.trim="config.mysql['engine']"
                  placeholder="默认为InnoDB"
                />
              </el-form-item>
              <el-form-item label="maxIdleConns">
                <el-input-number
                  v-model="config.mysql['max-idle-conns']"
                  :min="1"
                />
              </el-form-item>
              <el-form-item label="maxOpenConns">
                <el-input-number
                  v-model="config.mysql['max-open-conns']"
                  :min="1"
                />
              </el-form-item>
              <el-form-item label="日志模式">
                <el-select v-model="config.mysql['log-mode']">
                  <el-option value="off" label="关闭" />
                  <el-option value="fatal" label="致命" />
                  <el-option value="error" label="错误" />
                  <el-option value="warn" label="警告" />
                  <el-option value="info" label="信息" />
                  <el-option value="debug" label="调试" />
                  <el-option value="trace" label="跟踪" />
                </el-select>
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['db-type'] === 'pgsql'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">PostgreSQL</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="用户名">
                <el-input
                  v-model="config.pgsql.username"
                  placeholder="请输入用户名"
                />
              </el-form-item>
              <el-form-item label="密码">
                <el-input
                  v-model="config.pgsql.password"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item label="地址">
                <el-input
                  v-model.trim="config.pgsql.path"
                  placeholder="请输入地址"
                />
              </el-form-item>
              <el-form-item label="数据库">
                <el-input
                  v-model.trim="config.pgsql['db-name']"
                  placeholder="请输入数据库"
                />
              </el-form-item>
              <el-form-item label="前缀">
                <el-input
                  v-model.trim="config.pgsql['prefix']"
                  placeholder="请输入前缀"
                />
              </el-form-item>
              <el-form-item label="复数表">
                <el-switch v-model="config.pgsql['singular']" />
              </el-form-item>
              <el-form-item label="引擎">
                <el-input
                  v-model.trim="config.pgsql['engine']"
                  placeholder="请输入引擎"
                />
              </el-form-item>
              <el-form-item label="maxIdleConns">
                <el-input-number v-model="config.pgsql['max-idle-conns']" />
              </el-form-item>
              <el-form-item label="maxOpenConns">
                <el-input-number v-model="config.pgsql['max-open-conns']" />
              </el-form-item>
              <el-form-item label="日志模式">
                <el-select v-model="config.pgsql['log-mode']">
                  <el-option value="off" label="关闭" />
                  <el-option value="fatal" label="致命" />
                  <el-option value="error" label="错误" />
                  <el-option value="warn" label="警告" />
                  <el-option value="info" label="信息" />
                  <el-option value="debug" label="调试" />
                  <el-option value="trace" label="跟踪" />
                </el-select>
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['db-type'] === 'mssql'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">MsSQL</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="用户名">
                <el-input
                  v-model.trim="config.mssql.username"
                  placeholder="请输入用户名"
                />
              </el-form-item>
              <el-form-item label="密码">
                <el-input
                  v-model.trim="config.mssql.password"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item label="地址">
                <el-input
                  v-model.trim="config.mssql.path"
                  placeholder="请输入地址"
                />
              </el-form-item>
              <el-form-item label="端口">
                <el-input
                  v-model.trim="config.mssql.port"
                  placeholder="请输入端口"
                />
              </el-form-item>
              <el-form-item label="数据库">
                <el-input
                  v-model.trim="config.mssql['db-name']"
                  placeholder="请输入数据库"
                />
              </el-form-item>
              <el-form-item label="前缀">
                <el-input
                  v-model.trim="config.mssql['prefix']"
                  placeholder="请输入前缀"
                />
              </el-form-item>
              <el-form-item label="复数表">
                <el-switch v-model="config.mssql['singular']" />
              </el-form-item>
              <el-form-item label="引擎">
                <el-input
                  v-model.trim="config.mssql['engine']"
                  placeholder="请输入引擎"
                />
              </el-form-item>
              <el-form-item label="maxIdleConns">
                <el-input-number v-model="config.mssql['max-idle-conns']" />
              </el-form-item>
              <el-form-item label="maxOpenConns">
                <el-input-number v-model="config.mssql['max-open-conns']" />
              </el-form-item>
              <el-form-item label="日志模式">
                <el-select v-model="config.mssql['log-mode']">
                  <el-option value="off" label="关闭" />
                  <el-option value="fatal" label="致命" />
                  <el-option value="error" label="错误" />
                  <el-option value="warn" label="警告" />
                  <el-option value="info" label="信息" />
                  <el-option value="debug" label="调试" />
                  <el-option value="trace" label="跟踪" />
                </el-select>
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['db-type'] === 'sqlite'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">sqlite</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="用户名">
                <el-input
                  v-model.trim="config.sqlite.username"
                  placeholder="请输入用户名"
                />
              </el-form-item>
              <el-form-item label="密码">
                <el-input
                  v-model.trim="config.sqlite.password"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item label="地址">
                <el-input
                  v-model.trim="config.sqlite.path"
                  placeholder="请输入地址"
                />
              </el-form-item>
              <el-form-item label="端口">
                <el-input
                  v-model.trim="config.sqlite.port"
                  placeholder="请输入端口"
                />
              </el-form-item>
              <el-form-item label="数据库">
                <el-input
                  v-model.trim="config.sqlite['db-name']"
                  placeholder="请输入数据库"
                />
              </el-form-item>
              <el-form-item label="maxIdleConns">
                <el-input-number v-model="config.sqlite['max-idle-conns']" />
              </el-form-item>
              <el-form-item label="maxOpenConns">
                <el-input-number v-model="config.sqlite['max-open-conns']" />
              </el-form-item>
              <el-form-item label="日志模式">
                <el-select v-model="config.sqlite['log-mode']">
                  <el-option value="off" label="关闭" />
                  <el-option value="fatal" label="致命" />
                  <el-option value="error" label="错误" />
                  <el-option value="warn" label="警告" />
                  <el-option value="info" label="信息" />
                  <el-option value="debug" label="调试" />
                  <el-option value="trace" label="跟踪" />
                </el-select>
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['db-type'] === 'oracle'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">oracle</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="用户名">
                <el-input
                  v-model.trim="config.oracle.username"
                  placeholder="请输入用户名"
                />
              </el-form-item>
              <el-form-item label="密码">
                <el-input
                  v-model.trim="config.oracle.password"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item label="地址">
                <el-input
                  v-model.trim="config.oracle.path"
                  placeholder="请输入地址"
                />
              </el-form-item>
              <el-form-item label="数据库名称">
                <el-input
                  v-model.trim="config.oracle['db-name']"
                  placeholder="请输入数据库名称"
                />
              </el-form-item>
              <el-form-item label="前缀">
                <el-input
                  v-model.trim="config.oracle['prefix']"
                  placeholder="默认为空"
                />
              </el-form-item>
              <el-form-item label="复数表">
                <el-switch v-model="config.oracle['singular']" />
              </el-form-item>
              <el-form-item label="引擎">
                <el-input
                  v-model.trim="config.oracle['engine']"
                  placeholder="默认为InnoDB"
                />
              </el-form-item>
              <el-form-item label="maxIdleConns">
                <el-input-number
                  v-model="config.oracle['max-idle-conns']"
                  :min="1"
                />
              </el-form-item>
              <el-form-item label="maxOpenConns">
                <el-input-number
                  v-model="config.oracle['max-open-conns']"
                  :min="1"
                />
              </el-form-item>
              <el-form-item label="日志模式">
                <el-select v-model="config.oracle['log-mode']">
                  <el-option value="off" label="关闭" />
                  <el-option value="fatal" label="致命" />
                  <el-option value="error" label="错误" />
                  <el-option value="warn" label="警告" />
                  <el-option value="info" label="信息" />
                  <el-option value="debug" label="调试" />
                  <el-option value="trace" label="跟踪" />
                </el-select>
              </el-form-item>
            </div>
          </section>
        </template>

        <!--  OSS 配置  -->
        <template v-if="activeKey === 'oss'">
          <section
            v-if="config.system['oss-type'] === 'local'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">本地配置</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="本地文件访问路径">
                <el-input
                  v-model.trim="config.local.path"
                  placeholder="请输入本地文件访问路径"
                />
              </el-form-item>
              <el-form-item label="本地文件存储路径">
                <el-input
                  v-model.trim="config.local['store-path']"
                  placeholder="请输入本地文件存储路径"
                />
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['oss-type'] === 'qiniu'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">七牛上传配置</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="存储区域">
                <el-input
                  v-model.trim="config.qiniu.zone"
                  placeholder="请输入存储区域"
                />
              </el-form-item>
              <el-form-item label="空间名称">
                <el-input
                  v-model.trim="config.qiniu.bucket"
                  placeholder="请输入空间名称"
                />
              </el-form-item>
              <el-form-item label="CDN加速域名">
                <el-input
                  v-model.trim="config.qiniu['img-path']"
                  placeholder="请输入CDN加速域名"
                />
              </el-form-item>
              <el-form-item label="accessKey">
                <el-input
                  v-model.trim="config.qiniu['access-key']"
                  placeholder="请输入accessKey"
                />
              </el-form-item>
              <el-form-item label="secretKey">
                <el-input
                  v-model.trim="config.qiniu['secret-key']"
                  placeholder="请输入secretKey"
                />
              </el-form-item>
              <el-form-item label="是否使用https">
                <el-switch v-model="config.qiniu['use-https']" />
              </el-form-item>
              <el-form-item label="上传是否使用CDN上传加速">
                <el-switch v-model="config.qiniu['use-cdn-domains']" />
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['oss-type'] === 'tencent-cos'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">腾讯云COS上传配置</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="存储桶名称">
                <el-input
                  v-model.trim="config['tencent-cos']['bucket']"
                  placeholder="请输入存储桶名称"
                />
              </el-form-item>
              <el-form-item label="所属地域">
                <el-input
                  v-model.trim="config['tencent-cos'].region"
                  placeholder="请输入所属地域"
                />
              </el-form-item>
              <el-form-item label="secretID">
                <el-input
                  v-model.trim="config['tencent-cos']['secret-id']"
                  placeholder="请输入secretID"
                />
              </el-form-item>
              <el-form-item label="secretKey">
                <el-input
                  v-model.trim="config['tencent-cos']['secret-key']"
                  placeholder="请输入secretKey"
                />
              </el-form-item>
              <el-form-item label="路径前缀">
                <el-input
                  v-model.trim="config['tencent-cos']['path-prefix']"
                  placeholder="请输入路径前缀"
                />
              </el-form-item>
              <el-form-item label="访问域名">
                <el-input
                  v-model.trim="config['tencent-cos']['base-url']"
                  placeholder="请输入访问域名"
                />
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['oss-type'] === 'aliyun-oss'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">阿里云OSS上传配置</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="区域">
                <el-input
                  v-model.trim="config['aliyun-oss'].endpoint"
                  placeholder="请输入区域"
                />
              </el-form-item>
              <el-form-item label="accessKeyId">
                <el-input
                  v-model.trim="config['aliyun-oss']['access-key-id']"
                  placeholder="请输入accessKeyId"
                />
              </el-form-item>
              <el-form-item label="accessKeySecret">
                <el-input
                  v-model.trim="config['aliyun-oss']['access-key-secret']"
                  placeholder="请输入accessKeySecret"
                />
              </el-form-item>
              <el-form-item label="存储桶名称">
                <el-input
                  v-model.trim="config['aliyun-oss']['bucket-name']"
                  placeholder="请输入存储桶名称"
                />
              </el-form-item>
              <el-form-item label="访问域名">
                <el-input
                  v-model.trim="config['aliyun-oss']['bucket-url']"
                  placeholder="请输入访问域名"
                />
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['oss-type'] === 'huawei-obs'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">华为云OBS上传配置</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="路径">
                <el-input
                  v-model.trim="config['hua-wei-obs'].path"
                  placeholder="请输入路径"
                />
              </el-form-item>
              <el-form-item label="存储桶名称">
                <el-input
                  v-model.trim="config['hua-wei-obs'].bucket"
                  placeholder="请输入存储桶名称"
                />
              </el-form-item>
              <el-form-item label="区域">
                <el-input
                  v-model.trim="config['hua-wei-obs'].endpoint"
                  placeholder="请输入区域"
                />
              </el-form-item>
              <el-form-item label="accessKey">
                <el-input
                  v-model.trim="config['hua-wei-obs']['access-key']"
                  placeholder="请输入accessKey"
                />
              </el-form-item>
              <el-form-item label="secretKey">
                <el-input
                  v-model.trim="config['hua-wei-obs']['secret-key']"
                  placeholder="请输入secretKey"
                />
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['oss-type'] === 'cloudflare-r2'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">Cloudflare R2上传配置</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="路径">
                <el-input
                  v-model.trim="config['cloudflare-r2'].path"
                  placeholder="请输入路径"
                />
              </el-form-item>
              <el-form-item label="存储桶名称">
                <el-input
                  v-model.trim="config['cloudflare-r2'].bucket"
                  placeholder="请输入存储桶名称"
                />
              </el-form-item>
              <el-form-item label="Base URL">
                <el-input
                  v-model.trim="config['cloudflare-r2']['base-url']"
                  placeholder="请输入Base URL"
                />
              </el-form-item>
              <el-form-item label="Account ID">
                <el-input
                  v-model.trim="config['cloudflare-r2']['account-id']"
                  placeholder="请输入secretKey"
                />
              </el-form-item>
              <el-form-item label="Access Key ID">
                <el-input
                  v-model.trim="config['cloudflare-r2']['access-key-id']"
                  placeholder="请输入secretKey"
                />
              </el-form-item>
              <el-form-item label="Secret Access Key">
                <el-input
                  v-model.trim="config['cloudflare-r2']['secret-access-key']"
                  placeholder="请输入secretKey"
                />
              </el-form-item>
            </div>
          </section>
          <section
            v-if="config.system['oss-type'] === 'minio'"
            class="config-section"
          >
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">MinIO上传配置</h3>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="Endpoint">
                <el-input
                  v-model.trim="config.minio.endpoint"
                  placeholder="请输入Endpoint，如 127.0.0.1:9000"
                />
              </el-form-item>
              <el-form-item label="Access Key ID">
                <el-input
                  v-model.trim="config.minio['access-key-id']"
                  placeholder="请输入Access Key ID"
                />
              </el-form-item>
              <el-form-item label="Access Key Secret">
                <el-input
                  v-model.trim="config.minio['access-key-secret']"
                  placeholder="请输入Access Key Secret"
                />
              </el-form-item>
              <el-form-item label="存储桶名称">
                <el-input
                  v-model.trim="config.minio['bucket-name']"
                  placeholder="请输入存储桶名称"
                />
              </el-form-item>
              <el-form-item label="访问域名">
                <el-input
                  v-model.trim="config.minio['bucket-url']"
                  placeholder="请输入访问域名"
                />
              </el-form-item>
              <el-form-item label="Base Path">
                <el-input
                  v-model.trim="config.minio['base-path']"
                  placeholder="请输入Base Path"
                />
              </el-form-item>
              <el-form-item label="开启SSL">
                <el-switch v-model="config.minio['use-ssl']" />
              </el-form-item>
            </div>
          </section>
        </template>

        <!--  Mongo 数据库  -->
        <template v-if="activeKey === 'mongo' && config.system['use-mongo']">
          <section class="config-section">
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">Mongo 连接</h3>
                <p class="section-desc">MongoDB 连接与连接池参数</p>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item label="collection name(表名,一般不写)">
                <el-input
                  v-model.trim="config.mongo.coll"
                  placeholder="请输入collection name"
                />
              </el-form-item>
              <el-form-item label="mongodb 选项">
                <el-input
                  v-model.trim="config.mongo.options"
                  placeholder="请输入mongodb 选项"
                />
              </el-form-item>
              <el-form-item label="database name(数据库名)">
                <el-input
                  v-model.trim="config.mongo.database"
                  placeholder="请输入数据库名"
                />
              </el-form-item>
              <el-form-item label="用户名">
                <el-input
                  v-model.trim="config.mongo.username"
                  placeholder="请输入用户名"
                />
              </el-form-item>
              <el-form-item label="密码">
                <el-input
                  v-model.trim="config.mongo.password"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item label="是否开启zap日志">
                <el-switch v-model="config.mongo['is-zap']" />
              </el-form-item>
              <el-form-item label="最小连接池">
                <el-input-number
                  v-model="config.mongo['min-pool-size']"
                  min="0"
                />
              </el-form-item>
              <el-form-item label="最大连接池">
                <el-input-number
                  v-model="config.mongo['max-pool-size']"
                  min="100"
                />
              </el-form-item>
              <el-form-item label="socket超时时间">
                <el-input-number
                  v-model="config.mongo['socket-timeout-ms']"
                  min="0"
                />
              </el-form-item>
              <el-form-item label="连接超时时间">
                <el-input-number
                  v-model="config.mongo['connect-timeout-ms']"
                  min="0"
                />
              </el-form-item>
            </div>
          </section>
          <section class="config-section">
            <div class="section-header">
              <span class="section-bar"></span>
              <div>
                <h3 class="section-title">节点配置</h3>
                <p class="section-desc">MongoDB 节点列表，可添加多个节点</p>
              </div>
            </div>
            <div class="section-grid">
              <el-form-item
                v-for="(item, k) in config.mongo.hosts"
                :key="k"
                :label="`节点 ${k + 1}`"
              >
                <div class="w-full max-w-[440px]">
                  <el-form-item
                    v-for="(_, k2) in item"
                    :key="k + k2"
                    :label="k2"
                    label-width="60"
                    label-position="left"
                  >
                    <el-input
                      v-model.trim="item[k2]"
                      :placeholder="k2 === 'host' ? '请输入地址' : '请输入端口'"
                    />
                  </el-form-item>
                  <el-button
                    v-if="k > 0"
                    type="danger"
                    size="small"
                    plain
                    :icon="Minus"
                    @click="removeNode(k)"
                  />
                </div>
              </el-form-item>
              <el-form-item label="添加节点">
                <el-button
                  type="primary"
                  size="small"
                  plain
                  :icon="Plus"
                  @click="addNode"
                />
              </el-form-item>
            </div>
          </section>
        </template>

        <!--  自动化代码  -->
        <section v-if="activeKey === 'autocode'" class="config-section">
          <div class="section-header">
            <span class="section-bar"></span>
            <div>
              <h3 class="section-title">代码生成路径</h3>
              <p class="section-desc">自动化代码生成相关的目录配置</p>
            </div>
          </div>
          <div class="section-grid">
            <el-form-item label="root(项目根路径)">
              <el-input v-model="config.autocode.root" disabled />
            </el-form-item>
            <el-form-item label="Server(后端代码地址)">
              <el-input
                v-model.trim="config.autocode['server']"
                placeholder="请输入后端代码地址"
              />
            </el-form-item>
            <el-form-item label="Web(前端文件夹地址)">
              <el-input
                v-model.trim="config.autocode.web"
                placeholder="请输入前端文件夹地址"
              />
            </el-form-item>
            <el-form-item label="Module(后端 Go Module)">
              <el-input
                v-model.trim="config.autocode.module"
                placeholder="请输入后端 Go Module"
              />
            </el-form-item>
            <el-form-item label="AI服务路径">
              <el-input
                v-model.trim="config.autocode['ai-path']"
                placeholder="请输入AI服务路径"
              />
            </el-form-item>
          </div>
        </section>
      </el-form>
    </div>
  </div>
</template>

<script setup>
  import { getSystemConfig, reloadSystem, setSystemConfig } from '@/api/system'
  import { ref, watch } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { Minus, Plus } from '@element-plus/icons-vue'
  import { emailTest } from '@/api/email'
  import { CreateUUID } from '@/utils/format'

  defineOptions({
    name: 'Config'
  })

  const activeKey = ref('basic')
  const config = ref({
    system: {
      'iplimit-count': 0,
      'iplimit-time': 0
    },
    jwt: {},
    mysql: {},
    mssql: {},
    sqlite: {},
    pgsql: {},
    oracle: {},
    autocode: {},
    redis: {},
    mongo: {
      coll: '',
      options: '',
      database: '',
      username: '',
      password: '',
      'min-pool-size': '',
      'max-pool-size': '',
      'socket-timeout-ms': '',
      'connect-timeout-ms': '',
      'is-zap': false,
      hosts: [
        {
          host: '',
          port: ''
        }
      ]
    },
    qiniu: {},
    'tencent-cos': {},
    'aliyun-oss': {},
    'hua-wei-obs': {},
    'cloudflare-r2': {},
    minio: {},
    zap: {},
    local: {},
    email: {},
    timer: {
      detail: {}
    }
  })

  // Redis / Mongo 开关关闭时，左侧菜单项消失，当前页回退到基础设置，避免空白
  watch(
    () => config.value.system['use-redis'],
    (v) => {
      if (!v && activeKey.value === 'redis') activeKey.value = 'basic'
    }
  )
  watch(
    () => config.value.system['use-mongo'],
    (v) => {
      if (!v && activeKey.value === 'mongo') activeKey.value = 'basic'
    }
  )

  const initForm = async () => {
    const res = await getSystemConfig()
    if (res.code === 0) {
      config.value = res.data.config
    }
  }
  initForm()
  const reload = () => {
    ElMessageBox.confirm('确定要重载服务?', '警告', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
      .then(async () => {
        const res = await reloadSystem()
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '操作成功'
          })
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '取消重启'
        })
      })
  }

  const update = async () => {
    const res = await setSystemConfig({ config: config.value })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '配置文件设置成功'
      })
      await initForm()
    }
  }

  const email = async () => {
    const res = await emailTest()
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '邮件发送成功'
      })
      await initForm()
    } else {
      ElMessage({
        type: 'error',
        message: '邮件发送失败'
      })
    }
  }

  const getUUID = () => {
    config.value.jwt['signing-key'] = CreateUUID()
  }

  const addNode = () => {
    config.value.mongo.hosts.push({
      host: '',
      port: ''
    })
  }

  const removeNode = (index) => {
    config.value.mongo.hosts.splice(index, 1)
  }
</script>

<style lang="scss" scoped>
  .config-layout {
    @apply flex flex-col gap-5 lg:flex-row lg:gap-0;

    // 桌面端固定整体高度：切换左侧菜单时页面高度不再跳动，内容超出时右栏内部滚动
    @media (min-width: 1024px) {
      height: calc(100vh - 160px);
      min-height: 480px;
    }
  }

  .config-menu {
    @apply flex w-full flex-none flex-wrap gap-1 border-0 border-b border-solid border-border pb-3 lg:w-48 lg:flex-col lg:gap-0.5 lg:overflow-y-auto lg:border-b-0 lg:border-r lg:pb-0 lg:pr-4;
  }

  .menu-group {
    @apply contents lg:block;
  }

  .menu-group-title {
    @apply hidden px-3 pb-1 pt-3 text-xs text-muted-foreground lg:block;
  }

  .menu-item {
    @apply cursor-pointer select-none rounded-lg px-3 py-1.5 text-[13px] text-base-text transition-colors hover:bg-[rgb(var(--primary-color)/0.08)] lg:py-2;

    &.on {
      @apply bg-[rgb(var(--primary-color)/0.12)] font-medium text-active;
    }
  }

  .config-content {
    @apply min-w-0 flex-1 lg:overflow-y-auto lg:pl-5 lg:pr-1;
  }

  .config-form {
    :deep(.el-form-item__label) {
      @apply font-medium text-base-text;
    }
  }

  .config-section {
    @apply mb-6 last:mb-2;
  }

  .section-header {
    @apply mb-4 flex items-start gap-2.5;
  }

  .section-bar {
    @apply mt-1 h-4 w-1 rounded-full bg-primary;
  }

  .section-title {
    @apply m-0 text-[15px] font-semibold leading-5 text-base-text;
  }

  .section-desc {
    @apply m-0 mt-1 text-xs text-muted-foreground;
  }

  .section-grid {
    @apply grid grid-cols-1;

    :deep(.el-form-item__content > .el-input),
    :deep(.el-form-item__content > .el-select) {
      width: min(440px, 100%);
    }
  }

  .action-bar {
    // 零高度悬浮：吸附右上角但不占文档流高度，不会把标题顶下去；
    // 滚动时按钮始终钉在内容区右上角，内容从其下方穿过
    @apply sticky top-0 z-10 flex h-0 justify-end gap-2 overflow-visible;
  }
</style>
