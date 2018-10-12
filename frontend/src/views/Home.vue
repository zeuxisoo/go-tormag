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

        <transition enter-active-class="animated fadeInDown" leave-active-class="animated fadeOutUp" mode="out-in">
            <div class="card card-default" v-if="convertedFiles.length <= 0" key="converted-files-empty">
                <div class="card-header">Converted Result</div>
                <div class="card-body">
                    <div class="alert alert-info text-center" role="alert">
                        Please drop the files to the drop zone first
                    </div>
                </div>
            </div>
            <div v-else key="converted-files-results">
                <transition-group enter-active-class="animated bounceInUp" leave-active-class="animated bounceOutDown" tag="div">
                    <div class="card card-default mb-3" v-for="convertedFile in convertedFiles" v-bind:key="convertedFile.data.id">
                        <div class="card-header font-weight-bold">{{ convertedFile.data.file }}</div>
                        <div class="card-body p-0">
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mt-1 mb-1 font-weight-bold text-center bg-info text-white">ID</div>
                                <div class="col-lg-10 p-1">{{ convertedFile.data.id }}</div>
                            </div>
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mb-1 font-weight-bold text-center bg-info text-white">MD5</div>
                                <div class="col-lg-10 p-1">{{ convertedFile.data.md5 }}</div>
                            </div>
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mb-1 font-weight-bold text-center bg-info text-white">Magnet</div>
                                <div class="col-lg-10 p-1">
                                    <textarea class="form-control" v-bind:value="convertedFile.data.magnet"></textarea>
                                </div>
                            </div>
                        </div>
                    </div>
                </transition-group>
            </div>
        </transition>
    </div>
</template>

<style lang="scss">
@import "~filepond/dist/filepond.min.css";
@import "~animate.css/animate.min.css";

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
                this.convertedFiles.unshift(file.serverId);

                // Remove the completed file from drop zone
                this.$refs.uploadZone.removeFile(file.id);
            }
        }
    }
}
</script>
