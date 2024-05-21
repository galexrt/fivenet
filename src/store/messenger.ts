import Dexie, { type Table } from 'dexie';
import type { Message } from '~~/gen/ts/resources/messenger/message';
import type { Thread, ThreadUserState } from '~~/gen/ts/resources/messenger/thread';
import type {
    CreateOrUpdateThreadRequest,
    CreateOrUpdateThreadResponse,
    DeleteThreadRequest,
    DeleteThreadResponse,
    PostMessageRequest,
    PostMessageResponse,
} from '~~/gen/ts/services/messenger/messenger';

export class MessengerDexie extends Dexie {
    threads!: Table<Thread>;
    messages!: Table<Message>;

    constructor() {
        super('messenger');
        this.version(1).stores({
            threads: 'id',
            messages: 'id, threadId',
        });
    }

    // Thread
    async createOrUpdateThread(req: CreateOrUpdateThreadRequest): Promise<CreateOrUpdateThreadResponse> {
        try {
            const call = getGRPCMessengerClient().createOrUpdateThread(req);
            const { response } = await call;

            if (response.thread) {
                await this.threads.put(response.thread);
            }

            return response;
        } catch (e) {
            handleGRPCError(e);
            throw e;
        }
    }

    async deleteThread(req: DeleteThreadRequest): Promise<DeleteThreadResponse> {
        try {
            const call = getGRPCMessengerClient().deleteThread(req);
            const { response } = await call;

            await this.threads.delete(req.threadId);

            return response;
        } catch (e) {
            handleGRPCError(e);
            throw e;
        }
    }

    // Thread User State
    async getThreadUserState(threadId: string): Promise<ThreadUserState | undefined> {
        return (await this.threads.get(threadId))?.userState;
    }

    async setThreadUserState(state: Partial<ThreadUserState>): Promise<ThreadUserState | undefined> {
        const thread = await this.threads.get(state!.threadId);
        if (!thread) {
            return;
        }

        let update = false;
        if (!thread.userState) {
            update = true;

            thread.userState = {
                threadId: state.threadId!,
                userId: 0,
                lastRead: toTimestamp(),
                unread: false,
                important: false,
                favorite: false,
                muted: false,
            };
        } else {
            if (state.lastRead !== undefined && thread.userState.lastRead?.timestamp !== state.lastRead.timestamp) {
                update = true;
                thread.userState.lastRead = state.lastRead;
            }
            if (state.unread !== undefined && thread.userState.unread !== state.unread) {
                update = true;
                thread.userState.unread = state.unread;
            }
            if (state.important !== undefined && thread.userState.important !== state.important) {
                update = true;
                thread.userState.important = state.important;
            }
            if (state.favorite !== undefined && thread.userState.favorite !== state.favorite) {
                update = true;
                thread.userState.favorite = state.favorite;
            }
            if (state.muted !== undefined && thread.userState.muted !== state.muted) {
                update = true;
                thread.userState.muted = state.muted;
            }
        }

        if (update) {
            this.threads.update(thread.id, thread);
            await getGRPCMessengerClient().setThreadUserState({ state: thread.userState });
        }

        return thread.userState;
    }

    // Messages
    async postMessage(req: PostMessageRequest): Promise<PostMessageResponse> {
        try {
            const call = getGRPCMessengerClient().postMessage(req);
            const { response } = await call;

            if (response.message) {
                this.messages.add(response.message);
            }

            return response;
        } catch (e) {
            handleGRPCError(e);
            throw e;
        }
    }
}

export const messengerStore = new MessengerDexie();
