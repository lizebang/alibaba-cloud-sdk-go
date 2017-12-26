
package ram

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

func (client *Client) UpdateRole(request *UpdateRoleRequest) (response *UpdateRoleResponse, err error) {
response = CreateUpdateRoleResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) UpdateRoleWithChan(request *UpdateRoleRequest) (<-chan *UpdateRoleResponse, <-chan error) {
responseChan := make(chan *UpdateRoleResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.UpdateRole(request)
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

func (client *Client) UpdateRoleWithCallback(request *UpdateRoleRequest, callback func(response *UpdateRoleResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *UpdateRoleResponse
var err error
defer close(result)
response, err = client.UpdateRole(request)
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

type UpdateRoleRequest struct {
*requests.RpcRequest
                NewAssumeRolePolicyDocument  string `position:"Query" name:"NewAssumeRolePolicyDocument"`
                RoleName  string `position:"Query" name:"RoleName"`
}


type UpdateRoleResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Role struct {
            RoleId     string `json:"RoleId" xml:"RoleId"`
            RoleName     string `json:"RoleName" xml:"RoleName"`
            Arn     string `json:"Arn" xml:"Arn"`
            Description     string `json:"Description" xml:"Description"`
            AssumeRolePolicyDocument     string `json:"AssumeRolePolicyDocument" xml:"AssumeRolePolicyDocument"`
            CreateDate     string `json:"CreateDate" xml:"CreateDate"`
            UpdateDate     string `json:"UpdateDate" xml:"UpdateDate"`
            }  `json:"Role" xml:"Role"`
}

func CreateUpdateRoleRequest() (request *UpdateRoleRequest) {
request = &UpdateRoleRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ram", "2015-05-01", "UpdateRole", "", "")
return
}

func CreateUpdateRoleResponse() (response *UpdateRoleResponse) {
response = &UpdateRoleResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

