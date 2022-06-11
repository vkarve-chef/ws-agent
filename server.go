package main

import (
	"context"
	pb "ws-agent/io.chef"
)

type chefInfraServer struct {
	pb.UnimplementedChefInfraServer
}

func (s *chefInfraServer) GetCookbookVersion(ctx context.Context, cookbook *pb.Cookbook) (*pb.Version, error) {
	version := pb.Version{Major: 1, Minor: 2, Patch: 3}
	return &version, nil
}
