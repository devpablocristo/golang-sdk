package sdkmapdb

import (
	ports "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb/ports"
)

func Boostrap() ports.Repository {
	return newRepository()
}
