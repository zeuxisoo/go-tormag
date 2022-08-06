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
        <result-info
            :resultMode="viewState.resultMode"
            :fileList="viewState.fileList"
            @changeResultMode="handleChangeResultMode" />

        <transition enter-active-class="animate__animated animate__fadeInUp" leave-active-class="animate__animated animate__fadeOutUp" mode="out-in">
            <alert-block v-if="viewState.fileList.length <= 0" key="converted-files-empty">
                Please drop the files to the drop zone first
            </alert-block>
            <div v-else key="converted-files-results">
                <transition-group enter-active-class="animate__animated animate__bounceInUp" leave-active-class="animate__animated animate__bounceOutDown" tag="div" v-if="matchResultMode('list')">
                    <div class="card card-default mb-3" v-for="biggerFile in viewState.fileList" :key="biggerFile.data.id">
                        <div class="card-header font-weight-bold">{{ biggerFile.data.file }}</div>

                        <div class="card-body p-0" v-if="biggerFile.ok === true">
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mt-1 mb-1 font-weight-bold text-center bg-info text-white">ID</div>
                                <div class="col-lg-10 p-1 mt-1 mb-1">{{ biggerFile.data.id }}</div>
                            </div>
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mb-1 font-weight-bold text-center bg-info text-white">MD5</div>
                                <div class="col-lg-10 p-1 mb-1">{{ biggerFile.data.md5 }}</div>
                            </div>
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mb-1 font-weight-bold text-center bg-info text-white">Big File</div>
                                <div class="col-lg-10 p-1 mb-1">{{ biggerFile.data.big }}</div>
                            </div>
                        </div>

                        <div class="card-body p-0" v-else>
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mt-1 mb-1 font-weight-bold text-center bg-danger text-white">Error</div>
                                <div class="col-lg-10 p-1 mt-1 mb-1">{{ biggerFile.message }}</div>
                            </div>
                        </div>
                    </div>
                </transition-group>

                <textarea class="form-control" rows="20" v-if="matchResultMode('text')" :value="viewState.fileText"></textarea>
            </div>
        </transition>
    </div>
</template>

<style scoped>
</style>

<script setup>
import { ref, reactive } from "vue";
import config from "../config";
import {
    SectionHeader,
    ResultInfo,
    FileZone,
    AlertBlock,
} from "../components";

// State
const fileZoneRef = ref(null);

const viewState = reactive({
    resultMode: config.result_mode,
    fileList  : [],
    fileText  : "",
});

// Methods
const matchResultMode = mode => viewState.resultMode === mode;

// Methods components binding
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

const handleChangeResultMode = mode => viewState.resultMode = mode;
</script>
