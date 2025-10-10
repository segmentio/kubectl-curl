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
// Added: pod/resource lookup functions and verbosity parameter.
func ParseResourceTarget(
	requestURL *url.URL,
	resourceTypeMap map[string]string,
	isPodName func(string) bool,
	isDeploymentName func(string) bool,
	isStatefulSetName func(string) bool,
	isDaemonSetName func(string) bool,
	verbose bool,
) ResourceTarget {
	if verbose {
		println("[kubectl-curl] ParseResourceTarget called with Host=", requestURL.Host, ", Path=", requestURL.Path)
	}
	hostPort := requestURL.Host
	newPath := requestURL.Path

	var podName, podPort string
	var resourceType, resourceName string
	isResource := false

	// Unify podname:port handling for both host and hostless cases
	if hostPort == "" && newPath != "" {
		if strings.HasPrefix(newPath, "/") {
			newPath = newPath[1:]
		}
		hostPort = newPath
		newPath = ""
	}

	// 1. If host matches a resource type, parse path for resource name and port
	if canonicalType, ok := resourceTypeMap[strings.ToLower(hostPort)]; ok {
		if requestURL.Path != "" {
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
			if verbose {
				println("[kubectl-curl] Returning ResourceTarget:", resourceType, resourceName, podName, podPort, newPath)
			}
			return ResourceTarget{
				IsResource:   isResource,
				ResourceType: resourceType,
				ResourceName: resourceName,
				PodName:      "",
				PodPort:      podPort,
				NewPath:      newPath,
			}
		} else {
			// If path is empty, treat as fallback (e.g., just "deployment")
			resourceType = canonicalType
			isResource = true
			return ResourceTarget{
				IsResource:   isResource,
				ResourceType: resourceType,
				ResourceName: "",
				PodName:      "",
				PodPort:      "",
				NewPath:      "",
			}
		}
	}

	// 2. If host is type/name:port or type/name, parse accordingly
	if idx := strings.Index(hostPort, "/"); idx >= 0 {
		resourceAndMaybePort := hostPort
		resource := resourceAndMaybePort
		if colonIdx := strings.LastIndex(resourceAndMaybePort, ":"); colonIdx > -1 && colonIdx > idx {
			resource = resourceAndMaybePort[:colonIdx]
			podPort = resourceAndMaybePort[colonIdx+1:]
		}
		resourceParts := strings.SplitN(resource, "/", 2)
		if len(resourceParts) == 2 {
			resourceType, resourceName = resourceParts[0], resourceParts[1]
			if canonicalType, ok := resourceTypeMap[strings.ToLower(resourceType)]; ok {
				resourceType = canonicalType
				isResource = true
			}
			newPath = ""
		}
		return ResourceTarget{
			IsResource:   isResource,
			ResourceType: resourceType,
			ResourceName: resourceName,
			PodName:      "",
			PodPort:      podPort,
			NewPath:      newPath,
		}
	}

	// 3. If host is empty and path is set, treat as hostless input (e.g., mypod:8080 or deployment/mydeploy:3000)
	if hostPort == "" && newPath != "" {
		if strings.HasPrefix(newPath, "/") {
			newPath = newPath[1:]
		}
		hostPort = newPath
		newPath = ""
		// Instead of recursing, handle directly:
		parts := strings.SplitN(hostPort, ":", 2)
		name := parts[0]
		if len(parts) == 2 {
			podPort = parts[1]
		} else {
			podPort = ""
		}
		if isPodName != nil && isPodName(name) {
			podName = name
			if verbose {
				println("[kubectl-curl] Found pod ", podName, " port=", podPort)
			}
			return ResourceTarget{
				IsResource:   false,
				ResourceType: "",
				ResourceName: "",
				PodName:      podName,
				PodPort:      podPort,
				NewPath:      "",
			}
		} else if isDeploymentName != nil && isDeploymentName(name) {
			resourceType = "deployment"
			resourceName = name
			isResource = true
			if verbose {
				println("[kubectl-curl] Found deployment/", resourceName)
			}
			return ResourceTarget{
				IsResource:   isResource,
				ResourceType: resourceType,
				ResourceName: resourceName,
				PodName:      "",
				PodPort:      podPort,
				NewPath:      "",
			}
		} else if isStatefulSetName != nil && isStatefulSetName(name) {
			resourceType = "statefulset"
			resourceName = name
			isResource = true
			if verbose {
				println("[kubectl-curl] Found statefulset/", resourceName)
			}
			return ResourceTarget{
				IsResource:   isResource,
				ResourceType: resourceType,
				ResourceName: resourceName,
				PodName:      "",
				PodPort:      podPort,
				NewPath:      "",
			}
		} else if isDaemonSetName != nil && isDaemonSetName(name) {
			resourceType = "daemonset"
			resourceName = name
			isResource = true
			if verbose {
				println("[kubectl-curl] Found daemonset/", resourceName)
			}
			return ResourceTarget{
				IsResource:   isResource,
				ResourceType: resourceType,
				ResourceName: resourceName,
				PodName:      "",
				PodPort:      podPort,
				NewPath:      "",
			}
		} else {
			// fallback: treat as pod name, and set podPort if present
			podName = name
			if verbose {
				println("[kubectl-curl] Treating as pod ", podName, " port=", podPort)
			}
			return ResourceTarget{
				IsResource:   false,
				ResourceType: "",
				ResourceName: "",
				PodName:      podName,
				PodPort:      podPort,
				NewPath:      "",
			}
		}
	}

	// 4. Otherwise, treat as podname[:port] or fallback
	parts := strings.SplitN(hostPort, ":", 2)
	name := parts[0]
	if len(parts) == 2 {
		podPort = parts[1]
	} else {
		podPort = ""
	}

	if isPodName != nil && isPodName(name) {
		podName = name
		// Ensure podPort is set correctly
		if len(parts) == 2 {
			podPort = parts[1]
		}
		if verbose {
			println("[kubectl-curl] Found pod " + podName)
		}
		if requestURL.Path != "" && requestURL.Host != "" {
			newPath = requestURL.Path
		} else {
			newPath = ""
		}
		return ResourceTarget{
			IsResource:   false,
			ResourceType: "",
			ResourceName: "",
			PodName:      podName,
			PodPort:      podPort,
			NewPath:      newPath,
		}
	} else if isDeploymentName != nil && isDeploymentName(name) {
		resourceType = "deployment"
		resourceName = name
		isResource = true
		if verbose {
			println("[kubectl-curl] Found deployment/" + resourceName)
		}
		newPath = ""
		return ResourceTarget{
			IsResource:   isResource,
			ResourceType: resourceType,
			ResourceName: resourceName,
			PodName:      "",
			PodPort:      podPort,
			NewPath:      newPath,
		}
	} else if isStatefulSetName != nil && isStatefulSetName(name) {
		resourceType = "statefulset"
		resourceName = name
		isResource = true
		if verbose {
			println("[kubectl-curl] Found statefulset/" + resourceName)
		}
		newPath = ""
		return ResourceTarget{
			IsResource:   isResource,
			ResourceType: resourceType,
			ResourceName: resourceName,
			PodName:      "",
			PodPort:      podPort,
			NewPath:      newPath,
		}
	} else if isDaemonSetName != nil && isDaemonSetName(name) {
		resourceType = "daemonset"
		resourceName = name
		isResource = true
		if verbose {
			println("[kubectl-curl] Found daemonset/" + resourceName)
		}
		newPath = ""
		return ResourceTarget{
			IsResource:   isResource,
			ResourceType: resourceType,
			ResourceName: resourceName,
			PodName:      "",
			PodPort:      podPort,
			NewPath:      newPath,
		}
	} else {
		// fallback: treat as pod name, and set podPort if present
		podName = name
		if len(parts) == 2 {
			podPort = parts[1]
		}
		if verbose {
			println("[kubectl-curl] Treating as pod ", podName, " port=", podPort)
		}
		// Always set newPath to "" for fallback
		newPath = ""
		return ResourceTarget{
			IsResource:   false,
			ResourceType: "",
			ResourceName: "",
			PodName:      podName,
			PodPort:      podPort,
			NewPath:      newPath,
		}
	}

	// Special case: if Opaque is set (e.g., mypod:8080 with no scheme), treat Opaque as the input
	if requestURL.Opaque != "" && hostPort == "" && newPath == "" {
		hostPort = requestURL.Opaque
		newPath = ""
		// Handle as hostless input (same as above)
		parts := strings.SplitN(hostPort, ":", 2)
		name := parts[0]
		if len(parts) == 2 {
			podPort = parts[1]
		} else {
			podPort = ""
		}
		if isPodName != nil && isPodName(name) {
			podName = name
			if verbose {
				println("[kubectl-curl] Found pod (opaque)", podName, " port=", podPort)
			}
			return ResourceTarget{
				IsResource:   false,
				ResourceType: "",
				ResourceName: "",
				PodName:      podName,
				PodPort:      podPort,
				NewPath:      "",
			}
		} else if isDeploymentName != nil && isDeploymentName(name) {
			resourceType = "deployment"
			resourceName = name
			isResource = true
			if verbose {
				println("[kubectl-curl] Found deployment/ (opaque)", resourceName)
			}
			return ResourceTarget{
				IsResource:   isResource,
				ResourceType: resourceType,
				ResourceName: resourceName,
				PodName:      "",
				PodPort:      podPort,
				NewPath:      "",
			}
		} else if isStatefulSetName != nil && isStatefulSetName(name) {
			resourceType = "statefulset"
			resourceName = name
			isResource = true
			if verbose {
				println("[kubectl-curl] Found statefulset/ (opaque)", resourceName)
			}
			return ResourceTarget{
				IsResource:   isResource,
				ResourceType: resourceType,
				ResourceName: resourceName,
				PodName:      "",
				PodPort:      podPort,
				NewPath:      "",
			}
		} else if isDaemonSetName != nil && isDaemonSetName(name) {
			resourceType = "daemonset"
			resourceName = name
			isResource = true
			if verbose {
				println("[kubectl-curl] Found daemonset/ (opaque)", resourceName)
			}
			return ResourceTarget{
				IsResource:   isResource,
				ResourceType: resourceType,
				ResourceName: resourceName,
				PodName:      "",
				PodPort:      podPort,
				NewPath:      "",
			}
		} else {
			// fallback: treat as pod name, and set podPort if present
			podName = name
			if verbose {
				println("[kubectl-curl] Treating as pod (opaque)", podName, " port=", podPort)
			}
			return ResourceTarget{
				IsResource:   false,
				ResourceType: "",
				ResourceName: "",
				PodName:      podName,
				PodPort:      podPort,
				NewPath:      "",
			}
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
