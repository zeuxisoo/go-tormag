<template>
    <div id="home">
        <h4>Torrent file or Directory</h4>
        <hr />
        <file-zone
            ref="fileZoneRef"
            :baseUrl="config.api.base_url"
            :entryUrl="config.api.entry_urls.convert"
            :fileList="viewState.fileList"
            :fileText="viewState.fileText"
            @processFile="handleProcessFile" />

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

        <transition enter-active-class="animate__animated animate__fadeInUp" leave-active-class="animate__animated animate__fadeOutUp" mode="out-in">
            <div class="alert alert-info text-center" role="alert" v-if="viewState.fileList.length <= 0" key="converted-files-empty">
                Please drop the files to the drop zone first
            </div>
            <div v-else key="converted-files-results">
                <transition-group enter-active-class="animate__animated animate__bounceInUp" leave-active-class="animate__animated animate__bounceOutDown" tag="div" v-if="isTargetResultMode('list')">
                    <div class="card text-bg-light mb-3" v-for="convertedFile in viewState.fileList" v-bind:key="convertedFile.data.id">
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

                <textarea class="form-control" rows="20" v-if="isTargetResultMode('text') === true" v-bind:value="viewState.fileText"></textarea>
            </div>
        </transition>
    </div>
</template>

<style scoped>
</style>

<script setup>
import { ref, reactive, computed } from "vue";
import config from "../config";
import FileZone from "../components/FileZone.vue";

// Data
const fileZoneRef = ref(null);

const viewState = reactive({
    resultMode: config.result_mode,
    fileList  : [],
    fileText  : "",
});

// Computed
const convertedCount = computed(() => {
    const okCount = viewState.fileList
        .filter(file => file.ok === true)
        .length;

    return {
        ok   : okCount,
        error: viewState.fileList.length - okCount,
    }
});

// Methods
const handleProcessFile = (error, file) => {
    if (error) {
        console.log("Oops", error);
    }else{
        // Add the file to file list
        viewState.fileList.unshift(file.serverId);

        // Transform the file to text
        viewState.fileText = viewState.fileList
                .filter(file => file.ok === true)
                .map(file => file.data.magnet)
                .join("\n");

        // Remove the completed file from drop zone
        fileZoneRef.value.filePondRef.removeFile(file.id);
    }
};

const isTargetResultMode = (name) => viewState.resultMode === name;
const changeResultMode = (name) => viewState.resultMode = name.toLowerCase();
</script>
