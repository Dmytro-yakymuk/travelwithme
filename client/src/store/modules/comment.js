import axios from 'axios'
export default {
    actions: {
        async getAllComments({commit}, map) {
            //axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
            let query = '/comments?'
            if (map['user_id'] != null && map['user_id'] != '') {
                query += 'user_id='+map['user_id']
            }
            if (map['tour_id'] != null && map['tour_id'] != '') {
                query += '&tour_id='+map['tour_id']
            }

            var res = await axios
                .get(query);

            var comments = [];
            if (res.data != null) { 
                comments = res.data;
            }
            
            commit('updateCommentsMutation', comments);
        },
        async getOneComment({commit}, id) {
            let res = await axios
                .get('/comments/' + id);
            commit('updateCommentMutation', res.data[0])
        },
        createComment({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/comments', {
                    message: data.message,
                    star: data.star,
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
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                    reject(error)
                })
            })
        },
        updateComment({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.put('/comments/'+data.id, {
                    message: data.message,
                    star: data.star,
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
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                    reject(error)
                })
            })
        },
        deleteComment({commit}, id) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.delete('/comments/'+id)
                .then(response => {           
                    if (response.data.status == true) {
                        commit('updateSuccessMessage', response.data.message)
                        commit('deleteCommentMutation', id)
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
        updateCommentsMutation(state, comments) {
            state.comments = comments;
        },
        updateCommentMutation(state, comment) {
            state.comment = comment;
        },
        deleteCommentMutation(state, id) {
            let arr = state.comments;
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
        comments: [],
        comment: {},
    },
    getters: {
        allComments(state) {
            return state.comments
        },
        oneComment(state) {
            return state.comment
        },
    }
}