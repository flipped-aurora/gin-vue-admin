create table casbin_rule
(
    id    bigint unsigned auto_increment
        primary key,
    ptype varchar(100) null,
    v0    varchar(100) null,
    v1    varchar(100) null,
    v2    varchar(100) null,
    v3    varchar(100) null,
    v4    varchar(100) null,
    v5    varchar(100) null,
    constraint idx_casbin_rule
        unique (ptype, v0, v1, v2, v3, v4, v5)
);

create table cominfo
(
    id               bigint unsigned auto_increment
        primary key,
    created_at       datetime(3)  null,
    updated_at       datetime(3)  null,
    deleted_at       datetime(3)  null,
    com_title        varchar(100) null comment '比赛标题',
    com_introduction varchar(300) null comment '比赛简介',
    com_picture      varchar(100) null comment '比赛图片',
    com_model        bigint       null comment '比赛等级',
    com_start        datetime(3)  null comment '比赛开始时间',
    com_end          datetime(3)  null comment '比赛结束时间',
    com_number       bigint       null,
    com_type         bigint       null comment '比赛类型'
);

create index idx_comInfo_deleted_at
    on cominfo (deleted_at);

create table exa_customers
(
    id                    bigint unsigned auto_increment
        primary key,
    created_at            datetime(3)     null,
    updated_at            datetime(3)     null,
    deleted_at            datetime(3)     null,
    customer_name         varchar(191)    null comment '客户名',
    customer_phone_data   varchar(191)    null comment '客户手机号',
    sys_user_id           bigint unsigned null comment '管理ID',
    sys_user_authority_id bigint unsigned null comment '管理角色ID'
);

create index idx_exa_customers_deleted_at
    on exa_customers (deleted_at);

create table exa_file_chunks
(
    id                bigint unsigned auto_increment
        primary key,
    created_at        datetime(3)     null,
    updated_at        datetime(3)     null,
    deleted_at        datetime(3)     null,
    exa_file_id       bigint unsigned null,
    file_chunk_number bigint          null,
    file_chunk_path   varchar(191)    null
);

create index idx_exa_file_chunks_deleted_at
    on exa_file_chunks (deleted_at);

create table exa_file_upload_and_downloads
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)  null,
    updated_at datetime(3)  null,
    deleted_at datetime(3)  null,
    name       varchar(191) null comment '文件名',
    url        varchar(191) null comment '文件地址',
    tag        varchar(191) null comment '文件标签',
    `key`      varchar(191) null comment '编号'
);

create index idx_exa_file_upload_and_downloads_deleted_at
    on exa_file_upload_and_downloads (deleted_at);

create table exa_files
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime(3)  null,
    updated_at  datetime(3)  null,
    deleted_at  datetime(3)  null,
    file_name   varchar(191) null,
    file_md5    varchar(191) null,
    file_path   varchar(191) null,
    chunk_total bigint       null,
    is_finish   tinyint(1)   null
);

create index idx_exa_files_deleted_at
    on exa_files (deleted_at);

create table jwt_blacklists
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,
    jwt        text        null comment 'jwt'
);

create index idx_jwt_blacklists_deleted_at
    on jwt_blacklists (deleted_at);

create table sys_apis
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime(3)                 null,
    updated_at  datetime(3)                 null,
    deleted_at  datetime(3)                 null,
    path        varchar(191)                null comment 'api路径',
    description varchar(191)                null comment 'api中文描述',
    api_group   varchar(191)                null comment 'api组',
    method      varchar(191) default 'POST' null comment '方法'
);

create index idx_sys_apis_deleted_at
    on sys_apis (deleted_at);

create table sys_authorities
(
    created_at     datetime(3)                      null,
    updated_at     datetime(3)                      null,
    deleted_at     datetime(3)                      null,
    authority_id   bigint unsigned auto_increment comment '角色ID'
        primary key,
    authority_name varchar(191)                     null comment '角色名',
    parent_id      bigint unsigned                  null comment '父角色ID',
    default_router varchar(191) default 'dashboard' null comment '默认菜单',
    constraint authority_id
        unique (authority_id)
);

create table sys_authority_btns
(
    authority_id         bigint unsigned null comment '角色ID',
    sys_menu_id          bigint unsigned null comment '菜单ID',
    sys_base_menu_btn_id bigint unsigned null comment '菜单按钮ID'
);

create table sys_authority_menus
(
    sys_base_menu_id           bigint unsigned not null,
    sys_authority_authority_id bigint unsigned not null comment '角色ID',
    primary key (sys_base_menu_id, sys_authority_authority_id)
);

create table sys_auto_code_histories
(
    id             bigint unsigned auto_increment
        primary key,
    created_at     datetime(3)  null,
    updated_at     datetime(3)  null,
    deleted_at     datetime(3)  null,
    package        varchar(191) null,
    business_db    varchar(191) null,
    table_name     varchar(191) null,
    request_meta   text         null,
    auto_code_path text         null,
    injection_meta text         null,
    struct_name    varchar(191) null,
    struct_cn_name varchar(191) null,
    api_ids        varchar(191) null,
    flag           bigint       null
);

