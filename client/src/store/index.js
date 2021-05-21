import Vue from 'vue'
import Vuex from 'vuex'
import category from "./modules/category"
import auth from "./modules/auth"
import tour from "./modules/tour"
import region from "./modules/region"
import user from "./modules/user"
import trip from "./modules/trip"
import attachTrip from "./modules/attachTrip"
import order from "./modules/order"
import event from "./modules/event"
import comment from "./modules/comment"

Vue.use(Vuex)


const store = new Vuex.Store({
    state: {
        layout: 'default-layout',
        successMessage: '',
        failMessage: ''
    },
    mutations: {
        setLayout(state, payload) {
            state.layout = payload
        },
        updateSuccessMessage(state, successMessage) {
            state.successMessage = successMessage
        },
        updateFailMessage(state, failMessage) {
            state.failMessage = failMessage
        }
    },
    actions: {},
    getters: {
        layout(state) {
            return state.layout
        },
        getSuccessMessage(state) {
            return state.successMessage
        },
        getFailMessage(state) {
            return state.failMessage
        }
    },
    modules: {
        category,
        auth,
        tour,
        region,
        user,
        trip,
        order,
        attachTrip,
        event,
        comment
    }
})

export default store