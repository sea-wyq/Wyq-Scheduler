package wyq

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	Name = "wyq"
)

var (
	_ framework.QueueSortPlugin = &WYQ{} // 排序插件在集群中只能出现一个，如果存在默认的k8s调度器，则该插件无法使用.

	_ framework.FilterPlugin     = &WYQ{}
	_ framework.PreFilterPlugin  = &WYQ{}
	_ framework.PostFilterPlugin = &WYQ{}
	_ framework.PreScorePlugin   = &WYQ{}
	_ framework.ScorePlugin      = &WYQ{}
	_ framework.ReservePlugin    = &WYQ{}

	_ framework.PermitPlugin   = &WYQ{}
	_ framework.PreBindPlugin  = &WYQ{}
	_ framework.BindPlugin     = &WYQ{}
	_ framework.PostBindPlugin = &WYQ{}
)

type WYQ struct {
	h framework.Handle
}

func (y *WYQ) Name() string {
	return Name
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &WYQ{h: h}, nil
}

// 过滤带有wyq=Unschedulable的pod,无法调度
func (y *WYQ) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("PreFilter: %v", pod.Name)
	if pod.Labels["wyq"] == "Unschedulable" {
		return framework.NewStatus(framework.Unschedulable, "")
	}
	return framework.NewStatus(framework.Success, "")
}

func (y *WYQ) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}

// 如果节点打了wyq=Unschedulable 这个标签，则无法调度
func (y *WYQ) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("Filter : pod: %v, node: %v", pod.Name, nodeInfo.Node().GetName())
	for k, v := range nodeInfo.Node().Labels {
		if k == "wyq" && v == "Unschedulable" {
			return framework.NewStatus(framework.Unschedulable, "")
		}
	}
	return framework.NewStatus(framework.Success, "")
}

func (y *WYQ) PostFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, filteredNodeStatusMap framework.NodeToStatusMap) (*framework.PostFilterResult, *framework.Status) {
	klog.V(3).Infof("PostFilter: pod: %v", pod.Name)
	return nil, framework.NewStatus(framework.Success, "")
}

func (y *WYQ) Less(podInfo1 *framework.QueuedPodInfo, podInfo2 *framework.QueuedPodInfo) bool {
	klog.V(3).Infof("Less: %v, %v", podInfo1.Pod.Name, podInfo2.Pod.Name)
	return true
}

func (y *WYQ) PreScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodes []*v1.Node) *framework.Status {
	klog.V(3).Infof("PreScore: %v", pod.Name)
	for _, node := range nodes {
		klog.V(3).Infof("PreScore: %v", node.Name)
	}
	return framework.NewStatus(framework.Success, "")
}

func (y *WYQ) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.V(3).Infof("Score: pod: %v, node: %v", pod.Name, nodeName)
	nodeScore := int64(0)
	return nodeScore, framework.NewStatus(framework.Success, "")
}

func (y *WYQ) ScoreExtensions() framework.ScoreExtensions {
	return y
}

func (y *WYQ) NormalizeScore(_ context.Context, _ *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	klog.V(3).Infof("NormalizeScore: pod: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}

func (y *WYQ) Permit(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (*framework.Status, time.Duration) {
	klog.V(3).Infof("Permit: pod: %v, node: %v", pod.Name, nodeName)
	return nil, time.Hour
}

func (y *WYQ) PreBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("PreBind: pod: %v, node: %v", pod.Name, nodeName)
	return framework.NewStatus(framework.Success, "")
}

func (y *WYQ) Bind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("Bind: pod: %v, node: %v", pod.Name, nodeName)
	return framework.NewStatus(framework.Success, "")
}

func (y *WYQ) PostBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) {
	klog.V(3).Infof("PostBind: pod: %v, node: %v", pod.Name, nodeName)
}

func (y *WYQ) Reserve(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("Reserve: pod: %v, node: %v", pod.Name, nodeName)
	return framework.NewStatus(framework.Success, "")
}

func (y *WYQ) Unreserve(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) {
	klog.V(3).Infof("Unreserve: pod: %v, node: %v", pod.Name, nodeName)
}
