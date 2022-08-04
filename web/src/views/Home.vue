<template>
    <div id="home">
        <section-header>Torrent file or Directory</section-header>
        <file-zone
            ref="fileZoneRef"
            :baseUrl="config.api.base_url"
            :entryUrl="config.api.entry_urls.convert"
            :fileList="viewState.fileList"
            :fileText="viewState.fileText"
            @processFile="handleProcessFile" />

        <section-header>Converted Result</section-header>
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
                    <div class="card text-bg-light mb-3" v-for="convertedFile in viewState.fileList" :key="convertedFile.data.id">
                        <div class="card-header fw-bold">{{ convertedFile.data.file }}</div>

                        <div class="card-body p-0" v-if="convertedFile.ok">
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
                                    <div class="row">
                                        <div class="col-lg-11 m-0">
                                            <textarea class="form-control" :value="convertedFile.data.magnet"></textarea>
                                        </div>
                                        <div class="col-lg-1 m-0 d-grid gap-2">
                                            <button class="btn btn-secondary" @click="copyMagnet(convertedFile.data.magnet)">
                                                <i class="bi bi-clipboard" style="font-size: 18px"></i>
                                            </button>
                                        </div>
                                    </div>
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

                <textarea class="form-control" rows="20" v-if="matchResultMode('text')" :value="viewState.fileText"></textarea>
            </div>
        </transition>

        <div class="position-relative">
            <div class="toast-container bottom-0 start-50 translate-middle-x">
                <div class="toast align-items-center border-0">
                    <div class="d-flex">
                        <div class="toast-body">
                            Magnet Copied
                        </div>
                        <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast"></button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
</style>

<script setup>
import { ref, reactive } from "vue";
import { Toast } from "bootstrap";
import config from "../config";
import SectionHeader from "../components/SectionHeader.vue";
import FileZone from "../components/FileZone.vue";
import ResultInfo from "../components/ResultInfo.vue";
import AlertBlock from "../components/AlertBlock.vue";

// Data
const fileZoneRef = ref(null);

const viewState = reactive({
    resultMode: config.result_mode,
    fileList  : [],
    fileText  : "",
});

// Methods
const matchResultMode = mode => viewState.resultMode === mode;

const copyMagnet = magnet => {
    navigator.clipboard.writeText(magnet)
        .then(
            () => {
                new Toast(document.querySelector(".toast")).show();
            },
            () => console.log('Copy failed'),
        );
}

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
                .map(file => file.data.magnet)
                .join("\n");

        // Remove the completed file from drop zone
        fileZoneRef.value.filePondRef.removeFile(file.id);
    }
};

const handleChangeResultMode = mode => viewState.resultMode = mode;
</script>
