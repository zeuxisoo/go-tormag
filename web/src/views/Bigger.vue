<template>
    <div id="bigger">
        <section-header>Torrent file or Directory</section-header>
        <file-zone
            ref="fileZoneRef"
            :baseUrl="config.api.base_url"
            :entryUrl="config.api.entry_urls.bigger"
            :fileList="viewState.fileList"
            :fileText="viewState.fileText"
            @processFile="handleProcessFile" />

        <section-header>Find Result</section-header>
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
                    OK <span class="badge badge-light">{{ biggerCount.ok }}</span>
                </button>
                &nbsp;
                <button type="button" class="btn btn-danger">
                    Error <span class="badge badge-light">{{ biggerCount.error }}</span>
                </button>
            </div>
        </div>
        <hr />
    </div>
</template>

<style scoped>
</style>

<script setup>
import { ref, reactive, computed } from "vue";
import config from "../config";
import SectionHeader from "../components/SectionHeader.vue";
import FileZone from "../components/FileZone.vue";

// Data
const fileZoneRef = ref(null);

const viewState = reactive({
    resultMode: config.result_mode,
    fileList  : [],
    fileText  : "",
});

// Computed
const biggerCount = computed(() => {
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
                .map(file => `${file.data.file} => ${file.data.big}`)
                .join("\n");

        // Remove the completed file from drop zone
        fileZoneRef.value.filePondRef.removeFile(file.id);
    }
}

const isTargetResultMode = (name) => viewState.resultMode === name;
const changeResultMode = (name) => viewState.resultMode = name.toLowerCase();
</script>
