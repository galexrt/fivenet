// Code generated by protoc-gen-customizerweb. DO NOT EDIT.
// source: resources/users/users.proto

import * as enums from './users_pb';

// USER_ACTIVITY_TYPE
export class USER_ACTIVITY_TYPE_Util {
    public static toEnumKey(input: enums.USER_ACTIVITY_TYPE): string | undefined {
        const index = Object.values(enums.USER_ACTIVITY_TYPE).indexOf(input);
        if (index <= -1) {
            return "N/A";
        }
        return Object.keys(enums.USER_ACTIVITY_TYPE)[index];
    }
}
