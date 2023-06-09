## GVA LDAP 登陆功能插件
#### 开发者：GIN-VUE-ADMIN 官方

### 使用步骤

#### 1. 前往GVA主程序下的initialize/plugin.go 在InstallPlugin 方法最末尾按照你需要的及安全模式添加本插件
    例：
    本插件可以采用gva的配置文件 也可以直接写死内容作为配置 建议为gva添加配置文件结构 然后将配置传入
    
	PluginInit(PrivateGroup, ldap.CreateLdaplPlug(
        global.GlobalConfig.Host,
        global.GlobalConfig.BaseDN,
        global.GlobalConfig.BindDN,
        global.GlobalConfig.BindPwd,
        global.GlobalConfig.Filter,
        global.GlobalConfig.FieldMap,
        global.GlobalConfig.TLS,
		))

    同样也可以再传入时写死

   	PluginInit(PublicGroup, ldap.CreateLdaplPlug(
		"127.0.0.1:389",
		"dc=example,dc=org",
		"cn=admin,dc=example,dc=org",
		"admin",
		"(&(objectClass=organizationalPerson)(cn=%s))",
		"name=cn&email=mail&phone=telephoneNumber",
		false,
	))

### 2. 配置说明

#### 2-1 全局配置结构体说明

    type Ldap struct {
        BindDN   string `mapstructure:"bind_dn" json:"bind_dn" yaml:"bind_dn"`  // 管理员DN 例如："cn=admin,dc=example,dc=org"
        BindPwd  string `mapstructure:"bind_pwd" json:"bind_pwd" yaml:"bind_pwd"` // 管理员密码（明文） 
        Host     string `mapstructure:"host" json:"host" yaml:"host"` // ldap域名及端口，例如：127.0.0.1:389
        TLS      bool   `mapstructure:"tls" json:"tls" yaml:"tls"` // 是否SSL   是否开启SSL
        BaseDN   string `mapstructure:"base_dn" json:"base_dn" yaml:"base_dn"` // "dc=example,dc=org"
        Filter   string `mapstructure:"filter" json:"filter" yaml:"filter"`  // 查询用户使用的过滤器 例如："(&(objectClass=organizationalPerson)(cn=%s))"
        // ldap字段与系统的映射关系 例如："name=cn&email=mail&phone=telephoneNumber"，name是gva系统参数，cn是ldap配置的参数字段
        FieldMap string `mapstructure:"field_map" json:"field_map" yaml:"field_map"` 
    }



### 3. 方法API

    utils.Search(req *utils.ldap.SearchRequest, resp *utils.ldap.SearchResult) 查询ldap信息
    例:
    var resp *utils.ldap.SearchResult
    utils.Search(utils.BuildSearchRequest(username), resp)
    
    utils.SearchFunc(req *utils.ldap.SearchRequest, fn func(*ldap.SearchResult) error)) 查询ldap信息
    
    例:
    utils.Search(utils.BuildSearchRequest(username), func(resp *ldap.SearchResult) error){
        // 做业务处理
        return nil 
    })


### 4. 可直接调用的接口

    Ldap登陆：  /ldap/login [post] 已配置swagger

### 5. ldap和 phpldapadmin 快速安装
```sh
#!/bin/bash -e
docker run -p 389:389 -p 636:636 --env LDAP_ADMIN_PASSWORD="admin"  --name ldap-service --hostname ldap-service --detach osixia/openldap
docker run -p 8081:80 --privileged --name phpldapadmin-service --hostname phpldapadmin-service --link ldap-service:ldap-host --env PHPLDAPADMIN_HTTPS=false --env PHPLDAPADMIN_LDAP_HOSTS=ldap-host --detach osixia/phpldapadmin

PHPLDAP_IP=$(docker inspect -f "{{ .NetworkSettings.IPAddress }}" phpldapadmin-service)

echo "Go to: https://$PHPLDAP_IP"
echo "Login DN: cn=admin,dc=example,dc=org"
echo "Password: admin"
```

    –privileged 特权模式启动(使用该参数，container 内的 root 拥有真正的 root 权限。
    否则，container 内的 root 只是外部的一个普通用户权限。)
    
    –env PHPLDAPADMIN_HTTPS=false 禁用 HTTPS
    
    –env PHPLDAPADMIN_LDAP_HOSTS =127.0.0.1  配置 openLDAP 的 IP 或者域名
    此处设置访问端口为 8081



运行OpenLDAP镜像后，在LDAP容器中进行第一次搜索
```sh
docker exec ldap-service ldapsearch -x -H ldap://localhost -b dc=example,dc=org -D "cn=admin,dc=example,dc=org" -w admin
```

通过phpldapadmin创建LDAP用户，浏览器访问 http://127.0.0.1:8081/
登陆名：cn=admin,dc=example,dc=org
密码：admin

### 6. web 调试
    1. 前往GVA主程序下的 web/src/router/index.js, 更改登陆页的路由即可
    {
        path: '/login',
        name: 'Login',
        // component: () => import('@/view/login/index.vue')
        component: () => import('@/plugin/ldap/view/index.vue')
    }
    
    2. 为了复用登陆逻辑需要修改： web/src/pinia/modules/user.js
    根据不同的登陆提供者处理登陆逻辑    
    const LoginIn = async(loginInfo, provider) => {
        loadingInstance.value = ElLoading.service({
          fullscreen: true,
          text: '登录中，请稍候...',
        })
        try {
          const res = await provider(loginInfo)
    
   
### api插入语句

```sql
    INSERT INTO `你的数据库名字`.`sys_apis`(`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES ( '2021-08-25 23:09:12', '2021-08-25 23:09:12', NULL, '/ldap/login', 'LDAP登陆', 'ldap', 'POST');
```
