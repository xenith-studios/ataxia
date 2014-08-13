package main

import (
	gouuid "github.com/nu7hatch/gouuid"
)

func uuid() string {
	id, _ := gouuid.NewV4()
	return id.String()
}
