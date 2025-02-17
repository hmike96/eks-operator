/*
Copyright 2019 Wrangler Sample Controller Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/rancher/eks-operator/pkg/apis/eks.cattle.io/v1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type EKSClusterConfigHandler func(string, *v1.EKSClusterConfig) (*v1.EKSClusterConfig, error)

type EKSClusterConfigController interface {
	generic.ControllerMeta
	EKSClusterConfigClient

	OnChange(ctx context.Context, name string, sync EKSClusterConfigHandler)
	OnRemove(ctx context.Context, name string, sync EKSClusterConfigHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() EKSClusterConfigCache
}

type EKSClusterConfigClient interface {
	Create(*v1.EKSClusterConfig) (*v1.EKSClusterConfig, error)
	Update(*v1.EKSClusterConfig) (*v1.EKSClusterConfig, error)
	UpdateStatus(*v1.EKSClusterConfig) (*v1.EKSClusterConfig, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.EKSClusterConfig, error)
	List(namespace string, opts metav1.ListOptions) (*v1.EKSClusterConfigList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.EKSClusterConfig, err error)
}

type EKSClusterConfigCache interface {
	Get(namespace, name string) (*v1.EKSClusterConfig, error)
	List(namespace string, selector labels.Selector) ([]*v1.EKSClusterConfig, error)

	AddIndexer(indexName string, indexer EKSClusterConfigIndexer)
	GetByIndex(indexName, key string) ([]*v1.EKSClusterConfig, error)
}

type EKSClusterConfigIndexer func(obj *v1.EKSClusterConfig) ([]string, error)

type eKSClusterConfigController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewEKSClusterConfigController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) EKSClusterConfigController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &eKSClusterConfigController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromEKSClusterConfigHandlerToHandler(sync EKSClusterConfigHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.EKSClusterConfig
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.EKSClusterConfig))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *eKSClusterConfigController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.EKSClusterConfig))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateEKSClusterConfigDeepCopyOnChange(client EKSClusterConfigClient, obj *v1.EKSClusterConfig, handler func(obj *v1.EKSClusterConfig) (*v1.EKSClusterConfig, error)) (*v1.EKSClusterConfig, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *eKSClusterConfigController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *eKSClusterConfigController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *eKSClusterConfigController) OnChange(ctx context.Context, name string, sync EKSClusterConfigHandler) {
	c.AddGenericHandler(ctx, name, FromEKSClusterConfigHandlerToHandler(sync))
}

func (c *eKSClusterConfigController) OnRemove(ctx context.Context, name string, sync EKSClusterConfigHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromEKSClusterConfigHandlerToHandler(sync)))
}

func (c *eKSClusterConfigController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *eKSClusterConfigController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *eKSClusterConfigController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *eKSClusterConfigController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *eKSClusterConfigController) Cache() EKSClusterConfigCache {
	return &eKSClusterConfigCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *eKSClusterConfigController) Create(obj *v1.EKSClusterConfig) (*v1.EKSClusterConfig, error) {
	result := &v1.EKSClusterConfig{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *eKSClusterConfigController) Update(obj *v1.EKSClusterConfig) (*v1.EKSClusterConfig, error) {
	result := &v1.EKSClusterConfig{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *eKSClusterConfigController) UpdateStatus(obj *v1.EKSClusterConfig) (*v1.EKSClusterConfig, error) {
	result := &v1.EKSClusterConfig{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *eKSClusterConfigController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *eKSClusterConfigController) Get(namespace, name string, options metav1.GetOptions) (*v1.EKSClusterConfig, error) {
	result := &v1.EKSClusterConfig{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *eKSClusterConfigController) List(namespace string, opts metav1.ListOptions) (*v1.EKSClusterConfigList, error) {
	result := &v1.EKSClusterConfigList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *eKSClusterConfigController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *eKSClusterConfigController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.EKSClusterConfig, error) {
	result := &v1.EKSClusterConfig{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type eKSClusterConfigCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *eKSClusterConfigCache) Get(namespace, name string) (*v1.EKSClusterConfig, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.EKSClusterConfig), nil
}

func (c *eKSClusterConfigCache) List(namespace string, selector labels.Selector) (ret []*v1.EKSClusterConfig, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EKSClusterConfig))
	})

	return ret, err
}

func (c *eKSClusterConfigCache) AddIndexer(indexName string, indexer EKSClusterConfigIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.EKSClusterConfig))
		},
	}))
}

func (c *eKSClusterConfigCache) GetByIndex(indexName, key string) (result []*v1.EKSClusterConfig, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.EKSClusterConfig, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.EKSClusterConfig))
	}
	return result, nil
}

type EKSClusterConfigStatusHandler func(obj *v1.EKSClusterConfig, status v1.EKSClusterConfigStatus) (v1.EKSClusterConfigStatus, error)

type EKSClusterConfigGeneratingHandler func(obj *v1.EKSClusterConfig, status v1.EKSClusterConfigStatus) ([]runtime.Object, v1.EKSClusterConfigStatus, error)

func RegisterEKSClusterConfigStatusHandler(ctx context.Context, controller EKSClusterConfigController, condition condition.Cond, name string, handler EKSClusterConfigStatusHandler) {
	statusHandler := &eKSClusterConfigStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromEKSClusterConfigHandlerToHandler(statusHandler.sync))
}

func RegisterEKSClusterConfigGeneratingHandler(ctx context.Context, controller EKSClusterConfigController, apply apply.Apply,
	condition condition.Cond, name string, handler EKSClusterConfigGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &eKSClusterConfigGeneratingHandler{
		EKSClusterConfigGeneratingHandler: handler,
		apply:                             apply,
		name:                              name,
		gvk:                               controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterEKSClusterConfigStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type eKSClusterConfigStatusHandler struct {
	client    EKSClusterConfigClient
	condition condition.Cond
	handler   EKSClusterConfigStatusHandler
}

func (a *eKSClusterConfigStatusHandler) sync(key string, obj *v1.EKSClusterConfig) (*v1.EKSClusterConfig, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type eKSClusterConfigGeneratingHandler struct {
	EKSClusterConfigGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *eKSClusterConfigGeneratingHandler) Remove(key string, obj *v1.EKSClusterConfig) (*v1.EKSClusterConfig, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.EKSClusterConfig{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *eKSClusterConfigGeneratingHandler) Handle(obj *v1.EKSClusterConfig, status v1.EKSClusterConfigStatus) (v1.EKSClusterConfigStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.EKSClusterConfigGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
