import { defineStore, type StoreDefinition } from 'pinia';
import { CalendarEntry, RsvpResponses, type Calendar } from '~~/gen/ts/resources/calendar/calendar';
import type { UserShort } from '~~/gen/ts/resources/users/users';
import type {
    CreateOrUpdateCalendarEntryResponse,
    CreateOrUpdateCalendarResponse,
    GetCalendarEntryRequest,
    GetCalendarEntryResponse,
    GetCalendarRequest,
    GetCalendarResponse,
    ListCalendarEntriesRequest,
    ListCalendarEntriesResponse,
    ListCalendarEntryRSVPRequest,
    ListCalendarEntryRSVPResponse,
    ListCalendarsRequest,
    ListCalendarsResponse,
    RSVPCalendarEntryRequest,
    RSVPCalendarEntryResponse,
} from '~~/gen/ts/services/calendar/calendar';
import { useAuthStore } from './auth';

export interface CalendarState {
    activeCalendarIds: string[];
    weeklyView: boolean;
    calendars: Calendar[];
    entries: CalendarEntry[];
}

export const useCalendarStore = defineStore('calendar', {
    state: () =>
        ({
            activeCalendarIds: [],
            weeklyView: false,
            calendars: [],
            entries: [],
        }) as CalendarState,
    persist: {
        key(id) {
            return `state-${useAuthStore().activeChar?.userId}-${id}`;
        },
        paths: ['activeCalendarIds', 'weeklyView'],
    },
    actions: {
        // Calendars
        async getCalendar(req: GetCalendarRequest): Promise<GetCalendarResponse> {
            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getCalendarClient().getCalendar(req);
                const { response } = await call;

                if (response.calendar) {
                    const idx = this.calendars.findIndex((c) => c.id === response.calendar!.id);
                    if (idx > -1) {
                        this.calendars[idx] = response.calendar;
                    } else {
                        this.calendars.push(response.calendar);
                    }
                }

                return response;
            } catch (e) {
                throw e;
            }
        },
        async listCalendars(req: ListCalendarsRequest): Promise<ListCalendarsResponse> {
            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getCalendarClient().listCalendars(req);
                const { response } = await call;

                // Only "register" calendars in list when they are accessible by the user
                if (!req.onlyPublic) {
                    if (response.calendars.length > 0) {
                        response.calendars.forEach((calendar) => {
                            const idx = this.calendars.findIndex((c) => c.id === calendar!.id);
                            if (idx > -1) {
                                this.calendars[idx] = calendar;
                            } else {
                                this.calendars.push(calendar);
                            }
                        });
                    } else {
                        this.calendars.length = 0;
                    }
                }

                return response;
            } catch (e) {
                $grpc.handleError(e as RpcError);
                throw e;
            }
        },
        async createOrUpdateCalendar(calendar: Calendar): Promise<CreateOrUpdateCalendarResponse> {
            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getCalendarClient().createOrUpdateCalendar({
                    calendar: calendar,
                });
                const { response } = await call;

                if (response.calendar) {
                    const idx = this.calendars.findIndex((c) => c.id === response.calendar!.id);
                    if (idx > -1) {
                        this.calendars[idx] = response.calendar;
                    } else {
                        this.calendars.push(response.calendar);
                    }

                    this.activeCalendarIds.push(response.calendar.id);
                }

                return response;
            } catch (e) {
                throw e;
            }
        },
        async deleteCalendar(id: string): Promise<void> {
            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getCalendarClient().deleteCalendar({
                    calendarId: id,
                });
                await call;

                const idx = this.calendars.findIndex((c) => c.id === id);
                if (idx > -1) {
                    this.calendars.splice(idx, 1);
                }
            } catch (e) {
                $grpc.handleError(e as RpcError);
                throw e;
            }
        },

        // Entries
        async getCalendarEntry(req: GetCalendarEntryRequest): Promise<GetCalendarEntryResponse> {
            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getCalendarClient().getCalendarEntry(req);
                const { response } = await call;

                if (response.entry) {
                    const idx = this.entries.findIndex((c) => c.id === response.entry!.id);
                    if (idx > -1) {
                        this.entries[idx] = response.entry;
                    } else {
                        this.entries.push(response.entry);
                    }
                }

                return response;
            } catch (e) {
                throw e;
            }
        },
        async listCalendarEntries(req: ListCalendarEntriesRequest): Promise<ListCalendarEntriesResponse> {
            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getCalendarClient().listCalendarEntries(req);
                const { response } = await call;

                if (response.entries.length > 0) {
                    response.entries.forEach((entry) => {
                        const idx = this.entries.findIndex((c) => c.id === entry!.id);
                        if (idx > -1) {
                            this.entries[idx] = entry;
                        } else {
                            this.entries.push(entry);
                        }
                    });
                } else {
                    this.entries.length = 0;
                }

                return response;
            } catch (e) {
                $grpc.handleError(e as RpcError);
                throw e;
            }
        },
        async createOrUpdateCalendarEntry(
            entry: CalendarEntry,
            users?: UserShort[],
        ): Promise<CreateOrUpdateCalendarEntryResponse> {
            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getCalendarClient().createOrUpdateCalendarEntry({
                    entry: entry,
                });
                const { response } = await call;

                if (response.entry) {
                    const idx = this.entries.findIndex((e) => e.id === response.entry?.id);
                    if (idx > -1) {
                        this.entries[idx] = response.entry;
                    } else {
                        this.entries.push(response.entry);
                    }
                }

                if (users && users.length > 0) {
                    const call = $grpc.getCalendarClient().shareCalendarEntry({
                        entryId: entry.id,
                        userIds: users.map((u) => u.userId),
                    });
                    await call;
                }

                return response;
            } catch (e) {
                throw e;
            }
        },

        async deleteCalendarEntry(entryId: string): Promise<void> {
            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getCalendarClient().deleteCalendarEntry({
                    entryId: entryId,
                });
                await call;

                const idx = this.entries.findIndex((c) => c.id === entryId);
                if (idx > -1) {
                    this.entries.splice(idx, 1);
                }
            } catch (e) {
                $grpc.handleError(e as RpcError);
                throw e;
            }
        },

        // RSVP
        async listCalendarEntryRSVP(req: ListCalendarEntryRSVPRequest): Promise<ListCalendarEntryRSVPResponse> {
            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getCalendarClient().listCalendarEntryRSVP(req);
                const { response } = await call;

                return response;
            } catch (e) {
                $grpc.handleError(e as RpcError);
                throw e;
            }
        },
        async rsvpCalendarEntry(req: RSVPCalendarEntryRequest): Promise<RSVPCalendarEntryResponse> {
            const { $grpc } = useNuxtApp();
            try {
                const call = $grpc.getCalendarClient().rSVPCalendarEntry(req);
                const { response } = await call;

                // Retrieve calendar entry if a "should be visible" response and it is not in our list yet
                if (req.entry?.entryId && response.entry?.response && response.entry.response > RsvpResponses.HIDDEN) {
                    const entry = this.entries.find((e) => e.id === response.entry?.entryId);
                    if (!entry) {
                        await this.getCalendarEntry({ entryId: req.entry?.entryId });
                    } else {
                        entry.rsvp = response.entry;
                    }
                }

                return response;
            } catch (e) {
                $grpc.handleError(e as RpcError);
                throw e;
            }
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useCalendarStore as unknown as StoreDefinition, import.meta.hot));
}
