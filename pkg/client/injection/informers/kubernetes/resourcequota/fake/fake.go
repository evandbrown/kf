/*
Copyright 2019 The Knative Authors
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

package fake

import (
	"context"

	resourcequota "github.com/google/kf/pkg/client/injection/informers/kubernetes/resourcequota"

	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection"

	"knative.dev/pkg/injection/informers/kubeinformers/factory/fake"
)

var Get = resourcequota.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Core().V1().ResourceQuotas()
	return context.WithValue(ctx, resourcequota.Key{}, inf), inf.Informer()
}
