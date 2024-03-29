import { RpcError } from '@protobuf-ts/runtime-rpc';
import { defineStore, type StoreDefinition } from 'pinia';
import { v4 as uuidv4 } from 'uuid';
import { type Notification, type NotificationConfig, type NotificationType } from '~/composables/notifications/notifications';
import { useAuthStore } from '~/store/auth';
import { NotificationCategory } from '~~/gen/ts/resources/notifications/notifications';
import { MarkNotificationsRequest } from '~~/gen/ts/services/notificator/notificator';

// In seconds
const initialReconnectBackoffTime = 2;

export interface NotificationsState {
    lastId: string;
    notifications: Notification[];
    abort: AbortController | undefined;
    reconnecting: boolean;
    reconnectBackoffTime: number;
}

export const useNotificatorStore = defineStore('notifications', {
    state: () =>
        ({
            lastId: '0',
            notifications: [],
            abort: undefined,
            reconnecting: false,
            reconnectBackoffTime: 0,
        }) as NotificationsState,
    persist: {
        paths: ['lastId'],
    },
    actions: {
        setLastId(lastId: bigint): void {
            this.lastId = lastId.toString();
        },
        removeNotification(id: string): void {
            this.notifications = this.notifications.filter((notification) => notification.id !== id);
        },
        dispatchNotification({
            title,
            content,
            type,
            autoClose = true,
            duration = 4500,
            category = NotificationCategory.GENERAL,
            data = undefined,
            position = 'top-right',
            callback = undefined,
            onClick = undefined,
            onClickText = undefined,
        }: NotificationConfig) {
            const id = uuidv4();
            this.notifications.unshift({
                id,
                title,
                content,
                type,
                category,
                data,
                position,
                callback,
                onClick,
                onClickText,
            });

            if (autoClose) {
                setTimeout(() => {
                    this.removeNotification(id);
                }, duration);
            }
        },

        async startStream(): Promise<void> {
            if (this.abort !== undefined) {
                return;
            }

            console.debug('Notificator: Stream starting, starting at ID:', this.getLastId);

            this.abort = new AbortController();
            this.reconnecting = false;
            const { $grpc } = useNuxtApp();

            try {
                this.abort = new AbortController();

                const call = $grpc.getNotificatorClient().stream(
                    {
                        lastId: this.getLastId,
                    },
                    {
                        abort: this.abort.signal,
                    },
                );

                for await (const resp of call.responses) {
                    if (resp.lastId > this.getLastId) {
                        this.setLastId(resp.lastId);
                    }

                    if (resp.data.oneofKind !== undefined) {
                        if (resp.data.oneofKind === 'token') {
                            const tokenUpdate = resp.data.token;

                            const authStore = useAuthStore();

                            // Update active char when updated user info is received
                            if (tokenUpdate.userInfo) {
                                console.debug('Notificator: Updated UserInfo received');

                                authStore.setActiveChar(tokenUpdate.userInfo);
                                authStore.setPermissions(tokenUpdate.permissions);
                                authStore.setJobProps(tokenUpdate.jobProps);
                            }

                            if (tokenUpdate.newToken && tokenUpdate.expires) {
                                console.debug('Notificator: New Token received');

                                authStore.setAccessToken(tokenUpdate.newToken, toDate(tokenUpdate.expires) as null | Date);

                                this.dispatchNotification({
                                    title: { key: 'notifications.renewed_token.title', parameters: {} },
                                    content: { key: 'notifications.renewed_token.content', parameters: {} },
                                    type: 'info',
                                });
                            }
                        } else if (resp.data.oneofKind === 'notifications') {
                            resp.data.notifications.notifications.forEach((n) => {
                                const nType: NotificationType = (n.type as NotificationType) ?? 'info';

                                if (n.title === undefined || n.content === undefined) {
                                    return;
                                }

                                switch (n.category) {
                                    default: {
                                        const not: NotificationConfig = {
                                            title: { key: n.title.key, parameters: n.title.parameters },
                                            content: {
                                                key: n.content.key,
                                                parameters: n.content.parameters,
                                            },
                                            type: nType,
                                            category: n.category,
                                            data: n.data,
                                        };

                                        if (n.data?.link !== undefined) {
                                            if (n.data.link.external === true) {
                                                not.onClick = async () => {
                                                    navigateTo(n.data!.link!.to, { external: true });
                                                };
                                            } else {
                                                // @ts-ignore route from a notification is a string
                                                const route = useRouter().resolve(n.data!.link!.to);

                                                not.onClick = async () => {
                                                    navigateTo(route);
                                                };
                                            }
                                        }

                                        this.dispatchNotification(not);
                                        break;
                                    }
                                }
                            });
                        } else {
                            // @ts-ignore this is a catch all "unknown", so okay if it is technically "never" reached till it is
                            console.warn('Notificator: Unknown data received - Kind: ', resp.data.oneofKind, resp.data);
                        }
                    }

                    if (resp.restart) {
                        console.debug('Notificator: Server requested stream to be restarted');
                        this.reconnectBackoffTime = 0;
                        this.restartStream();
                        break;
                    }
                }
            } catch (e) {
                const error = e as RpcError;
                if (error.code !== 'CANCELLED' && error.code !== 'ABORTED') {
                    console.debug('Notificator: Stream failed', error.code, error.message, error.cause);
                    this.restartStream();
                }
            }

            console.debug('Notificator: Stream ended');
        },

        async stopStream(): Promise<void> {
            if (this.abort === undefined) {
                return;
            }

            this.abort?.abort();
            this.abort = undefined;
            console.debug('Notificator: Stopping Data Stream');
        },

        async restartStream(): Promise<void> {
            this.reconnecting = true;

            // Reset back off time when over 2 minutes
            if (this.reconnectBackoffTime > 120) {
                this.reconnectBackoffTime = initialReconnectBackoffTime;
            } else {
                this.reconnectBackoffTime += initialReconnectBackoffTime;
            }

            console.debug('Notificator: Restart back off time in', this.reconnectBackoffTime, 'seconds');
            await this.stopStream();

            setTimeout(async () => {
                if (this.reconnecting) {
                    this.startStream();
                }
            }, this.reconnectBackoffTime * 1000);
        },

        // Notification Actions
        async markNotifications(req: MarkNotificationsRequest): Promise<void> {
            const { $grpc } = useNuxtApp();

            try {
                await $grpc.getNotificatorClient().markNotifications(req);
            } catch (e) {
                $grpc.handleError(e as RpcError);
                throw e;
            }
        },
    },
    getters: {
        getLastId(): bigint {
            return BigInt(this.lastId);
        },
        getNotifications(): Notification[] {
            return [...this.notifications].slice(0, 4);
        },
        getNotificationsCount(): number {
            return this.notifications.length;
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useNotificatorStore as unknown as StoreDefinition, import.meta.hot));
}
