package defalut

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

var (
	_ framework.QueueSortPlugin = &Defalut{} // 排序插件在集群中只能出现一个，如果存在默认的k8s调度器，则该插件无法使用.

	_ framework.FilterPlugin     = &Defalut{}
	_ framework.PreFilterPlugin  = &Defalut{}
	_ framework.PostFilterPlugin = &Defalut{}
	_ framework.PreScorePlugin   = &Defalut{}
	_ framework.ReservePlugin    = &Defalut{}

	_ framework.PermitPlugin   = &Defalut{}
	_ framework.PreBindPlugin  = &Defalut{}
	_ framework.BindPlugin     = &Defalut{}
	_ framework.PostBindPlugin = &Defalut{}
)

type Defalut struct{}

const (
	Name = "defalut"
)

func (y *Defalut) Name() string {
	return Name
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &Defalut{}, nil
}

// Pod 的信息进行预处理 检查pod的合规性
// 过滤带有wyq=Unschedulable的pod,无法调度
// PreFilter在调度周期开始时调用。所有预过滤插件必须返回成功，否则pod将被拒绝。
func (y *Defalut) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("Defalut PreFilter: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}

func (y *Defalut) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}

// 排除那些不能运行该 Pod 的节点
// 如果节点打了wyq=Unschedulable 这个标签，则无法调度
func (y *Defalut) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("Defalut Filter : pod: %v, node: %v", pod.Name, nodeInfo.Node().GetName())
	return framework.NewStatus(framework.Success, "")
}

func (y *Defalut) PostFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, filteredNodeStatusMap framework.NodeToStatusMap) (*framework.PostFilterResult, *framework.Status) {
	klog.V(3).Infof("Defalut PostFilter: pod: %v", pod.Name)
	return nil, framework.NewStatus(framework.Success, "")
}

// 插件用于对调度队列中的pod进行排序。一次只能启用一个队列排序插件。
func (y *Defalut) Less(podInfo1 *framework.QueuedPodInfo, podInfo2 *framework.QueuedPodInfo) bool {
	klog.V(3).Infof("Defalut Less: %v, %v", podInfo1.Pod.Name, podInfo2.Pod.Name)
	return true
}

// 调度框架在节点列表通过筛选阶段后调用PreScore。所有的prescore插件必须返回成功，否则pod将被拒绝
func (y *Defalut) PreScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodes []*v1.Node) *framework.Status {
	klog.V(3).Infof("Defalut PreScore: %v", pod.Name)
	for _, node := range nodes {
		klog.V(3).Infof("Defalut PreScore: %v", node.Name)
	}
	return framework.NewStatus(framework.Success, "")
}

func (y *Defalut) NormalizeScore(_ context.Context, _ *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	klog.V(3).Infof("Defalut NormalizeScore: pod: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}

func (y *Defalut) Permit(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (*framework.Status, time.Duration) {
	klog.V(3).Infof("Defalut Permit: pod: %v, node: %v", pod.Name, nodeName)
	return nil, time.Hour
}

func (y *Defalut) PreBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("Defalut PreBind: pod: %v, node: %v", pod.Name, nodeName)
	return framework.NewStatus(framework.Success, "")
}

func (y *Defalut) Bind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("Defalut Bind: pod: %v, node: %v", pod.Name, nodeName)
	return framework.NewStatus(framework.Success, "")
}

func (y *Defalut) PostBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) {
	klog.V(3).Infof("Defalut PostBind: pod: %v, node: %v", pod.Name, nodeName)
}

func (y *Defalut) Reserve(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("Defalut Reserve: pod: %v, node: %v", pod.Name, nodeName)
	return framework.NewStatus(framework.Success, "")
}

func (y *Defalut) Unreserve(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) {
	klog.V(3).Infof("Defalut Unreserve: pod: %v, node: %v", pod.Name, nodeName)
}
