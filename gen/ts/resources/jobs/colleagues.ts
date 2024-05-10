// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "resources/jobs/colleagues.proto" (package "resources.jobs", syntax proto3)
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
import { File } from "../filestore/file";
/**
 * @generated from protobuf message resources.jobs.Colleague
 */
export interface Colleague {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: string identifier = 2;
     */
    identifier: string;
    /**
     * @generated from protobuf field: string job = 3;
     */
    job: string;
    /**
     * @generated from protobuf field: optional string job_label = 4;
     */
    jobLabel?: string;
    /**
     * @generated from protobuf field: int32 job_grade = 5;
     */
    jobGrade: number;
    /**
     * @generated from protobuf field: optional string job_grade_label = 6;
     */
    jobGradeLabel?: string;
    /**
     * @generated from protobuf field: string firstname = 7;
     */
    firstname: string;
    /**
     * @generated from protobuf field: string lastname = 8;
     */
    lastname: string;
    /**
     * @generated from protobuf field: string dateofbirth = 9;
     */
    dateofbirth: string;
    /**
     * @generated from protobuf field: optional string phone_number = 12;
     */
    phoneNumber?: string;
    /**
     * @generated from protobuf field: optional resources.filestore.File avatar = 17;
     */
    avatar?: File;
    /**
     * @generated from protobuf field: resources.jobs.JobsUserProps props = 18;
     */
    props?: JobsUserProps; // @gotags: alias:"fivenet_jobs_user_props"
}
/**
 * @generated from protobuf message resources.jobs.JobsUserProps
 */
export interface JobsUserProps {
    /**
     * @generated from protobuf field: int32 user_id = 1;
     */
    userId: number;
    /**
     * @generated from protobuf field: string job = 2;
     */
    job: string;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp absence_begin = 3;
     */
    absenceBegin?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp absence_end = 4;
     */
    absenceEnd?: Timestamp;
    /**
     * @sanitize: method=StripTags
     *
     * @generated from protobuf field: optional string note = 5;
     */
    note?: string;
}
/**
 * @generated from protobuf message resources.jobs.JobsUserActivity
 */
export interface JobsUserActivity {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string; // @gotags: sql:"primary_key" alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: string job = 4;
     */
    job: string;
    /**
     * @generated from protobuf field: int32 source_user_id = 5;
     */
    sourceUserId: number;
    /**
     * @generated from protobuf field: resources.jobs.Colleague source_user = 6;
     */
    sourceUser?: Colleague; // @gotags: alias:"source_user"
    /**
     * @generated from protobuf field: int32 target_user_id = 7;
     */
    targetUserId: number;
    /**
     * @generated from protobuf field: resources.jobs.Colleague target_user = 8;
     */
    targetUser?: Colleague; // @gotags: alias:"target_user"
    /**
     * @generated from protobuf field: resources.jobs.JobsUserActivityType activity_type = 9;
     */
    activityType: JobsUserActivityType;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string reason = 10;
     */
    reason: string;
    /**
     * @generated from protobuf field: resources.jobs.JobsUserActivityData data = 11;
     */
    data?: JobsUserActivityData;
}
/**
 * @generated from protobuf message resources.jobs.JobsUserActivityData
 */
