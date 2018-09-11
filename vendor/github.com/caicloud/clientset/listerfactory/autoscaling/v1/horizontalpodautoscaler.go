/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/autoscaling/v1"
)

var _ v1.HorizontalPodAutoscalerLister = &horizontalPodAutoscalerLister{}

var _ v1.HorizontalPodAutoscalerNamespaceLister = &horizontalPodAutoscalerNamespaceLister{}

// horizontalPodAutoscalerLister implements the HorizontalPodAutoscalerLister interface.
type horizontalPodAutoscalerLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewHorizontalPodAutoscalerLister returns a new HorizontalPodAutoscalerLister.
func NewHorizontalPodAutoscalerLister(client kubernetes.Interface) v1.HorizontalPodAutoscalerLister {
	return NewFilteredHorizontalPodAutoscalerLister(client, nil)
}

func NewFilteredHorizontalPodAutoscalerLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1.HorizontalPodAutoscalerLister {
	return &horizontalPodAutoscalerLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all HorizontalPodAutoscalers in the indexer.
func (s *horizontalPodAutoscalerLister) List(selector labels.Selector) (ret []*autoscalingv1.HorizontalPodAutoscaler, err error) {
	listopt := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.AutoscalingV1().HorizontalPodAutoscalers(metav1.NamespaceAll).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// HorizontalPodAutoscalers returns an object that can list and get HorizontalPodAutoscalers.
func (s *horizontalPodAutoscalerLister) HorizontalPodAutoscalers(namespace string) v1.HorizontalPodAutoscalerNamespaceLister {
	return horizontalPodAutoscalerNamespaceLister{client: s.client, tweakListOptions: s.tweakListOptions, namespace: namespace}
}

// horizontalPodAutoscalerNamespaceLister implements the HorizontalPodAutoscalerNamespaceLister
// interface.
type horizontalPodAutoscalerNamespaceLister struct {
	client           kubernetes.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// List lists all HorizontalPodAutoscalers in the indexer for a given namespace.
func (s horizontalPodAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*autoscalingv1.HorizontalPodAutoscaler, err error) {
	listopt := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.AutoscalingV1().HorizontalPodAutoscalers(s.namespace).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the HorizontalPodAutoscaler from the indexer for a given namespace and name.
func (s horizontalPodAutoscalerNamespaceLister) Get(name string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	return s.client.AutoscalingV1().HorizontalPodAutoscalers(s.namespace).Get(name, metav1.GetOptions{})
}