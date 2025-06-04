package helpers

import (
	"context"
	"time"
)

func GetAllocatedTime() (time.Duration, error) {
	db := GetDBQueries()
	user, err := db.GetFirstUser(context.Background())
	if err != nil {
		return 0, nil
	}
	return time.Duration(user.AllocatedTimeSeconds) * time.Second, nil
}
