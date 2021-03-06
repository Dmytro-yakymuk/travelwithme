import axios from 'axios'
export default {
    actions: {
        async getAllCategories({commit}, limit) {
            const res = await axios
                .get('/categories?limit=' + limit);

            const categories = res.data;
            commit('updateCategories', categories)
        }
    },
    mutations: {
        updateCategories(state, categories) {
            state.categories = categories
        }
    },
    state: {
        categories: []
    },
    getters: {
        allCategories(state) {
            return state.categories
        }
    }
}