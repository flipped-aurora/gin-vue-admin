<template>
  <div class="system">
    <el-form
        ref="form"
        :model="config"
        label-width="320px"
    >
      <!--  System start  -->
      <el-tabs v-model="activeNames">
        <el-tab-pane
            :label="t('view.systemTools.system.systemConfig')"
            name="1"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.portValue')">
            <el-input-number v-model.number="config.system.addr" placeholder="请输入端口值"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.dbType')">
            <el-select
                v-model="config.system['db-type']"
                class="w-full"
            >
              <el-option value="mysql"/>
              <el-option value="pgsql"/>
              <el-option value="mssql"/>
              <el-option value="sqlite"/>
              <el-option value="oracle"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.ossType')">
            <el-select
                v-model="config.system['oss-type']"
                class="w-full"
            >
              <el-option value="local">本地</el-option>
              <el-option value="qiniu">七牛</el-option>
              <el-option value="tencent-cos">腾讯云COS</el-option>
              <el-option value="aliyun-oss">阿里云OSS</el-option>
              <el-option value="huawei-obs">华为云OBS</el-option>
              <el-option value="cloudflare-r2">cloudflare R2</el-option>
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.blockMultiSignOn')">
            <el-switch v-model="config.system['use-multipoint']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.enableRedis')">
            <el-switch v-model="config.system['use-redis']"/>
          </el-form-item>
          <el-form-item label="开启Mongo">
            <el-switch v-model="config.system['use-mongo']"/>
          </el-form-item>
          <el-form-item label="严格角色模式">
            <el-switch v-model="config.system['use-strict-auth']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.ipLimitCount')">
            <el-input-number v-model.number="config.system['iplimit-count']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.ipLimitTime')">
            <el-input-number v-model.number="config.system['iplimit-time']"/>
          </el-form-item>
          <el-tooltip
              :content="t('view.systemTools.system.globalRoutePrefixNote')"
              placement="top-start"
          >
            <el-form-item :label="t('view.systemTools.system.globalRoutePrefix')">
              <el-input v-model.trim="config.system['router-prefix']" placeholder="请输入全局路由前缀"/>
            </el-form-item>
          </el-tooltip>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.jwtSignature')"
            name="2"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.jwtSignature')">
            <el-input v-model.trim="config.jwt['signing-key']" placeholder="请输入jwt签名">
              <template #append>
                <el-button @click="getUUID">生成</el-button>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.expirartionSec')">
            <el-input v-model.trim="config.jwt['expires-time']" placeholder="请输入有效期"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.bufferPeriodSec')">
            <el-input v-model.trim="config.jwt['buffer-time']" placeholder="请输入缓冲期"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.issuer')">
            <el-input v-model.trim="config.jwt.issuer" placeholder="请输入签发者"/>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.zapLogConfig')"
            name="3"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.level')">
            <el-select v-model="config.zap.level">
              <el-option value="off" label="关闭"/>
              <el-option value="fatal" label="致命"/>
              <el-option value="error" label="错误"/>
              <el-option value="warn" label="警告"/>
              <el-option value="info" label="信息"/>
              <el-option value="debug" label="调试"/>
              <el-option value="trace" label="跟踪"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.output')">
            <el-select v-model="config.zap.format">
              <el-option value="console" label="console"/>
              <el-option value="json" label="json"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.logPrefix')">
            <el-input v-model.trim="config.zap.prefix" placeholder="请输入日志前缀"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.logFolder')">
            <el-input v-model.trim="config.zap.director" placeholder="请输入日志文件夹"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.encodeLevel')">
            <el-select v-model="config.zap['encode-level']" class="w-6/12">
              <el-option value="LowercaseLevelEncoder" label="LowercaseLevelEncoder"/>
              <el-option value="LowercaseColorLevelEncoder" label="LowercaseColorLevelEncoder"/>
              <el-option value="CapitalLevelEncoder" label="CapitalLevelEncoder"/>
              <el-option value="CapitalColorLevelEncoder" label="CapitalColorLevelEncoder"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.stackName')">
            <el-input v-model.trim="config.zap['stacktrace-key']" placeholder="请输入栈名"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.logRetentionTime')">
            <el-input-number v-model="config.zap['retention-day']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.showLine')">
            <el-switch v-model="config.zap['show-line']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.outputConsole')">
            <el-switch v-model="config.zap['log-in-console']"/>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            label="Redis"
            name="4"
            class="mt-3.5"
            v-if="config.system['use-redis']"
        >
          <el-form-item :label="t('view.systemTools.system.storehouse')">
            <el-input-number v-model="config.redis.db" min="0" max="16"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.address')">
            <el-input v-model.trim="config.redis.addr" placeholder="请输入地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.password')">
            <el-input v-model.trim="config.redis.password" placeholder="请输入密码"/>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            label="邮箱配置"
            name="5"
            class="mt-3.5"
        >
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
            <el-input v-model.trim="config.email.from" placeholder="请输入发送者邮箱" />
          </el-form-item>
          <el-form-item label="host">
            <el-input v-model.trim="config.email.host" placeholder="请输入host" />
          </el-form-item>
          <el-form-item label="是否为ssl">
            <el-switch v-model="config.email['is-ssl']" />
          </el-form-item>
          <el-form-item label="secret">
            <el-input v-model.trim="config.email.secret" placeholder="请输入secret"/>
          </el-form-item>
          <el-form-item label="测试邮件">
            <el-button @click="email">测试邮件</el-button>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.mongoDbConfig')"
            name="14"
            class="mt-3.5"
            v-if="config.system['use-mongo']"
        >
          <el-form-item :label="t('view.systemTools.system.collectionName')">
            <el-input v-model.trim="config.mongo.coll" placeholder="请输入collection name"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.options')">
            <el-input v-model.trim="config.mongo.options" placeholder="请输入mongodb 选项"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.dbName')">
            <el-input v-model.trim="config.mongo.database" placeholder="请输入数据库名"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.userName')">
            <el-input v-model.trim="config.mongo.username" placeholder="请输入用户名"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.password')">
            <el-input v-model.trim="config.mongo.password" placeholder="请输入密码"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.minimumConnectionPool')">
            <el-input-number v-model="config.mongo['min-pool-size']" min="0"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.maximumConnectionPool')">
            <el-input-number v-model="config.mongo['max-pool-size']" min="100"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.socketTimeout')">
            <el-input-number v-model="config.mongo['socket-timeout-ms']" min="0"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.connectionTimeout')">
            <el-input-number v-model="config.mongo['socket-timeout-ms']" min="0"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.enableZapLog')">
            <el-switch v-model="config.mongo['is-zap']"/>
          </el-form-item>
          <template v-for="(item,k) in config.mongo.hosts">
            <el-form-item :label="`节点 ${k+1}`">
              <div v-for="(_,k2) in item" :key="k2">
                <el-form-item :key="k+k2" :label="k2" label-width="60">
                  <el-input v-model.trim="item[k2]" :placeholder="k2 === 'host' ? '请输入地址' : '请输入端口'"/>
                </el-form-item>
              </div>
              <el-form-item v-if="k > 0">
                <el-button type="danger" size="small" plain :icon="Minus" @click="removeNode(k)" class="ml-3"/>
              </el-form-item>
            </el-form-item>
          </template>
          <el-form-item>
            <el-button type="primary" size="small" plain :icon="Plus" @click="addNode"/>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.verCodeConfig')"
            name="7"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.charLength')">
            <el-input-number v-model="config.captcha['key-long']" min="4" max="6"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.imageWidth')">
            <el-input-number v-model.number="config.captcha['img-width']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.imageHeight')">
            <el-input-number v-model.number="config.captcha['img-height']"/>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.dbConfig')"
            name="9"
            class="mt-3.5"
        >
          <template v-if="config.system['db-type'] === 'mysql'">
            <el-form-item label="">
              <h3>MySQL</h3>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.userName')">
              <el-input v-model.trim="config.mysql.username" placeholder="请输入用户名"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.password')">
              <el-input v-model.trim="config.mysql.password" placeholder="请输入密码"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.address')">
              <el-input v-model.trim="config.mysql.path" placeholder="请输入地址"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.dbName')">
              <el-input v-model.trim="config.mysql['db-name']" placeholder="请输入数据库名称"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.prefix')">
              <el-input v-model.trim="config.mysql['prefix']" placeholder="默认为空"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.pluralTable')">
              <el-switch v-model="config.mysql['singular']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.engine')">
              <el-input v-model.trim="config.mysql['engine']" placeholder="默认为InnoDB"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.maxIdleConns')">
              <el-input-number v-model="config.mysql['max-idle-conns']" :min="1"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.maxOpenConns')">
              <el-input-number v-model="config.mysql['max-open-conns']" :min="1"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.enableZapLog')">
              <el-switch v-model="config.mysql['log-zap']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.logMode')">
              <el-select v-model="config.mysql['log-mode']">
                <el-option value="off" label="关闭"/>
                <el-option value="fatal" label="致命"/>
                <el-option value="error" label="错误"/>
                <el-option value="warn" label="警告"/>
                <el-option value="info" label="信息"/>
                <el-option value="debug" label="调试"/>
                <el-option value="trace" label="跟踪"/>
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'pgsql'">
            <el-form-item label="">
              <h3>PostgreSQL</h3>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.userName')">
              <el-input v-model="config.pgsql.username" placeholder="请输入用户名"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.password')">
              <el-input v-model="config.pgsql.password" placeholder="请输入密码"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.address')">
              <el-input v-model.trim="config.pgsql.path" placeholder="请输入地址"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.dbName')">
              <el-input v-model.trim="config.pgsql['db-name']" placeholder="请输入数据库"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.prefix')">
              <el-input v-model.trim="config.pgsql['prefix']" placeholder="请输入前缀"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.pluralTable')">
              <el-switch v-model="config.pgsql['singular']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.engine')">
              <el-input v-model.trim="config.pgsql['engine']" placeholder="请输入引擎"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.maxIdleConns')">
              <el-input-number v-model="config.pgsql['max-idle-conns']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.maxOpenConns')">
              <el-input-number v-model="config.pgsql['max-open-conns']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.enableZapLog')">
              <el-switch v-model="config.pgsql['log-zap']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.logMode')">
              <el-select v-model="config.pgsql['log-mode']">
                <el-option value="off" label="关闭"/>
                <el-option value="fatal" label="致命"/>
                <el-option value="error" label="错误"/>
                <el-option value="warn" label="警告"/>
                <el-option value="info" label="信息"/>
                <el-option value="debug" label="调试"/>
                <el-option value="trace" label="跟踪"/>
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'mssql'">
            <el-form-item label="">
              <h3>MsSQL</h3>
            </el-form-item>
            <el-form-item label="用户名">
              <el-input v-model.trim="config.mssql.username" placeholder="请输入用户名"/>
            </el-form-item>
            <el-form-item label.trim="密码">
              <el-input v-model.trim="config.mssql.password" placeholder="请输入密码"/>
            </el-form-item>
            <el-form-item label="地址">
              <el-input v-model.trim="config.mssql.path" placeholder="请输入地址"/>
            </el-form-item>
            <el-form-item label="端口">
              <el-input v-model.trim="config.mssql.port" placeholder="请输入端口"/>
            </el-form-item>
            <el-form-item label="数据库">
              <el-input v-model.trim="config.mssql['db-name']" placeholder="请输入数据库"/>
            </el-form-item>
            <el-form-item label="前缀">
              <el-input v-model.trim="config.mssql['prefix']" placeholder="请输入前缀"/>
            </el-form-item>
            <el-form-item label="复数表">
              <el-switch v-model="config.mssql['singular']"/>
            </el-form-item>
            <el-form-item label="引擎">
              <el-input v-model.trim="config.mssql['engine']" placeholder="请输入引擎"/>
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input-number v-model="config.mssql['max-idle-conns']"/>
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input-number v-model="config.mssql['max-open-conns']"/>
            </el-form-item>
            <el-form-item label="写入日志">
              <el-switch v-model="config.mssql['log-zap']"/>
            </el-form-item>
            <el-form-item label="日志模式">
              <el-select v-model="config.mssql['log-mode']">
                <el-option value="off" label="关闭"/>
                <el-option value="fatal" label="致命"/>
                <el-option value="error" label="错误"/>
                <el-option value="warn" label="警告"/>
                <el-option value="info" label="信息"/>
                <el-option value="debug" label="调试"/>
                <el-option value="trace" label="跟踪"/>
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'sqlite'">
            <el-form-item label="">
              <h3>sqlite</h3>
            </el-form-item>
            <el-form-item label="用户名">
              <el-input v-model.trim="config.sqlite.username" placeholder="请输入用户名"/>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model.trim="config.sqlite.password" placeholder="请输入密码"/>
            </el-form-item>
            <el-form-item label="地址">
              <el-input v-model.trim="config.sqlite.path" placeholder="请输入地址"/>
            </el-form-item>
            <el-form-item label="端口">
              <el-input v-model.trim="config.sqlite.port" placeholder="请输入端口"/>
            </el-form-item>
            <el-form-item label="数据库">
              <el-input v-model.trim="config.sqlite['db-name']" placeholder="请输入数据库"/>
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input-number v-model="config.sqlite['max-idle-conns']"/>
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input-number v-model="config.sqlite['max-open-conns']"/>
            </el-form-item>
            <el-form-item label="写入日志">
              <el-switch v-model="config.sqlite['log-zap']"/>
            </el-form-item>
            <el-form-item label="日志模式">
              <el-select v-model="config.sqlite['log-mode']">
                <el-option value="off" label="关闭"/>
                <el-option value="fatal" label="致命"/>
                <el-option value="error" label="错误"/>
                <el-option value="warn" label="警告"/>
                <el-option value="info" label="信息"/>
                <el-option value="debug" label="调试"/>
                <el-option value="trace" label="跟踪"/>
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'oracle'">
            <el-form-item label="">
              <h3>oracle</h3>
            </el-form-item>
            <el-form-item label="用户名">
              <el-input v-model.trim="config.oracle.username" placeholder="请输入用户名"/>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model.trim="config.oracle.password" placeholder="请输入密码"/>
            </el-form-item>
            <el-form-item label="地址">
              <el-input v-model.trim="config.oracle.path" placeholder="请输入地址"/>
            </el-form-item>
            <el-form-item label="数据库名称">
              <el-input v-model.trim="config.oracle['db-name']" placeholder="请输入数据库名称"/>
            </el-form-item>
            <el-form-item label="前缀">
              <el-input v-model.trim="config.oracle['prefix']" placeholder="默认为空"/>
            </el-form-item>
            <el-form-item label="复数表">
              <el-switch v-model="config.oracle['singular']"/>
            </el-form-item>
            <el-form-item label="引擎">
              <el-input v-model.trim="config.oracle['engine']" placeholder="默认为InnoDB"/>
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input-number v-model="config.oracle['max-idle-conns']" :min="1"/>
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input-number v-model="config.oracle['max-open-conns']" :min="1"/>
            </el-form-item>
            <el-form-item label="写入日志">
              <el-switch v-model="config.oracle['log-zap']"/>
            </el-form-item>
            <el-form-item label="日志模式">
              <el-select v-model="config.oracle['log-mode']">
                <el-option value="off" label="关闭"/>
                <el-option value="fatal" label="致命"/>
                <el-option value="error" label="错误"/>
                <el-option value="warn" label="警告"/>
                <el-option value="info" label="信息"/>
                <el-option value="debug" label="调试"/>
                <el-option value="trace" label="跟踪"/>
              </el-select>
            </el-form-item>
          </template>
        </el-tab-pane>
        <el-tab-pane
            label="oss配置"
            name="10"
            class="mt-3.5"
        >
          <template v-if="config.system['oss-type'] === 'local'">
            <h2>{{ t('view.systemTools.system.localFileConfig') }}</h2>
            <el-form-item label="本地文件访问路径">
              <el-input v-model.trim="config.local.path" placeholder="请输入本地文件访问路径"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.localFilePath')">
              <el-input v-model.trim="config.local['store-path']" placeholder="请输入本地文件存储路径"/>
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'qiniu'">
            <h2>七牛上传配置</h2>
            <el-form-item label="存储区域">
              <el-input v-model.trim="config.qiniu.zone" placeholder="请输入存储区域" />
            </el-form-item>
            <el-form-item label="空间名称">
              <el-input v-model.trim="config.qiniu.bucket" placeholder="请输入空间名称" />
            </el-form-item>
            <el-form-item label="CDN加速域名">
              <el-input v-model.trim="config.qiniu['img-path']" placeholder="请输入CDN加速域名" />
            </el-form-item>
            <el-form-item label="是否使用https">
              <el-switch v-model="config.qiniu['use-https']">开启</el-switch>
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model.trim="config.qiniu['access-key']" placeholder="请输入accessKey" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model.trim="config.qiniu['secret-key']" placeholder="请输入secretKey" />
            </el-form-item>
            <el-form-item label="上传是否使用CDN上传加速">
              <el-switch v-model="config.qiniu['use-cdn-domains']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'tencent-cos'">
            <h2>腾讯云COS上传配置</h2>
            <el-form-item label="存储桶名称">
              <el-input v-model.trim="config['tencent-cos']['bucket']" placeholder="请输入存储桶名称"/>
            </el-form-item>
            <el-form-item label="所属地域">
              <el-input v-model.trim="config['tencent-cos'].region" placeholder="请输入所属地域"/>
            </el-form-item>
            <el-form-item label="secretID">
              <el-input v-model.trim="config['tencent-cos']['secret-id']" placeholder="请输入secretID"/>
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model.trim="config['tencent-cos']['secret-key']" placeholder="请输入secretKey"/>
            </el-form-item>
            <el-form-item label="路径前缀">
              <el-input v-model.trim="config['tencent-cos']['path-prefix']" placeholder="请输入路径前缀"/>
            </el-form-item>
            <el-form-item label="访问域名">
              <el-input v-model.trim="config['tencent-cos']['base-url']" placeholder="请输入访问域名"/>
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'aliyun-oss'">
            <h2>阿里云OSS上传配置</h2>
            <el-form-item label="区域">
              <el-input v-model.trim="config['aliyun-oss'].endpoint" placeholder="请输入区域"/>
            </el-form-item>
            <el-form-item label="accessKeyId">
              <el-input v-model.trim="config['aliyun-oss']['access-key-id']" placeholder="请输入accessKeyId"/>
            </el-form-item>
            <el-form-item label="accessKeySecret">
              <el-input v-model.trim="config['aliyun-oss']['access-key-secret']" placeholder="请输入accessKeySecret"/>
            </el-form-item>
            <el-form-item label="存储桶名称">
              <el-input v-model.trim="config['aliyun-oss']['bucket-name']" placeholder="请输入存储桶名称" />
            </el-form-item>
            <el-form-item label="访问域名">
              <el-input v-model.trim="config['aliyun-oss']['bucket-url']" placeholder="请输入访问域名" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'huawei-obs'">
            <h2>华为云OBS上传配置</h2>
            <el-form-item label="路径">
              <el-input v-model.trim="config['hua-wei-obs'].path" placeholder="请输入路径" />
            </el-form-item>
            <el-form-item label="存储桶名称">
              <el-input v-model.trim="config['hua-wei-obs'].bucket" placeholder="请输入存储桶名称" />
            </el-form-item>
            <el-form-item label="区域">
              <el-input v-model.trim="config['hua-wei-obs'].endpoint" placeholder="请输入区域" />
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model.trim="config['hua-wei-obs']['access-key']" placeholder="请输入accessKey" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model.trim="config['hua-wei-obs']['secret-key']" placeholder="请输入secretKey" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'cloudflare-r2'">
            <h2>Cloudflare R2上传配置</h2>
            <el-form-item label="路径">
              <el-input v-model.trim="config['cloudflare-r2'].path" placeholder="请输入路径" />
            </el-form-item>
            <el-form-item label="存储桶名称">
              <el-input v-model.trim="config['cloudflare-r2'].bucket" placeholder="请输入存储桶名称" />
            </el-form-item>
            <el-form-item label="Base URL">
              <el-input v-model.trim="config['cloudflare-r2']['base-url']" placeholder="请输入Base URL" />
            </el-form-item>
            <el-form-item label="Account ID">
              <el-input v-model.trim="config['cloudflare-r2']['account-id']" placeholder="请输入secretKey" />
            </el-form-item>
            <el-form-item label="Access Key ID">
              <el-input v-model.trim="config['cloudflare-r2']['access-key-id']" placeholder="请输入secretKey" />
            </el-form-item>
            <el-form-item label="Secret Access Key">
              <el-input v-model.trim="config['cloudflare-r2']['secret-access-key']" placeholder="请输入secretKey" />
            </el-form-item>
          </template>

        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.excelUploadConfig')"
            name="11"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.excelDir')">
            <el-input v-model.trim="config.excel.dir" placeholder="请输入合成目标地址" />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.autoCodeConfig')"
            name="12"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.autoRestart')">
            <el-switch v-model="config.autocode['transfer-restart']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.projectRootPath')">
            <el-input
                v-model="config.autocode.root"
                disabled
            />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendCodePath')">
            <el-input v-model.trim="config.autocode['server']" placeholder="请输入后端代码地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendApiPath')">
            <el-input v-model.trim="config.autocode['server-api']" placeholder="请输入后端api文件夹地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendInitPath')">
            <el-input v-model.trim="config.autocode['server-initialize']" placeholder="请输入后端Initialize文件夹"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendModelPath')">
            <el-input v-model.trim="config.autocode['server-model']" placeholder="请输入后端Model文件地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendRequestPath')">
            <el-input v-model.trim="config.autocode['server-request']" placeholder="请输入后端Request文件夹地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendRouterPath')">
            <el-input v-model.trim="config.autocode['server-router']" placeholder="请输入后端Router文件夹地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendServicePath')">
            <el-input v-model.trim="config.autocode['server-service']" placeholder="请输入后端Service文件夹地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendCodePath')">
            <el-input v-model.trim="config.autocode.web" placeholder="请输入前端文件夹地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendApiPath')">
            <el-input v-model.trim="config.autocode['web-api']" placeholder="请输入后端WApi文件夹地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendFormPath')">
            <el-input v-model.trim="config.autocode['web-form']" placeholder="请输入后端WForm文件夹地址"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendTablePath')">
            <el-input v-model.trim="config.autocode['web-table']" placeholder="请输入后端WTable文件夹地址"/>
          </el-form-item>
        </el-tab-pane>
      </el-tabs>
    </el-form>
    <div class="mt-4">
      <el-button
          type="primary"
          @click="update"
      >{{ t('view.systemTools.system.updateNow') }}
      </el-button>
      <el-button
          type="primary"
          @click="reload"
      >{{ t('view.systemTools.system.restartService') }}
      </el-button>
    </div>
  </div>
