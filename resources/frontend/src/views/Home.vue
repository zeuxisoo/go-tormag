<template>
    <div class="home">
        <form method="post" action="" enctype="multipart/form-data"
            v-bind:class="['upload-box', draggingClass]"
            v-on:dragover.prevent="setIsDragging($event, true)"
            v-on:dragenter.prevent="setIsDragging($event, true)"
            v-on:dragleave.prevent="setIsDragging($event, false)"
            v-on:dragend.prevent="setIsDragging($event, false)"
            v-on:drop.prevent="handleDrop($event)">
            <div class="text-center">
                <input type="file" name="filess" id="file" class="input-file" v-on:change="handleChoose($event)" multiple allowdirs />
                <label for="file">
                    <p><strong>Choose a file</strong></p>
                    <p>- OR -</p>
                    <p><strong>Drag it here</strong></p>
                </label>
                <button type="submit" class="upload-button">Upload</button>
            </div>
        </form>
    </div>
</template>

<style lang="scss" scoped>
.upload-box {
    outline: 2px dashed #c7cbf8;
    outline-offset: -10px;
    transition: outline-offset .15s ease-in-out, background-color .15s linear;

    font-size: 1.25rem;
    background-color: #e5ecee;
    position: relative;
    padding: 100px 20px;

    &.is-dragging {
        outline: 2px dashed #0088FF;
        outline-offset: -20px;
        background: #e2efff;
    }

    .input-file {
        width: 0.1px;
        height: 0.1px;
        opacity: 0;
        overflow: hidden;
        position: absolute;
        z-index: -1;
    }

    .upload-button {
        display: none;
    }
}
</style>

<script>
export default {
    name: "home",

    data() {
        return {
            isDragging: false
        }
    },

    computed: {
        draggingClass() {
            return this.isDragging === true ? 'is-dragging' : '';
        }
    },

    methods: {
        setIsDragging(e, state) {
            this.isDragging = state;
        },

        handleDrop(e) {
            this.setIsDragging(e, false);

            // TODO: upload drag and drop file or directory
            var dt = event.dataTransfer;
            if (dt.items && dt.items.length && "webkitGetAsEntry" in dt.items[0]) {
                console.log(1, dt.items);
            } else if ("getFilesAndDirectories" in dt) {
                console.log(2, dt);
            } else if (dt.files) {
                console.log(3, dt.files);
            } else {
                console.log(4);
            }
        },

        handleChoose(e) {
            // TODO: upload single file
            console.log(e.target.files);
        }
    }
}
</script>
