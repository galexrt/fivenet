<script lang="ts" setup>
import {
    Combobox,
    ComboboxButton,
    ComboboxInput,
    ComboboxOption,
    ComboboxOptions,
    Listbox,
    ListboxButton,
    ListboxOption,
    ListboxOptions,
} from '@headlessui/vue';
import { listEnumValues } from '@protobuf-ts/runtime';
import { watchDebounced } from '@vueuse/core';
import { CheckIcon, ChevronDownIcon, CloseIcon } from 'mdi-vue3';
import { useCompletorStore } from '~/store/completor';
import { type ArrayElement } from '~/utils/types';
import { AccessLevel } from '~~/gen/ts/resources/documents/access';
import { Job, JobGrade } from '~~/gen/ts/resources/users/jobs';
import { UserShort } from '~~/gen/ts/resources/users/users';

type AccessType = { id: number; name: string };

const props = withDefaults(
    defineProps<{
        readOnly?: boolean;
        showRequired?: boolean;
        init: {
            id: string;
            type: number;
            values: {
                job?: string;
                char?: number;
                accessRole?: AccessLevel;
                minimumGrade?: number;
            };
            required?: boolean;
        };
        accessTypes: AccessType[];
        accessRoles?: undefined | AccessLevel[];
    }>(),
    {
        readOnly: false,
        showRequired: false,
        accessRoles: undefined,
    },
);

const emit = defineEmits<{
    (e: 'typeChange', payload: { id: string; type: number }): void;
    (
        e: 'nameChange',
        payload: {
            id: string;
            job: Job | undefined;
            char: UserShort | undefined;
            required?: boolean;
        },
    ): void;
    (e: 'rankChange', payload: { id: string; rank: JobGrade; required?: boolean }): void;
    (e: 'accessChange', payload: { id: string; access: AccessLevel; required?: boolean }): void;
    (e: 'deleteRequest', payload: { id: string }): void;
    (e: 'requiredChange', payload: { id: string; required?: boolean }): void;
}>();

const completorStore = useCompletorStore();
const { jobs } = storeToRefs(completorStore);

const { t } = useI18n();

const required = ref<boolean | undefined>(props.init.required);
const selectedAccessType = ref<AccessType>({
    id: -1,
    name: '',
});
const entriesChars = ref<UserShort[]>();
const queryCharRaw = ref('');
const queryChar = computed(() => queryCharRaw.value.toLowerCase());
const selectedChar = ref<undefined | UserShort>(undefined);

const queryJobRaw = ref('');
const queryJob = computed(() => queryJobRaw.value.toLowerCase());
const filteredJobs = computed(() =>
    jobs.value.filter((j) => j.name.toLowerCase().includes(queryJob.value) || j.label.toLowerCase().includes(queryJob.value)),
);
const selectedJob = ref<Job>();

const queryMinimumRankRaw = ref('');
const queryMinimumRank = computed(() => queryMinimumRankRaw.value.toLowerCase());
const entriesMinimumRank = ref<JobGrade[]>([]);
const filteredJobRanks = computed(() =>
    entriesMinimumRank.value.filter(
        (j) =>
            j.grade.toString().toLowerCase().includes(queryMinimumRank.value) ||
            j.label.toLowerCase().includes(queryMinimumRank.value),
    ),
);
const selectedMinimumRank = ref<JobGrade | undefined>(undefined);

const entriesAccessRoles: {
    id: AccessLevel;
    label: string;
    value: string;
}[] = [];
if (props.accessRoles === undefined || props.accessRoles.length === 0) {
    entriesAccessRoles.push(
        ...listEnumValues(AccessLevel)
            .map((e, k) => {
                return {
                    id: k,
                    label: t(`enums.docstore.AccessLevel.${e.name}`),
                    value: e.name,
                };
            })
            .filter((e) => e.id !== 0),
    );
} else {
    props.accessRoles.forEach((e) => {
        entriesAccessRoles.push({
            id: e,
            label: t(`enums.docstore.AccessLevel.${AccessLevel[e]}`),
            value: AccessLevel[e],
        });
    });
}

const queryAccessRole = ref('');
const selectedAccessRole = ref<ArrayElement<typeof entriesAccessRoles>>();

