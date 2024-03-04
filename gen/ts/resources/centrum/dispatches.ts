// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "resources/centrum/dispatches.proto" (package "resources.centrum", syntax proto3)
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
import { UserShort } from "../users/users";
import { Unit } from "./units";
import { User } from "../users/users";
import { Attributes } from "./general";
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.centrum.Dispatch
 */
export interface Dispatch {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string; // @gotags: sql:"primary_key" alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: string job = 4;
     */
    job: string;
    /**
     * @generated from protobuf field: optional resources.centrum.DispatchStatus status = 5;
     */
    status?: DispatchStatus;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string message = 7;
     */
    message: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string description = 8;
     */
    description?: string;
    /**
     * @generated from protobuf field: optional resources.centrum.Attributes attributes = 9;
     */
    attributes?: Attributes;
    /**
     * @generated from protobuf field: double x = 10;
     */
    x: number;
    /**
     * @generated from protobuf field: double y = 11;
     */
    y: number;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string postal = 12;
     */
    postal?: string;
    /**
     * @generated from protobuf field: bool anon = 13;
     */
    anon: boolean;
    /**
     * @generated from protobuf field: optional int32 creator_id = 14;
     */
    creatorId?: number;
    /**
     * @generated from protobuf field: optional resources.users.User creator = 15;
     */
    creator?: User;
    /**
     * @generated from protobuf field: repeated resources.centrum.DispatchAssignment units = 16;
     */
    units: DispatchAssignment[];
    /**
     * @generated from protobuf field: optional resources.centrum.DispatchReferences references = 17;
     */
    references?: DispatchReferences;
}
/**
 * @generated from protobuf message resources.centrum.DispatchAssignments
 */
export interface DispatchAssignments {
    /**
     * @generated from protobuf field: uint64 dispatch_id = 1 [jstype = JS_STRING];
     */
    dispatchId: string;
    /**
     * @generated from protobuf field: string job = 2;
     */
    job: string;
    /**
     * @generated from protobuf field: repeated resources.centrum.DispatchAssignment units = 3;
     */
    units: DispatchAssignment[];
}
/**
 * @generated from protobuf message resources.centrum.DispatchAssignment
 */
export interface DispatchAssignment {
    /**
     * @generated from protobuf field: uint64 dispatch_id = 1 [jstype = JS_STRING];
     */
    dispatchId: string; // @gotags: sql:"primary_key" alias:"dispatch_id"
    /**
     * @generated from protobuf field: uint64 unit_id = 2 [jstype = JS_STRING];
     */
    unitId: string; // @gotags: sql:"primary_key" alias:"unit_id"
    /**
     * @generated from protobuf field: optional resources.centrum.Unit unit = 3;
     */
    unit?: Unit;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 4;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp expires_at = 5;
     */
    expiresAt?: Timestamp;
}
/**
 * @generated from protobuf message resources.centrum.DispatchStatus
 */
export interface DispatchStatus {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string; // @gotags: sql:"primary_key" alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 dispatch_id = 3 [jstype = JS_STRING];
     */
    dispatchId: string;
    /**
     * @generated from protobuf field: optional uint64 unit_id = 4 [jstype = JS_STRING];
     */
    unitId?: string;
    /**
     * @generated from protobuf field: optional resources.centrum.Unit unit = 5;
     */
    unit?: Unit;
    /**
     * @generated from protobuf field: resources.centrum.StatusDispatch status = 6;
     */
    status: StatusDispatch;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string reason = 7;
     */
    reason?: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string code = 8;
     */
    code?: string;
    /**
     * @generated from protobuf field: optional int32 user_id = 9;
     */
    userId?: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort user = 10;
     */
    user?: UserShort;
    /**
     * @generated from protobuf field: optional double x = 11;
     */
    x?: number;
    /**
     * @generated from protobuf field: optional double y = 12;
     */
    y?: number;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string postal = 13;
     */
    postal?: string;
}
/**
 * @generated from protobuf message resources.centrum.DispatchReferences
 */
export interface DispatchReferences {
    /**
     * @generated from protobuf field: repeated resources.centrum.DispatchReference references = 1;
     */
    references: DispatchReference[];
}
/**
 * @generated from protobuf message resources.centrum.DispatchReference
 */
