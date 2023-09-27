package user

import (
	"context"
	"devops-super/api"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	"os"
	"path/filepath"
	"time"
)

type sHost struct{}

var cols = dao.Host.Columns()

func init() {
	service.RegisterHost(New())
}

func New() *sHost {
	return &sHost{}
}

func (*sHost) Add(ctx context.Context, in *entity.Host) (err error) {
	r, err := dao.Host.Ctx(ctx).Where(cols.Name, in.Name).One()
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.New("已存在该名称的主机")
	}
	_, err = dao.Host.Ctx(ctx).Insert(in)
	return
}

func (*sHost) Upt(ctx context.Context, in *do.Host) (err error) {
	if !gutil.IsEmpty(in.Name) {
		r, err := dao.Host.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Name, in.Name).One()
		if err != nil {
			return err
		}
		if r != nil {
			return gerror.New("已存在该名称的主机")
		}
	}
	_, err = dao.Host.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sHost) GetCountByHostGroupId(ctx context.Context, hostGroupId int) (int, error) {
	return dao.Host.Ctx(ctx).Where(cols.HostGroupId, hostGroupId).Count()
}

func (*sHost) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.Host], err error) {
	out = &api.PageLstRes[*entity.Host]{}
	m := dao.Host.Ctx(ctx).Safe(true)
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, in.SearchStr()).WhereOrLike(cols.HostAddr, in.SearchStr()))
	}

	if hostGroupId := in.Wheres.Get("hostGroupId"); !hostGroupId.IsNil() {
		m = m.Where(cols.HostGroupId, hostGroupId.Int())
	}

	err = m.Offset(in.Offset()).OrderDesc(cols.Id).Limit(in.Limit()).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (*sHost) Get(ctx context.Context, in *do.Host) (out *entity.Host, err error) {
	err = dao.Host.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sHost) Del(ctx context.Context, in *do.Host) (err error) {
	_, err = dao.Host.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}

func (s *sHost) TestSSH(ctx context.Context, in *entity.Host) (err error) {
	client, err := s.SshClient(in)
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Close(); err != nil {
			glog.Error(ctx, err)
		}
	}()
	return nil
}

func (s *sHost) DownloadFile(ctx context.Context, in *mid.DownloadFileIn) error {
	eHost, err := s.Get(ctx, &do.Host{Id: in.Id})
	if err != nil {
		return err
	}
	client, err := s.SftpClient(eHost)
	if err != nil {
		return err
	}

	file, err := client.OpenFile(in.Path, os.O_RDONLY)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			glog.Error(ctx, err)
		}
	}()

	//var length int64 = -1
	//if size, err := file.Seek(0, 2); err == nil {
	//	if _, err = file.Seek(0, 0); err == nil {
	//		length = size
	//	}
	//}

	res := g.RequestFromCtx(ctx).Response
	res.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(in.Path)))
	res.ServeContent(filepath.Base(in.Path), time.Now(), file)

	return nil
}

func (s *sHost) CanAccess(ctx context.Context, in *entity.Host) (bool, error) {
	if service.CurrentUser(ctx).IsAdmin() {
		return true, nil
	}

	// 1. 获取机器所属主机组 in
	eHostGroup, err := service.HostGroup().Get(ctx, &do.HostGroup{Id: in.HostGroupId})
	if err != nil {
		return false, err
	}
	// 2. 获取主机组授权的角色和用户 eHostGroup.RoleIds eHostGroup.UserIds
	// 3. 获取当前用户的角色
	eUser, err := service.User().Get(ctx, &do.User{Id: service.CurrentUser(ctx).UserId})
	if err != nil {
		return false, err
	}
	// 4. 如果当前用户存在于主机组授权的用户列表，则有权限
	for _, hostGroupUserId := range eHostGroup.UserIds.Array() {
		if eUser.Id == gconv.Int(hostGroupUserId) {
			return true, nil
		}
	}
	// 5. 如果当前用户拥有的角色存在于主机组授权的角色，则有权限
	for _, userRoleId := range eUser.RoleIds.Array() {
		for _, hostGroupRoleId := range eHostGroup.RoleIds.Array() {
			if userRoleId == hostGroupRoleId {
				return true, nil
			}
		}
	}

	return false, nil
}
