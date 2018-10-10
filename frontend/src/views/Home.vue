<template>
    <div class="home">
        <div class="card card-default">
            <div class="card-header">Torrent file or Directory</div>
            <div class="card-body">
                <file-pond
                    name="file"
                    ref="uploadZone"
                    class-name="upload-zone"
                    label-idle="Drop files here..."
                    allow-multiple="true"
                    accepted-file-types="application/x-bittorrent, application/octet-stream"
                    v-bind:server="serverOptions"
                    v-bind:files="uploadFiles"
                    v-on:processfile="handleProcessFile" />
            </div>
        </div>

        <hr>

        <div class="card card-default" v-if="convertedFiles.length <= 0">
            <div class="card-header">Converted Result</div>
            <div class="card-body">
                <div class="alert alert-info text-center" role="alert">
                    Please drop the files to the drop zone first
                </div>
            </div>
        </div>

        <div class="card card-default mb-3" v-for="convertedFile in convertedFiles" v-bind:key="convertedFile.data.file">
            <div class="card-header">{{ convertedFile.data.file }}</div>
            <div class="card-body">
                {{ convertedFile.data.magnet }}
            </div>
        </div>
    </div>
</template>

<style lang="scss">
@import "~filepond/dist/filepond.min.css";

.filepond--panel-root {
    background-color: rgb(227, 237, 243);
}

.filepond--drop-label {
    color: rgb(169, 166, 206);
}
</style>

<script>
import VueFilePond from 'vue-filepond';

const FilePond = VueFilePond();

export default {
    name: "home",

    components: {
        "file-pond": FilePond
    },

    data() {
        return {
            uploadFiles   : [],
            convertedFiles: [],
            serverOptions : {
                url    : 'http://127.0.0.1:3000',
                process: {
                    url            : '/convert',
                    method         : 'POST',
                    withCredentials: false,
                    headers        : {},
                    timeout        : 7000,
                    onload         : (response) => { return JSON.parse(response) },
                    onerror        : null
                }
            }
        }
    },

    methods: {
        handleProcessFile(error, file) {
            if (error) {
                console.log("Oops", error);
            }else{
                // Add the converted file to converted file list
                this.convertedFiles.push(file.serverId);

                // Remove the completed file from drop zone
                this.$refs.uploadZone.removeFile(file.id);
            }
        }
    }
}
</script>
