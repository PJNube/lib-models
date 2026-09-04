package datatypes

import (
	"fmt"
	"strings"
)

// ConnectionScope says what a broker connection is for. The device has
// exactly one Local connection (its own bus, owned by the device) and, for
// now, at most one Cloud connection (the uplink an enrolled platform's
// requests arrive on).
type ConnectionScope string

const (
	ConnectionScopeLocal ConnectionScope = "local"
	ConnectionScopeCloud ConnectionScope = "cloud"
)

var ConnectionScopeMap = map[ConnectionScope]struct{}{
	ConnectionScopeLocal: {},
	ConnectionScopeCloud: {},
}

func (s *ConnectionScope) Validate() error {
	if s != nil {
		if _, ok := ConnectionScopeMap[*s]; ok {
			return nil
		}
	}
	return fmt.Errorf("please provide a valid connection scope, i.e.: %s", join(ConnectionScopeMap))
}

// ConnectionKind is the transport a connection row describes. The row's
// config and credentials blobs are interpreted according to it.
type ConnectionKind string

const (
	ConnectionKindNats ConnectionKind = "nats"
)

var ConnectionKindMap = map[ConnectionKind]struct{}{
	ConnectionKindNats: {},
}

func (k *ConnectionKind) Validate() error {
	if k != nil {
		if _, ok := ConnectionKindMap[*k]; ok {
			return nil
		}
	}
	return fmt.Errorf("please provide a valid connection kind, i.e.: %s", join(ConnectionKindMap))
}

func join[T ~string](m map[T]struct{}) string {
	v := make([]string, 0, len(m))
	for k := range m {
		v = append(v, string(k))
	}
	return strings.Join(v, " or ")
}
