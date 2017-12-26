
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

func (client *Client) DescribeMigrateTasksForSQLServer(request *DescribeMigrateTasksForSQLServerRequest) (response *DescribeMigrateTasksForSQLServerResponse, err error) {
response = CreateDescribeMigrateTasksForSQLServerResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeMigrateTasksForSQLServerWithChan(request *DescribeMigrateTasksForSQLServerRequest) (<-chan *DescribeMigrateTasksForSQLServerResponse, <-chan error) {
responseChan := make(chan *DescribeMigrateTasksForSQLServerResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeMigrateTasksForSQLServer(request)
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

func (client *Client) DescribeMigrateTasksForSQLServerWithCallback(request *DescribeMigrateTasksForSQLServerRequest, callback func(response *DescribeMigrateTasksForSQLServerResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeMigrateTasksForSQLServerResponse
var err error
defer close(result)
response, err = client.DescribeMigrateTasksForSQLServer(request)
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

type DescribeMigrateTasksForSQLServerRequest struct {
*requests.RpcRequest
                EndTime  string `position:"Query" name:"EndTime"`
                PageSize  string `position:"Query" name:"PageSize"`
                DBInstanceId  string `position:"Query" name:"DBInstanceId"`
                StartTime  string `position:"Query" name:"StartTime"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type DescribeMigrateTasksForSQLServerResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceID     string `json:"DBInstanceID" xml:"DBInstanceID"`
            DBInstanceName     string `json:"DBInstanceName" xml:"DBInstanceName"`
            StartTime     string `json:"StartTime" xml:"StartTime"`
            EndTime     string `json:"EndTime" xml:"EndTime"`
            TotalRecordCount     int `json:"TotalRecordCount" xml:"TotalRecordCount"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageRecordCount     int `json:"PageRecordCount" xml:"PageRecordCount"`
                Items struct {
                    MigrateTask []struct {
            DBName     string `json:"DBName" xml:"DBName"`
            MigrateIaskId     string `json:"MigrateIaskId" xml:"MigrateIaskId"`
            CreateTime     string `json:"CreateTime" xml:"CreateTime"`
            EndTime     string `json:"EndTime" xml:"EndTime"`
            TaskType     string `json:"TaskType" xml:"TaskType"`
            Status     string `json:"Status" xml:"Status"`
            IsDBReplaced     string `json:"IsDBReplaced" xml:"IsDBReplaced"`
            Desc     string `json:"Desc" xml:"Desc"`
                    }   `json:"MigrateTask" xml:"MigrateTask"`
                } `json:"Items" xml:"Items"`
}

func CreateDescribeMigrateTasksForSQLServerRequest() (request *DescribeMigrateTasksForSQLServerRequest) {
request = &DescribeMigrateTasksForSQLServerRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeMigrateTasksForSQLServer", "", "")
return
}

func CreateDescribeMigrateTasksForSQLServerResponse() (response *DescribeMigrateTasksForSQLServerResponse) {
response = &DescribeMigrateTasksForSQLServerResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

