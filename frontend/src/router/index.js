import {createRouter, createWebHistory} from "vue-router";
import Hosts from "../components/hosts/Hosts.vue"
import Edit from "../components/hosts/Edit.vue"
import Check from "../components/Check/Index.vue"
import System from "../components/system/System.vue"
import Fileinfo from "../components/fileinfo/Fileinfo.vue"
import Editor from "../components/editor/Editor.vue"
import Home from "../components/Home.vue"

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home,
    }, {
        path: '/hosts',
        name: 'hosts',
        component: Hosts,
    }, {
        path: '/edithosts',
        name: 'edithosts',
        component: Edit,
    }, {
        path: '/check',
        name: 'check',
        component: Check,
    }, {
        path: '/system',
        name: 'system',
        component: System,
    }, {
        path: '/fileinfo',
        name: 'fileinfo',
        component: Fileinfo,
    }, {
        path: '/editor',
        name: 'editor',
        component: Editor,
    }
]
const router = createRouter({
    history: createWebHistory(), // 路由模式：history模式
    routes: routes
})
export default router