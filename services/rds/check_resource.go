
package rds

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

func (client *Client) CheckResource(request *CheckResourceRequest) (response *CheckResourceResponse, err error) {
response = CreateCheckResourceResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) CheckResourceWithChan(request *CheckResourceRequest) (<-chan *CheckResourceResponse, <-chan error) {
responseChan := make(chan *CheckResourceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CheckResource(request)
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

func (client *Client) CheckResourceWithCallback(request *CheckResourceRequest, callback func(response *CheckResourceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CheckResourceResponse
var err error
defer close(result)
response, err = client.CheckResource(request)
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

type CheckResourceRequest struct {
*requests.RpcRequest
                DBInstanceId  string `position:"Query" name:"DBInstanceId"`
                Engine  string `position:"Query" name:"Engine"`
                ZoneId  string `position:"Query" name:"ZoneId"`
                DBInstanceUseType  string `position:"Query" name:"DBInstanceUseType"`
                DBInstanceClass  string `position:"Query" name:"DBInstanceClass"`
                SpecifyCount  string `position:"Query" name:"SpecifyCount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                EngineVersion  string `position:"Query" name:"EngineVersion"`
}


type CheckResourceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            SpecifyCount     string `json:"SpecifyCount" xml:"SpecifyCount"`
                Resources struct {
                    Resource []struct {
            DBInstanceAvailable     string `json:"DBInstanceAvailable" xml:"DBInstanceAvailable"`
            Engine     string `json:"Engine" xml:"Engine"`
            EngineVersion     string `json:"EngineVersion" xml:"EngineVersion"`
                    }   `json:"Resource" xml:"Resource"`
                } `json:"Resources" xml:"Resources"`
}

func CreateCheckResourceRequest() (request *CheckResourceRequest) {
request = &CheckResourceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CheckResource", "", "")
return
}

func CreateCheckResourceResponse() (response *CheckResourceResponse) {
response = &CheckResourceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

