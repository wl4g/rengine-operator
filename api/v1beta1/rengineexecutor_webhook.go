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

package v1beta1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var rengineexecutorlog = logf.Log.WithName("rengineexecutor-resource")

func (r *RengineExecutor) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-rengine-wl4g-com-v1beta1-rengineexecutor,mutating=true,failurePolicy=fail,sideEffects=None,groups=rengine.wl4g.com,resources=rengineexecutors,verbs=create;update,versions=v1beta1,name=mrengineexecutor.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &RengineExecutor{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *RengineExecutor) Default() {
	rengineexecutorlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-rengine-wl4g-com-v1beta1-rengineexecutor,mutating=false,failurePolicy=fail,sideEffects=None,groups=rengine.wl4g.com,resources=rengineexecutors,verbs=create;update,versions=v1beta1,name=vrengineexecutor.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &RengineExecutor{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *RengineExecutor) ValidateCreate() error {
	rengineexecutorlog.Info("validate create", "name", r.Name)

	if r.Spec.Image == "" {
		return fmt.Errorf("image is required")
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *RengineExecutor) ValidateUpdate(old runtime.Object) error {
	rengineexecutorlog.Info("validate update", "name", r.Name)
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *RengineExecutor) ValidateDelete() error {
	rengineexecutorlog.Info("validate delete", "name", r.Name)
	return nil
}
