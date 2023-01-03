package mapper

import (
	"github.com/acool-kaz/post-crud-service-server/internal/models"
	post_crud_pb "github.com/acool-kaz/post-crud-service-server/pkg/post_crud"
)

func PostToProto(p models.Post) *post_crud_pb.Post {
	return &post_crud_pb.Post{
		Id:     int32(p.Id),
		UserId: int32(p.UserId),
		Title:  p.Title,
		Body:   p.Body,
	}
}

func ProtoToPost(post *post_crud_pb.Post) models.Post {
	return models.Post{
		Id:     int(post.GetId()),
		UserId: int(post.GetUserId()),
		Title:  post.GetTitle(),
		Body:   post.GetBody(),
	}
}
