<template>
    <div id="bigger">
        <h4>Torrent file or Directory</h4>
        <hr />
        <file-zone
            ref="fileZoneRef"
            :baseUrl="config.api.base_url"
            :entryUrl="config.api.entry_urls.bigger"
            :fileList="viewState.fileList"
            :fileText="viewState.fileText"
            :onProcessFile="handleProcessFile" />
    </div>
</template>

<style scoped>
</style>

<script setup>
import { ref, reactive } from "vue";
import config from "../config";
import FileZone from "../components/FileZone.vue";

const fileZoneRef = ref(null);

const viewState = reactive({
    resultMode: config.result_mode,

    fileList: [],
    fileText: "",
});

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
