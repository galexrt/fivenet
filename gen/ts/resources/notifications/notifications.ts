// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/notifications/notifications.proto" (package "resources.notifications", syntax proto3)
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
import { TranslateItem } from "../common/i18n";
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.notifications.Notification
 */
export interface Notification {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp read_at = 3;
     */
    readAt?: Timestamp;
    /**
     * @generated from protobuf field: int32 user_id = 4;
     */
    userId: number;
    /**
     * @sanitize
     *
     * @generated from protobuf field: resources.common.TranslateItem title = 5;
     */
    title?: TranslateItem;
    /**
     * @generated from protobuf field: resources.notifications.NotificationType type = 6;
     */
    type: NotificationType;
    /**
     * @sanitize
     *
     * @generated from protobuf field: resources.common.TranslateItem content = 7;
     */
    content?: TranslateItem;
    /**
     * @generated from protobuf field: resources.notifications.NotificationCategory category = 8;
     */
    category: NotificationCategory;
    /**
     * @generated from protobuf field: optional resources.notifications.Data data = 9;
     */
    data?: Data;
    /**
     * @generated from protobuf field: optional bool starred = 10;
     */
    starred?: boolean;
}
/**
 * @generated from protobuf message resources.notifications.Data
 */
export interface Data {
    /**
     * @generated from protobuf field: optional resources.notifications.Link link = 1;
     */
    link?: Link;
    /**
     * @generated from protobuf field: optional resources.users.UserShort caused_by = 2;
     */
    causedBy?: UserShort;
    /**
     * @generated from protobuf field: optional resources.notifications.CalendarData calendar = 3;
     */
    calendar?: CalendarData;
}
/**
 * @generated from protobuf message resources.notifications.Link
 */
export interface Link {
    /**
     * @generated from protobuf field: string to = 1;
     */
    to: string;
    /**
     * @generated from protobuf field: optional string title = 2;
     */
    title?: string;
    /**
     * @generated from protobuf field: optional bool external = 3;
     */
    external?: boolean;
}
/**
 * @generated from protobuf message resources.notifications.CalendarData
 */
export interface CalendarData {
    /**
     * @generated from protobuf field: optional uint64 calendar_id = 1 [jstype = JS_STRING];
     */
    calendarId?: string;
    /**
     * @generated from protobuf field: optional uint64 calendar_entry_id = 2 [jstype = JS_STRING];
     */
    calendarEntryId?: string;
}
/**
 * @generated from protobuf enum resources.notifications.NotificationType
 */
export enum NotificationType {
    /**
     * @generated from protobuf enum value: NOTIFICATION_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: NOTIFICATION_TYPE_ERROR = 1;
     */
    ERROR = 1,
    /**
     * @generated from protobuf enum value: NOTIFICATION_TYPE_WARNING = 2;
     */
    WARNING = 2,
    /**
     * @generated from protobuf enum value: NOTIFICATION_TYPE_INFO = 3;
     */
    INFO = 3,
    /**
     * @generated from protobuf enum value: NOTIFICATION_TYPE_SUCCESS = 4;
     */
    SUCCESS = 4
}
/**
 * @generated from protobuf enum resources.notifications.NotificationCategory
 */
