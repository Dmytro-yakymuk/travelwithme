<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            <div class="row">
                <div class="col-md-12">

                    <router-link
                        :to="{ name: 'adminCreateTour'}">
                            <a href="#"> Додати тур</a>
                    </router-link> 
                    
                    <div class="panel panel-default">
                        <div class="panel-heading"></div>
                        <div class="panel-body">
                            <div class="table-responsive">
                                <table class="table table-striped table-bordered table-hover" id="dataTables-example">
                                    <thead>
                                        <tr>
                                            <th>Назва</th>
                                            <th>Фото</th>
                                            <th>Статус</th>
                                            <th>Видалити</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        
                                        <tr v-for="tour in allTours" :key="tour.id" class="gradeX" >
                                            <td>
                                                <a href="#" @click="goTo(tour.slug)">{{ tour.title }}</a>
                                            </td>
                                            <td>
                                                <div class="img_upload">
                                                    <img 
                                                        v-for="image in tour.images.slice(0,1)"
                                                        :key="image.id"
                                                        :src="imageURL+image.name" 
                                                        alt="#">
                                                </div>
                                            </td>
                                            <td>{{ tour.status }}</td>
                                            <td class="center">
                                                <a @click="goDelete(tour.slug)" href="#"> X</a>
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
            }
        },
        
        computed: mapGetters(['allTours']),
        methods: {
            ...mapActions(['getAllTours']),
            goTo(slug) {
                this.$router.push({ name: 'adminUpdateTour', params: {slug: slug} })
            },
            goDelete(slug) {
                this.$router.push({ name: 'adminDeleteTour', params: {slug: slug} })
            },
        },
        async mounted() {
            this.getAllTours({
                'where': "user_id="+localStorage.getItem('id'),
                'order_by': "id desc",
                'limit': null
            });
        }
        
    }
</script>

<style scoped>
    .img_upload {
        height: 150px !important;
        width: 125px !important;
    }
</style>