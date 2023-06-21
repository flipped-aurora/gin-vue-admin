package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/clothing"

type JobGroupByProcessResponse struct {
	Process clothing.Process `json:"process"`
	Jobs    []clothing.Job   `json:"jobs"`
}
