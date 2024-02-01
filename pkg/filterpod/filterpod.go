package filterpod

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	Name = "filteringPod"
)

var (
	_ framework.PreFilterPlugin = &FilteringPod{}
)

type FilteringPod struct {
	h framework.Handle
}

func (y *FilteringPod) Name() string {
	return Name
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &FilteringPod{h: h}, nil
}

// 过滤带有wyq=Unschedulable的pod,无法调度
func (y *FilteringPod) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("FilteringPod PreFilter: %v", pod.Name)
	if pod.Labels["wyq"] == "Unschedulable" {
		return framework.NewStatus(framework.Unschedulable, "")
	}
	return framework.NewStatus(framework.Success, "")
}

func (y *FilteringPod) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}
