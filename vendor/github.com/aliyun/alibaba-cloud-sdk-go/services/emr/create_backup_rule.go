package emr

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

// CreateBackupRule invokes the emr.CreateBackupRule API synchronously
// api document: https://help.aliyun.com/api/emr/createbackuprule.html
func (client *Client) CreateBackupRule(request *CreateBackupRuleRequest) (response *CreateBackupRuleResponse, err error) {
	response = CreateCreateBackupRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBackupRuleWithChan invokes the emr.CreateBackupRule API asynchronously
// api document: https://help.aliyun.com/api/emr/createbackuprule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBackupRuleWithChan(request *CreateBackupRuleRequest) (<-chan *CreateBackupRuleResponse, <-chan error) {
	responseChan := make(chan *CreateBackupRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBackupRule(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateBackupRuleWithCallback invokes the emr.CreateBackupRule API asynchronously
// api document: https://help.aliyun.com/api/emr/createbackuprule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBackupRuleWithCallback(request *CreateBackupRuleRequest, callback func(response *CreateBackupRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBackupRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateBackupRule(request)
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

// CreateBackupRuleRequest is the request struct for api CreateBackupRule
type CreateBackupRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BackupMethodType string           `position:"Query" name:"BackupMethodType"`
	Description      string           `position:"Query" name:"Description"`
	BackupPlanId     string           `position:"Query" name:"BackupPlanId"`
	MetadataType     string           `position:"Query" name:"MetadataType"`
	Name             string           `position:"Query" name:"Name"`
}

// CreateBackupRuleResponse is the response struct for api CreateBackupRule
type CreateBackupRuleResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Id               string `json:"Id" xml:"Id"`
	Name             string `json:"Name" xml:"Name"`
	Description      string `json:"Description" xml:"Description"`
	MetadataType     string `json:"MetadataType" xml:"MetadataType"`
	BackupMethodType string `json:"BackupMethodType" xml:"BackupMethodType"`
	BackupPlanId     string `json:"BackupPlanId" xml:"BackupPlanId"`
}

// CreateCreateBackupRuleRequest creates a request to invoke CreateBackupRule API
func CreateCreateBackupRuleRequest() (request *CreateBackupRuleRequest) {
	request = &CreateBackupRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateBackupRule", "emr", "openAPI")
	return
}

// CreateCreateBackupRuleResponse creates a response to parse from CreateBackupRule response
func CreateCreateBackupRuleResponse() (response *CreateBackupRuleResponse) {
	response = &CreateBackupRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
