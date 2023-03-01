import * as jspb from 'google-protobuf'

import * as common_character_pb from '../common/character_pb';


export class LoginRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): LoginRequest;

  getPassword(): string;
  setPassword(value: string): LoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class LoginResponse extends jspb.Message {
  getToken(): string;
  setToken(value: string): LoginResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginResponse): LoginResponse.AsObject;
  static serializeBinaryToWriter(message: LoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginResponse;
  static deserializeBinaryFromReader(message: LoginResponse, reader: jspb.BinaryReader): LoginResponse;
}

export namespace LoginResponse {
  export type AsObject = {
    token: string,
  }
}

export class GetCharactersRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCharactersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCharactersRequest): GetCharactersRequest.AsObject;
  static serializeBinaryToWriter(message: GetCharactersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCharactersRequest;
  static deserializeBinaryFromReader(message: GetCharactersRequest, reader: jspb.BinaryReader): GetCharactersRequest;
}

export namespace GetCharactersRequest {
  export type AsObject = {
  }
}

export class GetCharactersResponse extends jspb.Message {
  getCharsList(): Array<common_character_pb.Character>;
  setCharsList(value: Array<common_character_pb.Character>): GetCharactersResponse;
  clearCharsList(): GetCharactersResponse;
  addChars(value?: common_character_pb.Character, index?: number): common_character_pb.Character;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCharactersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCharactersResponse): GetCharactersResponse.AsObject;
  static serializeBinaryToWriter(message: GetCharactersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCharactersResponse;
  static deserializeBinaryFromReader(message: GetCharactersResponse, reader: jspb.BinaryReader): GetCharactersResponse;
}

export namespace GetCharactersResponse {
  export type AsObject = {
    charsList: Array<common_character_pb.Character.AsObject>,
  }
}

export class ChooseCharacterRequest extends jspb.Message {
  getIdentifier(): string;
  setIdentifier(value: string): ChooseCharacterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChooseCharacterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChooseCharacterRequest): ChooseCharacterRequest.AsObject;
  static serializeBinaryToWriter(message: ChooseCharacterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChooseCharacterRequest;
  static deserializeBinaryFromReader(message: ChooseCharacterRequest, reader: jspb.BinaryReader): ChooseCharacterRequest;
}

export namespace ChooseCharacterRequest {
  export type AsObject = {
    identifier: string,
  }
}

export class ChooseCharacterResponse extends jspb.Message {
  getToken(): string;
  setToken(value: string): ChooseCharacterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChooseCharacterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ChooseCharacterResponse): ChooseCharacterResponse.AsObject;
  static serializeBinaryToWriter(message: ChooseCharacterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChooseCharacterResponse;
  static deserializeBinaryFromReader(message: ChooseCharacterResponse, reader: jspb.BinaryReader): ChooseCharacterResponse;
}

export namespace ChooseCharacterResponse {
  export type AsObject = {
    token: string,
  }
}

export class LogoutRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LogoutRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LogoutRequest): LogoutRequest.AsObject;
  static serializeBinaryToWriter(message: LogoutRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LogoutRequest;
  static deserializeBinaryFromReader(message: LogoutRequest, reader: jspb.BinaryReader): LogoutRequest;
}

export namespace LogoutRequest {
  export type AsObject = {
  }
}

export class LogoutResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): LogoutResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LogoutResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LogoutResponse): LogoutResponse.AsObject;
  static serializeBinaryToWriter(message: LogoutResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LogoutResponse;
  static deserializeBinaryFromReader(message: LogoutResponse, reader: jspb.BinaryReader): LogoutResponse;
}

export namespace LogoutResponse {
  export type AsObject = {
    success: boolean,
  }
}

