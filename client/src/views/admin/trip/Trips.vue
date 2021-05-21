<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            <div class="row">
                <div class="col-md-12">

                    <router-link
                        v-if="role != 'admin'"
                        :to="{name: 'adminCreateTrip'}">
                            <a class="btn btn-primary btn-color" href="#"> Додати поїздку</a>
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
                                            <th class="text-center">Дати</th>
                                            <th class="text-center">Ціна</th>
                                            <th class="text-center">Кількість</th>
                                            <th class="text-center">Для туру</th>
                                            <th class="text-center">Видалити</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        
                                        <tr v-for="trip in allTrips" :key="trip.id" class="gradeX" >
                                            <td class="text-center">
                                                {{ format_date(trip.start, 'DD MMM') }} - {{ format_date(trip.end, 'DD MMM YYYY') }}
                                            </td>
                                            <td class="text-center">
                                                {{ trip.price }}
                                            </td>
                                            <td class="text-center">
                                                {{ trip.count }}
                                            </td>
                                            <td class="text-center">
                                                {{ trip.tour_id }}
                                            </td>
                                            
                                            <td class="text-center">
                                                <a class="btn btn-primary btn-action" @click="goTo(trip.id)"><i class="fa fa-edit"></i></a>
                                                <a class="btn btn-primary btn-action" @click="goDelete(trip.id)"><i class="fa fa-trash"></i></a>
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
        name: "AdminTrips",
        data: function () {
            return {
                role: localStorage.getItem('role'),
            }
        },
        
        computed: mapGetters(['allTrips', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getAllTrips', 'deleteTrip']),
            goTo(id) {
                this.$router.push({ name: 'adminUpdateTrip', params: {id: id} })
            },
            goDelete(id) {
                this.deleteTrip(id);
            },
            format_date(value, format='DD MMM YYYY'){
                if (value) {
                    return moment(String(value)).format(format); 
                }
            },
        },
        async mounted() {
            this.getAllTrips();
        }
        
    }
</script>

<style scoped>

</style>