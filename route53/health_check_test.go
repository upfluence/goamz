package route53

import (
	"testing"
)

func TestCreateHealthCheck(t *testing.T) {
	testServer := makeTestServer()
	client := makeClient(testServer)
	testServer.Response(201, nil, CreateHealthCheckExample)

	req := &CreateHealthCheckRequest{
		CallerReference: "foo",
		HealthCheckConfig: HealthCheckConfig{
			IPAddress:                "127.0.0.1",
			Port:                     "80",
			Type:                     "HTTP",
			ResourcePath:             "/",
			FullyQualifiedDomainName: "upfluence.co",
		},
	}

	resp, err := client.CreateHealthCheck(req)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if resp.HealthCheck.ID != "abcdef11-2222-3333-4444-555555fedcba" {
		t.Fatalf("bad: %v", resp)
	}
	if resp.HealthCheck.HealthCheckConfig.ResourcePath != "/docs/route-53-health-check.html" {
		t.Fatalf("bad: %v", resp)
	}
	if resp.HealthCheck.HealthCheckConfig.FullyQualifiedDomainName != "example.com" {
		t.Fatalf("bad: %v", resp)
	}

	httpReq := testServer.WaitRequest()
	if httpReq.URL.Path != "/2013-04-01/healthcheck" {
		t.Fatalf("bad: %#v", httpReq)
	}
	if httpReq.Method != "POST" {
		t.Fatalf("bad: %#v", httpReq)
	}
	if httpReq.ContentLength == 0 {
		t.Fatalf("bad: %#v", httpReq)
	}
}

func TestGetHealthCheck(t *testing.T) {
	testServer := makeTestServer()
	client := makeClient(testServer)
	testServer.Response(200, nil, GetHealthCheckExample)

	resp, err := client.GetHealthCheck("foobarbaz")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if resp.HealthCheck.HealthCheckConfig.IPAddress != "192.0.2.17" {
		t.Fatalf("bad: %v", resp)
	}
}

func TestListHealthChecks(t *testing.T) {
	testServer := makeTestServer()
	client := makeClient(testServer)
	testServer.Response(200, nil, ListHealthChecksExample)

	resp, err := client.ListHealthChecks("", 0)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if resp.HealthChecks[0].HealthCheckConfig.IPAddress != "192.0.2.17" {
		t.Fatalf("bad: %v", resp)
	}
}
