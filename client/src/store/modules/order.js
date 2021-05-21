import axios from 'axios'
export default {
    actions: {
        async getAllOrders({commit}) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
                axios.get('/orders')
                .then(response => {
                    var orders = [];
                    if (response.data != null) {
                        var orders = response.data;
                    }
                    
                    commit('updateOrders', orders);
                    resolve(response)
                })
                
                .catch(error => {
                    reject(error)
                })
            })
        },
        async getOneOrder({commit}, id) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
                axios.get('/orders/'+id)
                .then(response => {
                    var order = {};
                    if (response.data != null) {
                        var order = response.data;
                    }
                    
                    commit('updateOrder', order);
                    resolve(response)
                })
                
                .catch(error => {
                    reject(error)
                })
            })
        },
        updatePaid({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.put('/orders/paid/'+data.id, {
                    paid: data.paid
                })
                .then(response => {
                    resolve(response)
                })
                .catch(error => {
                    reject(error)
                })
            })
        },

        createOrder({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/orders', {
                    id: data.id,
                    phone: data.phone,
                    text: data.text,
                    totalPrice: data.totalPrice
                })
                .then(response => {
                    if (response.data.status == true) {
                        commit('updateSuccessMessage', response.data.message)
                        setTimeout(() => {
                            commit('updateSuccessMessage', '')
                        }, 3000)
                    } else if(response.data.status == false) {
                        commit('updateFailMessage', response.data.message)
                        setTimeout(() => {
                            commit('updateFailMessage', '')
                        }, 3000)
                    }
                    resolve(response)
                })
                .catch(error => {
                    if(typeof error.response !== "undefined"){
                        if (error.response != null) {
                            if (error.response.status === 401) {
                                this.$router.push({ name: 'logout' })
                            }
                        }
                        if(error.response.data.status == false) {
                            commit('updateFailMessage', error.response.data.message)
                            setTimeout(() => {
                                commit('updateFailMessage', '')
                            }, 3000)
                        }
                    }
                    reject(error)
                })
            })
        },

        downloadOrder({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/orders/download', {
                    id: data.id,
                })
                .then(response => {
                    if (response.data.status == true) {
                        commit('updateSuccessMessage', response.data.message)
                        setTimeout(() => {
                            commit('updateSuccessMessage', '')
                        }, 3000)
                    } else if(response.data.status == false) {
                        commit('updateFailMessage', response.data.message)
                        setTimeout(() => {
                            commit('updateFailMessage', '')
                        }, 3000)
                    }
                    resolve(response)
                })
                .catch(error => {
                    if(typeof error.response !== "undefined"){
                        if (error.response != null) {
                            if (error.response.status === 401) {
                                this.$router.push({ name: 'logout' })
                            }
                        }
                        if(error.response.data.status == false) {
                            commit('updateFailMessage', error.response.data.message)
                            setTimeout(() => {
                                commit('updateFailMessage', '')
                            }, 3000)
                        }
                    }
                    reject(error)
                })
            })
        },
    },
    mutations: {
        updateOrders(state, orders) {
            state.orders = orders
        },
        updateOrder(state, order) {
            state.order = order
        },
    },
    state: {
        orders: [],
        order: {},
    },
    getters: {
        allOrders(state) {
            return state.orders
        },
        oneOrder(state) {
            return state.order
        },
    }
}