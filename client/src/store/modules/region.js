import axios from 'axios'
export default {
    actions: {
        async getAllRegions({commit}, limit) {
            const res = await axios
                .get('/regions?limit=' + limit);

            const regions = res.data;
            commit('updateRegions', regions)
        }
    },
    mutations: {
        updateRegions(state, regions) {
            state.regions = regions
        }
    },
    state: {
        regions: []
    },
    getters: {
        allRegions(state) {
            return state.regions
        }
    }
}