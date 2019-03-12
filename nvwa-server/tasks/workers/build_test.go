package workers

import (
    "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs"
    "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
    "github.com/astaxie/beego"
    "log"
    "path/filepath"
    "runtime"
    "testing"
    "time"
)

// reset appPath so beego can find correct configuration file.
func init() {
    _, file, _, _ := runtime.Caller(1)
    appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../../"+string(filepath.Separator))))
    beego.TestBeegoInit(appPath)
}

// deal build record
func Test_BuildDealOnce(t *testing.T) {
    DefaultBuildWorker.DealOnce()
    time.Sleep(time.Second * 60)
}

func Test_JenkinsTemp(t *testing.T) {
    sys := svrs.DefaultSystemSvr.Get()
    app, err := svrs.DefaultAppSvr.GetById(1)
    if err != nil {
        log.Fatal(err.Error())
    }

    content, err := libs.ParseTemplate(sys.JenkinsTemplate, struct {
        App                 *entities.AppEntity
        JenkinsCredentialId string
    }{
        App:                 app,
        JenkinsCredentialId: "jci",
    })
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println(content)
}
