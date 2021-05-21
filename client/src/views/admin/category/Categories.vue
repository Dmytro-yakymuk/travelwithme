<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            <div class="row">
                <div class="col-md-12">

                    <router-link
                        :to="{name: 'adminCreateCategory'}">
                            <a class="btn btn-primary btn-color" href="#"> Додати категорію</a>
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
                                            <th class="text-center">Назва</th>
                                            <th class="text-center">Іконка</th>
                                            <th class="text-center">Кількість турів в категорії</th>
                                            <th class="text-center">Видалити</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        
                                        <tr v-for="category in allCategories" :key="category.id" class="gradeX" >
                                            <td>
                                                {{ category.name }}
                                            </td>
                                            <td>
                                                {{ category.icon }}
                                            </td>
                                            <td class="text-center">
                                                {{ category.count_tours }}
                                            </td>
                                            
                                            <td class="text-center">
                                                <a class="btn btn-primary btn-action" @click="goTo(category.slug)"> <i class="fa fa-edit"></i></a>
                                                <a class="btn btn-primary btn-action" @click="goDelete(category.slug)"> <i class="fa fa-trash"></i></a>
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

    export default {
        name: "AdminCategories",
        
        computed: mapGetters(['allCategories', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getAllCategories', 'deleteCategory']),
            goTo(slug) {
                this.$router.push({ name: 'adminUpdateCategory', params: {slug: slug} })
            },
            goDelete(slug) {
                this.deleteCategory(slug);
            },
        },
        async mounted() {
            this.getAllCategories();
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