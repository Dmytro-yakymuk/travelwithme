<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            <div class="row">
                <div class="col-md-12">

                    <router-link
                        :to="{name: 'adminCreateEvent'}">
                            <a class="btn btn-primary btn-color" href="#"> Додати подію</a>
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
                                            <th class="text-center">Кількість турів з подією</th>
                                            <th class="text-center">Видалити</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        
                                        <tr v-for="event in allEvents" :key="event.id" class="gradeX" >
                                            <td>
                                                {{ event.text }}
                                            </td>
                                            <td class="text-center">
                                                {{ event.count_tours }}
                                            </td>
                                            
                                            <td class="text-center">
                                                <a class="btn btn-primary btn-action" @click="goTo(event.id)"><i class="fa fa-edit"></i></a>
                                                <a class="btn btn-primary btn-action" @click="goDelete(event.id)"><i class="fa fa-trash"></i></a>
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
        name: "AdminEvents",
        
        computed: mapGetters(['allEvents', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getAllEvents', 'deleteEvent']),
            goTo(id) {
                this.$router.push({ name: 'adminUpdateEvent', params: {id: id} })
            },
            goDelete(id) {
                this.deleteEvent(id);
            },
        },
        async mounted() {
            this.getAllEvents();
        }
    }
</script>

<style scoped></style>