create index idx_sys_auto_code_histories_deleted_at
    on sys_auto_code_histories (deleted_at);

create table sys_auto_codes
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime(3)  null,
    updated_at   datetime(3)  null,
    deleted_at   datetime(3)  null,
    package_name varchar(191) null comment '包名',
    label        varchar(191) null comment '展示名',
    `desc`       varchar(191) null comment '描述'
);

create index idx_sys_auto_codes_deleted_at
    on sys_auto_codes (deleted_at);

create table sys_base_menu_btns
(
    id               bigint unsigned auto_increment
        primary key,
    created_at       datetime(3)     null,
    updated_at       datetime(3)     null,
    deleted_at       datetime(3)     null,
    name             varchar(191)    null comment '按钮关键key',
    `desc`           varchar(191)    null,
    sys_base_menu_id bigint unsigned null comment '菜单ID'
);

create index idx_sys_base_menu_btns_deleted_at
    on sys_base_menu_btns (deleted_at);

create table sys_base_menu_parameters
(
    id               bigint unsigned auto_increment
        primary key,
    created_at       datetime(3)     null,
    updated_at       datetime(3)     null,
    deleted_at       datetime(3)     null,
    sys_base_menu_id bigint unsigned null,
    type             varchar(191)    null comment '地址栏携带参数为params还是query',
    `key`            varchar(191)    null comment '地址栏携带参数的key',
    value            varchar(191)    null comment '地址栏携带参数的值'
);

create index idx_sys_base_menu_parameters_deleted_at
    on sys_base_menu_parameters (deleted_at);

create table sys_base_menus
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime(3)     null,
    updated_at   datetime(3)     null,
    deleted_at   datetime(3)     null,
    menu_level   bigint unsigned null,
    parent_id    varchar(191)    null comment '父菜单ID',
    path         varchar(191)    null comment '路由path',
    name         varchar(191)    null comment '路由name',
    hidden       tinyint(1)      null comment '是否在列表隐藏',
    component    varchar(191)    null comment '对应前端文件路径',
    sort         bigint          null comment '排序标记',
    active_name  varchar(191)    null comment '附加属性',
    keep_alive   tinyint(1)      null comment '附加属性',
    default_menu tinyint(1)      null comment '附加属性',
    title        varchar(191)    null comment '附加属性',
    icon         varchar(191)    null comment '附加属性',
    close_tab    tinyint(1)      null comment '附加属性'
);

create index idx_sys_base_menus_deleted_at
    on sys_base_menus (deleted_at);

create table sys_data_authority_id
(
    sys_authority_authority_id     bigint unsigned not null comment '角色ID',
    data_authority_id_authority_id bigint unsigned not null comment '角色ID',
    primary key (sys_authority_authority_id, data_authority_id_authority_id)
);

create table sys_dictionaries
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)  null,
    updated_at datetime(3)  null,
    deleted_at datetime(3)  null,
    name       varchar(191) null comment '字典名（中）',
    type       varchar(191) null comment '字典名（英）',
    status     tinyint(1)   null comment '状态',
    `desc`     varchar(191) null comment '描述'
);

create index idx_sys_dictionaries_deleted_at
    on sys_dictionaries (deleted_at);

create table sys_dictionary_details
(
    id                bigint unsigned auto_increment
        primary key,
    created_at        datetime(3)     null,
    updated_at        datetime(3)     null,
    deleted_at        datetime(3)     null,
    label             varchar(191)    null comment '展示值',
    value             bigint          null comment '字典值',
    extend            varchar(191)    null comment '扩展值',
    status            tinyint(1)      null comment '启用状态',
    sort              bigint          null comment '排序标记',
    sys_dictionary_id bigint unsigned null comment '关联标记'
);

create index idx_sys_dictionary_details_deleted_at
    on sys_dictionary_details (deleted_at);

create table sys_export_templates
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime(3)  null,
    updated_at    datetime(3)  null,
    deleted_at    datetime(3)  null,
    name          varchar(191) null comment '模板名称',
    table_name    varchar(191) null comment '表名称',
    template_id   varchar(191) null comment '模板标识',
    template_info text         null
);

create index idx_sys_export_templates_deleted_at
    on sys_export_templates (deleted_at);

create table sys_operation_records
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime(3)     null,
    updated_at    datetime(3)     null,
    deleted_at    datetime(3)     null,
    ip            varchar(191)    null comment '请求ip',
    method        varchar(191)    null comment '请求方法',
    path          varchar(191)    null comment '请求路径',
    status        bigint          null comment '请求状态',
    latency       bigint          null comment '延迟',
    agent         varchar(191)    null comment '代理',
    error_message varchar(191)    null comment '错误信息',
    body          text            null comment '请求Body',
    resp          text            null comment '响应Body',
    user_id       bigint unsigned null comment '用户id'
);

create index idx_sys_operation_records_deleted_at
    on sys_operation_records (deleted_at);

