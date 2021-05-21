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
                                    <form role="form" @submit.prevent="startProcessingQueue">
                                        
                                       <div class="form-group">
                                            <label>Початок</label>
                                            <input type="date" v-model="oneTrip.start" class="form-control"
                                                min="2021-01-01" max="2022-01-01">
                                        </div>

                                        <div class="form-group">
                                            <label>Кінець</label>
                                            <input type="date" v-model="oneTrip.end" class="form-control"
                                                min="2021-01-01" max="2022-01-01">
                                        </div>

                                        <div class="form-group">
                                            <label>Ціна</label>
                                            <input type="number" v-model.number="oneTrip.price" min="0">
                                        </div>

                                        <div class="form-group">
                                            <label>Кількість</label>
                                            <input type="number" v-model.number="oneTrip.count" min="0">
                                        </div>
 

                                        <div class="form-group">
                                            <label>Тур</label>
                                            <select v-model="oneTrip.tour_id"  class="form-control">
                                                <option 
                                                    v-for="tour in allTours" :key="tour.id"
                                                    :value="tour.id"
                                                    :selected="tour.id == oneTrip.tour_id"
                                                >{{tour.title}}</option>
                                            </select>
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
        name: "AdminUpdateTrip",
        props: ['id'],
        computed: mapGetters(['oneTrip', 'allTours', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getOneTrip', 'getAllTours']),

            startProcessingQueue() {
                this.$store.dispatch('updateTrip', {
                    'id': this.oneTrip.id,
                    'start': this.format_date_unix(this.oneTrip.start),
                    'end': this.format_date_unix(this.oneTrip.end),
                    'price': this.oneTrip.price,
                    'count': this.oneTrip.count,
                    'tour_id': this.oneTrip.tour_id
                });
                
            },
            format_date_unix(value){
                var datum = Date.parse(value);
                return datum/1000;
			}
           
        },
        async mounted() {
            this.getOneTrip({
                'id': this.id
            });
            this.getAllTours({
                'user_id': localStorage.getItem('id'),
                'order_by': "id desc"
            });
        },
        
    }
</script>

<style scoped></style>