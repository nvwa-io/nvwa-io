package vo

import "github.com/nvwa-io/nvwa-io/nvwa-server/entities"

type ReqUser struct {
	User entities.UserEntity `json:"user"`
}
