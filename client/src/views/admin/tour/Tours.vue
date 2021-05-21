<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            <div class="row">
                <div class="col-md-12">

                    <router-link
                        v-if="role != 'admin'"
                        :to="{ name: 'adminCreateTour'}">
                            <a class="btn btn-primary btn-color" href="#"> Додати тур</a>
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
                                            <th class="text-center">Фото</th>
                                            <th class="text-center" v-if="role == 'admin'">Публікація</th>
                                            <th class="text-center" v-if="role == 'admin'">Власник</th>
                                            <th class="text-center" v-if="role == 'owner'">Активність</th>
                                            <th class="text-center"></th>
                                        </tr>
                                    </thead>
                                    <tbody class="tr-height">
                                        <tr v-for="tour in allTours" :key="tour.id" class="gradeX" >
                                            <td>{{ tour.title }}</td>
                                            <td class="text-center">
                                                <div class="img_upload">
                                                    <img 
                                                        v-for="image in tour.images.slice(0,1)"
                                                        :key="image.id"
                                                        :src="imageURL+image.name" 
                                                        alt="#">
                                                </div>
                                            </td>
                                            <td class="text-center" v-if="role == 'admin'"><input @change="goPublic(tour.slug)" type="checkbox" class="form-check-input" v-model="tour.public"></td>
                                            <td class="text-center" v-if="role == 'admin'"><p>{{ tour.user.name }}</p></td>
                                            <td class="text-center" v-if="role == 'owner'"><input type="checkbox" class="form-check-input" @change="goActiv(tour.slug)" v-model="tour.activ"></td>
                                            <td class="text-center">
                                                <a class="btn btn-primary btn-action" @click="goTo(tour.slug)"> <i class="fa fa-edit"></i></a>
                                                <a class="btn btn-primary btn-action" @click="goDelete(tour.slug)"> <i class="fa fa-trash"></i></a>
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
    import {ShowImageURL} from '../../../constants/index'
    export default {
        name: "AdminTours",
        data: function () {
            return {
                imageURL: ShowImageURL,
                role: localStorage.getItem('role'),
            }
        },
        
        computed: mapGetters(['allTours', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getAllTours', 'publicTour', 'activTour']),
            goTo(slug) {
                this.$router.push({ name: 'adminUpdateTour', params: {slug: slug} })
            },
            goPublic(slug) {
                this.publicTour(slug);
            },
            goActiv(slug) {
                this.activTour(slug);
            },
            goDelete(slug) {
                this.$router.push({ name: 'adminDeleteTour', params: {slug: slug} })
            },
        },
        async mounted() {
            this.getAllTours({
                'user_id': localStorage.getItem('id'),
                'order_by': "id desc"
            });
        },
		// watch: {
        //     category_id: function () {

        //     }
        // }
        
    }
</script>

<style scoped>
    .img_upload {
        height: 150px !important;
        width: 225px !important;
    }
    .img_upload img {
        height: 150px !important;
        width: 225px !important;
    }
</style>