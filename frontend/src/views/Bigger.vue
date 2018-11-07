<template>
    <div id="bigger">
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

        <h4>Find Result</h4>
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
            <div class="col-lg-3 text-right d-none d-lg-block"> <!-- display on >= lg size only -->
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

        <transition enter-active-class="animated fadeInUp" leave-active-class="animated fadeOutUp" mode="out-in">
            <div class="alert alert-info text-center" role="alert" v-if="biggerFiles.length <= 0" key="converted-files-empty">
                Please drop the files to the drop zone first
            </div>
            <div v-else key="converted-files-results">
                <transition-group enter-active-class="animated bounceInUp" leave-active-class="animated bounceOutDown" tag="div" v-if="isTargetResultMode('list') === true">
                    <div class="card card-default mb-3" v-for="biggerFile in biggerFiles" v-bind:key="biggerFile.data.id">
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

                <textarea class="form-control" rows="20" v-if="isTargetResultMode('text') === true" v-bind:value="biggerText"></textarea>
            </div>
        </transition>
    </div>
</template>

<style lang="scss">
@import "~filepond/dist/filepond.min.css";

.filepond--panel-root {
    background-color: rgb(238, 243, 227);
}

.filepond--drop-label {
    color: rgb(187, 206, 166);
}
</style>

<script>
import config from '../config';
import VueFilePond from 'vue-filepond';

const FilePond = VueFilePond();

export default {
    name: 'bigger',

    components: {
        "file-pond": FilePond
    },

    data() {
        return {
            resultMode : config.result_mode,

            uploadFiles   : [],
            biggerFiles   : [],
            biggerText    : "",
            serverOptions : {
                url    : config.api.base_url,
                process: {
                    url            : config.api.entry_urls.bigger,
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
        biggerCount() {
            const okCount = this.biggerFiles.filter(file => file.ok === true).length;

            return {
                ok   : okCount,
                error: this.biggerFiles.length - okCount
            }
        }
    },

    methods: {
        handleProcessFile(error, file) {
            if (error) {
                console.log("Oops", error);
            }else{
                // Add the bigger file to bigger file list
                this.biggerFiles.unshift(file.serverId);

                // Transform the bigger file to text
                this.biggerText = this.biggerFiles
                        .filter(file => file.ok === true)
                        .map(file => {
                            return `${file.data.file} => ${file.data.big}`
                        })
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
                    this.resultMode = config.result_mode;
                    break;
            }
        },

        isTargetResultMode(name) {
            return this.resultMode === name;
        }
    }
}
</script>
