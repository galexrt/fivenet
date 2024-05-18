// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "services/messenger/messenger.proto" (package "services.messenger", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Message } from "../../resources/messenger/message";
import { Thread } from "../../resources/messenger/thread";
import { PaginationResponse } from "../../resources/common/database/database";
import { Timestamp } from "../../resources/timestamp/timestamp";
import { PaginationRequest } from "../../resources/common/database/database";
/**
 * @generated from protobuf message services.messenger.ListThreadsRequest
 */
export interface ListThreadsRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp after = 2;
     */
    after?: Timestamp;
}
/**
 * @generated from protobuf message services.messenger.ListThreadsResponse
 */
export interface ListThreadsResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.messenger.Thread threads = 2;
     */
    threads: Thread[];
}
/**
 * @generated from protobuf message services.messenger.CreateOrUpdateThreadRequest
 */
export interface CreateOrUpdateThreadRequest {
    /**
     * @generated from protobuf field: resources.messenger.Thread thread = 1;
     */
    thread?: Thread;
}
/**
 * @generated from protobuf message services.messenger.CreateOrUpdateThreadResponse
 */
export interface CreateOrUpdateThreadResponse {
    /**
     * @generated from protobuf field: resources.messenger.Thread thread = 1;
     */
    thread?: Thread;
}
/**
 * @generated from protobuf message services.messenger.DeleteThreadRequest
 */
export interface DeleteThreadRequest {
    /**
     * @generated from protobuf field: uint64 thread_id = 1 [jstype = JS_STRING];
     */
    threadId: string;
}
/**
 * @generated from protobuf message services.messenger.DeleteThreadResponse
 */
export interface DeleteThreadResponse {
}
// Messages

/**
 * @generated from protobuf message services.messenger.GetThreadMessagesRequest
 */
export interface GetThreadMessagesRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * @generated from protobuf field: uint64 thread_id = 2 [jstype = JS_STRING];
     */
    threadId: string;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp start = 3;
     */
    start?: Timestamp;
}
/**
 * @generated from protobuf message services.messenger.GetThreadMessagesResponse
 */
export interface GetThreadMessagesResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.messenger.Message messages = 2;
     */
    messages: Message[];
}
/**
 * @generated from protobuf message services.messenger.PostMessageRequest
 */
export interface PostMessageRequest {
    /**
     * @generated from protobuf field: resources.messenger.Message message = 1;
     */
    message?: Message;
}
/**
 * @generated from protobuf message services.messenger.PostMessageResponse
 */
export interface PostMessageResponse {
    /**
     * @generated from protobuf field: resources.messenger.Message message = 1;
     */
    message?: Message;
}
/**
 * @generated from protobuf message services.messenger.DeleteMessageRequest
 */
export interface DeleteMessageRequest {
    /**
     * @generated from protobuf field: uint64 message_id = 1 [jstype = JS_STRING];
     */
    messageId: string;
}
/**
 * @generated from protobuf message services.messenger.DeleteMessageResponse
 */
