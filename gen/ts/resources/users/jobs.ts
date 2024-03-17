// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "resources/users/jobs.proto" (package "resources.users", syntax proto3)
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
import { File } from "../filestore/file";
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.users.Job
 */
export interface Job {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string; // @gotags: sql:"primary_key" alias:"name"
    /**
     * @generated from protobuf field: string label = 2;
     */
    label: string;
    /**
     * @generated from protobuf field: repeated resources.users.JobGrade grades = 3;
     */
    grades: JobGrade[];
}
/**
 * @generated from protobuf message resources.users.JobGrade
 */
export interface JobGrade {
    /**
     * @generated from protobuf field: optional string job_name = 1;
     */
    jobName?: string;
    /**
     * @generated from protobuf field: int32 grade = 2;
     */
    grade: number;
    /**
     * @generated from protobuf field: string label = 3;
     */
    label: string;
}
/**
 * @generated from protobuf message resources.users.JobProps
 */
export interface JobProps {
    /**
     * @generated from protobuf field: string job = 1;
     */
    job: string;
    /**
     * @generated from protobuf field: optional string job_label = 2;
     */
    jobLabel?: string;
    /**
     * @generated from protobuf field: string theme = 3;
     */
    theme: string;
    /**
     * @generated from protobuf field: string livemap_marker_color = 4;
     */
    livemapMarkerColor: string;
    /**
     * @generated from protobuf field: resources.users.QuickButtons quick_buttons = 5;
     */
    quickButtons?: QuickButtons;
    /**
     * @generated from protobuf field: optional string radio_frequency = 6;
     */
    radioFrequency?: string;
    /**
     * @generated from protobuf field: optional uint64 discord_guild_id = 7 [jstype = JS_STRING];
     */
    discordGuildId?: string;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp discord_last_sync = 8;
     */
    discordLastSync?: Timestamp;
    /**
     * @generated from protobuf field: resources.users.DiscordSyncSettings discord_sync_settings = 9;
     */
    discordSyncSettings?: DiscordSyncSettings;
    /**
     * @generated from protobuf field: optional string motd = 10;
     */
    motd?: string;
    /**
     * @generated from protobuf field: optional resources.filestore.File logo_url = 11;
     */
    logoUrl?: File;
    /**
     * @generated from protobuf field: resources.users.JobSettings settings = 12;
     */
    settings?: JobSettings;
}
/**
 * @generated from protobuf message resources.users.QuickButtons
 */
export interface QuickButtons {
    /**
     * @generated from protobuf field: bool penalty_calculator = 1;
     */
    penaltyCalculator: boolean;
    /**
     * @generated from protobuf field: bool body_checkup = 2;
     */
    bodyCheckup: boolean;
}
/**
 * @generated from protobuf message resources.users.DiscordSyncSettings
 */
export interface DiscordSyncSettings {
    /**
     * @generated from protobuf field: bool user_info_sync = 1;
     */
    userInfoSync: boolean;
    /**
     * @generated from protobuf field: optional resources.users.UserInfoSyncSettings user_info_sync_settings = 2;
     */
    userInfoSyncSettings?: UserInfoSyncSettings;
    /**
     * @generated from protobuf field: bool status_log = 3;
     */
    statusLog: boolean;
    /**
     * @generated from protobuf field: optional resources.users.StatusLogSettings status_log_settings = 4;
     */
    statusLogSettings?: StatusLogSettings;
    /**
     * @generated from protobuf field: bool jobs_absence = 5;
     */
    jobsAbsence: boolean;
    /**
     * @generated from protobuf field: optional resources.users.JobsAbsenceSettings jobs_absence_settings = 6;
     */
    jobsAbsenceSettings?: JobsAbsenceSettings;
}
/**
 * @generated from protobuf message resources.users.UserInfoSyncSettings
 */
export interface UserInfoSyncSettings {
    /**
     * @generated from protobuf field: bool employee_role_enabled = 1;
     */
    employeeRoleEnabled: boolean;
    /**
     * @generated from protobuf field: optional string employee_role_format = 2;
     */
    employeeRoleFormat?: string;
    /**
     * @generated from protobuf field: optional string grade_role_format = 3;
     */
    gradeRoleFormat?: string;
    /**
     * @generated from protobuf field: bool unemployed_enabled = 4;
     */
    unemployedEnabled: boolean;
    /**
     * @generated from protobuf field: resources.users.UserInfoSyncUnemployedMode unemployed_mode = 5;
     */
    unemployedMode: UserInfoSyncUnemployedMode;
    /**
     * @generated from protobuf field: optional string unemployed_role_name = 6;
     */
    unemployedRoleName?: string;
}
/**
 * @generated from protobuf message resources.users.StatusLogSettings
 */