export enum NotificationCategory {
    /**
     * @generated from protobuf enum value: NOTIFICATION_CATEGORY_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: NOTIFICATION_CATEGORY_GENERAL = 1;
     */
    GENERAL = 1,
    /**
     * @generated from protobuf enum value: NOTIFICATION_CATEGORY_DOCUMENT = 2;
     */
    DOCUMENT = 2,
    /**
     * @generated from protobuf enum value: NOTIFICATION_CATEGORY_CALENDAR = 3;
     */
    CALENDAR = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class Notification$Type extends MessageType<Notification> {
    constructor() {
        super("resources.notifications.Notification", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "read_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "title", kind: "message", T: () => TranslateItem },
            { no: 6, name: "type", kind: "enum", T: () => ["resources.notifications.NotificationType", NotificationType, "NOTIFICATION_TYPE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 7, name: "content", kind: "message", T: () => TranslateItem },
            { no: 8, name: "category", kind: "enum", T: () => ["resources.notifications.NotificationCategory", NotificationCategory, "NOTIFICATION_CATEGORY_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 9, name: "data", kind: "message", T: () => Data },
            { no: 10, name: "starred", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<Notification>): Notification {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.userId = 0;
        message.type = 0;
        message.category = 0;
        if (value !== undefined)
            reflectionMergePartial<Notification>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Notification): Notification {
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
                case /* resources.timestamp.Timestamp read_at */ 3:
                    message.readAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.readAt);
                    break;
                case /* int32 user_id */ 4:
                    message.userId = reader.int32();
                    break;
                case /* resources.common.TranslateItem title */ 5:
                    message.title = TranslateItem.internalBinaryRead(reader, reader.uint32(), options, message.title);
                    break;
                case /* resources.notifications.NotificationType type */ 6:
                    message.type = reader.int32();
                    break;
                case /* resources.common.TranslateItem content */ 7:
                    message.content = TranslateItem.internalBinaryRead(reader, reader.uint32(), options, message.content);
                    break;
                case /* resources.notifications.NotificationCategory category */ 8:
                    message.category = reader.int32();
                    break;
                case /* optional resources.notifications.Data data */ 9:
                    message.data = Data.internalBinaryRead(reader, reader.uint32(), options, message.data);
                    break;
                case /* optional bool starred */ 10:
                    message.starred = reader.bool();
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
    internalBinaryWrite(message: Notification, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* resources.timestamp.Timestamp created_at = 2; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* resources.timestamp.Timestamp read_at = 3; */
        if (message.readAt)
            Timestamp.internalBinaryWrite(message.readAt, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* int32 user_id = 4; */
        if (message.userId !== 0)
            writer.tag(4, WireType.Varint).int32(message.userId);
        /* resources.common.TranslateItem title = 5; */
        if (message.title)
            TranslateItem.internalBinaryWrite(message.title, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* resources.notifications.NotificationType type = 6; */
        if (message.type !== 0)
            writer.tag(6, WireType.Varint).int32(message.type);
        /* resources.common.TranslateItem content = 7; */
        if (message.content)
            TranslateItem.internalBinaryWrite(message.content, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* resources.notifications.NotificationCategory category = 8; */
        if (message.category !== 0)
            writer.tag(8, WireType.Varint).int32(message.category);
        /* optional resources.notifications.Data data = 9; */
        if (message.data)
            Data.internalBinaryWrite(message.data, writer.tag(9, WireType.LengthDelimited).fork(), options).join();
        /* optional bool starred = 10; */
        if (message.starred !== undefined)
            writer.tag(10, WireType.Varint).bool(message.starred);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.notifications.Notification
 */
export const Notification = new Notification$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Data$Type extends MessageType<Data> {
    constructor() {
        super("resources.notifications.Data", [
            { no: 1, name: "link", kind: "message", T: () => Link },
            { no: 2, name: "caused_by", kind: "message", T: () => UserShort },
            { no: 3, name: "calendar", kind: "message", T: () => CalendarData }
        ]);
    }
    create(value?: PartialMessage<Data>): Data {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<Data>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Data): Data {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* optional resources.notifications.Link link */ 1:
                    message.link = Link.internalBinaryRead(reader, reader.uint32(), options, message.link);
                    break;
                case /* optional resources.users.UserShort caused_by */ 2:
                    message.causedBy = UserShort.internalBinaryRead(reader, reader.uint32(), options, message.causedBy);
                    break;
                case /* optional resources.notifications.CalendarData calendar */ 3:
                    message.calendar = CalendarData.internalBinaryRead(reader, reader.uint32(), options, message.calendar);
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
    internalBinaryWrite(message: Data, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* optional resources.notifications.Link link = 1; */
        if (message.link)
            Link.internalBinaryWrite(message.link, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.users.UserShort caused_by = 2; */
        if (message.causedBy)
            UserShort.internalBinaryWrite(message.causedBy, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.notifications.CalendarData calendar = 3; */
        if (message.calendar)
            CalendarData.internalBinaryWrite(message.calendar, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.notifications.Data
 */
export const Data = new Data$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Link$Type extends MessageType<Link> {
    constructor() {
        super("resources.notifications.Link", [
            { no: 1, name: "to", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "title", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "external", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<Link>): Link {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.to = "";
        if (value !== undefined)
            reflectionMergePartial<Link>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Link): Link {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string to */ 1:
                    message.to = reader.string();
                    break;
                case /* optional string title */ 2:
                    message.title = reader.string();
                    break;
                case /* optional bool external */ 3:
                    message.external = reader.bool();
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
    internalBinaryWrite(message: Link, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string to = 1; */
        if (message.to !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.to);
        /* optional string title = 2; */
        if (message.title !== undefined)
            writer.tag(2, WireType.LengthDelimited).string(message.title);
        /* optional bool external = 3; */
        if (message.external !== undefined)
            writer.tag(3, WireType.Varint).bool(message.external);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.notifications.Link
 */
export const Link = new Link$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CalendarData$Type extends MessageType<CalendarData> {
    constructor() {
        super("resources.notifications.CalendarData", [
            { no: 1, name: "calendar_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "calendar_entry_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
    create(value?: PartialMessage<CalendarData>): CalendarData {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<CalendarData>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CalendarData): CalendarData {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* optional uint64 calendar_id = 1 [jstype = JS_STRING];*/ 1:
                    message.calendarId = reader.uint64().toString();
                    break;
                case /* optional uint64 calendar_entry_id = 2 [jstype = JS_STRING];*/ 2:
                    message.calendarEntryId = reader.uint64().toString();
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
    internalBinaryWrite(message: CalendarData, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* optional uint64 calendar_id = 1 [jstype = JS_STRING]; */
        if (message.calendarId !== undefined)
            writer.tag(1, WireType.Varint).uint64(message.calendarId);
        /* optional uint64 calendar_entry_id = 2 [jstype = JS_STRING]; */
        if (message.calendarEntryId !== undefined)
            writer.tag(2, WireType.Varint).uint64(message.calendarEntryId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.notifications.CalendarData
 */
export const CalendarData = new CalendarData$Type();
