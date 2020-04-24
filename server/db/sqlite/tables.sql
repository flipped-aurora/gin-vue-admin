-- casbin_rule definition

CREATE TABLE casbin_rule(
                            p_type varchar(100),
                            v0 varchar(100),
                            v1 varchar(100),
                            v2 varchar(100),
                            v3 varchar(100),
                            v4 varchar(100),
                            v5 varchar(100));


-- ch_cities definition

CREATE TABLE ch_cities(
                          id int,
                          province varchar(255),
                          city varchar(255),
                          code int);


-- ch_provinces definition

CREATE TABLE ch_provinces(
                             id int,
                             province varchar(255));


-- ch_t_d_areainfo definition

CREATE TABLE ch_t_d_areainfo(
                                id int NOT NULL,
                                name varchar(48) NOT NULL,
                                arealevel tinyint NOT NULL,
                                parent_id int,
                                PRIMARY KEY (id)
);


-- exa_customers definition

CREATE TABLE exa_customers(
                              id int NOT NULL,
                              created_at TIMESTAMP,
                              updated_at TIMESTAMP,
                              deleted_at TIMESTAMP,
                              customer_name varchar(255),
                              customer_phone_data varchar(255),
                              sys_user_id int,
                              sys_user_authority_id varchar(255),
                              PRIMARY KEY (id)
);


-- exa_file_chunks definition

CREATE TABLE exa_file_chunks(
                                id int NOT NULL,
                                created_at TIMESTAMP,
                                updated_at TIMESTAMP,
                                deleted_at TIMESTAMP,
                                exa_file_id int,
                                file_chunk_path varchar(255),
                                file_chunk_number int,
                                PRIMARY KEY (id)
);


-- exa_file_upload_and_downloads definition

CREATE TABLE exa_file_upload_and_downloads(
                                              id int NOT NULL,
                                              created_at TIMESTAMP,
                                              updated_at TIMESTAMP,
                                              deleted_at TIMESTAMP,
                                              name varchar(255),
                                              url varchar(255),
                                              tag varchar(255),
                                              "key" varchar(255),
                                              PRIMARY KEY (id)
);


-- exa_files definition

CREATE TABLE exa_files(
                          id int NOT NULL,
                          created_at TIMESTAMP,
                          updated_at TIMESTAMP,
                          deleted_at TIMESTAMP,
                          file_name varchar(255),
                          file_md5 varchar(255),
                          file_path varchar(255),
                          chunk_total int,
                          is_finish tinyint,
                          PRIMARY KEY (id)
);


-- jwt_blacklists definition

CREATE TABLE jwt_blacklists(
                               id int NOT NULL,
                               created_at TIMESTAMP,
                               updated_at TIMESTAMP,
                               deleted_at TIMESTAMP,
                               jwt text(65535),
                               PRIMARY KEY (id)
);


-- sys_apis definition

CREATE TABLE sys_apis(
                         id int NOT NULL,
                         created_at TIMESTAMP,
                         updated_at TIMESTAMP,
                         deleted_at TIMESTAMP,
                         authority_id int,
                         "path" varchar(255),
                         description varchar(255),
                         api_group varchar(255),
                         "method" varchar(255),
                         PRIMARY KEY (id)
);


-- sys_authorities definition

CREATE TABLE sys_authorities(
                                authority_id varchar(255) NOT NULL,
                                authority_name varchar(255),
                                parent_id varchar(255),
                                created_at TIMESTAMP,
                                updated_at TIMESTAMP,
                                deleted_at TIMESTAMP,
                                PRIMARY KEY (authority_id)
);


-- sys_authority_menus definition

CREATE TABLE sys_authority_menus(
                                    sys_authority_authority_id varchar(255) NOT NULL,
                                    sys_base_menu_id int NOT NULL,
                                    PRIMARY KEY (sys_authority_authority_id,sys_base_menu_id)
);


-- sys_base_menus definition

CREATE TABLE sys_base_menus(
                               id int NOT NULL,
                               created_at TIMESTAMP,
                               updated_at TIMESTAMP,
                               deleted_at TIMESTAMP,
                               menu_level int,
                               parent_id int,
                               "path" varchar(255),
                               name varchar(255),
                               hidden tinyint,
                               component varchar(255),
                               title varchar(255),
                               icon varchar(255),
                               nick_name varchar(255),
                               sort int,
                               PRIMARY KEY (id)
);


-- sys_data_authority_id definition

CREATE TABLE sys_data_authority_id(
                                      sys_authority_authority_id varchar(255) NOT NULL,
                                      data_authority_id varchar(255) NOT NULL,
                                      PRIMARY KEY (sys_authority_authority_id,data_authority_id)
);


-- sys_menus definition

CREATE TABLE sys_menus(
                          id int NOT NULL,
                          created_at TIMESTAMP,
                          updated_at TIMESTAMP,
                          deleted_at TIMESTAMP,
                          menu_level int,
                          authority_id int,
                          "path" varchar(255),
                          name varchar(255),
                          hidden tinyint,
                          component varchar(255),
                          title varchar(255),
                          icon varchar(255),
                          parent_id int,
                          menu_id varchar(255),
                          nick_name varchar(255),
                          sort varchar(255),
                          PRIMARY KEY (id)
);


-- sys_users definition

CREATE TABLE sys_users(
                          id int NOT NULL,
                          created_at TIMESTAMP,
                          updated_at TIMESTAMP,
                          deleted_at TIMESTAMP,
                          uuid BLOB,
                          user_name varchar(255),
                          pass_word varchar(255),
                          nick_name varchar(255),
                          header_img varchar(255),
                          authority_id double,
                          authority_name varchar(255),
                          username varchar(255),
                          password varchar(255),
                          phone_data varchar(255),
                          manager varchar(255),
                          PRIMARY KEY (id)
);


-- sys_workflow_step_infos definition

CREATE TABLE sys_workflow_step_infos(
                                        id int NOT NULL,
                                        created_at TIMESTAMP,
                                        updated_at TIMESTAMP,
                                        deleted_at TIMESTAMP,
                                        workflow_id int,
                                        is_strat tinyint,
                                        step_name varchar(255),
                                        step_no double,
                                        step_authority_id varchar(255),
                                        is_end tinyint,
                                        sys_workflow_id int,
                                        PRIMARY KEY (id)
);


-- sys_workflows definition

CREATE TABLE sys_workflows(
                              id int NOT NULL,
                              created_at TIMESTAMP,
                              updated_at TIMESTAMP,
                              deleted_at TIMESTAMP,
                              workflow_nick_name varchar(255),
                              workflow_name varchar(255),
                              workflow_description varchar(255),
                              PRIMARY KEY (id)
);
