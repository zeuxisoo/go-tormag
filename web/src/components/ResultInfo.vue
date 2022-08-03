<template>
    <div class="result-info">
        <div class="row">
            <div class="col-lg-9">
                <button type="button" :class="['btn', { 'btn-info': isTargetResultMode('list') }]" @click="setResultMode('list')">
                    <i class="fas fa-list-ul"></i> List
                </button>
                &nbsp;
                <button type="button" :class="['btn', { 'btn-info': isTargetResultMode('text') }]" @click="setResultMode('text')">
                    <i class="fas fa-align-justify"></i> Text Only
                </button>
            </div>
            <div class="col-lg-3 text-end d-none d-lg-block"> <!-- display on >= lg size only -->
                <button type="button" class="btn btn-success">
                    OK <span class="badge badge-light">{{ countState.ok }}</span>
                </button>
                &nbsp;
                <button type="button" class="btn btn-danger">
                    Error <span class="badge badge-light">{{ countState.error }}</span>
                </button>
            </div>
        </div>
        <hr />
    </div>
</template>

<script setup>
import { computed } from "vue";

// Properties
const props = defineProps({
    resultMode: { type: String, required: true },
    fileList  : { type: Array, required: true },
});

// Emit
const emit = defineEmits(['changeResultMode']);

// Computed
const countState = computed(() => {
    const okCount = props.fileList
        .filter(file => file.ok === true)
        .length;

    return {
        ok   : okCount,
        error: props.fileList.length - okCount,
    }
});

// Methods
const isTargetResultMode = name => props.resultMode === name;

const setResultMode = name => {
    emit('changeResultMode', name.toLowerCase());
};
</script>
