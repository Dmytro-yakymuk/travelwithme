import axios from 'axios'
import moment from 'moment'
export default {
    actions: {
        async getAllTrips({commit}) {
            axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
            let query = '/trips'

            var res = await axios
                .get(query);

            var trips = [];
            if (res.data != null) { 
                trips = res.data;
            }
            
            commit('updateTripsMutation', trips);
        },
        async getOneTrip({commit}, data) {
            axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
            var res = await axios
                .get('/trips/' + data.id);

            if (res.data == null) {
                var trip = null;
            } else {
                var trip = res.data;
                trip.start = moment(String(trip.start)).format('yyyy-MM-DD');
                trip.end = moment(String(trip.end)).format('yyyy-MM-DD');
            }
            commit('updateTripMutation', trip)
        },
        createTrip({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/trips', {
                    start: data.start,
                    end: data.end,
                    price: data.price,
                    count: data.count,
                    tour_id: data.tour_id
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
        updateTrip({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.put('/trips/' + data.id, {
                    start: data.start,
                    end: data.end,
                    price: data.price,
                    count: data.count,
                    tour_id: data.tour_id
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
        deleteTrip({commit}, id) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.delete('/trips/'+id)
                .then(response => {
                    if (response.data.status == true) {
                        commit('updateSuccessMessage', response.data.message)
                        commit('deleteTripMutation', id)
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
        }
    },
    mutations: {
        updateTripsMutation(state, trips) {
            state.trips = trips;
        },
        updateTripMutation(state, trip) {
            state.trip = trip;
        },
        deleteTripMutation(state, id) {
            let arr = state.trips;
            let attr = 'id';
            let value = id;
            let i = arr.length;
            while(i--){
                if( arr[i] && arr[i].hasOwnProperty(attr) && arr[i][attr] === value ){
                    arr.splice(i,1);
                }
            }
        },
    },
    state: {
        trips: [],
        trip: {}
    },
    getters: {
        allTrips(state) {
            return state.trips
        },
        oneTrip(state) {
            return state.trip
        },
    }
}