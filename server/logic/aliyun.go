package logic

import (
	"app/global"
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"time"
)

func AliyunClient() (client *sdk.Client, err error) {
	return sdk.NewClientWithAccessKey(global.Config.AdminConfig.Aliyun.RegionId, global.Config.AdminConfig.Aliyun.AccessKeyId, global.Config.AdminConfig.Aliyun.AccessKeySecret)
}

func AliyunRequest(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
	client, err := AliyunClient()
	if err != nil {
		return nil, err
	}
	return client.ProcessCommonRequest(request)
}

type AliyunListCdtInternetTrafficResponse struct {
	RequestId      string `json:"RequestId"`
	TrafficDetails []struct {
		Traffic            int `json:"Traffic"`
		TrafficTierDetails []struct {
			HighestTraffic int64 `json:"HighestTraffic"`
			Tier           int   `json:"Tier"`
			Traffic        int   `json:"Traffic"`
			LowestTraffic  int   `json:"LowestTraffic"`
		} `json:"TrafficTierDetails"`
		ISPType               string `json:"ISPType"`
		BusinessRegionId      string `json:"BusinessRegionId"`
		ProductTrafficDetails []struct {
			Traffic int    `json:"Traffic"`
			Product string `json:"Product"`
		} `json:"ProductTrafficDetails"`
	} `json:"TrafficDetails"`
}

// AliyunListCdtInternetTraffic 获取已使用流量
func AliyunListCdtInternetTraffic() (float64, error) {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Product = "CDT"
	request.Domain = "cdt.aliyuncs.com"
	request.Version = "2021-08-13"
	request.ApiName = "ListCdtInternetTraffic"

	result, err := AliyunRequest(request)
	if err != nil {
		return 0, err
	}
	response := &AliyunListCdtInternetTrafficResponse{}
	if err := json.Unmarshal(result.GetHttpContentBytes(), response); err != nil {
		return 0, err
	}

	flow := 0
	for _, v := range response.TrafficDetails {
		flow = flow + v.Traffic
	}
	return float64(flow) / 1024 / 1024 / 1024, nil
}

type AliyunDescribeSecurityGroupAttributeResponse struct {
	Description       string `json:"Description"`
	VpcId             string `json:"VpcId"`
	RequestId         string `json:"RequestId"`
	SecurityGroupName string `json:"SecurityGroupName"`
	SecurityGroupId   string `json:"SecurityGroupId"`
	Permissions       struct {
		Permission []struct {
			SourceGroupId           string    `json:"SourceGroupId"`
			Policy                  string    `json:"Policy"`
			Description             string    `json:"Description"`
			SourcePortRange         string    `json:"SourcePortRange"`
			Priority                int       `json:"Priority"`
			CreateTime              time.Time `json:"CreateTime"`
			DestPrefixListName      string    `json:"DestPrefixListName"`
			Ipv6SourceCidrIp        string    `json:"Ipv6SourceCidrIp"`
			NicType                 string    `json:"NicType"`
			Direction               string    `json:"Direction"`
			DestGroupId             string    `json:"DestGroupId"`
			SourceGroupName         string    `json:"SourceGroupName"`
			PortRange               string    `json:"PortRange"`
			DestGroupOwnerAccount   string    `json:"DestGroupOwnerAccount"`
			DestPrefixListId        string    `json:"DestPrefixListId"`
			SourceCidrIp            string    `json:"SourceCidrIp"`
			SourcePrefixListName    string    `json:"SourcePrefixListName"`
			IpProtocol              string    `json:"IpProtocol"`
			SecurityGroupRuleId     string    `json:"SecurityGroupRuleId"`
			DestCidrIp              string    `json:"DestCidrIp"`
			DestGroupName           string    `json:"DestGroupName"`
			Ipv6DestCidrIp          string    `json:"Ipv6DestCidrIp"`
			SourceGroupOwnerAccount string    `json:"SourceGroupOwnerAccount"`
			SourcePrefixListId      string    `json:"SourcePrefixListId"`
		} `json:"Permission"`
	} `json:"Permissions"`
	InnerAccessPolicy string `json:"InnerAccessPolicy"`
	RegionId          string `json:"RegionId"`
}

// AliyunDescribeSecurityGroupAttribute 获取安全组规则状态
func AliyunDescribeSecurityGroupAttribute() (bool, error) {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Product = "Ecs"
	request.Domain = "ecs." + global.Config.AdminConfig.Aliyun.RegionId + ".aliyuncs.com"
	request.Version = "2014-05-26"
	request.ApiName = "DescribeSecurityGroupAttribute"
	request.QueryParams["RegionId"] = global.Config.AdminConfig.Aliyun.RegionId
	request.QueryParams["SecurityGroupId"] = global.Config.AdminConfig.Aliyun.SecurityGroupId
	result, err := AliyunRequest(request)
	if err != nil {
		return false, err
	}
	response := &AliyunDescribeSecurityGroupAttributeResponse{}
	if err := json.Unmarshal(result.GetHttpContentBytes(), response); err != nil {
		return false, err
	}

	for _, v := range response.Permissions.Permission {
		if v.IpProtocol == "ALL" && v.SourceCidrIp == "0.0.0.0/0" && v.Policy == "Accept" && v.NicType == "intranet" && v.Direction == "ingress" {
			return true, nil
		}
	}
	return false, nil
}

// AliyunRevokeSecurityGroup 移除安全组规则
func AliyunRevokeSecurityGroup() (bool, error) {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Product = "Ecs"
	request.Domain = "ecs." + global.Config.AdminConfig.Aliyun.RegionId + ".aliyuncs.com"
	request.Version = "2014-05-26"
	request.ApiName = "RevokeSecurityGroup"
	request.QueryParams["RegionId"] = global.Config.AdminConfig.Aliyun.RegionId
	request.QueryParams["SecurityGroupId"] = global.Config.AdminConfig.Aliyun.SecurityGroupId
	request.QueryParams["IpProtocol"] = "all"
	request.QueryParams["PortRange"] = "-1/-1"
	request.QueryParams["SourceCidrIp"] = "0.0.0.0/0"
	_, err := AliyunRequest(request)
	if err != nil {
		return false, err
	}
	return true, nil
}

// AliyunAuthorizeSecurityGroup 添加安全组规则
func AliyunAuthorizeSecurityGroup() (bool, error) {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Product = "Ecs"
	request.Domain = "ecs." + global.Config.AdminConfig.Aliyun.RegionId + ".aliyuncs.com"
	request.Version = "2014-05-26"
	request.ApiName = "AuthorizeSecurityGroup"
	request.QueryParams["RegionId"] = global.Config.AdminConfig.Aliyun.RegionId
	request.QueryParams["SecurityGroupId"] = global.Config.AdminConfig.Aliyun.SecurityGroupId
	request.QueryParams["IpProtocol"] = "all"
	request.QueryParams["PortRange"] = "-1/-1"
	request.QueryParams["SourceCidrIp"] = "0.0.0.0/0"
	_, err := AliyunRequest(request)
	if err != nil {
		return false, err
	}
	return true, nil
}
