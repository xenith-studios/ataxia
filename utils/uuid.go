package utils

import (
	gouuid "github.com/nu7hatch/gouuid"
)

// UUID creates a new uuid and returns it as a string
func UUID() string {
	id, _ := gouuid.NewV4()
	return id.String()
}