create table sys_user_authority
(
    sys_user_id                bigint unsigned not null,
    sys_authority_authority_id bigint unsigned not null comment '角色ID',
    primary key (sys_user_id, sys_authority_authority_id)
);

create table sys_users
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime(3)                                                              null,
    updated_at   datetime(3)                                                              null,
    deleted_at   datetime(3)                                                              null,
    uuid         varchar(191)                                                             null comment '用户UUID',
    username     varchar(191)                                                             null comment '用户登录名',
    password     varchar(191)                                                             null comment '用户登录密码',
    nick_name    varchar(191)    default '系统用户'                                       null comment '用户昵称',
    side_mode    varchar(191)    default 'dark'                                           null comment '用户侧边主题',
    header_img   varchar(191)    default 'https://qmplusimg.henrongyi.top/gva_header.jpg' null comment '用户头像',
    base_color   varchar(191)    default '#fff'                                           null comment '基础颜色',
    active_color varchar(191)    default '#1890ff'                                        null comment '活跃颜色',
    authority_id bigint unsigned default '888'                                            null comment '用户角色ID',
    phone        varchar(191)                                                             null comment '用户手机号',
    email        varchar(191)                                                             null comment '用户邮箱',
    enable       bigint          default 1                                                null comment '用户是否被冻结 1正常 2冻结'
);

create index idx_sys_users_deleted_at
    on sys_users (deleted_at);

create index idx_sys_users_username
    on sys_users (username);

create index idx_sys_users_uuid
    on sys_users (uuid);

create table userinfo
(
    id                bigint unsigned auto_increment
        primary key,
    created_at        datetime(3)  null,
    updated_at        datetime(3)  null,
    deleted_at        datetime(3)  null,
    user_wxopenid     varchar(50)  null comment '用户opemid',
    user_nickname     varchar(40)  null comment '用户昵称',
    user_avatar_url   varchar(400) null,
    user_gender       tinyint      null comment '用户性别',
    user_grade        bigint       null comment '用户年级',
    user_profession   varchar(50)  null comment '用户专业',
    user_introduction varchar(200) null comment '用户简介',
    user_label        varchar(100) null comment '用户标签',
    user_city         varchar(40)  null comment '用户城市',
    love_number       bigint       null comment '用户获赞数量',
    session_key       longtext     null,
    user_model        bigint       null comment '用户身份'
);

create table commentinfo
(
    id                  bigint unsigned auto_increment
        primary key,
    created_at          datetime(3)     null,
    updated_at          datetime(3)     null,
    deleted_at          datetime(3)     null,
    comment_user_id     bigint unsigned null comment '评论人id',
    comment_dis_id      bigint unsigned null comment '关联帖子id',
    comment_text        varchar(300)    null comment '评论信息',
    comment_like_number bigint          null comment '评论的评论id',
    constraint fk_commentInfo_user
        foreign key (comment_user_id) references userinfo (id)
);

create index idx_commentInfo_deleted_at
    on commentinfo (deleted_at);

create table commentsoninfo
(
    id                 bigint unsigned auto_increment
        primary key,
    created_at         datetime(3)     null,
    updated_at         datetime(3)     null,
    deleted_at         datetime(3)     null,
    comment_comment_id bigint unsigned null comment '评论的评论id',
    comment_user_id    bigint unsigned null comment '评论人id',
    comment_text       varchar(300)    null comment '评论信息',
    constraint fk_commentInfo_comment_sons
        foreign key (comment_comment_id) references commentinfo (id),
    constraint fk_commentSonInfo_user
        foreign key (comment_user_id) references userinfo (id)
);

create index idx_commentSonInfo_deleted_at
    on commentsoninfo (deleted_at);

create table disinfo
(
    id                 bigint unsigned auto_increment
        primary key,
    created_at         datetime(3)      null,
    updated_at         datetime(3)      null,
    deleted_at         datetime(3)      null,
    dis_com_id         bigint unsigned  null comment '帖子关联比赛id',
    dis_user_id        bigint unsigned  null comment '发帖人id',
    dis_title          varchar(300)     not null comment '帖子标题',
    dis_content        text             null comment '帖子内容',
    dis_love_number    bigint           null comment '帖子点赞人数',
    dis_collect_number bigint           null comment '帖子收藏人数',
    dis_model          bigint           null comment '帖子类型',
    dis_picture        json             null comment '帖子图片',
    dis_look_number    bigint default 0 null comment '帖子阅读量',
    dis_status         bigint default 0 null comment '帖子状态',
    constraint fk_disInfo_com
        foreign key (dis_com_id) references cominfo (id),
    constraint fk_disInfo_user
        foreign key (dis_user_id) references userinfo (id)
);

create index idx_disInfo_deleted_at
    on disinfo (deleted_at);

create index idx_userInfo_deleted_at
    on userinfo (deleted_at);

create table usertest
(
    id   int auto_increment
        primary key,
    name varchar(255) not null,
    age  int          null
);


