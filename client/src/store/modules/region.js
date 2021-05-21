import axios from 'axios'
export default {
    actions: {
        async getAllRegions({commit}, limit) {
            let res = await axios
                .get('/regions?limit=' + limit);

            let regions = res.data.regions;
            commit('updateRegions', regions)
        },
        async getOneRegion({commit}, id) {
            let res = await axios
                .get('/regions/' + id);
            commit('updateRegion', res.data)
        },
        async getOneRegionBySlug({commit}, slug) {
            let res = await axios
                .get('/regions/slug/' + slug);
            commit('updateRegion', res.data)
        },
        createRegion({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/regions', {
                    name: data.name
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
        updateRegion({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.put('/regions/'+data.slug, {
                    name: data.name
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
        deleteRegion({commit}, slug) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.delete('/regions/'+slug)
                .then(response => {           
                    if (response.data.status == true) {
                        commit('updateSuccessMessage', response.data.message)
                        commit('deleteRegionMutation', slug)
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
        updateRegions(state, regions) {
            state.regions = regions
        },
        updateRegion(state, region) {
            state.region = region
        },
        deleteRegionMutation(state, slug) {
            let arr = state.regions;
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
        regions: [],
        region: {}
    },
    getters: {
        allRegions(state) {
            return state.regions
        },
        oneRegion(state) {
            return state.region
        }
    }
}