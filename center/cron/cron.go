package cron

import (
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/robfig/cron/v3"

	"github.com/follow/model"
	"github.com/follow/utils/dir"
	"github.com/follow/utils/times"
)

var (
	c              *cron.Cron
	initialization int32
)

var (
	ErrNotInitialization = errors.New("cron not initialization")
)

var (
	rTask = make(map[string]cron.EntryID) // key is username + script
	tLock sync.RWMutex
)

// RegistrateTask 注册任务
func RegistrateTask(spec, shell, taskName string) (cron.EntryID, error) {
	if atomic.LoadInt32(&initialization) != 1 {
		return 0, ErrNotInitialization
	}
	slog.Info(fmt.Sprintf("start registrate [%s] task", taskName))
	defer slog.Info(fmt.Sprintf("registrate [%s] task end", taskName))

	id, err := c.AddFunc(spec, func() {
		cmd := exec.Command("sh", "-c", shell)
		out, err := cmd.Output()
		if err != nil {
			slog.Error(fmt.Sprintf("run task [%s] failed: %v", taskName, err))
			return
		}
		// out 回存储到数据库，目前只是打印
		slog.Info(fmt.Sprintf("run task [%s] result: %v", taskName, string(out)))

		u, s := getUS(taskName)
		r := model.Result{
			ScriptName: s,
			Username:   u,
			Content:    string(out),
			Image:      "https://pic.netbian.com/uploads/allimg/240322/233416-1711121656e5bd.jpg",
			Link:       "https://pic.netbian.com/uploads/allimg/240322/233416-1711121656e5bd.jpg",
			Header:     "test",
			CreateTime: times.GetCurTimeInt(),
		}
		if err := r.Create(); err != nil {
			slog.Error(fmt.Sprintf("create failed: %v", err))
		}
	})

	if err != nil {
		return 0, fmt.Errorf("add func failed: %v", err)
	}
	tLock.Lock()
	rTask[taskName] = id
	tLock.Unlock()
	return id, nil
}

// RemoveTask 删除任务
func RemoveTask(taskName string) error {
	tLock.RLock()
	id, ok := rTask[taskName]
	if !ok {
		tLock.RUnlock()
		return nil
	}
	tLock.RUnlock()

	if atomic.LoadInt32(&initialization) != 1 {
		return ErrNotInitialization
	}
	c.Remove(id) // 简单处理

	tLock.Lock()
	_, ok = rTask[taskName]
	if ok {
		delete(rTask, taskName)
	}
	defer tLock.Unlock()

	return nil
}

func InitCron(sfa model.ScriptFileAddr) error {
	c = cron.New()
	c.Start()
	atomic.StoreInt32(&initialization, 1)

	s := model.Script{}
	Scripts, err := s.GetAll()
	if err != nil {
		return fmt.Errorf("get all scripts failed: %w", err)
	}
	// 创建文件夹
	if err := dir.CreateDir(sfa.Addr); err != nil {
		return fmt.Errorf("create dir [%s] failed", err)
	}
	var lock sync.RWMutex
	rsScripts := make([]*model.Script, 0) // 存储注册成功脚本
	var wg sync.WaitGroup
	for _, script := range Scripts {
		if !script.Status { // 只处理开启状态的脚本
			continue
		}
		wg.Add(1)
		// 直接全量刷新
		go func(s model.Script) {
			defer wg.Done()
			if err := s.Build(sfa); err != nil {
				slog.Info("build failed: %v", err)
				return
			}
			spec := fmt.Sprintf("*/%d * * * *", s.Cycle/60)
			entryID, err := RegistrateTask(spec, s.RunShell, SetTaskName(s.Username, s.Name))
			if err != nil {
				slog.Error(fmt.Sprintf("registrate task failed: %v", err))
				return
			}
			s.CronID = entryID
			lock.Lock()
			rsScripts = append(rsScripts, &s)
			lock.Unlock()
		}(script)
	}

	wg.Wait()

	fmt.Println("this is a test:", rTask)
	for _, script := range rsScripts {
		if err := script.CreateOrUpdate(); err != nil {
			slog.Error(fmt.Sprintf("update script failed: %v", err)) // 简单处理，只打印错误日志
		}
	}
	return nil
}

func SetTaskName(username, scriptName string) string {
	return username + "|" + scriptName
}

func getUS(taskName string) (string, string) {
	item := strings.Split(taskName, "|")
	return item[0], item[1]
}
