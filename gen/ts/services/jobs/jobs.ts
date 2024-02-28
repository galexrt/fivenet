// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "services/jobs/jobs.proto" (package "services.jobs", syntax proto3)
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
import { JobsUserProps } from "../../resources/jobs/colleagues";
import { JobsUserActivity } from "../../resources/jobs/colleagues";
import { Colleague } from "../../resources/jobs/colleagues";
import { PaginationResponse } from "../../resources/common/database/database";
import { PaginationRequest } from "../../resources/common/database/database";
// Colleagues

/**
 * @generated from protobuf message services.jobs.ListColleaguesRequest
 */
export interface ListColleaguesRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * Search params
     *
     * @generated from protobuf field: string search_name = 2;
     */
    searchName: string;
    /**
     * @generated from protobuf field: optional int32 user_id = 3;
     */
    userId?: number;
}
/**
 * @generated from protobuf message services.jobs.ListColleaguesResponse
 */
export interface ListColleaguesResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.jobs.Colleague colleagues = 2;
     */
    colleagues: Colleague[];
}
/**
 * @generated from protobuf message services.jobs.GetSelfRequest
 */
export interface GetSelfRequest {
}
/**
 * @generated from protobuf message services.jobs.GetSelfResponse
 */
export interface GetSelfResponse {
    /**
     * @generated from protobuf field: resources.jobs.Colleague colleague = 1;
     */
    colleague?: Colleague;
}
/**
 * @generated from protobuf message services.jobs.GetColleagueRequest
 */
export interface GetColleagueRequest {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number;
}
/**
 * @generated from protobuf message services.jobs.GetColleagueResponse
 */
export interface GetColleagueResponse {
    /**
     * @generated from protobuf field: resources.jobs.Colleague colleague = 1;
     */
    colleague?: Colleague;
}
/**
 * @generated from protobuf message services.jobs.ListColleagueActivityRequest
 */
export interface ListColleagueActivityRequest {
    /**
     * @generated from protobuf field: resources.common.database.PaginationRequest pagination = 1;
     */
    pagination?: PaginationRequest;
    /**
     * @generated from protobuf field: int32 user_id = 2;
     */
    userId: number;
}
/**
 * @generated from protobuf message services.jobs.ListColleagueActivityResponse
 */
export interface ListColleagueActivityResponse {
    /**
     * @generated from protobuf field: resources.common.database.PaginationResponse pagination = 1;
     */
    pagination?: PaginationResponse;
    /**
     * @generated from protobuf field: repeated resources.jobs.JobsUserActivity activity = 2;
     */
    activity: JobsUserActivity[];
}
/**
 * @generated from protobuf message services.jobs.SetJobsUserPropsRequest
 */
export interface SetJobsUserPropsRequest {
    /**
     * @generated from protobuf field: resources.jobs.JobsUserProps props = 1;
     */
    props?: JobsUserProps;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string reason = 2;
     */
    reason: string;
}
/**
 * @generated from protobuf message services.jobs.SetJobsUserPropsResponse
 */
export interface SetJobsUserPropsResponse {
    /**
     * @generated from protobuf field: resources.jobs.JobsUserProps props = 1;
     */
    props?: JobsUserProps;
}
// MOTD

/**
 * @generated from protobuf message services.jobs.GetMOTDRequest
 */
export interface GetMOTDRequest {
}
/**
 * @generated from protobuf message services.jobs.GetMOTDResponse
 */
export interface GetMOTDResponse {
    /**
     * @generated from protobuf field: string motd = 1;
     */
    motd: string;
}
/**
 * @generated from protobuf message services.jobs.SetMOTDRequest
 */
export interface SetMOTDRequest {
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: string motd = 1;
     */
    motd: string;
}
/**
 * @generated from protobuf message services.jobs.SetMOTDResponse
 */
