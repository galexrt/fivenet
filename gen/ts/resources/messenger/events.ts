// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/messenger/events.proto" (package "resources.messenger", syntax proto3)
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
import { Message } from "./message";
import { Thread } from "./thread";
/**
 * @generated from protobuf message resources.messenger.MessengerEvent
 */
export interface MessengerEvent {
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "threadUpdate";
        /**
         * @generated from protobuf field: resources.messenger.Thread thread_update = 1;
         */
        threadUpdate: Thread;
    } | {
        oneofKind: "threadDelete";
        /**
         * @generated from protobuf field: uint64 thread_delete = 2 [jstype = JS_STRING];
         */
        threadDelete: string;
    } | {
        oneofKind: "messageUpdate";
        /**
         * @generated from protobuf field: resources.messenger.Message message_update = 3;
         */
        messageUpdate: Message;
    } | {
        oneofKind: "messageDelete";
        /**
         * @generated from protobuf field: uint64 message_delete = 4 [jstype = JS_STRING];
         */
        messageDelete: string;
    } | {
        oneofKind: undefined;
    };
}
// @generated message type with reflection information, may provide speed optimized methods
class MessengerEvent$Type extends MessageType<MessengerEvent> {
    constructor() {
        super("resources.messenger.MessengerEvent", [
            { no: 1, name: "thread_update", kind: "message", oneof: "data", T: () => Thread },
            { no: 2, name: "thread_delete", kind: "scalar", oneof: "data", T: 4 /*ScalarType.UINT64*/ },
            { no: 3, name: "message_update", kind: "message", oneof: "data", T: () => Message },
            { no: 4, name: "message_delete", kind: "scalar", oneof: "data", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
    create(value?: PartialMessage<MessengerEvent>): MessengerEvent {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.data = { oneofKind: undefined };
        if (value !== undefined)
            reflectionMergePartial<MessengerEvent>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MessengerEvent): MessengerEvent {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.messenger.Thread thread_update */ 1:
                    message.data = {
                        oneofKind: "threadUpdate",
                        threadUpdate: Thread.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).threadUpdate)
                    };
                    break;
                case /* uint64 thread_delete = 2 [jstype = JS_STRING];*/ 2:
                    message.data = {
                        oneofKind: "threadDelete",
                        threadDelete: reader.uint64().toString()
                    };
                    break;
                case /* resources.messenger.Message message_update */ 3:
                    message.data = {
                        oneofKind: "messageUpdate",
                        messageUpdate: Message.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).messageUpdate)
                    };
                    break;
                case /* uint64 message_delete = 4 [jstype = JS_STRING];*/ 4:
                    message.data = {
                        oneofKind: "messageDelete",
                        messageDelete: reader.uint64().toString()
                    };
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
    internalBinaryWrite(message: MessengerEvent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.messenger.Thread thread_update = 1; */
        if (message.data.oneofKind === "threadUpdate")
            Thread.internalBinaryWrite(message.data.threadUpdate, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* uint64 thread_delete = 2 [jstype = JS_STRING]; */
        if (message.data.oneofKind === "threadDelete")
            writer.tag(2, WireType.Varint).uint64(message.data.threadDelete);
        /* resources.messenger.Message message_update = 3; */
        if (message.data.oneofKind === "messageUpdate")
            Message.internalBinaryWrite(message.data.messageUpdate, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* uint64 message_delete = 4 [jstype = JS_STRING]; */
        if (message.data.oneofKind === "messageDelete")
            writer.tag(4, WireType.Varint).uint64(message.data.messageDelete);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.messenger.MessengerEvent
 */
export const MessengerEvent = new MessengerEvent$Type();
