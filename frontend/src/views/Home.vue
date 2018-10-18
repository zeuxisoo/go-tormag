<template>
    <div class="home">
        <h4>Torrent file or Directory</h4>
        <hr />
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

        <h4>Converted Result</h4>
        <hr />
        <div class="row">
            <div class="col-lg-9">
                <button type="button" v-bind:class="['btn', { 'btn-info': isTargetResultMode('list') }]" v-on:click="changeResultMode('list')">
                    <i class="fas fa-list-ul"></i> List
                </button>
                &nbsp;
                <button type="button" v-bind:class="['btn', { 'btn-info': isTargetResultMode('text') }]" v-on:click="changeResultMode('text')">
                    <i class="fas fa-align-justify"></i> Text Only
                </button>
            </div>
            <div class="col-lg-3 text-right">
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

        <transition enter-active-class="animated fadeInUp" leave-active-class="animated fadeOutUp" mode="out-in">
            <div class="alert alert-info text-center" role="alert" v-if="convertedFiles.length <= 0" key="converted-files-empty">
                Please drop the files to the drop zone first
            </div>
            <div v-else key="converted-files-results">
                <transition-group enter-active-class="animated bounceInUp" leave-active-class="animated bounceOutDown" tag="div" v-if="isTargetResultMode('list') === true">
                    <div class="card card-default mb-3" v-for="convertedFile in convertedFiles" v-bind:key="convertedFile.data.id">
                        <div class="card-header font-weight-bold">{{ convertedFile.data.file }}</div>

                        <div class="card-body p-0" v-if="convertedFile.ok === true">
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mt-1 mb-1 font-weight-bold text-center bg-info text-white">ID</div>
                                <div class="col-lg-10 p-1 mt-1 mb-1">{{ convertedFile.data.id }}</div>
                            </div>
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mb-1 font-weight-bold text-center bg-info text-white">MD5</div>
                                <div class="col-lg-10 p-1 mb-1">{{ convertedFile.data.md5 }}</div>
                            </div>
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mb-1 font-weight-bold text-center bg-info text-white">Magnet</div>
                                <div class="col-lg-10 p-1 mb-1">
                                    <textarea class="form-control" v-bind:value="convertedFile.data.magnet"></textarea>
                                </div>
                            </div>
                        </div>

                        <div class="card-body p-0" v-else>
                            <div class="row m-0">
                                <div class="col-lg-2 p-1 mt-1 mb-1 font-weight-bold text-center bg-danger text-white">Error</div>
                                <div class="col-lg-10 p-1 mt-1 mb-1">{{ convertedFile.message }}</div>
                            </div>
                        </div>
                    </div>
                </transition-group>

                <textarea class="form-control" rows="20" v-if="isTargetResultMode('text') === true">{{ convertedText }}</textarea>
            </div>
        </transition>
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

const DefaultResultMode = 'list';

export default {
    name: "home",

    components: {
        "file-pond": FilePond
    },

    data() {
        return {
            resultMode : DefaultResultMode,

            uploadFiles   : [],
            convertedFiles: [],
            convertedText : '',
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

    computed: {
        convertedCount() {
            const okCount = this.convertedFiles.filter(file => file.ok === true).length;

            return {
                ok   : okCount,
                error: this.convertedFiles.length - okCount
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

                // Transform the converted file to text
                this.convertedText = this.convertedFiles
                        .filter(file => file.ok === true)
                        .map(file => file.data.magnet)
                        .join("\n");

                // Remove the completed file from drop zone
                this.$refs.uploadZone.removeFile(file.id);
            }
        },

        changeResultMode(name) {
            const mode = name.toLowerCase()

            switch(mode) {
                case 'list':
                    this.resultMode = mode;
                    break;
                case 'text':
                    this.resultMode = mode;
                    break;
                default:
                    this.resultMode = DefaultResultMode;
                    break;
            }
        },

        isTargetResultMode(name) {
            return this.resultMode === name;
        }
    }
}
</script>
