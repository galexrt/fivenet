// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/auth/auth.proto

package auth

import (
	grpc "google.golang.org/grpc"
)

func (s *Server) RegisterServer(srv *grpc.Server) {
	RegisterAuthServiceServer(srv, s)
}
