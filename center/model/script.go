package model

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/robfig/cron/v3"

	"github.com/follow/utils/dir"
	"github.com/follow/utils/file"
)

// Script 脚本
type Script struct {
	ID          int64        `json:"id" gorm:"primaryKey"`
	Username    string       `json:"username"`             // 用户名
	Name        string       `json:"name" gorm:"unique"`   // 脚本名称
	Type        string       `json:"type" gorm:"not null"` // 脚本类型
	Language    string       `json:"language"`             // 脚本所属语言
	Code        string       `json:"code"`
	Cycle       int          `json:"cycle"`                       // 运行周期，单位min，数据库默认为6min
	Status      bool         `json:"status"`                      // 脚本状态，true为开启，false为关闭，默认不开启
	CronID      cron.EntryID `json:"cron_id"`                     // 任务ID
	RunShell    string       `json:"run_shell"`                   // 运行命令
	CreateTime  int64        `json:"create_time"`                 // 创建时间
	UpdateTime  int64        `json:"update_time"`                 // 更新时间
	Description string       `json:"description" gorm:"not null"` // 描述
}

func (s *Script) Check() error {
	elem := reflect.ValueOf(s).Elem()
	elemType := elem.Type()
	for index := 0; index < elem.NumField(); index++ {
		field, fieldName := elem.Field(index), elemType.Field(index).Name
		switch field.Kind() {
		case reflect.String:
			if fieldName == "RunShell" {
				continue
			}
			if field.String() == "" {
				return fmt.Errorf("field [%s] expect", fieldName)
			}
		default:
			// do nothing
		}
	}
	return nil
}

// GetByID 根据ID获取脚本
func (s *Script) GetByID() error {
	cursor, err := getDB()
	if err != nil {
		return err
	}
	if cursor.Where("id=?", s.ID).Find(s).Error != nil {
		return fmt.Errorf("find failed: %v", err)
	}
	return err
}

// GetByUsername 获取用户名下的脚本
func (s *Script) GetByUsername() ([]Script, error) {
	cursor, err := getDB()
	if err != nil {
		return []Script{}, err
	}
	scripts := make([]Script, 0)
	if cursor.Where("username=?", s.Username).Find(&scripts).Error != nil {
		return []Script{}, fmt.Errorf("find failed: %v", err)
	}
	return scripts, nil
}

// Delete 根据ID删除脚本
func (s *Script) Delete() error {
	cursor, err := getDB()
	if err != nil {
		return err
	}
	return cursor.Delete("id=?", s.ID).Error
}

// Create 创建脚本或者更新脚本
func (s *Script) CreateOrUpdate() error {
	cursor, err := getDB()
	if err != nil {
		return err
	}

	return cursor.Where("id=?", s.ID).Save(s).Error
}

// Update 更新脚本
func (s *Script) Update() error {
	return nil
}

// Exist 判断脚本是否存在
func (s *Script) Exist() (bool, error) {
	cursor, err := getDB()
	if err != nil {
		return false, err
	}
	scripts := make([]Script, 0)
	err = cursor.Where("username=? and name=?", s.Username, s.Name).Find(&scripts).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if len(scripts) > 0 {
		return true, nil
	}
	return false, nil
}

// GetAll get all script
func (s *Script) GetAll() ([]Script, error) {
	cursor, err := getDB()
	if err != nil {
		return []Script{}, err
	}
	scripts := make([]Script, 0)
	if cursor.Find(&scripts).Error != nil {
		return []Script{}, fmt.Errorf("find failed: %v", err)
	}
	return scripts, nil
}

func (s *Script) Build(sfa ScriptFileAddr) error {
	tmpScriptDir := filepath.Join(sfa.Addr, s.Name)
	scriptDir, err := filepath.Abs(tmpScriptDir)
	if err != nil {
		return fmt.Errorf("get abs [%s] failed: %v", tmpScriptDir, err)
	}
	if err := goBuild(scriptDir, s.Code, sfa.RunFilename); err != nil {
		return fmt.Errorf("script [%s] go build failed: %v", s.Code, err)
	}
	s.RunShell = fmt.Sprintf("cd %s && ./%s", scriptDir, sfa.RunFilename)
	return nil
}

func goBuild(dirName, code, runFilename string) error {
	if err := dir.CreateDir(dirName); err != nil {
		return fmt.Errorf("create dir [%s] failed: %w", dirName, err)
	}
	sfilename := filepath.Join(dirName, "main.go")
	if err := file.WriteFile(sfilename, []byte(code)); err != nil {
		return fmt.Errorf("write [%s] file failed: %w", sfilename, err)
	}
	cmd := exec.Command("go", "build", "-o", filepath.Join(dirName, runFilename), sfilename)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run [%s] failed: %w", cmd.String(), err)
	}
	return nil
}
