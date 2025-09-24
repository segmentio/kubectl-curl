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
				NewPath:      "",
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
				NewPath:      "",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := url.Parse(tt.urlStr)
			if err != nil {
				t.Fatalf("url.Parse failed: %v", err)
			}
			got := ParseResourceTarget(u, resourceTypeMap)
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