export interface StatusLogSettings {
    /**
     * @generated from protobuf field: optional string channel_id = 1;
     */
    channelId?: string;
}
/**
 * @generated from protobuf message resources.users.JobsAbsenceSettings
 */
export interface JobsAbsenceSettings {
    /**
     * @generated from protobuf field: string absence_role = 1;
     */
    absenceRole: string;
}
/**
 * @generated from protobuf message resources.users.JobSettings
 */
export interface JobSettings {
}
/**
 * @generated from protobuf enum resources.users.UserInfoSyncUnemployedMode
 */
export enum UserInfoSyncUnemployedMode {
    /**
     * @generated from protobuf enum value: USER_INFO_SYNC_UNEMPLOYED_MODE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: USER_INFO_SYNC_UNEMPLOYED_MODE_GIVE_ROLE = 1;
     */
    GIVE_ROLE = 1,
    /**
     * @generated from protobuf enum value: USER_INFO_SYNC_UNEMPLOYED_MODE_KICK = 2;
     */
    KICK = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class Job$Type extends MessageType<Job> {
    constructor() {
        super("resources.users.Job", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 2, name: "label", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 3, name: "grades", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => JobGrade }
        ]);
    }
    create(value?: PartialMessage<Job>): Job {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.name = "";
        message.label = "";
        message.grades = [];
        if (value !== undefined)
            reflectionMergePartial<Job>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Job): Job {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* string label */ 2:
                    message.label = reader.string();
                    break;
                case /* repeated resources.users.JobGrade grades */ 3:
                    message.grades.push(JobGrade.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: Job, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* string label = 2; */
        if (message.label !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.label);
        /* repeated resources.users.JobGrade grades = 3; */
        for (let i = 0; i < message.grades.length; i++)
            JobGrade.internalBinaryWrite(message.grades[i], writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.Job
 */
export const Job = new Job$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobGrade$Type extends MessageType<JobGrade> {
    constructor() {
        super("resources.users.JobGrade", [
            { no: 1, name: "job_name", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 2, name: "grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 3, name: "label", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } }
        ]);
    }
    create(value?: PartialMessage<JobGrade>): JobGrade {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.grade = 0;
        message.label = "";
        if (value !== undefined)
            reflectionMergePartial<JobGrade>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JobGrade): JobGrade {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* optional string job_name */ 1:
                    message.jobName = reader.string();
                    break;
                case /* int32 grade */ 2:
                    message.grade = reader.int32();
                    break;
                case /* string label */ 3:
                    message.label = reader.string();
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
    internalBinaryWrite(message: JobGrade, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* optional string job_name = 1; */
        if (message.jobName !== undefined)
            writer.tag(1, WireType.LengthDelimited).string(message.jobName);
        /* int32 grade = 2; */
        if (message.grade !== 0)
            writer.tag(2, WireType.Varint).int32(message.grade);
        /* string label = 3; */
        if (message.label !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.label);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.JobGrade
 */
export const JobGrade = new JobGrade$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobProps$Type extends MessageType<JobProps> {
    constructor() {
        super("resources.users.JobProps", [
            { no: 1, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 2, name: "job_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 3, name: "theme", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 4, name: "livemap_marker_color", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "7", pattern: "^#[A-Fa-f0-9]{6}$" } } } },
            { no: 5, name: "quick_buttons", kind: "message", T: () => QuickButtons },
            { no: 6, name: "radio_frequency", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "24" } } } },
            { no: 7, name: "discord_guild_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/ },
            { no: 8, name: "discord_last_sync", kind: "message", T: () => Timestamp },
            { no: 9, name: "discord_sync_settings", kind: "message", T: () => DiscordSyncSettings },
            { no: 10, name: "motd", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "1024" } } } },
            { no: 11, name: "logo_url", kind: "message", T: () => File },
            { no: 12, name: "settings", kind: "message", T: () => JobSettings }
        ]);
    }
    create(value?: PartialMessage<JobProps>): JobProps {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.job = "";
        message.theme = "";
        message.livemapMarkerColor = "";
        if (value !== undefined)
            reflectionMergePartial<JobProps>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JobProps): JobProps {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string job */ 1:
                    message.job = reader.string();
                    break;
                case /* optional string job_label */ 2:
                    message.jobLabel = reader.string();
                    break;
                case /* string theme */ 3:
                    message.theme = reader.string();
                    break;
                case /* string livemap_marker_color */ 4:
                    message.livemapMarkerColor = reader.string();
                    break;
                case /* resources.users.QuickButtons quick_buttons */ 5:
                    message.quickButtons = QuickButtons.internalBinaryRead(reader, reader.uint32(), options, message.quickButtons);
                    break;
                case /* optional string radio_frequency */ 6:
                    message.radioFrequency = reader.string();
                    break;
                case /* optional uint64 discord_guild_id = 7 [jstype = JS_STRING];*/ 7:
                    message.discordGuildId = reader.uint64().toString();
                    break;
                case /* optional resources.timestamp.Timestamp discord_last_sync */ 8:
                    message.discordLastSync = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.discordLastSync);
                    break;
                case /* resources.users.DiscordSyncSettings discord_sync_settings */ 9:
                    message.discordSyncSettings = DiscordSyncSettings.internalBinaryRead(reader, reader.uint32(), options, message.discordSyncSettings);
                    break;
                case /* optional string motd */ 10:
                    message.motd = reader.string();
                    break;
                case /* optional resources.filestore.File logo_url */ 11:
                    message.logoUrl = File.internalBinaryRead(reader, reader.uint32(), options, message.logoUrl);
                    break;
                case /* resources.users.JobSettings settings */ 12:
                    message.settings = JobSettings.internalBinaryRead(reader, reader.uint32(), options, message.settings);
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
    internalBinaryWrite(message: JobProps, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string job = 1; */
        if (message.job !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.job);
        /* optional string job_label = 2; */
        if (message.jobLabel !== undefined)
            writer.tag(2, WireType.LengthDelimited).string(message.jobLabel);
        /* string theme = 3; */
        if (message.theme !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.theme);
        /* string livemap_marker_color = 4; */
        if (message.livemapMarkerColor !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.livemapMarkerColor);
        /* resources.users.QuickButtons quick_buttons = 5; */
        if (message.quickButtons)
            QuickButtons.internalBinaryWrite(message.quickButtons, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* optional string radio_frequency = 6; */
        if (message.radioFrequency !== undefined)
            writer.tag(6, WireType.LengthDelimited).string(message.radioFrequency);
        /* optional uint64 discord_guild_id = 7 [jstype = JS_STRING]; */
        if (message.discordGuildId !== undefined)
            writer.tag(7, WireType.Varint).uint64(message.discordGuildId);
        /* optional resources.timestamp.Timestamp discord_last_sync = 8; */
        if (message.discordLastSync)
            Timestamp.internalBinaryWrite(message.discordLastSync, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* resources.users.DiscordSyncSettings discord_sync_settings = 9; */
        if (message.discordSyncSettings)
            DiscordSyncSettings.internalBinaryWrite(message.discordSyncSettings, writer.tag(9, WireType.LengthDelimited).fork(), options).join();
        /* optional string motd = 10; */
        if (message.motd !== undefined)
            writer.tag(10, WireType.LengthDelimited).string(message.motd);
        /* optional resources.filestore.File logo_url = 11; */
        if (message.logoUrl)
            File.internalBinaryWrite(message.logoUrl, writer.tag(11, WireType.LengthDelimited).fork(), options).join();
        /* resources.users.JobSettings settings = 12; */
        if (message.settings)
            JobSettings.internalBinaryWrite(message.settings, writer.tag(12, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.JobProps
 */
export const JobProps = new JobProps$Type();
// @generated message type with reflection information, may provide speed optimized methods
class QuickButtons$Type extends MessageType<QuickButtons> {
    constructor() {
        super("resources.users.QuickButtons", [
            { no: 1, name: "penalty_calculator", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "body_checkup", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<QuickButtons>): QuickButtons {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.penaltyCalculator = false;
        message.bodyCheckup = false;
        if (value !== undefined)
            reflectionMergePartial<QuickButtons>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: QuickButtons): QuickButtons {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool penalty_calculator */ 1:
                    message.penaltyCalculator = reader.bool();
                    break;
                case /* bool body_checkup */ 2:
                    message.bodyCheckup = reader.bool();
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
    internalBinaryWrite(message: QuickButtons, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool penalty_calculator = 1; */
        if (message.penaltyCalculator !== false)
            writer.tag(1, WireType.Varint).bool(message.penaltyCalculator);
        /* bool body_checkup = 2; */
        if (message.bodyCheckup !== false)
            writer.tag(2, WireType.Varint).bool(message.bodyCheckup);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.QuickButtons
 */
export const QuickButtons = new QuickButtons$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DiscordSyncSettings$Type extends MessageType<DiscordSyncSettings> {
    constructor() {
        super("resources.users.DiscordSyncSettings", [
            { no: 1, name: "user_info_sync", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "user_info_sync_settings", kind: "message", T: () => UserInfoSyncSettings },
            { no: 3, name: "status_log", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "status_log_settings", kind: "message", T: () => StatusLogSettings },
            { no: 5, name: "jobs_absence", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "jobs_absence_settings", kind: "message", T: () => JobsAbsenceSettings }
        ]);
    }
    create(value?: PartialMessage<DiscordSyncSettings>): DiscordSyncSettings {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.userInfoSync = false;
        message.statusLog = false;
        message.jobsAbsence = false;
        if (value !== undefined)
            reflectionMergePartial<DiscordSyncSettings>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DiscordSyncSettings): DiscordSyncSettings {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool user_info_sync */ 1:
                    message.userInfoSync = reader.bool();
                    break;
                case /* optional resources.users.UserInfoSyncSettings user_info_sync_settings */ 2:
                    message.userInfoSyncSettings = UserInfoSyncSettings.internalBinaryRead(reader, reader.uint32(), options, message.userInfoSyncSettings);
                    break;
                case /* bool status_log */ 3:
                    message.statusLog = reader.bool();
                    break;
                case /* optional resources.users.StatusLogSettings status_log_settings */ 4:
                    message.statusLogSettings = StatusLogSettings.internalBinaryRead(reader, reader.uint32(), options, message.statusLogSettings);
                    break;
                case /* bool jobs_absence */ 5:
                    message.jobsAbsence = reader.bool();
                    break;
                case /* optional resources.users.JobsAbsenceSettings jobs_absence_settings */ 6:
                    message.jobsAbsenceSettings = JobsAbsenceSettings.internalBinaryRead(reader, reader.uint32(), options, message.jobsAbsenceSettings);
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
    internalBinaryWrite(message: DiscordSyncSettings, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool user_info_sync = 1; */
        if (message.userInfoSync !== false)
            writer.tag(1, WireType.Varint).bool(message.userInfoSync);
        /* optional resources.users.UserInfoSyncSettings user_info_sync_settings = 2; */
        if (message.userInfoSyncSettings)
            UserInfoSyncSettings.internalBinaryWrite(message.userInfoSyncSettings, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* bool status_log = 3; */
        if (message.statusLog !== false)
            writer.tag(3, WireType.Varint).bool(message.statusLog);
        /* optional resources.users.StatusLogSettings status_log_settings = 4; */
        if (message.statusLogSettings)
            StatusLogSettings.internalBinaryWrite(message.statusLogSettings, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* bool jobs_absence = 5; */
        if (message.jobsAbsence !== false)
            writer.tag(5, WireType.Varint).bool(message.jobsAbsence);
        /* optional resources.users.JobsAbsenceSettings jobs_absence_settings = 6; */
        if (message.jobsAbsenceSettings)
            JobsAbsenceSettings.internalBinaryWrite(message.jobsAbsenceSettings, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.DiscordSyncSettings
 */
export const DiscordSyncSettings = new DiscordSyncSettings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserInfoSyncSettings$Type extends MessageType<UserInfoSyncSettings> {
    constructor() {
        super("resources.users.UserInfoSyncSettings", [
            { no: 1, name: "employee_role_enabled", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "employee_role_format", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } },
            { no: 3, name: "grade_role_format", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } },
            { no: 4, name: "unemployed_enabled", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "unemployed_mode", kind: "enum", T: () => ["resources.users.UserInfoSyncUnemployedMode", UserInfoSyncUnemployedMode, "USER_INFO_SYNC_UNEMPLOYED_MODE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 6, name: "unemployed_role_name", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } }
        ]);
    }
    create(value?: PartialMessage<UserInfoSyncSettings>): UserInfoSyncSettings {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.employeeRoleEnabled = false;
        message.unemployedEnabled = false;
        message.unemployedMode = 0;
        if (value !== undefined)
            reflectionMergePartial<UserInfoSyncSettings>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserInfoSyncSettings): UserInfoSyncSettings {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool employee_role_enabled */ 1:
                    message.employeeRoleEnabled = reader.bool();
                    break;
                case /* optional string employee_role_format */ 2:
                    message.employeeRoleFormat = reader.string();
                    break;
                case /* optional string grade_role_format */ 3:
                    message.gradeRoleFormat = reader.string();
                    break;
                case /* bool unemployed_enabled */ 4:
                    message.unemployedEnabled = reader.bool();
                    break;
                case /* resources.users.UserInfoSyncUnemployedMode unemployed_mode */ 5:
                    message.unemployedMode = reader.int32();
                    break;
                case /* optional string unemployed_role_name */ 6:
                    message.unemployedRoleName = reader.string();
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
    internalBinaryWrite(message: UserInfoSyncSettings, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool employee_role_enabled = 1; */
        if (message.employeeRoleEnabled !== false)
            writer.tag(1, WireType.Varint).bool(message.employeeRoleEnabled);
        /* optional string employee_role_format = 2; */
        if (message.employeeRoleFormat !== undefined)
            writer.tag(2, WireType.LengthDelimited).string(message.employeeRoleFormat);
        /* optional string grade_role_format = 3; */
        if (message.gradeRoleFormat !== undefined)
            writer.tag(3, WireType.LengthDelimited).string(message.gradeRoleFormat);
        /* bool unemployed_enabled = 4; */
        if (message.unemployedEnabled !== false)
            writer.tag(4, WireType.Varint).bool(message.unemployedEnabled);
        /* resources.users.UserInfoSyncUnemployedMode unemployed_mode = 5; */
        if (message.unemployedMode !== 0)
            writer.tag(5, WireType.Varint).int32(message.unemployedMode);
        /* optional string unemployed_role_name = 6; */
        if (message.unemployedRoleName !== undefined)
            writer.tag(6, WireType.LengthDelimited).string(message.unemployedRoleName);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.UserInfoSyncSettings
 */
export const UserInfoSyncSettings = new UserInfoSyncSettings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StatusLogSettings$Type extends MessageType<StatusLogSettings> {
    constructor() {
        super("resources.users.StatusLogSettings", [
            { no: 1, name: "channel_id", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<StatusLogSettings>): StatusLogSettings {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<StatusLogSettings>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StatusLogSettings): StatusLogSettings {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* optional string channel_id */ 1:
                    message.channelId = reader.string();
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
    internalBinaryWrite(message: StatusLogSettings, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* optional string channel_id = 1; */
        if (message.channelId !== undefined)
            writer.tag(1, WireType.LengthDelimited).string(message.channelId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.StatusLogSettings
 */
export const StatusLogSettings = new StatusLogSettings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobsAbsenceSettings$Type extends MessageType<JobsAbsenceSettings> {
    constructor() {
        super("resources.users.JobsAbsenceSettings", [
            { no: 1, name: "absence_role", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } }
        ]);
    }
    create(value?: PartialMessage<JobsAbsenceSettings>): JobsAbsenceSettings {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.absenceRole = "";
        if (value !== undefined)
            reflectionMergePartial<JobsAbsenceSettings>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JobsAbsenceSettings): JobsAbsenceSettings {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string absence_role */ 1:
                    message.absenceRole = reader.string();
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
    internalBinaryWrite(message: JobsAbsenceSettings, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string absence_role = 1; */
        if (message.absenceRole !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.absenceRole);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.JobsAbsenceSettings
 */
export const JobsAbsenceSettings = new JobsAbsenceSettings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class JobSettings$Type extends MessageType<JobSettings> {
    constructor() {
        super("resources.users.JobSettings", []);
    }
    create(value?: PartialMessage<JobSettings>): JobSettings {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<JobSettings>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: JobSettings): JobSettings {
        return target ?? this.create();
    }
    internalBinaryWrite(message: JobSettings, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message resources.users.JobSettings
 */
export const JobSettings = new JobSettings$Type();
