<template>
    <div id="home">
        <h4>Torrent file or Directory</h4>
        <hr />
        <file-pond
            name="file"
            ref="uploadZone"
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
import { reactive } from "vue";
import VueFilePond from "vue-filepond";
import FilePondPluginFileValidateType from "filepond-plugin-file-validate-type";

// Component
const FilePond = VueFilePond(
    FilePondPluginFileValidateType
);

// Data
const filePondConfig = reactive({
    server: {},
    files : [],
})

// Methods
const handleProcessFile = (error, file) => {

};
</script>
