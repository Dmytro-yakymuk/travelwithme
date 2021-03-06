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
            <form class="twocol-form" action="#" @submit.prevent="login">
                <div class="row">
                    <div class="col-md-6">

                        <div class="form-holder">
                            <div class="wrap">
                                <div class="hold">
                                    <label for="email">Email</label>
                                    <input v-model="email" type="text" id="email" class="form-control" required="">
                                </div>
                                <div class="hold">
                                    <label for="password">Пароль</label>
                                    <input v-model="password" type="password" id="password" class="form-control" required="">
                                </div>
                                <div class="btn-hold">
                                    <button type="submit" class="btn btn-default">Увійти</button>
                                </div>
                            </div>
                        </div>
                    </div>
                    
                </div>
            </form>
        </div>
    </main>
</template>

<script>

export default {
    name: "login",
    data() {
        return {
            email: '',
            password: '',
        }
    },
    methods: {
        login() {
           this.$store.dispatch('retrieveToken', {
               email: this.email,
               password: this.password
           })
           .then(response => {
               this.$router.push({ name: 'admin' })
           })
        }  
    }
}
</script>

<style scoped>


</style>