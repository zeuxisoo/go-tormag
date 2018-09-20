<template>
    <div id="app">
        <nav class="navbar navbar-expand-lg navbar-light bg-light fixed-top">
            <router-link to="/" class="navbar-brand">Tormag</router-link>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbar-content" aria-controls="navbar-content" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>

            <div class="collapse navbar-collapse" id="navbar-content">
                <ul class="navbar-nav mr-auto">
                    <li class="nav-item active">
                        <router-link to="/" class="nav-link">Home</router-link>
                    </li>
                </ul>
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <router-link to="/about" class="nav-link">About</router-link>
                    </li>
                </ul>
            </div>
        </nav>

        <div class="main">
            <div class="container">
                <router-view />
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.main {
    padding-top: 70px;
    padding-bottom: 30px;
    margin-bottom: 0;
}
</style>

<script>
import $ from 'jquery';
import "bootstrap/dist/js/bootstrap.js";
import "bootstrap/dist/css/bootstrap.css";



export default {
    name: "app",

    mounted() {
        // Auto collapse then menu when navbar item clicked
        $(document).on("click", ".navbar-nav > li > a", () => {
            $('.navbar-collapse').collapse('hide');
        });

        // Auto active nav item when it clicked
        $(document).on('click', '.nav-item .nav-link', () => {
            // Remove active state on other nav item
            $(".nav-item").removeClass('active');

            // Add active state on clicked item parent
            $(this).parent().addClass('active');
        });
    }
};
</script>
