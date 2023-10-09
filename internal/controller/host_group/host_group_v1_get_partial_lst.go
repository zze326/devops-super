package host_group

import (
	"context"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"

	"devops-super/api/host_group/v1"
)

func (c *ControllerV1) GetPartialLst(ctx context.Context, req *v1.GetPartialLstReq) (res *v1.GetPartialLstRes, err error) {
	var resList []*mid.HostGroupPartial
	eHostGroupList, err := service.HostGroup().GetLst(ctx, "")
	if err != nil {
		return nil, err
	}
	for _, eHostGorup := range eHostGroupList {
		hostCount, err := service.Host().GetCountByHostGroupId(ctx, eHostGorup.Id)
		if err != nil {
			return nil, err
		}
		resList = append(resList, &mid.HostGroupPartial{
			Id:        eHostGorup.Id,
			Name:      eHostGorup.Name,
			ParentId:  eHostGorup.ParentId,
			HostCount: hostCount,
		})
	}

	res = &v1.GetPartialLstRes{
		List: resList,
	}

	return
}
