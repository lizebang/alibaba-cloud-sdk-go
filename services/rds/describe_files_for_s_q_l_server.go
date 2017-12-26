
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

func (client *Client) DescribeFilesForSQLServer(request *DescribeFilesForSQLServerRequest) (response *DescribeFilesForSQLServerResponse, err error) {
response = CreateDescribeFilesForSQLServerResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeFilesForSQLServerWithChan(request *DescribeFilesForSQLServerRequest) (<-chan *DescribeFilesForSQLServerResponse, <-chan error) {
responseChan := make(chan *DescribeFilesForSQLServerResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeFilesForSQLServer(request)
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

func (client *Client) DescribeFilesForSQLServerWithCallback(request *DescribeFilesForSQLServerRequest, callback func(response *DescribeFilesForSQLServerResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeFilesForSQLServerResponse
var err error
defer close(result)
response, err = client.DescribeFilesForSQLServer(request)
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

type DescribeFilesForSQLServerRequest struct {
*requests.RpcRequest
                EndTime  string `position:"Query" name:"EndTime"`
                PageSize  string `position:"Query" name:"PageSize"`
                DBInstanceId  string `position:"Query" name:"DBInstanceId"`
                StartTime  string `position:"Query" name:"StartTime"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type DescribeFilesForSQLServerResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            TotalRecordCount     int `json:"TotalRecordCount" xml:"TotalRecordCount"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageRecordCount     int `json:"PageRecordCount" xml:"PageRecordCount"`
                Items struct {
                    SQLServerUploadFile []struct {
            DBName     string `json:"DBName" xml:"DBName"`
            FileName     string `json:"FileName" xml:"FileName"`
            FileSize     int64 `json:"FileSize" xml:"FileSize"`
            InternetFtpServer     string `json:"InternetFtpServer" xml:"InternetFtpServer"`
            InternetPort     int `json:"InternetPort" xml:"InternetPort"`
            IntranetFtpserver     string `json:"IntranetFtpserver" xml:"IntranetFtpserver"`
            Intranetport     int `json:"Intranetport" xml:"Intranetport"`
            UserName     string `json:"UserName" xml:"UserName"`
            Password     string `json:"Password" xml:"Password"`
            FileStatus     string `json:"FileStatus" xml:"FileStatus"`
            Description     string `json:"Description" xml:"Description"`
            CreationTime     string `json:"CreationTime" xml:"CreationTime"`
                    }   `json:"SQLServerUploadFile" xml:"SQLServerUploadFile"`
                } `json:"Items" xml:"Items"`
}

func CreateDescribeFilesForSQLServerRequest() (request *DescribeFilesForSQLServerRequest) {
request = &DescribeFilesForSQLServerRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeFilesForSQLServer", "", "")
return
}

func CreateDescribeFilesForSQLServerResponse() (response *DescribeFilesForSQLServerResponse) {
response = &DescribeFilesForSQLServerResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

