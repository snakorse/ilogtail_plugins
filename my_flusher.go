package ilogtail_plugins

import (
	"github.com/alibaba/ilogtail/pkg/models"
	"github.com/alibaba/ilogtail/pkg/pipeline"
	"github.com/alibaba/ilogtail/pkg/protocol"
)

type MyFlusher struct {
}

func (m *MyFlusher) Init(context pipeline.Context) error {
	//TODO implement me
	panic("implement me")
}

func (m *MyFlusher) Description() string {
	//TODO implement me
	panic("implement me")
}

func (m *MyFlusher) IsReady(projectName string, logstoreName string, logstoreKey int64) bool {
	//TODO implement me
	panic("implement me")
}

func (m *MyFlusher) SetUrgent(flag bool) {
	//TODO implement me
	panic("implement me")
}

func (m *MyFlusher) Stop() error {
	//TODO implement me
	panic("implement me")
}

func (m *MyFlusher) Flush(projectName string, logstoreName string, configName string, logGroupList []*protocol.LogGroup) error {
	//TODO implement me
	panic("implement me")
}

func (m *MyFlusher) Export(events []*models.PipelineGroupEvents, context pipeline.PipelineContext) error {
	//TODO implement me
	panic("implement me")
}

func init() {
	pipeline.AddFlusherCreator("my_flusher", func() pipeline.Flusher {
		return &MyFlusher{}
	})
}
