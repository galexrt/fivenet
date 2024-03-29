<script lang="ts" setup>
import {
    Disclosure,
    DisclosureButton,
    DisclosurePanel,
    Listbox,
    ListboxButton,
    ListboxOption,
    ListboxOptions,
} from '@headlessui/vue';
import { CheckIcon, ChevronDownIcon } from 'mdi-vue3';
import { useCompletorStore } from '~/store/completor';
import { AttributeValues, Permission, RoleAttribute } from '~~/gen/ts/resources/permissions/permissions';
import { Job, JobGrade } from '~~/gen/ts/resources/users/jobs';

const props = defineProps<{
    attribute: RoleAttribute;
    states: Map<string, AttributeValues | undefined>;
    disabled?: boolean;
    permission: Permission;
}>();

const emit = defineEmits<{
    (e: 'update:states', payload: Map<string, AttributeValues | undefined>): void;
    (e: 'changed'): void;
}>();

const completorStore = useCompletorStore();
const { jobs } = storeToRefs(completorStore);
const { listJobs } = completorStore;

const jobGrades = ref<Map<string, JobGrade>>(new Map());

const states = ref<typeof props.states>(props.states);
const id = ref<string>(props.attribute.attrId);

let maxValues = props.attribute.maxValues;
if (maxValues === undefined) {
    switch (lowercaseFirstLetter(props.attribute.type)) {
        case 'stringList':
            maxValues = {
                validValues: {
                    oneofKind: 'stringList',
                    stringList: {
                        strings: [],
                    },
                },
            };
            break;

        case 'jobList':
            maxValues = {
                validValues: {
                    oneofKind: 'jobList',
                    jobList: {
                        strings: [],
                    },
                },
            };
            break;

        case 'jobGradeList':
            maxValues = {
                validValues: {
                    oneofKind: 'jobGradeList',
                    jobGradeList: {
                        jobs: {},
                    },
                },
            };
            break;
    }
}

if (!states.value.has(id.value) || states.value.get(id.value) === undefined) {
    switch (lowercaseFirstLetter(props.attribute.type)) {
        case 'stringList': {
            states.value.set(id.value, {
                validValues: {
                    oneofKind: 'stringList',
                    stringList: {
                        strings: [],
                    },
                },
            });
            break;
        }

        case 'jobList': {
            states.value.set(id.value, {
                validValues: {
                    oneofKind: 'jobList',
                    jobList: {
                        strings: [],
                    },
                },
            });
            break;
        }

        case 'jobGradeList': {
            states.value.set(id.value, {
                validValues: {
                    oneofKind: 'jobGradeList',
                    jobGradeList: {
                        jobs: {},
                    },
                },
            });
            break;
        }
    }
}

const currentValue = states.value.get(id.value)!;
const validValues = ref<AttributeValues | undefined>(props.attribute.validValues);

async function toggleStringListValue(value: string): Promise<void> {
    if (currentValue.validValues.oneofKind !== 'stringList') {
        return;
    }

    const array = currentValue.validValues.stringList.strings;
    if (!array.includes(value)) {
        array.push(value);
    } else {
        array.splice(array.indexOf(value), 1);
    }

    currentValue.validValues.stringList.strings = array;
    states.value.set(id.value, currentValue);
    emit('update:states', states.value);
    emit('changed');
}

async function toggleJobListValue(value: string): Promise<void> {
    if (currentValue.validValues.oneofKind !== 'jobList') {
        return;
    }

    const array = currentValue.validValues.jobList.strings;
    if (!array.includes(value)) {
        array.push(value);
    } else {
        array.splice(array.indexOf(value), 1);
    }

    currentValue.validValues.jobList.strings = array;
    states.value.set(id.value, currentValue);
    emit('update:states', states.value);
    emit('changed');
}

async function toggleJobGradeValue(job: Job, checked: boolean): Promise<void> {
    if (currentValue.validValues.oneofKind !== 'jobGradeList') {
        return;
    }

    const map = currentValue.validValues.jobGradeList.jobs;
    if (checked && !map[job.name]) {
        map[job.name] = 1;
        jobGrades.value.set(job.name, job.grades[0]);
    } else if (!checked && map[job.name]) {
        delete map[job.name];
        jobGrades.value.set(job.name, job.grades[0]);
    } else {
        return;
    }

    currentValue.validValues.jobGradeList.jobs = map;
    states.value.set(id.value, currentValue);
    emit('update:states', states.value);
    emit('changed');
}

