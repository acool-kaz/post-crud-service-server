package grpc

import (
	"context"
	"log"

	"github.com/acool-kaz/post-crud-service-server/internal/delivery/grpc/mapper"
	"github.com/acool-kaz/post-crud-service-server/internal/models"
	"github.com/acool-kaz/post-crud-service-server/internal/service"
	post_crud_pb "github.com/acool-kaz/post-crud-service-server/pkg/post_crud"
)

type PostCRUDHandler struct {
	post_crud_pb.UnimplementedPostCRUDServiceServer
	service *service.Service
}

func InitPostCRUDHandler(service *service.Service) *PostCRUDHandler {
	log.Println("init grpc post crud handler")
	return &PostCRUDHandler{
		service: service,
	}
}

func (p *PostCRUDHandler) Create(ctx context.Context, req *post_crud_pb.CreateRequest) (*post_crud_pb.CreateResponse, error) {
	post := mapper.ProtoToPost(req.Post)

	if err := p.service.Post.Create(ctx, post); err != nil {
		return &post_crud_pb.CreateResponse{Status: "failed"}, err
	}

	return &post_crud_pb.CreateResponse{Status: "post created"}, nil
}

func (p *PostCRUDHandler) Read(ctx context.Context, req *post_crud_pb.ReadRequest) (*post_crud_pb.ReadResponse, error) {
	if req.GetId() != 0 {
		ctx = context.WithValue(ctx, models.PostId, int(req.GetId()))
	}

	posts, err := p.service.Post.Read(ctx)
	if err != nil {
		return &post_crud_pb.ReadResponse{}, err
	}

	protoPosts := []*post_crud_pb.Post{}
	for _, p := range posts {
		protoPosts = append(protoPosts, mapper.PostToProto(p))
	}

	return &post_crud_pb.ReadResponse{Post: protoPosts}, nil
}

func (p *PostCRUDHandler) Update(ctx context.Context, req *post_crud_pb.UpdateRequest) (*post_crud_pb.UpdateResponse, error) {
	update := mapper.FromUpdatePostProtoToUpdatePost(req)
	id := int(req.GetId())

	if err := p.service.Post.Update(ctx, id, update); err != nil {
		return &post_crud_pb.UpdateResponse{Status: "failed"}, err
	}

	return &post_crud_pb.UpdateResponse{Status: "post updated"}, nil
}

func (p *PostCRUDHandler) Delete(ctx context.Context, req *post_crud_pb.DeleteRequest) (*post_crud_pb.DeleteResponse, error) {
	id := int(req.GetId())

	if err := p.service.Post.Delete(ctx, id); err != nil {
		return &post_crud_pb.DeleteResponse{Status: "failed"}, err
	}

	return &post_crud_pb.DeleteResponse{Status: "post deleted"}, nil
}
