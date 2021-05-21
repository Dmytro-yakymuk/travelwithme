<template>
    <div id="page-wrapper" >
        <div id="page-inner">
             
            <div v-if="getSuccessMessage != ''" class="alert alert-success" role="alert">{{getSuccessMessage}}</div>
            <div v-if="getFailMessage != ''" class="alert alert-danger" role="alert">{{getFailMessage}}</div>

            <div class="row">
                <div class="col-lg-12">
                    <div class="panel panel-default">
                        <div class="panel-body">
                            <div class="row">
                                <div class="col-lg-6">
                                    <form role="form" @submit.prevent="updateProfile">
                                        
                                        <div class="form-group">
                                            <label>Прізвище</label>
                                            <input v-model="oneUser.surname" class="form-control" placeholder="Введіть прізвище">
                                        </div>
                                        <div class="form-group">
                                            <label>Ім'я</label>
                                            <input v-model="oneUser.name" class="form-control" placeholder="Введіть ім'я">
                                        </div>
                                        <div class="form-group">
                                            <label>По батькові</label>
                                            <input v-model="oneUser.patronymic" class="form-control" placeholder="Введіть по батькові">
                                        </div>
                                        <div class="form-group">
                                            <label>Email</label>
                                            <input v-model="oneUser.email" class="form-control" placeholder="Введіть email">
                                        </div>
                                        <div class="form-group">
                                            <label>Старий пароль</label>
                                            <input v-model="old_password" class="form-control" placeholder="Введіть старий пароль">
                                        </div>
                                        <div class="form-group">
                                            <label>Новий пароль</label>
                                            <input v-model="new_password" class="form-control" placeholder="Введіть новий пароль">
                                        </div>
                                        <div class="form-group">
                                            <label>Номер телефону</label>
                                            <input v-model="oneUser.phone" class="form-control" placeholder="Введіть номер телефону">
                                        </div>
              
                                        <button type="submit" class="btn btn-default">Редагувати</button>
                                    </form>

                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapGetters, mapActions, mapMutations} from "vuex"

    export default {
        name: "AdminUpdateProfile",
        data: function () {
            return {
                old_password: '123456',
                new_password: '123456',
            }
        },
        computed: {
            ...mapGetters(['oneUser', 'getMessage', 'getSuccessMessage', 'getFailMessage'])
        },
        methods: {
            ...mapActions(['getOneUser']),
            ...mapMutations(['updateMessage']),

            updateProfile() {
                this.$store.dispatch('updateUser', {
                    'id': this.oneUser.id,
                    'surname': this.oneUser.surname,
                    'name': this.oneUser.name,
                    'patronymic': this.oneUser.patronymic,
                    'email': this.oneUser.email,
                    'old_password': this.old_password,
                    'new_password': this.new_password,
                    'phone': this.oneUser.phone
                   
                })
                setTimeout(() => {
                    this.updateMessage('')
                }, 3000)
                
            },
        },
        async mounted() {
            this.getOneUser();
        }
        
    }
</script>

<style scoped>
</style>