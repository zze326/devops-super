package host_terminal_session

import (
	"context"
	"devops-super/api"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
	"devops-super/utility/util"
	"github.com/gogf/gf/v2/util/gutil"
)

type sHostTerminalSession struct{}

var cols = dao.HostTerminalSession.Columns()

func init() {
	service.RegisterHostTerminalSession(New())
}

func New() *sHostTerminalSession {
	return &sHostTerminalSession{}
}

func (*sHostTerminalSession) Get(ctx context.Context, in *do.HostTerminalSession) (out *entity.HostTerminalSession, err error) {
	err = dao.HostTerminalSession.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sHostTerminalSession) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.HostTerminalSession], err error) {
	out = &api.PageLstRes[*entity.HostTerminalSession]{}
	m := dao.HostTerminalSession.Ctx(ctx).Safe(true)
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.HostName, in.SearchStr()).WhereOrLike(cols.HostAddr, in.SearchStr()))
	}
	err = m.Offset(in.Offset()).OrderDesc(cols.Id).Limit(in.Limit()).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (*sHostTerminalSession) CheckSessionFile(ctx context.Context, id int) (bool, error) {
	eSession := new(entity.HostTerminalSession)
	err := dao.HostTerminalSession.Ctx(ctx).WherePri(id).Scan(eSession)
	return util.FileExists(eSession.Filepath), err
}
