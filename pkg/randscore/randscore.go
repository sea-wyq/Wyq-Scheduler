package randscore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

// var NodeToIp = map[string]string{
// 	"yigou-dev-102-44": "100.95.137.154",
// 	"yigou-dev-102-45": "100.122.47.218",
// 	"yigou-dev-102-46": "100.73.234.168",
// }

type RandScore struct {
	h framework.Handle
}

func (y *RandScore) Name() string {
	return Name
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &RandScore{h: h}, nil
}

// 如果某节点存在某镜像打100分。否则0分
func (y *RandScore) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	nodeScore := int64(0)
	data, err := os.ReadFile("/etc/nodetoip/nodetoip.json")
	if err != nil {
		panic(err)
	}

	var NodeToIp map[string]string
	if err = json.Unmarshal(data, &NodeToIp); err != nil {
		panic(err)
	}

	for _, env := range pod.Spec.Containers[0].Env {
		if env.Name == "PUBLICIMAGE" {
			client := &http.Client{}
			url := fmt.Sprintf("http://%s:8088?reference=%s", NodeToIp[nodeName], env.Value)
			resp, err := client.Get(url)
			if err != nil {
				klog.V(3).ErrorS(err, "call scheduler-image-web failed")
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				klog.V(3).ErrorS(err, "get resp.Body filed")
			}
			if len(body) > 0 {
				klog.V(3).Infof("get docker image successed form  env PUBLICIMAGE.imageID: %s", string(body))
				nodeScore = int64(100)
			}
		}
	}
	klog.V(3).Infof("RandScore Score: pod: %v, node: %v, socer: %v", pod.Name, nodeName, nodeScore)
	return nodeScore, framework.NewStatus(framework.Success, "")
}

func (y *RandScore) ScoreExtensions() framework.ScoreExtensions {
	return nil
}
