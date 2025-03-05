package nuuid

import (
	"crypto/rand"
	"fmt"
	"io"
)

func MakeUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8]), nil
}

func SmallUUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return
	}
	uuid = fmt.Sprintf("id_%X%X", b[0:4], b[4:6])
	return
}

func ShortUUID(prefix ...string) string {
	u := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, u)
	if n != len(u) || err != nil {
		return "-error-uuid-"
	}
	uuid_ := fmt.Sprintf("%x%x", u[0:4], u[4:6])
	if len(prefix) > 0 {
		uuid_ = fmt.Sprintf("%s_%s", prefix[0], uuid_)
	}
	return uuid_
}
