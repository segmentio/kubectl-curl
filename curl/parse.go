package curl

import (
	"net/url"
	"strings"
)

type ResourceTarget struct {
	IsResource   bool
	ResourceType string
	ResourceName string
	PodName      string
	PodPort      string
	NewPath      string
}

// ParseResourceTarget parses the URL and returns resource/pod targeting info.
func ParseResourceTarget(requestURL *url.URL, resourceTypeMap map[string]string) ResourceTarget {
	hostPort := requestURL.Host
	var podName, podPort string
	var resourceType, resourceName string
	isResource := false
	newPath := requestURL.Path

	if canonicalType, ok := resourceTypeMap[strings.ToLower(hostPort)]; ok && requestURL.Path != "" {
		segments := strings.SplitN(strings.TrimLeft(requestURL.Path, "/"), "/", 2)
		resourceAndMaybePort := segments[0]
		resourceType = canonicalType
		if colonIdx := strings.LastIndex(resourceAndMaybePort, ":"); colonIdx > -1 {
			resourceName = resourceAndMaybePort[:colonIdx]
			podPort = resourceAndMaybePort[colonIdx+1:]
		} else {
			resourceName = resourceAndMaybePort
		}
		isResource = true
		if len(segments) > 1 {
			newPath = "/" + segments[1]
		} else {
			newPath = "/"
		}
	} else if idx := strings.Index(hostPort, "/"); idx >= 0 {
		resourceAndMaybePort := hostPort
		resource := resourceAndMaybePort
		if colonIdx := strings.LastIndex(resourceAndMaybePort, ":"); colonIdx > -1 && colonIdx > idx {
			resource = resourceAndMaybePort[:colonIdx]
			podPort = resourceAndMaybePort[colonIdx+1:]
		}
		resourceParts := strings.SplitN(resource, "/", 2)
		if len(resourceParts) == 2 {
			resourceType, resourceName = resourceParts[0], resourceParts[1]
			isResource = true
		}
	} else {
		// podname[:port]
		parts := strings.SplitN(hostPort, ":", 2)
		podName = parts[0]
		if len(parts) == 2 {
			podPort = parts[1]
		} else {
			podPort = ""
		}
	}

	return ResourceTarget{
		IsResource:   isResource,
		ResourceType: resourceType,
		ResourceName: resourceName,
		PodName:      podName,
		PodPort:      podPort,
		NewPath:      newPath,
	}
}
