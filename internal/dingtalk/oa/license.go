package oa

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkworkflow "github.com/alibabacloud-go/dingtalk/workflow_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/coffee377/autoctl/internal/dingtalk/app"
	"github.com/coffee377/autoctl/pkg/log"
)

type Lic struct {
	app.App
	client *dingtalkworkflow.Client
}

func NewOA(app app.App) (*Lic, error) {
	config := new(openapi.Config)
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, err := dingtalkworkflow.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Lic{App: app, client: client}, nil
}

func (l *Lic) Demo() {
	accessToken := l.GetAccessToken()
	headers := &dingtalkworkflow.ListProcessInstanceIdsHeaders{
		XAcsDingtalkAccessToken: tea.String(accessToken),
	}
	request := &dingtalkworkflow.ListProcessInstanceIdsRequest{}
	request.SetProcessCode("PROC-2C30849D-B138-47E8-90ED-F1BE217444AC")

	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).UnixMilli()

	log.Info("审批实例开始时间: %d", startTime)

	request.SetStartTime(startTime)
	request.SetNextToken(0)
	request.SetMaxResults(20)
	//request.SetUserIds(tea.StringSlice([]string{"3015582306805324"}))
	response, err := l.client.ListProcessInstanceIdsWithOptions(request, headers, &util.RuntimeOptions{})
	if err != nil {
		panic(err)
	}

	dir := "D:\\project\\personal\\Lic-grpc"

	// 删除 dir 目录下所有 *.Lic *.rif
	_ = filepath.WalkDir(dir, func(path string, file fs.DirEntry, err error) error {
		if !file.IsDir() && (filepath.Ext(path) == ".Lic" || filepath.Ext(path) == ".rif") {
			_ = os.Remove(path)
		}
		return nil
	})

	for _, instanceId := range response.Body.Result.List {

		// 获取审批实列详情
		getProcessInstanceHeaders := &dingtalkworkflow.GetProcessInstanceHeaders{
			XAcsDingtalkAccessToken: tea.String(accessToken),
		}
		getProcessInstanceRequest := &dingtalkworkflow.GetProcessInstanceRequest{
			ProcessInstanceId: instanceId,
		}
		res, err := l.client.GetProcessInstanceWithOptions(getProcessInstanceRequest, getProcessInstanceHeaders, &util.RuntimeOptions{})
		if err != nil {
			panic(err)
		}

		log.Info("")
		log.Info("审批实例：%s", *instanceId)
		log.Info("标   题：%s", *res.Body.Result.Title)
		log.Info("审批编号：%s", *res.Body.Result.BusinessId)
		log.Info("部   门：%s", *res.Body.Result.OriginatorDeptName)

		var attachments []attachment

		for _, e := range res.Body.Result.FormComponentValues {
			if *e.ComponentType == "DDAttachment" && *e.Name == "上传导出注册文件" {
				_ = json.Unmarshal([]byte(*e.Value), &attachments)
				for _, a := range attachments {
					if !a.isRegistrationInformationFile() {
						continue
					}
					// 注册信息文件

					grantProcessInstanceForDownloadFileHeaders := &dingtalkworkflow.GrantProcessInstanceForDownloadFileHeaders{
						XAcsDingtalkAccessToken: tea.String(accessToken),
					}
					grantProcessInstanceForDownloadFileRequest := &dingtalkworkflow.GrantProcessInstanceForDownloadFileRequest{
						ProcessInstanceId: instanceId,
						FileId:            tea.String(a.FileId),
					}
					fileRes, err := l.client.GrantProcessInstanceForDownloadFileWithOptions(grantProcessInstanceForDownloadFileRequest, grantProcessInstanceForDownloadFileHeaders, &util.RuntimeOptions{})
					if err != nil {
						panic(err)
					}
					downloadUri := fileRes.Body.Result.DownloadUri
					log.Info("注册信息文件下载 url：%s", *downloadUri)

					_, _ = a.downloadFile(*downloadUri, *res.Body.Result.BusinessId, &dir)
				}
				break
			}
			continue
		}

	}

}

type attachment struct {
	SpaceId  string  `json:"spaceId"`
	FileName string  `json:"fileName"`
	FileSize float64 `json:"fileSize"`
	FileType string  `json:"fileType"`
	FileId   string  `json:"fileId"`
}

func (receiver attachment) isRegistrationInformationFile() bool {
	reg := regexp.MustCompile("RegistrationInformation.*.rif")
	if receiver.FileType == "rif" && reg.MatchString(receiver.FileName) {
		return true
	}
	return false
}

func (receiver attachment) downloadFile(url, businessId string, dir *string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("文件下载出错，检查文件地址: %v", err)
	}
	if dir == nil {
		dir = tea.String("RegistrationInformation")
	}
	filename := filepath.Join(*dir, fmt.Sprintf("%s.rif", businessId))
	_ = os.MkdirAll(*dir, 0755)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(f, err)
		return "", fmt.Errorf("文件保存出错，检查目录: %v", err)
	}
	_, err = io.Copy(f, res.Body)
	if err != nil {
		return "", fmt.Errorf("文件保存出错: %v", err)
	}
	return filename, nil
}