</template>

<script setup>
import { getSystemConfig, reloadSystem, setSystemConfig } from '@/api/system'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Minus, Plus } from '@element-plus/icons-vue'
import { emailTest } from '@/api/email'
import {CreateUUID} from "@/utils/format";
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

defineOptions({
  name: 'Config',
})

const activeNames = ref('1')
const config = ref({
  system: {
    'iplimit-count': 0,
    'iplimit-time': 0,
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
        port: '',
      },
    ],
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
    detail: {},
  },
  language: {}
})

const initForm = async() => {
  const res = await getSystemConfig()
  if (res.code === 0) {
    config.value = res.data.config
  }
}
initForm()
const reload = () => {
  ElMessageBox.confirm(
      '确定要重启服务?',
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      },
  )
      .then(async() => {
        const res = await reloadSystem()
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '操作成功',
          })
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '取消重启',
        })
      })
}

const update = async() => {
  const res = await setSystemConfig({ config: config.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: t('view.systemTools.system.configSetupSuccess')
    })
    await initForm()
  }
}

const email = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: t('view.systemTools.system.emailSentSuccess')
    })
    await initForm()
  } else {
    ElMessage({
      type: 'error',
      message: t('view.systemTools.system.emailSentError')
    })
  }
}

const getUUID = () => {
  config.value.jwt['signing-key'] = CreateUUID()
}

const addNode = () => {
  config.value.mongo.hosts.push({
    host: '',
    port: '',
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
