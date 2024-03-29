import { type TypedRouteFromName } from '@typed-router';
import { useConfigStore } from '~/store/config';

// Checking for `GetParentResourceName` existance doesn't work (anymore) in FiveM NUI iframes
export function isNUIAvailable(): boolean {
    return useConfigStore().isNUIAvailable;
}

function getParentResourceName(): string {
    return useConfigStore().nuiResourceName ?? 'fivenet';
}

export async function fetchNUI<T = any, V = any>(event: string, data: T): Promise<V> {
    const body = jsonStringify(data);
    console.debug(`NUI: Fetch ${event}: ${body}`);
    // @ts-ignore FiveM NUI functions
    const resp = await fetch(`https://${getParentResourceName()}/${event}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json; charset=UTF-8',
        },
        body,
    });

    const parsed = resp.json();
    return parsed as V;
}

type NUIMessage =
    | {
          type: 'navigateTo';
          route: TypedRouteFromName<any>;
      }
    | {
          type: undefined;
      };

export async function onNUIMessage(event: MessageEvent<NUIMessage>): Promise<void> {
    if (event.data.type === 'navigateTo') {
        await navigateTo(event.data.route);
    } else {
        console.error('NUI Message: Unknown message type received', event.data);
    }
}

// NUI Callbacks

export async function toggleTablet(state: boolean): Promise<void> {
    if (!isNUIAvailable()) {
        return;
    }

    return await fetchNUI(state ? 'openTablet' : 'closeTablet', { ok: true });
}

export async function focusTablet(state: boolean): Promise<void> {
    if (!isNUIAvailable()) {
        return;
    }

    return await fetchNUI('focusTablet', { state });
}

export async function setWaypoint(x: number, y: number): Promise<void> {
    if (!isNUIAvailable()) {
        return;
    }

    return fetchNUI('setWaypoint', { x, y });
}

export async function phoneCallNumber(phoneNumber: string): Promise<void> {
    if (!isNUIAvailable()) {
        return;
    }

    return fetchNUI('phoneCallNumber', { phoneNumber });
}

export async function copyToClipboard(text: string): Promise<void> {
    if (!isNUIAvailable()) {
        return;
    }

    return fetchNUI('copyToClipboard', { text });
}

export async function setRadioFrequency(frequency: string): Promise<void> {
    if (!isNUIAvailable()) {
        return;
    }

    return fetchNUI('setRadioFrequency', { frequency });
}

export async function setWaypointPLZ(plz: string): Promise<void> {
    if (!isNUIAvailable()) {
        return;
    }

    return fetchNUI('setWaypointPLZ', { plz });
}
