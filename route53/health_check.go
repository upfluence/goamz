package route53

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type UpdateHealthCheckRequest struct {
	IPAddress                string `xml:"IPAddress,omitempty"`
	Port                     string `xml:"Port,omitempty"`
	ResourcePath             string `xml:"ResourcePath,omitempty"`
	FullyQualifiedDomainName string `xml:"FullyQualifiedDomainName,omitempty"`
	SearchString             string `xml:"SearchString,omitempty"`
	FailureThreshold         uint64 `xml:"FailureThreshold"`
}

type UpdateHealthCheckResponse struct {
	HealthCheck HealthCheck `xml:"HealthCheck"`
}

type CreateHealthCheckRequest struct {
	CallerReference   string            `xml:"CallerReference"`
	HealthCheckConfig HealthCheckConfig `xml:"HealthCheckConfig"`
}

type CreateHealthCheckResponse struct {
	HealthCheck HealthCheck `xml:"HealthCheck"`
}

type HealthCheckConfig struct {
	IPAddress                string `xml:"IPAddress,omitempty"`
	Port                     string `xml:"Port,omitempty"`
	Type                     string `xml:"Type"`
	ResourcePath             string `xml:"ResourcePath,omitempty"`
	FullyQualifiedDomainName string `xml:"FullyQualifiedDomainName,omitempty"`
	SearchString             string `xml:"SearchString,omitempty"`
	RequestInterval          uint64 `xml:"RequestInterval"`
	FailureThreshold         uint64 `xml:"FailureThreshold"`
}

type HealthCheck struct {
	ID                 string            `xml:"Id"`
	CallerReference    string            `xml:"CallerReference"`
	HealthCheckConfig  HealthCheckConfig `xml:"HealthCheckConfig"`
	HealthCheckVersion uint64            `xml:"HealthCheckVersion"`
}

func (r *Route53) UpdateHealthCheck(id string, req *UpdateHealthCheckRequest) (*UpdateHealthCheckResponse, error) {
	out := &UpdateHealthCheckResponse{}
	if err := r.query("POST", fmt.Sprintf("/%s/healthcheck/%s", APIVersion, id), req, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *Route53) CreateHealthCheck(req *CreateHealthCheckRequest) (*CreateHealthCheckResponse, error) {
	// Generate a unique caller reference if none provided
	if req.CallerReference == "" {
		req.CallerReference = time.Now().Format(time.RFC3339Nano)
	}
	out := &CreateHealthCheckResponse{}
	if err := r.query("POST", fmt.Sprintf("/%s/healthcheck", APIVersion), req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ListHealthChecksResponse struct {
	HealthChecks []HealthCheck `xml:"HealthChecks>HealthCheck"`
	Marker       string        `xml:"Marker"`
	IsTruncated  bool          `xml:"IsTruncated"`
	NextMarker   string        `xml:"NextMarker"`
	MaxItems     int           `xml:"MaxItems"`
}

func (r *Route53) ListHealthChecks(marker string, maxItems int) (*ListHealthChecksResponse, error) {
	values := url.Values{}

	if marker != "" {
		values.Add("marker", marker)
	}

	if maxItems != 0 {
		values.Add("maxItems", strconv.Itoa(maxItems))
	}

	out := &ListHealthChecksResponse{}
	err := r.query("GET", fmt.Sprintf("/%s/healthcheck/", APIVersion), values, out)
	if err != nil {
		return nil, err
	}
	return out, err
}

type GetHealthCheckResponse struct {
	HealthCheck HealthCheck `xml:"HealthCheck"`
}

func (r *Route53) GetHealthCheck(ID string) (*GetHealthCheckResponse, error) {
	out := &GetHealthCheckResponse{}
	err := r.query("GET", fmt.Sprintf("/%s/healthcheck/%s", APIVersion, ID), nil, out)
	if err != nil {
		return nil, err
	}
	return out, err
}
