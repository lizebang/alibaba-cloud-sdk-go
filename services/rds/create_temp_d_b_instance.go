
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

func (client *Client) CreateTempDBInstance(request *CreateTempDBInstanceRequest) (response *CreateTempDBInstanceResponse, err error) {
response = CreateCreateTempDBInstanceResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) CreateTempDBInstanceWithChan(request *CreateTempDBInstanceRequest) (<-chan *CreateTempDBInstanceResponse, <-chan error) {
responseChan := make(chan *CreateTempDBInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateTempDBInstance(request)
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

func (client *Client) CreateTempDBInstanceWithCallback(request *CreateTempDBInstanceRequest, callback func(response *CreateTempDBInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateTempDBInstanceResponse
var err error
defer close(result)
response, err = client.CreateTempDBInstance(request)
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

type CreateTempDBInstanceRequest struct {
*requests.RpcRequest
                DBInstanceId  string `position:"Query" name:"DBInstanceId"`
                RestoreTime  string `position:"Query" name:"RestoreTime"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                BackupId  string `position:"Query" name:"BackupId"`
}


type CreateTempDBInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TempDBInstanceId     string `json:"TempDBInstanceId" xml:"TempDBInstanceId"`
}

func CreateCreateTempDBInstanceRequest() (request *CreateTempDBInstanceRequest) {
request = &CreateTempDBInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CreateTempDBInstance", "", "")
return
}

func CreateCreateTempDBInstanceResponse() (response *CreateTempDBInstanceResponse) {
response = &CreateTempDBInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

