import axios from 'axios'
import router from '@/router'
export default {
    actions: {
        async getAllTripsWhichAttach({commit, state}) {
            if (state.token != null) {
                return new Promise((resolve, reject) => {
                    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                    axios.get('/attach_trips')
                    .then(response => {
                        let attach_trips = response.data;
                        commit('updateAttachTrips', attach_trips)

                        if (state.attach_trips != null && state.attach_trips.lenght != 0) {
                            var totalPrice = state.attach_trips.map(item => item).reduce((acc, item) => acc + item.price*item.count, 0);
                        }
                            
                        commit('updateTotalPrice', totalPrice)
                        resolve(response)
                    })
                    .catch(error => {
                        
                        if (error.response != null) {
                            if (error.response.status === 401) {
                                router.push({ name: 'logout' })
                            }
                        }
                        reject(error);
                    })
                })
            } 
        },
        incr({commit}, id) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
                axios.post('/attach_trips', {
                    id: id.toString(),   
                })
                .then(response => {
                    if (response.data.status == true) {
                        commit('incrMutation', id)
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
        decr({commit}, id) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
                axios.post('/attach_trips/decr', {
                    id: id.toString(),   
                })
                .then(response => {
                    if (response.data.status == true) {
                        commit('decrMutation', id)
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
        deleteAttachTrip({commit}, id) {
            return new Promise((resolve, reject) => {
         
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
                axios.delete('/attach_trips/'+id)
                .then(response => {
                    if (response.data.status == true) {
                        commit('deleteAttachTripsMutation', id)
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
                    reject(error);
                })
            })
        }
    },
    mutations: {
        updateAttachTrips(state, attach_trips) {
            state.attach_trips = attach_trips
        },
        deleteAttachTripsMutation(state, value) {
            var arr = state.attach_trips
            var attr = 'id'
            var i = arr.length;
            
            while(i--){
                if(arr[i] && arr[i][attr] === value) {
                    arr.splice(i,1);
                }
            }         
        },
        incrMutation(state, value) {
            var arr = state.attach_trips
            var attr = 'id'
            var i = arr.length;
            
            while(i--){
                if(arr[i] && arr[i][attr] === value) {
                    arr[i]['count']++;
                    state.totalPrice += arr[i]['price']
                }
            }
        },
        decrMutation(state, value) {
            var arr = state.attach_trips
            var attr = 'id'
            var i = arr.length;
            
            while(i--){
                if(arr[i] && arr[i][attr] === value && arr[i]['count'] > 1) {
                    arr[i]['count']--;
                    state.totalPrice -= arr[i]['price']
                }
            }
        },
        updateTotalPrice(state, totalPrice) {
            state.totalPrice = totalPrice
        },
        updateMessage(state, message) {
            state.message = message
        },
    },
    state: {
        attach_trips: [],
        totalPrice: 0,
        token: localStorage.getItem('token') || null
    },
    getters: {
        allTripsWhichAttach(state) {
            return state.attach_trips
        },
        getTotalPrice(state) {
            return state.totalPrice
        }
    }
}