async function findChars(userId?: number): Promise<UserShort[]> {
    if (queryChar.value === '' && userId === undefined) return [];

    return completorStore.completeCitizens({
        search: queryChar.value,
        userId,
    });
}

onMounted(async () => {
    const passedType = props.accessTypes.find((e) => e.id === props.init.type);
    if (passedType) {
        selectedAccessType.value = passedType;
    }

    if (props.init.type === 0 && props.init.values.char !== undefined && props.init.values.accessRole !== undefined) {
        const users = await findChars(props.init.values.char);
        selectedChar.value = users.find((char) => char.userId === props.init.values.char);
    } else if (
        props.init.type === 1 &&
        props.init.values.job !== undefined &&
        props.init.values.minimumGrade !== undefined &&
        props.init.values.accessRole !== undefined
    ) {
        selectedJob.value = await completorStore.getJobByName(props.init.values.job);
        if (selectedJob.value) {
            entriesMinimumRank.value = selectedJob.value.grades;
        }
        selectedMinimumRank.value = entriesMinimumRank.value.find((rank) => rank.grade === props.init.values.minimumGrade);
    }
    selectedAccessRole.value = entriesAccessRoles.find((type) => type.id === props.init.values.accessRole);
});

watchDebounced(queryChar, async () => (entriesChars.value = await findChars()), {
    debounce: 600,
    maxWait: 1750,
});

watch(required, () => {
    emit('requiredChange', { id: props.init.id, required: required.value });
});

watch(selectedAccessType, async () => {
    emit('typeChange', {
        id: props.init.id,
        type: selectedAccessType.value.id,
    });

    selectedChar.value = undefined;
    selectedJob.value = undefined;
    selectedMinimumRank.value = undefined;

    if (selectedAccessType.value.id === 0) {
        queryCharRaw.value = '';
    } else {
        queryJobRaw.value = '';
        queryMinimumRankRaw.value = '';
    }
});

watch(selectedJob, () => {
    if (!selectedJob.value) {
        return;
    }

    emit('nameChange', {
        id: props.init.id,
        job: selectedJob.value,
        char: undefined,
        required: required.value,
    });

    entriesMinimumRank.value = selectedJob.value.grades;
});

watch(selectedChar, () => {
    if (!selectedChar.value) {
        return;
    }

    emit('nameChange', {
        id: props.init.id,
        job: undefined,
        char: selectedChar.value,
        required: required.value,
    });
});

watch(selectedMinimumRank, () => {
    if (!selectedMinimumRank.value) {
        return;
    }

    emit('rankChange', { id: props.init.id, rank: selectedMinimumRank.value });
});

watch(selectedAccessRole, () => {
    if (!selectedAccessRole.value) {
        return;
    }

    emit('accessChange', {
        id: props.init.id,
        access: selectedAccessRole.value.id,
    });
});
</script>

