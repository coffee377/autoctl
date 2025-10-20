package oa

import (
	"cds/dingtalk/app"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/alibabacloud-go/tea/tea"

	"github.com/coffee377/autoctl/pkg/log"
)

type Lic struct {
	*Approval
}

func NewLic(app app.App) (*Lic, error) {
	approval, err := New(app)
	if err != nil {
		return nil, err
	}
	return &Lic{Approval: approval}, nil
}

func (l *Lic) Run(startTime string, licDir string) error {
	log.Info("审批实例开始时间: %s", startTime)

	ids, err := l.GetProcessInstanceIds("PROC-2C30849D-B138-47E8-90ED-F1BE217444AC", startTime, "", nil)
	if err != nil {
		return err
	}

	if licDir != "" {
		// 删除 licDir 目录下所有 *.Lic *.rif
		_ = filepath.WalkDir(licDir, func(path string, file fs.DirEntry, err error) error {
			if !file.IsDir() && (filepath.Ext(path) == ".Lic" || filepath.Ext(path) == ".rif") {
				_ = os.Remove(path)
			}
			return nil
		})
	}

	for _, instanceId := range ids {

		// 获取审批实列详情
		instance, err := l.GetProcessInstance(instanceId)
		if err != nil {
			return err
		}

		log.Info("")
		log.Info("审批实例：%s", instanceId)
		log.Info("标   题：%s", *instance.Title)
		log.Info("审批编号：%s", *instance.BusinessId)
		log.Info("部   门：%s", *instance.OriginatorDeptName)

		if licDir == "" {
			continue
		}
		var attachments []attachment

		for _, e := range instance.FormComponentValues {
			if *e.ComponentType == "DDAttachment" && *e.Name == "上传导出注册文件" {
				_ = json.Unmarshal([]byte(*e.Value), &attachments)
				for _, a := range attachments {
					if !a.isRegistrationInformationFile() {
						continue
					}
					downloadUri, err := l.GetAttachmentDownloadUri(instanceId, a.FileId)
					if err != nil {
						return err
					}
					if downloadUri == nil {
						continue
					}
					log.Info("注册信息文件下载 url：%s", downloadUri)
					_, _ = a.downloadFile(*downloadUri, *instance.BusinessId, &licDir)
				}
				break
			}
			continue
		}

	}

	return nil
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
