/*
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
package options

import (
	"github.com/tejabeta/kopy/internal/context"

	"k8s.io/client-go/rest"
)

type KopyOptions struct {
	Namespace          string
	AllResource        bool
	SourceContext      *rest.Config
	DestinationContext *rest.Config
}

func GetKopyOptions(sourceCName string, destCName string) (*KopyOptions, error) {

	var sContext, dContext *rest.Config
	var err error

	if sourceCName == "" {
		sContext, err = context.GetContext()
		if err != nil {
			return nil, err
		}
	} else {
		sContext, err = context.SwitchContext(sourceCName)
		if err != nil {
			return nil, err
		}
	}

	dContext, err = context.SwitchContext(destCName)
	if err != nil {
		return nil, err
	}

	return &KopyOptions{
		SourceContext:      sContext,
		DestinationContext: dContext,
	}, err
}
