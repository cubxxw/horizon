package manager

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"g.hz.netease.com/horizon/lib/orm"
	"g.hz.netease.com/horizon/pkg/pipelinerun/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	ctx context.Context
)

func Test(t *testing.T) {
	pr, err := Mgr.Create(ctx, &models.Pipelinerun{
		ID:               0,
		ClusterID:        1,
		Action:           models.ActionBuildDeploy,
		Status:           "created",
		Title:            "title",
		Description:      "description",
		GitURL:           "",
		GitBranch:        "",
		GitCommit:        "",
		ImageURL:         "",
		LastConfigCommit: "",
		ConfigCommit:     "1",
		S3Bucket:         "",
		LogObject:        "",
		PrObject:         "",
		CreatedBy:        0,
	})
	assert.Nil(t, err)
	t.Logf("%v", pr)

	prGet, err := Mgr.GetByID(ctx, pr.ID)
	assert.Nil(t, err)
	assert.Equal(t, "title", prGet.Title)
	assert.Equal(t, "1", prGet.ConfigCommit)

	err = Mgr.UpdateConfigCommitByID(ctx, prGet.ID, "2")
	assert.Nil(t, err)

	prGet, err = Mgr.GetByID(ctx, pr.ID)
	assert.Nil(t, err)
	assert.Equal(t, "2", prGet.ConfigCommit)

	prGet, err = Mgr.GetLatestByClusterIDAndAction(ctx, pr.ClusterID, models.ActionBuildDeploy)
	assert.Nil(t, err)
	assert.Equal(t, "2", prGet.ConfigCommit)

	err = Mgr.UpdateResultByID(ctx, pr.ID, &models.Result{
		S3Bucket:   "bucket",
		LogObject:  "log-obj",
		PrObject:   "pr-obj",
		Result:     "ok",
		StartedAt:  func() *time.Time { t := time.Now(); return &t }(),
		FinishedAt: func() *time.Time { t := time.Now(); return &t }(),
	})
	assert.Nil(t, err)

	prGet, err = Mgr.GetByID(ctx, pr.ID)
	assert.Nil(t, err)
	assert.Equal(t, prGet.S3Bucket, "bucket")
	assert.Equal(t, prGet.LogObject, "log-obj")
	assert.Equal(t, prGet.PrObject, "pr-obj")
	assert.Equal(t, prGet.Status, "ok")
	b, _ := json.Marshal(prGet)
	t.Logf("%v", string(b))

	err = Mgr.DeleteByID(ctx, pr.ID)
	assert.Nil(t, err)

	prGet, err = Mgr.GetByID(ctx, pr.ID)
	assert.Nil(t, err)
	assert.Nil(t, prGet)
}

func TestMain(m *testing.M) {
	db, _ = orm.NewSqliteDB("")
	if err := db.AutoMigrate(&models.Pipelinerun{}); err != nil {
		panic(err)
	}
	ctx = orm.NewContext(context.TODO(), db)
	os.Exit(m.Run())
}
