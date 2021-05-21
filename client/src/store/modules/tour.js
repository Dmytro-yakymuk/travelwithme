import axios from 'axios'
export default {
    actions: {
        async getAllTours({commit}, map) {

            let query = '/tours?'
            if (map['category_id'] != null && map['category_id'] != '') {
                query += 'category_id='+map['category_id']
            }

            if (map['region_id'] != null && map['region_id'] != '') {
                query += '&region_id='+map['region_id']
            }

            if (map['user_id'] != null && map['user_id'] != '') {
                query += '&user_id='+map['user_id']
            }

            if (map['price'] != null && map['price'] != '') {
                query += '&price='+map['price']
            }

            if (map['date'] != null && map['date'] != '') {
                query += '&date='+map['date']
            }

            if (map['order_by'] != null) {
                query += '&order_by='+map['order_by']
            }
            if (map['limit'] != null) {
                query += '&limit=' + map['limit']
            }

            var res = await axios
                .get(query);

            if (res.data == null) {
                var tours = [];
            } else {
                var tours = res.data;
            }
            
            
            commit('updateToursMutation', tours);
        },
        async getOneTour({commit}, data) {
            var res = await axios
                .get('/tours/' + data.slug);

            if (res.data == null) {
                var tour = null;
            } else {
                var tour = res.data;
            }
            commit('updateTourMutation', tour)
        },
        async getAuditsByTourId({commit}, slug) {
            var res = await axios
                .get('/audits/' + slug);

            if (res.data == null) {
                var audits = [];
            } else {
                var audits = res.data;
            }
            commit('getAuditsByTourIdMutation', audits)
        },
        createTour({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/tours', {
                    title: data.title,
                    description: data.description,
                    text: data.text,
                    x: data.x,
                    y: data.y,
                    status: data.status,
                    category_id: data.category_id,
                    region_id: data.region_id,
                    selected_events: data.selected_events
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
        updateTour({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.put('/tours/'+data.slug, {
                    title: data.title,
                    slug: data.slug,
                    description: data.description,
                    text: data.text,
                    x: data.x,
                    y: data.y,
                    status: data.status,
                    category_id: data.category_id,
                    region_id: data.region_id,
                    selected_events: data.selected_events
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
        addAudit({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/audits', {
                    tour_id: data.tour_id,
                    message: data.message
                })
                .then(response => {
                    if (response.data.status == true) {
                        commit('addAuditMutation', response.data.result)
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
        publicTour({commit}, slug) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/tours/public/'+slug, {
                    slug: slug,
                })
                .then(response => {
                    if (response.data.status == true) {
                        commit('publicTourMutation', slug)
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
        activTour({commit}, slug) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/tours/activ/'+slug, {
                    slug: slug,
                })
                .then(response => {
                    if (response.data.status == true) {
                        commit('activTourMutation', slug)
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
        deleteTour({commit}, slug) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.delete('/tours/'+slug)
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
        deleteImage({commit}, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.delete('/images/'+data.image_id)
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
        attach({commit}, id) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
                axios.post('/attach_trips', {
                    id: id.toString(),   
                })
                .then(response => {
                    if (response.data.status == true) {
                        commit('attachMutation', id)
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
        delAudit({commit}, auditId) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.delete('/audits/'+auditId)
                .then(response => {
                    if (response.data.status == true) {
                        commit('delAuditMutation', response.data.result)
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
    },
    mutations: {
        updateToursMutation(state, tours) {
            state.tours = tours;
            //state.countTours = tours.length
        },
        publicTourMutation(state, slug) {
            var puclicByAttr = function(arr= state.tours, attr='slug', value=slug){
                var i = arr.length;
                while(i--){
                    if( arr[i] && arr[i].hasOwnProperty(attr) && (arguments.length > 2 && arr[i][attr] === value ) ){
                       arr[i]['public'] = !arr[i]['public']
                    }
                }
                return arr;
            }
            puclicByAttr
        },
        activToursMutation(state, slug) {
            var puclicByAttr = function(arr= state.tours, attr='slug', value=slug){
                var i = arr.length;
                while(i--){
                    if( arr[i] && arr[i].hasOwnProperty(attr) && (arguments.length > 2 && arr[i][attr] === value ) ){
                       arr[i]['activ'] = !arr[i]['activ']
                    }
                }
                return arr;
            }
            puclicByAttr
        },
        deleteTourMutation(state, slug) {
            var removeByAttr = function(arr= state.tours, attr='slug', value=slug){
                var i = arr.length;
                while(i--){
                    if( arr[i] && arr[i].hasOwnProperty(attr) && (arguments.length > 2 && arr[i][attr] === value ) ){
                        arr.splice(i,1);
                    }
                }
                return arr;
            }
            removeByAttr
            
        },
        updateTourMutation(state, tour) {
            state.tour = tour
        },
        updatePageCount(state, pageCount) {
            state.pageCount = pageCount
        },
        updateCountTours(state, countTours) {
            state.countTours = countTours
        },
        attachMutation(state, value) {
            var arr = state.tour.trips
            var attr = 'id'
            var i = arr.length;
            
            while(i--){
                if(arr[i] && arr[i][attr] === value) {
                    arr[i]['freeCount']--;
                }
            }
        },
        setMapInstance(state, mapInstance) {
            state.mapInstance = mapInstance
        },
        getAuditsByTourIdMutation(state, audits) {
            state.audits = audits
        },
        addAuditMutation(state, audit) {
            state.audits.push(audit)
        }
    },
    state: {
        tours: [],
        tour: {},
        audits: [],
        prices: {
            0:{
                'min': '1',
                'max': '499'
            },
            1: {
                'min': '500',
                'max': '999'
            },
            2: {
                'min': '1000',
                'max': '1499'
            },
            '3': {
                'min': '1500',
                'max': '2999'
            },
            4: {
                'min': '3000',
                'max': ''
            },
        },
        pageCount: 0,
		countTours: 0,
        mapInstance: null,
    },
    getters: {
        allTours(state) {
            return state.tours
        },
        oneTour(state) {
            return state.tour
        },
        allAudits(state) {
            return state.audits
        },
        getPrices(state) {
            return state.prices
        },
        getPageCount(state) {
            return state.pageCount
        },
        getCountTours(state) {
            return state.countTours
        }
        
    }
}