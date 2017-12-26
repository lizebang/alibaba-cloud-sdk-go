
package ecs

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

func (client *Client) ModifyAutoSnapshotPolicyEx(request *ModifyAutoSnapshotPolicyExRequest) (response *ModifyAutoSnapshotPolicyExResponse, err error) {
response = CreateModifyAutoSnapshotPolicyExResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ModifyAutoSnapshotPolicyExWithChan(request *ModifyAutoSnapshotPolicyExRequest) (<-chan *ModifyAutoSnapshotPolicyExResponse, <-chan error) {
responseChan := make(chan *ModifyAutoSnapshotPolicyExResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyAutoSnapshotPolicyEx(request)
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

func (client *Client) ModifyAutoSnapshotPolicyExWithCallback(request *ModifyAutoSnapshotPolicyExRequest, callback func(response *ModifyAutoSnapshotPolicyExResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyAutoSnapshotPolicyExResponse
var err error
defer close(result)
response, err = client.ModifyAutoSnapshotPolicyEx(request)
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

type ModifyAutoSnapshotPolicyExRequest struct {
*requests.RpcRequest
                RepeatWeekdays  string `position:"Query" name:"repeatWeekdays"`
                AutoSnapshotPolicyName  string `position:"Query" name:"autoSnapshotPolicyName"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                TimePoints  string `position:"Query" name:"timePoints"`
                AutoSnapshotPolicyId  string `position:"Query" name:"autoSnapshotPolicyId"`
                RetentionDays  string `position:"Query" name:"retentionDays"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type ModifyAutoSnapshotPolicyExResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyAutoSnapshotPolicyExRequest() (request *ModifyAutoSnapshotPolicyExRequest) {
request = &ModifyAutoSnapshotPolicyExRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyAutoSnapshotPolicyEx", "", "")
return
}

func CreateModifyAutoSnapshotPolicyExResponse() (response *ModifyAutoSnapshotPolicyExResponse) {
response = &ModifyAutoSnapshotPolicyExResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

