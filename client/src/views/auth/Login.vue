<template>
    <!-- main container-->
    <main id="main">
        <!-- top information area -->
        <div class="inner-top">
            <div class="container">
                <h1 class="inner-main-heading">Вхід</h1>
                <!-- breadcrumb -->
                <nav class="breadcrumbs">
                    <ul>
                        <li><a href="#">Головна</a></li>
                        <li><span>Вхід</span></li>
                    </ul>
                </nav>
            </div>
        </div>
        
        <div class="inner-main common-spacing container">
            <!-- form -->
            <form class="twocol-form" action="#" @submit.prevent="login">
                <div class="row d-flex justify-content-center">
                    <div style="margin: 0 280px;width: 600px">

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
                        Немає акаунта? <a href="/register" data-toggle="dropdown">
                            <span class="text hidden-md hidden-lg">Зареєструватись</span>
                        </a>
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
                var numberOfEntries = window.history.length;
                if (numberOfEntries>2) {
                    this.$router.go(-2)
                }
                this.$router.push({ name: 'home' });
            })
        }  
    }
}
</script>

<style scoped>


</style>