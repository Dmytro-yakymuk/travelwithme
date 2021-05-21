<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            <div class="row">
                <div class="col-md-12">

                    <router-link
                        :to="{name: 'adminCreateRegion'}">
                            <a class="btn btn-primary btn-color" href="#"> Додати регіон</a>
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
                                            <th class="text-center">Кількість турів в категорії</th>
                                            <th class="text-center">Видалити</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        
                                        <tr v-for="region in allRegions" :key="region.id" class="gradeX" >
                                            <td>
                                                {{ region.name }}
                                            </td>
                                            <td class="text-center">
                                                {{ region.count_tours }}
                                            </td>
                                            
                                            <td class="text-center">
                                                <a class="btn btn-primary btn-action" @click="goTo(region.slug)"><i class="fa fa-edit"></i></a>
                                                <a class="btn btn-primary btn-action" @click="goDelete(region.slug)"><i class="fa fa-trash"></i></a>
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
        name: "AdminRegions",
        
        computed: mapGetters(['allRegions', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getAllRegions', 'deleteRegion']),
            goTo(slug) {
                this.$router.push({ name: 'adminUpdateRegion', params: {slug: slug} })
            },
            goDelete(slug) {
                this.deleteRegion(slug);
            },
        },
        async mounted() {
            this.getAllRegions();
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