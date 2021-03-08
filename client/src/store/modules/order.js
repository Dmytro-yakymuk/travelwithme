import axios from 'axios'
export default {
    actions: {
        async getAllOrders({commit}) {
            const res = await axios
                .get('/orders');

            const orders = res.data;
            commit('updateOrders', orders)
        }
    },
    mutations: {
        updateOrders(state, orders) {
            state.orders = orders
        }
    },
    state: {
        orders: []
    },
    getters: {
        allOrders(state) {
            return state.orders
        }
    }
}