import path from 'path';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { VitePWA  } from 'vite-plugin-pwa';

// https://vitejs.dev/config/
export default defineConfig({
    server: {
        port: 5000,
    },
    plugins: [
        vue(),
        VitePWA({
            scope: "/",
            registerType: "autoUpdate",
            manifest: {
                background_color: "#FFFFFF",
                theme_color: "#4DBA87",
                icons: [
                    {
                        "src": "/static/img/icons/android-chrome-512x512.png",
                        "sizes": "512x512",
                        "type": "image/png"
                    }
                ],
            },
        }),
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
