//go:build (pg && mysql) || (pg && sqlite) || (mysql && sqlite)

package storage

// This file triggers a compilation error if multiple database tags are provided.
func init() {
	var _ = invalid_build_multiple_db_tags_specified
}
