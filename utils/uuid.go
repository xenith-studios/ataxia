package utils

import (
	gouuid "github.com/nu7hatch/gouuid"
)

func UUID() string {
	id, _ := gouuid.NewV4()
	return id.String()
}
