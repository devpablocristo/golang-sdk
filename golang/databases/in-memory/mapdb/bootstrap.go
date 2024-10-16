package sdkmapdb

import (
	ports "github.com/devpablocristo/sdk/golang/sdk/databases/in-memory/mapdb/ports"
)

func Boostrap() ports.Repository {
	return newRepository()
}
