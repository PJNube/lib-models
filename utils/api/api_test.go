package api

import (
	"fmt"
	"testing"
)

func TestResolveRouteParams(t *testing.T) {
	pattern := "get.api.point.uuid.:uuid"
	route := "get.api.point.uuid.uuid_xxxxxx"

	fmt.Println(ResolveRouteParams(pattern, route))

	pattern = "get.api.point.name.:network_name.:device_name.:point_name"
	route = "get.api.point.name.net.dev.point"

	fmt.Println(ResolveRouteParams(pattern, route))

	pattern = "get.api.point.uuid.:uuid.tags"
	route = "get.api.point.uuid.uuid_xxxxxx.tags"

	fmt.Println(ResolveRouteParams(pattern, route))
}
