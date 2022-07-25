-- Table: public.system_user

-- DROP TABLE IF EXISTS public.system_user;

CREATE TABLE IF NOT EXISTS public.system_user
(
    id bigint NOT NULL DEFAULT nextval('system_user_id_seq'::regclass),
    dept_id smallint DEFAULT 0,
    role_id smallint DEFAULT 0,
    user_name character varying(36) COLLATE pg_catalog."default" NOT NULL,
    nick_name character varying(36) COLLATE pg_catalog."default",
    user_type smallint DEFAULT 0,
    email character varying(50) COLLATE pg_catalog."default",
    phone_number character varying(11) COLLATE pg_catalog."default",
    avatar character varying(128) COLLATE pg_catalog."default",
    password character varying(128) COLLATE pg_catalog."default",
    status boolean DEFAULT true,
    login_ip character varying(128) COLLATE pg_catalog."default",
    login_date timestamp without time zone,
    create_by character varying(36) COLLATE pg_catalog."default",
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone NOT NULL,
    CONSTRAINT system_user_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.system_user
    OWNER to postgres;

COMMENT ON TABLE public.system_user
    IS '用户表';

COMMENT ON COLUMN public.system_user.id
    IS '用户ID,  主键';

COMMENT ON COLUMN public.system_user.dept_id
    IS '部门ID';

COMMENT ON COLUMN public.system_user.role_id
    IS '角色ID';

COMMENT ON COLUMN public.system_user.user_name
    IS '用户账号';

COMMENT ON COLUMN public.system_user.nick_name
    IS '用户昵称';

COMMENT ON COLUMN public.system_user.user_type
    IS '用户类型（0系统用户）';

COMMENT ON COLUMN public.system_user.email
    IS '用户邮箱';

COMMENT ON COLUMN public.system_user.phone_number
    IS '手机号码';

COMMENT ON COLUMN public.system_user.avatar
    IS '头像地址';

COMMENT ON COLUMN public.system_user.password
    IS '密码';

COMMENT ON COLUMN public.system_user.status
    IS '帐号状态（true正常 false停用）';

COMMENT ON COLUMN public.system_user.login_ip
    IS '最后登录IP';

COMMENT ON COLUMN public.system_user.login_date
    IS '最后登录时间';

COMMENT ON COLUMN public.system_user.create_by
    IS '创建者';

COMMENT ON COLUMN public.system_user.created_at
    IS '创建时间';

COMMENT ON COLUMN public.system_user.updated_at
    IS '更新时间';

COMMENT ON COLUMN public.system_user.deleted_at
    IS '删除时间，为null则未删除';