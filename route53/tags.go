package route53

import (
	"fmt"
)

const (
	HealthcheckResourceType = "healthcheck"
	HostedZoneResourceType  = "hostedzones"
)

type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

type ResourceTagSet struct {
	ResourceId   string `xml:"ResourceId"`
	ResourceType string `xml:"ResourceType"`
	Tags         []Tag  `xml:"Tags>Tag"`
}

type ListTagsForResourcesRequest struct {
	ResourceIds []string `xml:"ResourceIds>ResourceId"`
}

type ListTagsForResourcesResponse struct {
	ResourceTagSets []ResourceTagSet `xml:"ResourceTagSets>ResourceTagSet"`
}

type ListTagsForResourceResponse struct {
	ResourceTagSet ResourceTagSet `xml:"ResourceTagSet"`
}

func (r *Route53) ListTagsForResource(resourceID, resourceType string) (*ListTagsForResourceResponse, error) {
	out := &ListTagsForResourceResponse{}
	err := r.query("GET", fmt.Sprintf("/%s/tags/%s/%s", APIVersion, resourceType, resourceID), nil, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *Route53) ListTagsForResources(resourceType string, req *ListTagsForResourcesRequest) (*ListTagsForResourcesResponse, error) {
	out := &ListTagsForResourcesResponse{}
	err := r.query("POST", fmt.Sprintf("/%s/tags/%s", APIVersion, resourceType), req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
