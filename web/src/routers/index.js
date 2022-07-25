import { createRouter, createWebHistory  } from 'vue-router';
import { Home, Bigger, About } from '../views';

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home,
    },
    {
        path: '/bigger',
        name: 'bigger',
        component: Bigger,
    },
    {
        path: '/about',
        name: 'about',
        component: About,
    },
    {
        path: '/:pathMatch(.*)',
        redirect: '/',
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;
