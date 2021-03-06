import axios from 'axios'
export default {
    actions: {
        async getAllTours({commit}, map) {

            let query = '/tours?'
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

            const tours = res.data;
            commit('updateToursMutation', tours)
        },
        async getOneTour({commit}, slug) {
            const res = await axios
                .get('/tours/' + slug);

            const tour = res.data;
            commit('updateTourMutation', tour)
        },
        createTour(context, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.post('/tours', {
                    title: data.title,
                    description: data.description,
                    status: data.status,
                    category_id: data.category_id,
                    region_id: data.region_id
                })
                .then(response => {
                    resolve(response)
                })
                .catch(error => {
                    reject(error)
                })
            })
        },
        updateTour(context, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.put('/tours/'+data.slug, {
                    title: data.title,
                    slug: data.slug,
                    description: data.description,
                    status: data.status,
                    category_id: data.category_id,
                    region_id: data.region_id
                })
                .then(response => {
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
                    commit('deleteTourMutation', slug)
                    resolve(response)
                })
                .catch(error => {
                    reject(error)
                })
            })
        },
        deleteImage(context, data) {
            return new Promise((resolve, reject) => {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

                axios.delete('/images/'+data.image_id)
                .then(response => {
                    resolve(response)
                })
                .catch(error => {
                    reject(error)
                })
            })
        },
        goOrder(context, id) {
            return new Promise((resolve, reject) => {
         
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
                axios.post('/orders', {
                    id: id, 
                    
                })
                .then(response => {
                    resolve(response)
                })
                .catch(error => {
                    reject(error)
                })
            })
        }
    },
    mutations: {
        updateToursMutation(state, tours) {
            state.tours = tours
        },
        deleteTourMutation(state, slug) {
            var removeByAttr = function(arr= state.tours, attr='slug', value=slug){
                var i = arr.length;
                while(i--){
                   if( arr[i] 
                       && arr[i].hasOwnProperty(attr) 
                       && (arguments.length > 2 && arr[i][attr] === value ) ){
                       arr.splice(i,1);
            
                   }
                }
                return arr;
            }
            removeByAttr
            
        },
        updateTourMutation(state, tour) {
            state.tour = tour
        }
    },
    state: {
        tours: [],
        tour: {}
    },
    getters: {
        allTours(state) {
            state.tours.map(t => t['minPrice'] = Math.min(...t.trips.map(i=>i.price)))
            return state.tours
        },
        oneTour(state) {
            var price = [];
            for (const color in state.tour.trips){
                price.push(state.tour.trips[color].price);
            }
            state.tour['minPrice'] = Math.min(...price)
            return state.tour
        }
    }
}