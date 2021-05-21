import axios from 'axios'
export default {
    actions: {
        async getAllCategories({commit}) {
            let res = await axios
                .get('/categories');

            let categories = res.data.categories;
            commit('updateCategories', categories)
        },
        async getOneCategory({commit}, id) {
            let res = await axios
                .get('/categories/' + id);

            let category = res.data;
            commit('updateCategory', category)
        },
        async getOneCategoryBySlug({commit}, slug) {
            let res = await axios
                .get('/categories/slug/' + slug);

            let category = res.data;
            commit('updateCategory', category)
        },
        createCategory({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/categories', {
                    name: data.name,
                    icon: data.icon
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
        updateCategory({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.put('/categories/'+data.slug, {
                    name: data.name,
                    icon: data.icon
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
        deleteCategory({commit}, slug) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.delete('/categories/'+slug)
                .then(response => {           
                    if (response.data.status == true) {
                        commit('updateSuccessMessage', response.data.message)
                        commit('deleteCategoryMutation', slug)
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
        updateCategories(state, categories) {
            state.categories = categories
        },
        updateCategory(state, category) {
            state.category = category
        },
        deleteCategoryMutation(state, slug) {
            let arr = state.categories;
            let attr = 'slug';
            let value = slug;
            let i = arr.length;
            while(i--){
                if( arr[i] && arr[i].hasOwnProperty(attr) && arr[i][attr] === value ){
                    arr.splice(i,1);
                }
            }
        },
    },
    state: {
        categories: [],
        category: {}
    },
    getters: {
        allCategories(state) {
            return state.categories
        },
        oneCategory(state) {
            return state.category
        }
    }
}