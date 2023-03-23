/*
Copyright 2023 James Wong<jameswong1376@gmail.com>.

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

package controllers

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	renginev1beta1 "github.com/wl4g/rengine-operator/api/v1beta1"
)

// RengineApiServerReconciler reconciles a RengineApiServer object
type RengineApiServerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=rengine.wl4g.com,resources=rengineapiservers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=rengine.wl4g.com,resources=rengineapiservers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=rengine.wl4g.com,resources=rengineapiservers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RengineApiServer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *RengineApiServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	instance := &renginev1beta1.RengineExecutor{}
	if err := r.Client.Get(ctx, req.NamespacedName, instance); err != nil {
		if k8sErrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		if instance.Spec.Image != "" {
			nl := &v1.NodeList{}
			if err := r.Client.List(ctx, nl); err != nil {
				return ctrl.Result{}, err
			}
			for _, node := range nl.Items {
				pod := v1.Pod{
					TypeMeta: metav1.TypeMeta{
						APIVersion: "v1",
						Kind:       "Pod",
					},
					ObjectMeta: metav1.ObjectMeta{
						GenerateName: fmt.Sprintf("%s-", node.Name),
					},
					Spec: v1.PodSpec{
						Containers: []v1.Container{
							{
								Image: instance.Spec.Image,
								Name:  "RengineApiServer",
							},
						},
					},
				}
				if err := r.Client.Create(ctx, &pod); err != nil {
					return ctrl.Result{}, err
				}
			}
		}

		return ctrl.Result{}, err
	}

	// return r.Do(ctx, instance)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RengineApiServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&renginev1beta1.RengineApiServer{}).
		Complete(r)
}
