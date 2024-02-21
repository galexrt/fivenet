// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "resources/vehicles/vehicles.proto" (package "resources.vehicles", syntax proto3)
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
/**
 * @generated from protobuf message resources.vehicles.Vehicle
 */
export interface Vehicle {
    /**
     * @generated from protobuf field: string plate = 1;
     */
    plate: string;
    /**
     * @generated from protobuf field: string model = 2;
     */
    model: string;
    /**
     * @generated from protobuf field: string type = 3;
     */
    type: string;
    /**
     * @generated from protobuf field: resources.users.UserShort owner = 4;
     */
    owner?: UserShort;
}
// @generated message type with reflection information, may provide speed optimized methods
class Vehicle$Type extends MessageType<Vehicle> {
    constructor() {
        super("resources.vehicles.Vehicle", [
            { no: 1, name: "plate", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "model", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "owner", kind: "message", T: () => UserShort }
        ]);
    }
    create(value?: PartialMessage<Vehicle>): Vehicle {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.plate = "";
        message.model = "";
        message.type = "";
        if (value !== undefined)
            reflectionMergePartial<Vehicle>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Vehicle): Vehicle {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string plate */ 1:
                    message.plate = reader.string();
                    break;
                case /* string model */ 2:
                    message.model = reader.string();
                    break;
                case /* string type */ 3:
                    message.type = reader.string();
                    break;
                case /* resources.users.UserShort owner */ 4:
                    message.owner = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.owner);
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
    internalBinaryWrite(message: Vehicle, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string plate = 1; */
        if (message.plate !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.plate);
        /* string model = 2; */
        if (message.model !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.model);
        /* string type = 3; */
        if (message.type !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.type);
        /* resources.users.UserShort owner = 4; */
        if (message.owner)
            UserShort.internalBinaryWrite(message.owner, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.vehicles.Vehicle
 */
export const Vehicle = new Vehicle$Type();
