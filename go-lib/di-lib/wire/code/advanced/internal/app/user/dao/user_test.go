package dao

import (
	"testing"

	"github.com/spf13/pflag"
	"hb.study/go-lib/di-lib/wire/code/advanced/pkg/database"
	"hb.study/go-lib/di-lib/wire/code/advanced/setting"

	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	pflag.Set("conf", "../app/cmd/config-httprouter_test.yaml")
	conf, err := setting.InitConfig()
	if err != nil {
		panic(err)
	}
	db, _ = database.InitMySQL(conf)

	m.Run()
}
func TestUserDao_Delete(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "httprouter_test dele",
			fields: fields{DB: db},
			args:   args{id: 1},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserDao{
				DB: tt.fields.DB,
			}
			if got := r.Delete(tt.args.id); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