async function updateJobGradeValue(job: Job, grade: JobGrade): Promise<void> {
    if (currentValue.validValues.oneofKind !== 'jobGradeList') {
        return;
    }

    const map = currentValue.validValues.jobGradeList.jobs;

    map[job.name] = grade.grade;
    jobGrades.value.set(job.name, job.grades[grade.grade - 1]);

    currentValue.validValues.jobGradeList.jobs = map;
    states.value.set(id.value, currentValue);
    emit('update:states', states.value);
    emit('changed');
}

onBeforeMount(async () => {
    if (currentValue.validValues.oneofKind === 'jobList' || currentValue.validValues.oneofKind === 'jobGradeList') {
        await listJobs();

        jobs.value.forEach((job) => {
            if (currentValue.validValues.oneofKind !== 'jobGradeList') {
                return;
            }

            if (maxValues && maxValues.validValues.oneofKind === 'jobGradeList') {
                if (!maxValues.validValues.jobGradeList.jobs[job.name]) {
                    return;
                }
            }
            jobGrades.value.set(job.name, job.grades[(currentValue.validValues?.jobGradeList.jobs[job.name] ?? 1) - 1]);
        });
    }
});
</script>

<template>
    <div v-if="attribute">
        <Disclosure
            v-slot="{ open }"
            as="div"
            :class="[disabled ? 'border-neutral/10 text-base-300' : 'border-neutral/20 text-neutral hover:border-neutral/70']"
        >
            <DisclosureButton
                :disabled="disabled"
                :class="[
                    open ? 'rounded-t-lg border-b-0' : 'rounded-lg',
                    disabled ? 'cursor-not-allowed' : '',
                    'flex w-full items-start justify-between border-2 border-inherit p-2 text-left transition-colors',
                ]"
            >
                <span class="text-base leading-7 transition-colors">
                    {{ $t(`perms.${attribute.category}.${attribute.name}.attrs_types.${attribute.key}`) }}
                </span>
                <span class="ml-6 flex h-7 items-center">
                    <ChevronDownIcon :class="[open ? 'upsidedown' : '', 'h-5 w-5 transition-transform']" aria-hidden="true" />
                </span>
            </DisclosureButton>
            <DisclosurePanel class="-mt-2 rounded-b-lg border-2 border-t-0 border-inherit px-4 pb-2 transition-colors">
                <div class="mx-auto my-2 flex max-w-4xl flex-col gap-2">
                    <div
                        v-if="
                            currentValue.validValues.oneofKind === 'stringList' &&
                            maxValues?.validValues &&
                            maxValues?.validValues.oneofKind === 'stringList'
                        "
                        class="flex flex-row flex-wrap gap-4"
                    >
                        <span v-if="maxValues.validValues.stringList.strings.length === 0">
                            {{ $t('common.not_found', [$t('common.attributes', 2)]) }}
                        </span>
                        <template v-else>
                            <div
                                v-for="value in maxValues.validValues.stringList.strings"
                                :key="value"
                                class="flex flex-initial flex-row flex-nowrap"
                            >
                                <input
                                    :id="value"
                                    :name="value"
                                    type="checkbox"
                                    :checked="!!currentValue.validValues.stringList.strings.find((v) => v === value)"
                                    class="my-auto h-4 w-4 rounded border-base-300 text-primary-500 focus:ring-primary-500"
                                    @click="toggleStringListValue(value)"
                                />
                                <span class="ml-1">{{
                                    $t(`perms.${permission.category}.${permission.name}.attrs.${value.replaceAll('.', '_')}`)
                                }}</span>
                            </div>
                        </template>
                    </div>
                    <div
                        v-else-if="
                            currentValue.validValues.oneofKind === 'jobList' &&
                            maxValues?.validValues &&
                            maxValues?.validValues.oneofKind === 'jobList'
                        "
                        class="flex flex-row flex-wrap gap-4"
                    >
                        <span v-if="maxValues.validValues.jobList.strings.length === 0">
                            {{ $t('common.not_found', [$t('common.attributes', 2)]) }}
                        </span>
                        <template v-else>
                            <div
                                v-for="job in jobs.filter(
                                    (j) =>
                                        maxValues?.validValues.oneofKind === 'jobList' &&
                                        (!maxValues?.validValues.jobList?.strings.length ||
                                            maxValues.validValues?.jobList?.strings.includes(j.name)),
                                )"
                                :key="job.name"
                                class="flex flex-initial flex-row flex-nowrap"
                            >
                                <input
                                    :id="job.name"
                                    :name="job.name"
                                    type="checkbox"
                                    :checked="!!currentValue.validValues.jobList?.strings.find((v) => v === job.name)"
                                    class="my-auto h-4 w-4 rounded border-base-300 text-primary-500 focus:ring-primary-500"
                                    @click="toggleJobListValue(job.name)"
                                />
                                <span class="ml-1">{{ job.label }}</span>
                            </div>
                        </template>
                    </div>
                    <div
                        v-else-if="
                            currentValue.validValues.oneofKind === 'jobGradeList' &&
                            maxValues?.validValues &&
                            maxValues.validValues.oneofKind === 'jobGradeList'
                        "
                        class="flex flex-col gap-2"
                    >
                        <span v-if="Object.keys(maxValues.validValues.jobGradeList.jobs).length === 0">
                            {{ $t('common.not_found', [$t('common.attributes', 2)]) }}
                        </span>
                        <template v-else>
                            <div
                                v-for="job in jobs.filter(
                                    (j) =>
                                        maxValues &&
                                        maxValues.validValues.oneofKind === 'jobGradeList' &&
                                        maxValues.validValues.jobGradeList.jobs[j.name],
                                )"
                                :key="job.name"
                                class="flex flex-initial flex-row flex-nowrap gap-2"
                            >
                                <input
                                    :id="job.name"
                                    :name="job.name"
                                    type="checkbox"
                                    :checked="!!currentValue.validValues?.jobGradeList.jobs[job.name]"
                                    class="my-auto h-4 w-4 rounded border-base-300 text-primary-500 focus:ring-primary-500"
                                    @change="toggleJobGradeValue(job, ($event.target as any).checked)"
                                />
                                <span class="my-auto flex-1">{{ job.label }}</span>
                                <Listbox
                                    as="div"
                                    class="flex-1"
                                    :model-value="jobGrades.get(job.name)"
                                    :disabled="!currentValue.validValues.jobGradeList?.jobs[job.name]"
                                    @update:model-value="updateJobGradeValue(job, $event)"
                                >
                                    <div class="relative">
                                        <ListboxButton
                                            class="block w-full rounded-md border-0 bg-base-700 py-1.5 pl-3 text-left text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 disabled:cursor-not-allowed disabled:bg-base-800 disabled:text-neutral/50 sm:text-sm sm:leading-6"
                                        >
                                            <span class="block truncate">{{ jobGrades.get(job.name)?.label }}</span>
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
                                                    v-for="grade in job.grades.filter(
                                                        (g) =>
                                                            maxValues &&
                                                            maxValues.validValues.oneofKind === 'jobGradeList' &&
                                                            maxValues.validValues.jobGradeList.jobs[job.name] + 1 > g.grade,
                                                    )"
                                                    :key="grade.grade"
                                                    v-slot="{ active, selected }"
                                                    as="template"
                                                    :value="grade"
                                                >
                                                    <li
                                                        :class="[
                                                            active ? 'bg-primary-500' : '',
                                                            'relative cursor-default select-none py-2 pl-8 pr-4 text-neutral',
                                                        ]"
                                                    >
                                                        <span
                                                            :class="[
                                                                selected ? 'font-semibold' : 'font-normal',
                                                                'block truncate',
                                                            ]"
                                                        >
                                                            {{ grade.label }}
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
                                                </ListboxOption>
                                            </ListboxOptions>
                                        </transition>
                                    </div>
                                </Listbox>
                            </div>
                        </template>
                    </div>
                    <div v-else>{{ currentValue.validValues.oneofKind }} {{ validValues }}</div>
                </div>
            </DisclosurePanel>
        </Disclosure>
    </div>
</template>
