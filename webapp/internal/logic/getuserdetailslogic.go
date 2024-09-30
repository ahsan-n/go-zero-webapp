package logic

import (
	"ahsan-n/go-zero-webapp/model"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"

	"ahsan-n/go-zero-webapp/webapp/internal/svc"
	"ahsan-n/go-zero-webapp/webapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var conn sqlx.SqlConn

func init() {
	connStr := "root:my-secret-pw@tcp(127.0.0.1:3306)/users"
	conn = sqlx.NewMysql(connStr)
	_, err := conn.RawDB()
	if err != nil {
		//logx.Must(err)
		log.Fatal(err)
	}

}

type GetUserDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailsLogic {
	return &GetUserDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailsLogic) GetUserDetails(req types.Request) (resp *types.UserDetails, err error) {

	// todo: add your logic here and delete this line

	m := model.NewUsersModel(conn)
	user, err := m.FindOne(l.ctx, req.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	resp = &types.UserDetails{
		Name:     user.Name,
		Age:      user.Age,
		Semester: int32(user.Semester),
		CGPA:     user.Cgpa,
	}

	return resp, nil
}
