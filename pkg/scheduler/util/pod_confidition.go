package util

import (
	"github.com/golang/glog"

	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
)

// PodConditionUpdater updates the condition of a pod based on the passed PodCondition
type PodConditionUpdater interface {
	Update(pod *v1.Pod, podCondition *v1.PodCondition) error
}

// DefaultPodConditionUpdater is the default implementation of the PodConditionUpdater interface
type DefaultPodConditionUpdater struct {
	Client clientset.Interface
}

// Update pod with podCondition
func (pcUpdater *DefaultPodConditionUpdater) Update(pod *v1.Pod, condition *v1.PodCondition) error {
	glog.V(3).Infof("Updating pod condition for %s/%s to (%s==%s)", pod.Namespace, pod.Name, condition.Type, condition.Status)
	if podutil.UpdatePodCondition(&pod.Status, condition) {
		_, err := pcUpdater.Client.CoreV1().Pods(pod.Namespace).UpdateStatus(pod)
		return err
	}
	return nil
}
