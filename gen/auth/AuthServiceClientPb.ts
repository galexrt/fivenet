/**
 * @fileoverview gRPC-Web generated client stub for gen.auth
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: auth/auth.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as auth_auth_pb from '../auth/auth_pb';


export class AccountServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorLogin = new grpcWeb.MethodDescriptor(
    '/gen.auth.AccountService/Login',
    grpcWeb.MethodType.UNARY,
    auth_auth_pb.LoginRequest,
    auth_auth_pb.LoginResponse,
    (request: auth_auth_pb.LoginRequest) => {
      return request.serializeBinary();
    },
    auth_auth_pb.LoginResponse.deserializeBinary
  );

  login(
    request: auth_auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null): Promise<auth_auth_pb.LoginResponse>;

  login(
    request: auth_auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_auth_pb.LoginResponse) => void): grpcWeb.ClientReadableStream<auth_auth_pb.LoginResponse>;

  login(
    request: auth_auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_auth_pb.LoginResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/gen.auth.AccountService/Login',
        request,
        metadata || {},
        this.methodDescriptorLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/gen.auth.AccountService/Login',
    request,
    metadata || {},
    this.methodDescriptorLogin);
  }

  methodDescriptorGetCharacters = new grpcWeb.MethodDescriptor(
    '/gen.auth.AccountService/GetCharacters',
    grpcWeb.MethodType.UNARY,
    auth_auth_pb.GetCharactersRequest,
    auth_auth_pb.GetCharactersResponse,
    (request: auth_auth_pb.GetCharactersRequest) => {
      return request.serializeBinary();
    },
    auth_auth_pb.GetCharactersResponse.deserializeBinary
  );

  getCharacters(
    request: auth_auth_pb.GetCharactersRequest,
    metadata: grpcWeb.Metadata | null): Promise<auth_auth_pb.GetCharactersResponse>;

  getCharacters(
    request: auth_auth_pb.GetCharactersRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_auth_pb.GetCharactersResponse) => void): grpcWeb.ClientReadableStream<auth_auth_pb.GetCharactersResponse>;

  getCharacters(
    request: auth_auth_pb.GetCharactersRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_auth_pb.GetCharactersResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/gen.auth.AccountService/GetCharacters',
        request,
        metadata || {},
        this.methodDescriptorGetCharacters,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/gen.auth.AccountService/GetCharacters',
    request,
    metadata || {},
    this.methodDescriptorGetCharacters);
  }

  methodDescriptorChooseCharacter = new grpcWeb.MethodDescriptor(
    '/gen.auth.AccountService/ChooseCharacter',
    grpcWeb.MethodType.UNARY,
    auth_auth_pb.ChooseCharacterRequest,
    auth_auth_pb.ChooseCharacterResponse,
    (request: auth_auth_pb.ChooseCharacterRequest) => {
      return request.serializeBinary();
    },
    auth_auth_pb.ChooseCharacterResponse.deserializeBinary
  );

  chooseCharacter(
    request: auth_auth_pb.ChooseCharacterRequest,
    metadata: grpcWeb.Metadata | null): Promise<auth_auth_pb.ChooseCharacterResponse>;

  chooseCharacter(
    request: auth_auth_pb.ChooseCharacterRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_auth_pb.ChooseCharacterResponse) => void): grpcWeb.ClientReadableStream<auth_auth_pb.ChooseCharacterResponse>;

  chooseCharacter(
    request: auth_auth_pb.ChooseCharacterRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_auth_pb.ChooseCharacterResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/gen.auth.AccountService/ChooseCharacter',
        request,
        metadata || {},
        this.methodDescriptorChooseCharacter,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/gen.auth.AccountService/ChooseCharacter',
    request,
    metadata || {},
    this.methodDescriptorChooseCharacter);
  }

  methodDescriptorLogout = new grpcWeb.MethodDescriptor(
    '/gen.auth.AccountService/Logout',
    grpcWeb.MethodType.UNARY,
    auth_auth_pb.LogoutRequest,
    auth_auth_pb.LogoutResponse,
    (request: auth_auth_pb.LogoutRequest) => {
      return request.serializeBinary();
    },
    auth_auth_pb.LogoutResponse.deserializeBinary
  );

  logout(
    request: auth_auth_pb.LogoutRequest,
    metadata: grpcWeb.Metadata | null): Promise<auth_auth_pb.LogoutResponse>;

  logout(
    request: auth_auth_pb.LogoutRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_auth_pb.LogoutResponse) => void): grpcWeb.ClientReadableStream<auth_auth_pb.LogoutResponse>;

  logout(
    request: auth_auth_pb.LogoutRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_auth_pb.LogoutResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/gen.auth.AccountService/Logout',
        request,
        metadata || {},
        this.methodDescriptorLogout,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/gen.auth.AccountService/Logout',
    request,
    metadata || {},
    this.methodDescriptorLogout);
  }

}

