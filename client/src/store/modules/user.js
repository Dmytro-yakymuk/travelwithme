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

            const users = res.data;
            commit('updateUsers', users)
        }
    },
    mutations: {
        updateUsers(state, users) {
            state.users = users
        }
    },
    state: {
        users: []
    },
    getters: {
        allUsers(state) {
            return state.users
        }
    }
}