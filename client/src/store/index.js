import Vue from 'vue'
import Vuex from 'vuex'
import category from "./modules/category"
import auth from "./modules/auth"
import tour from "./modules/tour"
import region from "./modules/region"
import user from "./modules/user"
import order from "./modules/order"

Vue.use(Vuex)


const store = new Vuex.Store({
    state: {
        layout: 'default-layout'
    },
    mutations: {
        setLayout(state, payload) {
            state.layout = payload
        }
    },
    actions: {},
    getters: {
        layout(state) {
            return state.layout
        }
    },
    modules: {
        category,
        auth,
        tour,
        region,
        user,
        order,
    }
})

export default store