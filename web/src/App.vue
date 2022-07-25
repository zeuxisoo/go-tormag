<template>
    <div id="main">
        <nav class="navbar navbar-expand-md bg-light-blue fixed-top">
            <div class="container">
                <router-link :to="{ name: 'home' }" class="navbar-brand">Tormag</router-link>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarDefault" aria-controls="navbarDefault" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarDefault">
                    <ul class="navbar-nav mr-auto mb-2 mb-md-0">
                        <li class="nav-item">
                            <router-link :to="{ name: 'home' }" class="nav-link">Home</router-link>
                        </li>
                         <li class="nav-item">
                            <router-link :to="{ name: 'bigger' }" class="nav-link">Bigger</router-link>
                        </li>
                         <li class="nav-item">
                            <router-link :to="{ name: 'about' }" class="nav-link">About</router-link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
        <main class="container">
            <router-view></router-view>
        </main>
    </div>
</template>

<script setup>
</script>

<style>
@import "~bootstrap/dist/css/bootstrap.css";

body {
    margin-top: 60px;
}

.bg-light-blue {
    background-color: #e3f2fd;
}
</style>
