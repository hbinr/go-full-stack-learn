// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	systemUserFieldNames          = builder.RawFieldNames(&SystemUser{}, true)
	systemUserRows                = strings.Join(systemUserFieldNames, ",")
	systemUserRowsExpectAutoSet   = strings.Join(stringx.Remove(systemUserFieldNames, "id", "create_time", "update_time", "create_t", "update_at"), ",")
	systemUserRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(systemUserFieldNames, "id", "create_time", "update_time", "create_at", "update_at"))
)

type (
	systemUserModel interface {
		Insert(ctx context.Context, data *SystemUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SystemUser, error)
		Update(ctx context.Context, data *SystemUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSystemUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SystemUser struct {
		Id          int64          `db:"id"`           // 用户ID,  主键
		DeptId      int64          `db:"dept_id"`      // 部门ID
		RoleId      int64          `db:"role_id"`      // 角色ID
		UserName    string         `db:"user_name"`    // 用户账号
		NickName    sql.NullString `db:"nick_name"`    // 用户昵称
		UserType    int64          `db:"user_type"`    // 用户类型（0系统用户）
		Email       sql.NullString `db:"email"`        // 用户邮箱
		PhoneNumber sql.NullString `db:"phone_number"` // 手机号码
		Avatar      sql.NullString `db:"avatar"`       // 头像地址
		Password    sql.NullString `db:"password"`     // 密码
		Status      bool           `db:"status"`       // 帐号状态（true正常 false停用）
		LoginIp     sql.NullString `db:"login_ip"`     // 最后登录IP
		LoginDate   sql.NullTime   `db:"login_date"`   // 最后登录时间
		CreateBy    sql.NullString `db:"create_by"`    // 创建者
		CreatedAt   sql.NullTime   `db:"created_at"`   // 创建时间
		UpdatedAt   sql.NullTime   `db:"updated_at"`   // 更新时间
		DeletedAt   time.Time      `db:"deleted_at"`   // 删除时间，为null则未删除
	}
)

func newSystemUserModel(conn sqlx.SqlConn) *defaultSystemUserModel {
	return &defaultSystemUserModel{
		conn:  conn,
		table: `"public"."system_user"`,
	}
}

func (m *defaultSystemUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSystemUserModel) FindOne(ctx context.Context, id int64) (*SystemUser, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", systemUserRows, m.table)
	var resp SystemUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSystemUserModel) Insert(ctx context.Context, data *SystemUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", m.table, systemUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeptId, data.RoleId, data.UserName, data.NickName, data.UserType, data.Email, data.PhoneNumber, data.Avatar, data.Password, data.Status, data.LoginIp, data.LoginDate, data.CreateBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	return ret, err
}

func (m *defaultSystemUserModel) Update(ctx context.Context, data *SystemUser) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, systemUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.DeptId, data.RoleId, data.UserName, data.NickName, data.UserType, data.Email, data.PhoneNumber, data.Avatar, data.Password, data.Status, data.LoginIp, data.LoginDate, data.CreateBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	return err
}

func (m *defaultSystemUserModel) tableName() string {
	return m.table
}
