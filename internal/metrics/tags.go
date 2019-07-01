package metrics

import (
	"go.opencensus.io/tag"
)

var (
	keyMethod      tag.Key = tag.MustNewKey("method")
	keyStatus      tag.Key = tag.MustNewKey("status")
	keyService     tag.Key = tag.MustNewKey("service")
	keyGRPCService tag.Key = tag.MustNewKey("grpc_service")
	keyHost        tag.Key = tag.MustNewKey("host")
)