<template>
    <!-- main container-->
    <main id="main">
        <!-- top information area -->
        <div class="inner-top">
            <div class="container">
                <h1 class="inner-main-heading">Login</h1>
                <!-- breadcrumb -->
                <nav class="breadcrumbs">
                    <ul>
                        <li><a href="#">Home</a></li>
                        <li><a href="#">pages</a></li>
                        <li><span>Login</span></li>
                    </ul>
                </nav>
            </div>
        </div>
        <div class="inner-main common-spacing container">
            <!-- form -->
            <form class="twocol-form" action="#" @submit.prevent="register">
                <div  style="margin: 0 280px;width: 600px">
                    <div class="top-box">
                        <span class="holder height">Реєстрація</span>
                    </div>
                    <div class="form-holder">
                        <div class="wrap">
                            <div class="hold">
                                <label for="name">Ім'я</label>
                                <input type="text" id="name" v-model="name" class="form-control">
                            </div>
                            <div class="hold">
                                <label for="email">Email</label>
                                <input type="email" id="email" v-model="email" class="form-control">
                            </div>
                            <div class="hold">
                                <label for="password">Пароль</label>
                                <input type="password" id="password" v-model="password" class="form-control">
                            </div>

                            <div class="hold">
                                <label for="role">Тип профілю</label>
                                <select v-model="role" id="role" class="trip dark">
                                    <option :value="'owner'">Я пропоную тури</option>
                                    <option :value="'client'">Я шукаю тури</option>
                                </select>
                            </div>

                            <div class="btn-hold">
                                <button type="submit" class="btn btn-default">Зареєтруватись</button>
                            </div>
                        </div>
                    </div>
                    Акаунт вже створений? <a href="/login" data-toggle="dropdown">
                        <span class="text">Увійти</span>
                    </a>
                </div>
            </form>
        </div>
       
    </main>
</template>

<script>
export default {
    name: "register",
    data() {
        return {
            name: '',
            email: '',
            password: '',
            role: ''
        }
    },
    methods: {
        register() {
           this.$store.dispatch('register', {
               name: this.name,
               email: this.email,
               password: this.password,
               role: this.role
           })
           .then(response => {
               this.$router.push({ name: 'login' })
           })
        }  
    }
}
</script>

<style scoped>


</style>