export interface SetMOTDResponse {
    /**
     * @generated from protobuf field: string motd = 1;
     */
    motd: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class ListColleaguesRequest$Type extends MessageType<ListColleaguesRequest> {
    constructor() {
        super("services.jobs.ListColleaguesRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "search_name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 3, name: "user_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<ListColleaguesRequest>): ListColleaguesRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.searchName = "";
        if (value !== undefined)
            reflectionMergePartial<ListColleaguesRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListColleaguesRequest): ListColleaguesRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationRequest pagination */ 1:
                    message.pagination = PaginationRequest.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* string search_name */ 2:
                    message.searchName = reader.string();
                    break;
                case /* optional int32 user_id */ 3:
                    message.userId = reader.int32();
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
    internalBinaryWrite(message: ListColleaguesRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationRequest pagination = 1; */
        if (message.pagination)
            PaginationRequest.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* string search_name = 2; */
        if (message.searchName !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.searchName);
        /* optional int32 user_id = 3; */
        if (message.userId !== undefined)
            writer.tag(3, WireType.Varint).int32(message.userId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ListColleaguesRequest
 */
export const ListColleaguesRequest = new ListColleaguesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListColleaguesResponse$Type extends MessageType<ListColleaguesResponse> {
    constructor() {
        super("services.jobs.ListColleaguesResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "colleagues", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Colleague }
        ]);
    }
    create(value?: PartialMessage<ListColleaguesResponse>): ListColleaguesResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.colleagues = [];
        if (value !== undefined)
            reflectionMergePartial<ListColleaguesResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListColleaguesResponse): ListColleaguesResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationResponse pagination */ 1:
                    message.pagination = PaginationResponse.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* repeated resources.jobs.Colleague colleagues */ 2:
                    message.colleagues.push(Colleague.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: ListColleaguesResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationResponse pagination = 1; */
        if (message.pagination)
            PaginationResponse.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated resources.jobs.Colleague colleagues = 2; */
        for (let i = 0; i < message.colleagues.length; i++)
            Colleague.internalBinaryWrite(message.colleagues[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ListColleaguesResponse
 */
export const ListColleaguesResponse = new ListColleaguesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetSelfRequest$Type extends MessageType<GetSelfRequest> {
    constructor() {
        super("services.jobs.GetSelfRequest", []);
    }
    create(value?: PartialMessage<GetSelfRequest>): GetSelfRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetSelfRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetSelfRequest): GetSelfRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: GetSelfRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.GetSelfRequest
 */
export const GetSelfRequest = new GetSelfRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetSelfResponse$Type extends MessageType<GetSelfResponse> {
    constructor() {
        super("services.jobs.GetSelfResponse", [
            { no: 1, name: "colleague", kind: "message", T: () => Colleague }
        ]);
    }
    create(value?: PartialMessage<GetSelfResponse>): GetSelfResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetSelfResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetSelfResponse): GetSelfResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.jobs.Colleague colleague */ 1:
                    message.colleague = Colleague.internalBinaryRead(reader, reader.uint32(), options, message.colleague);
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
    internalBinaryWrite(message: GetSelfResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.jobs.Colleague colleague = 1; */
        if (message.colleague)
            Colleague.internalBinaryWrite(message.colleague, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.GetSelfResponse
 */
export const GetSelfResponse = new GetSelfResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetColleagueRequest$Type extends MessageType<GetColleagueRequest> {
    constructor() {
        super("services.jobs.GetColleagueRequest", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<GetColleagueRequest>): GetColleagueRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        if (value !== undefined)
            reflectionMergePartial<GetColleagueRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetColleagueRequest): GetColleagueRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 user_id */ 1:
                    message.userId = reader.int32();
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
    internalBinaryWrite(message: GetColleagueRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 user_id = 1; */
        if (message.userId !== 0)
            writer.tag(1, WireType.Varint).int32(message.userId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.GetColleagueRequest
 */
export const GetColleagueRequest = new GetColleagueRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetColleagueResponse$Type extends MessageType<GetColleagueResponse> {
    constructor() {
        super("services.jobs.GetColleagueResponse", [
            { no: 1, name: "colleague", kind: "message", T: () => Colleague }
        ]);
    }
    create(value?: PartialMessage<GetColleagueResponse>): GetColleagueResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetColleagueResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetColleagueResponse): GetColleagueResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.jobs.Colleague colleague */ 1:
                    message.colleague = Colleague.internalBinaryRead(reader, reader.uint32(), options, message.colleague);
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
    internalBinaryWrite(message: GetColleagueResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.jobs.Colleague colleague = 1; */
        if (message.colleague)
            Colleague.internalBinaryWrite(message.colleague, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.GetColleagueResponse
 */
export const GetColleagueResponse = new GetColleagueResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListColleagueActivityRequest$Type extends MessageType<ListColleagueActivityRequest> {
    constructor() {
        super("services.jobs.ListColleagueActivityRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationRequest, options: { "validate.rules": { message: { required: true } } } },
            { no: 2, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<ListColleagueActivityRequest>): ListColleagueActivityRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        if (value !== undefined)
            reflectionMergePartial<ListColleagueActivityRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListColleagueActivityRequest): ListColleagueActivityRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationRequest pagination */ 1:
                    message.pagination = PaginationRequest.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* int32 user_id */ 2:
                    message.userId = reader.int32();
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
    internalBinaryWrite(message: ListColleagueActivityRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationRequest pagination = 1; */
        if (message.pagination)
            PaginationRequest.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* int32 user_id = 2; */
        if (message.userId !== 0)
            writer.tag(2, WireType.Varint).int32(message.userId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ListColleagueActivityRequest
 */
export const ListColleagueActivityRequest = new ListColleagueActivityRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListColleagueActivityResponse$Type extends MessageType<ListColleagueActivityResponse> {
    constructor() {
        super("services.jobs.ListColleagueActivityResponse", [
            { no: 1, name: "pagination", kind: "message", T: () => PaginationResponse },
            { no: 2, name: "activity", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => JobsUserActivity }
        ]);
    }
    create(value?: PartialMessage<ListColleagueActivityResponse>): ListColleagueActivityResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.activity = [];
        if (value !== undefined)
            reflectionMergePartial<ListColleagueActivityResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListColleagueActivityResponse): ListColleagueActivityResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.common.database.PaginationResponse pagination */ 1:
                    message.pagination = PaginationResponse.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* repeated resources.jobs.JobsUserActivity activity */ 2:
                    message.activity.push(JobsUserActivity.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: ListColleagueActivityResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.common.database.PaginationResponse pagination = 1; */
        if (message.pagination)
            PaginationResponse.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated resources.jobs.JobsUserActivity activity = 2; */
        for (let i = 0; i < message.activity.length; i++)
            JobsUserActivity.internalBinaryWrite(message.activity[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.ListColleagueActivityResponse
 */
export const ListColleagueActivityResponse = new ListColleagueActivityResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SetJobsUserPropsRequest$Type extends MessageType<SetJobsUserPropsRequest> {
    constructor() {
        super("services.jobs.SetJobsUserPropsRequest", [
            { no: 1, name: "props", kind: "message", T: () => JobsUserProps },
            { no: 2, name: "reason", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "255", ignoreEmpty: true } } } }
        ]);
    }
    create(value?: PartialMessage<SetJobsUserPropsRequest>): SetJobsUserPropsRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.reason = "";
        if (value !== undefined)
            reflectionMergePartial<SetJobsUserPropsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SetJobsUserPropsRequest): SetJobsUserPropsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.jobs.JobsUserProps props */ 1:
                    message.props = JobsUserProps.internalBinaryRead(reader, reader.uint32(), options, message.props);
                    break;
                case /* string reason */ 2:
                    message.reason = reader.string();
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
    internalBinaryWrite(message: SetJobsUserPropsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.jobs.JobsUserProps props = 1; */
        if (message.props)
            JobsUserProps.internalBinaryWrite(message.props, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* string reason = 2; */
        if (message.reason !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.reason);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.SetJobsUserPropsRequest
 */
export const SetJobsUserPropsRequest = new SetJobsUserPropsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SetJobsUserPropsResponse$Type extends MessageType<SetJobsUserPropsResponse> {
    constructor() {
        super("services.jobs.SetJobsUserPropsResponse", [
            { no: 1, name: "props", kind: "message", T: () => JobsUserProps }
        ]);
    }
    create(value?: PartialMessage<SetJobsUserPropsResponse>): SetJobsUserPropsResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<SetJobsUserPropsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SetJobsUserPropsResponse): SetJobsUserPropsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.jobs.JobsUserProps props */ 1:
                    message.props = JobsUserProps.internalBinaryRead(reader, reader.uint32(), options, message.props);
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
    internalBinaryWrite(message: SetJobsUserPropsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.jobs.JobsUserProps props = 1; */
        if (message.props)
            JobsUserProps.internalBinaryWrite(message.props, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.SetJobsUserPropsResponse
 */
export const SetJobsUserPropsResponse = new SetJobsUserPropsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetMOTDRequest$Type extends MessageType<GetMOTDRequest> {
    constructor() {
        super("services.jobs.GetMOTDRequest", []);
    }
    create(value?: PartialMessage<GetMOTDRequest>): GetMOTDRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetMOTDRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetMOTDRequest): GetMOTDRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: GetMOTDRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.GetMOTDRequest
 */
export const GetMOTDRequest = new GetMOTDRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetMOTDResponse$Type extends MessageType<GetMOTDResponse> {
    constructor() {
        super("services.jobs.GetMOTDResponse", [
            { no: 1, name: "motd", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<GetMOTDResponse>): GetMOTDResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.motd = "";
        if (value !== undefined)
            reflectionMergePartial<GetMOTDResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetMOTDResponse): GetMOTDResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string motd */ 1:
                    message.motd = reader.string();
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
    internalBinaryWrite(message: GetMOTDResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string motd = 1; */
        if (message.motd !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.motd);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.GetMOTDResponse
 */
export const GetMOTDResponse = new GetMOTDResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SetMOTDRequest$Type extends MessageType<SetMOTDRequest> {
    constructor() {
        super("services.jobs.SetMOTDRequest", [
            { no: 1, name: "motd", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "1024" } } } }
        ]);
    }
    create(value?: PartialMessage<SetMOTDRequest>): SetMOTDRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.motd = "";
        if (value !== undefined)
            reflectionMergePartial<SetMOTDRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SetMOTDRequest): SetMOTDRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string motd */ 1:
                    message.motd = reader.string();
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
    internalBinaryWrite(message: SetMOTDRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string motd = 1; */
        if (message.motd !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.motd);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.SetMOTDRequest
 */
export const SetMOTDRequest = new SetMOTDRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SetMOTDResponse$Type extends MessageType<SetMOTDResponse> {
    constructor() {
        super("services.jobs.SetMOTDResponse", [
            { no: 1, name: "motd", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SetMOTDResponse>): SetMOTDResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.motd = "";
        if (value !== undefined)
            reflectionMergePartial<SetMOTDResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SetMOTDResponse): SetMOTDResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string motd */ 1:
                    message.motd = reader.string();
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
    internalBinaryWrite(message: SetMOTDResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string motd = 1; */
        if (message.motd !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.motd);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message services.jobs.SetMOTDResponse
 */
export const SetMOTDResponse = new SetMOTDResponse$Type();
/**
 * @generated ServiceType for protobuf service services.jobs.JobsService
 */
export const JobsService = new ServiceType("services.jobs.JobsService", [
    { name: "ListColleagues", options: {}, I: ListColleaguesRequest, O: ListColleaguesResponse },
    { name: "GetSelf", options: {}, I: GetSelfRequest, O: GetSelfResponse },
    { name: "GetColleague", options: {}, I: GetColleagueRequest, O: GetColleagueResponse },
    { name: "ListColleagueActivity", options: {}, I: ListColleagueActivityRequest, O: ListColleagueActivityResponse },
    { name: "SetJobsUserProps", options: {}, I: SetJobsUserPropsRequest, O: SetJobsUserPropsResponse },
    { name: "GetMOTD", options: {}, I: GetMOTDRequest, O: GetMOTDResponse },
    { name: "SetMOTD", options: {}, I: SetMOTDRequest, O: SetMOTDResponse }
]);
