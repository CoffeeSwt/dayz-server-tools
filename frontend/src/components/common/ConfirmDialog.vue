<template>
    <Teleport to="body">
        <div v-if="modelValue" fixed inset-0 z-1000 flex-center class="bg-black/45" @click.self="handleCancel">
            <div w-96 max-w-90vw rounded-xl bg-light-white-2 dark:bg-dark-black-2 p-6 box-border shadow-xl>
                <div text-lg font-bold text-light-black-1 dark:text-white>{{ title }}</div>
                <div mt-3 text-sm leading-6 text-light-white-1 dark:text-dark-gray-1>{{ message }}</div>
                <div mt-6 flex justify-end gap-3 bg-light-white-2 text-light-black-1 dark:bg-dark-black-2
                    dark:text-dark-gray-1>
                    <div @click="handleCancel" flex-center cursor-pointer hover:bg-light-black-2
                        dark:hover:bg-dark-black-1 rounded-2xl p-4 py-2 box-border duration-150>
                        {{ cancelText }}
                    </div>
                    <div @click="handleConfirm" flex-center cursor-pointer hover:bg-light-black-2
                        dark:hover:bg-dark-black-1 rounded-2xl p-4 py-2 box-border duration-150>
                        {{ confirmText }}
                    </div>
                </div>
            </div>
        </div>
    </Teleport>
</template>

<script setup lang="ts">
const props = withDefaults(defineProps<{
    modelValue: boolean
    title?: string
    message?: string
    confirmText?: string
    cancelText?: string
}>(), {
    title: '确认操作',
    message: '确认继续吗？',
    confirmText: '确认',
    cancelText: '取消',
})

const emit = defineEmits<{
    (e: 'update:modelValue', value: boolean): void
    (e: 'confirm'): void
    (e: 'cancel'): void
}>()

const handleConfirm = () => {
    emit('confirm')
    emit('update:modelValue', false)
}

const handleCancel = () => {
    emit('cancel')
    emit('update:modelValue', false)
}
</script>
