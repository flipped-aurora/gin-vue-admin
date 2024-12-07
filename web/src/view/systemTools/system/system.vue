<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="240px">
      <!--  System start  -->
      <el-tabs v-model="activeNames">
        <el-tab-pane label="系统配置" name="1" class="mt-3.5">
          <el-form-item label="端口值">
            <el-input-number
              v-model="config.system.addr"
              placeholder="请输入端口值"
            />
          </el-form-item>
          <el-form-item label="数据库类型">
            <el-select v-model="config.system['db-type']" class="w-full">
              <el-option value="mysql" />
              <el-option value="pgsql" />
              <el-option value="mssql" />
              <el-option value="sqlite" />
              <el-option value="oracle" />
            </el-select>
          </el-form-item>
          <el-form-item label="Oss类型">
            <el-select v-model="config.system['oss-type']" class="w-full">
              <el-option value="local">本地</el-option>
              <el-option value="qiniu">七牛</el-option>
              <el-option value="tencent-cos">腾讯云COS</el-option>
              <el-option value="aliyun-oss">阿里云OSS</el-option>
              <el-option value="huawei-obs">华为云OBS</el-option>
              <el-option value="cloudflare-r2">cloudflare R2</el-option>
            </el-select>
          </el-form-item>
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
        </el-tab-pane>
        <el-tab-pane label="jwt签名" name="2" class="mt-3.5">
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
        </el-tab-pane>
        <el-tab-pane label="Zap日志配置" name="3" class="mt-3.5">
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
          <el-form-item label="编码级">
            <el-select v-model="config.zap['encode-level']" class="w-6/12">
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
          <el-form-item label="日志留存时间(默认以天为单位)">
            <el-input-number v-model="config.zap['retention-day']" />
          </el-form-item>
          <el-form-item label="显示行">
            <el-switch v-model="config.zap['show-line']" />
          </el-form-item>
          <el-form-item label="输出控制台">
            <el-switch v-model="config.zap['log-in-console']" />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
          label="Redis"
          name="4"
          class="mt-3.5"
          v-if="config.system['use-redis']"
        >
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
        </el-tab-pane>
        <el-tab-pane label="邮箱配置" name="5" class="mt-3.5">
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
          <el-form-item label="是否为ssl">
            <el-switch v-model="config.email['is-ssl']" />
          </el-form-item>
          <el-form-item label="secret">
            <el-input
              v-model.trim="config.email.secret"
              placeholder="请输入secret"
            />
          </el-form-item>
          <el-form-item label="测试邮件">
            <el-button @click="email">测试邮件</el-button>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
          label="Mongo 数据库配置"
          name="14"
          class="mt-3.5"
          v-if="config.system['use-mongo']"
        >
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
          <el-form-item label="最小连接池">
            <el-input-number v-model="config.mongo['min-pool-size']" min="0" />
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
              v-model="config.mongo['socket-timeout-ms']"
              min="0"
            />
          </el-form-item>
          <el-form-item label="是否开启zap日志">
            <el-switch v-model="config.mongo['is-zap']" />
          </el-form-item>
          <el-form-item
            v-for="(item, k) in config.mongo.hosts"
            :key="k"
            :label="`节点 ${k + 1}`"
          >
            <div v-for="(_, k2) in item" :key="k2">
              <el-form-item :key="k + k2" :label="k2" label-width="60">
                <el-input
                  v-model.trim="item[k2]"
                  :placeholder="k2 === 'host' ? '请输入地址' : '请输入端口'"
                />
              </el-form-item>
            </div>
            <el-form-item v-if="k > 0">
              <el-button
                type="danger"
                size="small"
                plain
                :icon="Minus"
                @click="removeNode(k)"
                class="ml-3"
              />
            </el-form-item>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              size="small"
              plain
              :icon="Plus"
              @click="addNode"
            />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane label="验证码配置" name="7" class="mt-3.5">
          <el-form-item label="字符长度">
            <el-input-number
              v-model="config.captcha['key-long']"
              :min="4"
              :max="6"
            />
          </el-form-item>
          <el-form-item label="图片宽度">
            <el-input-number v-model.number="config.captcha['img-width']" />
          </el-form-item>
          <el-form-item label="图片高度">
            <el-input-number v-model.number="config.captcha['img-height']" />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane label="数据库配置" name="9" class="mt-3.5">
          <template v-if="config.system['db-type'] === 'mysql'">
            <el-form-item label="">
              <h3>MySQL</h3>
            </el-form-item>
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
            <el-form-item label="写入日志">
              <el-switch v-model="config.mysql['log-zap']" />
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
          </template>
          <template v-if="config.system['db-type'] === 'pgsql'">
            <el-form-item label="">
              <h3>PostgreSQL</h3>
            </el-form-item>
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
            <el-form-item label="写入日志">
              <el-switch v-model="config.pgsql['log-zap']" />
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
          </template>
          <template v-if="config.system['db-type'] === 'mssql'">
            <el-form-item label="">
              <h3>MsSQL</h3>
            </el-form-item>
            <el-form-item label="用户名">
              <el-input
                v-model.trim="config.mssql.username"
                placeholder="请输入用户名"
              />
            </el-form-item>
            <el-form-item label.trim="密码">
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
            <el-form-item label="写入日志">
              <el-switch v-model="config.mssql['log-zap']" />
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
          </template>
          <template v-if="config.system['db-type'] === 'sqlite'">
            <el-form-item label="">
              <h3>sqlite</h3>
            </el-form-item>
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
            <el-form-item label="写入日志">
              <el-switch v-model="config.sqlite['log-zap']" />
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
          </template>
          <template v-if="config.system['db-type'] === 'oracle'">
            <el-form-item label="">
              <h3>oracle</h3>
            </el-form-item>
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
            <el-form-item label="写入日志">
              <el-switch v-model="config.oracle['log-zap']" />
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
          </template>
        </el-tab-pane>
        <el-tab-pane label="oss配置" name="10" class="mt-3.5">
          <template v-if="config.system['oss-type'] === 'local'">
            <h2>本地配置</h2>
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
          </template>
          <template v-if="config.system['oss-type'] === 'qiniu'">
            <h2>七牛上传配置</h2>
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
            <el-form-item label="是否使用https">
              <el-switch v-model="config.qiniu['use-https']">开启</el-switch>
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
            <el-form-item label="上传是否使用CDN上传加速">
              <el-switch v-model="config.qiniu['use-cdn-domains']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'tencent-cos'">
            <h2>腾讯云COS上传配置</h2>
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
          </template>
          <template v-if="config.system['oss-type'] === 'aliyun-oss'">
            <h2>阿里云OSS上传配置</h2>
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
          </template>
          <template v-if="config.system['oss-type'] === 'huawei-obs'">
            <h2>华为云OBS上传配置</h2>
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
          </template>
          <template v-if="config.system['oss-type'] === 'cloudflare-r2'">
            <h2>Cloudflare R2上传配置</h2>
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
          </template>
        </el-tab-pane>
        <el-tab-pane label="Excel上传配置" name="11" class="mt-3.5">
          <el-form-item label="合成目标地址">
            <el-input
              v-model.trim="config.excel.dir"
              placeholder="请输入合成目标地址"
            />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane label="自动化代码配置" name="12" class="mt-3.5">
          <el-form-item label="是否自动重启(linux)">
            <el-switch v-model="config.autocode['transfer-restart']" />
          </el-form-item>
          <el-form-item label="root(项目根路径)">
            <el-input v-model="config.autocode.root" disabled />
          </el-form-item>
          <el-form-item label="Server(后端代码地址)">
            <el-input
              v-model.trim="config.autocode['server']"
              placeholder="请输入后端代码地址"
            />
          </el-form-item>
          <el-form-item label="SApi(后端api文件夹地址)">
            <el-input
              v-model.trim="config.autocode['server-api']"
              placeholder="请输入后端api文件夹地址"
            />
          </el-form-item>
          <el-form-item label="SInitialize(后端Initialize文件夹)">
            <el-input
              v-model.trim="config.autocode['server-initialize']"
              placeholder="请输入后端Initialize文件夹"
            />
          </el-form-item>
          <el-form-item label="SModel(后端Model文件地址)">
            <el-input
              v-model.trim="config.autocode['server-model']"
              placeholder="请输入后端Model文件地址"
            />
          </el-form-item>
          <el-form-item label="SRequest(后端Request文件夹地址)">
            <el-input
              v-model.trim="config.autocode['server-request']"
              placeholder="请输入后端Request文件夹地址"
            />
          </el-form-item>
          <el-form-item label="SRouter(后端Router文件夹地址)">
            <el-input
              v-model.trim="config.autocode['server-router']"
              placeholder="请输入后端Router文件夹地址"
            />
          </el-form-item>
          <el-form-item label="SService(后端Service文件夹地址)">
            <el-input
              v-model.trim="config.autocode['server-service']"
              placeholder="请输入后端Service文件夹地址"
            />
          </el-form-item>
          <el-form-item label="Web(前端文件夹地址)">
            <el-input
              v-model.trim="config.autocode.web"
              placeholder="请输入前端文件夹地址"
            />
          </el-form-item>
          <el-form-item label="WApi(后端WApi文件夹地址)">
            <el-input
              v-model.trim="config.autocode['web-api']"
              placeholder="请输入后端WApi文件夹地址"
            />
          </el-form-item>
          <el-form-item label="WForm(后端WForm文件夹地址)">
            <el-input
              v-model.trim="config.autocode['web-form']"
              placeholder="请输入后端WForm文件夹地址"
            />
          </el-form-item>
          <el-form-item label="WTable(后端WTable文件夹地址)">
            <el-input
              v-model.trim="config.autocode['web-table']"
              placeholder="请输入后端WTable文件夹地址"
            />
          </el-form-item>
        </el-tab-pane>
      </el-tabs>
    </el-form>
    <div class="mt-4">
      <el-button type="primary" @click="update">立即更新 </el-button>
      <el-button type="primary" @click="reload">重启服务 </el-button>
    </div>
  </div>
</template>

<script setup>
  import { getSystemConfig, reloadSystem, setSystemConfig } from '@/api/system'
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { Minus, Plus } from '@element-plus/icons-vue'
  import { emailTest } from '@/api/email'
  import { CreateUUID } from '@/utils/format'

  defineOptions({
    name: 'Config'
  })

  const activeNames = ref('1')
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
    excel: {},
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
    captcha: {},
    zap: {},
    local: {},
    email: {},
    timer: {
      detail: {}
    }
  })

  const initForm = async () => {
    const res = await getSystemConfig()
    if (res.code === 0) {
      config.value = res.data.config
    }
  }
  initForm()
  const reload = () => {
    ElMessageBox.confirm('确定要重启服务?', '警告', {
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
  .system {
    @apply bg-white p-9 rounded dark:bg-slate-900;
    h2 {
      @apply p-2.5 my-2.5 text-lg shadow;
    }
  }
</style>
