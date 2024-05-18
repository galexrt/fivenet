// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/messenger/mail.proto" (package "resources.messenger", syntax proto3)
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
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.messenger.Mail
 */
export interface Mail {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp deleted_at = 4;
     */
    deletedAt?: Timestamp;
    /**
     * @generated from protobuf field: bool unread = 5;
     */
    unread: boolean;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string subject = 6;
     */
    subject: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string body = 9;
     */
    body: string;
    /**
     * @generated from protobuf field: optional int32 from_id = 11;
     */
    fromId?: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort from = 12;
     */
    from?: UserShort; // @gotags: alias:"from"
}
// @generated message type with reflection information, may provide speed optimized methods
class Mail$Type extends MessageType<Mail> {
    constructor() {
        super("resources.messenger.Mail", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "deleted_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "unread", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "subject", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "1024" } } } },
            { no: 9, name: "body", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "20", maxBytes: "1750000" } } } },
            { no: 11, name: "from_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "from", kind: "message", T: () => UserShort }
        ]);
    }
    create(value?: PartialMessage<Mail>): Mail {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.unread = false;
        message.subject = "";
        message.body = "";
        if (value !== undefined)
            reflectionMergePartial<Mail>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mail): Mail {
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
                case /* optional resources.timestamp.Timestamp updated_at */ 3:
                    message.updatedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.updatedAt);
                    break;
                case /* optional resources.timestamp.Timestamp deleted_at */ 4:
                    message.deletedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.deletedAt);
                    break;
                case /* bool unread */ 5:
                    message.unread = reader.bool();
                    break;
                case /* string subject */ 6:
                    message.subject = reader.string();
                    break;
                case /* string body */ 9:
                    message.body = reader.string();
                    break;
                case /* optional int32 from_id */ 11:
                    message.fromId = reader.int32();
                    break;
                case /* optional resources.users.UserShort from */ 12:
                    message.from = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.from);
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
    internalBinaryWrite(message: Mail, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* resources.timestamp.Timestamp created_at = 2; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp updated_at = 3; */
        if (message.updatedAt)
            Timestamp.internalBinaryWrite(message.updatedAt, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp deleted_at = 4; */
        if (message.deletedAt)
            Timestamp.internalBinaryWrite(message.deletedAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* bool unread = 5; */
        if (message.unread !== false)
            writer.tag(5, WireType.Varint).bool(message.unread);
        /* string subject = 6; */
        if (message.subject !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.subject);
        /* string body = 9; */
        if (message.body !== "")
            writer.tag(9, WireType.LengthDelimited).string(message.body);
        /* optional int32 from_id = 11; */
        if (message.fromId !== undefined)
            writer.tag(11, WireType.Varint).int32(message.fromId);
        /* optional resources.users.UserShort from = 12; */
        if (message.from)
            UserShort.internalBinaryWrite(message.from, writer.tag(12, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.messenger.Mail
 */
export const Mail = new Mail$Type();