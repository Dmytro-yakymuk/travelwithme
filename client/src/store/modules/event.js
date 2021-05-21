import axios from 'axios'
export default {
    actions: {
        async getAllEvents({commit}, limit) {
            let res = await axios
                .get('/events?limit=' + limit);

            let events = res.data.events;
            commit('updateEvents', events)
        },
        async getOneEvent({commit}, id) {
            let res = await axios
                .get('/events/' + id);
            commit('updateEvent', res.data)
        },
        createEvent({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/events', {
                    text: data.text
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
                    reject(error)
                })
            })
        },
        updateEvent({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.put('/events/'+data.id, {
                    text: data.text
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
                    reject(error)
                })
            })
        },
        deleteEvent({commit}, id) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.delete('/events/'+id)
                .then(response => {           
                    if (response.data.status == true) {
                        commit('updateSuccessMessage', response.data.message)
                        commit('deleteEventMutation', id)
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
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                    reject(error)
                })
            })
        },
    },
    mutations: {
        updateEvents(state, events) {
            state.events = events
        },
        updateEvent(state, event) {
            state.event = event
        },
        deleteEventMutation(state, id) {
            let arr = state.events;
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
        events: [],
        event: {}
    },
    getters: {
        allEvents(state) {
            return state.events
        },
        oneEvent(state) {
            return state.event
        }
    }
}