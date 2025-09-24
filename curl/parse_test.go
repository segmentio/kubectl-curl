package curl

import (
	"net/url"
	"testing"
)

var resourceTypeMap = map[string]string{
	"deployment":   "deployment",
	"deploy":       "deployment",
	"deployments":  "deployment",
	"ds":           "daemonset",
	"daemonset":    "daemonset",
	"daemonsets":   "daemonset",
	"sts":          "statefulset",
	"statefulset":  "statefulset",
	"statefulsets": "statefulset",
}

func TestParseResourceTarget(t *testing.T) {
	// Mock resource name sets for testing fallback logic
	pods := map[string]bool{"mypod": true, "foobar": false, "podonly": true}
	deployments := map[string]bool{"mydeploy": true, "foobar": true, "deployonly": true}
	statefulsets := map[string]bool{"mysts": true, "foobar": true, "stsonly": true}
	daemonsets := map[string]bool{"myds": true, "foobar": true, "dsonly": true}

	isPodName := func(name string) bool { return pods[name] }
	isDeploymentName := func(name string) bool { return deployments[name] }
	isStatefulSetName := func(name string) bool { return statefulsets[name] }
	isDaemonSetName := func(name string) bool { return daemonsets[name] }

	tests := []struct {
		name   string
		urlStr string
		want   ResourceTarget
	}{
		{
			name:   "podname only",
			urlStr: "http://mypod:8080/path",
			want: ResourceTarget{
				IsResource: false,
				PodName:    "mypod",
				PodPort:    "8080",
				NewPath:    "/path",
			},
		},
		{
			name:   "deployment as host, name:port as path",
			urlStr: "http://deployment/mydeploy:3000/foo",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "deployment",
				ResourceName: "mydeploy",
				PodPort:      "3000",
				NewPath:      "/foo",
			},
		},
		{
			name:   "daemonset abbreviation as host, name as path",
			urlStr: "http://ds/myds/bar",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "daemonset",
				ResourceName: "myds",
				PodPort:      "",
				NewPath:      "/bar",
			},
		},
		{
			name:   "statefulset as host, name:port as path",
			urlStr: "http://statefulset/mysts:1234/",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "statefulset",
				ResourceName: "mysts",
				PodPort:      "1234",
				NewPath:      "/",
			},
		},
		{
			name:   "type/name:port in host",
			urlStr: "http://deployment/mydeploy:3000",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "deployment",
				ResourceName: "mydeploy",
				PodPort:      "3000",
				NewPath:      "/",
			},
		},
		{
			name:   "type/name in host, no port",
			urlStr: "http://daemonset/myds",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "daemonset",
				ResourceName: "myds",
				PodPort:      "",
				NewPath:      "/",
			},
		},
		{
			name:   "podname only, no port",
			urlStr: "http://mypod",
			want: ResourceTarget{
				IsResource: false,
				PodName:    "mypod",
				PodPort:    "",
				NewPath:    "",
			},
		},
		{
			name:   "podname:port, no scheme",
			urlStr: "mypod:8080",
			want: ResourceTarget{
				IsResource: false,
				PodName:    "mypod",
				PodPort:    "8080",
				NewPath:    "",
			},
		},
		{
			name:   "resourceType/resourceName:port, no scheme",
			urlStr: "deployment/mydeploy:3000",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "deployment",
				ResourceName: "mydeploy",
				PodPort:      "3000",
				NewPath:      "",
			},
		},
		{
			name:   "resourceType/resourceName, no port, no scheme",
			urlStr: "ds/myds",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "daemonset",
				ResourceName: "myds",
				PodPort:      "",
				NewPath:      "",
			},
		},
		{
			name:   "fallback: pod preferred over deployment",
			urlStr: "podonly",
			want: ResourceTarget{
				IsResource: false,
				PodName:    "podonly",
				PodPort:    "",
				NewPath:    "",
			},
		},
		{
			name:   "fallback: deployment preferred over statefulset",
			urlStr: "deployonly",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "deployment",
				ResourceName: "deployonly",
				PodPort:      "",
				NewPath:      "",
			},
		},
		{
			name:   "fallback: statefulset preferred over daemonset",
			urlStr: "stsonly",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "statefulset",
				ResourceName: "stsonly",
				PodPort:      "",
				NewPath:      "",
			},
		},
		{
			name:   "fallback: daemonset if only match",
			urlStr: "dsonly",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "daemonset",
				ResourceName: "dsonly",
				PodPort:      "",
				NewPath:      "",
			},
		},
		{
			name:   "fallback: deployment preferred over statefulset and daemonset (foobar)",
			urlStr: "foobar",
			want: ResourceTarget{
				IsResource:   true,
				ResourceType: "deployment",
				ResourceName: "foobar",
				PodPort:      "",
				NewPath:      "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := error(nil)
			var u *url.URL
			if tt.urlStr == "mypod:8080" {
				u = &url.URL{Host: "mypod:8080"}
			} else {
				u, err = url.Parse(tt.urlStr)
				if err != nil {
					t.Fatalf("url.Parse failed: %v", err)
				}
			}
			if tt.urlStr == "mypod:8080" {
				t.Logf("DEBUG: url.Parse(%q) => Host=%q, Path=%q, Opaque=%q, RawPath=%q", tt.urlStr, u.Host, u.Path, u.Opaque, u.RawPath)
				t.Logf("DEBUG: isPodName('mypod') = %v", isPodName("mypod"))
			}
			got := ParseResourceTarget(u, resourceTypeMap, isPodName, isDeploymentName, isStatefulSetName, isDaemonSetName, false)
			if got.IsResource != tt.want.IsResource ||
				got.ResourceType != tt.want.ResourceType ||
				got.ResourceName != tt.want.ResourceName ||
				got.PodName != tt.want.PodName ||
				got.PodPort != tt.want.PodPort ||
				got.NewPath != tt.want.NewPath {
				t.Errorf("ParseResourceTarget(%q) = %+v, want %+v", tt.urlStr, got, tt.want)
			}
		})
	}
}
