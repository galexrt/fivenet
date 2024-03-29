import {
    CalendarCheckIcon,
    CalendarRemoveIcon,
    CancelIcon,
    CarBackIcon,
    CheckBoldIcon,
    CoffeeIcon,
    HelpCircleIcon,
    MarkerCheckIcon,
} from 'mdi-vue3';
import { type DefineComponent } from 'vue';
import { StatusDispatch } from '~~/gen/ts/resources/centrum/dispatches';
import { StatusUnit, Unit } from '~~/gen/ts/resources/centrum/units';

export type GroupedUnits = { status: StatusUnit; key: string; units: Unit[] }[];

export function dispatchStatusToFillColor(status: StatusDispatch | undefined): string {
    switch (status) {
        case StatusDispatch.UNSPECIFIED:
        case StatusDispatch.NEW:
        case StatusDispatch.UNASSIGNED:
        case StatusDispatch.UNIT_DECLINED:
            return 'fill-error-600';
        case StatusDispatch.EN_ROUTE:
            return 'fill-info-500';
        case StatusDispatch.ON_SCENE:
            return 'fill-info-700';
        case StatusDispatch.NEED_ASSISTANCE:
            return 'fill-warn-600';
        case StatusDispatch.COMPLETED:
            return 'fill-success-600';
        case StatusDispatch.CANCELLED:
            return 'fill-success-800';
        case StatusDispatch.ARCHIVED:
            return 'fill-base-600';
        case StatusDispatch.UNIT_ACCEPTED:
            return 'fill-info-600';
        default:
            return 'fill-info-500';
    }
}

export function dispatchStatusToBGColor(status: StatusDispatch | undefined): string {
    switch (status) {
        case StatusDispatch.UNSPECIFIED:
        case StatusDispatch.NEW:
        case StatusDispatch.UNASSIGNED:
        case StatusDispatch.UNIT_DECLINED:
            return 'bg-error-600';
        case StatusDispatch.EN_ROUTE:
            return 'bg-info-500';
        case StatusDispatch.ON_SCENE:
            return 'bg-info-700';
        case StatusDispatch.NEED_ASSISTANCE:
            return 'bg-warn-600';
        case StatusDispatch.COMPLETED:
            return 'bg-success-600';
        case StatusDispatch.CANCELLED:
            return 'bg-success-800';
        case StatusDispatch.ARCHIVED:
            return 'bg-base-600';
        case StatusDispatch.UNIT_ACCEPTED:
            return 'bg-info-600';
        default:
            return 'bg-info-500';
    }
}

export const animateStates = [
    StatusDispatch.NEW.valueOf(),
    StatusDispatch.UNASSIGNED.valueOf(),
    StatusDispatch.NEED_ASSISTANCE.valueOf(),
];

export function dispatchStatusAnimate(status: StatusDispatch | undefined): boolean {
    return animateStates.includes((status ?? StatusDispatch.NEW).valueOf());
}

export function unitStatusToBGColor(status: StatusUnit | undefined): string {
    switch (status) {
        case StatusUnit.ON_BREAK:
        case StatusUnit.USER_ADDED:
        case StatusUnit.USER_REMOVED:
            return 'bg-info-500';
        case StatusUnit.AVAILABLE:
            return 'bg-success-600';
        case StatusUnit.BUSY:
            return 'bg-warn-600';
        case StatusUnit.UNSPECIFIED:
        case StatusUnit.UNKNOWN:
        case StatusUnit.UNAVAILABLE:
        default:
            return 'bg-error-600';
    }
}

export const statusOrder = [
    StatusUnit.AVAILABLE,
    StatusUnit.ON_BREAK,
    StatusUnit.BUSY,
    StatusUnit.USER_ADDED,
    StatusUnit.USER_REMOVED,
    StatusUnit.UNAVAILABLE,
    StatusUnit.UNKNOWN,
    StatusUnit.UNSPECIFIED,
];

export const unitStatuses: {
    icon: DefineComponent;
    name: string;
    action?: Function;
    class?: string;
    status?: StatusUnit;
}[] = [
    { icon: markRaw(CancelIcon), name: 'Unavailable', status: StatusUnit.UNAVAILABLE },
    { icon: markRaw(CalendarCheckIcon), name: 'Available', status: StatusUnit.AVAILABLE },
    { icon: markRaw(CoffeeIcon), name: 'On Break', status: StatusUnit.ON_BREAK },
    { icon: markRaw(CalendarRemoveIcon), name: 'Busy', status: StatusUnit.BUSY },
];

export const dispatchStatuses: {
    icon: DefineComponent;
    name: string;
    action?: Function;
    class?: string;
    status?: StatusDispatch;
}[] = [
    { icon: markRaw(CarBackIcon), name: 'En Route', status: StatusDispatch.EN_ROUTE },
    { icon: markRaw(MarkerCheckIcon), name: 'On Scene', status: StatusDispatch.ON_SCENE },
    { icon: markRaw(HelpCircleIcon), name: 'Need Assistance', status: StatusDispatch.NEED_ASSISTANCE },
    { icon: markRaw(CheckBoldIcon), name: 'Completed', status: StatusDispatch.COMPLETED },
    { icon: markRaw(CancelIcon), name: 'Cancelled', status: StatusDispatch.CANCELLED },
];

export function isStatusDispatchCompleted(status: StatusDispatch): boolean {
    return status === StatusDispatch.ARCHIVED || status === StatusDispatch.CANCELLED || status === StatusDispatch.COMPLETED;
}