export interface DispatchReference {
    /**
     * @generated from protobuf field: uint64 target_dispatch_id = 1;
     */
    targetDispatchId: bigint;
    /**
     * @generated from protobuf field: resources.centrum.DispatchReferenceType reference_type = 2;
     */
    referenceType: DispatchReferenceType;
}
/**
 * @generated from protobuf enum resources.centrum.StatusDispatch
 */
export enum StatusDispatch {
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_NEW = 1;
     */
    NEW = 1,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNASSIGNED = 2;
     */
    UNASSIGNED = 2,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UPDATED = 3;
     */
    UPDATED = 3,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNIT_ASSIGNED = 4;
     */
    UNIT_ASSIGNED = 4,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNIT_UNASSIGNED = 5;
     */
    UNIT_UNASSIGNED = 5,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNIT_ACCEPTED = 6;
     */
    UNIT_ACCEPTED = 6,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_UNIT_DECLINED = 7;
     */
    UNIT_DECLINED = 7,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_EN_ROUTE = 8;
     */
    EN_ROUTE = 8,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_ON_SCENE = 9;
     */
    ON_SCENE = 9,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_NEED_ASSISTANCE = 10;
     */
    NEED_ASSISTANCE = 10,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_COMPLETED = 11;
     */
    COMPLETED = 11,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_CANCELLED = 12;
     */
    CANCELLED = 12,
    /**
     * @generated from protobuf enum value: STATUS_DISPATCH_ARCHIVED = 13;
     */
    ARCHIVED = 13
}
/**
 * @generated from protobuf enum resources.centrum.TakeDispatchResp
 */
export enum TakeDispatchResp {
    /**
     * @generated from protobuf enum value: TAKE_DISPATCH_RESP_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: TAKE_DISPATCH_RESP_TIMEOUT = 1;
     */
    TIMEOUT = 1,
    /**
     * @generated from protobuf enum value: TAKE_DISPATCH_RESP_ACCEPTED = 2;
     */
    ACCEPTED = 2,
    /**
     * @generated from protobuf enum value: TAKE_DISPATCH_RESP_DECLINED = 3;
     */
    DECLINED = 3
}
/**
 * @generated from protobuf enum resources.centrum.DispatchReferenceType
 */
