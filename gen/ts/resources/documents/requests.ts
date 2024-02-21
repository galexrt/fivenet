// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "resources/documents/requests.proto" (package "resources.documents", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { DocActivityData } from "./activity";
import { UserShort } from "../users/users";
import { DocActivityType } from "./activity";
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.documents.DocRequest
 */
export interface DocRequest {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 document_id = 4 [jstype = JS_STRING];
     */
    documentId: string;
    /**
     * @generated from protobuf field: resources.documents.DocActivityType request_type = 5;
     */
    requestType: DocActivityType;
    /**
     * @generated from protobuf field: optional int32 creator_id = 6;
     */
    creatorId?: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 7;
     */
    creator?: UserShort; // @gotags: alias:"creator"
    /**
     * @generated from protobuf field: string creator_job = 8;
     */
    creatorJob: string;
    /**
     * @generated from protobuf field: optional string creator_job_label = 9;
     */
    creatorJobLabel?: string;
    /**
     * @generated from protobuf field: optional string reason = 10;
     */
    reason?: string;
    /**
     * @generated from protobuf field: resources.documents.DocActivityData data = 11;
     */
    data?: DocActivityData;
    /**
     * @generated from protobuf field: optional bool accepted = 12;
     */
    accepted?: boolean;
}
// @generated message type with reflection information, may provide speed optimized methods
class DocRequest$Type extends MessageType<DocRequest> {
    constructor() {
        super("resources.documents.DocRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "document_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 5, name: "request_type", kind: "enum", T: () => ["resources.documents.DocActivityType", DocActivityType, "DOC_ACTIVITY_TYPE_"], options: { "validate.rules": { enum: { in: [13, 14, 15, 16, 17, 18] } } } },
            { no: 6, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "creator", kind: "message", T: () => UserShort },
            { no: 8, name: "creator_job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 9, name: "creator_job_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 10, name: "reason", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 11, name: "data", kind: "message", T: () => DocActivityData },
            { no: 12, name: "accepted", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<DocRequest>): DocRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.documentId = "0";
        message.requestType = 0;
        message.creatorJob = "";
        if (value !== undefined)
            reflectionMergePartial<DocRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DocRequest): DocRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* resources.timestamp.Timestamp created_at */ 2:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* resources.timestamp.Timestamp updated_at */ 3:
                    message.updatedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.updatedAt);
                    break;
                case /* uint64 document_id = 4 [jstype = JS_STRING];*/ 4:
                    message.documentId = reader.uint64().toString();
                    break;
                case /* resources.documents.DocActivityType request_type */ 5:
                    message.requestType = reader.int32();
                    break;
                case /* optional int32 creator_id */ 6:
                    message.creatorId = reader.int32();
                    break;
                case /* optional resources.users.UserShort creator */ 7:
                    message.creator = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.creator);
                    break;
                case /* string creator_job */ 8:
                    message.creatorJob = reader.string();
                    break;
                case /* optional string creator_job_label */ 9:
                    message.creatorJobLabel = reader.string();
                    break;
                case /* optional string reason */ 10:
                    message.reason = reader.string();
                    break;
                case /* resources.documents.DocActivityData data */ 11:
                    message.data = DocActivityData.internalBinaryRead(reader, reader.uint32(), options, message.data);
                    break;
                case /* optional bool accepted */ 12:
                    message.accepted = reader.bool();
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
    internalBinaryWrite(message: DocRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* resources.timestamp.Timestamp created_at = 2; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* resources.timestamp.Timestamp updated_at = 3; */
        if (message.updatedAt)
            Timestamp.internalBinaryWrite(message.updatedAt, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* uint64 document_id = 4 [jstype = JS_STRING]; */
        if (message.documentId !== "0")
            writer.tag(4, WireType.Varint).uint64(message.documentId);
        /* resources.documents.DocActivityType request_type = 5; */
        if (message.requestType !== 0)
            writer.tag(5, WireType.Varint).int32(message.requestType);
        /* optional int32 creator_id = 6; */
        if (message.creatorId !== undefined)
            writer.tag(6, WireType.Varint).int32(message.creatorId);
        /* optional resources.users.UserShort creator = 7; */
        if (message.creator)
            UserShort.internalBinaryWrite(message.creator, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* string creator_job = 8; */
        if (message.creatorJob !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.creatorJob);
        /* optional string creator_job_label = 9; */
        if (message.creatorJobLabel !== undefined)
            writer.tag(9, WireType.LengthDelimited).string(message.creatorJobLabel);
        /* optional string reason = 10; */
        if (message.reason !== undefined)
            writer.tag(10, WireType.LengthDelimited).string(message.reason);
        /* resources.documents.DocActivityData data = 11; */
        if (message.data)
            DocActivityData.internalBinaryWrite(message.data, writer.tag(11, WireType.LengthDelimited).fork(), options).join();
        /* optional bool accepted = 12; */
        if (message.accepted !== undefined)
            writer.tag(12, WireType.Varint).bool(message.accepted);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.DocRequest
 */
export const DocRequest = new DocRequest$Type();
