// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "resources/centrum/settings.proto" (package "resources.centrum", syntax proto3)
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
/**
 * @generated from protobuf message resources.centrum.Settings
 */
export interface Settings {
    /**
     * @generated from protobuf field: string job = 1;
     */
    job: string;
    /**
     * @generated from protobuf field: bool enabled = 2;
     */
    enabled: boolean;
    /**
     * @generated from protobuf field: resources.centrum.CentrumMode mode = 3;
     */
    mode: CentrumMode;
    /**
     * @generated from protobuf field: resources.centrum.CentrumMode fallback_mode = 4;
     */
    fallbackMode: CentrumMode;
    /**
     * @generated from protobuf field: optional resources.centrum.PredefinedStatus predefined_status = 5;
     */
    predefinedStatus?: PredefinedStatus;
}
/**
 * @generated from protobuf message resources.centrum.PredefinedStatus
 */
export interface PredefinedStatus {
    /**
     * @generated from protobuf field: repeated string unit_status = 1;
     */
    unitStatus: string[];
    /**
     * @generated from protobuf field: repeated string dispatch_status = 2;
     */
    dispatchStatus: string[];
}
/**
 * @generated from protobuf enum resources.centrum.CentrumMode
 */
export enum CentrumMode {
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_MANUAL = 1;
     */
    MANUAL = 1,
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_CENTRAL_COMMAND = 2;
     */
    CENTRAL_COMMAND = 2,
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_AUTO_ROUND_ROBIN = 3;
     */
    AUTO_ROUND_ROBIN = 3,
    /**
     * @generated from protobuf enum value: CENTRUM_MODE_SIMPLIFIED = 4;
     */
    SIMPLIFIED = 4
}
// @generated message type with reflection information, may provide speed optimized methods
class Settings$Type extends MessageType<Settings> {
    constructor() {
        super("resources.centrum.Settings", [
            { no: 1, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 2, name: "enabled", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "mode", kind: "enum", T: () => ["resources.centrum.CentrumMode", CentrumMode, "CENTRUM_MODE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 4, name: "fallback_mode", kind: "enum", T: () => ["resources.centrum.CentrumMode", CentrumMode, "CENTRUM_MODE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 5, name: "predefined_status", kind: "message", T: () => PredefinedStatus }
        ]);
    }
    create(value?: PartialMessage<Settings>): Settings {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.job = "";
        message.enabled = false;
        message.mode = 0;
        message.fallbackMode = 0;
        if (value !== undefined)
            reflectionMergePartial<Settings>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Settings): Settings {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string job */ 1:
                    message.job = reader.string();
                    break;
                case /* bool enabled */ 2:
                    message.enabled = reader.bool();
                    break;
                case /* resources.centrum.CentrumMode mode */ 3:
                    message.mode = reader.int32();
                    break;
                case /* resources.centrum.CentrumMode fallback_mode */ 4:
                    message.fallbackMode = reader.int32();
                    break;
                case /* optional resources.centrum.PredefinedStatus predefined_status */ 5:
                    message.predefinedStatus = PredefinedStatus.internalBinaryRead(reader, reader.uint32(), options, message.predefinedStatus);
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
    internalBinaryWrite(message: Settings, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string job = 1; */
        if (message.job !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.job);
        /* bool enabled = 2; */
        if (message.enabled !== false)
            writer.tag(2, WireType.Varint).bool(message.enabled);
        /* resources.centrum.CentrumMode mode = 3; */
        if (message.mode !== 0)
            writer.tag(3, WireType.Varint).int32(message.mode);
        /* resources.centrum.CentrumMode fallback_mode = 4; */
        if (message.fallbackMode !== 0)
            writer.tag(4, WireType.Varint).int32(message.fallbackMode);
        /* optional resources.centrum.PredefinedStatus predefined_status = 5; */
        if (message.predefinedStatus)
            PredefinedStatus.internalBinaryWrite(message.predefinedStatus, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.Settings
 */
export const Settings = new Settings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PredefinedStatus$Type extends MessageType<PredefinedStatus> {
    constructor() {
        super("resources.centrum.PredefinedStatus", [
            { no: 1, name: "unit_status", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { repeated: { maxItems: "5" } } } },
            { no: 2, name: "dispatch_status", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { repeated: { maxItems: "5" } } } }
        ]);
    }
    create(value?: PartialMessage<PredefinedStatus>): PredefinedStatus {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.unitStatus = [];
        message.dispatchStatus = [];
        if (value !== undefined)
            reflectionMergePartial<PredefinedStatus>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PredefinedStatus): PredefinedStatus {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated string unit_status */ 1:
                    message.unitStatus.push(reader.string());
                    break;
                case /* repeated string dispatch_status */ 2:
                    message.dispatchStatus.push(reader.string());
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
    internalBinaryWrite(message: PredefinedStatus, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated string unit_status = 1; */
        for (let i = 0; i < message.unitStatus.length; i++)
            writer.tag(1, WireType.LengthDelimited).string(message.unitStatus[i]);
        /* repeated string dispatch_status = 2; */
        for (let i = 0; i < message.dispatchStatus.length; i++)
            writer.tag(2, WireType.LengthDelimited).string(message.dispatchStatus[i]);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.centrum.PredefinedStatus
 */
export const PredefinedStatus = new PredefinedStatus$Type();
