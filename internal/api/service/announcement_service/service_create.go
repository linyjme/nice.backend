package announcement_service

import (
	"fmt"
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/repository/db_repo/announcement_repo"
)

func (s *service) Create(ctx core.Context) (err error) {
	var annList []announcement_repo.Announcement
	fmt.Println(annList)
	return nil
}
