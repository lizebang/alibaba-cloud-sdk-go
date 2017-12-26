
package cs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
response = CreateCreateClusterResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) CreateClusterWithChan(request *CreateClusterRequest) (<-chan *CreateClusterResponse, <-chan error) {
responseChan := make(chan *CreateClusterResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateCluster(request)
responseChan <- response
errChan <- err
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

func (client *Client) CreateClusterWithCallback(request *CreateClusterRequest, callback func(response *CreateClusterResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateClusterResponse
var err error
defer close(result)
response, err = client.CreateCluster(request)
callback(response, err)
result <- 1
})
if err != nil {
defer close(result)
callback(nil, err)
result <- 0
}
return result
}

type CreateClusterRequest struct {
*requests.RoaRequest
}


type CreateClusterResponse struct {
*responses.BaseResponse
}

func CreateCreateClusterRequest() (request *CreateClusterRequest) {
request = &CreateClusterRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("CS", "2015-12-15", "CreateCluster", "/clusters", "", "")
return
}

func CreateCreateClusterResponse() (response *CreateClusterResponse) {
response = &CreateClusterResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

