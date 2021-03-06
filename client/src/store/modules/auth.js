import axios from 'axios'
export default {
    actions: {
        register(context, data) {
            return new Promise((resolve, reject) => {
              axios.post('/singUp', {
                name: data.name,
                email: data.email,
                password: data.password,
                role: data.role,
              })
                .then(response => {
                  resolve(response)
                })
                .catch(error => {
                  reject(error)
                })
            })
        },
        retrieveToken({commit}, credentials) {

            return new Promise((resolve, reject) => { // оболочка яка перерве виконання, якщо є помилка, до того як відбудиться редірект в Login.vue
                axios.post('/singIn', {
                    email: credentials.email,
                    password: credentials.password,
                })
                .then(response => {
                    const token = response.data.token
                    const id = response.data.id
                    const role = response.data.role

                    localStorage.setItem('token', token)
                    localStorage.setItem('id', id)
                    localStorage.setItem('role', role)
                    commit('retrieveToken', token)
                    resolve(response)
                })
                .catch(error => {
                    console.log(error)
                    reject(error)
                })
            })
        },
        destroyToken(context) {
            // axios.defaults.headers.common['Authorization'] = 'Bearer ' + context.state.token
      
            if (context.getters.loggedIn) {
                localStorage.removeItem('token')
                localStorage.removeItem('id')
                localStorage.removeItem('role')
                context.commit('destroyToken')
            }
        },
    },
    mutations: {
        retrieveToken(state, token) {   
            state.token = token
        },
        destroyToken(state) {
            state.token = null
        }
    },
    state: {
        token: localStorage.getItem('token') || null
    },
    getters: {
        loggedIn(state) {
            return state.token !== null
        },
    }
}