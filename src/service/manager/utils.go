package manager

import (
	"fmt"
	"hash/crc64"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	hashTable = crc64.MakeTable(crc64.ISO)
)

// TODO
func hash(meta v1.ObjectMeta) string {
	return fmt.Sprintf("%d00000", crc64.Checksum([]byte(meta.Name), hashTable))[:5]
}