<template>
    <div class="my-2 flex flex-row items-center">
        <div v-if="showRequired" class="mr-2 flex-initial">
            <input
                v-model="required"
                :disabled="readOnly"
                :title="$t('common.require')"
                type="checkbox"
                name="required"
                data-te-toggle="tooltip"
                class="h-8 w-8 rounded-md border-0 bg-base-700 text-primary-600 transition-colors hover:bg-base-600 hover:text-primary-500 focus:ring-2 focus:ring-inset focus:ring-base-300"
                :class="readOnly ? 'disabled' : ''"
            />
        </div>
        <div class="mr-2 w-60 flex-initial">
            <input
                v-if="accessTypes.length === 1"
                type="text"
                disabled
                :value="accessTypes[0].name"
                class="block w-full rounded-md border-0 bg-base-700 py-1.5 pl-3 text-left text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                @focusin="focusTablet(true)"
                @focusout="focusTablet(false)"
            />
            <Listbox v-else v-model="selectedAccessType" as="div" :disabled="readOnly">
                <div class="relative">
                    <ListboxButton
                        class="block w-full rounded-md border-0 bg-base-700 py-1.5 pl-3 text-left text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                        :class="readOnly ? 'disabled' : ''"
                    >
                        <span class="block truncate">{{ selectedAccessType?.name }}</span>
                        <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                            <ChevronDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
                        </span>
                    </ListboxButton>

                    <transition
                        leave-active-class="transition duration-100 ease-in"
                        leave-from-class="opacity-100"
                        leave-to-class="opacity-0"
                    >
                        <ListboxOptions
                            class="absolute z-10 mt-1 max-h-44 w-full overflow-auto rounded-md bg-base-700 py-1 text-base sm:text-sm"
                        >
                            <ListboxOption
                                v-for="accessType in accessTypes"
                                :key="accessType.id"
                                v-slot="{ active, selected }"
                                as="template"
                                :value="accessType"
                            >
                                <li
                                    :class="[
                                        active ? 'bg-primary-500' : '',
                                        'relative cursor-default select-none py-2 pl-8 pr-4 text-neutral',
                                    ]"
                                >
                                    <span :class="[selected ? 'font-semibold' : 'font-normal', 'block truncate']">{{
                                        accessType.name
                                    }}</span>

                                    <span
                                        v-if="selected"
                                        :class="[
                                            active ? 'text-neutral' : 'text-primary-500',
                                            'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                        ]"
                                    >
                                        <CheckIcon class="h-5 w-5" aria-hidden="true" />
                                    </span>
                                </li>
                            </ListboxOption>
                        </ListboxOptions>
                    </transition>
                </div>
            </Listbox>
        </div>
        <div v-if="selectedAccessType?.id === 0" class="flex flex-grow">
            <div class="mr-2 flex-1">
                <Combobox v-model="selectedChar" as="div" :disabled="readOnly">
                    <div class="relative">
                        <ComboboxButton as="div">
                            <ComboboxInput
                                autocomplete="off"
                                class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                :display-value="(char: any) => `${char?.firstname} ${char?.lastname} (${char?.dateofbirth})`"
                                :class="readOnly ? 'disabled' : ''"
                                @change="queryCharRaw = $event.target.value"
                                @focusin="focusTablet(true)"
                                @focusout="focusTablet(false)"
                            />
                        </ComboboxButton>

                        <ComboboxOptions
                            v-if="entriesChars && entriesChars.length > 0"
                            class="absolute z-10 mt-1 max-h-44 w-full overflow-auto rounded-md bg-base-700 py-1 text-base sm:text-sm"
                        >
                            <ComboboxOption
                                v-for="char in entriesChars"
                                :key="char.identifier"
                                v-slot="{ active, selected }"
                                :value="char"
                                as="char"
                            >
                                <li
                                    :class="[
                                        'relative cursor-default select-none py-2 pl-8 pr-4 text-neutral',
                                        active ? 'bg-primary-500' : '',
                                    ]"
                                >
                                    <span :class="['block truncate', selected && 'font-semibold']">
                                        {{ char.firstname }} {{ char.lastname }} ({{ char.dateofbirth }})
                                    </span>

                                    <span
                                        v-if="selected"
                                        :class="[
                                            active ? 'text-neutral' : 'text-primary-500',
                                            'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                        ]"
                                    >
                                        <CheckIcon class="h-5 w-5" aria-hidden="true" />
                                    </span>
                                </li>
                            </ComboboxOption>
                        </ComboboxOptions>
                    </div>
                </Combobox>
            </div>
        </div>
        <div v-else class="flex flex-grow">
            <div class="mr-2 flex-1">
                <Combobox v-model="selectedJob" as="div" :disabled="readOnly">
                    <div class="relative">
                        <ComboboxButton as="div">
                            <ComboboxInput
                                autocomplete="off"
                                class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                :display-value="(job: any) => job?.label"
                                :class="readOnly ? 'disabled' : ''"
                                @change="queryJobRaw = $event.target.value"
                                @focusin="focusTablet(true)"
                                @focusout="focusTablet(false)"
                            />
                        </ComboboxButton>

                        <ComboboxOptions
                            v-if="filteredJobs.length > 0"
                            class="absolute z-10 mt-1 max-h-44 w-full overflow-auto rounded-md bg-base-700 py-1 text-base sm:text-sm"
                        >
                            <ComboboxOption
                                v-for="job in filteredJobs"
                                :key="job.name"
                                v-slot="{ active, selected }"
                                :value="job"
                            >
                                <li
                                    :class="[
                                        'relative cursor-default select-none py-2 pl-8 pr-4 text-neutral',
                                        active ? 'bg-primary-500' : '',
                                    ]"
                                >
                                    <span :class="['block truncate', selected && 'font-semibold']">
                                        {{ job.label }}
                                    </span>

                                    <span
                                        v-if="selected"
                                        :class="[
                                            active ? 'text-neutral' : 'text-primary-500',
                                            'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                        ]"
                                    >
                                        <CheckIcon class="h-5 w-5" aria-hidden="true" />
                                    </span>
                                </li>
                            </ComboboxOption>
                        </ComboboxOptions>
                    </div>
                </Combobox>
            </div>
            <div class="mr-2 flex-1">
                <Combobox v-model="selectedMinimumRank" as="div" :disabled="readOnly || selectedJob === undefined">
                    <div class="relative">
                        <ComboboxButton as="div">
                            <ComboboxInput
                                autocomplete="off"
                                class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                :class="readOnly ? 'disabled' : ''"
                                :display-value="(rank: any) => rank?.label"
                                @change="queryMinimumRankRaw = $event.target.value"
                                @focusin="focusTablet(true)"
                                @focusout="focusTablet(false)"
                            />
                        </ComboboxButton>

                        <ComboboxOptions
                            v-if="filteredJobRanks.length > 0"
                            class="absolute z-10 mt-1 max-h-44 w-full overflow-auto rounded-md bg-base-700 py-1 text-base sm:text-sm"
                        >
                            <ComboboxOption
                                v-for="rank in filteredJobRanks"
                                :key="rank.grade"
                                v-slot="{ active, selected }"
                                :value="rank"
                                as="minimumGrade"
                            >
                                <li
                                    :class="[
                                        'relative cursor-default select-none py-2 pl-8 pr-4 text-neutral',
                                        active ? 'bg-primary-500' : '',
                                    ]"
                                >
                                    <span :class="['block truncate', selected && 'font-semibold']">
                                        {{ rank.label }}
                                    </span>

                                    <span
                                        v-if="selected"
                                        :class="[
                                            active ? 'text-neutral' : 'text-primary-500',
                                            'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                        ]"
                                    >
                                        <CheckIcon class="h-5 w-5" aria-hidden="true" />
                                    </span>
                                </li>
                            </ComboboxOption>
                        </ComboboxOptions>
                    </div>
                </Combobox>
            </div>
        </div>
        <div class="flex-inital mr-2 w-60">
            <Combobox v-model="selectedAccessRole" as="div" :disabled="readOnly">
                <div class="relative">
                    <ComboboxButton as="div">
                        <ComboboxInput
                            autocomplete="off"
                            class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                            :class="readOnly ? 'disabled' : ''"
                            :display-value="(role: any) => role.label"
                            @change="queryAccessRole = $event.target.value"
                            @focusin="focusTablet(true)"
                            @focusout="focusTablet(false)"
                        />
                    </ComboboxButton>

                    <ComboboxOptions
                        v-if="entriesAccessRoles.length > 0"
                        class="absolute z-10 mt-1 max-h-44 w-full overflow-auto rounded-md bg-base-700 py-1 text-base sm:text-sm"
                    >
                        <ComboboxOption
                            v-for="role in entriesAccessRoles"
                            :key="role.id"
                            v-slot="{ active, selected }"
                            :value="role"
                            as="accessRole"
                        >
                            <li
                                :class="[
                                    'relative cursor-default select-none py-2 pl-8 pr-4 text-neutral',
                                    active ? 'bg-primary-500' : '',
                                ]"
                            >
                                <span :class="['block truncate', selected && 'font-semibold']">
                                    {{ role.label }}
                                </span>

                                <span
                                    v-if="selected"
                                    :class="[
                                        active ? 'text-neutral' : 'text-primary-500',
                                        'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                    ]"
                                >
                                    <CheckIcon class="h-5 w-5" aria-hidden="true" />
                                </span>
                            </li>
                        </ComboboxOption>
                    </ComboboxOptions>
                </div>
            </Combobox>
        </div>
        <div v-if="!readOnly" class="flex-initial">
            <button
                type="button"
                class="rounded-full bg-primary-500 p-1.5 text-neutral hover:bg-primary-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-500"
                @click="$emit('deleteRequest', { id: props.init.id })"
            >
                <CloseIcon class="h-5 w-5" aria-hidden="true" />
            </button>
        </div>
    </div>
</template>
