<template>
    <div id="about">
        <h4>About</h4>
        <hr />
        <div class="card card-default card-about">
            <div class="card-body p-0">
                <div class="row m-0">
                    <div class="col-lg-2 bg-info text-white font-weight-bold">Name</div>
                    <div class="col-lg-10">{{ name }}</div>
                </div>
                <div class="row m-0">
                    <div class="col-lg-2 bg-info text-white font-weight-bold">Mode</div>
                    <div class="col-lg-10">{{ mode }}</div>
                </div>
                <div class="row m-0">
                    <div class="col-lg-2 bg-info text-white font-weight-bold">Build Time</div>
                    <div class="col-lg-10">{{ build.time | defaultEmpty }}</div>
                </div>
                <div class="row m-0">
                    <div class="col-lg-2 bg-info text-white font-weight-bold">Build Hash</div>
                    <div class="col-lg-10">{{ build.hash | defaultEmpty }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.card-about .card-body .row:not(:first-child) {
    padding-top: 2px;
}
</style>

<script>
import Vue from 'vue';
import config from '../config';

export default {
    name: "about",

    data() {
        return {
            name: "",
            mode: "",
            build: {
                time: "",
                hash: "",
            }
        }
    },

    created() {
        this.fetchAboutData();
    },

    methods: {
        async fetchAboutData() {
            const entryApi = config.api.base_url + config.api.entry_urls.about;

            try {
                const response = await Vue.axios.post(entryApi);
                const data     = response.data;

                this.name  = data.name;
                this.mode  = data.mode;
                this.build = data.build;
            }catch(e) {
                console.log("About page: ", e);
            }
        }
    }
}
</script>
