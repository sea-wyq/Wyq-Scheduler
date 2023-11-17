package randscore

import (
	"context"
	"math/rand"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	Name = "randScore"
)

var (
	_ framework.ScorePlugin = &RandScore{}
)

type RandScore struct {
	h framework.Handle
}

func (y *RandScore) Name() string {
	return Name
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &RandScore{h: h}, nil
}

// 为所有可选节点进行随机打分
func (y *RandScore) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	nodeScore := rand.Int63n(101)
	klog.V(3).Infof("RandScore Score: pod: %v, node: %v, socer: %v", pod.Name, nodeName, nodeScore)
	return nodeScore, framework.NewStatus(framework.Success, "")
}

func (y *RandScore) ScoreExtensions() framework.ScoreExtensions {
	return nil
}
