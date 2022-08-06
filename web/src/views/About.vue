<template>
    <div id="home">
       <section-header>About</section-header>
       <div class="card card-about">
            <div class="card-header">Info</div>
            <div class="card-body p-0">
                <div class="row m-0">
                    <div class="col-lg-2 bg-info text-white fw-bold">Name</div>
                    <div class="col-lg-10">{{ viewState.name }}</div>
                </div>
                <div class="row m-0">
                    <div class="col-lg-2 bg-info text-white fw-bold">Mode</div>
                    <div class="col-lg-10">{{ viewState.mode }}</div>
                </div>
                <div class="row m-0">
                    <div class="col-lg-2 bg-info text-white fw-bold">Build Time</div>
                    <div class="col-lg-10">{{ defaultEmpty(viewState.build.time) }}</div>
                </div>
                <div class="row m-0">
                    <div class="col-lg-2 bg-info text-white fw-bold">Build Hash</div>
                    <div class="col-lg-10">{{ defaultEmpty(viewState.build.hash) }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.card-about > .card-body > .row > div {
    padding-top: 2px;
    padding-bottom: 2px;
}
</style>

<script setup>
import { reactive } from "vue";
import config from "../config";
import {
    SectionHeader,
}from "../components";

// State
const entryApi = config.api.base_url + config.api.entry_urls.about;
const viewState = reactive({
    name : "",
    mode : "",
    build: {},
});

// Lifecycle
(async () => {
    try {
        const response = await fetch(entryApi, { method: "POST"});
        const data     = await response.json();

        viewState.name  = data.name;
        viewState.mode  = data.mode;
        viewState.build = data.build;
    }catch(e) {
        console.log("About page: ", e);
    }
})();

// Methods
const defaultEmpty = value => {
    if (typeof(value) === "string" && value.trim().length > 0) {
        return value;
    }

    return "---";
};
</script>
