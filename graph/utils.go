package graph

import "hash/fnv"

func getHash(val string) uint64 {
	h := fnv.New64()
	h.Write([]byte(val))
	return h.Sum64()
}