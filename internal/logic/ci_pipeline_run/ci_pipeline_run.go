package ci_pipeline_run

import (
	"context"
	"devops-super/api"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"devops-super/utility/thirdclients/kubernetes"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sCiPipelineRun struct{}

var cols = dao.CiPipelineRun.Columns()

func init() {
	service.RegisterCiPipelineRun(New())
}

func New() *sCiPipelineRun {
	return &sCiPipelineRun{}
}

func (*sCiPipelineRun) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.CiPipelineRun], err error) {
	out = &api.PageLstRes[*entity.CiPipelineRun]{}
	m := dao.CiPipelineRun.Ctx(ctx).Safe(true)
	if pipelineId := in.Wheres.Get("pipelineId"); !pipelineId.IsNil() {
		m = m.Where(m.Builder().Where(cols.PipelineId, pipelineId.Int()))
	}
	err = m.Offset(in.Offset()).Limit(in.Limit()).OrderDesc(cols.Id).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (s *sCiPipelineRun) GetWithKubernetesClient(ctx context.Context, id int) (*entity.CiPipelineRun, *kubernetes.Client, error) {
	e, err := s.Get(ctx, &do.CiPipelineRun{Id: id})
	if err != nil {
		return nil, nil, err
	}

	eCiPipeline, err := service.CiPipeline().Get(ctx, &do.CiPipeline{Id: e.PipelineId})
	if err != nil {
		return nil, nil, err
	}

	if eCiPipeline == nil {
		return nil, nil, gerror.New("找不到源流水线")
	}

	eSecret, err := service.Secret().Get(ctx, &do.Secret{Id: eCiPipeline.KubernetesConfigId})
	if err != nil {
		return nil, nil, err
	}
	if eSecret == nil {
		return nil, nil, gerror.New("找不到 Kubernetes 配置")
	}

	kubeConfig := new(mid.TextContent)
	if err = eSecret.Content.Scan(kubeConfig); err != nil {
		return nil, nil, err
	}

	kubeClient, err := kubernetes.NewClient(ctx, kubeConfig.Text)
	if err != nil {
		return nil, nil, gerror.Wrap(err, "创建 kubernetes client 失败")
	}
	return e, kubeClient, nil
}

func (*sCiPipelineRun) Get(ctx context.Context, in *do.CiPipelineRun) (out *entity.CiPipelineRun, err error) {
	err = dao.CiPipelineRun.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (s *sCiPipelineRun) Cancel(ctx context.Context, id int) error {
	e, kubeClient, err := s.GetWithKubernetesClient(ctx, id)
	if err != nil {
		return err
	}

	if err := kubeClient.DeletePodForce(ctx, e.Namespace, e.PodName); err != nil && !kubernetes.IsNotFoundError(err) {
		return err
	}

	e.Status = 3

	if _, err := dao.CiPipelineRun.Ctx(ctx).Fields(cols.Status).Data(e).WherePri(e.Id).Update(); err != nil {
		return err
	}

	return nil
}
