<template>
    <div id="bigger">
        <h4>Torrent file or Directory</h4>
        <hr />
        <file-pond
            name="file"
            ref="filePondUploadZone"
            class-name="upload-zone"
            label-idle="Drop files here..."
            allow-multiple="true"
            accepted-file-types="application/x-bittorrent, application/octet-stream"
            :server="filePondConfig.server"
            :files="filePondConfig.files"
            @processfile="handleProcessFile" />
    </div>
</template>

<style scoped>
</style>

<script setup>
import { ref, reactive } from "vue";
import VueFilePond from "vue-filepond";
import FilePondPluginFileValidateType from "filepond-plugin-file-validate-type";
import config from "../config";

// Component
const FilePond = VueFilePond(
    FilePondPluginFileValidateType
);

// Data
const viewState = reactive({
    resultMode : config.result_mode,
    biggerFiles: [],
    biggerText : "",
});

const filePondUploadZone = ref(null);

const filePondConfig = reactive({
    server: {
        url    : config.api.base_url,
        process: {
            url            : config.api.entry_urls.bigger,
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

// Methods
const handleProcessFile = (error, file) => {
    if (error) {
        console.log("Oops", error);
    }else{
        // Add the bigger file to bigger file list
        viewState.biggerFiles.unshift(file.serverId);

        // Transform the bigger file to text
        viewState.biggerText = viewState.biggerFiles
                .filter(file => file.ok === true)
                .map(file => `${file.data.file} => ${file.data.big}`)
                .join("\n");

        // Remove the completed file from drop zone
        filePondUploadZone.value.removeFile(file.id);
    }
};
</script>
