import path from 'path';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

// https://vitejs.dev/config/
export default defineConfig({
    server: {
        port: 5000,
    },
    plugins: [
        vue(),
    ],
    resolve: {
        alias: {
            "~bootstrap": path.resolve(__dirname, "node_modules/bootstrap"),
            "~bootstrap-icons": path.resolve(__dirname, "node_modules/bootstrap-icons"),
            "~filepond": path.resolve(__dirname, "node_modules/filepond"),
            "~animate.css": path.resolve(__dirname, "node_modules/animate.css"),
        },
    },
});