export interface DeleteMessageResponse {
}
// @generated message type with reflection information, may provide speed optimized methods
class ListThreadsRequest$Type extends MessageType<ListThreadsRequest> {
    constructor() {
        super("services.messenger.ListThreadsRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "after", kind: "message", T: () => Timestamp }
        ]);
    }
    create(value?: PartialMessage<ListThreadsRequest>): ListThreadsRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<ListThreadsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListThreadsRequest): ListThreadsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationRequest pagination */ 1:
                    message.pagination = PaginationRequest.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* optional resources.timestamp.Timestamp after */ 2:
                    message.after = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.after);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ListThreadsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationRequest pagination = 1; */
        if (message.pagination)
            PaginationRequest.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp after = 2; */
        if (message.after)
            Timestamp.internalBinaryWrite(message.after, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.ListThreadsRequest
 */
export const ListThreadsRequest = new ListThreadsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListThreadsResponse$Type extends MessageType<ListThreadsResponse> {
    constructor() {
        super("services.messenger.ListThreadsResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "threads", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Thread }
        ]);
    }
    create(value?: PartialMessage<ListThreadsResponse>): ListThreadsResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.threads = [];
        if (value !== undefined)
            reflectionMergePartial<ListThreadsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListThreadsResponse): ListThreadsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationResponse pagination */ 1:
                    message.pagination = PaginationResponse.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* repeated resources.messenger.Thread threads */ 2:
                    message.threads.push(Thread.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ListThreadsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationResponse pagination = 1; */
        if (message.pagination)
            PaginationResponse.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated resources.messenger.Thread threads = 2; */
        for (let i = 0; i < message.threads.length; i++)
            Thread.internalBinaryWrite(message.threads[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.ListThreadsResponse
 */
export const ListThreadsResponse = new ListThreadsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateOrUpdateThreadRequest$Type extends MessageType<CreateOrUpdateThreadRequest> {
    constructor() {
        super("services.messenger.CreateOrUpdateThreadRequest", [
            { no: 1, name: "thread", kind: "message", T: () => Thread, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
    create(value?: PartialMessage<CreateOrUpdateThreadRequest>): CreateOrUpdateThreadRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<CreateOrUpdateThreadRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateOrUpdateThreadRequest): CreateOrUpdateThreadRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.messenger.Thread thread */ 1:
                    message.thread = Thread.internalBinaryRead(reader, reader.uint32(), options, message.thread);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: CreateOrUpdateThreadRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.messenger.Thread thread = 1; */
        if (message.thread)
            Thread.internalBinaryWrite(message.thread, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.CreateOrUpdateThreadRequest
 */
export const CreateOrUpdateThreadRequest = new CreateOrUpdateThreadRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateOrUpdateThreadResponse$Type extends MessageType<CreateOrUpdateThreadResponse> {
    constructor() {
        super("services.messenger.CreateOrUpdateThreadResponse", [
            { no: 1, name: "thread", kind: "message", T: () => Thread }
        ]);
    }
    create(value?: PartialMessage<CreateOrUpdateThreadResponse>): CreateOrUpdateThreadResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<CreateOrUpdateThreadResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateOrUpdateThreadResponse): CreateOrUpdateThreadResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.messenger.Thread thread */ 1:
                    message.thread = Thread.internalBinaryRead(reader, reader.uint32(), options, message.thread);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: CreateOrUpdateThreadResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.messenger.Thread thread = 1; */
        if (message.thread)
            Thread.internalBinaryWrite(message.thread, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.CreateOrUpdateThreadResponse
 */
export const CreateOrUpdateThreadResponse = new CreateOrUpdateThreadResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteThreadRequest$Type extends MessageType<DeleteThreadRequest> {
    constructor() {
        super("services.messenger.DeleteThreadRequest", [
            { no: 1, name: "thread_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
    create(value?: PartialMessage<DeleteThreadRequest>): DeleteThreadRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.threadId = "0";
        if (value !== undefined)
            reflectionMergePartial<DeleteThreadRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteThreadRequest): DeleteThreadRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 thread_id = 1 [jstype = JS_STRING];*/ 1:
                    message.threadId = reader.uint64().toString();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: DeleteThreadRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 thread_id = 1 [jstype = JS_STRING]; */
        if (message.threadId !== "0")
            writer.tag(1, WireType.Varint).uint64(message.threadId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.DeleteThreadRequest
 */
export const DeleteThreadRequest = new DeleteThreadRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteThreadResponse$Type extends MessageType<DeleteThreadResponse> {
    constructor() {
        super("services.messenger.DeleteThreadResponse", []);
    }
    create(value?: PartialMessage<DeleteThreadResponse>): DeleteThreadResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<DeleteThreadResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteThreadResponse): DeleteThreadResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: DeleteThreadResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.DeleteThreadResponse
 */
export const DeleteThreadResponse = new DeleteThreadResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetThreadMessagesRequest$Type extends MessageType<GetThreadMessagesRequest> {
    constructor() {
        super("services.messenger.GetThreadMessagesRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "thread_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 3, name: "start", kind: "message", T: () => Timestamp }
        ]);
    }
    create(value?: PartialMessage<GetThreadMessagesRequest>): GetThreadMessagesRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.threadId = "0";
        if (value !== undefined)
            reflectionMergePartial<GetThreadMessagesRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetThreadMessagesRequest): GetThreadMessagesRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationRequest pagination */ 1:
                    message.pagination = PaginationRequest.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* uint64 thread_id = 2 [jstype = JS_STRING];*/ 2:
                    message.threadId = reader.uint64().toString();
                    break;
                case /* resources.timestamp.Timestamp start */ 3:
                    message.start = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.start);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetThreadMessagesRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationRequest pagination = 1; */
        if (message.pagination)
            PaginationRequest.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* uint64 thread_id = 2 [jstype = JS_STRING]; */
        if (message.threadId !== "0")
            writer.tag(2, WireType.Varint).uint64(message.threadId);
        /* resources.timestamp.Timestamp start = 3; */
        if (message.start)
            Timestamp.internalBinaryWrite(message.start, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.GetThreadMessagesRequest
 */
export const GetThreadMessagesRequest = new GetThreadMessagesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetThreadMessagesResponse$Type extends MessageType<GetThreadMessagesResponse> {
    constructor() {
        super("services.messenger.GetThreadMessagesResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "messages", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Message }
        ]);
    }
    create(value?: PartialMessage<GetThreadMessagesResponse>): GetThreadMessagesResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.messages = [];
        if (value !== undefined)
            reflectionMergePartial<GetThreadMessagesResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetThreadMessagesResponse): GetThreadMessagesResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationResponse pagination */ 1:
                    message.pagination = PaginationResponse.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* repeated resources.messenger.Message messages */ 2:
                    message.messages.push(Message.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetThreadMessagesResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationResponse pagination = 1; */
        if (message.pagination)
            PaginationResponse.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated resources.messenger.Message messages = 2; */
        for (let i = 0; i < message.messages.length; i++)
            Message.internalBinaryWrite(message.messages[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.GetThreadMessagesResponse
 */
export const GetThreadMessagesResponse = new GetThreadMessagesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PostMessageRequest$Type extends MessageType<PostMessageRequest> {
    constructor() {
        super("services.messenger.PostMessageRequest", [
            { no: 1, name: "message", kind: "message", T: () => Message, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
    create(value?: PartialMessage<PostMessageRequest>): PostMessageRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<PostMessageRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PostMessageRequest): PostMessageRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.messenger.Message message */ 1:
                    message.message = Message.internalBinaryRead(reader, reader.uint32(), options, message.message);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: PostMessageRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.messenger.Message message = 1; */
        if (message.message)
            Message.internalBinaryWrite(message.message, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.PostMessageRequest
 */
export const PostMessageRequest = new PostMessageRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PostMessageResponse$Type extends MessageType<PostMessageResponse> {
    constructor() {
        super("services.messenger.PostMessageResponse", [
            { no: 1, name: "message", kind: "message", T: () => Message, options: { "validate.rules": { message: { required: true } } } }
        ]);
    }
    create(value?: PartialMessage<PostMessageResponse>): PostMessageResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<PostMessageResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PostMessageResponse): PostMessageResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.messenger.Message message */ 1:
                    message.message = Message.internalBinaryRead(reader, reader.uint32(), options, message.message);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: PostMessageResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.messenger.Message message = 1; */
        if (message.message)
            Message.internalBinaryWrite(message.message, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.PostMessageResponse
 */
export const PostMessageResponse = new PostMessageResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteMessageRequest$Type extends MessageType<DeleteMessageRequest> {
    constructor() {
        super("services.messenger.DeleteMessageRequest", [
            { no: 1, name: "message_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
    create(value?: PartialMessage<DeleteMessageRequest>): DeleteMessageRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.messageId = "0";
        if (value !== undefined)
            reflectionMergePartial<DeleteMessageRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteMessageRequest): DeleteMessageRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 message_id = 1 [jstype = JS_STRING];*/ 1:
                    message.messageId = reader.uint64().toString();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: DeleteMessageRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 message_id = 1 [jstype = JS_STRING]; */
        if (message.messageId !== "0")
            writer.tag(1, WireType.Varint).uint64(message.messageId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.DeleteMessageRequest
 */
export const DeleteMessageRequest = new DeleteMessageRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteMessageResponse$Type extends MessageType<DeleteMessageResponse> {
    constructor() {
        super("services.messenger.DeleteMessageResponse", []);
    }
    create(value?: PartialMessage<DeleteMessageResponse>): DeleteMessageResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<DeleteMessageResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteMessageResponse): DeleteMessageResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: DeleteMessageResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.messenger.DeleteMessageResponse
 */
export const DeleteMessageResponse = new DeleteMessageResponse$Type();
/**
 * @generated ServiceType for protobuf service services.messenger.MessengerService
 */
export const MessengerService = new ServiceType("services.messenger.MessengerService", [
    { name: "ListThreads", options: {}, I: ListThreadsRequest, O: ListThreadsResponse },
    { name: "CreateOrUpdateThread", options: {}, I: CreateOrUpdateThreadRequest, O: CreateOrUpdateThreadResponse },
    { name: "DeleteThread", options: {}, I: DeleteThreadRequest, O: DeleteThreadResponse },
    { name: "GetThreadMessages", options: {}, I: GetThreadMessagesRequest, O: GetThreadMessagesResponse },
    { name: "PostMessage", options: {}, I: PostMessageRequest, O: PostMessageResponse },
    { name: "DeleteMessage", options: {}, I: DeleteMessageRequest, O: DeleteMessageResponse }
]);