package common

import (
	"context"
	"devops-super/internal/consts"
	"devops-super/internal/model/do"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"devops-super/utility/util"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type sCommon struct{}

func init() {
	service.RegisterCommon(New())
}

func New() *sCommon {
	return &sCommon{}
}

func (*sCommon) GetGitBranchLst(ctx context.Context, gitUrl string, secretId int) ([]string, error) {
	eSecret, err := service.Secret().Get(ctx, &do.Secret{Id: secretId})
	if err != nil {
		return nil, err
	}

	if eSecret.Type == consts.SECRET_TYPE_GIT_BASIC_AUTH {
		secretContent := new(mid.UsernamePasswordContent)
		if err := eSecret.Content.Scan(secretContent); err != nil {
			return nil, err
		}

		auth := &http.BasicAuth{
			Username: secretContent.Username,
			Password: secretContent.Password,
		}

		return util.GetRemoteBranchList(gitUrl, auth)
	}

	return nil, nil
}
