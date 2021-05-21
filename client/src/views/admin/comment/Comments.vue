<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            <div class="row">
                <div class="col-md-12">

                    <router-link
                        v-if="role == 'client'"
                        :to="{name: 'adminCreateComment'}">
                            <a class="btn btn-primary btn-color"> Додати комент</a>
                    </router-link> 
                    
                    <div v-if="getSuccessMessage != ''" class="alert alert-success" role="alert">{{getSuccessMessage}}</div>
                    <div v-if="getFailMessage != ''" class="alert alert-danger" role="alert">{{getFailMessage}}</div>

                    <div class="panel panel-default">
                        <div class="panel-heading"></div>
                        <div class="panel-body">
                            <div class="table-responsive">
                                <table class="table table-striped table-bordered table-hover" id="dataTables-example">
                                    <thead>
                                        <tr>
                                            <th class="text-center">Комент</th>
                                            <th class="text-center">Зірки</th>
                                            <th class="text-center">Тур</th>
                                            <th v-if="role != 'client'" class="text-center">Юзер</th>
                                            <th  class="text-center">Дата коментування</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        
                                        <tr v-for="comment in allComments" :key="comment.id" class="gradeX" >
                                            <td>
                                                {{ comment.message }}
                                            </td>
                                            <td class="text-center">
                                                {{ comment.star }}
                                            </td>
                                            <td>
                                                {{ comment.tour.title }}
                                            </td>
                                            <td v-if="role != 'client'" class="text-center">
                                                {{ comment.user.name }}
                                            </td>
                                            <td class="text-center">
                                                {{ format_date(comment.created_at, 'DD MMM YYYY hh:ss') }}
                                            </td>
                                            <td class="text-center">
                                                <a class="btn btn-primary btn-action" @click="goView(comment.id)"> <i class="fa fa-book"></i></a>
                                                <a v-if="role == 'client'" class="btn btn-primary btn-action" @click="goEdit(comment.id)"> <i class="fa fa-edit"></i></a>
                                                <a v-if="role == 'client'" class="btn btn-primary btn-action" @click="goDelete(comment.id)"> <i class="fa fa-trash"></i></a>
                                            </td>
                                        </tr>

                                    </tbody>
                                </table>
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
    import moment from 'moment'

    export default {
        name: "AdminComments",
        data: function () {
            return {
                role: localStorage.getItem('role'),
            }
        },
        computed: mapGetters(['allComments', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getAllComments', 'deleteComment']),
            goView(id) {
                this.$router.push({ name: 'adminViewComment', params: {id: id} })
            },
            goEdit(id) {
                this.$router.push({ name: 'adminUpdateComment', params: {id: id} })
            },
            goDelete(id) {
                this.deleteComment(id);
            },
            format_date(value, format='DD MMM YYYY'){
                if (value) {
                    return moment(String(value)).format(format); 
                }
            },
        },
        async mounted() {
            this.getAllComments({
                'user_id': localStorage.getItem('id')
            });
        }
    }
</script>

<style scoped>
    .btn-color {
        background-color: #6b6957!important;
        color: blanchedalmond;
    }
    .btn-delete {
        min-width: 40px!important;
    }
</style>