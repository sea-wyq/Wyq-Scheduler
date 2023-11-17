# Wyq-Scheduler

通过k8s的scheduling framework实现一个自定义调度器Wyq-Scheduler

目标: 了解k8s pod调度的流程，了解每个扩展点的功能。
## 开始
- 构建 Wyq-Scheduler
```shell
make push 
```
- 部署 Wyq-Scheduler:
```shell
kubectl apply -f deploy/
```
- 检查调度器状态:
```shell
kubectl get pods -n kube-system 
```
- 部署测试pod

```shell
kubectl apply -f example/test-pod.yaml
```
## 调度流程
- 查看调度器的调度日志
```
# kubectl logs -f -n kube-system pod/test-scheduling-6bff584984-c4xg4
# ...
# I1117 05:56:05.241710       1 scheduler.go:443] "Attempting to schedule pod" pod="default/test"
# I1117 05:56:05.242050       1 scheduler.go:47] PreFilter: test
# I1117 05:56:05.242180       1 scheduler.go:60] Filter : pod: test, node: yigou-dev-102-45
# I1117 05:56:05.242192       1 scheduler.go:60] Filter : pod: test, node: yigou-dev-102-46
# I1117 05:56:05.243484       1 scheduler.go:80] PreScore: test
# I1117 05:56:05.243516       1 scheduler.go:82] PreScore: yigou-dev-102-45
# I1117 05:56:05.243526       1 scheduler.go:82] PreScore: yigou-dev-102-46
# I1117 05:56:05.243619       1 scheduler.go:88] Score: pod: test, node: yigou-dev-102-45
# I1117 05:56:05.243662       1 scheduler.go:88] Score: pod: test, node: yigou-dev-102-46
# I1117 05:56:05.243766       1 scheduler.go:98] NormalizeScore: pod: test
# I1117 05:56:05.243949       1 scheduler.go:122] Reserve: pod: test, node: yigou-dev-102-46
# I1117 05:56:05.244022       1 scheduler.go:103] Permit: pod: test, node: yigou-dev-102-46
# I1117 05:56:05.244097       1 scheduler.go:108] PreBind: pod: test, node: yigou-dev-102-46
# I1117 05:56:05.244159       1 default_binder.go:52] "Attempting to bind pod to node" pod="default/test" node="yigou-dev-102-46"
# I1117 05:56:05.275976       1 cache.go:385] "Finished binding for pod, can be expired" pod="default/test"
# I1117 05:56:05.276040       1 scheduler.go:621] "Successfully bound pod to node" pod="default/test" node="yigou-dev-102-46" evaluatedNodes=3 feasibleNodes=2
# I1117 05:56:05.276079       1 scheduler.go:118] PostBind: pod: test, node: yigou-dev-102-46
# I1117 05:56:05.276213       1 eventhandlers.go:161] "Delete event for unscheduled pod" pod="default/test"
# I1117 05:56:05.276327       1 eventhandlers.go:186] "Add event for scheduled pod" pod="default/test"
```