export enum DispatchReferenceType {
    /**
     * @generated from protobuf enum value: DISPATCH_REFERENCE_UNSPECIFIED = 0;
     */
    DISPATCH_REFERENCE_UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: DISPATCH_REFERENCE_REFERENCED = 1;
     */
    DISPATCH_REFERENCE_REFERENCED = 1,
    /**
     * @generated from protobuf enum value: DISPATCH_REFERENCE_DUPLICATED_BY = 2;
     */
    DISPATCH_REFERENCE_DUPLICATED_BY = 2,
    /**
     * @generated from protobuf enum value: DISPATCH_REFERENCE_DUPLICATES = 3;
     */
    DISPATCH_REFERENCE_DUPLICATES = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class Dispatch$Type extends MessageType<Dispatch> {
    constructor() {
        super("resources.centrum.Dispatch", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 5, name: "status", kind: "message", T: () => DispatchStatus },
            { no: 7, name: "message", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 8, name: "description", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "1024" } } } },
            { no: 9, name: "attributes", kind: "message", T: () => Attributes },
            { no: 10, name: "x", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 11, name: "y", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 12, name: "postal", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } },
            { no: 13, name: "anon", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 14, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 15, name: "creator", kind: "message", T: () => User },
            { no: 16, name: "units", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => DispatchAssignment },
            { no: 17, name: "references", kind: "message", T: () => DispatchReferences }
        ]);
    }
    create(value?: PartialMessage<Dispatch>): Dispatch {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.job = "";
        message.message = "";
        message.x = 0;
        message.y = 0;
        message.anon = false;
        message.units = [];
        if (value !== undefined)
            reflectionMergePartial<Dispatch>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Dispatch): Dispatch {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* optional resources.timestamp.Timestamp created_at */ 2:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* optional resources.timestamp.Timestamp updated_at */ 3:
                    message.updatedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.updatedAt);
                    break;
                case /* string job */ 4:
                    message.job = reader.string();
                    break;
                case /* optional resources.centrum.DispatchStatus status */ 5:
                    message.status = DispatchStatus.internalBinaryRead(reader, reader.uint32(), options, message.status);
                    break;
                case /* string message */ 7:
                    message.message = reader.string();
                    break;
                case /* optional string description */ 8:
                    message.description = reader.string();
                    break;
                case /* optional resources.centrum.Attributes attributes */ 9:
                    message.attributes = Attributes.internalBinaryRead(reader, reader.uint32(), options, message.attributes);
                    break;
                case /* double x */ 10:
                    message.x = reader.double();
                    break;
                case /* double y */ 11:
                    message.y = reader.double();
                    break;
                case /* optional string postal */ 12:
                    message.postal = reader.string();
                    break;
                case /* bool anon */ 13:
                    message.anon = reader.bool();
                    break;
                case /* optional int32 creator_id */ 14:
                    message.creatorId = reader.int32();
                    break;
                case /* optional resources.users.User creator */ 15:
                    message.creator = User.internalBinaryRead(reader, reader.uint32(), options, message.creator);
                    break;
                case /* repeated resources.centrum.DispatchAssignment units */ 16:
                    message.units.push(DispatchAssignment.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* optional resources.centrum.DispatchReferences references */ 17:
                    message.references = DispatchReferences.internalBinaryRead(reader, reader.uint32(), options, message.references);
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
    internalBinaryWrite(message: Dispatch, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* optional resources.timestamp.Timestamp created_at = 2; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp updated_at = 3; */
        if (message.updatedAt)
            Timestamp.internalBinaryWrite(message.updatedAt, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* string job = 4; */
        if (message.job !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.job);
        /* optional resources.centrum.DispatchStatus status = 5; */
        if (message.status)
            DispatchStatus.internalBinaryWrite(message.status, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* string message = 7; */
        if (message.message !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.message);
        /* optional string description = 8; */
        if (message.description !== undefined)
            writer.tag(8, WireType.LengthDelimited).string(message.description);
        /* optional resources.centrum.Attributes attributes = 9; */
        if (message.attributes)
            Attributes.internalBinaryWrite(message.attributes, writer.tag(9, WireType.LengthDelimited).fork(), options).join();
        /* double x = 10; */
        if (message.x !== 0)
            writer.tag(10, WireType.Bit64).double(message.x);
        /* double y = 11; */
        if (message.y !== 0)
            writer.tag(11, WireType.Bit64).double(message.y);
        /* optional string postal = 12; */
        if (message.postal !== undefined)
            writer.tag(12, WireType.LengthDelimited).string(message.postal);
        /* bool anon = 13; */
        if (message.anon !== false)
            writer.tag(13, WireType.Varint).bool(message.anon);
        /* optional int32 creator_id = 14; */
        if (message.creatorId !== undefined)
            writer.tag(14, WireType.Varint).int32(message.creatorId);
        /* optional resources.users.User creator = 15; */
        if (message.creator)
            User.internalBinaryWrite(message.creator, writer.tag(15, WireType.LengthDelimited).fork(), options).join();
        /* repeated resources.centrum.DispatchAssignment units = 16; */
        for (let i = 0; i < message.units.length; i++)
            DispatchAssignment.internalBinaryWrite(message.units[i], writer.tag(16, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.centrum.DispatchReferences references = 17; */
        if (message.references)
            DispatchReferences.internalBinaryWrite(message.references, writer.tag(17, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.Dispatch
 */
export const Dispatch = new Dispatch$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DispatchAssignments$Type extends MessageType<DispatchAssignments> {
    constructor() {
        super("resources.centrum.DispatchAssignments", [
            { no: 1, name: "dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 3, name: "units", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => DispatchAssignment }
        ]);
    }
    create(value?: PartialMessage<DispatchAssignments>): DispatchAssignments {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.dispatchId = "0";
        message.job = "";
        message.units = [];
        if (value !== undefined)
            reflectionMergePartial<DispatchAssignments>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DispatchAssignments): DispatchAssignments {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 dispatch_id = 1 [jstype = JS_STRING];*/ 1:
                    message.dispatchId = reader.uint64().toString();
                    break;
                case /* string job */ 2:
                    message.job = reader.string();
                    break;
                case /* repeated resources.centrum.DispatchAssignment units */ 3:
                    message.units.push(DispatchAssignment.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: DispatchAssignments, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 dispatch_id = 1 [jstype = JS_STRING]; */
        if (message.dispatchId !== "0")
            writer.tag(1, WireType.Varint).uint64(message.dispatchId);
        /* string job = 2; */
        if (message.job !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.job);
        /* repeated resources.centrum.DispatchAssignment units = 3; */
        for (let i = 0; i < message.units.length; i++)
            DispatchAssignment.internalBinaryWrite(message.units[i], writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.DispatchAssignments
 */
export const DispatchAssignments = new DispatchAssignments$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DispatchAssignment$Type extends MessageType<DispatchAssignment> {
    constructor() {
        super("resources.centrum.DispatchAssignment", [
            { no: 1, name: "dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 3, name: "unit", kind: "message", T: () => Unit },
            { no: 4, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "expires_at", kind: "message", T: () => Timestamp }
        ]);
    }
    create(value?: PartialMessage<DispatchAssignment>): DispatchAssignment {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.dispatchId = "0";
        message.unitId = "0";
        if (value !== undefined)
            reflectionMergePartial<DispatchAssignment>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DispatchAssignment): DispatchAssignment {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 dispatch_id = 1 [jstype = JS_STRING];*/ 1:
                    message.dispatchId = reader.uint64().toString();
                    break;
                case /* uint64 unit_id = 2 [jstype = JS_STRING];*/ 2:
                    message.unitId = reader.uint64().toString();
                    break;
                case /* optional resources.centrum.Unit unit */ 3:
                    message.unit = Unit.internalBinaryRead(reader, reader.uint32(), options, message.unit);
                    break;
                case /* optional resources.timestamp.Timestamp created_at */ 4:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* optional resources.timestamp.Timestamp expires_at */ 5:
                    message.expiresAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.expiresAt);
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
    internalBinaryWrite(message: DispatchAssignment, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 dispatch_id = 1 [jstype = JS_STRING]; */
        if (message.dispatchId !== "0")
            writer.tag(1, WireType.Varint).uint64(message.dispatchId);
        /* uint64 unit_id = 2 [jstype = JS_STRING]; */
        if (message.unitId !== "0")
            writer.tag(2, WireType.Varint).uint64(message.unitId);
        /* optional resources.centrum.Unit unit = 3; */
        if (message.unit)
            Unit.internalBinaryWrite(message.unit, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp created_at = 4; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp expires_at = 5; */
        if (message.expiresAt)
            Timestamp.internalBinaryWrite(message.expiresAt, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.DispatchAssignment
 */
export const DispatchAssignment = new DispatchAssignment$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DispatchStatus$Type extends MessageType<DispatchStatus> {
    constructor() {
        super("resources.centrum.DispatchStatus", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 4, name: "unit_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/ },
            { no: 5, name: "unit", kind: "message", T: () => Unit },
            { no: 6, name: "status", kind: "enum", T: () => ["resources.centrum.StatusDispatch", StatusDispatch, "STATUS_DISPATCH_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 7, name: "reason", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 8, name: "code", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 9, name: "user_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 10, name: "user", kind: "message", T: () => UserShort },
            { no: 11, name: "x", kind: "scalar", opt: true, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 12, name: "y", kind: "scalar", opt: true, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 13, name: "postal", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } }
        ]);
    }
    create(value?: PartialMessage<DispatchStatus>): DispatchStatus {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.dispatchId = "0";
        message.status = 0;
        if (value !== undefined)
            reflectionMergePartial<DispatchStatus>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DispatchStatus): DispatchStatus {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id = 1 [jstype = JS_STRING];*/ 1:
                    message.id = reader.uint64().toString();
                    break;
                case /* optional resources.timestamp.Timestamp created_at */ 2:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* uint64 dispatch_id = 3 [jstype = JS_STRING];*/ 3:
                    message.dispatchId = reader.uint64().toString();
                    break;
                case /* optional uint64 unit_id = 4 [jstype = JS_STRING];*/ 4:
                    message.unitId = reader.uint64().toString();
                    break;
                case /* optional resources.centrum.Unit unit */ 5:
                    message.unit = Unit.internalBinaryRead(reader, reader.uint32(), options, message.unit);
                    break;
                case /* resources.centrum.StatusDispatch status */ 6:
                    message.status = reader.int32();
                    break;
                case /* optional string reason */ 7:
                    message.reason = reader.string();
                    break;
                case /* optional string code */ 8:
                    message.code = reader.string();
                    break;
                case /* optional int32 user_id */ 9:
                    message.userId = reader.int32();
                    break;
                case /* optional resources.users.UserShort user */ 10:
                    message.user = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.user);
                    break;
                case /* optional double x */ 11:
                    message.x = reader.double();
                    break;
                case /* optional double y */ 12:
                    message.y = reader.double();
                    break;
                case /* optional string postal */ 13:
                    message.postal = reader.string();
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
    internalBinaryWrite(message: DispatchStatus, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* optional resources.timestamp.Timestamp created_at = 2; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* uint64 dispatch_id = 3 [jstype = JS_STRING]; */
        if (message.dispatchId !== "0")
            writer.tag(3, WireType.Varint).uint64(message.dispatchId);
        /* optional uint64 unit_id = 4 [jstype = JS_STRING]; */
        if (message.unitId !== undefined)
            writer.tag(4, WireType.Varint).uint64(message.unitId);
        /* optional resources.centrum.Unit unit = 5; */
        if (message.unit)
            Unit.internalBinaryWrite(message.unit, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* resources.centrum.StatusDispatch status = 6; */
        if (message.status !== 0)
            writer.tag(6, WireType.Varint).int32(message.status);
        /* optional string reason = 7; */
        if (message.reason !== undefined)
            writer.tag(7, WireType.LengthDelimited).string(message.reason);
        /* optional string code = 8; */
        if (message.code !== undefined)
            writer.tag(8, WireType.LengthDelimited).string(message.code);
        /* optional int32 user_id = 9; */
        if (message.userId !== undefined)
            writer.tag(9, WireType.Varint).int32(message.userId);
        /* optional resources.users.UserShort user = 10; */
        if (message.user)
            UserShort.internalBinaryWrite(message.user, writer.tag(10, WireType.LengthDelimited).fork(), options).join();
        /* optional double x = 11; */
        if (message.x !== undefined)
            writer.tag(11, WireType.Bit64).double(message.x);
        /* optional double y = 12; */
        if (message.y !== undefined)
            writer.tag(12, WireType.Bit64).double(message.y);
        /* optional string postal = 13; */
        if (message.postal !== undefined)
            writer.tag(13, WireType.LengthDelimited).string(message.postal);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.DispatchStatus
 */
export const DispatchStatus = new DispatchStatus$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DispatchReferences$Type extends MessageType<DispatchReferences> {
    constructor() {
        super("resources.centrum.DispatchReferences", [
            { no: 1, name: "references", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => DispatchReference }
        ]);
    }
    create(value?: PartialMessage<DispatchReferences>): DispatchReferences {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.references = [];
        if (value !== undefined)
            reflectionMergePartial<DispatchReferences>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DispatchReferences): DispatchReferences {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated resources.centrum.DispatchReference references */ 1:
                    message.references.push(DispatchReference.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: DispatchReferences, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated resources.centrum.DispatchReference references = 1; */
        for (let i = 0; i < message.references.length; i++)
            DispatchReference.internalBinaryWrite(message.references[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.DispatchReferences
 */
export const DispatchReferences = new DispatchReferences$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DispatchReference$Type extends MessageType<DispatchReference> {
    constructor() {
        super("resources.centrum.DispatchReference", [
            { no: 1, name: "target_dispatch_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "reference_type", kind: "enum", T: () => ["resources.centrum.DispatchReferenceType", DispatchReferenceType], options: { "validate.rules": { enum: { definedOnly: true } } } }
        ]);
    }
    create(value?: PartialMessage<DispatchReference>): DispatchReference {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.targetDispatchId = 0n;
        message.referenceType = 0;
        if (value !== undefined)
            reflectionMergePartial<DispatchReference>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DispatchReference): DispatchReference {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 target_dispatch_id */ 1:
                    message.targetDispatchId = reader.uint64().toBigInt();
                    break;
                case /* resources.centrum.DispatchReferenceType reference_type */ 2:
                    message.referenceType = reader.int32();
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
    internalBinaryWrite(message: DispatchReference, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 target_dispatch_id = 1; */
        if (message.targetDispatchId !== 0n)
            writer.tag(1, WireType.Varint).uint64(message.targetDispatchId);
        /* resources.centrum.DispatchReferenceType reference_type = 2; */
        if (message.referenceType !== 0)
            writer.tag(2, WireType.Varint).int32(message.referenceType);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.DispatchReference
 */
export const DispatchReference = new DispatchReference$Type();
