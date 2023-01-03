package mapper

import (
	"github.com/acool-kaz/post-crud-service-server/internal/models"
	post_crud_pb "github.com/acool-kaz/post-crud-service-server/pkg/post_crud"
)

func FromUpdatePostProtoToUpdatePost(req *post_crud_pb.UpdateRequest) models.UpdatePost {
	return models.UpdatePost{
		Id:     int(req.GetNewId()),
		UserId: int(req.GetNewUserId()),
		Title:  req.GetNewTitle(),
		Body:   req.GetNewBody(),
	}
}