export interface JobsUserActivityData {
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "absenceDate";
        /**
         * @generated from protobuf field: resources.jobs.ColleagueAbsenceDate absence_date = 1;
         */
        absenceDate: ColleagueAbsenceDate;
    } | {
        oneofKind: "gradeChange";
        /**
         * @generated from protobuf field: resources.jobs.ColleagueGradeChange grade_change = 2;
         */
        gradeChange: ColleagueGradeChange;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message resources.jobs.ColleagueAbsenceDate
 */
export interface ColleagueAbsenceDate {
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp absence_begin = 1;
     */
    absenceBegin?: Timestamp;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp absence_end = 2;
     */
    absenceEnd?: Timestamp;
}
/**
 * @generated from protobuf message resources.jobs.ColleagueGradeChange
 */
export interface ColleagueGradeChange {
    /**
     * @generated from protobuf field: int32 grade = 1;
     */
    grade: number;
    /**
     * @generated from protobuf field: string grade_label = 2;
     */
    gradeLabel: string;
}
/**
 * @generated from protobuf enum resources.jobs.JobsUserActivityType
 */
export enum JobsUserActivityType {
    /**
     * @generated from protobuf enum value: JOBS_USER_ACTIVITY_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: JOBS_USER_ACTIVITY_TYPE_HIRED = 1;
     */
    HIRED = 1,
    /**
     * @generated from protobuf enum value: JOBS_USER_ACTIVITY_TYPE_FIRED = 2;
     */
    FIRED = 2,
    /**
     * @generated from protobuf enum value: JOBS_USER_ACTIVITY_TYPE_PROMOTED = 3;
     */
    PROMOTED = 3,
    /**
     * @generated from protobuf enum value: JOBS_USER_ACTIVITY_TYPE_DEMOTED = 4;
     */
    DEMOTED = 4,
    /**
     * @generated from protobuf enum value: JOBS_USER_ACTIVITY_TYPE_ABSENCE_DATE = 5;
     */
    ABSENCE_DATE = 5,
    /**
     * @generated from protobuf enum value: JOBS_USER_ACTIVITY_TYPE_NOTE = 6;
     */
    NOTE = 6
}
// @generated message type with reflection information, may provide speed optimized methods
class Colleague$Type extends MessageType<Colleague> {
    constructor() {
        super("resources.jobs.Colleague", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 2, name: "identifier", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "46" } } } },
            { no: 3, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 4, name: "job_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 5, name: "job_grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: -1 } } } },
            { no: 6, name: "job_grade_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 7, name: "firstname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } },
            { no: 8, name: "lastname", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "1", maxLen: "50" } } } },
            { no: 9, name: "dateofbirth", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "10" } } } },
            { no: 12, name: "phone_number", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 17, name: "avatar", kind: "message", T: () => File },
            { no: 18, name: "props", kind: "message", T: () => JobsUserProps }
        ]);
    }
    create(value?: PartialMessage<Colleague>): Colleague {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        message.identifier = "";
        message.job = "";
        message.jobGrade = 0;
        message.firstname = "";
        message.lastname = "";
        message.dateofbirth = "";
        if (value !== undefined)
            reflectionMergePartial<Colleague>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Colleague): Colleague {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 user_id */ 1:
                    message.userId = reader.int32();
                    break;
                case /* string identifier */ 2:
                    message.identifier = reader.string();
                    break;
                case /* string job */ 3:
                    message.job = reader.string();
                    break;
                case /* optional string job_label */ 4:
                    message.jobLabel = reader.string();
                    break;
                case /* int32 job_grade */ 5:
                    message.jobGrade = reader.int32();
                    break;
                case /* optional string job_grade_label */ 6:
                    message.jobGradeLabel = reader.string();
                    break;
                case /* string firstname */ 7:
                    message.firstname = reader.string();
                    break;
                case /* string lastname */ 8:
                    message.lastname = reader.string();
                    break;
                case /* string dateofbirth */ 9:
                    message.dateofbirth = reader.string();
                    break;
                case /* optional string phone_number */ 12:
                    message.phoneNumber = reader.string();
                    break;
                case /* optional resources.filestore.File avatar */ 17:
                    message.avatar = File.internalBinaryRead(reader, reader.uint32(), options, message.avatar);
                    break;
                case /* resources.jobs.JobsUserProps props */ 18:
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
    internalBinaryWrite(message: Colleague, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 user_id = 1; */
        if (message.userId !== 0)
            writer.tag(1, WireType.Varint).int32(message.userId);
        /* string identifier = 2; */
        if (message.identifier !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.identifier);
        /* string job = 3; */
        if (message.job !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.job);
        /* optional string job_label = 4; */
        if (message.jobLabel !== undefined)
            writer.tag(4, WireType.LengthDelimited).string(message.jobLabel);
        /* int32 job_grade = 5; */
        if (message.jobGrade !== 0)
            writer.tag(5, WireType.Varint).int32(message.jobGrade);
        /* optional string job_grade_label = 6; */
        if (message.jobGradeLabel !== undefined)
            writer.tag(6, WireType.LengthDelimited).string(message.jobGradeLabel);
        /* string firstname = 7; */
        if (message.firstname !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.firstname);
        /* string lastname = 8; */
        if (message.lastname !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.lastname);
        /* string dateofbirth = 9; */
        if (message.dateofbirth !== "")
            writer.tag(9, WireType.LengthDelimited).string(message.dateofbirth);
        /* optional string phone_number = 12; */
        if (message.phoneNumber !== undefined)
            writer.tag(12, WireType.LengthDelimited).string(message.phoneNumber);
        /* optional resources.filestore.File avatar = 17; */
        if (message.avatar)
            File.internalBinaryWrite(message.avatar, writer.tag(17, WireType.LengthDelimited).fork(), options).join();
        /* resources.jobs.JobsUserProps props = 18; */
        if (message.props)
            JobsUserProps.internalBinaryWrite(message.props, writer.tag(18, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.jobs.Colleague
 */
export const Colleague = new Colleague$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobsUserProps$Type extends MessageType<JobsUserProps> {
    constructor() {
        super("resources.jobs.JobsUserProps", [
            { no: 1, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 2, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 3, name: "absence_begin", kind: "message", T: () => Timestamp },
            { no: 4, name: "absence_end", kind: "message", T: () => Timestamp },
            { no: 5, name: "note", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<JobsUserProps>): JobsUserProps {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userId = 0;
        message.job = "";
        if (value !== undefined)
            reflectionMergePartial<JobsUserProps>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JobsUserProps): JobsUserProps {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 user_id */ 1:
                    message.userId = reader.int32();
                    break;
                case /* string job */ 2:
                    message.job = reader.string();
                    break;
                case /* optional resources.timestamp.Timestamp absence_begin */ 3:
                    message.absenceBegin = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.absenceBegin);
                    break;
                case /* optional resources.timestamp.Timestamp absence_end */ 4:
                    message.absenceEnd = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.absenceEnd);
                    break;
                case /* optional string note */ 5:
                    message.note = reader.string();
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
    internalBinaryWrite(message: JobsUserProps, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 user_id = 1; */
        if (message.userId !== 0)
            writer.tag(1, WireType.Varint).int32(message.userId);
        /* string job = 2; */
        if (message.job !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.job);
        /* optional resources.timestamp.Timestamp absence_begin = 3; */
        if (message.absenceBegin)
            Timestamp.internalBinaryWrite(message.absenceBegin, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* optional resources.timestamp.Timestamp absence_end = 4; */
        if (message.absenceEnd)
            Timestamp.internalBinaryWrite(message.absenceEnd, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* optional string note = 5; */
        if (message.note !== undefined)
            writer.tag(5, WireType.LengthDelimited).string(message.note);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.jobs.JobsUserProps
 */
export const JobsUserProps = new JobsUserProps$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobsUserActivity$Type extends MessageType<JobsUserActivity> {
    constructor() {
        super("resources.jobs.JobsUserActivity", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 5, name: "source_user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 6, name: "source_user", kind: "message", T: () => Colleague },
            { no: 7, name: "target_user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 8, name: "target_user", kind: "message", T: () => Colleague },
            { no: 9, name: "activity_type", kind: "enum", T: () => ["resources.jobs.JobsUserActivityType", JobsUserActivityType, "JOBS_USER_ACTIVITY_TYPE_"] },
            { no: 10, name: "reason", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 11, name: "data", kind: "message", T: () => JobsUserActivityData }
        ]);
    }
    create(value?: PartialMessage<JobsUserActivity>): JobsUserActivity {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "0";
        message.job = "";
        message.sourceUserId = 0;
        message.targetUserId = 0;
        message.activityType = 0;
        message.reason = "";
        if (value !== undefined)
            reflectionMergePartial<JobsUserActivity>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JobsUserActivity): JobsUserActivity {
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
                case /* string job */ 4:
                    message.job = reader.string();
                    break;
                case /* int32 source_user_id */ 5:
                    message.sourceUserId = reader.int32();
                    break;
                case /* resources.jobs.Colleague source_user */ 6:
                    message.sourceUser = Colleague.internalBinaryRead(reader, reader.uint32(), options, message.sourceUser);
                    break;
                case /* int32 target_user_id */ 7:
                    message.targetUserId = reader.int32();
                    break;
                case /* resources.jobs.Colleague target_user */ 8:
                    message.targetUser = Colleague.internalBinaryRead(reader, reader.uint32(), options, message.targetUser);
                    break;
                case /* resources.jobs.JobsUserActivityType activity_type */ 9:
                    message.activityType = reader.int32();
                    break;
                case /* string reason */ 10:
                    message.reason = reader.string();
                    break;
                case /* resources.jobs.JobsUserActivityData data */ 11:
                    message.data = JobsUserActivityData.internalBinaryRead(reader, reader.uint32(), options, message.data);
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
    internalBinaryWrite(message: JobsUserActivity, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1 [jstype = JS_STRING]; */
        if (message.id !== "0")
            writer.tag(1, WireType.Varint).uint64(message.id);
        /* optional resources.timestamp.Timestamp created_at = 2; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* string job = 4; */
        if (message.job !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.job);
        /* int32 source_user_id = 5; */
        if (message.sourceUserId !== 0)
            writer.tag(5, WireType.Varint).int32(message.sourceUserId);
        /* resources.jobs.Colleague source_user = 6; */
        if (message.sourceUser)
            Colleague.internalBinaryWrite(message.sourceUser, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* int32 target_user_id = 7; */
        if (message.targetUserId !== 0)
            writer.tag(7, WireType.Varint).int32(message.targetUserId);
        /* resources.jobs.Colleague target_user = 8; */
        if (message.targetUser)
            Colleague.internalBinaryWrite(message.targetUser, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* resources.jobs.JobsUserActivityType activity_type = 9; */
        if (message.activityType !== 0)
            writer.tag(9, WireType.Varint).int32(message.activityType);
        /* string reason = 10; */
        if (message.reason !== "")
            writer.tag(10, WireType.LengthDelimited).string(message.reason);
        /* resources.jobs.JobsUserActivityData data = 11; */
        if (message.data)
            JobsUserActivityData.internalBinaryWrite(message.data, writer.tag(11, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.jobs.JobsUserActivity
 */
export const JobsUserActivity = new JobsUserActivity$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobsUserActivityData$Type extends MessageType<JobsUserActivityData> {
    constructor() {
        super("resources.jobs.JobsUserActivityData", [
            { no: 1, name: "absence_date", kind: "message", oneof: "data", T: () => ColleagueAbsenceDate },
            { no: 2, name: "grade_change", kind: "message", oneof: "data", T: () => ColleagueGradeChange }
        ]);
    }
    create(value?: PartialMessage<JobsUserActivityData>): JobsUserActivityData {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.data = { oneofKind: undefined };
        if (value !== undefined)
            reflectionMergePartial<JobsUserActivityData>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JobsUserActivityData): JobsUserActivityData {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.jobs.ColleagueAbsenceDate absence_date */ 1:
                    message.data = {
                        oneofKind: "absenceDate",
                        absenceDate: ColleagueAbsenceDate.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).absenceDate)
                    };
                    break;
                case /* resources.jobs.ColleagueGradeChange grade_change */ 2:
                    message.data = {
                        oneofKind: "gradeChange",
                        gradeChange: ColleagueGradeChange.internalBinaryRead(reader, reader.uint32(), options, (message.data as any).gradeChange)
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
    internalBinaryWrite(message: JobsUserActivityData, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.jobs.ColleagueAbsenceDate absence_date = 1; */
        if (message.data.oneofKind === "absenceDate")
            ColleagueAbsenceDate.internalBinaryWrite(message.data.absenceDate, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* resources.jobs.ColleagueGradeChange grade_change = 2; */
        if (message.data.oneofKind === "gradeChange")
            ColleagueGradeChange.internalBinaryWrite(message.data.gradeChange, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.jobs.JobsUserActivityData
 */
export const JobsUserActivityData = new JobsUserActivityData$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ColleagueAbsenceDate$Type extends MessageType<ColleagueAbsenceDate> {
    constructor() {
        super("resources.jobs.ColleagueAbsenceDate", [
            { no: 1, name: "absence_begin", kind: "message", T: () => Timestamp },
            { no: 2, name: "absence_end", kind: "message", T: () => Timestamp }
        ]);
    }
    create(value?: PartialMessage<ColleagueAbsenceDate>): ColleagueAbsenceDate {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<ColleagueAbsenceDate>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ColleagueAbsenceDate): ColleagueAbsenceDate {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* resources.timestamp.Timestamp absence_begin */ 1:
                    message.absenceBegin = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.absenceBegin);
                    break;
                case /* resources.timestamp.Timestamp absence_end */ 2:
                    message.absenceEnd = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.absenceEnd);
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
    internalBinaryWrite(message: ColleagueAbsenceDate, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* resources.timestamp.Timestamp absence_begin = 1; */
        if (message.absenceBegin)
            Timestamp.internalBinaryWrite(message.absenceBegin, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* resources.timestamp.Timestamp absence_end = 2; */
        if (message.absenceEnd)
            Timestamp.internalBinaryWrite(message.absenceEnd, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.jobs.ColleagueAbsenceDate
 */
export const ColleagueAbsenceDate = new ColleagueAbsenceDate$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ColleagueGradeChange$Type extends MessageType<ColleagueGradeChange> {
    constructor() {
        super("resources.jobs.ColleagueGradeChange", [
            { no: 1, name: "grade", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "grade_label", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<ColleagueGradeChange>): ColleagueGradeChange {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.grade = 0;
        message.gradeLabel = "";
        if (value !== undefined)
            reflectionMergePartial<ColleagueGradeChange>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ColleagueGradeChange): ColleagueGradeChange {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 grade */ 1:
                    message.grade = reader.int32();
                    break;
                case /* string grade_label */ 2:
                    message.gradeLabel = reader.string();
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
    internalBinaryWrite(message: ColleagueGradeChange, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 grade = 1; */
        if (message.grade !== 0)
            writer.tag(1, WireType.Varint).int32(message.grade);
        /* string grade_label = 2; */
        if (message.gradeLabel !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.gradeLabel);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.jobs.ColleagueGradeChange
 */
export const ColleagueGradeChange = new ColleagueGradeChange$Type();
