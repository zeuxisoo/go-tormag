<template>
    <file-pond
        name="file"
        ref="filePondRef"
        class-name="upload-zone"
        label-idle="Drop files here..."
        allow-multiple="true"
        accepted-file-types="application/x-bittorrent, application/octet-stream"
        :server="filePondConfig.server"
        :files="filePondConfig.files"
        @processfile="props.onProcessFile" />
</template>

<script setup>
import { ref, reactive } from "vue";
import VueFilePond from "vue-filepond";
import FilePondPluginFileValidateType from "filepond-plugin-file-validate-type";

// Properties
const props = defineProps({
    baseUrl      : { type: String, required: true },
    entryUrl     : { type: String, required: true },
    fileList     : { type: Array, required: true },
    fileText     : { type: String, required: true },
    onProcessFile: { type: Function, required: true },
});

// Component
const FilePond = VueFilePond(
    FilePondPluginFileValidateType
);

const filePondRef = ref(null);

const filePondConfig = reactive({
    server: {
        url    : props.baseUrl,
        process: {
            url            : props.entryUrl,
            method         : 'POST',
            withCredentials: false,
            headers        : {},
            timeout        : 7000,
            onload         : (response) => JSON.parse(response),
            onerror        : null
        }
    },
    files : [],
});

// Expose
defineExpose({
    filePondRef
});
</script>
