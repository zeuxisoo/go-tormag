<template>
    <div id="home">
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

        <h4>Converted Result</h4>
        <hr />
        <div class="row">
            <div class="col-lg-9">
                <button type="button" v-bind:class="['btn', { 'btn-info': isTargetResultMode('list') }]" @click="changeResultMode('list')">
                    <i class="fas fa-list-ul"></i> List
                </button>
                &nbsp;
                <button type="button" v-bind:class="['btn', { 'btn-info': isTargetResultMode('text') }]" @click="changeResultMode('text')">
                    <i class="fas fa-align-justify"></i> Text Only
                </button>
            </div>
            <div class="col-lg-3 text-end d-none d-lg-block"> <!-- display on >= lg size only -->
                <button type="button" class="btn btn-success">
                    OK <span class="badge badge-light">{{ convertedCount.ok }}</span>
                </button>
                &nbsp;
                <button type="button" class="btn btn-danger">
                    Error <span class="badge badge-light">{{ convertedCount.error }}</span>
                </button>
            </div>
        </div>
        <hr />

        <div class="alert alert-info text-center" role="alert" v-if="viewState.convertedFiles.length <= 0" key="converted-files-empty">
            Please drop the files to the drop zone first
        </div>
        <div v-else key="converted-files-results">
            <transition-group enter-active-class="animate__animated animate__bounceInUp" leave-active-class="animate__animated animate__bounceOutDown" tag="div" v-if="isTargetResultMode('list')">
                <div class="card text-bg-light mb-3" v-for="convertedFile in viewState.convertedFiles" v-bind:key="convertedFile.data.id">
                    <div class="card-header fw-bold">{{ convertedFile.data.file }}</div>

                    <div class="card-body p-0" v-if="convertedFile.ok === true">
                        <div class="row m-0">
                            <div class="col-lg-2 p-1 mt-1 mb-1 fw-bold text-center bg-info text-white">ID</div>
                            <div class="col-lg-10 p-1 mt-1 mb-1">{{ convertedFile.data.id }}</div>
                        </div>
                        <div class="row m-0">
                            <div class="col-lg-2 p-1 mb-1 fw-bold text-center bg-info text-white">MD5</div>
                            <div class="col-lg-10 p-1 mb-1">{{ convertedFile.data.md5 }}</div>
                        </div>
                        <div class="row m-0">
                            <div class="col-lg-2 p-1 mb-1 fw-bold text-center bg-info text-white">Magnet</div>
                            <div class="col-lg-10 p-1 mb-1">
                                <textarea class="form-control" v-bind:value="convertedFile.data.magnet"></textarea>
                            </div>
                        </div>
                    </div>

                    <div class="card-body p-0" v-else>
                        <div class="row m-0">
                            <div class="col-lg-2 p-1 mt-1 mb-1 fw-bold text-center bg-danger text-white">Error</div>
                            <div class="col-lg-10 p-1 mt-1 mb-1">{{ convertedFile.message }}</div>
                        </div>
                    </div>
                </div>
            </transition-group>

            <textarea class="form-control" rows="20" v-if="isTargetResultMode('text') === true" v-bind:value="viewState.convertedText"></textarea>
        </div>
    </div>
</template>

<style scoped>
</style>

<script setup>
import { ref, reactive, computed } from "vue";
import VueFilePond from "vue-filepond";
import FilePondPluginFileValidateType from "filepond-plugin-file-validate-type";
import config from "../config";

// Component
const FilePond = VueFilePond(
    FilePondPluginFileValidateType
);

// Data
const viewState = reactive({
    resultMode    : config.result_mode,
    convertedFiles: [],
    convertedText : "",

});

const filePondUploadZone = ref(null);

const filePondConfig = reactive({
    server: {
        url    : config.api.base_url,
        process: {
            url            : config.api.entry_urls.convert,
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

// Computed
const convertedCount = computed(() => {
    const okCount = viewState.convertedFiles
        .filter(file => file.ok === true)
        .length;

    return {
        ok   : okCount,
        error: viewState.convertedFiles.length - okCount,
    }
});

// Methods
const handleProcessFile = (error, file) => {
    if (error) {
        console.log("Oops", error);
    }else{
        // Add the converted file to converted file list
        viewState.convertedFiles.unshift(file.serverId);

        // Transform the converted file to text
        viewState.convertedText = viewState.convertedFiles
                .filter(file => file.ok === true)
                .map(file => file.data.magnet)
                .join("\n");

        // Remove the completed file from drop zone
        filePondUploadZone.value.removeFile(file.id);
    }
};

const isTargetResultMode = (name) => viewState.resultMode === name;
const changeResultMode = (name) => viewState.resultMode = name.toLowerCase();
</script>
