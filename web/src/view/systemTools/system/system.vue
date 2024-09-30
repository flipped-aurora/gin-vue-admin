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
            <el-input-number v-model.number="config.system.addr" :placeholder="t('view.systemTools.system.portValue')"/>
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
              <el-option value="local">{{t('view.systemTools.system.local') }}</el-option>
              <el-option value="qiniu">{{t('view.systemTools.system.qiniu') }}</el-option>
              <el-option value="tencent-cos">{{ t('view.systemTools.system.tencentCOS') }}</el-option>
              <el-option value="aliyun-oss">{{ t('view.systemTools.system.alibabaOSS') }}</el-option>
              <el-option value="huawei-obs">{{ t('view.systemTools.system.huaweiOBS') }}</el-option>
              <el-option value="cloudflare-r2">cloudflare R2</el-option>
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.blockMultiSignOn')">
            <el-switch v-model="config.system['use-multipoint']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.enableRedis')">
            <el-switch v-model="config.system['use-redis']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.enableMongo')">
            <el-switch v-model="config.system['use-mongo']"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.strictRoleMode')">
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
              <el-input v-model.trim="config.system['router-prefix']" :placeholder="t('view.systemTools.system.globalRoutePrefix')"/>
            </el-form-item>
          </el-tooltip>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.jwtSignature')"
            name="2"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.jwtSignature')">
            <el-input v-model.trim="config.jwt['signing-key']" :placeholder="t('view.systemTools.system.jwtSignature')">
              <template #append>
                <el-button @click="getUUID">{{t('view.systemTools.autoPkg.generate')}}</el-button>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.expirartionSec')">
            <el-input v-model.trim="config.jwt['expires-time']" :placeholder="t('view.systemTools.system.expiration')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.bufferPeriodSec')">
            <el-input v-model.trim="config.jwt['buffer-time']" :placeholder="t('view.systemTools.system.bufferPeriod')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.issuer')">
            <el-input v-model.trim="config.jwt.issuer" :placeholder="t('view.systemTools.system.issuer')"/>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.zapLogConfig')"
            name="3"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.level')">
            <el-select v-model="config.zap.level">
              <el-option value="off" :label="t('components.commandMenu.close')"/>
              <el-option value="fatal" :label="t('view.systemTools.system.fatal')"/>
              <el-option value="error" :label="t('view.dashboard.error')"/>
              <el-option value="warn" :label="t('general.warning')"/>
              <el-option value="info" :label="t('view.dashboard.information')"/>
              <el-option value="debug" :label="t('view.systemTools.system.debug')"/>
              <el-option value="trace" :label="t('view.systemTools.system.trace')"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.output')">
            <el-select v-model="config.zap.format">
              <el-option value="console" label="console"/>
              <el-option value="json" label="json"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.logPrefix')">
            <el-input v-model.trim="config.zap.prefix" :placeholder="t('view.systemTools.system.logPrefix')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.logFolder')">
            <el-input v-model.trim="config.zap.director" :placeholder="t('view.systemTools.system.logFolder')"/>
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
            <el-input v-model.trim="config.zap['stacktrace-key']" :placeholder="t('view.systemTools.system.stackName')"/>
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
            <el-input v-model.trim="config.redis.addr" :placeholder="t('view.systemTools.system.address')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.password')">
            <el-input v-model.trim="config.redis.password" :placeholder="t('view.systemTools.system.password')"/>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.emailConfig')"
            name="5"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.recipientEmail')">
            <el-input
                v-model="config.email.to"
                :placeholder="t('view.systemTools.system.multipleEmailsNote')"
            />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.port')">
            <el-input-number v-model="config.email.port" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.senderEmail')">
            <el-input v-model.trim="config.email.from" :placeholder="t('view.systemTools.system.enterSenderEmail')" />
          </el-form-item>
          <el-form-item label="host">
            <el-input v-model.trim="config.email.host" :placeholder="t('view.systemTools.system.enterHost')" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.isSSL')">
            <el-switch v-model="config.email['is-ssl']" />
          </el-form-item>
          <el-form-item label="secret">
            <el-input v-model.trim="config.email.secret" :placeholder="t('view.systemTools.system.enterSecret')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.testEmail')">
            <el-button @click="email">{{ t('view.systemTools.system.testEmail') }}</el-button>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.mongoDbConfig')"
            name="14"
            class="mt-3.5"
            v-if="config.system['use-mongo']"
        >
          <el-form-item :label="t('view.systemTools.system.collectionName')">
            <el-input v-model.trim="config.mongo.coll" :placeholder="t('view.systemTools.system.collectionName')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.options')">
            <el-input v-model.trim="config.mongo.options" :placeholder="t('view.systemTools.system.mongodbOptions')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.dbName')">
            <el-input v-model.trim="config.mongo.database" :placeholder="t('view.systemTools.system.databaseName')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.userName')">
            <el-input v-model.trim="config.mongo.username" :placeholder="t('view.systemTools.system.userName')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.password')">
            <el-input v-model.trim="config.mongo.password" :placeholder="t('view.systemTools.system.typePassword')"/>
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
                  <el-input v-model.trim="item[k2]" :placeholder="k2 === 'host' ? t('view.systemTools.system.address') : t('view.systemTools.system.portValue')"/>
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
              <el-input v-model.trim="config.mysql.username" :placeholder="t('view.systemTools.system.userName')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.password')">
              <el-input v-model.trim="config.mysql.password" :placeholder="t('view.systemTools.system.typePassword')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.address')">
              <el-input v-model.trim="config.mysql.path" :placeholder="t('view.systemTools.system.address')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.dbName')">
              <el-input v-model.trim="config.mysql['db-name']" :placeholder="t('init.enterDBName')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.prefix')">
              <el-input v-model.trim="config.mysql['prefix']" :placeholder="t('view.systemTools.system.defaultEmpty')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.pluralTable')">
              <el-switch v-model="config.mysql['singular']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.engine')">
              <el-input v-model.trim="config.mysql['engine']" :placeholder="t('view.systemTools.system.defaultInnoDB')"/>
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
                <el-option value="off" :label="t('components.commandMenu.close')"/>
                <el-option value="fatal" :label="t('view.systemTools.system.fatal')"/>
                <el-option value="error" :label="t('view.dashboard.error')"/>
                <el-option value="warn" :label="t('general.warning')"/>
                <el-option value="info" :label="t('view.dashboard.information')"/>
                <el-option value="debug" :label="t('view.systemTools.system.debug')"/>
                <el-option value="trace" :label="t('view.systemTools.system.trace')"/>
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'pgsql'">
            <el-form-item label="">
              <h3>PostgreSQL</h3>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.userName')">
              <el-input v-model="config.pgsql.username" :placeholder="t('view.systemTools.system.userName')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.password')">
              <el-input v-model="config.pgsql.password" :placeholder="t('view.systemTools.system.typePassword')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.address')">
              <el-input v-model.trim="config.pgsql.path" :placeholder="t('view.systemTools.system.address')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.dbName')">
              <el-input v-model.trim="config.pgsql['db-name']" :placeholder="t('view.systemTools.system.databaseName')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.prefix')">
              <el-input v-model.trim="config.pgsql['prefix']" :placeholder="t('view.systemTools.system.enterPrefix')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.pluralTable')">
              <el-switch v-model="config.pgsql['singular']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.engine')">
              <el-input v-model.trim="config.pgsql['engine']" :placeholder="t('view.systemTools.system.enterEngine')"/>
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
                <el-option value="off" :label="t('components.commandMenu.close')"/>
                <el-option value="fatal" :label="t('view.systemTools.system.fatal')"/>
                <el-option value="error" :label="t('view.dashboard.error')"/>
                <el-option value="warn" :label="t('general.warning')"/>
                <el-option value="info" :label="t('view.dashboard.information')"/>
                <el-option value="debug" :label="t('view.systemTools.system.debug')"/>
                <el-option value="trace" :label="t('view.systemTools.system.trace')"/>
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'mssql'">
            <el-form-item label="">
              <h3>MsSQL</h3>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.userName')">
              <el-input v-model.trim="config.mssql.username" :placeholder="t('view.systemTools.system.userName')"/>
            </el-form-item>
            <el-form-item :label.trim="t('view.systemTools.system.password')">
              <el-input v-model.trim="config.mssql.password" :placeholder="t('view.systemTools.system.typePassword')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.address')">
              <el-input v-model.trim="config.mssql.path" :placeholder="t('view.systemTools.system.address')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.port')">
              <el-input v-model.trim="config.mssql.port" :placeholder="t('view.systemTools.system.portValue')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.autoCode.dbName')">
              <el-input v-model.trim="config.mssql['db-name']" :placeholder="t('view.systemTools.system.databaseName')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.prefix')">
              <el-input v-model.trim="config.mssql['prefix']" :placeholder="t('view.systemTools.system.enterPrefix')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.pluralTable')">
              <el-switch v-model="config.mssql['singular']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.engine')">
              <el-input v-model.trim="config.mssql['engine']" :placeholder="t('view.systemTools.system.enterEngine')"/>
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input-number v-model="config.mssql['max-idle-conns']"/>
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input-number v-model="config.mssql['max-open-conns']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.writeLog')">
              <el-switch v-model="config.mssql['log-zap']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.logMode')">
              <el-select v-model="config.mssql['log-mode']">
                <el-option value="off" :label="t('components.commandMenu.close')"/>
                <el-option value="fatal" :label="t('view.systemTools.system.fatal')"/>
                <el-option value="error" :label="t('view.dashboard.error')"/>
                <el-option value="warn" :label="t('general.warning')"/>
                <el-option value="info" :label="t('view.dashboard.information')"/>
                <el-option value="debug" :label="t('view.systemTools.system.debug')"/>
                <el-option value="trace" :label="t('view.systemTools.system.trace')"/>
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'sqlite'">
            <el-form-item label="">
              <h3>sqlite</h3>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.userName')">
              <el-input v-model.trim="config.sqlite.username" :placeholder="t('view.systemTools.system.userName')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.password')">
              <el-input v-model.trim="config.sqlite.password" :placeholder="t('view.systemTools.system.typePassword')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.address')">
              <el-input v-model.trim="config.sqlite.path" :placeholder="t('view.systemTools.system.address')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.port')">
              <el-input v-model.trim="config.sqlite.port" :placeholder="t('view.systemTools.system.portValue')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.autoCode.dbName')">
              <el-input v-model.trim="config.sqlite['db-name']" :placeholder="t('view.systemTools.system.databaseName')"/>
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input-number v-model="config.sqlite['max-idle-conns']"/>
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input-number v-model="config.sqlite['max-open-conns']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.writeLog')">
              <el-switch v-model="config.sqlite['log-zap']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.logMode')">
              <el-select v-model="config.sqlite['log-mode']">
                <el-option value="off" :label="t('components.commandMenu.close')"/>
                <el-option value="fatal" :label="t('view.systemTools.system.fatal')"/>
                <el-option value="error" :label="t('view.dashboard.error')"/>
                <el-option value="warn" :label="t('general.warning')"/>
                <el-option value="info" :label="t('view.dashboard.information')"/>
                <el-option value="debug" :label="t('view.systemTools.system.debug')"/>
                <el-option value="trace" :label="t('view.systemTools.system.trace')"/>
              </el-select>
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'oracle'">
            <el-form-item label="">
              <h3>oracle</h3>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.userName')">
              <el-input v-model.trim="config.oracle.username" :placeholder="t('view.systemTools.system.userName')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.password')">
              <el-input v-model.trim="config.oracle.password" :placeholder="t('view.systemTools.system.typePassword')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.address')">
              <el-input v-model.trim="config.oracle.path" :placeholder="t('view.systemTools.system.address')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.dbName')">
              <el-input v-model.trim="config.oracle['db-name']" :placeholder="t('init.enterDBName')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.prefix')">
              <el-input v-model.trim="config.oracle['prefix']" :placeholder="t('view.systemTools.system.defaultEmpty')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.pluralTable')">
              <el-switch v-model="config.oracle['singular']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.engine')">
              <el-input v-model.trim="config.oracle['engine']" :placeholder="t('view.systemTools.system.defaultInnoDB')"/>
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input-number v-model="config.oracle['max-idle-conns']" :min="1"/>
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input-number v-model="config.oracle['max-open-conns']" :min="1"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.writeLog')">
              <el-switch v-model="config.oracle['log-zap']"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.logMode')">
              <el-select v-model="config.oracle['log-mode']">
                <el-option value="off" :label="t('components.commandMenu.close')"/>
                <el-option value="fatal" :label="t('view.systemTools.system.fatal')"/>
                <el-option value="error" :label="t('view.dashboard.error')"/>
                <el-option value="warn" :label="t('general.warning')"/>
                <el-option value="info" :label="t('view.dashboard.information')"/>
                <el-option value="debug" :label="t('view.systemTools.system.debug')"/>
                <el-option value="trace" :label="t('view.systemTools.system.trace')"/>
              </el-select>
            </el-form-item>
          </template>
        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.ossConfig')"
            name="10"
            class="mt-3.5"
        >
          <template v-if="config.system['oss-type'] === 'local'">
            <h2>{{ t('view.systemTools.system.localFileConfig') }}</h2>
            <el-form-item :label="t('view.systemTools.system.localFileAccessPath')">
              <el-input v-model.trim="config.local.path" :placeholder="t('view.systemTools.system.enterLocalFileAccessPath')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.localFilePath')">
              <el-input v-model.trim="config.local['store-path']" :placeholder="t('view.systemTools.system.enterLocalFileStoragePath')"/>
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'qiniu'">
            <h2>{{ t('view.systemTools.system.qiniuUploadConfig') }}</h2>
            <el-form-item :label="t('view.systemTools.system.storageRegion')">
              <el-input v-model.trim="config.qiniu.zone" :placeholder="t('view.systemTools.system.enterStorageRegion')" />
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.spaceName')">
              <el-input v-model.trim="config.qiniu.bucket" :placeholder="t('view.systemTools.system.enterSpaceName')" />
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.cdnAcceleratedDomain')">
              <el-input v-model.trim="config.qiniu['img-path']" :placeholder="t('view.systemTools.system.enterCdnAcceleratedDomain')" />
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.useHttps')">
              <el-switch v-model="config.qiniu['use-https']">{{t('general.enable')}}</el-switch>
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model.trim="config.qiniu['access-key']" :placeholder="t('view.systemTools.system.enterAccessKey')" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model.trim="config.qiniu['secret-key']" :placeholder="t('view.systemTools.system.enterSecretKey')" />
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.cdnUploadAccelerated')">
              <el-switch v-model="config.qiniu['use-cdn-domains']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'tencent-cos'">
            <h2>{{ t('view.systemTools.system.tencentCosUploadConfig') }}</h2>
            <el-form-item :label="t('view.systemTools.system.bucketName')">
              <el-input v-model.trim="config['tencent-cos']['bucket']" :placeholder="t('view.systemTools.system.enterBucketName')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.region')">
              <el-input v-model.trim="config['tencent-cos'].region" :placeholder="t('view.systemTools.system.enterRegion')"/>
            </el-form-item>
            <el-form-item label="secretID">
              <el-input v-model.trim="config['tencent-cos']['secret-id']" :placeholder="t('view.systemTools.system.enterSecretID')"/>
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model.trim="config['tencent-cos']['secret-key']" :placeholder="t('view.systemTools.system.enterSecretKey')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.pathPrefix')">
              <el-input v-model.trim="config['tencent-cos']['path-prefix']" :placeholder="t('view.systemTools.system.enterPathPrefix')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.accessDomain')">
              <el-input v-model.trim="config['tencent-cos']['base-url']" :placeholder="t('view.systemTools.system.enterAccessDomain')"/>
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'aliyun-oss'">
            <h2>{{ t('view.systemTools.system.aliyunOssUploadConfig') }}</h2>
            <el-form-item :label="t('view.systemTools.system.region')">
              <el-input v-model.trim="config['aliyun-oss'].endpoint" :label="t('view.systemTools.system.enterRegion')"/>
            </el-form-item>
            <el-form-item label="accessKeyId">
              <el-input v-model.trim="config['aliyun-oss']['access-key-id']" :placeholder="t('view.systemTools.system.enterSecretID')"/>
            </el-form-item>
            <el-form-item label="accessKeySecret">
              <el-input v-model.trim="config['aliyun-oss']['access-key-secret']" :placeholder="t('view.systemTools.system.enterSecretKey')"/>
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.bucketName')">
              <el-input v-model.trim="config['aliyun-oss']['bucket-name']" :placeholder="t('view.systemTools.system.enterBucketName')" />
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.accessDomain')">
              <el-input v-model.trim="config['aliyun-oss']['bucket-url']" :placeholder="t('view.systemTools.system.enterAccessDomain')" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'huawei-obs'">
            <h2>{{ t('view.systemTools.system.huaweiObsUploadConfig') }}</h2>
            <el-form-item :label="t('view.api.path')">
              <el-input v-model.trim="config['hua-wei-obs'].path" :placeholder="t('view.systemTools.system.enterPath')" />
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.bucketName')">
              <el-input v-model.trim="config['hua-wei-obs'].bucket" :placeholder="t('view.systemTools.system.enterBucketName')" />
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.region')">
              <el-input v-model.trim="config['hua-wei-obs'].endpoint" :label="t('view.systemTools.system.enterRegion')" />
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model.trim="config['hua-wei-obs']['access-key']" :placeholder="t('view.systemTools.system.enterAccessKey')" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model.trim="config['hua-wei-obs']['secret-key']" :placeholder="t('view.systemTools.system.enterSecretKey')" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'cloudflare-r2'">
            <h2>{{ t('view.systemTools.system.cloudflareR2UploadConfig') }}</h2>
            <el-form-item :label="t('view.api.path')">
              <el-input v-model.trim="config['cloudflare-r2'].path" :placeholder="t('view.systemTools.system.enterPath')" />
            </el-form-item>
            <el-form-item :label="t('view.systemTools.system.bucketName')">
              <el-input v-model.trim="config['cloudflare-r2'].bucket" :placeholder="t('view.systemTools.system.enterBucketName')" />
            </el-form-item>
            <el-form-item label="Base URL">
              <el-input v-model.trim="config['cloudflare-r2']['base-url']" :placeholder="t('view.systemTools.system.enterBaseUrl')" />
            </el-form-item>
            <el-form-item label="Account ID">
              <el-input v-model.trim="config['cloudflare-r2']['account-id']" :placeholder="t('view.systemTools.system.enterSecretKey')" />
            </el-form-item>
            <el-form-item label="Access Key ID">
              <el-input v-model.trim="config['cloudflare-r2']['access-key-id']" :placeholder="t('view.systemTools.system.enterSecretKey')" />
            </el-form-item>
            <el-form-item label="Secret Access Key">
              <el-input v-model.trim="config['cloudflare-r2']['secret-access-key']" :placeholder="t('view.systemTools.system.enterSecretKey')" />
            </el-form-item>
          </template>

        </el-tab-pane>
        <el-tab-pane
            :label="t('view.systemTools.system.excelUploadConfig')"
            name="11"
            class="mt-3.5"
        >
          <el-form-item :label="t('view.systemTools.system.excelDir')">
            <el-input v-model.trim="config.excel.dir" :placeholder="t('view.systemTools.system.enterCompositeTargetAddress')" />
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
            <el-input v-model.trim="config.autocode['server']" :placeholder="t('view.systemTools.system.enterBackendCodeAddress')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendApiPath')">
            <el-input v-model.trim="config.autocode['server-api']" :placeholder="t('view.systemTools.system.enterBackendApiFolderAddress')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendInitPath')">
            <el-input v-model.trim="config.autocode['server-initialize']" :placeholder="t('view.systemTools.system.enterBackendInitializeFolder')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendModelPath')">
            <el-input v-model.trim="config.autocode['server-model']" :placeholder="t('view.systemTools.system.enterBackendModelFileAddress')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendRequestPath')">
            <el-input v-model.trim="config.autocode['server-request']" :placeholder="t('view.systemTools.system.enterBackendRequestFolderAddress')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendRouterPath')">
            <el-input v-model.trim="config.autocode['server-router']" :placeholder="t('view.systemTools.system.enterBackendRouterFolderAddress')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendServicePath')">
            <el-input v-model.trim="config.autocode['server-service']" :placeholder="t('view.systemTools.system.enterBackendServiceFolderAddress')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendCodePath')">
            <el-input v-model.trim="config.autocode.web" :placeholder="t('view.systemTools.system.enterFrontendFolderAddress')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendApiPath')">
            <el-input v-model.trim="config.autocode['web-api']" :placeholder="t('view.systemTools.system.enterBackendWApiFolderAddress')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendFormPath')">
            <el-input v-model.trim="config.autocode['web-form']" :placeholder="t('view.systemTools.system.enterBackendWFormFolderAddress')"/>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendTablePath')">
            <el-input v-model.trim="config.autocode['web-table']" :placeholder="t('view.systemTools.system.enterBackendWTableFolderAddress')"/>
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
import { ElMessage, ElMessageBox } from 'element-plus'
import { Minus, Plus } from '@element-plus/icons-vue'
import { emailTest } from '@/api/email'
import {CreateUUID} from "@/utils/format";
import { ref, reactive } from 'vue'
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
      t('view.systemTools.system.confirmRestartService'),
      t('general.warning'),
      {
        confirmButtonText: t('general.confirm'),
        cancelButtonText: t('general.cancel'),
        type: 'warning',
      },
  )
      .then(async() => {
        const res = await reloadSystem()
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: t('view.systemTools.system.operationSuccess'),
          })
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: t('view.systemTools.system.cancelRestart'),
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
