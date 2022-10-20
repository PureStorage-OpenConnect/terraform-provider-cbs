/*

	Copyright 2021, Pure Storage Inc.

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

package appcatalog

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type SearchClient struct {
	client searchClient
}

func NewSearchClient() *SearchClient {
	client_options := policy.ClientOptions{}
	client_options.Telemetry.Disabled = true
	pipeline := runtime.NewPipeline("", "", runtime.PipelineOptions{}, &client_options)

	client := newSearchClient(pipeline)

	return &SearchClient{
		client: *client,
	}
}

func (client *SearchClient) Get(ctx context.Context, language string, market string, xmsApp string, selectParam []SearchV2FieldName, publisherIds []string) (SearchClientGetResponse, error) {
	search_options := searchClientGetOptions{}
	search_options.PublisherIDs = publisherIds
	return client.client.Get(ctx, language, selectParam, market, xmsApp, &search_options)
}
