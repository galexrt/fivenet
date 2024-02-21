// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "resources/centrum/general.proto" (package "resources.centrum", syntax proto3)
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
import { Timestamp } from "../timestamp/timestamp";
import { UserShort } from "../users/users";
/**
 * @generated from protobuf message resources.centrum.Attributes
 */
export interface Attributes {
    /**
     * @generated from protobuf field: repeated string list = 1;
     */
    list: string[];
}
/**
 * @generated from protobuf message resources.centrum.Disponents
 */
export interface Disponents {
    /**
     * @generated from protobuf field: string job = 1;
     */
    job: string;
    /**
     * @generated from protobuf field: repeated resources.users.UserShort disponents = 2;
     */
    disponents: UserShort[];
}
/**
 * @generated from protobuf message resources.centrum.UserUnitMapping
 */
export interface UserUnitMapping {
    /**
     * @generated from protobuf field: uint64 unit_id = 1;
     */
    unitId: bigint;
    /**
     * @generated from protobuf field: string job = 2;
     */
    job: string;
    /**
     * @generated from protobuf field: int32 user_id = 3;
     */
    userId: number;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 4;
     */
    createdAt?: Timestamp;
}
// @generated message type with reflection information, may provide speed optimized methods
class Attributes$Type extends MessageType<Attributes> {
    constructor() {
        super("resources.centrum.Attributes", [
            { no: 1, name: "list", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Attributes>): Attributes {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.list = [];
        if (value !== undefined)
            reflectionMergePartial<Attributes>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Attributes): Attributes {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated string list */ 1:
                    message.list.push(reader.string());
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
    internalBinaryWrite(message: Attributes, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated string list = 1; */
        for (let i = 0; i < message.list.length; i++)
            writer.tag(1, WireType.LengthDelimited).string(message.list[i]);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.Attributes
 */
export const Attributes = new Attributes$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Disponents$Type extends MessageType<Disponents> {
    constructor() {
        super("resources.centrum.Disponents", [
            { no: 1, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 2, name: "disponents", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UserShort }
        ]);
    }
    create(value?: PartialMessage<Disponents>): Disponents {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.job = "";
        message.disponents = [];
        if (value !== undefined)
            reflectionMergePartial<Disponents>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Disponents): Disponents {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string job */ 1:
                    message.job = reader.string();
                    break;
                case /* repeated resources.users.UserShort disponents */ 2:
                    message.disponents.push(UserShort.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: Disponents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string job = 1; */
        if (message.job !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.job);
        /* repeated resources.users.UserShort disponents = 2; */
        for (let i = 0; i < message.disponents.length; i++)
            UserShort.internalBinaryWrite(message.disponents[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.Disponents
 */
export const Disponents = new Disponents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserUnitMapping$Type extends MessageType<UserUnitMapping> {
    constructor() {
        super("resources.centrum.UserUnitMapping", [
            { no: 1, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "created_at", kind: "message", T: () => Timestamp }
        ]);
    }
    create(value?: PartialMessage<UserUnitMapping>): UserUnitMapping {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.unitId = 0n;
        message.job = "";
        message.userId = 0;
        if (value !== undefined)
            reflectionMergePartial<UserUnitMapping>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserUnitMapping): UserUnitMapping {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 unit_id */ 1:
                    message.unitId = reader.uint64().toBigInt();
                    break;
                case /* string job */ 2:
                    message.job = reader.string();
                    break;
                case /* int32 user_id */ 3:
                    message.userId = reader.int32();
                    break;
                case /* resources.timestamp.Timestamp created_at */ 4:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
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
    internalBinaryWrite(message: UserUnitMapping, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 unit_id = 1; */
        if (message.unitId !== 0n)
            writer.tag(1, WireType.Varint).uint64(message.unitId);
        /* string job = 2; */
        if (message.job !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.job);
        /* int32 user_id = 3; */
        if (message.userId !== 0)
            writer.tag(3, WireType.Varint).int32(message.userId);
        /* resources.timestamp.Timestamp created_at = 4; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.UserUnitMapping
 */
export const UserUnitMapping = new UserUnitMapping$Type();
