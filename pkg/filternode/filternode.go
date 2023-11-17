package filternode

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	Name = "filteringNode"
)

var (
	_ framework.FilterPlugin = &FilteringNode{}
)

type FilteringNode struct {
	h framework.Handle
}

func (y *FilteringNode) Name() string {
	return Name
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &FilteringNode{h: h}, nil
}

// 如果节点打了wyq=Unschedulable 这个标签，则无法调度
func (y *FilteringNode) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("FilteringNode Filter : pod: %v, node: %v", pod.Name, nodeInfo.Node().GetName())
	for k, v := range nodeInfo.Node().Labels {
		if k == "wyq" && v == "Unschedulable" {
			return framework.NewStatus(framework.Unschedulable, "")
		}
	}
	return framework.NewStatus(framework.Success, "")
}
