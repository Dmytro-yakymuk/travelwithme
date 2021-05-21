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
                                    <form role="form" @submit.prevent="submitForm">
                                        
                                        <div class="form-group">
                                            <label>Назва</label>
                                            <input type="text" v-model="name" class="form-control">
                                        </div>
                                         <div class="form-group">
                                            <label>Іконка</label>
                                            <input type="text" v-model="icon" class="form-control">
                                        </div>
  
                                        <button type="submit" class="btn btn-default">Зберегти</button>
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
    import {mapGetters, mapActions} from "vuex"

    export default {
        name: "AdminCreateCategory",
        data: function () {
            return {
                name: 'test_name',
                icon: 'test_icon'
            }
        },
        computed: mapGetters(['getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['createCategory']),

            submitForm() {
                this.createCategory({
                    'name': this.name,
                    'icon': this.icon
                })
                .then(response => {})
                .catch(error => {
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                })  
            },
        },
    }
</script>

<style scoped></style>