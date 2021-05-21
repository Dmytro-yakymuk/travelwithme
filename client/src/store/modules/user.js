import axios from 'axios'
export default {
    actions: {
        async getAllUsers({commit}, map) {

            let query = '/users?'
            if (map['where'] != null) {
                query += 'where='+map['where']
            }
            if (map['order_by'] != null) {
                query += '&order_by='+map['order_by']
            }
            if (map['limit'] != null) {
                query += '&limit=' + map['limit']
            }

            const res = await axios
                .get(query);

            var users = [];
            if (res.data != null) { 
                users = res.data;
            }
            commit('updateUsers', users)
        },
        async getOneUser({commit}) {
            axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
            const res = await axios
                .get("/users/byToken");

            var user = {};
            if (res.data != null) { 
                user = res.data;
            }
            commit('updateUser', user)
        },
        updateUser({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.put('/users/'+data.id, {
                    surname: data.surname,
                    name: data.name,
                    patronymic: data.patronymic,
                    email: data.email,
                    old_password: data.old_password,
                    new_password: data.new_password,
                    phone: data.phone
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
        }
    },
    mutations: {
        updateUsers(state, users) {
            state.users = users
        },
        updateUser(state, user) {
            state.user = user
        }
    },
    state: {
        users: [],
        user: {}
    },
    getters: {
        allUsers(state) {
            return state.users
        },
        oneUser(state) {
            return state.user
        }
    }
}