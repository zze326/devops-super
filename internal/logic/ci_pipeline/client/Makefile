image.push:
	rm -f devops-super-cli && go env -w 'GOPROXY=https://goproxy.cn,direct' && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o devops-super-cli
	docker build --platform linux/amd64 -t registry.cn-shenzhen.aliyuncs.com/zze/devops-super-ci-client:tmp .
	docker push registry.cn-shenzhen.aliyuncs.com/zze/devops-super-ci-client:tmp
	rm -f devops-super-cli