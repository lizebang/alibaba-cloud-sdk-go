
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

func (client *Client) DescribeBinlogFiles(request *DescribeBinlogFilesRequest) (response *DescribeBinlogFilesResponse, err error) {
response = CreateDescribeBinlogFilesResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeBinlogFilesWithChan(request *DescribeBinlogFilesRequest) (<-chan *DescribeBinlogFilesResponse, <-chan error) {
responseChan := make(chan *DescribeBinlogFilesResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeBinlogFiles(request)
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

func (client *Client) DescribeBinlogFilesWithCallback(request *DescribeBinlogFilesRequest, callback func(response *DescribeBinlogFilesResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeBinlogFilesResponse
var err error
defer close(result)
response, err = client.DescribeBinlogFiles(request)
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

type DescribeBinlogFilesRequest struct {
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


type DescribeBinlogFilesResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalRecordCount     int `json:"TotalRecordCount" xml:"TotalRecordCount"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageRecordCount     int `json:"PageRecordCount" xml:"PageRecordCount"`
            TotalFileSize     int64 `json:"TotalFileSize" xml:"TotalFileSize"`
                Items struct {
                    BinLogFile []struct {
            FileSize     int64 `json:"FileSize" xml:"FileSize"`
            LogBeginTime     string `json:"LogBeginTime" xml:"LogBeginTime"`
            LogEndTime     string `json:"LogEndTime" xml:"LogEndTime"`
            DownloadLink     string `json:"DownloadLink" xml:"DownloadLink"`
            IntranetDownloadLink     string `json:"IntranetDownloadLink" xml:"IntranetDownloadLink"`
            LinkExpiredTime     string `json:"LinkExpiredTime" xml:"LinkExpiredTime"`
            Checksum     string `json:"Checksum" xml:"Checksum"`
            HostInstanceID     string `json:"HostInstanceID" xml:"HostInstanceID"`
                    }   `json:"BinLogFile" xml:"BinLogFile"`
                } `json:"Items" xml:"Items"`
}

func CreateDescribeBinlogFilesRequest() (request *DescribeBinlogFilesRequest) {
request = &DescribeBinlogFilesRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeBinlogFiles", "", "")
return
}

func CreateDescribeBinlogFilesResponse() (response *DescribeBinlogFilesResponse) {
response = &DescribeBinlogFilesResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

