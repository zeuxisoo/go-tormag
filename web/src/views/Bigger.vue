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
            v-model:resultMode="viewState.resultMode"
            :fileList="viewState.fileList" />
    </div>
</template>

<style scoped>
</style>

<script setup>
import { ref, reactive, computed } from "vue";
import config from "../config";
import SectionHeader from "../components/SectionHeader.vue";
import ResultInfo from "../components/ResultInfo.vue";
import FileZone from "../components/FileZone.vue";

// Data
const fileZoneRef = ref(null);

const viewState = reactive({
    resultMode: config.result_mode,
    fileList  : [],
    fileText  : "",
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
</script>
